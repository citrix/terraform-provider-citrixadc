package appfwsettings

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwsettingsDataSourceModel is the data-source-specific model, decoupled from
// AppfwsettingsResourceModel. appfwsettings is a singleton (no lookup key), so a
// data source is a pure read surface: the configurable attributes are surfaced
// as Computed outputs AND the read-only attributes the resource deliberately
// omits (learning, builtin, feature) are exposed. Every non-key attribute is
// Computed.
type AppfwsettingsDataSourceModel struct {
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
	ProxypasswordWo          types.String `tfsdk:"proxypassword_wo"`
	ProxypasswordWoVersion   types.Int64  `tfsdk:"proxypassword_wo_version"`
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

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/appfwsettings.json). Never settable; populated from
	// GET, Null when the appliance omits them.
	Learning types.String `tfsdk:"learning"`
	Builtin  types.List   `tfsdk:"builtin"`
	Feature  types.String `tfsdk:"feature"`
}

func AppfwsettingsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
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
				Computed:    true,
				Description: "Add the specified flags to AppFW cookies. Available setttings function as follows:\n* None - Do not add flags to AppFW cookies.\n* HTTP Only - Add the HTTP Only flag to AppFW cookies, which prevent scripts from accessing them.\n* Secure - Add Secure flag to AppFW cookies.\n* All - Add both HTTPOnly and Secure flag to AppFW cookies.",
			},
			"cookiepostencryptprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String that is prepended to all encrypted cookie values.",
			},
			"defaultprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "Maximum cumulative size in bytes of all objects imported to Netscaler. The user is not allowed to import an object if the operation exceeds the currently configured limit.",
			},
			"learnratelimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of connections per second that the application firewall learning engine examines to generate new relaxations for learning-enabled security checks. The application firewall drops any connections above this limit from the list of connections used by the learning engine.",
			},
			"logmalformedreq": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
			"proxypassword_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Password with which proxy user logs on.",
			},
			"proxypassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a proxypassword_wo update.",
			},
			"proxyport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "Maximum number of sessions that the application firewall allows to be active, regardless of user activity. After the max_limit reaches, No more user session will be created .",
			},
			"sessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout, in seconds, after which a user session is terminated. Before continuing to use the protected web site, the user must establish a new session by opening a designated start URL.",
			},
			"signatureautoupdate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable auto update signatures",
			},
			"signatureurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to download the mapping file from server",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to use when an application firewall policy evaluates to undefined (UNDEF).\nAn UNDEF event indicates an internal error condition. The APPFW_BLOCK built-in profile is the default setting. You can specify a different built-in or user-created profile as the UNDEF profile.",
			},
			"useconfigurablesecretkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use configurable secret key in AppFw operations",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"learning": schema.StringAttribute{
				Computed:    true,
				Description: "Global learning option that overrides the profile level learning. Possible values: ON, OFF.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if application firewall settings is built-in or not. Possible values: MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// appfwsettingsDataSourceSetAttrFromGet projects a NITRO appfwsettings GET
// response onto the data-source model via the shared utils.MapGet* helpers.
// appfwsettings is a singleton, so the ID is a fixed synthetic value.
func appfwsettingsDataSourceSetAttrFromGet(ctx context.Context, data *AppfwsettingsDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwsettingsDataSourceSetAttrFromGet Function")

	data.Id = types.StringValue("appfwsettings-config")

	data.Ceflogging = utils.MapGetString(g, "ceflogging")
	data.Centralizedlearning = utils.MapGetString(g, "centralizedlearning")
	data.Clientiploggingheader = utils.MapGetString(g, "clientiploggingheader")
	data.Cookieflags = utils.MapGetString(g, "cookieflags")
	data.Cookiepostencryptprefix = utils.MapGetString(g, "cookiepostencryptprefix")
	data.Defaultprofile = utils.MapGetString(g, "defaultprofile")
	data.Entitydecoding = utils.MapGetString(g, "entitydecoding")
	data.Geolocationlogging = utils.MapGetString(g, "geolocationlogging")
	data.Importsizelimit = utils.MapGetInt64(g, "importsizelimit")
	data.Learnratelimit = utils.MapGetInt64(g, "learnratelimit")
	data.Logmalformedreq = utils.MapGetString(g, "logmalformedreq")
	data.Malformedreqaction = utils.MapGetStringList(g, "malformedreqaction")
	data.Proxyport = utils.MapGetInt64(g, "proxyport")
	data.Proxyserver = utils.MapGetString(g, "proxyserver")
	data.Proxyusername = utils.MapGetString(g, "proxyusername")
	data.Sessioncookiename = utils.MapGetString(g, "sessioncookiename")
	data.Sessionlifetime = utils.MapGetInt64(g, "sessionlifetime")
	data.Sessionlimit = utils.MapGetInt64(g, "sessionlimit")
	data.Sessiontimeout = utils.MapGetInt64(g, "sessiontimeout")
	data.Signatureautoupdate = utils.MapGetString(g, "signatureautoupdate")
	data.Signatureurl = utils.MapGetString(g, "signatureurl")
	data.Undefaction = utils.MapGetString(g, "undefaction")
	data.Useconfigurablesecretkey = utils.MapGetString(g, "useconfigurablesecretkey")

	// proxypassword / proxypassword_wo(+version) are write-only secret inputs
	// the GET never returns -> Null.
	data.Proxypassword = types.StringNull()
	data.ProxypasswordWo = types.StringNull()
	data.ProxypasswordWoVersion = types.Int64Null()

	// Read-only attributes.
	data.Learning = utils.MapGetString(g, "learning")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
