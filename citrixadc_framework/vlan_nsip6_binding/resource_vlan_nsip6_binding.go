package vlan_nsip6_binding

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
var _ resource.Resource = &VlanNsip6BindingResource{}
var _ resource.ResourceWithConfigure = (*VlanNsip6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*VlanNsip6BindingResource)(nil)

func NewVlanNsip6BindingResource() resource.Resource {
	return &VlanNsip6BindingResource{}
}

// VlanNsip6BindingResource defines the resource implementation.
type VlanNsip6BindingResource struct {
	client *service.NitroClient
}

func (r *VlanNsip6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VlanNsip6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_nsip6_binding"
}

func (r *VlanNsip6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VlanNsip6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VlanNsip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vlan_nsip6_binding resource")
	vlan_nsip6_binding := vlan_nsip6_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vlan_nsip6_binding.Type(), &vlan_nsip6_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vlan_nsip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vlan_nsip6_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(vlan_nsip6_bindingComposeId(&data))

	// Read the updated state back
	r.readVlanNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanNsip6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VlanNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vlan_nsip6_binding resource")

	r.readVlanNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanNsip6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VlanNsip6BindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// No NITRO update endpoint for this binding; every attribute is RequiresReplace,
	// so Update is a documented no-op.
	tflog.Debug(ctx, "Update is a no-op for vlan_nsip6_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readVlanNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanNsip6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VlanNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vlan_nsip6_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vlanid", "ipaddress"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	vlanid, ok := idMap["vlanid"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'vlanid' not found in ID")
		return
	}

	// Build delete args. ipaddress (IPv6, contains '/' and ':') is URL-encoded; the
	// optional filters (netmask/ownergroup/td) come from the resolved state so a
	// binding with a non-default value is targeted correctly.
	args := make([]string, 0)
	if val, ok := idMap["ipaddress"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(val)))
	}
	if !data.Netmask.IsNull() && data.Netmask.ValueString() != "" {
		args = append(args, fmt.Sprintf("netmask:%s", utils.UrlEncode(data.Netmask.ValueString())))
	}
	if !data.Ownergroup.IsNull() && data.Ownergroup.ValueString() != "" {
		args = append(args, fmt.Sprintf("ownergroup:%s", utils.UrlEncode(data.Ownergroup.ValueString())))
	}
	if !data.Td.IsNull() {
		args = append(args, fmt.Sprintf("td:%v", data.Td.ValueInt64()))
	}

	err = r.client.DeleteResourceWithArgs(service.Vlan_nsip6_binding.Type(), vlanid, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vlan_nsip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vlan_nsip6_binding binding")
}

// Helper function to read vlan_nsip6_binding data from API
func (r *VlanNsip6BindingResource) readVlanNsip6BindingFromApi(ctx context.Context, data *VlanNsip6BindingResourceModel, diags *diag.Diagnostics) {

	// Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vlanid", "ipaddress"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	vlanid, ok := idMap["vlanid"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'vlanid' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vlan_nsip6_binding.Type(),
		ResourceName:             vlanid,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vlan_nsip6_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vlan_nsip6_binding returned empty array.")
		return
	}

	// Iterate through results to find the binding matching ipaddress (and vlanid).
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ipaddress
		if idVal, ok := idMap["ipaddress"]; ok {
			if val, ok := v["ipaddress"].(string); ok {
				if val != idVal {
					match = false
				}
			} else {
				match = false
			}
		}

		// Check vlanid against the NITRO "id" field
		if match {
			if val, ok := v["id"]; ok {
				valInt64, _ := utils.ConvertToInt64(val)
				idValInt64, _ := strconv.ParseInt(vlanid, 10, 64)
				if valInt64 != idValInt64 {
					match = false
				}
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", "vlan_nsip6_binding not found with the provided ID attributes")
		return
	}

	vlan_nsip6_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
