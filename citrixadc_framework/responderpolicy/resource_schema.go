package responderpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward and
// removal is a silent no-op. Unlike a schema Default it injects no value into the
// add/update payload (unknown values are skipped), which matters here because the
// NITRO defaults are pseudo-values rejected by add/update.
// defaultValue is the value the attribute reverts to after a NITRO unset (the
// appliance echoes it back on GET). The modifier skips the unknown-forcing when
// the prior state already equals this value, otherwise removal from config would
// perpetually re-plan (the post-unset value is non-empty).
type unsetOnRemoveStringModifier struct{ defaultValue string }

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-default value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	sv := req.StateValue.ValueString()
	if req.ConfigValue.IsNull() && sv != "" && sv != m.defaultValue {
		resp.PlanValue = types.StringUnknown()
	}
}

// ResponderpolicyGlobalbindingModel is one element of the globalbinding set.
type ResponderpolicyGlobalbindingModel struct {
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`
}

var responderpolicyGlobalbindingAttrTypes = map[string]attr.Type{
	"gotopriorityexpression": types.StringType,
	"invoke":                 types.BoolType,
	"labelname":              types.StringType,
	"labeltype":              types.StringType,
	"policyname":             types.StringType,
	"priority":               types.Int64Type,
	"type":                   types.StringType,
}

// ResponderpolicyLbvserverbindingModel is one element of the lbvserverbinding set.
type ResponderpolicyLbvserverbindingModel struct {
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Name                   types.String `tfsdk:"name"`
	Priority               types.Int64  `tfsdk:"priority"`
}

var responderpolicyLbvserverbindingAttrTypes = map[string]attr.Type{
	"bindpoint":              types.StringType,
	"gotopriorityexpression": types.StringType,
	"invoke":                 types.BoolType,
	"labelname":              types.StringType,
	"labeltype":              types.StringType,
	"name":                   types.StringType,
	"priority":               types.Int64Type,
}

// ResponderpolicyCsvserverbindingModel is one element of the csvserverbinding set.
type ResponderpolicyCsvserverbindingModel struct {
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Name                   types.String `tfsdk:"name"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Targetlbvserver        types.String `tfsdk:"targetlbvserver"`
}

var responderpolicyCsvserverbindingAttrTypes = map[string]attr.Type{
	"bindpoint":              types.StringType,
	"gotopriorityexpression": types.StringType,
	"invoke":                 types.BoolType,
	"labelname":              types.StringType,
	"labeltype":              types.StringType,
	"name":                   types.StringType,
	"policyname":             types.StringType,
	"priority":               types.Int64Type,
	"targetlbvserver":        types.StringType,
}

// ResponderpolicyResourceModel describes the resource data model.
type ResponderpolicyResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Action           types.String `tfsdk:"action"`
	Appflowaction    types.String `tfsdk:"appflowaction"`
	Comment          types.String `tfsdk:"comment"`
	Logaction        types.String `tfsdk:"logaction"`
	Name             types.String `tfsdk:"name"`
	Rule             types.String `tfsdk:"rule"`
	Undefaction      types.String `tfsdk:"undefaction"`
	Globalbinding    types.Set    `tfsdk:"globalbinding"`
	Lbvserverbinding types.Set    `tfsdk:"lbvserverbinding"`
	Csvserverbinding types.Set    `tfsdk:"csvserverbinding"`
}

