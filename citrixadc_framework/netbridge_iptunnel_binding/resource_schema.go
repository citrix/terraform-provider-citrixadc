package netbridge_iptunnel_binding

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

// NetbridgeIptunnelBindingResourceModel describes the resource data model.
type NetbridgeIptunnelBindingResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Tunnel types.String `tfsdk:"tunnel"`
}

func (r *NetbridgeIptunnelBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the netbridge_iptunnel_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the network bridge.",
			},
			"tunnel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the tunnel that is a part of this bridge.",
			},
		},
	}
}

func netbridge_iptunnel_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NetbridgeIptunnelBindingResourceModel) network.Netbridgeiptunnelbinding {
	tflog.Debug(ctx, "In netbridge_iptunnel_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	netbridge_iptunnel_binding := network.Netbridgeiptunnelbinding{}
	if !data.Name.IsNull() {
		netbridge_iptunnel_binding.Name = data.Name.ValueString()
	}
	if !data.Tunnel.IsNull() {
		netbridge_iptunnel_binding.Tunnel = data.Tunnel.ValueString()
	}

	return netbridge_iptunnel_binding
}

func netbridge_iptunnel_bindingSetAttrFromGet(ctx context.Context, data *NetbridgeIptunnelBindingResourceModel, getResponseData map[string]interface{}) *NetbridgeIptunnelBindingResourceModel {
	tflog.Debug(ctx, "In netbridge_iptunnel_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["tunnel"]; ok && val != nil {
		data.Tunnel = types.StringValue(val.(string))
	} else {
		data.Tunnel = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("tunnel:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Tunnel.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
