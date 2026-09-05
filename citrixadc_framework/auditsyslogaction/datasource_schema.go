package auditsyslogaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuditsyslogactionDataSourceModel is the data-source-specific model, decoupled
// from AuditsyslogactionResourceModel. A data source is a pure read surface, so it
// can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes the resource deliberately omits.
type AuditsyslogactionDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Acl                  types.String `tfsdk:"acl"`
	Alg                  types.String `tfsdk:"alg"`
	Appflowexport        types.String `tfsdk:"appflowexport"`
	Contentinspectionlog types.String `tfsdk:"contentinspectionlog"`
	Dateformat           types.String `tfsdk:"dateformat"`
	Denylistviolations   types.String `tfsdk:"denylistviolations"`
	Dns                  types.String `tfsdk:"dns"`
	Domainresolvenow     types.Bool   `tfsdk:"domainresolvenow"`
	Domainresolveretry   types.Int64  `tfsdk:"domainresolveretry"`
	Httpauthtoken        types.String `tfsdk:"httpauthtoken"`
	Httpendpointurl      types.String `tfsdk:"httpendpointurl"`
	Httpschemafile       types.String `tfsdk:"httpschemafile"`
	Lbvservername        types.String `tfsdk:"lbvservername"`
	Logfacility          types.String `tfsdk:"logfacility"`
	Loglevel             types.Set    `tfsdk:"loglevel"`
	Lsn                  types.String `tfsdk:"lsn"`
	Managementlog        types.List   `tfsdk:"managementlog"`
	Maxlogdatasizetohold types.Int64  `tfsdk:"maxlogdatasizetohold"`
	Mgmtloglevel         types.List   `tfsdk:"mgmtloglevel"`
	Name                 types.String `tfsdk:"name"` // Required lookup key
	Netprofile           types.String `tfsdk:"netprofile"`
	Protocolviolations   types.String `tfsdk:"protocolviolations"`
	Serverdomainname     types.String `tfsdk:"serverdomainname"`
	Serverip             types.String `tfsdk:"serverip"`
	Serverport           types.Int64  `tfsdk:"serverport"`
	Sslinterception      types.String `tfsdk:"sslinterception"`
	Streamanalytics      types.String `tfsdk:"streamanalytics"`
	Subscriberlog        types.String `tfsdk:"subscriberlog"`
	Syslogcompliance     types.String `tfsdk:"syslogcompliance"`
	Tcp                  types.String `tfsdk:"tcp"`
	Tcpprofilename       types.String `tfsdk:"tcpprofilename"`
	Timezone             types.String `tfsdk:"timezone"`
	Transport            types.String `tfsdk:"transport"`
	Urlfiltering         types.String `tfsdk:"urlfiltering"`
	Userdefinedauditlog  types.String `tfsdk:"userdefinedauditlog"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/auditsyslogaction.json). Never settable; populated from GET.
	Ip      types.String `tfsdk:"ip"`
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AuditsyslogactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"acl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log access control list (ACL) messages.",
			},
			"alg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log alg info",
			},
			"appflowexport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Export log messages to AppFlow collectors.\nAppflow collectors are entities to which log messages can be sent so that some action can be performed on them.",
			},
			"contentinspectionlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log Content Inspection event information",
			},
			"dateformat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of dates in the logs.\nSupported formats are:\n* MMDDYYYY. -U.S. style month/date/year format.\n* DDMMYYYY - European style date/month/year format.\n* YYYYMMDD - ISO style year/month/date format.",
			},
			"denylistviolations": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log denylist violations.",
			},
			"dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log DNS related syslog messages",
			},
			"domainresolvenow": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Immediately send a DNS query to resolve the server's domain name.",
			},
			"domainresolveretry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the syslog server if the last query failed.",
			},
			"httpauthtoken": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Token for authenticating with the endpoint. If the endpoint requires the Authorization header in a particular format, specify the complete format as the value to this parameter. For eg., in case of splunk, the Authorization header is required to be of the form - Splunk <auth-token>.",
			},
			"httpendpointurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The URL at which to upload the logs messages on the endpoint",
			},
			"httpschemafile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP Schema file to input tokens to be sent in log message to log server.",
			},
			"lbvservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LB vserver. Mutually exclusive with syslog serverIP/serverName",
			},
			"logfacility": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Facility value, as defined in RFC 3164, assigned to the log message.\nLog facility values are numbers 0 to 7 (LOCAL0 through LOCAL7). Each number indicates where a specific message originated from, such as the Citrix ADC itself, the VPN, or external.",
			},
			"loglevel": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Audit log level, which specifies the types of events to log.\nAvailable values function as follows:\n* ALL - All events.\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.\n* NONE - No events.",
			},
			"lsn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log lsn info",
			},
			"managementlog": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Management log specifies the categories of log files to be exported.\nIt use destination and transport from PE params.\nAvailable values function as follows:\n* ALL - All categories (SHELL, NSMGMT and ACCESS).\n* SHELL -  bash.log, and sh.log.\n* ACCESS - auth.log, nsvpn.log, httpaccess.log, httperror.log, httpaccess-vpn.log and httperror-vpn.log.\n* NSMGMT - notice.log and ns.log.\n* NONE - No logs.",
			},
			"maxlogdatasizetohold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Max size of log data that can be held in NSB chain of server info.",
			},
			"mgmtloglevel": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Management log level, which specifies the types of events to log.\nAvailable values function as follows:\n* ALL - All events.\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.\n* NONE - No events.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the syslog action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the syslog action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my syslog action\" or 'my syslog action').",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network profile.\nThe SNIP configured in the network profile will be used as source IP while sending log messages.",
			},
			"protocolviolations": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log protocol violations",
			},
			"serverdomainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SYSLOG server name as a FQDN.  Mutually exclusive with serverIP/lbVserverName",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the syslog server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which the syslog server accepts connections.",
			},
			"sslinterception": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log SSL Interception event information",
			},
			"streamanalytics": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Export log stream analytics statistics to syslog server.",
			},
			"subscriberlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log subscriber session event information",
			},
			"syslogcompliance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setting this parameter ensures that all the Audit Logs generated for this Syslog Action comply with an RFC. For example, set it to RFC5424 to ensure RFC 5424 compliance",
			},
			"tcp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log TCP messages.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile whose settings are to be applied to the audit server info to tune the TCP connection parameters.",
			},
			"timezone": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time zone used for date and timestamps in the logs.\nSupported settings are:\n* GMT_TIME. Coordinated Universal time.\n* LOCAL_TIME. Use the server's timezone setting.",
			},
			"transport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Transport type used to send auditlogs to syslog server. Default type is UDP.",
			},
			"urlfiltering": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log URL filtering event information",
			},
			"userdefinedauditlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log user-configurable log messages to syslog.\nSetting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "The resolved IP address of the syslog server.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// auditsyslogactionDataSourceSetAttrFromGet projects a NITRO auditsyslogaction GET
// response onto the data-source model. Attributes are filled from the GET (or left
// Null when the GET omits them) using the shared utils.MapGet* helpers.
func auditsyslogactionDataSourceSetAttrFromGet(ctx context.Context, data *AuditsyslogactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In auditsyslogactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Acl = utils.MapGetString(g, "acl")
	data.Alg = utils.MapGetString(g, "alg")
	data.Appflowexport = utils.MapGetString(g, "appflowexport")
	data.Contentinspectionlog = utils.MapGetString(g, "contentinspectionlog")
	data.Dateformat = utils.MapGetString(g, "dateformat")
	data.Denylistviolations = utils.MapGetString(g, "denylistviolations")
	data.Dns = utils.MapGetString(g, "dns")
	data.Domainresolvenow = utils.MapGetBool(g, "domainresolvenow")
	data.Domainresolveretry = utils.MapGetInt64(g, "domainresolveretry")
	data.Httpendpointurl = utils.MapGetString(g, "httpendpointurl")
	data.Httpschemafile = utils.MapGetString(g, "httpschemafile")
	data.Lbvservername = utils.MapGetString(g, "lbvservername")
	data.Logfacility = utils.MapGetString(g, "logfacility")
	data.Lsn = utils.MapGetString(g, "lsn")
	data.Managementlog = utils.MapGetStringList(g, "managementlog")
	data.Maxlogdatasizetohold = utils.MapGetInt64(g, "maxlogdatasizetohold")
	data.Mgmtloglevel = utils.MapGetStringList(g, "mgmtloglevel")
	data.Netprofile = utils.MapGetString(g, "netprofile")
	data.Protocolviolations = utils.MapGetString(g, "protocolviolations")
	data.Serverdomainname = utils.MapGetString(g, "serverdomainname")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Serverport = utils.MapGetInt64(g, "serverport")
	data.Sslinterception = utils.MapGetString(g, "sslinterception")
	data.Streamanalytics = utils.MapGetString(g, "streamanalytics")
	data.Subscriberlog = utils.MapGetString(g, "subscriberlog")
	data.Syslogcompliance = utils.MapGetString(g, "syslogcompliance")
	data.Tcp = utils.MapGetString(g, "tcp")
	data.Tcpprofilename = utils.MapGetString(g, "tcpprofilename")
	data.Timezone = utils.MapGetString(g, "timezone")
	data.Transport = utils.MapGetString(g, "transport")
	data.Urlfiltering = utils.MapGetString(g, "urlfiltering")
	data.Userdefinedauditlog = utils.MapGetString(g, "userdefinedauditlog")

	// loglevel is a Set attribute; project the GET array into a string set.
	if v, ok := g["loglevel"]; ok && v != nil {
		if arr, ok := v.([]interface{}); ok {
			if sv, d := types.SetValueFrom(ctx, types.StringType, utils.ToStringList(arr)); !d.HasError() {
				data.Loglevel = sv
			} else {
				data.Loglevel = types.SetNull(types.StringType)
			}
		} else {
			data.Loglevel = types.SetNull(types.StringType)
		}
	} else {
		data.Loglevel = types.SetNull(types.StringType)
	}

	// httpauthtoken is a secret input the GET never returns -> Null.
	data.Httpauthtoken = types.StringNull()

	// Read-only metadata.
	data.Ip = utils.MapGetString(g, "ip")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
