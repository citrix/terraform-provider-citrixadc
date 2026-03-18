package authenticationcaptchaaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationcaptchaactionResourceModel describes the resource data model.
type AuthenticationcaptchaactionResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Name                       types.String `tfsdk:"name"`
	Scorethreshold             types.Int64  `tfsdk:"scorethreshold"`
	Secretkey                  types.String `tfsdk:"secretkey"`
	Serverurl                  types.String `tfsdk:"serverurl"`
	Sitekey                    types.String `tfsdk:"sitekey"`
}

func (r *AuthenticationcaptchaactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationcaptchaaction resource.",
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
				Default:     int64default.StaticInt64(5),
				Description: "This is the score threshold value for recaptcha v3.",
			},
			"secretkey": schema.StringAttribute{
				Required:    true,
				Description: "Secret of gateway as established at the captcha source.",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the endpoint at which captcha response is validated.",
			},
			"sitekey": schema.StringAttribute{
				Required:    true,
				Description: "Sitekey to identify gateway fqdn while loading captcha.",
			},
		},
	}
}

func authenticationcaptchaactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationcaptchaactionResourceModel) authentication.Authenticationcaptchaaction {
	tflog.Debug(ctx, "In authenticationcaptchaactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	authenticationcaptchaaction := authentication.Authenticationcaptchaaction{}
	if !data.Defaultauthenticationgroup.IsNull() {
		authenticationcaptchaaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Name.IsNull() {
		authenticationcaptchaaction.Name = data.Name.ValueString()
	}
	if !data.Scorethreshold.IsNull() {
		authenticationcaptchaaction.Scorethreshold = utils.IntPtr(int(data.Scorethreshold.ValueInt64()))
	}
	if !data.Secretkey.IsNull() {
		authenticationcaptchaaction.Secretkey = data.Secretkey.ValueString()
	}
	if !data.Serverurl.IsNull() {
		authenticationcaptchaaction.Serverurl = data.Serverurl.ValueString()
	}
	if !data.Sitekey.IsNull() {
		authenticationcaptchaaction.Sitekey = data.Sitekey.ValueString()
	}

	return authenticationcaptchaaction
}

func authenticationcaptchaactionSetAttrFromGet(ctx context.Context, data *AuthenticationcaptchaactionResourceModel, getResponseData map[string]interface{}) *AuthenticationcaptchaactionResourceModel {
	tflog.Debug(ctx, "In authenticationcaptchaactionSetAttrFromGet Function")

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
	if val, ok := getResponseData["secretkey"]; ok && val != nil {
		data.Secretkey = types.StringValue(val.(string))
	} else {
		data.Secretkey = types.StringNull()
	}
	if val, ok := getResponseData["serverurl"]; ok && val != nil {
		data.Serverurl = types.StringValue(val.(string))
	} else {
		data.Serverurl = types.StringNull()
	}
	if val, ok := getResponseData["sitekey"]; ok && val != nil {
		data.Sitekey = types.StringValue(val.(string))
	} else {
		data.Sitekey = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
