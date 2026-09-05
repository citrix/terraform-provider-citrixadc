package botpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. Because these attributes revert to no value
// (absent from GET) after unset, marking the plan unknown also avoids a
// "provider produced inconsistent result" error, which a static Default would
// trigger.
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

// BotpolicyResourceModel describes the resource data model.
type BotpolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Name        types.String `tfsdk:"name"`
	Newname     types.String `tfsdk:"newname"`
	Profilename types.String `tfsdk:"profilename"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *BotpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botpolicy resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Any type of information about this bot policy.",
			},
			"logaction": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// NITRO always echoes logaction with the non-empty default "None",
				// and unset reverts it to "None". A static Default (not the
				// unset-on-remove modifier, which is for attributes that revert to
				// absent) keeps the plan stable when logaction is unconfigured while
				// still producing a diff -> unset when a set value is removed.
				Default:     stringdefault.StaticString("None"),
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required: true,
				// SDK v2 marked name as ForceNew.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the bot policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the bot policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my bot policy\" or 'my bot policy').",
			},
			"newname": schema.StringAttribute{
				// newname is a pure user input used only by the NITRO rename action.
				// It is never echoed back by GET, so it must NOT be Computed (would
				// cause known-after-apply churn) and must NOT be RequiresReplace
				// (would force recreation instead of an in-place rename).
				Optional:    true,
				Description: "New name for the bot policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. \n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my bot policy\" or 'my bot policy').",
			},
			"profilename": schema.StringAttribute{
				Required: true,
				// SDK v2 marked profilename as ForceNew.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the bot profile to apply if the request matches this bot policy.",
			},
			"rule": schema.StringAttribute{
				Required: true,
				// SDK v2 marked rule as ForceNew.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression that the policy uses to determine whether to apply bot profile on the specified request.",
			},
			"undefaction": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// NITRO always echoes undefaction with the non-empty default "None",
				// and unset reverts it to "None". Use a static Default so an
				// unconfigured undefaction stays stable across plans while a removed
				// set value still diffs -> unset.
				Default:     stringdefault.StaticString("None"),
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition.",
			},
		},
	}
}

func botpolicyGetThePayloadFromthePlan(ctx context.Context, data *BotpolicyResourceModel) bot.Botpolicy {
	tflog.Debug(ctx, "In botpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model. Both the NITRO add (POST) and
	// update (PUT) endpoints accept the same field set: name, rule, profilename,
	// undefaction, comment, logaction. newname is intentionally excluded — it is
	// only ever sent via the rename action.
	botpolicy := bot.Botpolicy{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		botpolicy.Comment = data.Comment.ValueString()
	}
	// "None" is NITRO's read-back sentinel for "no action"; it is not a valid
	// input value (POST/PUT reject it), so it is never sent in the payload. It is
	// also the value the appliance reverts to on unset.
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() && data.Logaction.ValueString() != "None" {
		botpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		botpolicy.Name = data.Name.ValueString()
	}
	// Skip newname: it is rename-only and must not be part of add/update payloads.
	if !data.Profilename.IsNull() && !data.Profilename.IsUnknown() {
		botpolicy.Profilename = data.Profilename.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		botpolicy.Rule = data.Rule.ValueString()
	}
	// "None" is the read-back sentinel / unset revert value for undefaction and is
	// rejected as an input, so it is never sent in the payload.
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() && data.Undefaction.ValueString() != "None" {
		botpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return botpolicy
}

func botpolicySetAttrFromGet(ctx context.Context, data *BotpolicyResourceModel, getResponseData map[string]interface{}) *BotpolicyResourceModel {
	tflog.Debug(ctx, "In botpolicySetAttrFromGet Function")

	// Convert API response to model
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

	// name is the primary key and also the live object name. After a rename the
	// live name (GET response) differs from the configured name attribute, so we
	// only adopt the GET value when the configured value is absent (import case);
	// otherwise we preserve the configured value to avoid a spurious
	// RequiresReplace diff.
	var liveName string
	if val, ok := getResponseData["name"]; ok && val != nil {
		liveName = val.(string)
		if data.Name.IsNull() || data.Name.ValueString() == "" {
			data.Name = types.StringValue(liveName)
		}
	}

	// newname is rename-only and is never returned by GET. Preserve the existing
	// plan/state value (do not null it) to keep it consistent for Optional inputs.

	if val, ok := getResponseData["profilename"]; ok && val != nil {
		data.Profilename = types.StringValue(val.(string))
	} else {
		data.Profilename = types.StringNull()
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

	// Set ID for the resource. The ID tracks the live object name so that Read
	// and Delete continue to work after a rename.
	if liveName != "" {
		data.Id = types.StringValue(liveName)
	} else {
		data.Id = types.StringValue(data.Name.ValueString())
	}

	return data
}
