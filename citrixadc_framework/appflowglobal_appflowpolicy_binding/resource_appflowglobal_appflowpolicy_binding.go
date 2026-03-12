package appflowglobal_appflowpolicy_binding

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
var _ resource.Resource = &AppflowglobalAppflowpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppflowglobalAppflowpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppflowglobalAppflowpolicyBindingResource)(nil)

func NewAppflowglobalAppflowpolicyBindingResource() resource.Resource {
	return &AppflowglobalAppflowpolicyBindingResource{}
}

// AppflowglobalAppflowpolicyBindingResource defines the resource implementation.
type AppflowglobalAppflowpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AppflowglobalAppflowpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppflowglobalAppflowpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appflowglobal_appflowpolicy_binding"
}

func (r *AppflowglobalAppflowpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppflowglobalAppflowpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppflowglobalAppflowpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appflowglobal_appflowpolicy_binding resource")

	// appflowglobal_appflowpolicy_binding := appflowglobal_appflowpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appflowglobal_appflowpolicy_binding.Type(), &appflowglobal_appflowpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appflowglobal_appflowpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appflowglobal_appflowpolicy_binding-config")

	tflog.Trace(ctx, "Created appflowglobal_appflowpolicy_binding resource")

	// Read the updated state back
	r.readAppflowglobalAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowglobalAppflowpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppflowglobalAppflowpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appflowglobal_appflowpolicy_binding resource")

	r.readAppflowglobalAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowglobalAppflowpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppflowglobalAppflowpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appflowglobal_appflowpolicy_binding resource")

	// Create API request body from the model
	// appflowglobal_appflowpolicy_binding := appflowglobal_appflowpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appflowglobal_appflowpolicy_binding.Type(), &appflowglobal_appflowpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appflowglobal_appflowpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appflowglobal_appflowpolicy_binding resource")

	// Read the updated state back
	r.readAppflowglobalAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowglobalAppflowpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppflowglobalAppflowpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appflowglobal_appflowpolicy_binding resource")

	// For appflowglobal_appflowpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appflowglobal_appflowpolicy_binding resource from state")
}

// Helper function to read appflowglobal_appflowpolicy_binding data from API
func (r *AppflowglobalAppflowpolicyBindingResource) readAppflowglobalAppflowpolicyBindingFromApi(ctx context.Context, data *AppflowglobalAppflowpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appflowglobal_appflowpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appflowglobal_appflowpolicy_binding, got error: %s", err))
		return
	}

	appflowglobal_appflowpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
