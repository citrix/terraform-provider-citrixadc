package vxlanvlanmap_vxlan_binding

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
var _ resource.Resource = &VxlanvlanmapVxlanBindingResource{}
var _ resource.ResourceWithConfigure = (*VxlanvlanmapVxlanBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VxlanvlanmapVxlanBindingResource)(nil)

func NewVxlanvlanmapVxlanBindingResource() resource.Resource {
	return &VxlanvlanmapVxlanBindingResource{}
}

// VxlanvlanmapVxlanBindingResource defines the resource implementation.
type VxlanvlanmapVxlanBindingResource struct {
	client *service.NitroClient
}

func (r *VxlanvlanmapVxlanBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VxlanvlanmapVxlanBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vxlanvlanmap_vxlan_binding"
}

func (r *VxlanvlanmapVxlanBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VxlanvlanmapVxlanBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VxlanvlanmapVxlanBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vxlanvlanmap_vxlan_binding resource")

	// vxlanvlanmap_vxlan_binding := vxlanvlanmap_vxlan_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vxlanvlanmap_vxlan_binding.Type(), &vxlanvlanmap_vxlan_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vxlanvlanmap_vxlan_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vxlanvlanmap_vxlan_binding-config")

	tflog.Trace(ctx, "Created vxlanvlanmap_vxlan_binding resource")

	// Read the updated state back
	r.readVxlanvlanmapVxlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanvlanmapVxlanBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VxlanvlanmapVxlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vxlanvlanmap_vxlan_binding resource")

	r.readVxlanvlanmapVxlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanvlanmapVxlanBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VxlanvlanmapVxlanBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vxlanvlanmap_vxlan_binding resource")

	// Create API request body from the model
	// vxlanvlanmap_vxlan_binding := vxlanvlanmap_vxlan_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vxlanvlanmap_vxlan_binding.Type(), &vxlanvlanmap_vxlan_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vxlanvlanmap_vxlan_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vxlanvlanmap_vxlan_binding resource")

	// Read the updated state back
	r.readVxlanvlanmapVxlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanvlanmapVxlanBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VxlanvlanmapVxlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vxlanvlanmap_vxlan_binding resource")

	// For vxlanvlanmap_vxlan_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vxlanvlanmap_vxlan_binding resource from state")
}

// Helper function to read vxlanvlanmap_vxlan_binding data from API
func (r *VxlanvlanmapVxlanBindingResource) readVxlanvlanmapVxlanBindingFromApi(ctx context.Context, data *VxlanvlanmapVxlanBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vxlanvlanmap_vxlan_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vxlanvlanmap_vxlan_binding, got error: %s", err))
		return
	}

	vxlanvlanmap_vxlan_bindingSetAttrFromGet(ctx, data, getResponseData)

}
