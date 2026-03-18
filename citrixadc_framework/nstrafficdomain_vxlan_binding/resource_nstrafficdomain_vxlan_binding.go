package nstrafficdomain_vxlan_binding

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
var _ resource.Resource = &NstrafficdomainVxlanBindingResource{}
var _ resource.ResourceWithConfigure = (*NstrafficdomainVxlanBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NstrafficdomainVxlanBindingResource)(nil)

func NewNstrafficdomainVxlanBindingResource() resource.Resource {
	return &NstrafficdomainVxlanBindingResource{}
}

// NstrafficdomainVxlanBindingResource defines the resource implementation.
type NstrafficdomainVxlanBindingResource struct {
	client *service.NitroClient
}

func (r *NstrafficdomainVxlanBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstrafficdomainVxlanBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstrafficdomain_vxlan_binding"
}

func (r *NstrafficdomainVxlanBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstrafficdomainVxlanBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstrafficdomainVxlanBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstrafficdomain_vxlan_binding resource")

	// nstrafficdomain_vxlan_binding := nstrafficdomain_vxlan_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nstrafficdomain_vxlan_binding.Type(), &nstrafficdomain_vxlan_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstrafficdomain_vxlan_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nstrafficdomain_vxlan_binding-config")

	tflog.Trace(ctx, "Created nstrafficdomain_vxlan_binding resource")

	// Read the updated state back
	r.readNstrafficdomainVxlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainVxlanBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstrafficdomainVxlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstrafficdomain_vxlan_binding resource")

	r.readNstrafficdomainVxlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainVxlanBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NstrafficdomainVxlanBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nstrafficdomain_vxlan_binding resource")

	// Create API request body from the model
	// nstrafficdomain_vxlan_binding := nstrafficdomain_vxlan_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nstrafficdomain_vxlan_binding.Type(), &nstrafficdomain_vxlan_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nstrafficdomain_vxlan_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nstrafficdomain_vxlan_binding resource")

	// Read the updated state back
	r.readNstrafficdomainVxlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainVxlanBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstrafficdomainVxlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstrafficdomain_vxlan_binding resource")

	// For nstrafficdomain_vxlan_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nstrafficdomain_vxlan_binding resource from state")
}

// Helper function to read nstrafficdomain_vxlan_binding data from API
func (r *NstrafficdomainVxlanBindingResource) readNstrafficdomainVxlanBindingFromApi(ctx context.Context, data *NstrafficdomainVxlanBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nstrafficdomain_vxlan_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstrafficdomain_vxlan_binding, got error: %s", err))
		return
	}

	nstrafficdomain_vxlan_bindingSetAttrFromGet(ctx, data, getResponseData)

}
