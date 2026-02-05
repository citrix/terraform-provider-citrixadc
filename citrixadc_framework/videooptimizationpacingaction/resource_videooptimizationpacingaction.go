package videooptimizationpacingaction

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
var _ resource.Resource = &VideooptimizationpacingactionResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationpacingactionResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationpacingactionResource)(nil)

func NewVideooptimizationpacingactionResource() resource.Resource {
	return &VideooptimizationpacingactionResource{}
}

// VideooptimizationpacingactionResource defines the resource implementation.
type VideooptimizationpacingactionResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationpacingactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationpacingactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationpacingaction"
}

func (r *VideooptimizationpacingactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationpacingactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationpacingactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationpacingaction resource")

	// videooptimizationpacingaction := videooptimizationpacingactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Videooptimizationpacingaction.Type(), &videooptimizationpacingaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationpacingaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("videooptimizationpacingaction-config")

	tflog.Trace(ctx, "Created videooptimizationpacingaction resource")

	// Read the updated state back
	r.readVideooptimizationpacingactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationpacingactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationpacingaction resource")

	r.readVideooptimizationpacingactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VideooptimizationpacingactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating videooptimizationpacingaction resource")

	// Create API request body from the model
	// videooptimizationpacingaction := videooptimizationpacingactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Videooptimizationpacingaction.Type(), &videooptimizationpacingaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update videooptimizationpacingaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated videooptimizationpacingaction resource")

	// Read the updated state back
	r.readVideooptimizationpacingactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationpacingactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationpacingaction resource")

	// For videooptimizationpacingaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted videooptimizationpacingaction resource from state")
}

// Helper function to read videooptimizationpacingaction data from API
func (r *VideooptimizationpacingactionResource) readVideooptimizationpacingactionFromApi(ctx context.Context, data *VideooptimizationpacingactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Videooptimizationpacingaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationpacingaction, got error: %s", err))
		return
	}

	videooptimizationpacingactionSetAttrFromGet(ctx, data, getResponseData)

}
