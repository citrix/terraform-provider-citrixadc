package authenticationldappolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationldappolicyResourceModel describes the resource data model.
type AuthenticationldappolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Reqaction types.String `tfsdk:"reqaction"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *AuthenticationldappolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationldappolicy resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LDAP policy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after LDAP policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy\" or 'my authentication policy').",
			},
			"reqaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LDAP action to perform if the policy matches.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the LDAP server.",
			},
		},
	}
}

func authenticationldappolicyGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationldappolicyResourceModel) authentication.Authenticationldappolicy {
	tflog.Debug(ctx, "In authenticationldappolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationldappolicy := authentication.Authenticationldappolicy{}
	if !data.Name.IsNull() {
		authenticationldappolicy.Name = data.Name.ValueString()
	}
	if !data.Reqaction.IsNull() {
		authenticationldappolicy.Reqaction = data.Reqaction.ValueString()
	}
	if !data.Rule.IsNull() {
		authenticationldappolicy.Rule = data.Rule.ValueString()
	}

	return authenticationldappolicy
}

func authenticationldappolicySetAttrFromGet(ctx context.Context, data *AuthenticationldappolicyResourceModel, getResponseData map[string]interface{}) *AuthenticationldappolicyResourceModel {
	tflog.Debug(ctx, "In authenticationldappolicySetAttrFromGet Function")

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
