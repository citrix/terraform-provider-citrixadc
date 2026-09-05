package bridgegroup_vlan_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// BridgegroupVlanBindingResourceModel describes the resource data model.
type BridgegroupVlanBindingResourceModel struct {
	Id            types.String `tfsdk:"id"`
	BridgegroupId types.Int64  `tfsdk:"bridgegroup_id"`
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
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The integer that uniquely identifies the bridge group.",
			},
			"vlan": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Names of all member VLANs.",
			},
		},
	}
}

func bridgegroup_vlan_bindingGetThePayloadFromthePlan(ctx context.Context, data *BridgegroupVlanBindingResourceModel) network.Bridgegroupvlanbinding {
	tflog.Debug(ctx, "In bridgegroup_vlan_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	bridgegroup_vlan_binding := network.Bridgegroupvlanbinding{}
	if !data.BridgegroupId.IsNull() && !data.BridgegroupId.IsUnknown() {
		bridgegroup_vlan_binding.Id = utils.IntPtr(int(data.BridgegroupId.ValueInt64()))
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		bridgegroup_vlan_binding.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}

	return bridgegroup_vlan_binding
}

// bridgegroup_vlan_bindingSetAttrFromGet preserves the user-configured key
// attributes (resource flow) and recomputes the composite ID. Because the
// NITRO GET returns the bridge group integer under the key "id" and the vlan
// under "vlan", the datasource path uses a dedicated setter below.
func bridgegroup_vlan_bindingSetAttrFromGet(ctx context.Context, data *BridgegroupVlanBindingResourceModel, getResponseData map[string]interface{}) *BridgegroupVlanBindingResourceModel {
	tflog.Debug(ctx, "In bridgegroup_vlan_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.BridgegroupId = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	}

	// Set ID for the resource
	// Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bridgegroup_id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BridgegroupId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("vlan:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// bridgegroup_vlan_bindingSetAttrFromGetForDatasource faithfully copies every
// field from the GET response and sets the composite ID, since the datasource
// has no prior plan/state to preserve.
func bridgegroup_vlan_bindingSetAttrFromGetForDatasource(ctx context.Context, data *BridgegroupVlanBindingResourceModel, getResponseData map[string]interface{}) *BridgegroupVlanBindingResourceModel {
	tflog.Debug(ctx, "In bridgegroup_vlan_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.BridgegroupId = types.Int64Value(intVal)
		} else {
			data.BridgegroupId = types.Int64Null()
		}
	} else {
		data.BridgegroupId = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		} else {
			data.Vlan = types.Int64Null()
		}
	} else {
		data.Vlan = types.Int64Null()
	}

	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bridgegroup_id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BridgegroupId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("vlan:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlan.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
