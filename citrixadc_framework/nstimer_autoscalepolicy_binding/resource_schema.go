package nstimer_autoscalepolicy_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NstimerAutoscalepolicyBindingResourceModel describes the resource data model.
type NstimerAutoscalepolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Name                   types.String `tfsdk:"name"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Samplesize             types.Int64  `tfsdk:"samplesize"`
	Threshold              types.Int64  `tfsdk:"threshold"`
	Vserver                types.String `tfsdk:"vserver"`
}

func (r *NstimerAutoscalepolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nstimer_autoscalepolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Timer name.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The timer policy associated with the timer.",
			},
			"priority": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Specifies the priority of the timer policy.",
			},
			"samplesize": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Denotes the sample size. Sample size value of 'x' means that previous '(x - 1)' policy's rule evaluation results and the current evaluation result are present with the binding. For example, sample size of 10 means that there is a state of previous 9 policy evaluation results and also the current policy evaluation result.",
			},
			"threshold": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Denotes the threshold. If the rule of the policy in the binding relation evaluates 'threshold size' number of times in 'sample size' to true, then the corresponding action is taken. Its value needs to be less than or equal to the sample size value.",
			},
			"vserver": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the vserver which provides the context for the rule in timer policy. When not specified it is treated as a Global Default context.",
			},
		},
	}
}

func nstimer_autoscalepolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *NstimerAutoscalepolicyBindingResourceModel) ns.Nstimerautoscalepolicybinding {
	tflog.Debug(ctx, "In nstimer_autoscalepolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nstimer_autoscalepolicy_binding := ns.Nstimerautoscalepolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		nstimer_autoscalepolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nstimer_autoscalepolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		nstimer_autoscalepolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		nstimer_autoscalepolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Samplesize.IsNull() && !data.Samplesize.IsUnknown() {
		nstimer_autoscalepolicy_binding.Samplesize = utils.IntPtr(int(data.Samplesize.ValueInt64()))
	}
	if !data.Threshold.IsNull() && !data.Threshold.IsUnknown() {
		nstimer_autoscalepolicy_binding.Threshold = utils.IntPtr(int(data.Threshold.ValueInt64()))
	}
	if !data.Vserver.IsNull() && !data.Vserver.IsUnknown() {
		nstimer_autoscalepolicy_binding.Vserver = data.Vserver.ValueString()
	}

	return nstimer_autoscalepolicy_binding
}

// nstimer_autoscalepolicy_bindingSetAttrFromGet populates the resource model
// from a GET response. The non-Computed Optional attributes are RequiresReplace
// and are preserved from the existing plan/state to avoid "inconsistent result
// after apply" errors (e.g. NITRO defaults samplesize/threshold to 3 when the
// user leaves them null). The composite ID is set once in Create, not here.
func nstimer_autoscalepolicy_bindingSetAttrFromGet(ctx context.Context, data *NstimerAutoscalepolicyBindingResourceModel, getResponseData map[string]interface{}) *NstimerAutoscalepolicyBindingResourceModel {
	tflog.Debug(ctx, "In nstimer_autoscalepolicy_bindingSetAttrFromGet Function")

	// name and policyname are the binding keys; adopt them from GET only when
	// the model has no value yet (covers import).
	if data.Name.IsNull() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	if data.Policyname.IsNull() || data.Policyname.ValueString() == "" {
		if val, ok := getResponseData["policyname"]; ok && val != nil {
			data.Policyname = types.StringValue(val.(string))
		}
	}

	// priority is Required and echoed back by GET - adopt it.
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	// gotopriorityexpression, samplesize, threshold, vserver are preserved from
	// plan/state (RequiresReplace, not Computed).

	return data
}
