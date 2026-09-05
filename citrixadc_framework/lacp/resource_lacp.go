package lacp

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
var _ resource.Resource = &LacpResource{}
var _ resource.ResourceWithConfigure = (*LacpResource)(nil)
var _ resource.ResourceWithImportState = (*LacpResource)(nil)

func NewLacpResource() resource.Resource {
	return &LacpResource{}
}

// LacpResource defines the resource implementation.
type LacpResource struct {
	client *service.NitroClient
}

func (r *LacpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LacpResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lacp"
}

func (r *LacpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LacpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LacpResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lacp resource")

	// Build the NITRO payload from the plan.
	lacp := lacpGetThePayloadFromtheConfig(ctx, &data)

	// lacp is a singleton-style global config (NITRO exposes only update/get):
	// use UpdateUnnamedResource, matching the SDK v2 resource.
	err := r.client.UpdateUnnamedResource(service.Lacp.Type(), &lacp)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lacp, got error: %s", err))
		return
	}

	// ID is the ownernode value (matches SDK v2 d.SetId(strconv.Itoa(ownernode))).
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Ownernode.ValueInt64()))

	tflog.Trace(ctx, "Created lacp resource")

	// Read the updated state back
	if !r.readLacpFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lacp not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LacpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LacpResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lacp resource")

	found := r.readLacpFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LacpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LacpResourceModel

	// Read Terraform prior state to detect changes / preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (recomputed below in case ownernode changed).
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lacp resource")

	// Detect changes on the writable attributes (syspriority is the updatable
	// one in SDK v2; ownernode selects the target node).
	hasChange := false
	if !data.Syspriority.Equal(state.Syspriority) {
		tflog.Debug(ctx, "syspriority has changed for lacp")
		hasChange = true
	}
	if !data.Ownernode.Equal(state.Ownernode) {
		tflog.Debug(ctx, "ownernode has changed for lacp")
		hasChange = true
	}

	if hasChange {
		lacp := lacpGetThePayloadFromtheConfig(ctx, &data)
		err := r.client.UpdateUnnamedResource(service.Lacp.Type(), &lacp)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lacp, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated lacp resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lacp resource, skipping update")
	}

	// Keep the ID in sync with the (possibly changed) ownernode.
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Ownernode.ValueInt64()))

	// Read the updated state back
	if !r.readLacpFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lacp not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LacpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LacpResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lacp resource")

	// lacp is a global configuration and NITRO exposes no delete operation
	// (matches SDK v2). Just remove it from Terraform state.
	tflog.Trace(ctx, "Deleted lacp resource from state")
}

// Helper function to read lacp data from API.
// Returns false when the resource is not found on the ADC.
func (r *LacpResource) readLacpFromApi(ctx context.Context, data *LacpResourceModel, diags *diag.Diagnostics) bool {
	// ID is the ownernode value; read the lacp entry for that node
	// (matches SDK v2 FindResource(service.Lacp.Type(), strconv.Itoa(ownernode))).
	ownernode := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lacp.Type(), ownernode)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lacp, got error: %s", err))
		return false
	}

	lacpSetAttrFromGet(ctx, data, getResponseData)

	return true
}
