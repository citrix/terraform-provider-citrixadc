package nstrafficdomain_vxlan_binding

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

// NstrafficdomainVxlanBindingResourceModel describes the resource data model.
type NstrafficdomainVxlanBindingResourceModel struct {
	Id    types.String `tfsdk:"id"`
	Td    types.Int64  `tfsdk:"td"`
	Vxlan types.Int64  `tfsdk:"vxlan"`
}

func (r *NstrafficdomainVxlanBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstrafficdomain_vxlan_binding resource.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a traffic domain.",
			},
			"vxlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the VXLAN to bind to this traffic domain. More than one VXLAN can be bound to a traffic domain, but the same VXLAN cannot be a part of multiple traffic domains.",
			},
		},
	}
}

func nstrafficdomain_vxlan_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NstrafficdomainVxlanBindingResourceModel) ns.Nstrafficdomainvxlanbinding {
	tflog.Debug(ctx, "In nstrafficdomain_vxlan_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nstrafficdomain_vxlan_binding := ns.Nstrafficdomainvxlanbinding{}
	if !data.Td.IsNull() {
		nstrafficdomain_vxlan_binding.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Vxlan.IsNull() {
		nstrafficdomain_vxlan_binding.Vxlan = utils.IntPtr(int(data.Vxlan.ValueInt64()))
	}

	return nstrafficdomain_vxlan_binding
}

func nstrafficdomain_vxlan_bindingSetAttrFromGet(ctx context.Context, data *NstrafficdomainVxlanBindingResourceModel, getResponseData map[string]interface{}) *NstrafficdomainVxlanBindingResourceModel {
	tflog.Debug(ctx, "In nstrafficdomain_vxlan_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["vxlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vxlan = types.Int64Value(intVal)
		}
	} else {
		data.Vxlan = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("td:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Td.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("vxlan:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vxlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
