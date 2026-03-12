package appfwpolicylabel_appfwpolicy_binding

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
var _ resource.Resource = &AppfwpolicylabelAppfwpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwpolicylabelAppfwpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwpolicylabelAppfwpolicyBindingResource)(nil)

func NewAppfwpolicylabelAppfwpolicyBindingResource() resource.Resource {
	return &AppfwpolicylabelAppfwpolicyBindingResource{}
}

// AppfwpolicylabelAppfwpolicyBindingResource defines the resource implementation.
type AppfwpolicylabelAppfwpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwpolicylabelAppfwpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwpolicylabelAppfwpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwpolicylabel_appfwpolicy_binding"
}

func (r *AppfwpolicylabelAppfwpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwpolicylabelAppfwpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwpolicylabelAppfwpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwpolicylabel_appfwpolicy_binding resource")

	// appfwpolicylabel_appfwpolicy_binding := appfwpolicylabel_appfwpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwpolicylabel_appfwpolicy_binding.Type(), &appfwpolicylabel_appfwpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwpolicylabel_appfwpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwpolicylabel_appfwpolicy_binding-config")

	tflog.Trace(ctx, "Created appfwpolicylabel_appfwpolicy_binding resource")

	// Read the updated state back
	r.readAppfwpolicylabelAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwpolicylabelAppfwpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwpolicylabelAppfwpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwpolicylabel_appfwpolicy_binding resource")

	r.readAppfwpolicylabelAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwpolicylabelAppfwpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwpolicylabelAppfwpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwpolicylabel_appfwpolicy_binding resource")

	// Create API request body from the model
	// appfwpolicylabel_appfwpolicy_binding := appfwpolicylabel_appfwpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwpolicylabel_appfwpolicy_binding.Type(), &appfwpolicylabel_appfwpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwpolicylabel_appfwpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwpolicylabel_appfwpolicy_binding resource")

	// Read the updated state back
	r.readAppfwpolicylabelAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwpolicylabelAppfwpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwpolicylabelAppfwpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwpolicylabel_appfwpolicy_binding resource")

	// For appfwpolicylabel_appfwpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwpolicylabel_appfwpolicy_binding resource from state")
}

// Helper function to read appfwpolicylabel_appfwpolicy_binding data from API
func (r *AppfwpolicylabelAppfwpolicyBindingResource) readAppfwpolicylabelAppfwpolicyBindingFromApi(ctx context.Context, data *AppfwpolicylabelAppfwpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwpolicylabel_appfwpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwpolicylabel_appfwpolicy_binding, got error: %s", err))
		return
	}

	appfwpolicylabel_appfwpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
