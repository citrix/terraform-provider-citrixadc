package botsettings

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// BotsettingsResourceModel describes the resource data model.
type BotsettingsResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultnonintrusiveprofile types.String `tfsdk:"defaultnonintrusiveprofile"`
	Defaultprofile             types.String `tfsdk:"defaultprofile"`
	Dfprequestlimit            types.Int64  `tfsdk:"dfprequestlimit"`
	Javascriptname             types.String `tfsdk:"javascriptname"`
	Proxypassword              types.String `tfsdk:"proxypassword"`
	Proxyport                  types.Int64  `tfsdk:"proxyport"`
	Proxyserver                types.String `tfsdk:"proxyserver"`
	Proxyusername              types.String `tfsdk:"proxyusername"`
	Sessioncookiename          types.String `tfsdk:"sessioncookiename"`
	Sessiontimeout             types.Int64  `tfsdk:"sessiontimeout"`
	Signatureautoupdate        types.String `tfsdk:"signatureautoupdate"`
	Signatureurl               types.String `tfsdk:"signatureurl"`
	Trapurlautogenerate        types.String `tfsdk:"trapurlautogenerate"`
	Trapurlinterval            types.Int64  `tfsdk:"trapurlinterval"`
	Trapurllength              types.Int64  `tfsdk:"trapurllength"`
}

func (r *BotsettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botsettings resource.",
			},
			"defaultnonintrusiveprofile": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("BOT_STATS"),
				Description: "Profile to use when the feature is not enabled but feature is licensed. NonIntrusive checks will be disabled and IPRep cronjob(24 Hours) will be removed if this is set to BOT_BYPASS.",
			},
			"defaultprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to use when a connection does not match any policy. Default setting is \" \", which sends unmatched connections back to the Citrix ADC without attempting to filter them further.",
			},
			"dfprequestlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of requests to allow without bot session cookie if device fingerprint is enabled",
			},
			"javascriptname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the JavaScript that the Bot Management feature  uses in response.\nMust begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cookie name\" or 'my cookie name').",
			},
			"proxypassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password with which user logs on.",
			},
			"proxyport": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(8080),
				Description: "Proxy Server Port to get updated signatures from AWS.",
			},
			"proxyserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Server IP to get updated signatures from AWS.",
			},
			"proxyusername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy Username",
			},
			"sessioncookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SessionCookie that the Bot Management feature uses for tracking.\nMust begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cookie name\" or 'my cookie name').",
			},
			"sessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout, in seconds, after which a user session is terminated.",
			},
			"signatureautoupdate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable bot auto update signatures",
			},
			"signatureurl": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("https://nsbotsignatures.s3.amazonaws.com/BotSignatureMapping.json"),
				Description: "URL to download the bot signature mapping file from server",
			},
			"trapurlautogenerate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/disable trap URL auto generation. When enabled, trap URL is updated within the configured interval.",
			},
			"trapurlinterval": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3600),
				Description: "Time in seconds after which trap URL is updated.",
			},
			"trapurllength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(32),
				Description: "Length of the auto-generated trap URL.",
			},
		},
	}
}

