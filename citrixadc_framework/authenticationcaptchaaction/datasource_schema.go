package authenticationcaptchaaction

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

func AuthenticationcaptchaactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the group that is added to user sessions that match current policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new captcha action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"scorethreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the score threshold value for recaptcha v3.",
			},
			"secretkey": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Secret of gateway as established at the captcha source.",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the endpoint at which captcha response is validated.",
			},
			"sitekey": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Sitekey to identify gateway fqdn while loading captcha.",
			},
		},
	}
}

type AuthenticationcaptchaactionDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Name                       types.String `tfsdk:"name"`
	Scorethreshold             types.Int64  `tfsdk:"scorethreshold"`
	Secretkey                  types.String `tfsdk:"secretkey"`
	Serverurl                  types.String `tfsdk:"serverurl"`
	Sitekey                    types.String `tfsdk:"sitekey"`
}

func authenticationcaptchaactionDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationcaptchaactionDataSourceModel, getResponseData map[string]interface{}) *AuthenticationcaptchaactionDataSourceModel {
	tflog.Debug(ctx, "In authenticationcaptchaactionDataSourceSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["scorethreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Scorethreshold = types.Int64Value(intVal)
		}
	} else {
		data.Scorethreshold = types.Int64Null()
	}
	// secretkey is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["serverurl"]; ok && val != nil {
		data.Serverurl = types.StringValue(val.(string))
	} else {
		data.Serverurl = types.StringNull()
	}
	// sitekey is not returned by NITRO API (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
