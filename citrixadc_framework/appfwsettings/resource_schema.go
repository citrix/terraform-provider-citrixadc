package appfwsettings

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwsettingsResourceModel describes the resource data model.
type AppfwsettingsResourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Ceflogging               types.String `tfsdk:"ceflogging"`
	Centralizedlearning      types.String `tfsdk:"centralizedlearning"`
	Clientiploggingheader    types.String `tfsdk:"clientiploggingheader"`
	Cookieflags              types.String `tfsdk:"cookieflags"`
	Cookiepostencryptprefix  types.String `tfsdk:"cookiepostencryptprefix"`
	Defaultprofile           types.String `tfsdk:"defaultprofile"`
	Entitydecoding           types.String `tfsdk:"entitydecoding"`
	Geolocationlogging       types.String `tfsdk:"geolocationlogging"`
	Importsizelimit          types.Int64  `tfsdk:"importsizelimit"`
	Learnratelimit           types.Int64  `tfsdk:"learnratelimit"`
	Logmalformedreq          types.String `tfsdk:"logmalformedreq"`
	Malformedreqaction       types.List   `tfsdk:"malformedreqaction"`
	Proxypassword            types.String `tfsdk:"proxypassword"`
	Proxyport                types.Int64  `tfsdk:"proxyport"`
	Proxyserver              types.String `tfsdk:"proxyserver"`
	Proxyusername            types.String `tfsdk:"proxyusername"`
	Sessioncookiename        types.String `tfsdk:"sessioncookiename"`
	Sessionlifetime          types.Int64  `tfsdk:"sessionlifetime"`
	Sessionlimit             types.Int64  `tfsdk:"sessionlimit"`
	Sessiontimeout           types.Int64  `tfsdk:"sessiontimeout"`
	Signatureautoupdate      types.String `tfsdk:"signatureautoupdate"`
	Signatureurl             types.String `tfsdk:"signatureurl"`
	Undefaction              types.String `tfsdk:"undefaction"`
	Useconfigurablesecretkey types.String `tfsdk:"useconfigurablesecretkey"`
}

func (r *AppfwsettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwsettings resource.",
			},
			"ceflogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable CEF format logs.",
			},
			"centralizedlearning": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable ADM centralized learning",
			},
			"clientiploggingheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of an HTTP header that contains the IP address that the client used to connect to the protected web site or service.",
			},
			"cookieflags": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "Add the specified flags to AppFW cookies. Available setttings function as follows:\n* None - Do not add flags to AppFW cookies.\n* HTTP Only - Add the HTTP Only flag to AppFW cookies, which prevent scripts from accessing them.\n* Secure - Add Secure flag to AppFW cookies.\n* All - Add both HTTPOnly and Secure flag to AppFW cookies.",
			},
			"cookiepostencryptprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String that is prepended to all encrypted cookie values.",
			},
			"defaultprofile": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("APPFW_BYPASS"),
				Description: "Profile to use when a connection does not match any policy. Default setting is APPFW_BYPASS, which sends unmatched connections back to the Citrix ADC without attempting to filter them further.",
			},
			"entitydecoding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Transform multibyte (double- or half-width) characters to single width characters.",
			},
			"geolocationlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable Geo-Location Logging in CEF format logs.",
			},
			"importsizelimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(134217728),
				Description: "Maximum cumulative size in bytes of all objects imported to Netscaler. The user is not allowed to import an object if the operation exceeds the currently configured limit.",
			},
			"learnratelimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(400),
				Description: "Maximum number of connections per second that the application firewall learning engine examines to generate new relaxations for learning-enabled security checks. The application firewall drops any connections above this limit from the list of connections used by the learning engine.",
			},
			"logmalformedreq": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Log requests that are so malformed that application firewall parsing doesn't occur.",
			},
			"malformedreqaction": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "flag to define action on malformed requests that application firewall cannot parse",
			},
			"proxypassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password with which proxy user logs on.",
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
				Description: "Name of the session cookie that the application firewall uses to track user sessions.\nMust begin with a letter or number, and can consist of from 1 to 31 letters, numbers, and the hyphen (-) and underscore (_) symbols.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cookie name\" or 'my cookie name').",
			},
			"sessionlifetime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum amount of time (in seconds) that the application firewall allows a user session to remain active, regardless of user activity. After this time, the user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL. A value of 0 represents infinite time.",
			},
			"sessionlimit": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(100000),
				Description: "Maximum number of sessions that the application firewall allows to be active, regardless of user activity. After the max_limit reaches, No more user session will be created .",
			},
			"sessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(900),
				Description: "Timeout, in seconds, after which a user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL.",
			},
			"signatureautoupdate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable auto update signatures",
			},
			"signatureurl": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("https://s3.amazonaws.com/NSAppFwSignatures/SignaturesMapping.xml"),
				Description: "URL to download the mapping file from server",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("APPFW_BLOCK"),
				Description: "Profile to use when an application firewall policy evaluates to undefined (UNDEF).\nAn UNDEF event indicates an internal error condition. The APPFW_BLOCK built-in profile is the default setting. You can specify a different built-in or user-created profile as the UNDEF profile.",
			},
			"useconfigurablesecretkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use configurable secret key in AppFw operations",
			},
		},
	}
}

