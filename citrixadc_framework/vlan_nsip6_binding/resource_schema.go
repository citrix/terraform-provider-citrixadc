package vlan_nsip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
				Required:    true,
				Description: "Specifies the virtual LAN ID.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IP address assigned to the VLAN.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet mask for the network address defined for this VLAN.",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DEFAULT_NG"),
				Description: "The owner node group in a Cluster for this vlan.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}

func vlan_nsip6_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VlanNsip6BindingResourceModel) network.Vlannsip6binding {
	tflog.Debug(ctx, "In vlan_nsip6_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vlan_nsip6_binding := network.Vlannsip6binding{}
	if !data.Vlanid.IsNull() {
		vlan_nsip6_binding.Id = utils.IntPtr(int(data.Vlanid.ValueInt64()))
	}
	if !data.Ipaddress.IsNull() {
		vlan_nsip6_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Netmask.IsNull() {
		vlan_nsip6_binding.Netmask = data.Netmask.ValueString()
	}
	if !data.Ownergroup.IsNull() {
		vlan_nsip6_binding.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Td.IsNull() {
		vlan_nsip6_binding.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return vlan_nsip6_binding
}

func vlan_nsip6_bindingSetAttrFromGet(ctx context.Context, data *VlanNsip6BindingResourceModel, getResponseData map[string]interface{}) *VlanNsip6BindingResourceModel {
	tflog.Debug(ctx, "In vlan_nsip6_bindingSetAttrFromGet Function")

	// Convert API response to model
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vlanid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlanid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipaddress.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("netmask:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Netmask.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ownergroup:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ownergroup.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("td:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Td.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
