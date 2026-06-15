package feoglobal_feopolicy_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/feo"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// FeoglobalFeopolicyBindingResourceModel describes the resource data model.
type FeoglobalFeopolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
}

func (r *FeoglobalFeopolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the feoglobal_feopolicy_binding resource.",
			},
			"globalbindtype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the globally bound front end optimization policy.",
			},
			"priority": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The priority assigned to the policy binding.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bindpoint to which the policy is bound.",
			},
		},
	}
}

func feoglobal_feopolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *FeoglobalFeopolicyBindingResourceModel) feo.Feoglobalfeopolicybinding {
	tflog.Debug(ctx, "In feoglobal_feopolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	feoglobal_feopolicy_binding := feo.Feoglobalfeopolicybinding{}
	if !data.Globalbindtype.IsNull() && !data.Globalbindtype.IsUnknown() {
		feoglobal_feopolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		feoglobal_feopolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		feoglobal_feopolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		feoglobal_feopolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		feoglobal_feopolicy_binding.Type = data.Type.ValueString()
	}

	return feoglobal_feopolicy_binding
}

// feoglobal_feopolicy_bindingSetAttrFromGet is the RESOURCE setter. The identity
// inputs (policyname, priority, type) are RequiresReplace user inputs; we preserve
// the values already in state/plan rather than overwriting them with the GET
// response so the post-apply state matches the user config (avoids
// "inconsistent result after apply"). Only the computed read-back fields
// (globalbindtype, gotopriorityexpression) are populated from the response. The ID
// is set exactly once in Create, so it is not recomputed here.
func feoglobal_feopolicy_bindingSetAttrFromGet(ctx context.Context, data *FeoglobalFeopolicyBindingResourceModel, getResponseData map[string]interface{}) *FeoglobalFeopolicyBindingResourceModel {
	tflog.Debug(ctx, "In feoglobal_feopolicy_bindingSetAttrFromGet Function")

	// Computed read-back fields - take from the GET response.
	if val, ok := getResponseData["globalbindtype"]; ok && val != nil {
		data.Globalbindtype = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}

	// policyname, priority and type are RequiresReplace identity inputs - preserved
	// from prior state/plan (not overwritten by the GET response).

	return data
}

// feoglobal_feopolicy_bindingSetAttrFromGetForDatasource is the DATASOURCE setter.
// The datasource has no prior plan/state to preserve, so it faithfully copies every
// field from the GET response and sets the ID (plain policyname, matching Create).
func feoglobal_feopolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *FeoglobalFeopolicyBindingResourceModel, getResponseData map[string]interface{}) *FeoglobalFeopolicyBindingResourceModel {
	tflog.Debug(ctx, "In feoglobal_feopolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["globalbindtype"]; ok && val != nil {
		data.Globalbindtype = types.StringValue(val.(string))
	} else {
		data.Globalbindtype = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// ID is the plain policyname (backward-compatible with SDK v2).
	data.Id = types.StringValue(data.Policyname.ValueString())

	return data
}
