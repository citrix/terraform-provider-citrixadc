package botprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/bot"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// BotprofileResourceModel describes the resource data model.
type BotprofileResourceModel struct {
	Id                                     types.String `tfsdk:"id"`
	Addcookieflags                         types.String `tfsdk:"addcookieflags"`
	BotEnableBlackList                     types.String `tfsdk:"bot_enable_black_list"`
	BotEnableIpReputation                  types.String `tfsdk:"bot_enable_ip_reputation"`
	BotEnableRateLimit                     types.String `tfsdk:"bot_enable_rate_limit"`
	BotEnableTps                           types.String `tfsdk:"bot_enable_tps"`
	BotEnableWhiteList                     types.String `tfsdk:"bot_enable_white_list"`
	Clientipexpression                     types.String `tfsdk:"clientipexpression"`
	Comment                                types.String `tfsdk:"comment"`
	Devicefingerprint                      types.String `tfsdk:"devicefingerprint"`
	Devicefingerprintaction                types.List   `tfsdk:"devicefingerprintaction"`
	Devicefingerprintmobile                types.List   `tfsdk:"devicefingerprintmobile"`
	Dfprequestlimit                        types.Int64  `tfsdk:"dfprequestlimit"`
	Errorurl                               types.String `tfsdk:"errorurl"`
	Headlessbrowserdetection               types.String `tfsdk:"headlessbrowserdetection"`
	Kmdetection                            types.String `tfsdk:"kmdetection"`
	Kmeventspostbodylimit                  types.Int64  `tfsdk:"kmeventspostbodylimit"`
	Kmjavascriptname                       types.String `tfsdk:"kmjavascriptname"`
	Name                                   types.String `tfsdk:"name"`
	Sessioncookiename                      types.String `tfsdk:"sessioncookiename"`
	Sessiontimeout                         types.Int64  `tfsdk:"sessiontimeout"`
	Signature                              types.String `tfsdk:"signature"`
	Signaturemultipleuseragentheaderaction types.List   `tfsdk:"signaturemultipleuseragentheaderaction"`
	Signaturenouseragentheaderaction       types.List   `tfsdk:"signaturenouseragentheaderaction"`
	Spoofedreqaction                       types.List   `tfsdk:"spoofedreqaction"`
	Trap                                   types.String `tfsdk:"trap"`
	Trapaction                             types.List   `tfsdk:"trapaction"`
	Trapurl                                types.String `tfsdk:"trapurl"`
	Verboseloglevel                        types.String `tfsdk:"verboseloglevel"`
}

func (r *BotprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the botprofile resource.",
			},
			"addcookieflags": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("httpOnly"),
				Description: "Add the specified flags to bot session cookies. Available settings function as follows:\n* None - Do not add flags to cookies.\n* HTTP Only - Add the HTTP Only flag to cookies, which prevents scripts from accessing cookies.\n* Secure - Add Secure flag to cookies.\n* All - Add both HTTPOnly and Secure flags to cookies.",
			},
			"bot_enable_black_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Enable black-list bot detection.",
			},
			"bot_enable_ip_reputation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Enable IP-reputation bot detection.",
			},
			"bot_enable_rate_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Enable rate-limit bot detection.",
			},
			"bot_enable_tps": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Enable TPS.",
			},
			"bot_enable_white_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Enable white-list bot detection.",
			},
			"clientipexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression to get the client IP.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"devicefingerprint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Enable device-fingerprint bot detection",
			},
			"devicefingerprintaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Action to be taken for device-fingerprint based bot detection.",
			},
			"devicefingerprintmobile": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Enabling bot device fingerprint protection for mobile clients",
			},
			"dfprequestlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of requests to allow without bot session cookie if device fingerprint is enabled",
			},
			"errorurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL that Bot protection uses as the Error URL.",
			},
			"headlessbrowserdetection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Enable Headless Browser detection.",
			},
			"kmdetection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable keyboard-mouse based bot detection.",
			},
			"kmeventspostbodylimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Size of the KM data send by the browser, needs to be processed on ADC",
			},
			"kmjavascriptname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the JavaScript file that the Bot Management feature will insert in the response for keyboard-mouse based detection.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my javascript file name\" or 'my javascript file name').",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
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
			"signature": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of object containing bot static signature details.",
			},
			"signaturemultipleuseragentheaderaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Actions to be taken if multiple User-Agent headers are seen in a request (Applicable if Signature check is enabled). Log action should be combined with other actions",
			},
			"signaturenouseragentheaderaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Actions to be taken if no User-Agent header in the request (Applicable if Signature check is enabled).",
			},
			"spoofedreqaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Actions to be taken on a spoofed request (A request spoofing good bot user agent string).",
			},
			"trap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "Enable trap bot detection.",
			},
			"trapaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Action to be taken for bot trap based bot detection.",
			},
			"trapurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL that Bot protection uses as the Trap URL.",
			},
			"verboseloglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NONE"),
				Description: "Bot verbose Logging. Based on the log level, ADC will log additional information whenever client is detected as a bot.",
			},
		},
	}
}

