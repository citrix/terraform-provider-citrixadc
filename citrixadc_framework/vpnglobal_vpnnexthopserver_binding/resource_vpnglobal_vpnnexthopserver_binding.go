package vpnglobal_vpnnexthopserver_binding

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
var _ resource.Resource = &VpnglobalVpnnexthopserverBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalVpnnexthopserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalVpnnexthopserverBindingResource)(nil)

func NewVpnglobalVpnnexthopserverBindingResource() resource.Resource {
	return &VpnglobalVpnnexthopserverBindingResource{}
}

// VpnglobalVpnnexthopserverBindingResource defines the resource implementation.
type VpnglobalVpnnexthopserverBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalVpnnexthopserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalVpnnexthopserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_vpnnexthopserver_binding"
}

func (r *VpnglobalVpnnexthopserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalVpnnexthopserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalVpnnexthopserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_vpnnexthopserver_binding resource")

	// vpnglobal_vpnnexthopserver_binding := vpnglobal_vpnnexthopserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnnexthopserver_binding.Type(), &vpnglobal_vpnnexthopserver_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_vpnnexthopserver_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnglobal_vpnnexthopserver_binding-config")

	tflog.Trace(ctx, "Created vpnglobal_vpnnexthopserver_binding resource")

	// Read the updated state back
	r.readVpnglobalVpnnexthopserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnnexthopserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalVpnnexthopserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_vpnnexthopserver_binding resource")

	r.readVpnglobalVpnnexthopserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnnexthopserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnglobalVpnnexthopserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnglobal_vpnnexthopserver_binding resource")

	// Create API request body from the model
	// vpnglobal_vpnnexthopserver_binding := vpnglobal_vpnnexthopserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnnexthopserver_binding.Type(), &vpnglobal_vpnnexthopserver_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_vpnnexthopserver_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnglobal_vpnnexthopserver_binding resource")

	// Read the updated state back
	r.readVpnglobalVpnnexthopserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnnexthopserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalVpnnexthopserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_vpnnexthopserver_binding resource")

	// For vpnglobal_vpnnexthopserver_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnglobal_vpnnexthopserver_binding resource from state")
}

// Helper function to read vpnglobal_vpnnexthopserver_binding data from API
func (r *VpnglobalVpnnexthopserverBindingResource) readVpnglobalVpnnexthopserverBindingFromApi(ctx context.Context, data *VpnglobalVpnnexthopserverBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnglobal_vpnnexthopserver_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_vpnnexthopserver_binding, got error: %s", err))
		return
	}

	vpnglobal_vpnnexthopserver_bindingSetAttrFromGet(ctx, data, getResponseData)

}
