package systemparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemparameterDataSourceModel is the data-source-specific model, decoupled
// from SystemparameterResourceModel. systemparameter is a singleton (keyless)
// config object. A data source is a pure read surface, so it can expose the FULL
// GET projection: the configurable attributes (as Computed outputs) AND the
// read-only metadata the appliance returns on GET (allowdefaultpartition) that
// the resource omits.
type SystemparameterDataSourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Basicauth               types.String `tfsdk:"basicauth"`
	Cliloglevel             types.String `tfsdk:"cliloglevel"`
	Daystoexpire            types.Int64  `tfsdk:"daystoexpire"`
	Denylist                types.String `tfsdk:"denylist"`
	Denylistlogging         types.String `tfsdk:"denylistlogging"`
	Doppler                 types.String `tfsdk:"doppler"`
	Fipsusermode            types.String `tfsdk:"fipsusermode"`
	Forcepasswordchange     types.String `tfsdk:"forcepasswordchange"`
	Googleanalytics         types.String `tfsdk:"googleanalytics"`
	Localauth               types.String `tfsdk:"localauth"`
	Maxclient               types.String `tfsdk:"maxclient"`
	Maxsessionperuser       types.Int64  `tfsdk:"maxsessionperuser"`
	Minpasswordlen          types.Int64  `tfsdk:"minpasswordlen"`
	Natpcbforceflushlimit   types.Int64  `tfsdk:"natpcbforceflushlimit"`
	Natpcbrstontimeout      types.String `tfsdk:"natpcbrstontimeout"`
	Passwordhistorycontrol  types.String `tfsdk:"passwordhistorycontrol"`
	Promptstring            types.String `tfsdk:"promptstring"`
	Pwdhistorycount         types.Int64  `tfsdk:"pwdhistorycount"`
	Rbaonresponse           types.String `tfsdk:"rbaonresponse"`
	Reauthonauthparamchange types.String `tfsdk:"reauthonauthparamchange"`
	Removesensitivefiles    types.String `tfsdk:"removesensitivefiles"`
	Restrictedtimeout       types.String `tfsdk:"restrictedtimeout"`
	Strongpassword          types.String `tfsdk:"strongpassword"`
	Timeout                 types.Int64  `tfsdk:"timeout"`
	Totalauthtimeout        types.Int64  `tfsdk:"totalauthtimeout"`
	Wafprotection           types.List   `tfsdk:"wafprotection"`
	Warnpriorndays          types.Int64  `tfsdk:"warnpriorndays"`

	// Read-only (GET-only) metadata from the NITRO read-only set
	// (zion73x_readonly/systemparameter.json). Never settable; populated from GET.
	Allowdefaultpartition types.String `tfsdk:"allowdefaultpartition"`
}

func SystemparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"basicauth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable basic authentication for Nitro API.",
			},
			"cliloglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audit log level, which specifies the types of events to log for cli executed commands.\nAvailable values function as follows:\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.",
			},
			"daystoexpire": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Password expiry days for all the system users. The daystoexpire value ranges from 30 to 255.",
			},
			"denylist": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable denylist protection.",
			},
			"denylistlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable denylist protection logging.",
			},
			"doppler": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable Doppler",
			},
			"fipsusermode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to set the FIPS mode for key user-land processes. When enabled, these user-land processes will operate in FIPS mode. In this mode, these processes will use FIPS 140-2 certified crypto algorithms.\nWith a FIPS license, it is enabled by default and cannot be disabled.\nWithout a FIPS license, it is disabled by default, wherein these user-land processes will not operate in FIPS mode.",
			},
			"forcepasswordchange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable force password change for nsroot user",
			},
			"googleanalytics": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable Google analytics",
			},
			"localauth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When enabled, local users can access Citrix ADC even when external authentication is configured. When disabled, local users are not allowed to access the Citrix ADC, Local users can access the Citrix ADC only when the configured external authentication servers are unavailable. This parameter is not applicable to SSH Key-based authentication",
			},
			"maxclient": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of client connection allowed per user. Exposed for backward compatibility with the SDK v2 resource; this is a read-only NITRO parameter.",
			},
			"maxsessionperuser": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of client connection allowed per user.The maxsessionperuser value ranges from 1 to 40",
			},
			"minpasswordlen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum length of system user password. When strong password is enabled default minimum length is 8. User entered value can be greater than or equal to 8. Default mininum value is 1 when strong password is disabled. Maximum value is 127 in both cases.",
			},
			"natpcbforceflushlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush the system if the number of Network Address Translation Protocol Control Blocks (NATPCBs) exceeds this value.",
			},
			"natpcbrstontimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send a reset signal to client and server connections when their NATPCBs time out. Avoids the buildup of idle TCP connections on both the sides.",
			},
			"passwordhistorycontrol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables or disable password expiry feature for system users.\nIf the feature is ENABLED, by default the last 6 passwords of users will be maintained and will not be allowed to reuse same.\nWhen the feature is enabled the daystoexpire, warnpriorndays and pwdhistoryCount will be set with default values. The values can only be set in system\nfor system parameter. It cannot be unset. It is possible to set and unset the values for daytoexpire and warnpriorndays in system groups.\nDefault values if feature is ENABLED:\ndaystoexpire: 30\nwarnpriorndays: 5\npwdhistoryCount: 6\nIf the feature is DISABLED the values cannot be set or unset in system parameter and system groups",
			},
			"promptstring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String to display at the command-line prompt. Can consist of letters, numbers, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), underscore (_), and the following variables:\n* %u - Will be replaced by the user name.\n* %h - Will be replaced by the hostname of the Citrix ADC.\n* %t - Will be replaced by the current time in 12-hour format.\n* %T - Will be replaced by the current time in 24-hour format.\n* %d - Will be replaced by the current date.\n* %s - Will be replaced by the state of the Citrix ADC.\n\nNote: The 63-character limit for the length of the string does not apply to the characters that replace the variables.",
			},
			"pwdhistorycount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of passwords to be maintained as history for system users. The pwdhistorycount value ranges from 1 to 10.",
			},
			"rbaonresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable Role-Based Authentication (RBA) on responses.",
			},
			"reauthonauthparamchange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable External user reauthentication when authentication parameter changes",
			},
			"removesensitivefiles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to remove the sensitive files from the system like authorise keys, public keys etc. The commands which will remove sensitive files when this system paramter is enabled are rm cluster instance, rm cluster node, rm ha node, clear config full, join cluster and add cluster instance.",
			},
			"restrictedtimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable the restricted timeout behaviour. When enabled, timeout cannot be configured beyond admin configured timeout  and also it will have the [minimum - maximum] range check. When disabled, timeout will have the old behaviour. By default the value is disabled",
			},
			"strongpassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "After enabling strong password (enableall / enablelocal - not included in exclude list), all the passwords / sensitive information must have - Atleast 1 Lower case character, Atleast 1 Upper case character, Atleast 1 numeric character, Atleast 1 special character ( ~, `, !, @, #, $, %, ^, &, *, -, _, =, +, {, }, [, ], |, \\, :, <, >, /, ., ,, \" \"). Exclude list in case of enablelocal is - NS_FIPS, NS_CRL, NS_RSAKEY, NS_PKCS12, NS_PKCS8, NS_LDAP, NS_TACACS, NS_TACACSACTION, NS_RADIUS, NS_RADIUSACTION, NS_ENCRYPTION_PARAMS. So no Strong Password checks will be performed on these ObjectType commands for enablelocal case.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "CLI session inactivity timeout, in seconds. If Restrictedtimeout argument is enabled, Timeout can have values in the range [300-86400] seconds.\nIf Restrictedtimeout argument is disabled, Timeout can have values in the range [0, 10-100000000] seconds. Default value is 900 seconds.",
			},
			"totalauthtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Total time a request can take for authentication/authorization",
			},
			"wafprotection": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Configure WAF protection for endpoints used by NetScaler management interfaces. The available options are:\n* DEFAULT - NetScaler determines which endpoints have WAF protection enabled or disabled. In the current release, WAF protection is disabled for all endpoints when this option is used. The behavior of this option may change in future releases.\n* GUI - Endpoints used by the Management GUI Interface are WAF protected.\n* DISABLED - WAF protection is disabled for all endpoints.",
			},
			"warnpriorndays": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of days before which password expiration warning would be thrown with respect to daystoexpire. The warnpriorndays value ranges from 5 to 40.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (this is
			// intentionally NOT modeled on the resource). Computed.
			"allowdefaultpartition": schema.StringAttribute{
				Computed:    true,
				Description: "Enable/Disable the allowing partition users to access default partition. Possible values: YES, NO.",
			},
		},
	}
}

// systemparameterDataSourceSetAttrFromGet projects a NITRO systemparameter GET
// response onto the data-source model. Attributes are filled from the GET (or
// left Null when the GET omits them) via the shared utils.MapGet* helpers.
func systemparameterDataSourceSetAttrFromGet(ctx context.Context, data *SystemparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In systemparameterDataSourceSetAttrFromGet Function")

	// Configurable attributes as read-back outputs.
	data.Basicauth = utils.MapGetString(g, "basicauth")
	data.Cliloglevel = utils.MapGetString(g, "cliloglevel")
	data.Daystoexpire = utils.MapGetInt64(g, "daystoexpire")
	data.Denylist = utils.MapGetString(g, "denylist")
	data.Denylistlogging = utils.MapGetString(g, "denylistlogging")
	data.Doppler = utils.MapGetString(g, "doppler")
	data.Fipsusermode = utils.MapGetString(g, "fipsusermode")
	data.Forcepasswordchange = utils.MapGetString(g, "forcepasswordchange")
	data.Googleanalytics = utils.MapGetString(g, "googleanalytics")
	data.Localauth = utils.MapGetString(g, "localauth")
	data.Maxclient = utils.MapGetString(g, "maxclient")
	data.Maxsessionperuser = utils.MapGetInt64(g, "maxsessionperuser")
	data.Minpasswordlen = utils.MapGetInt64(g, "minpasswordlen")
	data.Natpcbforceflushlimit = utils.MapGetInt64(g, "natpcbforceflushlimit")
	data.Natpcbrstontimeout = utils.MapGetString(g, "natpcbrstontimeout")
	data.Passwordhistorycontrol = utils.MapGetString(g, "passwordhistorycontrol")
	data.Promptstring = utils.MapGetString(g, "promptstring")
	data.Pwdhistorycount = utils.MapGetInt64(g, "pwdhistorycount")
	data.Rbaonresponse = utils.MapGetString(g, "rbaonresponse")
	data.Reauthonauthparamchange = utils.MapGetString(g, "reauthonauthparamchange")
	data.Removesensitivefiles = utils.MapGetString(g, "removesensitivefiles")
	data.Restrictedtimeout = utils.MapGetString(g, "restrictedtimeout")
	data.Strongpassword = utils.MapGetString(g, "strongpassword")
	data.Timeout = utils.MapGetInt64(g, "timeout")
	data.Totalauthtimeout = utils.MapGetInt64(g, "totalauthtimeout")
	data.Wafprotection = utils.MapGetStringList(g, "wafprotection")
	data.Warnpriorndays = utils.MapGetInt64(g, "warnpriorndays")

	// Read-only metadata.
	data.Allowdefaultpartition = utils.MapGetString(g, "allowdefaultpartition")

	// Singleton resource -> static ID (matches the resource/datasource contract).
	data.Id = types.StringValue("systemparameter-config")
}
