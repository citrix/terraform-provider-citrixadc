package auditnslogaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuditnslogactionDataSourceModel is the data-source-specific model, decoupled
// from AuditnslogactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (ip, builtin, feature). Every non-key attribute is Computed; the
// Framework's per-attribute model <-> schema reflection requires this model to
// have exactly the attributes the data-source schema declares, which is why it
// cannot reuse the resource model.
type AuditnslogactionDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Acl                  types.String `tfsdk:"acl"`
	Alg                  types.String `tfsdk:"alg"`
	Appflowexport        types.String `tfsdk:"appflowexport"`
	Contentinspectionlog types.String `tfsdk:"contentinspectionlog"`
	Dateformat           types.String `tfsdk:"dateformat"`
	Denylistviolations   types.String `tfsdk:"denylistviolations"`
	Domainresolvenow     types.Bool   `tfsdk:"domainresolvenow"`
	Domainresolveretry   types.Int64  `tfsdk:"domainresolveretry"`
	Logfacility          types.String `tfsdk:"logfacility"`
	Loglevel             types.List   `tfsdk:"loglevel"`
	Lsn                  types.String `tfsdk:"lsn"`
	Name                 types.String `tfsdk:"name"` // Required lookup key
	Protocolviolations   types.String `tfsdk:"protocolviolations"`
	Serverdomainname     types.String `tfsdk:"serverdomainname"`
	Serverip             types.String `tfsdk:"serverip"`
	Serverport           types.Int64  `tfsdk:"serverport"`
	Sslinterception      types.String `tfsdk:"sslinterception"`
	Subscriberlog        types.String `tfsdk:"subscriberlog"`
	Tcp                  types.String `tfsdk:"tcp"`
	Timezone             types.String `tfsdk:"timezone"`
	Urlfiltering         types.String `tfsdk:"urlfiltering"`
	Userdefinedauditlog  types.String `tfsdk:"userdefinedauditlog"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/auditnslogaction.json). Never settable; populated from GET.
	Ip      types.String `tfsdk:"ip"`
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AuditnslogactionDataSourceSchema() schema.Schema {
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
				Description: "Log Content Inspection event information",
			},
			"dateformat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Format of dates in the logs.\nSupported formats are:\n* MMDDYYYY - U.S. style month/date/year format.\n* DDMMYYYY - European style date/month/year format.\n* YYYYMMDD - ISO style year/month/date format.",
			},
			"denylistviolations": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log denylist violations.",
			},
			"domainresolvenow": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Immediately send a DNS query to resolve the server's domain name.",
			},
			"domainresolveretry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the audit server if the last query failed.",
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
				Description: "Audit log level, which specifies the types of events to log.\nAvailable settings function as follows:\n* ALL - All events.\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.\n* NONE - No events.",
			},
			"lsn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log the LSN messages",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the nslog action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the nslog action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my nslog action\" or 'my nslog action').",
			},
			"protocolviolations": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log protocol violations",
			},
			"serverdomainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Auditserver name as a FQDN. Mutually exclusive with serverIP",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the nslog server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which the nslog server accepts connections.",
			},
			"sslinterception": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log SSL Interception event information",
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
				Description: "Time zone used for date and timestamps in the logs.\nAvailable settings function as follows:\n* GMT_TIME. Coordinated Universal Time.\n* LOCAL_TIME. The server's timezone setting.",
			},
			"urlfiltering": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log URL filtering event information",
			},
			"userdefinedauditlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log user-configurable log messages to nslog.\nSetting this parameter to NO causes auditing to ignore all user-configured message actions. Setting this parameter to YES causes auditing to log user-configured message actions that meet the other logging criteria.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "The resolved IP address of the auditserver.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values: MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// auditnslogactionDataSourceSetAttrFromGet projects a NITRO auditnslogaction GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func auditnslogactionDataSourceSetAttrFromGet(ctx context.Context, data *AuditnslogactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In auditnslogactionDataSourceSetAttrFromGet Function")

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
	data.Domainresolvenow = utils.MapGetBool(g, "domainresolvenow")
	data.Domainresolveretry = utils.MapGetInt64(g, "domainresolveretry")
	data.Logfacility = utils.MapGetString(g, "logfacility")
	data.Loglevel = utils.MapGetStringList(g, "loglevel")
	data.Lsn = utils.MapGetString(g, "lsn")
	data.Protocolviolations = utils.MapGetString(g, "protocolviolations")
	data.Serverdomainname = utils.MapGetString(g, "serverdomainname")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Serverport = utils.MapGetInt64(g, "serverport")
	data.Sslinterception = utils.MapGetString(g, "sslinterception")
	data.Subscriberlog = utils.MapGetString(g, "subscriberlog")
	data.Tcp = utils.MapGetString(g, "tcp")
	data.Timezone = utils.MapGetString(g, "timezone")
	data.Urlfiltering = utils.MapGetString(g, "urlfiltering")
	data.Userdefinedauditlog = utils.MapGetString(g, "userdefinedauditlog")

	// Read-only metadata.
	data.Ip = utils.MapGetString(g, "ip")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
