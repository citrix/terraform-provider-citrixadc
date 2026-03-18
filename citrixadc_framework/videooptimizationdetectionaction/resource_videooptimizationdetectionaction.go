package videooptimizationdetectionaction

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
var _ resource.Resource = &VideooptimizationdetectionactionResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationdetectionactionResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationdetectionactionResource)(nil)

func NewVideooptimizationdetectionactionResource() resource.Resource {
	return &VideooptimizationdetectionactionResource{}
}

// VideooptimizationdetectionactionResource defines the resource implementation.
type VideooptimizationdetectionactionResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationdetectionactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationdetectionactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationdetectionaction"
}

func (r *VideooptimizationdetectionactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationdetectionactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationdetectionactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationdetectionaction resource")

	// videooptimizationdetectionaction := videooptimizationdetectionactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Videooptimizationdetectionaction.Type(), &videooptimizationdetectionaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationdetectionaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("videooptimizationdetectionaction-config")

	tflog.Trace(ctx, "Created videooptimizationdetectionaction resource")

	// Read the updated state back
	r.readVideooptimizationdetectionactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationdetectionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationdetectionaction resource")

	r.readVideooptimizationdetectionactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VideooptimizationdetectionactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating videooptimizationdetectionaction resource")

	// Create API request body from the model
	// videooptimizationdetectionaction := videooptimizationdetectionactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Videooptimizationdetectionaction.Type(), &videooptimizationdetectionaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update videooptimizationdetectionaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated videooptimizationdetectionaction resource")

	// Read the updated state back
	r.readVideooptimizationdetectionactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationdetectionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationdetectionaction resource")

	// For videooptimizationdetectionaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted videooptimizationdetectionaction resource from state")
}

// Helper function to read videooptimizationdetectionaction data from API
func (r *VideooptimizationdetectionactionResource) readVideooptimizationdetectionactionFromApi(ctx context.Context, data *VideooptimizationdetectionactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Videooptimizationdetectionaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationdetectionaction, got error: %s", err))
		return
	}

	videooptimizationdetectionactionSetAttrFromGet(ctx, data, getResponseData)

}
