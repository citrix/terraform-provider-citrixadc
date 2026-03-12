package appflowpolicylabel_appflowpolicy_binding

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
var _ resource.Resource = &AppflowpolicylabelAppflowpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AppflowpolicylabelAppflowpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppflowpolicylabelAppflowpolicyBindingResource)(nil)

func NewAppflowpolicylabelAppflowpolicyBindingResource() resource.Resource {
	return &AppflowpolicylabelAppflowpolicyBindingResource{}
}

// AppflowpolicylabelAppflowpolicyBindingResource defines the resource implementation.
type AppflowpolicylabelAppflowpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appflowpolicylabel_appflowpolicy_binding"
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppflowpolicylabelAppflowpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appflowpolicylabel_appflowpolicy_binding resource")

	// appflowpolicylabel_appflowpolicy_binding := appflowpolicylabel_appflowpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appflowpolicylabel_appflowpolicy_binding.Type(), &appflowpolicylabel_appflowpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appflowpolicylabel_appflowpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appflowpolicylabel_appflowpolicy_binding-config")

	tflog.Trace(ctx, "Created appflowpolicylabel_appflowpolicy_binding resource")

	// Read the updated state back
	r.readAppflowpolicylabelAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppflowpolicylabelAppflowpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appflowpolicylabel_appflowpolicy_binding resource")

	r.readAppflowpolicylabelAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppflowpolicylabelAppflowpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appflowpolicylabel_appflowpolicy_binding resource")

	// Create API request body from the model
	// appflowpolicylabel_appflowpolicy_binding := appflowpolicylabel_appflowpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appflowpolicylabel_appflowpolicy_binding.Type(), &appflowpolicylabel_appflowpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appflowpolicylabel_appflowpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appflowpolicylabel_appflowpolicy_binding resource")

	// Read the updated state back
	r.readAppflowpolicylabelAppflowpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowpolicylabelAppflowpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppflowpolicylabelAppflowpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appflowpolicylabel_appflowpolicy_binding resource")

	// For appflowpolicylabel_appflowpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appflowpolicylabel_appflowpolicy_binding resource from state")
}

// Helper function to read appflowpolicylabel_appflowpolicy_binding data from API
func (r *AppflowpolicylabelAppflowpolicyBindingResource) readAppflowpolicylabelAppflowpolicyBindingFromApi(ctx context.Context, data *AppflowpolicylabelAppflowpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appflowpolicylabel_appflowpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appflowpolicylabel_appflowpolicy_binding, got error: %s", err))
		return
	}

	appflowpolicylabel_appflowpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
