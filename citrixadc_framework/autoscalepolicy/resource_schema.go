package autoscalepolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/autoscale"

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
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. It does nothing when the config still carries a
// value, on create (no prior state), or when the prior value is already empty.
type unsetOnRemoveStringModifier struct{}

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-empty value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" {
		resp.PlanValue = types.StringUnknown()
	}
}

// AutoscalepolicyResourceModel describes the resource data model.
type AutoscalepolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Action    types.String `tfsdk:"action"`
	Comment   types.String `tfsdk:"comment"`
	Logaction types.String `tfsdk:"logaction"`
	Name      types.String `tfsdk:"name"`
	Newname   types.String `tfsdk:"newname"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *AutoscalepolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the autoscalepolicy resource.",
			},
			// Backward compat with SDK v2: action was Optional+Computed there, not
			// Required. The auto-generated code regressed it to Required.
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The autoscale profile associated with the policy.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Comments associated with this autoscale policy.",
			},
			"logaction": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "The log action associated with the autoscale policy",
			},
			// name is the primary key. SDK v2 marked it Required + ForceNew, so it
			// must be Required + RequiresReplace here.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the autoscale policy.",
			},
			// newname is the rename trigger (NITRO ?action=rename). It must be
			// Optional only: changing it drives an in-place rename via Update (NOT a
			// replace), and it is never echoed back by GET (so NOT Computed - a
			// Computed value the server never returns stays unknown after apply).
			"newname": schema.StringAttribute{
				Optional:    true,
				Description: "The new name of the autoscale policy.",
			},
			// Backward compat with SDK v2: rule was Optional+Computed there, not
			// Required. The auto-generated code regressed it to Required.
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The rule associated with the policy.",
			},
		},
	}
}

func autoscalepolicyGetThePayloadFromthePlan(ctx context.Context, data *AutoscalepolicyResourceModel) autoscale.Autoscalepolicy {
	tflog.Debug(ctx, "In autoscalepolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	autoscalepolicy := autoscale.Autoscalepolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		autoscalepolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		autoscalepolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		autoscalepolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		autoscalepolicy.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		autoscalepolicy.Rule = data.Rule.ValueString()
	}

	return autoscalepolicy
}

func autoscalepolicySetAttrFromGet(ctx context.Context, data *AutoscalepolicyResourceModel, getResponseData map[string]interface{}) *AutoscalepolicyResourceModel {
	tflog.Debug(ctx, "In autoscalepolicySetAttrFromGet Function")

	// Convert API response to model
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
	// name is the user-facing key. Once a rename has happened (via newname), the
	// live object name (tracked by data.Id) diverges from the configured name, and
	// GET returns the live (new) name. Overwriting name from GET would clobber the
	// user's configured value and trigger a spurious RequiresReplace diff. So only
	// adopt the GET value when we don't already have one (e.g. on import, where
	// state carries only the ID); otherwise preserve.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	return data
}

// autoscalepolicySetAttrFromGetForDatasource faithfully copies every field from
// the GET response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func autoscalepolicySetAttrFromGetForDatasource(ctx context.Context, data *AutoscalepolicyResourceModel, getResponseData map[string]interface{}) *AutoscalepolicyResourceModel {
	tflog.Debug(ctx, "In autoscalepolicySetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
