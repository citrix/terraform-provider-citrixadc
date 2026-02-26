package vpnglobal_authenticationldappolicy_binding

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
var _ resource.Resource = &VpnglobalAuthenticationldappolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalAuthenticationldappolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalAuthenticationldappolicyBindingResource)(nil)

func NewVpnglobalAuthenticationldappolicyBindingResource() resource.Resource {
	return &VpnglobalAuthenticationldappolicyBindingResource{}
}

// VpnglobalAuthenticationldappolicyBindingResource defines the resource implementation.
type VpnglobalAuthenticationldappolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalAuthenticationldappolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalAuthenticationldappolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_authenticationldappolicy_binding"
}

func (r *VpnglobalAuthenticationldappolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalAuthenticationldappolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalAuthenticationldappolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_authenticationldappolicy_binding resource")

	// vpnglobal_authenticationldappolicy_binding := vpnglobal_authenticationldappolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationldappolicy_binding.Type(), &vpnglobal_authenticationldappolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_authenticationldappolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnglobal_authenticationldappolicy_binding-config")

	tflog.Trace(ctx, "Created vpnglobal_authenticationldappolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalAuthenticationldappolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationldappolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalAuthenticationldappolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_authenticationldappolicy_binding resource")

	r.readVpnglobalAuthenticationldappolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationldappolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnglobalAuthenticationldappolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnglobal_authenticationldappolicy_binding resource")

	// Create API request body from the model
	// vpnglobal_authenticationldappolicy_binding := vpnglobal_authenticationldappolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationldappolicy_binding.Type(), &vpnglobal_authenticationldappolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_authenticationldappolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnglobal_authenticationldappolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalAuthenticationldappolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationldappolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalAuthenticationldappolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_authenticationldappolicy_binding resource")

	// For vpnglobal_authenticationldappolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnglobal_authenticationldappolicy_binding resource from state")
}

// Helper function to read vpnglobal_authenticationldappolicy_binding data from API
func (r *VpnglobalAuthenticationldappolicyBindingResource) readVpnglobalAuthenticationldappolicyBindingFromApi(ctx context.Context, data *VpnglobalAuthenticationldappolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnglobal_authenticationldappolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_authenticationldappolicy_binding, got error: %s", err))
		return
	}

	vpnglobal_authenticationldappolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
