package videooptimizationpacingpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VideooptimizationpacingpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationpacingpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationpacingpolicylabelResource)(nil)

func NewVideooptimizationpacingpolicylabelResource() resource.Resource {
	return &VideooptimizationpacingpolicylabelResource{}
}

// VideooptimizationpacingpolicylabelResource defines the resource implementation.
// NOTE: videooptimization pacing functionality is deprecated on the NetScaler
// (NITRO/CLI) side; this resource is retained for compatibility.
type VideooptimizationpacingpolicylabelResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationpacingpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationpacingpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationpacingpolicylabel"
}

func (r *VideooptimizationpacingpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationpacingpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationpacingpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationpacingpolicylabel resource")
	videooptimizationpacingpolicylabel := videooptimizationpacingpolicylabelGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	labelname_value := data.Labelname.ValueString()
	_, err := r.client.AddResource(service.Videooptimizationpacingpolicylabel.Type(), labelname_value, &videooptimizationpacingpolicylabel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationpacingpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created videooptimizationpacingpolicylabel resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	// Read the updated state back
	r.readVideooptimizationpacingpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationpacingpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationpacingpolicylabel resource")

	r.readVideooptimizationpacingpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VideooptimizationpacingpolicylabelResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// NITRO exposes no update/set endpoint for videooptimizationpacingpolicylabel
	// (only add/delete/get/rename). All schema attributes use RequiresReplace, so
	// Terraform never invokes Update for an attribute change. This body is a
	// documented no-op that simply refreshes state.
	// NOTE: videooptimization pacing is deprecated on the NITRO/CLI side.
	tflog.Debug(ctx, "Update is a no-op for videooptimizationpacingpolicylabel; all attributes are RequiresReplace and there is no NITRO update endpoint")

	// Read the updated state back
	r.readVideooptimizationpacingpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationpacingpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationpacingpolicylabel resource")
	// Named resource - delete using DeleteResource
	labelname_value := data.Labelname.ValueString()
	err := r.client.DeleteResource(service.Videooptimizationpacingpolicylabel.Type(), labelname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete videooptimizationpacingpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted videooptimizationpacingpolicylabel resource")
}

// Helper function to read videooptimizationpacingpolicylabel data from API
func (r *VideooptimizationpacingpolicylabelResource) readVideooptimizationpacingpolicylabelFromApi(ctx context.Context, data *VideooptimizationpacingpolicylabelResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	labelname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Videooptimizationpacingpolicylabel.Type(), labelname_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationpacingpolicylabel, got error: %s", err))
		return
	}

	videooptimizationpacingpolicylabelSetAttrFromGet(ctx, data, getResponseData)

}
