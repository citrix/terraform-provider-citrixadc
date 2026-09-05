package snmpcommunity

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
var _ resource.Resource = &SnmpcommunityResource{}
var _ resource.ResourceWithConfigure = (*SnmpcommunityResource)(nil)
var _ resource.ResourceWithImportState = (*SnmpcommunityResource)(nil)

func NewSnmpcommunityResource() resource.Resource {
	return &SnmpcommunityResource{}
}

// SnmpcommunityResource defines the resource implementation.
type SnmpcommunityResource struct {
	client *service.NitroClient
}

func (r *SnmpcommunityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmpcommunityResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpcommunity"
}

func (r *SnmpcommunityResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmpcommunityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmpcommunityResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmpcommunity resource")

	snmpcommunity := snmpcommunityGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	communityname_value := data.Communityname.ValueString()
	_, err := r.client.AddResource(service.Snmpcommunity.Type(), communityname_value, &snmpcommunity)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmpcommunity, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created snmpcommunity resource")

	// Set ID for the resource before reading state (single unique attribute - plain value)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Communityname.ValueString()))

	// Read the updated state back
	if !r.readSnmpcommunityFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmpcommunity not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpcommunityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmpcommunityResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmpcommunity resource")

	found := r.readSnmpcommunityFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SnmpcommunityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SnmpcommunityResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating snmpcommunity resource")

	// snmpcommunity has no NITRO update operation and both attributes
	// (communityname, permissions) are ForceNew/RequiresReplace, so any value
	// change triggers a destroy/recreate rather than reaching Update. This
	// method therefore only reads the current state back.

	// Read the updated state back
	if !r.readSnmpcommunityFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "snmpcommunity not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpcommunityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmpcommunityResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmpcommunity resource")

	// Named resource - delete using DeleteResource
	communityname_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Snmpcommunity.Type(), communityname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete snmpcommunity, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted snmpcommunity resource")
}

// Helper function to read snmpcommunity data from API
func (r *SnmpcommunityResource) readSnmpcommunityFromApi(ctx context.Context, data *SnmpcommunityResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	communityname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Snmpcommunity.Type(), communityname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmpcommunity, got error: %s", err))
		return false
	}

	snmpcommunitySetAttrFromGet(ctx, data, getResponseData)

	return true
}
