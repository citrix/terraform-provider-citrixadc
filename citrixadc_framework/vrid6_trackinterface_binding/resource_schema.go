package vrid6_trackinterface_binding

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

// Vrid6TrackinterfaceBindingResourceModel describes the resource data model.
// The NITRO key attribute is named "id"; it is exposed in Terraform as "vrid_id"
// so it does not collide with the synthetic resource ID attribute "id".
type Vrid6TrackinterfaceBindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	VridId     types.Int64  `tfsdk:"vrid_id"`
	Trackifnum types.String `tfsdk:"trackifnum"`
}

func (r *Vrid6TrackinterfaceBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vrid6_trackinterface_binding resource.",
			},
			"vrid_id": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies a VMAC6 address.",
			},
			"trackifnum": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interfaces which need to be tracked for this vrID.",
			},
		},
	}
}

func vrid6_trackinterface_bindingGetThePayloadFromthePlan(ctx context.Context, data *Vrid6TrackinterfaceBindingResourceModel) network.Vrid6trackinterfacebinding {
	tflog.Debug(ctx, "In vrid6_trackinterface_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vrid6_trackinterface_binding := network.Vrid6trackinterfacebinding{}
	if !data.VridId.IsNull() && !data.VridId.IsUnknown() {
		vrid6_trackinterface_binding.Id = utils.IntPtr(int(data.VridId.ValueInt64()))
	}
	if !data.Trackifnum.IsNull() && !data.Trackifnum.IsUnknown() {
		vrid6_trackinterface_binding.Trackifnum = data.Trackifnum.ValueString()
	}

	return vrid6_trackinterface_binding
}

// vrid6_trackinterface_bindingSetAttrFromGet is used by the resource Read/Create flow.
// It preserves the plan/state-supplied identity values (vrid_id, trackifnum are both
// RequiresReplace) and does NOT recompute the ID, which is set exactly once in Create.
func vrid6_trackinterface_bindingSetAttrFromGet(ctx context.Context, data *Vrid6TrackinterfaceBindingResourceModel, getResponseData map[string]interface{}) *Vrid6TrackinterfaceBindingResourceModel {
	tflog.Debug(ctx, "In vrid6_trackinterface_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.VridId = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["trackifnum"]; ok && val != nil {
		data.Trackifnum = types.StringValue(val.(string))
	}

	return data
}

// vrid6_trackinterface_bindingSetAttrFromGetForDatasource faithfully copies every
// field from the GET response (including the read-only flags) and composes the ID,
// since the datasource has no Create to seed those values. Note: trackinterface
// bindings have no vlan field.
func vrid6_trackinterface_bindingSetAttrFromGetForDatasource(ctx context.Context, data *Vrid6TrackinterfaceBindingDataSourceModel, getResponseData map[string]interface{}) *Vrid6TrackinterfaceBindingDataSourceModel {
	tflog.Debug(ctx, "In vrid6_trackinterface_bindingSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.VridId = types.Int64Value(intVal)
		}
	} else {
		data.VridId = types.Int64Null()
	}
	if val, ok := getResponseData["trackifnum"]; ok && val != nil {
		data.Trackifnum = types.StringValue(val.(string))
	} else {
		data.Trackifnum = types.StringNull()
	}
	if val, ok := getResponseData["flags"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Flags = types.Int64Value(intVal)
		}
	} else {
		data.Flags = types.Int64Null()
	}

	// Set ID for the datasource
	// Composite key: id,trackifnum (key:UrlEncode(value) pairs)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.VridId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("trackifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Trackifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
