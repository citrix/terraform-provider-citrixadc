package aaaparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaaparameterDataSourceModel is the data-source-specific model, decoupled from
// AaaparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type AaaparameterDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Aaadloglevel               types.String `tfsdk:"aaadloglevel"`
	Aaadnatip                  types.String `tfsdk:"aaadnatip"`
	Aaasessionloglevel         types.String `tfsdk:"aaasessionloglevel"`
	Apitokencache              types.String `tfsdk:"apitokencache"`
	Classicendpoints           types.String `tfsdk:"classicendpoints"`
	Defaultauthtype            types.String `tfsdk:"defaultauthtype"`
	Defaultcspheader           types.String `tfsdk:"defaultcspheader"`
	Dynaddr                    types.String `tfsdk:"dynaddr"`
	Enableenhancedauthfeedback types.String `tfsdk:"enableenhancedauthfeedback"`
	Enablesessionstickiness    types.String `tfsdk:"enablesessionstickiness"`
	Enablestaticpagecaching    types.String `tfsdk:"enablestaticpagecaching"`
	Enhancedepa                types.String `tfsdk:"enhancedepa"`
	Failedlogintimeout         types.Int64  `tfsdk:"failedlogintimeout"`
	Ftmode                     types.String `tfsdk:"ftmode"`
	Httponlycookie             types.String `tfsdk:"httponlycookie"`
	Loginencryption            types.String `tfsdk:"loginencryption"`
	Maxaaausers                types.Int64  `tfsdk:"maxaaausers"`
	Maxkbquestions             types.Int64  `tfsdk:"maxkbquestions"`
	Maxloginattempts           types.Int64  `tfsdk:"maxloginattempts"`
	Maxsamldeflatesize         types.Int64  `tfsdk:"maxsamldeflatesize"`
	Persistentloginattempts    types.String `tfsdk:"persistentloginattempts"`
	Pwdexpirynotificationdays  types.Int64  `tfsdk:"pwdexpirynotificationdays"`
	Samesite                   types.String `tfsdk:"samesite"`
	Securityinsights           types.String `tfsdk:"securityinsights"`
	Tokenintrospectioninterval types.Int64  `tfsdk:"tokenintrospectioninterval"`
	Wafprotection              types.List   `tfsdk:"wafprotection"`
	Webviewendpoints           types.String `tfsdk:"webviewendpoints"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/aaaparameter.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AaaparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aaadloglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "AAAD log level, which specifies the types of AAAD events to log in nsvpn.log.\nAvailable values function as follows:\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.",
			},
			"aaadnatip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IP address to use for traffic that is sent to the authentication server.",
			},
			"aaasessionloglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audit log level, which specifies the types of events to log for cli executed commands.\nAvailable values function as follows:\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.",
			},
			"apitokencache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to enable/disable API cache feature.",
			},
			"classicendpoints": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to enable/disable classic endpoints.",
			},
			"defaultauthtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The default authentication server type.",
			},
			"defaultcspheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to enable/disable default CSP header",
			},
			"dynaddr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set by the DHCP client when the IP address was fetched dynamically.",
			},
			"enableenhancedauthfeedback": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enhanced auth feedback provides more information to the end user about the reason for an authentication failure.  The default value is set to NO.",
			},
			"enablesessionstickiness": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables/Disables stickiness to authentication servers",
			},
			"enablestaticpagecaching": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The default state of VPN Static Page caching. Static Page caching is enabled by default.",
			},
			"enhancedepa": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to enable/disable EPA v2 functionality",
			},
			"failedlogintimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of minutes an account will be locked if user exceeds maximum permissible attempts",
			},
			"ftmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "First time user mode determines which configuration options are shown by default when logging in to the GUI. This setting is controlled by the GUI.",
			},
			"httponlycookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to set/reset HttpOnly Flag for NSC_AAAC/NSC_TMAS cookies in nfactor",
			},
			"loginencryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to encrypt login information for nFactor flow",
			},
			"maxaaausers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent users allowed to log on to VPN simultaneously.",
			},
			"maxkbquestions": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This will set maximum number of Questions to be asked for KB Validation. Default value is 2, Max Value is 6",
			},
			"maxloginattempts": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum Number of login Attempts",
			},
			"maxsamldeflatesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This will set the maximum deflate size in case of SAML Redirect binding.",
			},
			"persistentloginattempts": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistent storage of unsuccessful user login attempts",
			},
			"pwdexpirynotificationdays": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This will set the threshold time in days for password expiry notification. Default value is 0, which means no notification is sent",
			},
			"samesite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SameSite attribute value for Cookies generated in AAATM context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite",
			},
			"securityinsights": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the security insight records to the configured collectors when request comes to Authentication endpoint.\n* If cs vserver is frontend with Authentication vserver as target for cs action, then record is sent using Authentication vserver name.\n* If vpn/lb/cs vserver are configured with Authentication ON, then then record is sent using vpn/lb/cs vserver name accordingly.\n* If authentication vserver is frontend, then record is sent using Authentication vserver name.",
			},
			"tokenintrospectioninterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Frequency at which a token must be verified at the Authorization Server (AS) despite being found in cache.",
			},
			"wafprotection": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Entities for which WAF Protection need to be applied.\nAvailable settings function as follows:\n* DEFAULT - No Endpoint WAF protection.\n* AUTH - Endpoints used for Authentication applicable for both AAATM, IDP, GATEWAY use cases.\n* VPN - Endpoints used for Gateway use cases.\n* PORTAL - Endpoints related to web portal.\n* DISABLED - No Endpoint WAF protection.\nCurrently supported only in default partition",
			},
			"webviewendpoints": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to enable/disable webview endpoints.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if aaa param is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// aaaparameterDataSourceSetAttrFromGet projects a NITRO aaaparameter GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func aaaparameterDataSourceSetAttrFromGet(ctx context.Context, data *AaaparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaaparameterDataSourceSetAttrFromGet Function")

	// aaaparameter is a singleton; use a static ID.
	data.Id = types.StringValue("aaaparameter-config")

	// Read/write attributes as read-back outputs.
	data.Aaadloglevel = utils.MapGetString(g, "aaadloglevel")
	data.Aaadnatip = utils.MapGetString(g, "aaadnatip")
	data.Aaasessionloglevel = utils.MapGetString(g, "aaasessionloglevel")
	data.Apitokencache = utils.MapGetString(g, "apitokencache")
	data.Classicendpoints = utils.MapGetString(g, "classicendpoints")
	data.Defaultauthtype = utils.MapGetString(g, "defaultauthtype")
	data.Defaultcspheader = utils.MapGetString(g, "defaultcspheader")
	data.Dynaddr = utils.MapGetString(g, "dynaddr")
	data.Enableenhancedauthfeedback = utils.MapGetString(g, "enableenhancedauthfeedback")
	data.Enablesessionstickiness = utils.MapGetString(g, "enablesessionstickiness")
	data.Enablestaticpagecaching = utils.MapGetString(g, "enablestaticpagecaching")
	data.Enhancedepa = utils.MapGetString(g, "enhancedepa")
	data.Failedlogintimeout = utils.MapGetInt64(g, "failedlogintimeout")
	data.Ftmode = utils.MapGetString(g, "ftmode")
	data.Httponlycookie = utils.MapGetString(g, "httponlycookie")
	data.Loginencryption = utils.MapGetString(g, "loginencryption")
	data.Maxaaausers = utils.MapGetInt64(g, "maxaaausers")
	data.Maxkbquestions = utils.MapGetInt64(g, "maxkbquestions")
	data.Maxloginattempts = utils.MapGetInt64(g, "maxloginattempts")
	data.Maxsamldeflatesize = utils.MapGetInt64(g, "maxsamldeflatesize")
	data.Persistentloginattempts = utils.MapGetString(g, "persistentloginattempts")
	data.Pwdexpirynotificationdays = utils.MapGetInt64(g, "pwdexpirynotificationdays")
	data.Samesite = utils.MapGetString(g, "samesite")
	data.Securityinsights = utils.MapGetString(g, "securityinsights")
	data.Tokenintrospectioninterval = utils.MapGetInt64(g, "tokenintrospectioninterval")
	data.Wafprotection = utils.MapGetStringList(g, "wafprotection")
	data.Webviewendpoints = utils.MapGetString(g, "webviewendpoints")

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
