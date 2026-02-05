package videooptimizationpacingpolicy

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
var _ resource.Resource = &VideooptimizationpacingpolicyResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationpacingpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationpacingpolicyResource)(nil)

func NewVideooptimizationpacingpolicyResource() resource.Resource {
	return &VideooptimizationpacingpolicyResource{}
}

// VideooptimizationpacingpolicyResource defines the resource implementation.
type VideooptimizationpacingpolicyResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationpacingpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationpacingpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationpacingpolicy"
}

func (r *VideooptimizationpacingpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationpacingpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationpacingpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationpacingpolicy resource")

	// videooptimizationpacingpolicy := videooptimizationpacingpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Videooptimizationpacingpolicy.Type(), &videooptimizationpacingpolicy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationpacingpolicy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("videooptimizationpacingpolicy-config")

	tflog.Trace(ctx, "Created videooptimizationpacingpolicy resource")

	// Read the updated state back
	r.readVideooptimizationpacingpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationpacingpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationpacingpolicy resource")

	r.readVideooptimizationpacingpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VideooptimizationpacingpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating videooptimizationpacingpolicy resource")

	// Create API request body from the model
	// videooptimizationpacingpolicy := videooptimizationpacingpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Videooptimizationpacingpolicy.Type(), &videooptimizationpacingpolicy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update videooptimizationpacingpolicy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated videooptimizationpacingpolicy resource")

	// Read the updated state back
	r.readVideooptimizationpacingpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationpacingpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationpacingpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationpacingpolicy resource")

	// For videooptimizationpacingpolicy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted videooptimizationpacingpolicy resource from state")
}

// Helper function to read videooptimizationpacingpolicy data from API
func (r *VideooptimizationpacingpolicyResource) readVideooptimizationpacingpolicyFromApi(ctx context.Context, data *VideooptimizationpacingpolicyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Videooptimizationpacingpolicy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationpacingpolicy, got error: %s", err))
		return
	}

	videooptimizationpacingpolicySetAttrFromGet(ctx, data, getResponseData)

}
