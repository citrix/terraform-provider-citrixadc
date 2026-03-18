package vpnglobal_authenticationradiuspolicy_binding

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
var _ resource.Resource = &VpnglobalAuthenticationradiuspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalAuthenticationradiuspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalAuthenticationradiuspolicyBindingResource)(nil)

func NewVpnglobalAuthenticationradiuspolicyBindingResource() resource.Resource {
	return &VpnglobalAuthenticationradiuspolicyBindingResource{}
}

// VpnglobalAuthenticationradiuspolicyBindingResource defines the resource implementation.
type VpnglobalAuthenticationradiuspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalAuthenticationradiuspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalAuthenticationradiuspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_authenticationradiuspolicy_binding"
}

func (r *VpnglobalAuthenticationradiuspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalAuthenticationradiuspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_authenticationradiuspolicy_binding resource")

	// vpnglobal_authenticationradiuspolicy_binding := vpnglobal_authenticationradiuspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationradiuspolicy_binding.Type(), &vpnglobal_authenticationradiuspolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_authenticationradiuspolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnglobal_authenticationradiuspolicy_binding-config")

	tflog.Trace(ctx, "Created vpnglobal_authenticationradiuspolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalAuthenticationradiuspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationradiuspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_authenticationradiuspolicy_binding resource")

	r.readVpnglobalAuthenticationradiuspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationradiuspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnglobal_authenticationradiuspolicy_binding resource")

	// Create API request body from the model
	// vpnglobal_authenticationradiuspolicy_binding := vpnglobal_authenticationradiuspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationradiuspolicy_binding.Type(), &vpnglobal_authenticationradiuspolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_authenticationradiuspolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnglobal_authenticationradiuspolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalAuthenticationradiuspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationradiuspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalAuthenticationradiuspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_authenticationradiuspolicy_binding resource")

	// For vpnglobal_authenticationradiuspolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnglobal_authenticationradiuspolicy_binding resource from state")
}

// Helper function to read vpnglobal_authenticationradiuspolicy_binding data from API
func (r *VpnglobalAuthenticationradiuspolicyBindingResource) readVpnglobalAuthenticationradiuspolicyBindingFromApi(ctx context.Context, data *VpnglobalAuthenticationradiuspolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnglobal_authenticationradiuspolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_authenticationradiuspolicy_binding, got error: %s", err))
		return
	}

	vpnglobal_authenticationradiuspolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
