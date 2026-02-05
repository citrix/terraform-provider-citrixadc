package transformpolicylabel

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
var _ resource.Resource = &TransformpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*TransformpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*TransformpolicylabelResource)(nil)

func NewTransformpolicylabelResource() resource.Resource {
	return &TransformpolicylabelResource{}
}

// TransformpolicylabelResource defines the resource implementation.
type TransformpolicylabelResource struct {
	client *service.NitroClient
}

func (r *TransformpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TransformpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transformpolicylabel"
}

func (r *TransformpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TransformpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TransformpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating transformpolicylabel resource")

	// transformpolicylabel := transformpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Transformpolicylabel.Type(), &transformpolicylabel)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create transformpolicylabel, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("transformpolicylabel-config")

	tflog.Trace(ctx, "Created transformpolicylabel resource")

	// Read the updated state back
	r.readTransformpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TransformpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading transformpolicylabel resource")

	r.readTransformpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data TransformpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating transformpolicylabel resource")

	// Create API request body from the model
	// transformpolicylabel := transformpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Transformpolicylabel.Type(), &transformpolicylabel)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update transformpolicylabel, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated transformpolicylabel resource")

	// Read the updated state back
	r.readTransformpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TransformpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting transformpolicylabel resource")

	// For transformpolicylabel, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted transformpolicylabel resource from state")
}

// Helper function to read transformpolicylabel data from API
func (r *TransformpolicylabelResource) readTransformpolicylabelFromApi(ctx context.Context, data *TransformpolicylabelResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Transformpolicylabel.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read transformpolicylabel, got error: %s", err))
		return
	}

	transformpolicylabelSetAttrFromGet(ctx, data, getResponseData)

}
