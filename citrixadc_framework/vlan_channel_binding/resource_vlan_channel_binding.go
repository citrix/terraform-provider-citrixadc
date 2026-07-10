package vlan_channel_binding

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VlanChannelBindingResource{}
var _ resource.ResourceWithConfigure = (*VlanChannelBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VlanChannelBindingResource)(nil)

func NewVlanChannelBindingResource() resource.Resource {
	return &VlanChannelBindingResource{}
}

// VlanChannelBindingResource defines the resource implementation.
type VlanChannelBindingResource struct {
	client *service.NitroClient
}

func (r *VlanChannelBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VlanChannelBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_channel_binding"
}

func (r *VlanChannelBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VlanChannelBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VlanChannelBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vlan_channel_binding resource")
	vlan_channel_binding := vlan_channel_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add verb is HTTP PUT, use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vlan_channel_binding.Type(), &vlan_channel_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vlan_channel_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vlan_channel_binding resource")

	// Set ID for the resource before reading state (legacy order: vlanid,ifnum)
	data.Id = types.StringValue(vlan_channel_bindingComposeId(&data))

	// Read the updated state back
	r.readVlanChannelBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanChannelBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VlanChannelBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vlan_channel_binding resource")

	r.readVlanChannelBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanChannelBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VlanChannelBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// No NITRO update endpoint exists for this binding; every attribute is
	// RequiresReplace so Terraform never reaches Update with a real change.
	tflog.Debug(ctx, "Update is a no-op for vlan_channel_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readVlanChannelBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanChannelBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VlanChannelBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vlan_channel_binding resource")

	// Parse the parent vlan id (legacy order: vlanid,ifnum)
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

	// Delete args per NITRO doc: ifnum, tagged, ownergroup.
	// Use the disambiguating values held in state (matches SDK v2 behaviour).
	args := make([]string, 0)
	if !data.Ifnum.IsNull() && data.Ifnum.ValueString() != "" {
		args = append(args, fmt.Sprintf("ifnum:%s", utils.UrlEncode(data.Ifnum.ValueString())))
	} else if val, ok := idMap["ifnum"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ifnum:%s", utils.UrlEncode(val)))
	}
	if !data.Tagged.IsNull() {
		args = append(args, fmt.Sprintf("tagged:%s", strconv.FormatBool(data.Tagged.ValueBool())))
	}
	if !data.Ownergroup.IsNull() && data.Ownergroup.ValueString() != "" {
		args = append(args, fmt.Sprintf("ownergroup:%s", utils.UrlEncode(data.Ownergroup.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Vlan_channel_binding.Type(), vlanid, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vlan_channel_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vlan_channel_binding binding")
}

// Helper function to read vlan_channel_binding data from API
func (r *VlanChannelBindingResource) readVlanChannelBindingFromApi(ctx context.Context, data *VlanChannelBindingResourceModel, diags *diag.Diagnostics) {

	// Array filter with parent ID - parse from ID (legacy order: vlanid,ifnum)
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

	findParams := service.FindParams{
		ResourceType:             service.Vlan_channel_binding.Type(),
		ResourceName:             vlanid,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vlan_channel_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vlan_channel_binding returned empty array.")
		return
	}

	// Iterate through results to find the one matching ifnum
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["ifnum"].(string); ok && val == ifnum {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", "vlan_channel_binding not found with the provided ID attributes")
		return
	}

	vlan_channel_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
