package bridgegroup_vlan_binding

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

// BridgegroupVlanBindingResourceModel describes the resource data model.
type BridgegroupVlanBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	Bridgegroupid types.Int64  `tfsdk:"bridgegroup_id"`
	Vlan          types.Int64  `tfsdk:"vlan"`
}

func (r *BridgegroupVlanBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the bridgegroup_vlan_binding resource.",
			},
			"bridgegroup_id": schema.Int64Attribute{
				Required:    true,
				Description: "The integer that uniquely identifies the bridge group.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Names of all member VLANs.",
			},
		},
	}
}

func bridgegroup_vlan_bindingGetThePayloadFromtheConfig(ctx context.Context, data *BridgegroupVlanBindingResourceModel) network.Bridgegroupvlanbinding {
	tflog.Debug(ctx, "In bridgegroup_vlan_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	bridgegroup_vlan_binding := network.Bridgegroupvlanbinding{}
	if !data.Bridgegroupid.IsNull() {
		bridgegroup_vlan_binding.Id = utils.IntPtr(int(data.Bridgegroupid.ValueInt64()))
	}
	if !data.Vlan.IsNull() {
		bridgegroup_vlan_binding.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return bridgegroup_vlan_binding
}

func bridgegroup_vlan_bindingSetAttrFromGet(ctx context.Context, data *BridgegroupVlanBindingResourceModel, getResponseData map[string]interface{}) *BridgegroupVlanBindingResourceModel {
	tflog.Debug(ctx, "In bridgegroup_vlan_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bridgegroupid = types.Int64Value(intVal)
		}
	} else {
		data.Bridgegroupid = types.Int64Null()
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
	idParts = append(idParts, fmt.Sprintf("bridgegroup_id:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Bridgegroupid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("vlan:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Vlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
