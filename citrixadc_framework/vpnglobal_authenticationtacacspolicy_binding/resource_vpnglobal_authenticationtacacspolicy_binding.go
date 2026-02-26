package vpnglobal_authenticationtacacspolicy_binding

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
var _ resource.Resource = &VpnglobalAuthenticationtacacspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalAuthenticationtacacspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalAuthenticationtacacspolicyBindingResource)(nil)

func NewVpnglobalAuthenticationtacacspolicyBindingResource() resource.Resource {
	return &VpnglobalAuthenticationtacacspolicyBindingResource{}
}

// VpnglobalAuthenticationtacacspolicyBindingResource defines the resource implementation.
type VpnglobalAuthenticationtacacspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalAuthenticationtacacspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalAuthenticationtacacspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_authenticationtacacspolicy_binding"
}

func (r *VpnglobalAuthenticationtacacspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalAuthenticationtacacspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_authenticationtacacspolicy_binding resource")

	// vpnglobal_authenticationtacacspolicy_binding := vpnglobal_authenticationtacacspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationtacacspolicy_binding.Type(), &vpnglobal_authenticationtacacspolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_authenticationtacacspolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnglobal_authenticationtacacspolicy_binding-config")

	tflog.Trace(ctx, "Created vpnglobal_authenticationtacacspolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationtacacspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_authenticationtacacspolicy_binding resource")

	r.readVpnglobalAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationtacacspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnglobalAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnglobal_authenticationtacacspolicy_binding resource")

	// Create API request body from the model
	// vpnglobal_authenticationtacacspolicy_binding := vpnglobal_authenticationtacacspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationtacacspolicy_binding.Type(), &vpnglobal_authenticationtacacspolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_authenticationtacacspolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnglobal_authenticationtacacspolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationtacacspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_authenticationtacacspolicy_binding resource")

	// For vpnglobal_authenticationtacacspolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnglobal_authenticationtacacspolicy_binding resource from state")
}

// Helper function to read vpnglobal_authenticationtacacspolicy_binding data from API
func (r *VpnglobalAuthenticationtacacspolicyBindingResource) readVpnglobalAuthenticationtacacspolicyBindingFromApi(ctx context.Context, data *VpnglobalAuthenticationtacacspolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnglobal_authenticationtacacspolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_authenticationtacacspolicy_binding, got error: %s", err))
		return
	}

	vpnglobal_authenticationtacacspolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
