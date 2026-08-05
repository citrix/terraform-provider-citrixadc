package ntpsync

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ntp"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NtpsyncResource{}
var _ resource.ResourceWithConfigure = (*NtpsyncResource)(nil)
var _ resource.ResourceWithImportState = (*NtpsyncResource)(nil)

// ntpsyncStaticId is the ID assigned to newly-created ntpsync resources.
//
// ntpsync is a singleton without any unique key attribute; NITRO reads it via
// GET (all) with an empty name and never keys off the ID. The SDK v2 resource
// assigned a random "tf-ntpsync-<n>" ID; the Framework Read preserves whatever
// ID is already in state, so migrated SDK v2 state keeps its original ID while
// newly-created resources get this stable singleton ID (no churn either way).
const ntpsyncStaticId = "ntpsync-config"

func NewNtpsyncResource() resource.Resource {
	return &NtpsyncResource{}
}

// NtpsyncResource defines the resource implementation.
type NtpsyncResource struct {
	client *service.NitroClient
}

func (r *NtpsyncResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NtpsyncResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ntpsync"
}

func (r *NtpsyncResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// doNtpsyncChange applies the desired state via the NITRO enable/disable
// actions. This mirrors the SDK v2 doNtpsyncChange helper exactly: ENABLED ->
// action=enable, DISABLED -> action=disable, anything else is an error at
// apply time (SDK v2 did not enforce the enum at the schema level).
func (r *NtpsyncResource) doNtpsyncChange(newstate string) error {
	ntpsync := ntp.Ntpsync{}
	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Ntpsync.Type(), &ntpsync, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Ntpsync.Type(), &ntpsync, "disable")
	default:
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

func (r *NtpsyncResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NtpsyncResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ntpsync resource")

	// Action-only resource: apply state via enable/disable action.
	if err := r.doNtpsyncChange(data.State.ValueString()); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ntpsync, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created ntpsync resource")

	// Singleton ID (no unique key attribute).
	data.Id = types.StringValue(ntpsyncStaticId)

	// Read the updated state back
	if !r.readNtpsyncFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ntpsync not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NtpsyncResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NtpsyncResourceModel

	// Read Terraform prior state data into the model (preserves ID)
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ntpsync resource")

	found := r.readNtpsyncFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NtpsyncResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NtpsyncResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating ntpsync resource")

	// SDK v2 parity: only re-apply the action when "state" changed.
	if !data.State.Equal(state.State) {
		tflog.Debug(ctx, "state has changed for ntpsync, starting update")
		if err := r.doNtpsyncChange(data.State.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ntpsync, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated ntpsync resource")
	} else {
		tflog.Debug(ctx, "No changes detected for ntpsync resource, skipping update")
	}

	// Read the updated state back
	if !r.readNtpsyncFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "ntpsync not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NtpsyncResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// ntpsync does not support a DELETE operation (NITRO exposes only
	// enable/disable/get). SDK v2 delete was a no-op that just cleared the ID;
	// the Framework removes the resource from state automatically on return.
	tflog.Debug(ctx, "Deleting ntpsync resource (state-only removal)")
}

// readNtpsyncFromApi reads the ntpsync singleton via NITRO GET (all).
func (r *NtpsyncResource) readNtpsyncFromApi(ctx context.Context, data *NtpsyncResourceModel, diags *diag.Diagnostics) bool {
	// Singleton read: FindResource with an empty name (matches SDK v2).
	getResponseData, err := r.client.FindResource(service.Ntpsync.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ntpsync, got error: %s", err))
		return false
	}

	ntpsyncSetAttrFromGet(ctx, data, getResponseData)

	return true
}
