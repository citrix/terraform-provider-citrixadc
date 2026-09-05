package vpnurlpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

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
// post-unset value NITRO reports is empty/absent (it cannot be sent in add/update).
// defaultValue is the value the attribute reverts to after a NITRO unset. The
// modifier skips the unknown-forcing when the prior state already equals this
// value, otherwise removal from config would perpetually re-plan.
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

// VpnurlpolicyResourceModel describes the resource data model.
type VpnurlpolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Action    types.String `tfsdk:"action"`
	Comment   types.String `tfsdk:"comment"`
	Logaction types.String `tfsdk:"logaction"`
	Name      types.String `tfsdk:"name"`
	Newname   types.String `tfsdk:"newname"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *VpnurlpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnurlpolicy resource.",
			},
			"action": schema.StringAttribute{
				// SDK v2 parity: Required, updateable (no ForceNew).
				Required:    true,
				Description: "Action to be applied by the new urlPolicy if the rule criteria are met.",
			},
			"comment": schema.StringAttribute{
				// SDK v2 parity: Optional+Computed, updateable, no default.
				// Removal from config must trigger an unset (revert to the NITRO
				// default: comment is dropped/absent on GET). The post-unset value is
				// empty, so a schema Default cannot be used; this modifier forces the
				// plan to unknown on removal so Update runs the ?action=unset.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{defaultValue: ""},
				},
				Description: "Any comments to preserve information about this policy.",
			},
			"logaction": schema.StringAttribute{
				// SDK v2 parity: Optional+Computed, updateable, no default.
				// Removal from config must trigger an unset (revert to the NITRO
				// default: logaction is dropped/absent on GET). The post-unset value is
				// empty, so a schema Default cannot be used; this modifier forces the
				// plan to unknown on removal so Update runs the ?action=unset.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{defaultValue: ""},
				},
				Description: "Name of messagelog action to use when a request matches this policy.",
			},
			"name": schema.StringAttribute{
				// SDK v2 parity: Required + ForceNew -> Required + RequiresReplace.
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the new urlPolicy.",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it
				// must NOT force replacement - it drives an in-place rename via Update.
				// Not Computed: it is a pure user input, never echoed back by GET.
				Optional:    true,
				Description: "New name for the vpn urlPolicy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vpnurl policy\" or 'my vpnurl policy').",
			},
			"rule": schema.StringAttribute{
				// SDK v2 parity: Required, updateable (no ForceNew).
				Required:    true,
				Description: "Expression, or name of a named expression, specifying the traffic that matches the policy.\n\nThe following requirements apply only to the NetScaler CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

func vpnurlpolicyGetThePayloadFromthePlan(ctx context.Context, data *VpnurlpolicyResourceModel) vpn.Vpnurlpolicy {
	tflog.Debug(ctx, "In vpnurlpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnurlpolicy := vpn.Vpnurlpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		vpnurlpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		vpnurlpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		vpnurlpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnurlpolicy.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		vpnurlpolicy.Rule = data.Rule.ValueString()
	}

	return vpnurlpolicy
}

func vpnurlpolicyGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *VpnurlpolicyResourceModel) vpn.Vpnurlpolicy {
	tflog.Debug(ctx, "In vpnurlpolicyGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model, restricted to NITRO-updatable fields.
	vpnurlpolicy := vpn.Vpnurlpolicy{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnurlpolicy.Name = data.Name.ValueString()
	}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		vpnurlpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		vpnurlpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		vpnurlpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		vpnurlpolicy.Rule = data.Rule.ValueString()
	}

	return vpnurlpolicy
}

func vpnurlpolicySetAttrFromGet(ctx context.Context, data *VpnurlpolicyResourceModel, getResponseData map[string]interface{}) *VpnurlpolicyResourceModel {
	tflog.Debug(ctx, "In vpnurlpolicySetAttrFromGet Function")

	// Convert API response to model.
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

	// NOTE: data.Id is intentionally NOT set here. It is managed by the CRUD
	// functions (Create sets it to name, Update tracks renames, Read preserves it
	// from prior state / nulls it on not-found).

	return data
}

// vpnurlpolicySetAttrFromGetForDatasource faithfully copies every field from the
// GET response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func vpnurlpolicySetAttrFromGetForDatasource(ctx context.Context, data *VpnurlpolicyResourceModel, getResponseData map[string]interface{}) *VpnurlpolicyResourceModel {
	tflog.Debug(ctx, "In vpnurlpolicySetAttrFromGetForDatasource Function")

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
