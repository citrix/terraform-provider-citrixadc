package vrid6_channel_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Vrid6ChannelBindingResourceModel describes the resource data model.
// The NITRO key attribute is named "id"; it is exposed in Terraform as "vrid_id"
// so it does not collide with the synthetic resource ID attribute "id".
type Vrid6ChannelBindingResourceModel struct {
	Id     types.String `tfsdk:"id"`
	VridId types.Int64  `tfsdk:"vrid_id"`
	Ifnum  types.String `tfsdk:"ifnum"`
}

func (r *Vrid6ChannelBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vrid6_channel_binding resource.",
			},
			"vrid_id": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies a VMAC6 address.",
			},
			"ifnum": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interfaces to bind to the VMAC6, specified in (slot/port) notation (for example, 1/2).Use spaces to separate multiple entries.",
			},
		},
	}
}

func vrid6_channel_bindingGetThePayloadFromthePlan(ctx context.Context, data *Vrid6ChannelBindingResourceModel) network.Vrid6channelbinding {
	tflog.Debug(ctx, "In vrid6_channel_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vrid6_channel_binding := network.Vrid6channelbinding{}
	if !data.VridId.IsNull() && !data.VridId.IsUnknown() {
		vrid6_channel_binding.Id = utils.IntPtr(int(data.VridId.ValueInt64()))
	}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		vrid6_channel_binding.Ifnum = data.Ifnum.ValueString()
	}

	return vrid6_channel_binding
}

// vrid6_channel_bindingSetAttrFromGet is used by the resource Read/Create flow.
// It preserves the plan/state-supplied identity values (vrid_id, ifnum are both
// RequiresReplace) and does NOT recompute the ID, which is set exactly once in Create.
func vrid6_channel_bindingSetAttrFromGet(ctx context.Context, data *Vrid6ChannelBindingResourceModel, getResponseData map[string]interface{}) *Vrid6ChannelBindingResourceModel {
	tflog.Debug(ctx, "In vrid6_channel_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.VridId = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	}

	return data
}

// vrid6_channel_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response (including the read-only flags/vlan) and composes the ID,
// since the datasource has no Create to seed those values.
func vrid6_channel_bindingSetAttrFromGetForDatasource(ctx context.Context, data *Vrid6ChannelBindingDataSourceModel, getResponseData map[string]interface{}) *Vrid6ChannelBindingDataSourceModel {
	tflog.Debug(ctx, "In vrid6_channel_bindingSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.VridId = types.Int64Value(intVal)
		}
	} else {
		data.VridId = types.Int64Null()
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		if s, ok := val.(string); ok {
			data.Ifnum = types.StringValue(s)
		}
	}
	// NOTE: the firmware does NOT echo "ifnum" for these rows; when absent we RETAIN
	// the config-supplied ifnum (the datasource lookup key) instead of nulling it,
	// so the composite ID and the ifnum output stay correct.
	if val, ok := getResponseData["flags"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Flags = types.Int64Value(intVal)
		}
	} else {
		data.Flags = types.Int64Null()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else {
		data.Vlan = types.Int64Null()
	}

	// Set ID for the datasource
	// Composite key: id,ifnum (key:UrlEncode(value) pairs)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.VridId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
