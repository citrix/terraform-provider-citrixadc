package authenticationwebauthpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationwebauthpolicyResourceModel describes the resource data model.
type AuthenticationwebauthpolicyResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"`
	Rule   types.String `tfsdk:"rule"`
}

func (r *AuthenticationwebauthpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationwebauthpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the WebAuth action to perform if the policy matches.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the WebAuth policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after LDAP policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the Web server.",
			},
		},
	}
}

func authenticationwebauthpolicyGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationwebauthpolicyResourceModel) authentication.Authenticationwebauthpolicy {
	tflog.Debug(ctx, "In authenticationwebauthpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationwebauthpolicy := authentication.Authenticationwebauthpolicy{}
	if !data.Action.IsNull() {
		authenticationwebauthpolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationwebauthpolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() {
		authenticationwebauthpolicy.Rule = data.Rule.ValueString()
	}

	return authenticationwebauthpolicy
}

func authenticationwebauthpolicySetAttrFromGet(ctx context.Context, data *AuthenticationwebauthpolicyResourceModel, getResponseData map[string]interface{}) *AuthenticationwebauthpolicyResourceModel {
	tflog.Debug(ctx, "In authenticationwebauthpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
