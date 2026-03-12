package vpnglobal_authenticationnegotiatepolicy_binding

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
var _ resource.Resource = &VpnglobalAuthenticationnegotiatepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalAuthenticationnegotiatepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalAuthenticationnegotiatepolicyBindingResource)(nil)

func NewVpnglobalAuthenticationnegotiatepolicyBindingResource() resource.Resource {
	return &VpnglobalAuthenticationnegotiatepolicyBindingResource{}
}

// VpnglobalAuthenticationnegotiatepolicyBindingResource defines the resource implementation.
type VpnglobalAuthenticationnegotiatepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_authenticationnegotiatepolicy_binding"
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_authenticationnegotiatepolicy_binding resource")

	// vpnglobal_authenticationnegotiatepolicy_binding := vpnglobal_authenticationnegotiatepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationnegotiatepolicy_binding.Type(), &vpnglobal_authenticationnegotiatepolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_authenticationnegotiatepolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnglobal_authenticationnegotiatepolicy_binding-config")

	tflog.Trace(ctx, "Created vpnglobal_authenticationnegotiatepolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_authenticationnegotiatepolicy_binding resource")

	r.readVpnglobalAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnglobalAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnglobal_authenticationnegotiatepolicy_binding resource")

	// Create API request body from the model
	// vpnglobal_authenticationnegotiatepolicy_binding := vpnglobal_authenticationnegotiatepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_authenticationnegotiatepolicy_binding.Type(), &vpnglobal_authenticationnegotiatepolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_authenticationnegotiatepolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnglobal_authenticationnegotiatepolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_authenticationnegotiatepolicy_binding resource")

	// For vpnglobal_authenticationnegotiatepolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnglobal_authenticationnegotiatepolicy_binding resource from state")
}

// Helper function to read vpnglobal_authenticationnegotiatepolicy_binding data from API
func (r *VpnglobalAuthenticationnegotiatepolicyBindingResource) readVpnglobalAuthenticationnegotiatepolicyBindingFromApi(ctx context.Context, data *VpnglobalAuthenticationnegotiatepolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnglobal_authenticationnegotiatepolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_authenticationnegotiatepolicy_binding, got error: %s", err))
		return
	}

	vpnglobal_authenticationnegotiatepolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
