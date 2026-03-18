package nstrafficdomain_bridgegroup_binding

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

// NstrafficdomainBridgegroupBindingResourceModel describes the resource data model.
type NstrafficdomainBridgegroupBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Bridgegroup types.Int64  `tfsdk:"bridgegroup"`
	Td          types.Int64  `tfsdk:"td"`
}

func (r *NstrafficdomainBridgegroupBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstrafficdomain_bridgegroup_binding resource.",
			},
			"bridgegroup": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the configured bridge to bind to this traffic domain. More than one bridge group can be bound to a traffic domain, but the same bridge group cannot be a part of multiple traffic domains.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a traffic domain.",
			},
		},
	}
}

func nstrafficdomain_bridgegroup_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NstrafficdomainBridgegroupBindingResourceModel) ns.Nstrafficdomainbridgegroupbinding {
	tflog.Debug(ctx, "In nstrafficdomain_bridgegroup_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nstrafficdomain_bridgegroup_binding := ns.Nstrafficdomainbridgegroupbinding{}
	if !data.Bridgegroup.IsNull() {
		nstrafficdomain_bridgegroup_binding.Bridgegroup = utils.IntPtr(int(data.Bridgegroup.ValueInt64()))
	}
	if !data.Td.IsNull() {
		nstrafficdomain_bridgegroup_binding.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return nstrafficdomain_bridgegroup_binding
}

func nstrafficdomain_bridgegroup_bindingSetAttrFromGet(ctx context.Context, data *NstrafficdomainBridgegroupBindingResourceModel, getResponseData map[string]interface{}) *NstrafficdomainBridgegroupBindingResourceModel {
	tflog.Debug(ctx, "In nstrafficdomain_bridgegroup_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bridgegroup"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bridgegroup = types.Int64Value(intVal)
		}
	} else {
		data.Bridgegroup = types.Int64Null()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bridgegroup:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Bridgegroup.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("td:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Td.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
