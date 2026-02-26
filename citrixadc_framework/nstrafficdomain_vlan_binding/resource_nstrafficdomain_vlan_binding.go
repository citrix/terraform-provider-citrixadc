package nstrafficdomain_vlan_binding

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
var _ resource.Resource = &NstrafficdomainVlanBindingResource{}
var _ resource.ResourceWithConfigure = (*NstrafficdomainVlanBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NstrafficdomainVlanBindingResource)(nil)

func NewNstrafficdomainVlanBindingResource() resource.Resource {
	return &NstrafficdomainVlanBindingResource{}
}

// NstrafficdomainVlanBindingResource defines the resource implementation.
type NstrafficdomainVlanBindingResource struct {
	client *service.NitroClient
}

func (r *NstrafficdomainVlanBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstrafficdomainVlanBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstrafficdomain_vlan_binding"
}

func (r *NstrafficdomainVlanBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstrafficdomainVlanBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NstrafficdomainVlanBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstrafficdomain_vlan_binding resource")

	// nstrafficdomain_vlan_binding := nstrafficdomain_vlan_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nstrafficdomain_vlan_binding.Type(), &nstrafficdomain_vlan_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstrafficdomain_vlan_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nstrafficdomain_vlan_binding-config")

	tflog.Trace(ctx, "Created nstrafficdomain_vlan_binding resource")

	// Read the updated state back
	r.readNstrafficdomainVlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainVlanBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstrafficdomainVlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstrafficdomain_vlan_binding resource")

	r.readNstrafficdomainVlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainVlanBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NstrafficdomainVlanBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nstrafficdomain_vlan_binding resource")

	// Create API request body from the model
	// nstrafficdomain_vlan_binding := nstrafficdomain_vlan_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nstrafficdomain_vlan_binding.Type(), &nstrafficdomain_vlan_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nstrafficdomain_vlan_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nstrafficdomain_vlan_binding resource")

	// Read the updated state back
	r.readNstrafficdomainVlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstrafficdomainVlanBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstrafficdomainVlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstrafficdomain_vlan_binding resource")

	// For nstrafficdomain_vlan_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nstrafficdomain_vlan_binding resource from state")
}

// Helper function to read nstrafficdomain_vlan_binding data from API
func (r *NstrafficdomainVlanBindingResource) readNstrafficdomainVlanBindingFromApi(ctx context.Context, data *NstrafficdomainVlanBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nstrafficdomain_vlan_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstrafficdomain_vlan_binding, got error: %s", err))
		return
	}

	nstrafficdomain_vlan_bindingSetAttrFromGet(ctx, data, getResponseData)

}
