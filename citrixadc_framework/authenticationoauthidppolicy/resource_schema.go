package authenticationoauthidppolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationoauthidppolicyResourceModel describes the resource data model.
type AuthenticationoauthidppolicyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Action      types.String `tfsdk:"action"`
	Comment     types.String `tfsdk:"comment"`
	Logaction   types.String `tfsdk:"logaction"`
	Name        types.String `tfsdk:"name"`
	Newname     types.String `tfsdk:"newname"`
	Rule        types.String `tfsdk:"rule"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *AuthenticationoauthidppolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationoauthidppolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the profile to apply to requests or connections that match this policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of messagelog action to use when a request matches this policy.",
			},
			"name": schema.StringAttribute{
				Required: true,
				// SDK v2 had ForceNew on name -> RequiresReplace. Changing the primary
				// key itself is a recreate. An in-place name change goes through newname.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the OAuth Identity Provider (IdP) authentication policy. This is used for configuring Citrix ADC as OAuth Identity Provider. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				// newname is the rename trigger (NITRO ?action=rename). Changing it must
				// NOT force replacement - it drives an in-place rename via Update. It is a
				// pure user input, never echoed back by GET, so it is NOT Computed.
				Description: "New name for the OAuth IdentityProvider policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my oauthidppolicy policy\" or 'my oauthidppolicy policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression that the policy uses to determine whether to respond to the specified request.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only DROP/RESET actions can be used.",
			},
		},
	}
}

func authenticationoauthidppolicyGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationoauthidppolicyResourceModel) authentication.Authenticationoauthidppolicy {
	tflog.Debug(ctx, "In authenticationoauthidppolicyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationoauthidppolicy := authentication.Authenticationoauthidppolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		authenticationoauthidppolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		authenticationoauthidppolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() && !data.Logaction.IsUnknown() {
		authenticationoauthidppolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationoauthidppolicy.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of the
	// add/update payload, so it is deliberately excluded from the create/update body.
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		authenticationoauthidppolicy.Rule = data.Rule.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		authenticationoauthidppolicy.Undefaction = data.Undefaction.ValueString()
	}

	return authenticationoauthidppolicy
}

// authenticationoauthidppolicySetAttrFromGet updates the resource model from a GET
// response while PRESERVING configured/planned values that NITRO omits or that must
// not be clobbered (Pattern 7). Optional+Computed string attributes that NITRO omits
// when empty are only nulled when still unknown, so a configured value is preserved.
func authenticationoauthidppolicySetAttrFromGet(ctx context.Context, data *AuthenticationoauthidppolicyResourceModel, getResponseData map[string]interface{}) *AuthenticationoauthidppolicyResourceModel {
	tflog.Debug(ctx, "In authenticationoauthidppolicySetAttrFromGet Function")

	// action is Required and always echoed by GET.
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
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
	// (tracked by data.Id) diverges from the configured name and GET returns the live
	// name; overwriting name from GET would clobber the user's configured value and
	// trigger a spurious RequiresReplace diff. Only adopt the GET value when we do not
	// already have one (e.g. on import, where state carries only the ID).
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	// rule is Required and always echoed by GET.
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else if data.Undefaction.IsUnknown() {
		data.Undefaction = types.StringNull()
	}

	return data
}

// authenticationoauthidppolicySetAttrFromGetForDatasource faithfully copies every
// field from the GET response. The datasource has no prior plan/state to preserve, so
// it populates the model directly from the API response and sets the ID itself.
func authenticationoauthidppolicySetAttrFromGetForDatasource(ctx context.Context, data *AuthenticationoauthidppolicyResourceModel, getResponseData map[string]interface{}) *AuthenticationoauthidppolicyResourceModel {
	tflog.Debug(ctx, "In authenticationoauthidppolicySetAttrFromGetForDatasource Function")

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
