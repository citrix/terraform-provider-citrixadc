package vlan_interface_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VlanInterfaceBindingResource{}
var _ resource.ResourceWithConfigure = (*VlanInterfaceBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VlanInterfaceBindingResource)(nil)

func NewVlanInterfaceBindingResource() resource.Resource {
	return &VlanInterfaceBindingResource{}
}

// VlanInterfaceBindingResource defines the resource implementation.
type VlanInterfaceBindingResource struct {
	client *service.NitroClient
}

func (r *VlanInterfaceBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VlanInterfaceBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_interface_binding"
}

func (r *VlanInterfaceBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VlanInterfaceBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VlanInterfaceBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vlan_interface_binding resource")
	vlan_interface_binding := vlan_interface_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource (NITRO add for this binding is PUT)
	err := r.client.UpdateUnnamedResource(service.Vlan_interface_binding.Type(), &vlan_interface_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vlan_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vlan_interface_binding resource")

	// Set ID for the resource before reading state (vlanid:..,ifnum:..)
	data.Id = types.StringValue(vlanInterfaceBindingComposeId(&data))

	// Read the updated state back
	r.readVlanInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanInterfaceBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VlanInterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vlan_interface_binding resource")

	r.readVlanInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanInterfaceBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VlanInterfaceBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state. All attributes are RequiresReplace, so this
	// Update is a documented no-op (NITRO has no update endpoint for this binding).
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for vlan_interface_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readVlanInterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanInterfaceBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VlanInterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vlan_interface_binding resource")

	// Parse the ID (handles both new "vlanid:..,ifnum:.." and legacy "vlanid,ifnum")
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vlanid", "ifnum"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	vlanid, ok := idMap["vlanid"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'vlanid' not found in ID")
		return
	}
	ifnum, ok := idMap["ifnum"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Bound attribute 'ifnum' not found in ID")
		return
	}

	// Mirror the SDK v2 contract: delete the binding identified by vlanid with the
	// single ifnum arg, URL-encoding the value (ifnum may contain a slash, e.g. "1/1").
	args := []string{fmt.Sprintf("ifnum:%s", utils.UrlEncode(ifnum))}
	err = r.client.DeleteResourceWithArgs(service.Vlan_interface_binding.Type(), vlanid, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vlan_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vlan_interface_binding binding")
}

// Helper function to read vlan_interface_binding data from API
func (r *VlanInterfaceBindingResource) readVlanInterfaceBindingFromApi(ctx context.Context, data *VlanInterfaceBindingResourceModel, diags *diag.Diagnostics) {

	// Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vlanid", "ifnum"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	vlanid, ok := idMap["vlanid"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'vlanid' not found in ID string")
		return
	}
	ifnum, ok := idMap["ifnum"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'ifnum' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vlan_interface_binding.Type(),
		ResourceName:             vlanid,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vlan_interface_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vlan_interface_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right ifnum
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["ifnum"].(string); ok && val == ifnum {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", "vlan_interface_binding not found with the provided ID attributes")
		return
	}

	vlan_interface_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
