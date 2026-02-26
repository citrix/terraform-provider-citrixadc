package vpnglobal_authenticationsamlpolicy_binding

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
var _ resource.Resource = &VpnglobalAuthenticationsamlpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalAuthenticationsamlpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalAuthenticationsamlpolicyBindingResource)(nil)

func NewVpnglobalAuthenticationsamlpolicyBindingResource() resource.Resource {
	return &VpnglobalAuthenticationsamlpolicyBindingResource{}
}

// VpnglobalAuthenticationsamlpolicyBindingResource defines the resource implementation.
type VpnglobalAuthenticationsamlpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalAuthenticationsamlpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalAuthenticationsamlpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_authenticationsamlpolicy_binding"
}

func (r *VpnglobalAuthenticationsamlpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalAuthenticationsamlpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalAuthenticationsamlpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_authenticationsamlpolicy_binding resource")

	// vpnglobal_authenticationsamlpolicy_binding := vpnglobal_authenticationsamlpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationsamlpolicy_binding.Type(), &vpnglobal_authenticationsamlpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_authenticationsamlpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnglobal_authenticationsamlpolicy_binding-config")

	tflog.Trace(ctx, "Created vpnglobal_authenticationsamlpolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalAuthenticationsamlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationsamlpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalAuthenticationsamlpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_authenticationsamlpolicy_binding resource")

	r.readVpnglobalAuthenticationsamlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationsamlpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnglobalAuthenticationsamlpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnglobal_authenticationsamlpolicy_binding resource")

	// Create API request body from the model
	// vpnglobal_authenticationsamlpolicy_binding := vpnglobal_authenticationsamlpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationsamlpolicy_binding.Type(), &vpnglobal_authenticationsamlpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_authenticationsamlpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnglobal_authenticationsamlpolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalAuthenticationsamlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationsamlpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalAuthenticationsamlpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_authenticationsamlpolicy_binding resource")

	// For vpnglobal_authenticationsamlpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnglobal_authenticationsamlpolicy_binding resource from state")
}

// Helper function to read vpnglobal_authenticationsamlpolicy_binding data from API
func (r *VpnglobalAuthenticationsamlpolicyBindingResource) readVpnglobalAuthenticationsamlpolicyBindingFromApi(ctx context.Context, data *VpnglobalAuthenticationsamlpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnglobal_authenticationsamlpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_authenticationsamlpolicy_binding, got error: %s", err))
		return
	}

	vpnglobal_authenticationsamlpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
