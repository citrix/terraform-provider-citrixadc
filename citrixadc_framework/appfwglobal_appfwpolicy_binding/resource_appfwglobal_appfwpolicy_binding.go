package appfwglobal_appfwpolicy_binding

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
var _ resource.Resource = &AppfwglobalAppfwpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppfwglobalAppfwpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwglobalAppfwpolicyBindingResource)(nil)

func NewAppfwglobalAppfwpolicyBindingResource() resource.Resource {
	return &AppfwglobalAppfwpolicyBindingResource{}
}

// AppfwglobalAppfwpolicyBindingResource defines the resource implementation.
type AppfwglobalAppfwpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AppfwglobalAppfwpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwglobalAppfwpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwglobal_appfwpolicy_binding"
}

func (r *AppfwglobalAppfwpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwglobalAppfwpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwglobalAppfwpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwglobal_appfwpolicy_binding resource")

	// appfwglobal_appfwpolicy_binding := appfwglobal_appfwpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwglobal_appfwpolicy_binding.Type(), &appfwglobal_appfwpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwglobal_appfwpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwglobal_appfwpolicy_binding-config")

	tflog.Trace(ctx, "Created appfwglobal_appfwpolicy_binding resource")

	// Read the updated state back
	r.readAppfwglobalAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAppfwpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwglobalAppfwpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwglobal_appfwpolicy_binding resource")

	r.readAppfwglobalAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAppfwpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwglobalAppfwpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwglobal_appfwpolicy_binding resource")

	// Create API request body from the model
	// appfwglobal_appfwpolicy_binding := appfwglobal_appfwpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwglobal_appfwpolicy_binding.Type(), &appfwglobal_appfwpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwglobal_appfwpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwglobal_appfwpolicy_binding resource")

	// Read the updated state back
	r.readAppfwglobalAppfwpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwglobalAppfwpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwglobalAppfwpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwglobal_appfwpolicy_binding resource")

	// For appfwglobal_appfwpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwglobal_appfwpolicy_binding resource from state")
}

// Helper function to read appfwglobal_appfwpolicy_binding data from API
func (r *AppfwglobalAppfwpolicyBindingResource) readAppfwglobalAppfwpolicyBindingFromApi(ctx context.Context, data *AppfwglobalAppfwpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwglobal_appfwpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwglobal_appfwpolicy_binding, got error: %s", err))
		return
	}

	appfwglobal_appfwpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
