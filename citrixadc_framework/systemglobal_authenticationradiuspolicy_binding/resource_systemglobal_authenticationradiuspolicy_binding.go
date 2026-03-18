package systemglobal_authenticationradiuspolicy_binding

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
var _ resource.Resource = &SystemglobalAuthenticationradiuspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemglobalAuthenticationradiuspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemglobalAuthenticationradiuspolicyBindingResource)(nil)

func NewSystemglobalAuthenticationradiuspolicyBindingResource() resource.Resource {
	return &SystemglobalAuthenticationradiuspolicyBindingResource{}
}

// SystemglobalAuthenticationradiuspolicyBindingResource defines the resource implementation.
type SystemglobalAuthenticationradiuspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemglobal_authenticationradiuspolicy_binding"
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemglobal_authenticationradiuspolicy_binding resource")

	// systemglobal_authenticationradiuspolicy_binding := systemglobal_authenticationradiuspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemglobal_authenticationradiuspolicy_binding.Type(), &systemglobal_authenticationradiuspolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemglobal_authenticationradiuspolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("systemglobal_authenticationradiuspolicy_binding-config")

	tflog.Trace(ctx, "Created systemglobal_authenticationradiuspolicy_binding resource")

	// Read the updated state back
	r.readSystemglobalAuthenticationradiuspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemglobal_authenticationradiuspolicy_binding resource")

	r.readSystemglobalAuthenticationradiuspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SystemglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating systemglobal_authenticationradiuspolicy_binding resource")

	// Create API request body from the model
	// systemglobal_authenticationradiuspolicy_binding := systemglobal_authenticationradiuspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemglobal_authenticationradiuspolicy_binding.Type(), &systemglobal_authenticationradiuspolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemglobal_authenticationradiuspolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated systemglobal_authenticationradiuspolicy_binding resource")

	// Read the updated state back
	r.readSystemglobalAuthenticationradiuspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationradiuspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemglobal_authenticationradiuspolicy_binding resource")

	// For systemglobal_authenticationradiuspolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted systemglobal_authenticationradiuspolicy_binding resource from state")
}

// Helper function to read systemglobal_authenticationradiuspolicy_binding data from API
func (r *SystemglobalAuthenticationradiuspolicyBindingResource) readSystemglobalAuthenticationradiuspolicyBindingFromApi(ctx context.Context, data *SystemglobalAuthenticationradiuspolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Systemglobal_authenticationradiuspolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemglobal_authenticationradiuspolicy_binding, got error: %s", err))
		return
	}

	systemglobal_authenticationradiuspolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
