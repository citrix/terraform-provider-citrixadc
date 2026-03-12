package authenticationcertaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationcertactionResourceModel describes the resource data model.
type AuthenticationcertactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Groupnamefield             types.String `tfsdk:"groupnamefield"`
	Name                       types.String `tfsdk:"name"`
	Twofactor                  types.String `tfsdk:"twofactor"`
	Usernamefield              types.String `tfsdk:"usernamefield"`
}

func (r *AuthenticationcertactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationcertaction resource.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupnamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client-cert field from which the group is extracted.  Must be set to either \"\"Subject\"\" and \"\"Issuer\"\" (include both sets of double quotation marks).\nFormat: <field>:<subfield>",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the client cert authentication server profile (action).\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after certifcate action is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"twofactor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables or disables two-factor authentication.\nTwo factor authentication is client cert authentication followed by password authentication.",
			},
			"usernamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client-cert field from which the username is extracted. Must be set to either \"\"Subject\"\" and \"\"Issuer\"\" (include both sets of double quotation marks).\nFormat: <field>:<subfield>.",
			},
		},
	}
}

func authenticationcertactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationcertactionResourceModel) authentication.Authenticationcertaction {
	tflog.Debug(ctx, "In authenticationcertactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationcertaction := authentication.Authenticationcertaction{}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationcertaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Groupnamefield.IsNull() {
		authenticationcertaction.Groupnamefield = data.Groupnamefield.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationcertaction.Name = data.Name.ValueString()
	}
	if !data.Twofactor.IsNull() {
		authenticationcertaction.Twofactor = data.Twofactor.ValueString()
	}
	if !data.Usernamefield.IsNull() {
		authenticationcertaction.Usernamefield = data.Usernamefield.ValueString()
	}

	return authenticationcertaction
}

func authenticationcertactionSetAttrFromGet(ctx context.Context, data *AuthenticationcertactionResourceModel, getResponseData map[string]interface{}) *AuthenticationcertactionResourceModel {
	tflog.Debug(ctx, "In authenticationcertactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["groupnamefield"]; ok && val != nil {
		data.Groupnamefield = types.StringValue(val.(string))
	} else {
		data.Groupnamefield = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["twofactor"]; ok && val != nil {
		data.Twofactor = types.StringValue(val.(string))
	} else {
		data.Twofactor = types.StringNull()
	}
	if val, ok := getResponseData["usernamefield"]; ok && val != nil {
		data.Usernamefield = types.StringValue(val.(string))
	} else {
		data.Usernamefield = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
