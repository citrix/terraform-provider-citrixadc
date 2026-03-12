package videooptimizationdetectionpolicy

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
var _ resource.Resource = &VideooptimizationdetectionpolicyResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationdetectionpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationdetectionpolicyResource)(nil)

func NewVideooptimizationdetectionpolicyResource() resource.Resource {
	return &VideooptimizationdetectionpolicyResource{}
}

// VideooptimizationdetectionpolicyResource defines the resource implementation.
type VideooptimizationdetectionpolicyResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationdetectionpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationdetectionpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationdetectionpolicy"
}

func (r *VideooptimizationdetectionpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationdetectionpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationdetectionpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationdetectionpolicy resource")

	// videooptimizationdetectionpolicy := videooptimizationdetectionpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Videooptimizationdetectionpolicy.Type(), &videooptimizationdetectionpolicy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationdetectionpolicy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("videooptimizationdetectionpolicy-config")

	tflog.Trace(ctx, "Created videooptimizationdetectionpolicy resource")

	// Read the updated state back
	r.readVideooptimizationdetectionpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationdetectionpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationdetectionpolicy resource")

	r.readVideooptimizationdetectionpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VideooptimizationdetectionpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating videooptimizationdetectionpolicy resource")

	// Create API request body from the model
	// videooptimizationdetectionpolicy := videooptimizationdetectionpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Videooptimizationdetectionpolicy.Type(), &videooptimizationdetectionpolicy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update videooptimizationdetectionpolicy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated videooptimizationdetectionpolicy resource")

	// Read the updated state back
	r.readVideooptimizationdetectionpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationdetectionpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationdetectionpolicy resource")

	// For videooptimizationdetectionpolicy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted videooptimizationdetectionpolicy resource from state")
}

// Helper function to read videooptimizationdetectionpolicy data from API
func (r *VideooptimizationdetectionpolicyResource) readVideooptimizationdetectionpolicyFromApi(ctx context.Context, data *VideooptimizationdetectionpolicyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Videooptimizationdetectionpolicy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationdetectionpolicy, got error: %s", err))
		return
	}

	videooptimizationdetectionpolicySetAttrFromGet(ctx, data, getResponseData)

}
