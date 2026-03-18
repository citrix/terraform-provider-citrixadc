package appfwprofile_crosssitescripting_binding

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
var _ resource.Resource = &AppfwprofileCrosssitescriptingBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwprofileCrosssitescriptingBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwprofileCrosssitescriptingBindingResource)(nil)

func NewAppfwprofileCrosssitescriptingBindingResource() resource.Resource {
	return &AppfwprofileCrosssitescriptingBindingResource{}
}

// AppfwprofileCrosssitescriptingBindingResource defines the resource implementation.
type AppfwprofileCrosssitescriptingBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwprofileCrosssitescriptingBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwprofile_crosssitescripting_binding"
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwprofileCrosssitescriptingBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwprofile_crosssitescripting_binding resource")

	// appfwprofile_crosssitescripting_binding := appfwprofile_crosssitescripting_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_crosssitescripting_binding.Type(), &appfwprofile_crosssitescripting_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwprofile_crosssitescripting_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwprofile_crosssitescripting_binding-config")

	tflog.Trace(ctx, "Created appfwprofile_crosssitescripting_binding resource")

	// Read the updated state back
	r.readAppfwprofileCrosssitescriptingBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwprofileCrosssitescriptingBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwprofile_crosssitescripting_binding resource")

	r.readAppfwprofileCrosssitescriptingBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwprofileCrosssitescriptingBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwprofile_crosssitescripting_binding resource")

	// Create API request body from the model
	// appfwprofile_crosssitescripting_binding := appfwprofile_crosssitescripting_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwprofile_crosssitescripting_binding.Type(), &appfwprofile_crosssitescripting_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwprofile_crosssitescripting_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwprofile_crosssitescripting_binding resource")

	// Read the updated state back
	r.readAppfwprofileCrosssitescriptingBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwprofileCrosssitescriptingBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwprofileCrosssitescriptingBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwprofile_crosssitescripting_binding resource")

	// For appfwprofile_crosssitescripting_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwprofile_crosssitescripting_binding resource from state")
}

// Helper function to read appfwprofile_crosssitescripting_binding data from API
func (r *AppfwprofileCrosssitescriptingBindingResource) readAppfwprofileCrosssitescriptingBindingFromApi(ctx context.Context, data *AppfwprofileCrosssitescriptingBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwprofile_crosssitescripting_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwprofile_crosssitescripting_binding, got error: %s", err))
		return
	}

	appfwprofile_crosssitescripting_bindingSetAttrFromGet(ctx, data, getResponseData)

}
