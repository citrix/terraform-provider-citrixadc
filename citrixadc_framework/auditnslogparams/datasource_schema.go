package auditnslogparams

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuditnslogparamsDataSourceModel is the data-source-specific model, decoupled
// from AuditnslogparamsResourceModel.
//
// auditnslogparams is a singleton config entity, so the data source has no
// lookup key. It is a pure read surface (Read only; no plan/apply lifecycle),
// exposing the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (builtin, feature). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type AuditnslogparamsDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Acl                  types.String `tfsdk:"acl"`
	Alg                  types.String `tfsdk:"alg"`
	Appflowexport        types.String `tfsdk:"appflowexport"`
	Contentinspectionlog types.String `tfsdk:"contentinspectionlog"`
	Dateformat           types.String `tfsdk:"dateformat"`
	Denylistviolations   types.String `tfsdk:"denylistviolations"`
	Logfacility          types.String `tfsdk:"logfacility"`
	Loglevel             types.List   `tfsdk:"loglevel"`
	Lsn                  types.String `tfsdk:"lsn"`
	Protocolviolations   types.String `tfsdk:"protocolviolations"`
	Serverip             types.String `tfsdk:"serverip"`
	Serverport           types.Int64  `tfsdk:"serverport"`
	Sslinterception      types.String `tfsdk:"sslinterception"`
	Subscriberlog        types.String `tfsdk:"subscriberlog"`
	Tcp                  types.String `tfsdk:"tcp"`
	Timezone             types.String `tfsdk:"timezone"`
	Urlfiltering         types.String `tfsdk:"urlfiltering"`
	Userdefinedauditlog  types.String `tfsdk:"userdefinedauditlog"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/auditnslogparams.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func AuditnslogparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"acl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure auditing to log access control list (ACL) messages.",
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
				Description: "Log denylist violations",
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
				Description: "Types of information to be logged.\nAvailable settings function as follows:\n* ALL - All events.\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.\n* NONE - No events.",
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
				Description: "Configure auditing to log TCP messages.",
			},
			"timezone": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time zone used for date and timestamps in the logs.\nSupported settings are:\n* GMT_TIME - Coordinated Universal Time.\n* LOCAL_TIME - Use the server's timezone setting.",
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

// auditnslogparamsDataSourceSetAttrFromGet projects a NITRO auditnslogparams GET
// response onto the data-source model. auditnslogparams is a singleton, so the
// ID is the same static value used by the resource. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func auditnslogparamsDataSourceSetAttrFromGet(ctx context.Context, data *AuditnslogparamsDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In auditnslogparamsDataSourceSetAttrFromGet Function")

	// Singleton config entity - static ID (matches the resource).
	data.Id = types.StringValue("auditnslogparams-config")

	// Read/write attributes as read-back outputs.
	data.Acl = utils.MapGetString(g, "acl")
	data.Alg = utils.MapGetString(g, "alg")
	data.Appflowexport = utils.MapGetString(g, "appflowexport")
	data.Contentinspectionlog = utils.MapGetString(g, "contentinspectionlog")
	data.Dateformat = utils.MapGetString(g, "dateformat")
	data.Denylistviolations = utils.MapGetString(g, "denylistviolations")
	data.Logfacility = utils.MapGetString(g, "logfacility")
	data.Loglevel = utils.MapGetStringList(g, "loglevel")
	data.Lsn = utils.MapGetString(g, "lsn")
	data.Protocolviolations = utils.MapGetString(g, "protocolviolations")
	data.Serverip = utils.MapGetString(g, "serverip")
	data.Serverport = utils.MapGetInt64(g, "serverport")
	data.Sslinterception = utils.MapGetString(g, "sslinterception")
	data.Subscriberlog = utils.MapGetString(g, "subscriberlog")
	data.Tcp = utils.MapGetString(g, "tcp")
	data.Timezone = utils.MapGetString(g, "timezone")
	data.Urlfiltering = utils.MapGetString(g, "urlfiltering")
	data.Userdefinedauditlog = utils.MapGetString(g, "userdefinedauditlog")

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
