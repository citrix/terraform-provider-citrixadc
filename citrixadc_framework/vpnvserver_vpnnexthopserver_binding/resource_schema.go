package vpnvserver_vpnnexthopserver_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnvserverVpnnexthopserverBindingResourceModel describes the resource data model.
type VpnvserverVpnnexthopserverBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	Nexthopserver types.String `tfsdk:"nexthopserver"`
}

func (r *VpnvserverVpnnexthopserverBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_vpnnexthopserver_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"nexthopserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the next hop server bound to the VPN virtual server.",
			},
		},
	}
}

func vpnvserver_vpnnexthopserver_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverVpnnexthopserverBindingResourceModel) vpn.Vpnvservervpnnexthopserverbinding {
	tflog.Debug(ctx, "In vpnvserver_vpnnexthopserver_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver_vpnnexthopserver_binding := vpn.Vpnvservervpnnexthopserverbinding{}
	if !data.Name.IsNull() {
		vpnvserver_vpnnexthopserver_binding.Name = data.Name.ValueString()
	}
	if !data.Nexthopserver.IsNull() {
		vpnvserver_vpnnexthopserver_binding.Nexthopserver = data.Nexthopserver.ValueString()
	}

	return vpnvserver_vpnnexthopserver_binding
}

func vpnvserver_vpnnexthopserver_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverVpnnexthopserverBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverVpnnexthopserverBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_vpnnexthopserver_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["nexthopserver"]; ok && val != nil {
		data.Nexthopserver = types.StringValue(val.(string))
	} else {
		data.Nexthopserver = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("nexthopserver:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Nexthopserver.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
