package crpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cr"

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

// CrpolicyResourceModel describes the resource data model.
type CrpolicyResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Action     types.String `tfsdk:"action"`
	Logaction  types.String `tfsdk:"logaction"`
	Newname    types.String `tfsdk:"newname"`
	Policyname types.String `tfsdk:"policyname"`
	Rule       types.String `tfsdk:"rule"`
}

func (r *CrpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the crpolicy resource.",
			},
			"action": schema.StringAttribute{
				// SDK v2 parity: Optional + Computed.
				Optional:    true,
				Computed:    true,
				Description: "Name of the built-in cache redirection action: CACHE/ORIGIN.",
			},
			"logaction": schema.StringAttribute{
				// SDK v2 parity: Optional + Computed.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "The log action associated with the cache redirection policy",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it
				// must NOT force replacement - it drives an in-place rename via Update.
				// Not Computed: it is a pure user input, never echoed back by GET.
				Optional:    true,
				Description: "The new name of the content switching policy.",
			},
			"policyname": schema.StringAttribute{
				// SDK v2 parity: Required + ForceNew -> RequiresReplace. Cannot be
				// changed after the policy is created (renaming is done via newname).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the cache redirection policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the policy is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"rule": schema.StringAttribute{
				// SDK v2 parity: Required (Computed: false).
				Required:    true,
				Description: "Expression, or name of a named expression, against which traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n*  If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n*  If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n*  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

// crpolicyGetThePayloadFromthePlan builds the add (create) payload. newname is a
// rename-only argument and is deliberately excluded from the create POST body.
func crpolicyGetThePayloadFromthePlan(ctx context.Context, data *CrpolicyResourceModel) cr.Crpolicy {
	tflog.Debug(ctx, "In crpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	crpolicy := cr.Crpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		crpolicy.Action = data.Action.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		crpolicy.Logaction = data.Logaction.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		crpolicy.Policyname = data.Policyname.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		crpolicy.Rule = data.Rule.ValueString()
	}

	return crpolicy
}

// crpolicyGetTheUpdatablePayloadFromThePlan builds the update (PUT) payload,
// restricted to the NITRO-updatable, non-key attributes. The caller sets
// Policyname to the current live name before issuing the request.
func crpolicyGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *CrpolicyResourceModel) cr.Crpolicy {
	tflog.Debug(ctx, "In crpolicyGetTheUpdatablePayloadFromThePlan Function")

	crpolicy := cr.Crpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		crpolicy.Action = data.Action.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		crpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		crpolicy.Rule = data.Rule.ValueString()
	}

	return crpolicy
}

// crpolicySetAttrFromGet populates the model from a GET response for the RESOURCE.
// It preserves the user-facing key (policyname) and the rename-only newname so a
// rename does not clobber the configured values (see lbpolicylabel reference).
func crpolicySetAttrFromGet(ctx context.Context, data *CrpolicyResourceModel, getResponseData map[string]interface{}) *CrpolicyResourceModel {
	tflog.Debug(ctx, "In crpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else {
		data.Logaction = types.StringNull()
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	// policyname is the user-facing key. After a rename the live object name
	// (tracked by data.Id) diverges from the configured policyname, and GET returns
	// the live (new) name. Overwriting policyname from GET would clobber the user's
	// configured value and trigger a spurious RequiresReplace diff. Only adopt the
	// GET value when we don't already have one (e.g. on import, where state carries
	// only the ID); otherwise preserve.
	if data.Policyname.IsNull() || data.Policyname.IsUnknown() || data.Policyname.ValueString() == "" {
		if val, ok := getResponseData["policyname"]; ok && val != nil {
			data.Policyname = types.StringValue(val.(string))
		}
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	return data
}

// crpolicySetAttrFromGetForDatasource faithfully copies every field from the GET
// response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func crpolicySetAttrFromGetForDatasource(ctx context.Context, data *CrpolicyResourceModel, getResponseData map[string]interface{}) *CrpolicyResourceModel {
	tflog.Debug(ctx, "In crpolicySetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else {
		data.Logaction = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(data.Policyname.ValueString())

	return data
}
