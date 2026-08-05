package sslpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslpolicyResourceModel describes the resource data model.
type SslpolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Name        types.String `tfsdk:"name"`
	Reqaction   types.String `tfsdk:"reqaction"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *SslpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslpolicy resource.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the built-in or user-defined action to perform on the request. Available built-in actions are NOOP, RESET, DROP, CLIENTAUTH, NOCLIENTAUTH, INTERCEPT AND BYPASS.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with this policy.",
			},
			"name": schema.StringAttribute{
				Required: true,
				// SDK v2 had ForceNew:true on name (cannot be changed after creation).
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the new SSL policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.  Cannot be changed after the policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"reqaction": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// SDK v2 did NOT set ForceNew on reqaction, so no RequiresReplace here
				// (auto-gen wrongly added it). reqaction is write-on-create only: NITRO
				// neither returns it in GET nor accepts it in update.
				Description: "The name of the action to be performed on the request. Refer to 'add ssl action' command to add a new action. Builtin actions like NOOP, RESET, DROP, CLIENTAUTH and NOCLIENTAUTH are also allowed.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression, against which traffic is evaluated.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the  character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the action to be performed when the result of rule evaluation is undefined. Possible values for control policies: CLIENTAUTH, NOCLIENTAUTH, NOOP, RESET, DROP. Possible values for data policies: NOOP, RESET, DROP and BYPASS",
			},
		},
	}
}

// sslpolicyGetThePayloadFromthePlan builds the full add payload (create).
func sslpolicyGetThePayloadFromthePlan(ctx context.Context, data *SslpolicyResourceModel) ssl.Sslpolicy {
	tflog.Debug(ctx, "In sslpolicyGetThePayloadFromthePlan Function")

	sslpolicy := ssl.Sslpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		sslpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		sslpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sslpolicy.Name = data.Name.ValueString()
	}
	if !data.Reqaction.IsNull() && !data.Reqaction.IsUnknown() {
		sslpolicy.Reqaction = data.Reqaction.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		sslpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		sslpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return sslpolicy
}

// sslpolicyGetTheUpdatablePayloadFromThePlan builds the update (PUT) payload.
// Per the NITRO doc, the sslpolicy update payload accepts name, rule, action,
// undefaction and comment ONLY - reqaction is not settable via update, so it is
// excluded here.
func sslpolicyGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *SslpolicyResourceModel) ssl.Sslpolicy {
	tflog.Debug(ctx, "In sslpolicyGetTheUpdatablePayloadFromThePlan Function")

	sslpolicy := ssl.Sslpolicy{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		sslpolicy.Name = data.Name.ValueString()
	}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		sslpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		sslpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		sslpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		sslpolicy.Undefaction = data.Undefaction.ValueString()
	}
	// reqaction intentionally excluded - not accepted by NITRO update.

	return sslpolicy
}

// sslpolicySetAttrFromGet updates the resource state from a GET response.
// Else-branches only null values that are still unknown so a configured value
// that NITRO omits from GET is never clobbered (omit-on-default guard).
func sslpolicySetAttrFromGet(ctx context.Context, data *SslpolicyResourceModel, getResponseData map[string]interface{}) *SslpolicyResourceModel {
	tflog.Debug(ctx, "In sslpolicySetAttrFromGet Function")

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
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// reqaction is never returned by NITRO GET; preserve the configured/prior
	// value and only resolve an unknown (Computed, unconfigured) to null.
	if data.Reqaction.IsUnknown() {
		data.Reqaction = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else if data.Undefaction.IsUnknown() {
		data.Undefaction = types.StringNull()
	}

	// Set ID for the resource - single unique attribute (name) as plain value.
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}

// sslpolicySetAttrFromGetForDatasource copies every value returned by the GET
// response into the model and sets the ID. Used by the datasource, which has no
// prior configured state to preserve.
func sslpolicySetAttrFromGetForDatasource(ctx context.Context, data *SslpolicyResourceModel, getResponseData map[string]interface{}) *SslpolicyResourceModel {
	tflog.Debug(ctx, "In sslpolicySetAttrFromGetForDatasource Function")

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
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// reqaction is not returned by NITRO GET.
	if val, ok := getResponseData["reqaction"]; ok && val != nil {
		data.Reqaction = types.StringValue(val.(string))
	} else {
		data.Reqaction = types.StringNull()
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

	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
