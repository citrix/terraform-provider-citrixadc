package nstrafficdomain_vlan_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NstrafficdomainVlanBindingResourceModel describes the resource data model.
type NstrafficdomainVlanBindingResourceModel struct {
	Id   types.String `tfsdk:"id"`
	Td   types.Int64  `tfsdk:"td"`
	Vlan types.Int64  `tfsdk:"vlan"`
}

func (r *NstrafficdomainVlanBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstrafficdomain_vlan_binding resource.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a traffic domain.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VLAN to bind to this traffic domain. More than one VLAN can be bound to a traffic domain, but the same VLAN cannot be a part of multiple traffic domains.",
			},
		},
	}
}

func nstrafficdomain_vlan_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NstrafficdomainVlanBindingResourceModel) ns.Nstrafficdomainvlanbinding {
	tflog.Debug(ctx, "In nstrafficdomain_vlan_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nstrafficdomain_vlan_binding := ns.Nstrafficdomainvlanbinding{}
	if !data.Td.IsNull() {
		nstrafficdomain_vlan_binding.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vlan.IsNull() {
		nstrafficdomain_vlan_binding.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return nstrafficdomain_vlan_binding
}

func nstrafficdomain_vlan_bindingSetAttrFromGet(ctx context.Context, data *NstrafficdomainVlanBindingResourceModel, getResponseData map[string]interface{}) *NstrafficdomainVlanBindingResourceModel {
	tflog.Debug(ctx, "In nstrafficdomain_vlan_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
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
	idParts = append(idParts, fmt.Sprintf("td:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Td.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("vlan:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Vlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
