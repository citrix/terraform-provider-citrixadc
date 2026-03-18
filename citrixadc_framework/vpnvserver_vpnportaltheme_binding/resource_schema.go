package vpnvserver_vpnportaltheme_binding

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

// VpnvserverVpnportalthemeBindingResourceModel describes the resource data model.
type VpnvserverVpnportalthemeBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Portaltheme types.String `tfsdk:"portaltheme"`
}

func (r *VpnvserverVpnportalthemeBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_vpnportaltheme_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"portaltheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the portal theme bound to VPN vserver",
			},
		},
	}
}

func vpnvserver_vpnportaltheme_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverVpnportalthemeBindingResourceModel) vpn.Vpnvservervpnportalthemebinding {
	tflog.Debug(ctx, "In vpnvserver_vpnportaltheme_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver_vpnportaltheme_binding := vpn.Vpnvservervpnportalthemebinding{}
	if !data.Name.IsNull() {
		vpnvserver_vpnportaltheme_binding.Name = data.Name.ValueString()
	}
	if !data.Portaltheme.IsNull() {
		vpnvserver_vpnportaltheme_binding.Portaltheme = data.Portaltheme.ValueString()
	}

	return vpnvserver_vpnportaltheme_binding
}

func vpnvserver_vpnportaltheme_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverVpnportalthemeBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverVpnportalthemeBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_vpnportaltheme_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["portaltheme"]; ok && val != nil {
		data.Portaltheme = types.StringValue(val.(string))
	} else {
		data.Portaltheme = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("portaltheme:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Portaltheme.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
