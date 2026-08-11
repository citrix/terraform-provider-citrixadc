package lbpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

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
// prior) and call Update, which issues the NITRO ?action=unset — mirroring the
// SDK v2 unset-on-remove contract. Without it an Optional+Computed attribute is
// "sticky": the prior value is carried forward and removal is a silent no-op.
// It intentionally does nothing when the config still carries a value, on create
// (no prior state), or when the prior value is already empty (avoids churn).
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

// LbpolicyResourceModel describes the resource data model.
type LbpolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Name        types.String `tfsdk:"name"`
	Newname     types.String `tfsdk:"newname"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *LbpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbpolicy resource.",
			},
			"action": schema.StringAttribute{
				// SDK v2 parity: Required (not Computed).
				Required:    true,
				Description: "Name of action to use if the request matches this LB policy.",
			},
			"comment": schema.StringAttribute{
				// SDK v2 parity: Optional+Computed.
				// Option B unset: removing comment from config marks it unknown via the
				// plan modifier (not a static Default), producing a plan diff that lets
				// Update fire the NITRO unset while still round-tripping cleanly on import
				// (NITRO omits comment from GET when unset, so a Default would break import).
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Any type of information about this LB policy.",
			},
			"logaction": schema.StringAttribute{
				// SDK v2 parity: Optional+Computed.
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				// SDK v2 parity: Required + ForceNew -> RequiresReplace. Changing the
				// name itself recreates the resource, exactly like SDK v2. In-place
				// renames are driven by the separate `newname` attribute (see below).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the LB policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the LB policy is added.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb policy\" or 'my lb policy').",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). It is a pure
				// user input, never echoed back by GET, so it is Optional-only:
				// NOT Computed (avoids known-after-apply churn) and NOT RequiresReplace
				// (a change must reach Update to drive an in-place rename, not recreate).
				Optional:    true,
				Description: "New name for the LB policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb policy\" or 'my lb policy').",
			},
			"rule": schema.StringAttribute{
				// SDK v2 parity: Required (not Computed).
				Required:    true,
				Description: "Expression against which traffic is evaluated.",
			},
			"undefaction": schema.StringAttribute{
				// SDK v2 parity: Optional+Computed.
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Available settings function as follows:\n* NOLBACTION - Does not consider LB actions in making LB decision.\n* RESET - Reset the request and notify the user, so that the user can resend the request.\n* DROP - Drop the request without sending a response to the user.",
			},
		},
	}
}

// lbpolicyGetThePayloadFromthePlan builds the add/create POST body. newname is a
// rename-only argument and is deliberately excluded from the create payload.
func lbpolicyGetThePayloadFromthePlan(ctx context.Context, data *LbpolicyResourceModel) lb.Lbpolicy {
	tflog.Debug(ctx, "In lbpolicyGetThePayloadFromthePlan Function")

	lbpolicy := lb.Lbpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		lbpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		lbpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		lbpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		lbpolicy.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename); excluded from add.
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		lbpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		lbpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return lbpolicy
}

// lbpolicyGetTheUpdatablePayloadFromThePlan builds the UpdateResource body,
// restricted to NITRO-updatable fields. newname is excluded (rename-only).
func lbpolicyGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *LbpolicyResourceModel) lb.Lbpolicy {
	tflog.Debug(ctx, "In lbpolicyGetTheUpdatablePayloadFromThePlan Function")

	lbpolicy := lb.Lbpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		lbpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		lbpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		lbpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		lbpolicy.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename); excluded from update.
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		lbpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		lbpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return lbpolicy
}

// lbpolicySetAttrFromGet populates the resource model from a GET response.
// It preserves configured/plan values that NITRO omits from GET (omit-on-default
// trap guard: only resolve to null when the field is still Unknown), and never
// clobbers the user-facing name after a rename.
func lbpolicySetAttrFromGet(ctx context.Context, data *LbpolicyResourceModel, getResponseData map[string]interface{}) *LbpolicyResourceModel {
	tflog.Debug(ctx, "In lbpolicySetAttrFromGet Function")

	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else if data.Action.IsUnknown() {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else if data.Logaction.IsUnknown() {
		data.Logaction = types.StringNull()
	}
	// name is the user-facing key. After a rename (via newname) the live object name
	// (tracked by data.Id) diverges from the configured name, and GET returns the
	// live name. Overwriting name from GET would clobber the user's configured value
	// and trigger a spurious RequiresReplace diff. Only adopt the GET value when we
	// don't already have one (e.g. on import, where state carries only the ID).
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else if data.Rule.IsUnknown() {
		data.Rule = types.StringNull()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else if data.Undefaction.IsUnknown() {
		data.Undefaction = types.StringNull()
	}

	return data
}

// lbpolicySetAttrFromGetForDatasource faithfully copies every field from the GET
// response. The datasource has no prior plan/state to preserve, so it populates
// the model directly from the API response and sets the ID itself.
func lbpolicySetAttrFromGetForDatasource(ctx context.Context, data *LbpolicyResourceModel, getResponseData map[string]interface{}) *LbpolicyResourceModel {
	tflog.Debug(ctx, "In lbpolicySetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
