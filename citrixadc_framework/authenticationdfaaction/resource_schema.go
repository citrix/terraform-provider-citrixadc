package authenticationdfaaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationdfaactionResourceModel describes the resource data model.
type AuthenticationdfaactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Clientid                   types.String `tfsdk:"clientid"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Name                       types.String `tfsdk:"name"`
	Passphrase                 types.String `tfsdk:"passphrase"`
	Serverurl                  types.String `tfsdk:"serverurl"`
}

func (r *AuthenticationdfaactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationdfaaction resource.",
			},
			"clientid": schema.StringAttribute{
				Required:    true,
				Description: "If configured, this string is sent to the DFA server as the X-Citrix-Exchange header value.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DFA action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the DFA action is added.",
			},
			"passphrase": schema.StringAttribute{
				Required:    true,
				Description: "Key shared between the DFA server and the Citrix ADC.\nRequired to allow the Citrix ADC to communicate with the DFA server.",
			},
			"serverurl": schema.StringAttribute{
				Required:    true,
				Description: "DFA Server URL",
			},
		},
	}
}

func authenticationdfaactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationdfaactionResourceModel) authentication.Authenticationdfaaction {
	tflog.Debug(ctx, "In authenticationdfaactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationdfaaction := authentication.Authenticationdfaaction{}
	if !data.Clientid.IsNull() {
		authenticationdfaaction.Clientid = data.Clientid.ValueString()
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationdfaaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationdfaaction.Name = data.Name.ValueString()
	}
	if !data.Passphrase.IsNull() {
		authenticationdfaaction.Passphrase = data.Passphrase.ValueString()
	}
	if !data.Serverurl.IsNull() {
		authenticationdfaaction.Serverurl = data.Serverurl.ValueString()
	}

	return authenticationdfaaction
}

func authenticationdfaactionSetAttrFromGet(ctx context.Context, data *AuthenticationdfaactionResourceModel, getResponseData map[string]interface{}) *AuthenticationdfaactionResourceModel {
	tflog.Debug(ctx, "In authenticationdfaactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clientid"]; ok && val != nil {
		data.Clientid = types.StringValue(val.(string))
	} else {
		data.Clientid = types.StringNull()
	}
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["passphrase"]; ok && val != nil {
		data.Passphrase = types.StringValue(val.(string))
	} else {
		data.Passphrase = types.StringNull()
	}
	if val, ok := getResponseData["serverurl"]; ok && val != nil {
		data.Serverurl = types.StringValue(val.(string))
	} else {
		data.Serverurl = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