func botprofileGetThePayloadFromthePlan(ctx context.Context, data *BotprofileResourceModel) bot.Botprofile {
	tflog.Debug(ctx, "In botprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	botprofile := bot.Botprofile{}
	if !data.Addcookieflags.IsNull() && !data.Addcookieflags.IsUnknown() {
		botprofile.Addcookieflags = data.Addcookieflags.ValueString()
	}
	if !data.BotEnableBlackList.IsNull() && !data.BotEnableBlackList.IsUnknown() {
		botprofile.Botenableblacklist = data.BotEnableBlackList.ValueString()
	}
	if !data.BotEnableIpReputation.IsNull() && !data.BotEnableIpReputation.IsUnknown() {
		botprofile.Botenableipreputation = data.BotEnableIpReputation.ValueString()
	}
	if !data.BotEnableRateLimit.IsNull() && !data.BotEnableRateLimit.IsUnknown() {
		botprofile.Botenableratelimit = data.BotEnableRateLimit.ValueString()
	}
	if !data.BotEnableTps.IsNull() && !data.BotEnableTps.IsUnknown() {
		botprofile.Botenabletps = data.BotEnableTps.ValueString()
	}
	if !data.BotEnableWhiteList.IsNull() && !data.BotEnableWhiteList.IsUnknown() {
		botprofile.Botenablewhitelist = data.BotEnableWhiteList.ValueString()
	}
	if !data.Clientipexpression.IsNull() && !data.Clientipexpression.IsUnknown() {
		botprofile.Clientipexpression = data.Clientipexpression.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		botprofile.Comment = data.Comment.ValueString()
	}
	if !data.Devicefingerprint.IsNull() && !data.Devicefingerprint.IsUnknown() {
		botprofile.Devicefingerprint = data.Devicefingerprint.ValueString()
	}
	if !data.Devicefingerprintaction.IsNull() && !data.Devicefingerprintaction.IsUnknown() {
		var listValue []string
		data.Devicefingerprintaction.ElementsAs(ctx, &listValue, false)
		botprofile.Devicefingerprintaction = listValue
	}
	if !data.Devicefingerprintmobile.IsNull() && !data.Devicefingerprintmobile.IsUnknown() {
		var listValue []string
		data.Devicefingerprintmobile.ElementsAs(ctx, &listValue, false)
		botprofile.Devicefingerprintmobile = listValue
	}
	if !data.Dfprequestlimit.IsNull() && !data.Dfprequestlimit.IsUnknown() {
		botprofile.Dfprequestlimit = utils.IntPtr(int(data.Dfprequestlimit.ValueInt64()))
	}
	if !data.Errorurl.IsNull() && !data.Errorurl.IsUnknown() {
		botprofile.Errorurl = data.Errorurl.ValueString()
	}
	if !data.Headlessbrowserdetection.IsNull() && !data.Headlessbrowserdetection.IsUnknown() {
		botprofile.Headlessbrowserdetection = data.Headlessbrowserdetection.ValueString()
	}
	if !data.Kmdetection.IsNull() && !data.Kmdetection.IsUnknown() {
		botprofile.Kmdetection = data.Kmdetection.ValueString()
	}
	if !data.Kmeventspostbodylimit.IsNull() && !data.Kmeventspostbodylimit.IsUnknown() {
		botprofile.Kmeventspostbodylimit = utils.IntPtr(int(data.Kmeventspostbodylimit.ValueInt64()))
	}
	if !data.Kmjavascriptname.IsNull() && !data.Kmjavascriptname.IsUnknown() {
		botprofile.Kmjavascriptname = data.Kmjavascriptname.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		botprofile.Name = data.Name.ValueString()
	}
	if !data.Sessioncookiename.IsNull() && !data.Sessioncookiename.IsUnknown() {
		botprofile.Sessioncookiename = data.Sessioncookiename.ValueString()
	}
	if !data.Sessiontimeout.IsNull() && !data.Sessiontimeout.IsUnknown() {
		botprofile.Sessiontimeout = utils.IntPtr(int(data.Sessiontimeout.ValueInt64()))
	}
	if !data.Signature.IsNull() && !data.Signature.IsUnknown() {
		botprofile.Signature = data.Signature.ValueString()
	}
	if !data.Signaturemultipleuseragentheaderaction.IsNull() && !data.Signaturemultipleuseragentheaderaction.IsUnknown() {
		var listValue []string
		data.Signaturemultipleuseragentheaderaction.ElementsAs(ctx, &listValue, false)
		botprofile.Signaturemultipleuseragentheaderaction = listValue
	}
	if !data.Signaturenouseragentheaderaction.IsNull() && !data.Signaturenouseragentheaderaction.IsUnknown() {
		var listValue []string
		data.Signaturenouseragentheaderaction.ElementsAs(ctx, &listValue, false)
		botprofile.Signaturenouseragentheaderaction = listValue
	}
	if !data.Spoofedreqaction.IsNull() && !data.Spoofedreqaction.IsUnknown() {
		var listValue []string
		data.Spoofedreqaction.ElementsAs(ctx, &listValue, false)
		botprofile.Spoofedreqaction = listValue
	}
	if !data.Trap.IsNull() && !data.Trap.IsUnknown() {
		botprofile.Trap = data.Trap.ValueString()
	}
	if !data.Trapaction.IsNull() && !data.Trapaction.IsUnknown() {
		var listValue []string
		data.Trapaction.ElementsAs(ctx, &listValue, false)
		botprofile.Trapaction = listValue
	}
	if !data.Trapurl.IsNull() && !data.Trapurl.IsUnknown() {
		botprofile.Trapurl = data.Trapurl.ValueString()
	}
	if !data.Verboseloglevel.IsNull() && !data.Verboseloglevel.IsUnknown() {
		botprofile.Verboseloglevel = data.Verboseloglevel.ValueString()
	}

	return botprofile
}

func botprofileSetAttrFromGet(ctx context.Context, data *BotprofileResourceModel, getResponseData map[string]interface{}) *BotprofileResourceModel {
	tflog.Debug(ctx, "In botprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["addcookieflags"]; ok && val != nil {
		data.Addcookieflags = types.StringValue(val.(string))
	} else {
		data.Addcookieflags = types.StringNull()
	}
	if val, ok := getResponseData["bot_enable_black_list"]; ok && val != nil {
		data.BotEnableBlackList = types.StringValue(val.(string))
	} else {
		data.BotEnableBlackList = types.StringNull()
	}
	if val, ok := getResponseData["bot_enable_ip_reputation"]; ok && val != nil {
		data.BotEnableIpReputation = types.StringValue(val.(string))
	} else {
		data.BotEnableIpReputation = types.StringNull()
	}
	if val, ok := getResponseData["bot_enable_rate_limit"]; ok && val != nil {
		data.BotEnableRateLimit = types.StringValue(val.(string))
	} else {
		data.BotEnableRateLimit = types.StringNull()
	}
	if val, ok := getResponseData["bot_enable_tps"]; ok && val != nil {
		data.BotEnableTps = types.StringValue(val.(string))
	} else {
		data.BotEnableTps = types.StringNull()
	}
	if val, ok := getResponseData["bot_enable_white_list"]; ok && val != nil {
		data.BotEnableWhiteList = types.StringValue(val.(string))
	} else {
		data.BotEnableWhiteList = types.StringNull()
	}
	if val, ok := getResponseData["clientipexpression"]; ok && val != nil {
		data.Clientipexpression = types.StringValue(val.(string))
	} else {
		data.Clientipexpression = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["devicefingerprint"]; ok && val != nil {
		data.Devicefingerprint = types.StringValue(val.(string))
	} else {
		data.Devicefingerprint = types.StringNull()
	}
	if val, ok := getResponseData["devicefingerprintaction"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			stringList := utils.ToStringList(v)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Devicefingerprintaction = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Devicefingerprintaction = listValue
		default:
			if data.Devicefingerprintaction.IsUnknown() {
				data.Devicefingerprintaction = types.ListNull(types.StringType)
			}
		}
	} else if data.Devicefingerprintaction.IsUnknown() {
		data.Devicefingerprintaction = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["devicefingerprintmobile"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			stringList := utils.ToStringList(v)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Devicefingerprintmobile = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Devicefingerprintmobile = listValue
		default:
			if data.Devicefingerprintmobile.IsUnknown() {
				data.Devicefingerprintmobile = types.ListNull(types.StringType)
			}
		}
	} else if data.Devicefingerprintmobile.IsUnknown() {
		data.Devicefingerprintmobile = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["dfprequestlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dfprequestlimit = types.Int64Value(intVal)
		}
	} else if data.Dfprequestlimit.IsUnknown() {
		data.Dfprequestlimit = types.Int64Null()
	}
	if val, ok := getResponseData["errorurl"]; ok && val != nil {
		data.Errorurl = types.StringValue(val.(string))
	} else {
		data.Errorurl = types.StringNull()
	}
	if val, ok := getResponseData["headlessbrowserdetection"]; ok && val != nil {
		data.Headlessbrowserdetection = types.StringValue(val.(string))
	} else {
		data.Headlessbrowserdetection = types.StringNull()
	}
	if val, ok := getResponseData["kmdetection"]; ok && val != nil {
		data.Kmdetection = types.StringValue(val.(string))
	} else {
		data.Kmdetection = types.StringNull()
	}
	if val, ok := getResponseData["kmeventspostbodylimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Kmeventspostbodylimit = types.Int64Value(intVal)
		}
	} else if data.Kmeventspostbodylimit.IsUnknown() {
		data.Kmeventspostbodylimit = types.Int64Null()
	}
	if val, ok := getResponseData["kmjavascriptname"]; ok && val != nil {
		data.Kmjavascriptname = types.StringValue(val.(string))
	} else {
		data.Kmjavascriptname = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
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
	} else if data.Sessiontimeout.IsUnknown() {
		data.Sessiontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["signature"]; ok && val != nil {
		data.Signature = types.StringValue(val.(string))
	} else {
		data.Signature = types.StringNull()
	}
	if val, ok := getResponseData["signaturemultipleuseragentheaderaction"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			stringList := utils.ToStringList(v)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Signaturemultipleuseragentheaderaction = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Signaturemultipleuseragentheaderaction = listValue
		default:
			if data.Signaturemultipleuseragentheaderaction.IsUnknown() {
				data.Signaturemultipleuseragentheaderaction = types.ListNull(types.StringType)
			}
		}
	} else if data.Signaturemultipleuseragentheaderaction.IsUnknown() {
		data.Signaturemultipleuseragentheaderaction = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["signaturenouseragentheaderaction"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			stringList := utils.ToStringList(v)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Signaturenouseragentheaderaction = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Signaturenouseragentheaderaction = listValue
		default:
			if data.Signaturenouseragentheaderaction.IsUnknown() {
				data.Signaturenouseragentheaderaction = types.ListNull(types.StringType)
			}
		}
	} else if data.Signaturenouseragentheaderaction.IsUnknown() {
		data.Signaturenouseragentheaderaction = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["spoofedreqaction"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			stringList := utils.ToStringList(v)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Spoofedreqaction = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Spoofedreqaction = listValue
		default:
			if data.Spoofedreqaction.IsUnknown() {
				data.Spoofedreqaction = types.ListNull(types.StringType)
			}
		}
	} else if data.Spoofedreqaction.IsUnknown() {
		data.Spoofedreqaction = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["trap"]; ok && val != nil {
		data.Trap = types.StringValue(val.(string))
	} else {
		data.Trap = types.StringNull()
	}
	if val, ok := getResponseData["trapaction"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			stringList := utils.ToStringList(v)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Trapaction = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Trapaction = listValue
		default:
			if data.Trapaction.IsUnknown() {
				data.Trapaction = types.ListNull(types.StringType)
			}
		}
	} else if data.Trapaction.IsUnknown() {
		data.Trapaction = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["trapurl"]; ok && val != nil {
		data.Trapurl = types.StringValue(val.(string))
	} else {
		data.Trapurl = types.StringNull()
	}
	if val, ok := getResponseData["verboseloglevel"]; ok && val != nil {
		data.Verboseloglevel = types.StringValue(val.(string))
	} else {
		data.Verboseloglevel = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value (name) as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
