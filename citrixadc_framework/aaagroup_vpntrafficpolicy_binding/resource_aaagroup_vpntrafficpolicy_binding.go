package aaagroup_vpntrafficpolicy_binding

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
var _ resource.Resource = &AaagroupVpntrafficpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AaagroupVpntrafficpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AaagroupVpntrafficpolicyBindingResource)(nil)

func NewAaagroupVpntrafficpolicyBindingResource() resource.Resource {
	return &AaagroupVpntrafficpolicyBindingResource{}
}

// AaagroupVpntrafficpolicyBindingResource defines the resource implementation.
type AaagroupVpntrafficpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AaagroupVpntrafficpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaagroupVpntrafficpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaagroup_vpntrafficpolicy_binding"
}

func (r *AaagroupVpntrafficpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaagroupVpntrafficpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaagroupVpntrafficpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaagroup_vpntrafficpolicy_binding resource")

	// aaagroup_vpntrafficpolicy_binding := aaagroup_vpntrafficpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaagroup_vpntrafficpolicy_binding.Type(), &aaagroup_vpntrafficpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaagroup_vpntrafficpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaagroup_vpntrafficpolicy_binding-config")

	tflog.Trace(ctx, "Created aaagroup_vpntrafficpolicy_binding resource")

	// Read the updated state back
	r.readAaagroupVpntrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupVpntrafficpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaagroupVpntrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaagroup_vpntrafficpolicy_binding resource")

	r.readAaagroupVpntrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupVpntrafficpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaagroupVpntrafficpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaagroup_vpntrafficpolicy_binding resource")

	// Create API request body from the model
	// aaagroup_vpntrafficpolicy_binding := aaagroup_vpntrafficpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaagroup_vpntrafficpolicy_binding.Type(), &aaagroup_vpntrafficpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaagroup_vpntrafficpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaagroup_vpntrafficpolicy_binding resource")

	// Read the updated state back
	r.readAaagroupVpntrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaagroupVpntrafficpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaagroupVpntrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaagroup_vpntrafficpolicy_binding resource")

	// For aaagroup_vpntrafficpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaagroup_vpntrafficpolicy_binding resource from state")
}

// Helper function to read aaagroup_vpntrafficpolicy_binding data from API
func (r *AaagroupVpntrafficpolicyBindingResource) readAaagroupVpntrafficpolicyBindingFromApi(ctx context.Context, data *AaagroupVpntrafficpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaagroup_vpntrafficpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaagroup_vpntrafficpolicy_binding, got error: %s", err))
		return
	}

	aaagroup_vpntrafficpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
