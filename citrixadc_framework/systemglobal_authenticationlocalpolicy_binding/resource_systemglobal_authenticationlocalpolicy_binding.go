package systemglobal_authenticationlocalpolicy_binding

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
var _ resource.Resource = &SystemglobalAuthenticationlocalpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*SystemglobalAuthenticationlocalpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SystemglobalAuthenticationlocalpolicyBindingResource)(nil)

func NewSystemglobalAuthenticationlocalpolicyBindingResource() resource.Resource {
	return &SystemglobalAuthenticationlocalpolicyBindingResource{}
}

// SystemglobalAuthenticationlocalpolicyBindingResource defines the resource implementation.
type SystemglobalAuthenticationlocalpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *SystemglobalAuthenticationlocalpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemglobalAuthenticationlocalpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemglobal_authenticationlocalpolicy_binding"
}

func (r *SystemglobalAuthenticationlocalpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemglobalAuthenticationlocalpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemglobalAuthenticationlocalpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemglobal_authenticationlocalpolicy_binding resource")

	// systemglobal_authenticationlocalpolicy_binding := systemglobal_authenticationlocalpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemglobal_authenticationlocalpolicy_binding.Type(), &systemglobal_authenticationlocalpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemglobal_authenticationlocalpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("systemglobal_authenticationlocalpolicy_binding-config")

	tflog.Trace(ctx, "Created systemglobal_authenticationlocalpolicy_binding resource")

	// Read the updated state back
	r.readSystemglobalAuthenticationlocalpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationlocalpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemglobalAuthenticationlocalpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemglobal_authenticationlocalpolicy_binding resource")

	r.readSystemglobalAuthenticationlocalpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationlocalpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SystemglobalAuthenticationlocalpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating systemglobal_authenticationlocalpolicy_binding resource")

	// Create API request body from the model
	// systemglobal_authenticationlocalpolicy_binding := systemglobal_authenticationlocalpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Systemglobal_authenticationlocalpolicy_binding.Type(), &systemglobal_authenticationlocalpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemglobal_authenticationlocalpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated systemglobal_authenticationlocalpolicy_binding resource")

	// Read the updated state back
	r.readSystemglobalAuthenticationlocalpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemglobalAuthenticationlocalpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemglobalAuthenticationlocalpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemglobal_authenticationlocalpolicy_binding resource")

	// For systemglobal_authenticationlocalpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted systemglobal_authenticationlocalpolicy_binding resource from state")
}

// Helper function to read systemglobal_authenticationlocalpolicy_binding data from API
func (r *SystemglobalAuthenticationlocalpolicyBindingResource) readSystemglobalAuthenticationlocalpolicyBindingFromApi(ctx context.Context, data *SystemglobalAuthenticationlocalpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Systemglobal_authenticationlocalpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemglobal_authenticationlocalpolicy_binding, got error: %s", err))
		return
	}

	systemglobal_authenticationlocalpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