func appfwsettingsGetThePayloadFromtheConfig(ctx context.Context, data *AppfwsettingsResourceModel) appfw.Appfwsettings {
	tflog.Debug(ctx, "In appfwsettingsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwsettings := appfw.Appfwsettings{}
	if !data.Ceflogging.IsNull() {
		appfwsettings.Ceflogging = data.Ceflogging.ValueString()
	}
	if !data.Centralizedlearning.IsNull() {
		appfwsettings.Centralizedlearning = data.Centralizedlearning.ValueString()
	}
	if !data.Clientiploggingheader.IsNull() {
		appfwsettings.Clientiploggingheader = data.Clientiploggingheader.ValueString()
	}
	if !data.Cookieflags.IsNull() {
		appfwsettings.Cookieflags = data.Cookieflags.ValueString()
	}
	if !data.Cookiepostencryptprefix.IsNull() {
		appfwsettings.Cookiepostencryptprefix = data.Cookiepostencryptprefix.ValueString()
	}
	if !data.Defaultprofile.IsNull() {
		appfwsettings.Defaultprofile = data.Defaultprofile.ValueString()
	}
	if !data.Entitydecoding.IsNull() {
		appfwsettings.Entitydecoding = data.Entitydecoding.ValueString()
	}
	if !data.Geolocationlogging.IsNull() {
		appfwsettings.Geolocationlogging = data.Geolocationlogging.ValueString()
	}
	if !data.Importsizelimit.IsNull() {
		appfwsettings.Importsizelimit = utils.IntPtr(int(data.Importsizelimit.ValueInt64()))
	}
	if !data.Learnratelimit.IsNull() {
		appfwsettings.Learnratelimit = utils.IntPtr(int(data.Learnratelimit.ValueInt64()))
	}
	if !data.Logmalformedreq.IsNull() {
		appfwsettings.Logmalformedreq = data.Logmalformedreq.ValueString()
	}
	if !data.Proxypassword.IsNull() {
		appfwsettings.Proxypassword = data.Proxypassword.ValueString()
	}
	if !data.Proxyport.IsNull() {
		appfwsettings.Proxyport = utils.IntPtr(int(data.Proxyport.ValueInt64()))
	}
	if !data.Proxyserver.IsNull() {
		appfwsettings.Proxyserver = data.Proxyserver.ValueString()
	}
	if !data.Proxyusername.IsNull() {
		appfwsettings.Proxyusername = data.Proxyusername.ValueString()
	}
	if !data.Sessioncookiename.IsNull() {
		appfwsettings.Sessioncookiename = data.Sessioncookiename.ValueString()
	}
	if !data.Sessionlifetime.IsNull() {
		appfwsettings.Sessionlifetime = utils.IntPtr(int(data.Sessionlifetime.ValueInt64()))
	}
	if !data.Sessionlimit.IsNull() {
		appfwsettings.Sessionlimit = utils.IntPtr(int(data.Sessionlimit.ValueInt64()))
	}
	if !data.Sessiontimeout.IsNull() {
		appfwsettings.Sessiontimeout = utils.IntPtr(int(data.Sessiontimeout.ValueInt64()))
	}
	if !data.Signatureautoupdate.IsNull() {
		appfwsettings.Signatureautoupdate = data.Signatureautoupdate.ValueString()
	}
	if !data.Signatureurl.IsNull() {
		appfwsettings.Signatureurl = data.Signatureurl.ValueString()
	}
	if !data.Undefaction.IsNull() {
		appfwsettings.Undefaction = data.Undefaction.ValueString()
	}
	if !data.Useconfigurablesecretkey.IsNull() {
		appfwsettings.Useconfigurablesecretkey = data.Useconfigurablesecretkey.ValueString()
	}

	return appfwsettings
}

func appfwsettingsSetAttrFromGet(ctx context.Context, data *AppfwsettingsResourceModel, getResponseData map[string]interface{}) *AppfwsettingsResourceModel {
	tflog.Debug(ctx, "In appfwsettingsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ceflogging"]; ok && val != nil {
		data.Ceflogging = types.StringValue(val.(string))
	} else {
		data.Ceflogging = types.StringNull()
	}
	if val, ok := getResponseData["centralizedlearning"]; ok && val != nil {
		data.Centralizedlearning = types.StringValue(val.(string))
	} else {
		data.Centralizedlearning = types.StringNull()
	}
	if val, ok := getResponseData["clientiploggingheader"]; ok && val != nil {
		data.Clientiploggingheader = types.StringValue(val.(string))
	} else {
		data.Clientiploggingheader = types.StringNull()
	}
	if val, ok := getResponseData["cookieflags"]; ok && val != nil {
		data.Cookieflags = types.StringValue(val.(string))
	} else {
		data.Cookieflags = types.StringNull()
	}
	if val, ok := getResponseData["cookiepostencryptprefix"]; ok && val != nil {
		data.Cookiepostencryptprefix = types.StringValue(val.(string))
	} else {
		data.Cookiepostencryptprefix = types.StringNull()
	}
	if val, ok := getResponseData["defaultprofile"]; ok && val != nil {
		data.Defaultprofile = types.StringValue(val.(string))
	} else {
		data.Defaultprofile = types.StringNull()
	}
	if val, ok := getResponseData["entitydecoding"]; ok && val != nil {
		data.Entitydecoding = types.StringValue(val.(string))
	} else {
		data.Entitydecoding = types.StringNull()
	}
	if val, ok := getResponseData["geolocationlogging"]; ok && val != nil {
		data.Geolocationlogging = types.StringValue(val.(string))
	} else {
		data.Geolocationlogging = types.StringNull()
	}
	if val, ok := getResponseData["importsizelimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Importsizelimit = types.Int64Value(intVal)
		}
	} else {
		data.Importsizelimit = types.Int64Null()
	}
	if val, ok := getResponseData["learnratelimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Learnratelimit = types.Int64Value(intVal)
		}
	} else {
		data.Learnratelimit = types.Int64Null()
	}
	if val, ok := getResponseData["logmalformedreq"]; ok && val != nil {
		data.Logmalformedreq = types.StringValue(val.(string))
	} else {
		data.Logmalformedreq = types.StringNull()
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
	if val, ok := getResponseData["sessionlifetime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sessionlifetime = types.Int64Value(intVal)
		}
	} else {
		data.Sessionlifetime = types.Int64Null()
	}
	if val, ok := getResponseData["sessionlimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sessionlimit = types.Int64Value(intVal)
		}
	} else {
		data.Sessionlimit = types.Int64Null()
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
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}
	if val, ok := getResponseData["useconfigurablesecretkey"]; ok && val != nil {
		data.Useconfigurablesecretkey = types.StringValue(val.(string))
	} else {
		data.Useconfigurablesecretkey = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("appfwsettings-config")

	return data
}
