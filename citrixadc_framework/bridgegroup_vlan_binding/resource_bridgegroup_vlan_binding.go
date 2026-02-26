package bridgegroup_vlan_binding

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
var _ resource.Resource = &BridgegroupVlanBindingResource{}
var _ resource.ResourceWithConfigure = (*BridgegroupVlanBindingResource)(nil)
var _ resource.ResourceWithImportState = (*BridgegroupVlanBindingResource)(nil)

func NewBridgegroupVlanBindingResource() resource.Resource {
	return &BridgegroupVlanBindingResource{}
}

// BridgegroupVlanBindingResource defines the resource implementation.
type BridgegroupVlanBindingResource struct {
	client *service.NitroClient
}

func (r *BridgegroupVlanBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BridgegroupVlanBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bridgegroup_vlan_binding"
}

func (r *BridgegroupVlanBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BridgegroupVlanBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BridgegroupVlanBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating bridgegroup_vlan_binding resource")

	// bridgegroup_vlan_binding := bridgegroup_vlan_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Bridgegroup_vlan_binding.Type(), &bridgegroup_vlan_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create bridgegroup_vlan_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("bridgegroup_vlan_binding-config")

	tflog.Trace(ctx, "Created bridgegroup_vlan_binding resource")

	// Read the updated state back
	r.readBridgegroupVlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupVlanBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BridgegroupVlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading bridgegroup_vlan_binding resource")

	r.readBridgegroupVlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupVlanBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data BridgegroupVlanBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating bridgegroup_vlan_binding resource")

	// Create API request body from the model
	// bridgegroup_vlan_binding := bridgegroup_vlan_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Bridgegroup_vlan_binding.Type(), &bridgegroup_vlan_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update bridgegroup_vlan_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated bridgegroup_vlan_binding resource")

	// Read the updated state back
	r.readBridgegroupVlanBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupVlanBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BridgegroupVlanBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting bridgegroup_vlan_binding resource")

	// For bridgegroup_vlan_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted bridgegroup_vlan_binding resource from state")
}

// Helper function to read bridgegroup_vlan_binding data from API
func (r *BridgegroupVlanBindingResource) readBridgegroupVlanBindingFromApi(ctx context.Context, data *BridgegroupVlanBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Bridgegroup_vlan_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read bridgegroup_vlan_binding, got error: %s", err))
		return
	}

	bridgegroup_vlan_bindingSetAttrFromGet(ctx, data, getResponseData)

}
