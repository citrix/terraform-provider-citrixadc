package vlan_nsip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VlanNsip6BindingResourceModel describes the resource data model.
type VlanNsip6BindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Vlanid     types.Int64  `tfsdk:"vlanid"`
	Ipaddress  types.String `tfsdk:"ipaddress"`
	Netmask    types.String `tfsdk:"netmask"`
	Ownergroup types.String `tfsdk:"ownergroup"`
	Td         types.Int64  `tfsdk:"td"`
}

func (r *VlanNsip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vlan_nsip6_binding resource.",
			},
			"vlanid": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specifies the virtual LAN ID.",
			},
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IP address assigned to the VLAN.",
			},
			// netmask/ownergroup/td are Optional only (no Computed): the
			// vlan_nsip6_binding GET endpoint never echoes them back, so a Computed
			// flag would leave them unknown after apply ("invalid result object").
			"netmask": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet mask for the network address defined for this VLAN.",
			},
			"ownergroup": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The owner node group in a Cluster for this vlan.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}

func vlan_nsip6_bindingGetThePayloadFromthePlan(ctx context.Context, data *VlanNsip6BindingResourceModel) network.Vlannsip6binding {
	tflog.Debug(ctx, "In vlan_nsip6_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model. The NITRO field for the VLAN id is "id".
	vlan_nsip6_binding := network.Vlannsip6binding{}
	if !data.Vlanid.IsNull() && !data.Vlanid.IsUnknown() {
		vlan_nsip6_binding.Id = utils.IntPtr(int(data.Vlanid.ValueInt64()))
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		vlan_nsip6_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		vlan_nsip6_binding.Netmask = data.Netmask.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		vlan_nsip6_binding.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		vlan_nsip6_binding.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return vlan_nsip6_binding
}

// vlan_nsip6_bindingComposeId builds the composite key:value ID from the model.
// Matches the SDK v2 positional order (vlanid,ipaddress) so ParseIdString can
// decode both the new key:value form and legacy comma-separated imports.
func vlan_nsip6_bindingComposeId(data *VlanNsip6BindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vlanid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlanid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipaddress.ValueString()))))
	return strings.Join(idParts, ",")
}

// vlan_nsip6_bindingSetAttrFromGet updates the resource state from the GET
// response while preserving the composite ID (set once in Create).
func vlan_nsip6_bindingSetAttrFromGet(ctx context.Context, data *VlanNsip6BindingResourceModel, getResponseData map[string]interface{}) *VlanNsip6BindingResourceModel {
	tflog.Debug(ctx, "In vlan_nsip6_bindingSetAttrFromGet Function")

	// Convert API response to model. The NITRO field for the VLAN id is "id".
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlanid = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	}

	// Re-derive the canonical id so a legacy SDK v2 id is upgraded to the new format on Read.
	data.Id = types.StringValue(vlan_nsip6_bindingComposeId(data))

	return data
}

// vlan_nsip6_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and sets the composite ID, since the datasource has no
// Create to establish state.
func vlan_nsip6_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VlanNsip6BindingResourceModel, getResponseData map[string]interface{}) *VlanNsip6BindingResourceModel {
	tflog.Debug(ctx, "In vlan_nsip6_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlanid = types.Int64Value(intVal)
		}
	} else {
		data.Vlanid = types.Int64Null()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}

	data.Id = types.StringValue(vlan_nsip6_bindingComposeId(data))

	return data
}
