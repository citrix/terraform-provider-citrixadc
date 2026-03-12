package vpnvserver_intranetip_binding

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

// VpnvserverIntranetipBindingResourceModel describes the resource data model.
type VpnvserverIntranetipBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Intranetip types.String `tfsdk:"intranetip"`
	Name       types.String `tfsdk:"name"`
	Netmask    types.String `tfsdk:"netmask"`
}

func (r *VpnvserverIntranetipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver_intranetip_binding resource.",
			},
			"intranetip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The network ID for the range of intranet IP addresses or individual intranet IP addresses to be bound to the virtual server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the virtual server.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The netmask of the intranet IP address or range.",
			},
		},
	}
}

func vpnvserver_intranetip_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverIntranetipBindingResourceModel) vpn.Vpnvserverintranetipbinding {
	tflog.Debug(ctx, "In vpnvserver_intranetip_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver_intranetip_binding := vpn.Vpnvserverintranetipbinding{}
	if !data.Intranetip.IsNull() {
		vpnvserver_intranetip_binding.Intranetip = data.Intranetip.ValueString()
	}
	if !data.Name.IsNull() {
		vpnvserver_intranetip_binding.Name = data.Name.ValueString()
	}
	if !data.Netmask.IsNull() {
		vpnvserver_intranetip_binding.Netmask = data.Netmask.ValueString()
	}

	return vpnvserver_intranetip_binding
}

func vpnvserver_intranetip_bindingSetAttrFromGet(ctx context.Context, data *VpnvserverIntranetipBindingResourceModel, getResponseData map[string]interface{}) *VpnvserverIntranetipBindingResourceModel {
	tflog.Debug(ctx, "In vpnvserver_intranetip_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["intranetip"]; ok && val != nil {
		data.Intranetip = types.StringValue(val.(string))
	} else {
		data.Intranetip = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("intranetip:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Intranetip.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
