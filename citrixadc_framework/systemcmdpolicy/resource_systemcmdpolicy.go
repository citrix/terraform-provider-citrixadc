package systemcmdpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemcmdpolicyResource{}
var _ resource.ResourceWithConfigure = (*SystemcmdpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*SystemcmdpolicyResource)(nil)

func NewSystemcmdpolicyResource() resource.Resource {
	return &SystemcmdpolicyResource{}
}

// SystemcmdpolicyResource defines the resource implementation.
type SystemcmdpolicyResource struct {
	client *service.NitroClient
}

func (r *SystemcmdpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemcmdpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemcmdpolicy"
}

func (r *SystemcmdpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemcmdpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemcmdpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemcmdpolicy resource")

	systemcmdpolicy := systemcmdpolicyGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	policyname := data.Policyname.ValueString()
	_, err := r.client.AddResource(service.Systemcmdpolicy.Type(), policyname, &systemcmdpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemcmdpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created systemcmdpolicy resource")

	// Set ID for the resource before reading state (single unique attribute - plain value)
	data.Id = types.StringValue(policyname)

	// Read the updated state back
	if !r.readSystemcmdpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "systemcmdpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemcmdpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemcmdpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemcmdpolicy resource")

	found := r.readSystemcmdpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SystemcmdpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemcmdpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemcmdpolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for systemcmdpolicy")
		hasChange = true
	}
	if !data.Cmdspec.Equal(state.Cmdspec) {
		tflog.Debug(ctx, "cmdspec has changed for systemcmdpolicy")
		hasChange = true
	}

	if hasChange {
		systemcmdpolicy := systemcmdpolicyGetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource
		policyname := data.Policyname.ValueString()
		_, err := r.client.UpdateResource(service.Systemcmdpolicy.Type(), policyname, &systemcmdpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemcmdpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated systemcmdpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for systemcmdpolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readSystemcmdpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "systemcmdpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemcmdpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemcmdpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemcmdpolicy resource")

	// Named resource - delete using DeleteResource
	policyname := data.Id.ValueString()
	err := r.client.DeleteResource(service.Systemcmdpolicy.Type(), policyname)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systemcmdpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted systemcmdpolicy resource")
}

// Helper function to read systemcmdpolicy data from API
func (r *SystemcmdpolicyResource) readSystemcmdpolicyFromApi(ctx context.Context, data *SystemcmdpolicyResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value
	policyname := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Systemcmdpolicy.Type(), policyname)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemcmdpolicy, got error: %s", err))
		return false
	}

	systemcmdpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
