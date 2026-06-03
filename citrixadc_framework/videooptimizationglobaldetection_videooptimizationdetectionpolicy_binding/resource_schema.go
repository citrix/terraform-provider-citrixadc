package videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding

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

// VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel describes the resource data model.
type VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
}

func (r *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding resource.",
			},
			"globalbindtype": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "0",
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
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or\nevaluate the specified policy label.",
			},
			"labelname": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel.",
			},
			"labeltype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of invocation, Available settings function as follows:\n* vserver - Forward the request to the specified virtual server.\n* policylabel - Invoke the specified policy label.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the videooptimization detection policy.",
			},
			"priority": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specifies the bind point whose policies you want to display.",
			},
		},
	}
}

func videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel) videooptimization.Videooptimizationglobaldetectionvideooptimizationdetectionpolicybinding {
	tflog.Debug(ctx, "In videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding := videooptimization.Videooptimizationglobaldetectionvideooptimizationdetectionpolicybinding{}
	// globalbindtype is a read-only response field (default SYSTEM_GLOBAL); it is not
	// accepted in the bind/add payload (Pattern 15 sanctioned exception) so it is excluded.
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.Type = data.Type.ValueString()
	}

	return videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding
}

func videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingSetAttrFromGet(ctx context.Context, data *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel, getResponseData map[string]interface{}) *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel {
	tflog.Debug(ctx, "In videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model. Per live GET inspection of this global binding:
	//   Case A (GET echoes a server-assigned default): gotopriorityexpression ("END") and
	//     type ("RES_DEFAULT" bindpoint). Both are Optional+Computed and read from response.
	//   Case B (GET omits the field): invoke, labelname, labeltype. The `ok` guard below
	//     leaves the existing plan/state value untouched, avoiding a perpetual diff
	//     (Pattern 7).
	// globalbindtype is a Computed read-only field ("SYSTEM_GLOBAL"), always read.
	if val, ok := getResponseData["globalbindtype"]; ok && val != nil {
		data.Globalbindtype = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	}

	// NOTE: ID is set once in Create; it is intentionally not recomputed here to avoid
	// wiping it when a key field is absent from the GET response (Pattern 6).

	return data
}

// videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingSetAttrFromGetForDatasource faithfully
// copies every field from the GET response (the datasource has no prior plan/state to
// preserve) and sets the composite ID (Pattern 7 datasource regression guard).
func videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel, getResponseData map[string]interface{}) *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingResourceModel {
	tflog.Debug(ctx, "In videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingSetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["invoke"]; ok && val != nil {
		data.Invoke = types.BoolValue(val.(bool))
	} else {
		data.Invoke = types.BoolNull()
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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
