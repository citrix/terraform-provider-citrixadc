package appfwwsdl

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
var _ resource.Resource = &AppfwwsdlResource{}
var _ resource.ResourceWithConfigure = (*AppfwwsdlResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwwsdlResource)(nil)

func NewAppfwwsdlResource() resource.Resource {
	return &AppfwwsdlResource{}
}

// AppfwwsdlResource defines the resource implementation.
type AppfwwsdlResource struct {
	client *service.NitroClient
}

func (r *AppfwwsdlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwwsdlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwwsdl"
}

func (r *AppfwwsdlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwwsdlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwwsdlResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwwsdl resource")

	// appfwwsdl := appfwwsdlGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwwsdl.Type(), &appfwwsdl)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwwsdl, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwwsdl-config")

	tflog.Trace(ctx, "Created appfwwsdl resource")

	// Read the updated state back
	r.readAppfwwsdlFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwwsdlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwwsdlResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwwsdl resource")

	r.readAppfwwsdlFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwwsdlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwwsdlResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwwsdl resource")

	// Create API request body from the model
	// appfwwsdl := appfwwsdlGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwwsdl.Type(), &appfwwsdl)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwwsdl, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwwsdl resource")

	// Read the updated state back
	r.readAppfwwsdlFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwwsdlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwwsdlResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwwsdl resource")

	// For appfwwsdl, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwwsdl resource from state")
}

// Helper function to read appfwwsdl data from API
func (r *AppfwwsdlResource) readAppfwwsdlFromApi(ctx context.Context, data *AppfwwsdlResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwwsdl.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwwsdl, got error: %s", err))
		return
	}

	appfwwsdlSetAttrFromGet(ctx, data, getResponseData)

}
