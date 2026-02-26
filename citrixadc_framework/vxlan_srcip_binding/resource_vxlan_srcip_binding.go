package vxlan_srcip_binding

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
var _ resource.Resource = &VxlanSrcipBindingResource{}
var _ resource.ResourceWithConfigure = (*VxlanSrcipBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VxlanSrcipBindingResource)(nil)

func NewVxlanSrcipBindingResource() resource.Resource {
	return &VxlanSrcipBindingResource{}
}

// VxlanSrcipBindingResource defines the resource implementation.
type VxlanSrcipBindingResource struct {
	client *service.NitroClient
}

func (r *VxlanSrcipBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VxlanSrcipBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vxlan_srcip_binding"
}

func (r *VxlanSrcipBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VxlanSrcipBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VxlanSrcipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vxlan_srcip_binding resource")

	// vxlan_srcip_binding := vxlan_srcip_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vxlan_srcip_binding.Type(), &vxlan_srcip_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vxlan_srcip_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vxlan_srcip_binding-config")

	tflog.Trace(ctx, "Created vxlan_srcip_binding resource")

	// Read the updated state back
	r.readVxlanSrcipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanSrcipBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VxlanSrcipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vxlan_srcip_binding resource")

	r.readVxlanSrcipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanSrcipBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VxlanSrcipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vxlan_srcip_binding resource")

	// Create API request body from the model
	// vxlan_srcip_binding := vxlan_srcip_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vxlan_srcip_binding.Type(), &vxlan_srcip_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vxlan_srcip_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vxlan_srcip_binding resource")

	// Read the updated state back
	r.readVxlanSrcipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanSrcipBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VxlanSrcipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vxlan_srcip_binding resource")

	// For vxlan_srcip_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vxlan_srcip_binding resource from state")
}

// Helper function to read vxlan_srcip_binding data from API
func (r *VxlanSrcipBindingResource) readVxlanSrcipBindingFromApi(ctx context.Context, data *VxlanSrcipBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vxlan_srcip_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vxlan_srcip_binding, got error: %s", err))
		return
	}

	vxlan_srcip_bindingSetAttrFromGet(ctx, data, getResponseData)

}
