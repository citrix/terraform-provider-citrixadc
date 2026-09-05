package auditsyslogparams

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuditsyslogparamsDataSourceModel is the data-source-specific model, decoupled
// from AuditsyslogparamsResourceModel. A data source is a pure read surface, so it
// can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes the resource deliberately omits.
type AuditsyslogparamsDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Acl                  types.String `tfsdk:"acl"`
	Alg                  types.String `tfsdk:"alg"`
	Appflowexport        types.String `tfsdk:"appflowexport"`
	Contentinspectionlog types.String `tfsdk:"contentinspectionlog"`
	Dateformat           types.String `tfsdk:"dateformat"`
	Denylistviolations   types.String `tfsdk:"denylistviolations"`
	Dns                  types.String `tfsdk:"dns"`
	Logfacility          types.String `tfsdk:"logfacility"`
	Loglevel             types.List   `tfsdk:"loglevel"`
	Lsn                  types.String `tfsdk:"lsn"`
	Protocolviolations   types.String `tfsdk:"protocolviolations"`
	Serverip             types.String `tfsdk:"serverip"`
	Serverport           types.Int64  `tfsdk:"serverport"`
	Sslinterception      types.String `tfsdk:"sslinterception"`
	Streamanalytics      types.String `tfsdk:"streamanalytics"`
	Subscriberlog        types.String `tfsdk:"subscriberlog"`
	Tcp                  types.String `tfsdk:"tcp"`
	Timezone             types.String `tfsdk:"timezone"`
	Urlfiltering         types.String `tfsdk:"urlfiltering"`
	Userdefinedauditlog  types.String `tfsdk:"userdefinedauditlog"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/auditsyslogparams.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AuditsyslogparamsDataSourceSchema() schema.Schema {
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
				Description: "Log the ALG messages",
			},
			"appflowexport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Export log messages to AppFlow collectors.\nAppflow collectors are entities to which log messages can be sent so that some action can be performed on them.",
			},
			"contentinspectionlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log Content Inspection event ifnormation",
			},
			"dateformat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of dates in the logs.\nSupported formats are:\n* MMDDYYYY - U.S. style month/date/year format.\n* DDMMYYYY. European style  -date/month/year format.\n* YYYYMMDD - ISO style year/month/date format.",
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
			"logfacility": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Facility value, as defined in RFC 3164, assigned to the log message.\nLog facility values are numbers 0 to 7 (LOCAL0 through LOCAL7). Each number indicates where a specific message originated from, such as the Citrix ADC itself, the VPN, or external.",
			},
			"loglevel": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Types of information to be logged.\nAvailable settings function as follows:\n* ALL - All events.\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* NONE - No events.",
			},
			"lsn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log the LSN messages",
			},
			"protocolviolations": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log protocol violations",
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
				Description: "Log SSL Interceptionn event information",
			},
			"streamanalytics": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Export log stream analytics statistics to syslog server",
			},
			"subscriberlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log subscriber session event information",
			},
			"tcp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log TCP messages.",
			},
			"timezone": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time zone used for date and timestamps in the logs.\nAvailable settings function as follows:\n* GMT_TIME - Coordinated Universal Time.\n* LOCAL_TIME  Use the server's timezone setting.",
			},
			"urlfiltering": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log URL filtering event information",
			},
			"userdefinedauditlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log user-configurable log messages to syslog.\nSetting this parameter to NO causes audit to ignore all user-configured message actions. Setting this parameter to YES causes audit to log user-configured message actions that meet the other logging criteria.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
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

// auditsyslogparamsDataSourceSetAttrFromGet projects a NITRO auditsyslogparams GET
// response onto the data-source model. Attributes are filled from the GET (or left
// Null when the GET omits them) using the shared utils.MapGet* helpers.
func auditsyslogparamsDataSourceSetAttrFromGet(ctx context.Context, data *AuditsyslogparamsDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In auditsyslogparamsDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Acl = utils.MapGetString(g, "acl")
	data.Alg = utils.MapGetString(g, "alg")
	data.Appflowexport = utils.MapGetString(g, "appflowexport")
	data.Contentinspectionlog = utils.MapGetString(g, "contentinspectionlog")
	data.Dateformat = utils.MapGetString(g, "dateformat")
	data.Denylistviolations = utils.MapGetString(g, "denylistviolations")
	data.Dns = utils.MapGetString(g, "dns")
	data.Logfacility = utils.MapGetString(g, "logfacility")
	data.Loglevel = utils.MapGetStringList(g, "loglevel")
	data.Lsn = utils.MapGetString(g, "lsn")
	data.Protocolviolations = utils.MapGetString(g, "protocolviolations")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Serverport = utils.MapGetInt64(g, "serverport")
	data.Sslinterception = utils.MapGetString(g, "sslinterception")
	data.Streamanalytics = utils.MapGetString(g, "streamanalytics")
	data.Subscriberlog = utils.MapGetString(g, "subscriberlog")
	data.Tcp = utils.MapGetString(g, "tcp")
	data.Timezone = utils.MapGetString(g, "timezone")
	data.Urlfiltering = utils.MapGetString(g, "urlfiltering")
	data.Userdefinedauditlog = utils.MapGetString(g, "userdefinedauditlog")

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")

	// Singleton config resource has no key; use a static ID.
	data.Id = types.StringValue("auditsyslogparams-config")
}
