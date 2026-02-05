package authenticationcertpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationcertpolicyResourceModel describes the resource data model.
type AuthenticationcertpolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Reqaction types.String `tfsdk:"reqaction"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *AuthenticationcertpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationcertpolicy resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the client certificate authentication policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after cert authentication policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"reqaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the client cert authentication action to be performed if the policy matches.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the authentication server.",
			},
		},
	}
}

func authenticationcertpolicyGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationcertpolicyResourceModel) authentication.Authenticationcertpolicy {
	tflog.Debug(ctx, "In authenticationcertpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationcertpolicy := authentication.Authenticationcertpolicy{}
	if !data.Name.IsNull() {
		authenticationcertpolicy.Name = data.Name.ValueString()
	}
	if !data.Reqaction.IsNull() {
		authenticationcertpolicy.Reqaction = data.Reqaction.ValueString()
	}
	if !data.Rule.IsNull() {
		authenticationcertpolicy.Rule = data.Rule.ValueString()
	}

	return authenticationcertpolicy
}

func authenticationcertpolicySetAttrFromGet(ctx context.Context, data *AuthenticationcertpolicyResourceModel, getResponseData map[string]interface{}) *AuthenticationcertpolicyResourceModel {
	tflog.Debug(ctx, "In authenticationcertpolicySetAttrFromGet Function")

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
