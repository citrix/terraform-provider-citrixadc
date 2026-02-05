package nscapacity

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
var _ resource.Resource = &NscapacityResource{}
var _ resource.ResourceWithConfigure = (*NscapacityResource)(nil)
var _ resource.ResourceWithImportState = (*NscapacityResource)(nil)

func NewNscapacityResource() resource.Resource {
	return &NscapacityResource{}
}

// NscapacityResource defines the resource implementation.
type NscapacityResource struct {
	client *service.NitroClient
}

func (r *NscapacityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NscapacityResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nscapacity"
}

func (r *NscapacityResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NscapacityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NscapacityResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nscapacity resource")

	// nscapacity := nscapacityGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nscapacity.Type(), &nscapacity)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nscapacity, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nscapacity-config")

	tflog.Trace(ctx, "Created nscapacity resource")

	// Read the updated state back
	r.readNscapacityFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscapacityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NscapacityResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nscapacity resource")

	r.readNscapacityFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscapacityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NscapacityResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nscapacity resource")

	// Create API request body from the model
	// nscapacity := nscapacityGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nscapacity.Type(), &nscapacity)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nscapacity, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nscapacity resource")

	// Read the updated state back
	r.readNscapacityFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscapacityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NscapacityResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nscapacity resource")

	// For nscapacity, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nscapacity resource from state")
}

// Helper function to read nscapacity data from API
func (r *NscapacityResource) readNscapacityFromApi(ctx context.Context, data *NscapacityResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nscapacity.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nscapacity, got error: %s", err))
		return
	}

	nscapacitySetAttrFromGet(ctx, data, getResponseData)

}
