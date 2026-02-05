package authenticationsamlpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationsamlpolicyResourceModel describes the resource data model.
type AuthenticationsamlpolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Reqaction types.String `tfsdk:"reqaction"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *AuthenticationsamlpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationsamlpolicy resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the SAML policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after SAML policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"reqaction": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SAML authentication action to be performed if the policy matches.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the SAML server.",
			},
		},
	}
}

func authenticationsamlpolicyGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationsamlpolicyResourceModel) authentication.Authenticationsamlpolicy {
	tflog.Debug(ctx, "In authenticationsamlpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationsamlpolicy := authentication.Authenticationsamlpolicy{}
	if !data.Name.IsNull() {
		authenticationsamlpolicy.Name = data.Name.ValueString()
	}
	if !data.Reqaction.IsNull() {
		authenticationsamlpolicy.Reqaction = data.Reqaction.ValueString()
	}
	if !data.Rule.IsNull() {
		authenticationsamlpolicy.Rule = data.Rule.ValueString()
	}

	return authenticationsamlpolicy
}

func authenticationsamlpolicySetAttrFromGet(ctx context.Context, data *AuthenticationsamlpolicyResourceModel, getResponseData map[string]interface{}) *AuthenticationsamlpolicyResourceModel {
	tflog.Debug(ctx, "In authenticationsamlpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
