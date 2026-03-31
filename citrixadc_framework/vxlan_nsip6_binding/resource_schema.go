package vxlan_nsip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VxlanNsip6BindingResourceModel describes the resource data model.
type VxlanNsip6BindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Vxlanid   types.Int64  `tfsdk:"vxlanid"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Netmask   types.String `tfsdk:"netmask"`
}

func (r *VxlanNsip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vxlan_nsip6_binding resource.",
			},
			"vxlanid": schema.Int64Attribute{
				Required:    true,
				Description: "A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IP address assigned to the VXLAN.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet mask for the network address defined for this VXLAN.",
			},
		},
	}
}

func vxlan_nsip6_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VxlanNsip6BindingResourceModel) network.Vxlannsip6binding {
	tflog.Debug(ctx, "In vxlan_nsip6_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vxlan_nsip6_binding := network.Vxlannsip6binding{}
	if !data.Vxlanid.IsNull() {
		vxlan_nsip6_binding.Id = utils.IntPtr(int(data.Vxlanid.ValueInt64()))
	}
	if !data.Ipaddress.IsNull() {
		vxlan_nsip6_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Netmask.IsNull() {
		vxlan_nsip6_binding.Netmask = data.Netmask.ValueString()
	}

	return vxlan_nsip6_binding
}

func vxlan_nsip6_bindingSetAttrFromGet(ctx context.Context, data *VxlanNsip6BindingResourceModel, getResponseData map[string]interface{}) *VxlanNsip6BindingResourceModel {
	tflog.Debug(ctx, "In vxlan_nsip6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlanid = types.Int64Value(intVal)
		}
	} else {
		data.Vxlanid = types.Int64Null()
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vxlanid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vxlanid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipaddress.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("netmask:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Netmask.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
