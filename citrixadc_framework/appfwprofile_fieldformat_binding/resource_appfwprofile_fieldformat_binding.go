package appfwprofile_fieldformat_binding

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
var _ resource.Resource = &AppfwprofileFieldformatBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileFieldformatBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileFieldformatBindingResource)(nil)

func NewAppfwprofileFieldformatBindingResource() resource.Resource {
	return &AppfwprofileFieldformatBindingResource{}
}

// AppfwprofileFieldformatBindingResource defines the resource implementation.
type AppfwprofileFieldformatBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileFieldformatBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileFieldformatBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_fieldformat_binding"
}

func (r *AppfwprofileFieldformatBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileFieldformatBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileFieldformatBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_fieldformat_binding resource")

	// appfwprofile_fieldformat_binding := appfwprofile_fieldformat_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_fieldformat_binding.Type(), &appfwprofile_fieldformat_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_fieldformat_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_fieldformat_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_fieldformat_binding resource")

	// Read the updated state back
	r.readAppfwprofileFieldformatBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFieldformatBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileFieldformatBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_fieldformat_binding resource")

	r.readAppfwprofileFieldformatBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFieldformatBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileFieldformatBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_fieldformat_binding resource")

	// Create API request body from the model
	// appfwprofile_fieldformat_binding := appfwprofile_fieldformat_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_fieldformat_binding.Type(), &appfwprofile_fieldformat_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_fieldformat_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_fieldformat_binding resource")

	// Read the updated state back
	r.readAppfwprofileFieldformatBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileFieldformatBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileFieldformatBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_fieldformat_binding resource")

	// For appfwprofile_fieldformat_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_fieldformat_binding resource from state")
}

// Helper function to read appfwprofile_fieldformat_binding data from API
func (r *AppfwprofileFieldformatBindingResource) readAppfwprofileFieldformatBindingFromApi(ctx context.Context, data *AppfwprofileFieldformatBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_fieldformat_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_fieldformat_binding, got error: %s", err))
		return
	}

	appfwprofile_fieldformat_bindingSetAttrFromGet(ctx, data, getResponseData)

}
