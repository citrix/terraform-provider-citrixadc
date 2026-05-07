package authenticationcaptchaaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
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
	SecretkeyWo                types.String `tfsdk:"secretkey_wo"`
	SecretkeyWoVersion         types.Int64  `tfsdk:"secretkey_wo_version"`
	Serverurl                  types.String `tfsdk:"serverurl"`
	Sitekey                    types.String `tfsdk:"sitekey"`
	SitekeyWo                  types.String `tfsdk:"sitekey_wo"`
	SitekeyWoVersion           types.Int64  `tfsdk:"sitekey_wo_version"`
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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the new captcha action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"scorethreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the score threshold value for recaptcha v3.",
			},
			"secretkey": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Secret of gateway as established at the captcha source.",
			},
			"secretkey_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Secret of gateway as established at the captcha source.",
			},
			"secretkey_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a secretkey_wo update.",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the endpoint at which captcha response is validated.",
			},
			"sitekey": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Sitekey to identify gateway fqdn while loading captcha.",
			},
			"sitekey_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Sitekey to identify gateway fqdn while loading captcha.",
			},
			"sitekey_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a sitekey_wo update.",
			},
		},
	}
}

func authenticationcaptchaactionGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationcaptchaactionResourceModel) authentication.Authenticationcaptchaaction {
	tflog.Debug(ctx, "In authenticationcaptchaactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationcaptchaaction := authentication.Authenticationcaptchaaction{}
	if !data.Defaultauthenticationgroup.IsNull() && !data.Defaultauthenticationgroup.IsUnknown() {
		authenticationcaptchaaction.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationcaptchaaction.Name = data.Name.ValueString()
	}
	if !data.Scorethreshold.IsNull() && !data.Scorethreshold.IsUnknown() {
		authenticationcaptchaaction.Scorethreshold = utils.IntPtr(int(data.Scorethreshold.ValueInt64()))
	}
	if !data.Secretkey.IsNull() && !data.Secretkey.IsUnknown() {
		authenticationcaptchaaction.Secretkey = data.Secretkey.ValueString()
	}
	// Skip write-only attribute: secretkey_wo
	// Skip version tracker attribute: secretkey_wo_version
	if !data.Serverurl.IsNull() && !data.Serverurl.IsUnknown() {
		authenticationcaptchaaction.Serverurl = data.Serverurl.ValueString()
	}
	if !data.Sitekey.IsNull() && !data.Sitekey.IsUnknown() {
		authenticationcaptchaaction.Sitekey = data.Sitekey.ValueString()
	}
	// Skip write-only attribute: sitekey_wo
	// Skip version tracker attribute: sitekey_wo_version

	return authenticationcaptchaaction
}

func authenticationcaptchaactionGetThePayloadFromtheConfig(ctx context.Context, data *AuthenticationcaptchaactionResourceModel, payload *authentication.Authenticationcaptchaaction) {
	tflog.Debug(ctx, "In authenticationcaptchaactionGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: secretkey_wo -> secretkey
	if !data.SecretkeyWo.IsNull() {
		secretkeyWo := data.SecretkeyWo.ValueString()
		if secretkeyWo != "" {
			payload.Secretkey = secretkeyWo
		}
	}
	// Handle write-only secret attribute: sitekey_wo -> sitekey
	if !data.SitekeyWo.IsNull() {
		sitekeyWo := data.SitekeyWo.ValueString()
		if sitekeyWo != "" {
			payload.Sitekey = sitekeyWo
		}
	}
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
	// secretkey is not returned by NITRO API (secret/ephemeral) - retain from config
	// secretkey_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// secretkey_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["serverurl"]; ok && val != nil {
		data.Serverurl = types.StringValue(val.(string))
	} else {
		data.Serverurl = types.StringNull()
	}
	// sitekey is not returned by NITRO API (secret/ephemeral) - retain from config
	// sitekey_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// sitekey_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
