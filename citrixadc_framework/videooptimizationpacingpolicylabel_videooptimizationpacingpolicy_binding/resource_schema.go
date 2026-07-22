package videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NOTE: The NetScaler CLI marks the videooptimization pacing feature as deprecated.
// This binding resource is retained for backward compatibility.

// VideooptimizationpacingpolicylabelVideooptimizationpacingpolicyBindingResourceModel describes the resource data model.
type VideooptimizationpacingpolicylabelVideooptimizationpacingpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *VideooptimizationpacingpolicylabelVideooptimizationpacingpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				// GET echoes this with a server-assigned default ("END") even when unset,
				// so it must be Optional+Computed (read from response) to avoid an
				// "inconsistent result after apply" error.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				// GET echoes this with a server-assigned default (false) even when unset,
				// so it must be Optional+Computed (read from response).
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label and evaluate the specified policy label.",
			},
			"invoke_labelname": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "* If labelType is policylabel, name of the policy label to invoke.\n* If labelType is reqvserver or resvserver, name of the virtual server.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the videooptimization pacing policy label to which to bind the policy.",
			},
			"labeltype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of policy label to invoke. Available settings function as follows:\n* vserver - Invoke an unnamed policy label associated with a virtual server.\n* policylabel - Invoke a user-defined policy label.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the videooptimization policy.",
			},
			"priority": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}

func videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VideooptimizationpacingpolicylabelVideooptimizationpacingpolicyBindingResourceModel) videooptimization.Videooptimizationpacingpolicylabelvideooptimizationpacingpolicybinding {
	tflog.Debug(ctx, "In videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding := videooptimization.Videooptimizationpacingpolicylabelvideooptimizationpacingpolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.InvokeLabelname.IsNull() && !data.InvokeLabelname.IsUnknown() {
		videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.Invokelabelname = data.InvokeLabelname.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_binding
}

func videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingSetAttrFromGet(ctx context.Context, data *VideooptimizationpacingpolicylabelVideooptimizationpacingpolicyBindingResourceModel, getResponseData map[string]interface{}) *VideooptimizationpacingpolicylabelVideooptimizationpacingpolicyBindingResourceModel {
	tflog.Debug(ctx, "In videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model. Structurally identical to the detection label binding.
	//   Case A (GET echoes a server-assigned default): gotopriorityexpression ("END"),
	//     invoke (false). These are Optional+Computed and are read from the response.
	//   Case B (GET omits the field): invoke_labelname, labeltype. The `ok` guard below
	//     leaves the existing plan/state value untouched, avoiding a perpetual diff
	//     (Pattern 7).
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["invoke_labelname"]; ok && val != nil {
		data.InvokeLabelname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	return data
}

// videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingSetAttrFromGetForDatasource faithfully
// copies every field from the GET response (the datasource has no prior plan/state to
// preserve) and sets the composite ID.
func videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VideooptimizationpacingpolicylabelVideooptimizationpacingpolicyBindingResourceModel, getResponseData map[string]interface{}) *VideooptimizationpacingpolicylabelVideooptimizationpacingpolicyBindingResourceModel {
	tflog.Debug(ctx, "In videooptimizationpacingpolicylabel_videooptimizationpacingpolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	} else {
		data.Invoke = types.BoolNull()
	}
	if val, ok := getResponseData["invoke_labelname"]; ok && val != nil {
		data.InvokeLabelname = types.StringValue(val.(string))
	} else {
		data.InvokeLabelname = types.StringNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["labeltype"]; ok && val != nil {
		data.Labeltype = types.StringValue(val.(string))
	} else {
		data.Labeltype = types.StringNull()
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

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
