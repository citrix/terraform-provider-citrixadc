package vpnglobal_vpnclientlessaccesspolicy_binding

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
var _ resource.Resource = &VpnglobalVpnclientlessaccesspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalVpnclientlessaccesspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalVpnclientlessaccesspolicyBindingResource)(nil)

func NewVpnglobalVpnclientlessaccesspolicyBindingResource() resource.Resource {
	return &VpnglobalVpnclientlessaccesspolicyBindingResource{}
}

// VpnglobalVpnclientlessaccesspolicyBindingResource defines the resource implementation.
type VpnglobalVpnclientlessaccesspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_vpnclientlessaccesspolicy_binding"
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_vpnclientlessaccesspolicy_binding resource")

	// vpnglobal_vpnclientlessaccesspolicy_binding := vpnglobal_vpnclientlessaccesspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnclientlessaccesspolicy_binding.Type(), &vpnglobal_vpnclientlessaccesspolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_vpnclientlessaccesspolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnglobal_vpnclientlessaccesspolicy_binding-config")

	tflog.Trace(ctx, "Created vpnglobal_vpnclientlessaccesspolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalVpnclientlessaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_vpnclientlessaccesspolicy_binding resource")

	r.readVpnglobalVpnclientlessaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnglobalVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnglobal_vpnclientlessaccesspolicy_binding resource")

	// Create API request body from the model
	// vpnglobal_vpnclientlessaccesspolicy_binding := vpnglobal_vpnclientlessaccesspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnclientlessaccesspolicy_binding.Type(), &vpnglobal_vpnclientlessaccesspolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_vpnclientlessaccesspolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnglobal_vpnclientlessaccesspolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalVpnclientlessaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalVpnclientlessaccesspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_vpnclientlessaccesspolicy_binding resource")

	// For vpnglobal_vpnclientlessaccesspolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnglobal_vpnclientlessaccesspolicy_binding resource from state")
}

// Helper function to read vpnglobal_vpnclientlessaccesspolicy_binding data from API
func (r *VpnglobalVpnclientlessaccesspolicyBindingResource) readVpnglobalVpnclientlessaccesspolicyBindingFromApi(ctx context.Context, data *VpnglobalVpnclientlessaccesspolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnglobal_vpnclientlessaccesspolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_vpnclientlessaccesspolicy_binding, got error: %s", err))
		return
	}

	vpnglobal_vpnclientlessaccesspolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
