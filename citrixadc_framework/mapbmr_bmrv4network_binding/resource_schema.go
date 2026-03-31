package mapbmr_bmrv4network_binding

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

// MapbmrBmrv4networkBindingResourceModel describes the resource data model.
type MapbmrBmrv4networkBindingResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Netmask types.String `tfsdk:"netmask"`
	Network types.String `tfsdk:"network"`
}

func (r *MapbmrBmrv4networkBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the mapbmr_bmrv4network_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Basic Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Basic Mapping Rule is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"add network MapBmr bmr1 -natprefix 2005::/64 -EAbitLength 16 -psidoffset 6 -portsharingratio 8\" ).\n			The Basic Mapping Rule information allows a MAP BR to determine source IPv4 address from the IPv6 packet sent from MAP CE device.\n			Also it allows to determine destination IPv6 address of MAP CE before sending packets to MAP CE",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet mask for the IPv4 address specified in the Network parameter.",
			},
			"network": schema.StringAttribute{
				Required:    true,
				Description: "IPv4 NAT address range of Customer Edge (CE).",
			},
		},
	}
}

func mapbmr_bmrv4network_bindingGetThePayloadFromtheConfig(ctx context.Context, data *MapbmrBmrv4networkBindingResourceModel) network.Mapbmrbmrv4networkbinding {
	tflog.Debug(ctx, "In mapbmr_bmrv4network_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	mapbmr_bmrv4network_binding := network.Mapbmrbmrv4networkbinding{}
	if !data.Name.IsNull() {
		mapbmr_bmrv4network_binding.Name = data.Name.ValueString()
	}
	if !data.Netmask.IsNull() {
		mapbmr_bmrv4network_binding.Netmask = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() {
		mapbmr_bmrv4network_binding.Network = data.Network.ValueString()
	}

	return mapbmr_bmrv4network_binding
}

func mapbmr_bmrv4network_bindingSetAttrFromGet(ctx context.Context, data *MapbmrBmrv4networkBindingResourceModel, getResponseData map[string]interface{}) *MapbmrBmrv4networkBindingResourceModel {
	tflog.Debug(ctx, "In mapbmr_bmrv4network_bindingSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["network"]; ok && val != nil {
		data.Network = types.StringValue(val.(string))
	} else {
		data.Network = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("network:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Network.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