func botsettingsGetThePayloadFromtheConfig(ctx context.Context, data *BotsettingsResourceModel) bot.Botsettings {
	tflog.Debug(ctx, "In botsettingsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	botsettings := bot.Botsettings{}
	if !data.Defaultnonintrusiveprofile.IsNull() {
		botsettings.Defaultnonintrusiveprofile = data.Defaultnonintrusiveprofile.ValueString()
	}
	if !data.Defaultprofile.IsNull() {
		botsettings.Defaultprofile = data.Defaultprofile.ValueString()
	}
	if !data.Dfprequestlimit.IsNull() {
		botsettings.Dfprequestlimit = utils.IntPtr(int(data.Dfprequestlimit.ValueInt64()))
	}
	if !data.Javascriptname.IsNull() {
		botsettings.Javascriptname = data.Javascriptname.ValueString()
	}
	if !data.Proxypassword.IsNull() {
		botsettings.Proxypassword = data.Proxypassword.ValueString()
	}
	if !data.Proxyport.IsNull() {
		botsettings.Proxyport = utils.IntPtr(int(data.Proxyport.ValueInt64()))
	}
	if !data.Proxyserver.IsNull() {
		botsettings.Proxyserver = data.Proxyserver.ValueString()
	}
	if !data.Proxyusername.IsNull() {
		botsettings.Proxyusername = data.Proxyusername.ValueString()
	}
	if !data.Sessioncookiename.IsNull() {
		botsettings.Sessioncookiename = data.Sessioncookiename.ValueString()
	}
	if !data.Sessiontimeout.IsNull() {
		botsettings.Sessiontimeout = utils.IntPtr(int(data.Sessiontimeout.ValueInt64()))
	}
	if !data.Signatureautoupdate.IsNull() {
		botsettings.Signatureautoupdate = data.Signatureautoupdate.ValueString()
	}
	if !data.Signatureurl.IsNull() {
		botsettings.Signatureurl = data.Signatureurl.ValueString()
	}
	if !data.Trapurlautogenerate.IsNull() {
		botsettings.Trapurlautogenerate = data.Trapurlautogenerate.ValueString()
	}
	if !data.Trapurlinterval.IsNull() {
		botsettings.Trapurlinterval = utils.IntPtr(int(data.Trapurlinterval.ValueInt64()))
	}
	if !data.Trapurllength.IsNull() {
		botsettings.Trapurllength = utils.IntPtr(int(data.Trapurllength.ValueInt64()))
	}

	return botsettings
}

func botsettingsSetAttrFromGet(ctx context.Context, data *BotsettingsResourceModel, getResponseData map[string]interface{}) *BotsettingsResourceModel {
	tflog.Debug(ctx, "In botsettingsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["defaultnonintrusiveprofile"]; ok && val != nil {
		data.Defaultnonintrusiveprofile = types.StringValue(val.(string))
	} else {
		data.Defaultnonintrusiveprofile = types.StringNull()
	}
	if val, ok := getResponseData["defaultprofile"]; ok && val != nil {
		data.Defaultprofile = types.StringValue(val.(string))
	} else {
		data.Defaultprofile = types.StringNull()
	}
	if val, ok := getResponseData["dfprequestlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dfprequestlimit = types.Int64Value(intVal)
		}
	} else {
		data.Dfprequestlimit = types.Int64Null()
	}
	if val, ok := getResponseData["javascriptname"]; ok && val != nil {
		data.Javascriptname = types.StringValue(val.(string))
	} else {
		data.Javascriptname = types.StringNull()
	}
	if val, ok := getResponseData["proxypassword"]; ok && val != nil {
		data.Proxypassword = types.StringValue(val.(string))
	} else {
		data.Proxypassword = types.StringNull()
	}
	if val, ok := getResponseData["proxyport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Proxyport = types.Int64Value(intVal)
		}
	} else {
		data.Proxyport = types.Int64Null()
	}
	if val, ok := getResponseData["proxyserver"]; ok && val != nil {
		data.Proxyserver = types.StringValue(val.(string))
	} else {
		data.Proxyserver = types.StringNull()
	}
	if val, ok := getResponseData["proxyusername"]; ok && val != nil {
		data.Proxyusername = types.StringValue(val.(string))
	} else {
		data.Proxyusername = types.StringNull()
	}
	if val, ok := getResponseData["sessioncookiename"]; ok && val != nil {
		data.Sessioncookiename = types.StringValue(val.(string))
	} else {
		data.Sessioncookiename = types.StringNull()
	}
	if val, ok := getResponseData["sessiontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sessiontimeout = types.Int64Value(intVal)
		}
	} else {
		data.Sessiontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["signatureautoupdate"]; ok && val != nil {
		data.Signatureautoupdate = types.StringValue(val.(string))
	} else {
		data.Signatureautoupdate = types.StringNull()
	}
	if val, ok := getResponseData["signatureurl"]; ok && val != nil {
		data.Signatureurl = types.StringValue(val.(string))
	} else {
		data.Signatureurl = types.StringNull()
	}
	if val, ok := getResponseData["trapurlautogenerate"]; ok && val != nil {
		data.Trapurlautogenerate = types.StringValue(val.(string))
	} else {
		data.Trapurlautogenerate = types.StringNull()
	}
	if val, ok := getResponseData["trapurlinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Trapurlinterval = types.Int64Value(intVal)
		}
	} else {
		data.Trapurlinterval = types.Int64Null()
	}
	if val, ok := getResponseData["trapurllength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Trapurllength = types.Int64Value(intVal)
		}
	} else {
		data.Trapurllength = types.Int64Null()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("botsettings-config")

	return data
}
