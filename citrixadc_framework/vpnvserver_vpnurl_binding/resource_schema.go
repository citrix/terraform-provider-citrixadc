package vpnvserver_vpnurl_binding

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

// VpnvserverVpnurlBindingResourceModel describes the resource data model.
type VpnvserverVpnurlBindingResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Urlname types.String `tfsdk:"urlname"`
}

func (r *VpnvserverVpnurlBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_vpnurl_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"urlname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The intranet URL.",
			},
		},
	}
}

func vpnvserver_vpnurl_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverVpnurlBindingResourceModel) vpn.Vpnvservervpnurlbinding {
	tflog.Debug(ctx, "In vpnvserver_vpnurl_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver_vpnurl_binding := vpn.Vpnvservervpnurlbinding{}
	if !data.Name.IsNull() {
		vpnvserver_vpnurl_binding.Name = data.Name.ValueString()
	}
	if !data.Urlname.IsNull() {
		vpnvserver_vpnurl_binding.Urlname = data.Urlname.ValueString()
	}

	return vpnvserver_vpnurl_binding
}

func vpnvserver_vpnurl_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverVpnurlBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverVpnurlBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_vpnurl_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["urlname"]; ok && val != nil {
		data.Urlname = types.StringValue(val.(string))
	} else {
		data.Urlname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("urlname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Urlname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
