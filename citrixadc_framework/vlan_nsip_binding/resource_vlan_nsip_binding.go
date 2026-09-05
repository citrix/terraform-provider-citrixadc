package vlan_nsip_binding

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
var _ resource.Resource = &VlanNsipBindingResource{}
var _ resource.ResourceWithConfigure = (*VlanNsipBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VlanNsipBindingResource)(nil)

func NewVlanNsipBindingResource() resource.Resource {
	return &VlanNsipBindingResource{}
}

// VlanNsipBindingResource defines the resource implementation.
type VlanNsipBindingResource struct {
	client *service.NitroClient
}

func (r *VlanNsipBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VlanNsipBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_nsip_binding"
}

func (r *VlanNsipBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VlanNsipBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VlanNsipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vlan_nsip_binding resource")
	vlan_nsip_binding := vlan_nsip_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vlan_nsip_binding.Type(), &vlan_nsip_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vlan_nsip_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vlan_nsip_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(buildVlanNsipBindingId(&data))

	// Read the updated state back
	r.readVlanNsipBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vlan_nsip_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanNsipBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VlanNsipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vlan_nsip_binding resource")

	r.readVlanNsipBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanNsipBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VlanNsipBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for vlan_nsip_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readVlanNsipBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vlan_nsip_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanNsipBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VlanNsipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vlan_nsip_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgsMap.
	// ParseIdString handles both the new key:value ID and the legacy "vlanid,ipaddress" form.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vlanid", "ipaddress"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	parentName, ok := idMap["vlanid"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'vlanid' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["ipaddress"]; ok && val != "" {
		argsMap["ipaddress"] = val
	}
	if val, ok := idMap["netmask"]; ok && val != "" {
		argsMap["netmask"] = val
	}
	if val, ok := idMap["ownergroup"]; ok && val != "" {
		argsMap["ownergroup"] = val
	}
	if val, ok := idMap["td"]; ok && val != "" {
		argsMap["td"] = val
	}

	err = r.client.DeleteResourceWithArgsMap(service.Vlan_nsip_binding.Type(), parentName, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vlan_nsip_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vlan_nsip_binding binding")
}

// Helper function to read vlan_nsip_binding data from API
func (r *VlanNsipBindingResource) readVlanNsipBindingFromApi(ctx context.Context, data *VlanNsipBindingResourceModel, diags *diag.Diagnostics) {

	// Parse the composite ID. ParseIdString accepts both the new key:value form and the
	// legacy positional "vlanid,ipaddress" form.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vlanid", "ipaddress"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	parentName, ok := idMap["vlanid"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'vlanid' not found in ID string")
		return
	}

	ipaddress, ok := idMap["ipaddress"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'ipaddress' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vlan_nsip_binding.Type(),
		ResourceName:             parentName,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vlan_nsip_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Identify the bound record by ipaddress (the user-controlled bound entity). The td
	// value, when present in the ID, further disambiguates.
	foundIndex := -1
	for i, v := range dataArr {
		val, ok := v["ipaddress"].(string)
		if !ok || val != ipaddress {
			continue
		}
		// Match td when it is part of the ID and the record carries it.
		if tdVal, ok := idMap["td"]; ok && tdVal != "" {
			if recTd, ok := v["td"]; ok {
				recTdInt, _ := utils.ConvertToInt64(recTd)
				if fmt.Sprintf("%v", recTdInt) != tdVal {
					continue
				}
			}
		}
		foundIndex = i
		break
	}

	//  Resource is missing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	vlan_nsip_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
