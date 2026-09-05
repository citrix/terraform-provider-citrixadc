package authenticationpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationpolicyResourceModel describes the resource data model.
type AuthenticationpolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Name        types.String `tfsdk:"name"`
	Newname     types.String `tfsdk:"newname"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *AuthenticationpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the authentication action to be performed if the policy matches.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Default "" makes removing comment from config produce a plan diff
				// (an Optional+Computed attr with no Default is sticky on removal, so
				// Update would never run and the unset would never fire). NITRO's
				// server default for comment is empty (GET omits it when unset).
				Default:     stringdefault.StaticString(""),
				Description: "Any comments to preserve information about this policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of messagelog action to use when a request matches this policy.",
			},
			"name": schema.StringAttribute{
				Required: true,
				// SDK v2 marked name ForceNew=true (cannot be changed after policy
				// is created), so a change to name forces recreation. Renaming an
				// existing policy in place is done via the separate newname attribute.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the advance AUTHENTICATION policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after AUTHENTICATION policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				// newname is the rename trigger (NITRO ?action=rename). Changing it
				// must NOT force replacement - it drives an in-place rename via Update.
				// Not Computed: it is a pure user input, never echoed back by GET.
				Description: "New name for the authentication policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the AUTHENTICATION server.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.",
			},
		},
	}
}

func authenticationpolicyGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationpolicyResourceModel) authentication.Authenticationpolicy {
	tflog.Debug(ctx, "In authenticationpolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationpolicy := authentication.Authenticationpolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		authenticationpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		authenticationpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		authenticationpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationpolicy.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add/update payload, so it is deliberately excluded here.
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		authenticationpolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		authenticationpolicy.Undefaction = data.Undefaction.ValueString()
	}

	return authenticationpolicy
}

func authenticationpolicySetAttrFromGet(ctx context.Context, data *AuthenticationpolicyResourceModel, getResponseData map[string]interface{}) *AuthenticationpolicyResourceModel {
	tflog.Debug(ctx, "In authenticationpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		// comment is Optional+Computed with Default "". NITRO omits comment from
		// GET when it is unset, so map the absent value to "" (not null) to match
		// the schema default and avoid an inconsistent-result error after an unset.
		data.Comment = types.StringValue("")
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
	// state carries only the ID); otherwise preserve the existing value.
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
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}

	return data
}

// authenticationpolicySetAttrFromGetForDatasource faithfully copies every field
// from the GET response. The datasource has no prior plan/state to preserve, so it
// must populate the model directly from the API response and set the ID itself.
func authenticationpolicySetAttrFromGetForDatasource(ctx context.Context, data *AuthenticationpolicyResourceModel, getResponseData map[string]interface{}) *AuthenticationpolicyResourceModel {
	tflog.Debug(ctx, "In authenticationpolicySetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		// comment is Optional+Computed with Default "". NITRO omits comment from
		// GET when it is unset, so map the absent value to "" (not null) to match
		// the schema default and avoid an inconsistent-result error after an unset.
		data.Comment = types.StringValue("")
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
