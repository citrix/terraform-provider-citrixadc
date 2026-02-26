package vxlan_nsip6_binding

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
var _ resource.Resource = &VxlanNsip6BindingResource{}
var _ resource.ResourceWithConfigure = (*VxlanNsip6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*VxlanNsip6BindingResource)(nil)

func NewVxlanNsip6BindingResource() resource.Resource {
	return &VxlanNsip6BindingResource{}
}

// VxlanNsip6BindingResource defines the resource implementation.
type VxlanNsip6BindingResource struct {
	client *service.NitroClient
}

func (r *VxlanNsip6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VxlanNsip6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vxlan_nsip6_binding"
}

func (r *VxlanNsip6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VxlanNsip6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VxlanNsip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vxlan_nsip6_binding resource")

	// vxlan_nsip6_binding := vxlan_nsip6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vxlan_nsip6_binding.Type(), &vxlan_nsip6_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vxlan_nsip6_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vxlan_nsip6_binding-config")

	tflog.Trace(ctx, "Created vxlan_nsip6_binding resource")

	// Read the updated state back
	r.readVxlanNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanNsip6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VxlanNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vxlan_nsip6_binding resource")

	r.readVxlanNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanNsip6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VxlanNsip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vxlan_nsip6_binding resource")

	// Create API request body from the model
	// vxlan_nsip6_binding := vxlan_nsip6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vxlan_nsip6_binding.Type(), &vxlan_nsip6_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vxlan_nsip6_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vxlan_nsip6_binding resource")

	// Read the updated state back
	r.readVxlanNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanNsip6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VxlanNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vxlan_nsip6_binding resource")

	// For vxlan_nsip6_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vxlan_nsip6_binding resource from state")
}

// Helper function to read vxlan_nsip6_binding data from API
func (r *VxlanNsip6BindingResource) readVxlanNsip6BindingFromApi(ctx context.Context, data *VxlanNsip6BindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vxlan_nsip6_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vxlan_nsip6_binding, got error: %s", err))
		return
	}

	vxlan_nsip6_bindingSetAttrFromGet(ctx, data, getResponseData)

}
