package authenticationstorefrontauthaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationstorefrontauthactionResourceModel describes the resource data model.
type AuthenticationstorefrontauthactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Domain                     types.String `tfsdk:"domain"`
	Name                       types.String `tfsdk:"name"`
	Serverurl                  types.String `tfsdk:"serverurl"`
}

func (r *AuthenticationstorefrontauthactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationstorefrontauthaction resource.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain of the server that is used for authentication. If users enter name without domain, this parameter is added to username in the authentication request to server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Storefront Authentication action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication action\" or 'my authentication action').",
			},
			"serverurl": schema.StringAttribute{
				Required:    true,
				Description: "URL of the Storefront server. This is the FQDN of the Storefront server. example: https://storefront.com/.  Authentication endpoints are learned dynamically by Gateway.",
			},
		},
	}
}

func authenticationstorefrontauthactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationstorefrontauthactionResourceModel) authentication.Authenticationstorefrontauthaction {
	tflog.Debug(ctx, "In authenticationstorefrontauthactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationstorefrontauthaction := authentication.Authenticationstorefrontauthaction{}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationstorefrontauthaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Domain.IsNull() {
		authenticationstorefrontauthaction.Domain = data.Domain.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationstorefrontauthaction.Name = data.Name.ValueString()
	}
	if !data.Serverurl.IsNull() {
		authenticationstorefrontauthaction.Serverurl = data.Serverurl.ValueString()
	}

	return authenticationstorefrontauthaction
}

func authenticationstorefrontauthactionSetAttrFromGet(ctx context.Context, data *AuthenticationstorefrontauthactionResourceModel, getResponseData map[string]interface{}) *AuthenticationstorefrontauthactionResourceModel {
	tflog.Debug(ctx, "In authenticationstorefrontauthactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
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
