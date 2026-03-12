package hafailover

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
var _ resource.Resource = &HafailoverResource{}
var _ resource.ResourceWithConfigure = (*HafailoverResource)(nil)
var _ resource.ResourceWithImportState = (*HafailoverResource)(nil)

func NewHafailoverResource() resource.Resource {
	return &HafailoverResource{}
}

// HafailoverResource defines the resource implementation.
type HafailoverResource struct {
	client *service.NitroClient
}

func (r *HafailoverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *HafailoverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hafailover"
}

func (r *HafailoverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HafailoverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HafailoverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hafailover resource")

	// hafailover := hafailoverGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Hafailover.Type(), &hafailover)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create hafailover, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("hafailover-config")

	tflog.Trace(ctx, "Created hafailover resource")

	// Read the updated state back
	r.readHafailoverFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafailoverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HafailoverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading hafailover resource")

	r.readHafailoverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafailoverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data HafailoverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating hafailover resource")

	// Create API request body from the model
	// hafailover := hafailoverGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Hafailover.Type(), &hafailover)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update hafailover, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated hafailover resource")

	// Read the updated state back
	r.readHafailoverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HafailoverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HafailoverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting hafailover resource")

	// For hafailover, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted hafailover resource from state")
}

// Helper function to read hafailover data from API
func (r *HafailoverResource) readHafailoverFromApi(ctx context.Context, data *HafailoverResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Hafailover.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read hafailover, got error: %s", err))
		return
	}

	hafailoverSetAttrFromGet(ctx, data, getResponseData)

}
