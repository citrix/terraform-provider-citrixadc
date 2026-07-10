package rewritepolicylabel_rewritepolicy_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/rewrite"

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

// RewritepolicylabelRewritepolicyBindingResourceModel describes the resource data model.
type RewritepolicylabelRewritepolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Invokelabelname        types.String `tfsdk:"invokelabelname"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
}

func (r *RewritepolicylabelRewritepolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rewritepolicylabel_rewritepolicy_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Suspend evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.",
			},
			"invokelabelname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "* If labelType is policylabel, name of the policy label to invoke. \n* If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request or response.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the rewrite policy label to which to bind the policy.",
			},
			"labeltype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of invocation. Available settings function as follows:\n* reqvserver - Forward the request to the specified request virtual server.\n* resvserver - Forward the response to the specified response virtual server.\n* policylabel - Invoke the specified policy label.",
			},
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the rewrite policy to bind to the policy label.",
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

func rewritepolicylabel_rewritepolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *RewritepolicylabelRewritepolicyBindingResourceModel) rewrite.Rewritepolicylabelrewritepolicybinding {
	tflog.Debug(ctx, "In rewritepolicylabel_rewritepolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	rewritepolicylabel_rewritepolicy_binding := rewrite.Rewritepolicylabelrewritepolicybinding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		rewritepolicylabel_rewritepolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Invoke.IsNull() && !data.Invoke.IsUnknown() {
		rewritepolicylabel_rewritepolicy_binding.Invoke = data.Invoke.ValueBool()
	}
	if !data.Invokelabelname.IsNull() && !data.Invokelabelname.IsUnknown() {
		rewritepolicylabel_rewritepolicy_binding.Invokelabelname = data.Invokelabelname.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		rewritepolicylabel_rewritepolicy_binding.Labelname = data.Labelname.ValueString()
	}
	if !data.Labeltype.IsNull() && !data.Labeltype.IsUnknown() {
		rewritepolicylabel_rewritepolicy_binding.Labeltype = data.Labeltype.ValueString()
	}
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		rewritepolicylabel_rewritepolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		rewritepolicylabel_rewritepolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return rewritepolicylabel_rewritepolicy_binding
}

// rewritepolicylabel_rewritepolicy_bindingSetAttrFromGet is used by the resource Read flow.
// It does NOT recompute data.Id (the ID is set once in Create — see FeatureDeveloper
// Pattern 6). The NITRO binding GET reliably echoes back every field that was supplied
// on the add call, so Optional+Computed fields that the user did not set are simply
// absent from the response and are resolved to null (matching the SDK v2 contract where
// these attributes are Optional+Computed). This keeps every Computed attribute known
// after apply (avoids the "still indicated an unknown value" error).
func rewritepolicylabel_rewritepolicy_bindingSetAttrFromGet(ctx context.Context, data *RewritepolicylabelRewritepolicyBindingResourceModel, getResponseData map[string]interface{}) *RewritepolicylabelRewritepolicyBindingResourceModel {
	tflog.Debug(ctx, "In rewritepolicylabel_rewritepolicy_bindingSetAttrFromGet Function")

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
		data.Invokelabelname = types.StringValue(val.(string))
	} else {
		data.Invokelabelname = types.StringNull()
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

	return data
}

// rewritepolicylabel_rewritepolicy_bindingSetAttrFromGetForDatasource is used by the
// datasource Read flow. Unlike the resource setter, it faithfully copies every field
// from the GET response (a datasource has no prior plan/state to preserve) and nulls
// fields absent from the response. The datasource Read sets data.Id itself.
func rewritepolicylabel_rewritepolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *RewritepolicylabelRewritepolicyBindingResourceModel, getResponseData map[string]interface{}) *RewritepolicylabelRewritepolicyBindingResourceModel {
	tflog.Debug(ctx, "In rewritepolicylabel_rewritepolicy_bindingSetAttrFromGetForDatasource Function")

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
		data.Invokelabelname = types.StringValue(val.(string))
	} else {
		data.Invokelabelname = types.StringNull()
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

	return data
}
