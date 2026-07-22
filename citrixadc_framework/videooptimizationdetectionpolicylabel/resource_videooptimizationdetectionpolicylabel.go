package videooptimizationdetectionpolicylabel

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
var _ resource.Resource = &VideooptimizationdetectionpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationdetectionpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationdetectionpolicylabelResource)(nil)

func NewVideooptimizationdetectionpolicylabelResource() resource.Resource {
	return &VideooptimizationdetectionpolicylabelResource{}
}

// VideooptimizationdetectionpolicylabelResource defines the resource implementation.
type VideooptimizationdetectionpolicylabelResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationdetectionpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationdetectionpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationdetectionpolicylabel"
}

func (r *VideooptimizationdetectionpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationdetectionpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationdetectionpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationdetectionpolicylabel resource")
	videooptimizationdetectionpolicylabel := videooptimizationdetectionpolicylabelGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	labelname_value := data.Labelname.ValueString()
	_, err := r.client.AddResource(service.Videooptimizationdetectionpolicylabel.Type(), labelname_value, &videooptimizationdetectionpolicylabel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationdetectionpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created videooptimizationdetectionpolicylabel resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	// Read the updated state back
	r.readVideooptimizationdetectionpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationdetectionpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationdetectionpolicylabel resource")

	r.readVideooptimizationdetectionpolicylabelFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Object was deleted out-of-band; remove it from state so a subsequent apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VideooptimizationdetectionpolicylabelResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// NITRO exposes no update/set endpoint for videooptimizationdetectionpolicylabel
	// (only add/delete/get/rename). All schema attributes use RequiresReplace, so
	// Terraform never invokes Update for an attribute change. This body is a
	// documented no-op that simply refreshes state.
	tflog.Debug(ctx, "Update is a no-op for videooptimizationdetectionpolicylabel; all attributes are RequiresReplace and there is no NITRO update endpoint")

	// Read the updated state back
	r.readVideooptimizationdetectionpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationdetectionpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationdetectionpolicylabel resource")
	// Named resource - delete using DeleteResource
	labelname_value := data.Labelname.ValueString()
	err := r.client.DeleteResource(service.Videooptimizationdetectionpolicylabel.Type(), labelname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete videooptimizationdetectionpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted videooptimizationdetectionpolicylabel resource")
}

// Helper function to read videooptimizationdetectionpolicylabel data from API
func (r *VideooptimizationdetectionpolicylabelResource) readVideooptimizationdetectionpolicylabelFromApi(ctx context.Context, data *VideooptimizationdetectionpolicylabelResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	labelname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Videooptimizationdetectionpolicylabel.Type(), labelname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationdetectionpolicylabel, got error: %s", err))
		return
	}

	videooptimizationdetectionpolicylabelSetAttrFromGet(ctx, data, getResponseData)

}
