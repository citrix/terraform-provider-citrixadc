package netbridge_vlan_binding

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

// NetbridgeVlanBindingResourceModel describes the resource data model.
type NetbridgeVlanBindingResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Vlan types.Int64  `tfsdk:"vlan"`
}

func (r *NetbridgeVlanBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the netbridge_vlan_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the network bridge.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The VLAN that is extended by this network bridge.",
			},
		},
	}
}

func netbridge_vlan_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NetbridgeVlanBindingResourceModel) network.Netbridgevlanbinding {
	tflog.Debug(ctx, "In netbridge_vlan_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	netbridge_vlan_binding := network.Netbridgevlanbinding{}
	if !data.Name.IsNull() {
		netbridge_vlan_binding.Name = data.Name.ValueString()
	}
	if !data.Vlan.IsNull() {
		netbridge_vlan_binding.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return netbridge_vlan_binding
}

func netbridge_vlan_bindingSetAttrFromGet(ctx context.Context, data *NetbridgeVlanBindingResourceModel, getResponseData map[string]interface{}) *NetbridgeVlanBindingResourceModel {
	tflog.Debug(ctx, "In netbridge_vlan_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vlan:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
