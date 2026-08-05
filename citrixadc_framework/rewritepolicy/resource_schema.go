package rewritepolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/rewrite"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RewritepolicyGlobalbindingModel is one element of the globalbinding set.
type RewritepolicyGlobalbindingModel struct {
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
}

var rewritepolicyGlobalbindingAttrTypes = map[string]attr.Type{
	"globalbindtype":         types.StringType,
	"gotopriorityexpression": types.StringType,
	"invoke":                 types.BoolType,
	"labelname":              types.StringType,
	"labeltype":              types.StringType,
	"policyname":             types.StringType,
	"priority":               types.Int64Type,
	"type":                   types.StringType,
}

// RewritepolicyLbvserverbindingModel is one element of the lbvserverbinding set.
type RewritepolicyLbvserverbindingModel struct {
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Name                   types.String `tfsdk:"name"`
	Priority               types.Int64  `tfsdk:"priority"`
}

var rewritepolicyLbvserverbindingAttrTypes = map[string]attr.Type{
	"bindpoint":              types.StringType,
	"gotopriorityexpression": types.StringType,
	"invoke":                 types.BoolType,
	"labelname":              types.StringType,
	"labeltype":              types.StringType,
	"name":                   types.StringType,
	"priority":               types.Int64Type,
}

// RewritepolicyCsvserverbindingModel is one element of the csvserverbinding set.
type RewritepolicyCsvserverbindingModel struct {
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Name                   types.String `tfsdk:"name"`
	Priority               types.Int64  `tfsdk:"priority"`
	Targetlbvserver        types.String `tfsdk:"targetlbvserver"`
}

var rewritepolicyCsvserverbindingAttrTypes = map[string]attr.Type{
	"bindpoint":              types.StringType,
	"gotopriorityexpression": types.StringType,
	"invoke":                 types.BoolType,
	"labelname":              types.StringType,
	"labeltype":              types.StringType,
	"name":                   types.StringType,
	"priority":               types.Int64Type,
	"targetlbvserver":        types.StringType,
}

// RewritepolicyResourceModel describes the resource data model.
type RewritepolicyResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Action           types.String `tfsdk:"action"`
	Comment          types.String `tfsdk:"comment"`
	Logaction        types.String `tfsdk:"logaction"`
	Name             types.String `tfsdk:"name"`
	Rule             types.String `tfsdk:"rule"`
	Undefaction      types.String `tfsdk:"undefaction"`
	Globalbinding    types.Set    `tfsdk:"globalbinding"`
	Lbvserverbinding types.Set    `tfsdk:"lbvserverbinding"`
	Csvserverbinding types.Set    `tfsdk:"csvserverbinding"`
}

func (r *RewritepolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rewritepolicy resource.",
			},
			// SDK v2 backward-compat: action/comment/logaction/name/rule/undefaction
			// were all Optional+Computed with NO ForceNew, so none carry RequiresReplace.
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the rewrite action to perform if the request or response matches this rewrite policy.\nThere are also some built-in actions which can be used. These are:\n* NOREWRITE - Send the request from the client to the server or response from the server to the client without making any changes in the message.\n* RESET - Resets the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.\n* DROP - Drop the request without sending a response to the user.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this rewrite policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of messagelog action to use when a request matches this policy.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Name for the rewrite policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my rewrite policy\" or 'my rewrite policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression against which traffic is evaluated.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character. \n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.",
			},
		},
		Blocks: map[string]schema.Block{
			// Convenience block: bind this rewrite policy globally (rewriteglobal_rewritepolicy_binding).
			"globalbinding": schema.SetNestedBlock{
				Description: "Bind this rewrite policy to a global bind point.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"globalbindtype":         schema.StringAttribute{Optional: true, Computed: true, Description: "Global bind point to which the policy is bound."},
						"gotopriorityexpression": schema.StringAttribute{Optional: true, Computed: true, Description: "Expression or priority to determine the next policy to evaluate."},
						"invoke":                 schema.BoolAttribute{Optional: true, Computed: true, Description: "Invoke a policy label if the current policy evaluates to TRUE."},
						"labelname":              schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the policy label to invoke."},
						"labeltype":              schema.StringAttribute{Optional: true, Computed: true, Description: "Type of policy label to invoke."},
						"policyname":             schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the rewrite policy."},
						"priority":               schema.Int64Attribute{Optional: true, Computed: true, Description: "Priority of the policy binding."},
						"type":                   schema.StringAttribute{Optional: true, Computed: true, Description: "Global bind point to which the policy is bound."},
					},
				},
			},
			// Convenience block: bind this rewrite policy to lb vservers (lbvserver_rewritepolicy_binding).
			"lbvserverbinding": schema.SetNestedBlock{
				Description: "Bind this rewrite policy to load balancing virtual servers.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"bindpoint":              schema.StringAttribute{Optional: true, Computed: true, Description: "Bind point to which to bind the policy."},
						"gotopriorityexpression": schema.StringAttribute{Optional: true, Computed: true, Description: "Expression or priority to determine the next policy to evaluate."},
						"invoke":                 schema.BoolAttribute{Optional: true, Computed: true, Description: "Invoke a policy label if the current policy evaluates to TRUE."},
						"labelname":              schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the policy label to invoke."},
						"labeltype":              schema.StringAttribute{Optional: true, Computed: true, Description: "Type of policy label to invoke."},
						"name":                   schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the load balancing virtual server to bind the policy to."},
						"priority":               schema.Int64Attribute{Optional: true, Computed: true, Description: "Priority of the policy binding."},
					},
				},
			},
			// Convenience block: bind this rewrite policy to cs vservers (csvserver_rewritepolicy_binding).
			"csvserverbinding": schema.SetNestedBlock{
				Description: "Bind this rewrite policy to content switching virtual servers.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"bindpoint":              schema.StringAttribute{Optional: true, Computed: true, Description: "Bind point to which to bind the policy."},
						"gotopriorityexpression": schema.StringAttribute{Optional: true, Computed: true, Description: "Expression or priority to determine the next policy to evaluate."},
						"invoke":                 schema.BoolAttribute{Optional: true, Computed: true, Description: "Invoke a policy label if the current policy evaluates to TRUE."},
						"labelname":              schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the policy label to invoke."},
						"labeltype":              schema.StringAttribute{Optional: true, Computed: true, Description: "Type of policy label to invoke."},
						"name":                   schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the content switching virtual server to bind the policy to."},
						"priority":               schema.Int64Attribute{Optional: true, Computed: true, Description: "Priority of the policy binding."},
						"targetlbvserver":        schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the LB vserver to bind the policy to."},
					},
				},
			},
		},
	}
}

func rewritepolicyGetThePayloadFromthePlan(ctx context.Context, data *RewritepolicyResourceModel) rewrite.Rewritepolicy {
	tflog.Debug(ctx, "In rewritepolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	rewritepolicy := rewrite.Rewritepolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		rewritepolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		rewritepolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		rewritepolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		rewritepolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		rewritepolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		rewritepolicy.Undefaction = data.Undefaction.ValueString()
	}

	return rewritepolicy
}

func rewritepolicySetAttrFromGet(ctx context.Context, data *RewritepolicyResourceModel, getResponseData map[string]interface{}) *RewritepolicyResourceModel {
	tflog.Debug(ctx, "In rewritepolicySetAttrFromGet Function")

	// Convert API response to model. Only the scalar (base rewritepolicy)
	// attributes are handled here; the convenience-block sets
	// (globalbinding/lbvserverbinding/csvserverbinding) are reconciled
	// separately in the binding helpers.
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else {
		data.Logaction = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
