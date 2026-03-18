package nd6ravariables_onlinkipv6prefix_binding

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

// Nd6ravariablesOnlinkipv6prefixBindingResourceModel describes the resource data model.
type Nd6ravariablesOnlinkipv6prefixBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Ipv6prefix types.String `tfsdk:"ipv6prefix"`
	Vlan       types.Int64  `tfsdk:"vlan"`
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nd6ravariables_onlinkipv6prefix_binding resource.",
			},
			"ipv6prefix": schema.StringAttribute{
				Required:    true,
				Description: "Onlink prefixes for RA messages.",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "The VLAN number.",
			},
		},
	}
}

func nd6ravariables_onlinkipv6prefix_bindingGetThePayloadFromtheConfig(ctx context.Context, data *Nd6ravariablesOnlinkipv6prefixBindingResourceModel) network.Nd6ravariablesonlinkipv6prefixbinding {
	tflog.Debug(ctx, "In nd6ravariables_onlinkipv6prefix_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nd6ravariables_onlinkipv6prefix_binding := network.Nd6ravariablesonlinkipv6prefixbinding{}
	if !data.Ipv6prefix.IsNull() {
		nd6ravariables_onlinkipv6prefix_binding.Ipv6prefix = data.Ipv6prefix.ValueString()
	}
	if !data.Vlan.IsNull() {
		nd6ravariables_onlinkipv6prefix_binding.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return nd6ravariables_onlinkipv6prefix_binding
}

func nd6ravariables_onlinkipv6prefix_bindingSetAttrFromGet(ctx context.Context, data *Nd6ravariablesOnlinkipv6prefixBindingResourceModel, getResponseData map[string]interface{}) *Nd6ravariablesOnlinkipv6prefixBindingResourceModel {
	tflog.Debug(ctx, "In nd6ravariables_onlinkipv6prefix_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipv6prefix"]; ok && val != nil {
		data.Ipv6prefix = types.StringValue(val.(string))
	} else {
		data.Ipv6prefix = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ipv6prefix:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Ipv6prefix.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vlan:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Vlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