func (r *ResponderpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the responderpolicy resource.",
			},
			// SDK v2 backward-compat: action/appflowaction/comment/logaction/name/rule/
			// undefaction were all Optional+Computed with NO ForceNew, so none carry
			// RequiresReplace. (The auto-gen wrongly made action/name/rule Required and
			// added a rename-only newname attribute that SDK v2 never exposed.)
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the responder action to perform if the request matches this responder policy.",
			},
			"appflowaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "AppFlow action to invoke for requests that match this policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any type of information about this responder policy.",
			},
			"logaction": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Removal from config must trigger an unset (revert to NITRO default
				// "None"). The NITRO default is a pseudo-value that cannot be sent in
				// an add/update payload, so a schema Default cannot be used; instead
				// this modifier forces the plan to unknown on removal so Update runs.
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{defaultValue: "None"},
				},
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Name for the responder policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that the policy uses to determine whether to respond to the specified request.",
			},
			"undefaction": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Removal from config must trigger an unset (revert to NITRO default
				// "Use Global"). That pseudo-value cannot be sent in an add/update
				// payload, so a schema Default cannot be used; this modifier forces
				// the plan to unknown on removal so Update runs.
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{defaultValue: "Use Global"},
				},
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF).",
			},
		},
		Blocks: map[string]schema.Block{
			// Convenience block: responderglobal_responderpolicy_binding.
			"globalbinding": schema.SetNestedBlock{
				Description: "Bind this responder policy to a global bind point.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"gotopriorityexpression": schema.StringAttribute{Optional: true, Computed: true, Description: "Expression or priority to determine the next policy to evaluate."},
						"invoke":                 schema.BoolAttribute{Optional: true, Computed: true, Description: "Invoke a policy label if the current policy evaluates to TRUE."},
						"labelname":              schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the policy label to invoke."},
						"labeltype":              schema.StringAttribute{Optional: true, Computed: true, Description: "Type of policy label to invoke."},
						"policyname":             schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the responder policy."},
						"priority":               schema.Int64Attribute{Optional: true, Computed: true, Description: "Priority of the policy binding."},
						"type":                   schema.StringAttribute{Optional: true, Computed: true, Description: "Global bind point to which the policy is bound."},
					},
				},
			},
			// Convenience block: lbvserver_responderpolicy_binding.
			"lbvserverbinding": schema.SetNestedBlock{
				Description: "Bind this responder policy to load balancing virtual servers.",
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
			// Convenience block: csvserver_responderpolicy_binding.
			"csvserverbinding": schema.SetNestedBlock{
				Description: "Bind this responder policy to content switching virtual servers.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"bindpoint":              schema.StringAttribute{Optional: true, Computed: true, Description: "Bind point to which to bind the policy."},
						"gotopriorityexpression": schema.StringAttribute{Optional: true, Computed: true, Description: "Expression or priority to determine the next policy to evaluate."},
						"invoke":                 schema.BoolAttribute{Optional: true, Computed: true, Description: "Invoke a policy label if the current policy evaluates to TRUE."},
						"labelname":              schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the policy label to invoke."},
						"labeltype":              schema.StringAttribute{Optional: true, Computed: true, Description: "Type of policy label to invoke."},
						"name":                   schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the content switching virtual server to bind the policy to."},
						"policyname":             schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the responder policy."},
						"priority":               schema.Int64Attribute{Optional: true, Computed: true, Description: "Priority of the policy binding."},
						"targetlbvserver":        schema.StringAttribute{Optional: true, Computed: true, Description: "Name of the LB vserver to bind the policy to."},
					},
				},
			},
		},
	}
}

func responderpolicyGetThePayloadFromthePlan(ctx context.Context, data *ResponderpolicyResourceModel) responder.Responderpolicy {
	tflog.Debug(ctx, "In responderpolicyGetThePayloadFromthePlan Function")

	responderpolicy := responder.Responderpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		responderpolicy.Action = data.Action.ValueString()
	}
	if !data.Appflowaction.IsNull() && !data.Appflowaction.IsUnknown() {
		responderpolicy.Appflowaction = data.Appflowaction.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		responderpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		responderpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		responderpolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		responderpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		responderpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return responderpolicy
}

func responderpolicySetAttrFromGet(ctx context.Context, data *ResponderpolicyResourceModel, getResponseData map[string]interface{}) *ResponderpolicyResourceModel {
	tflog.Debug(ctx, "In responderpolicySetAttrFromGet Function")

	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["appflowaction"]; ok && val != nil {
		data.Appflowaction = types.StringValue(val.(string))
	} else {
		data.Appflowaction = types.StringNull()
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

	// Case 2: Single unique attribute - use plain value (the name) as ID.
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
