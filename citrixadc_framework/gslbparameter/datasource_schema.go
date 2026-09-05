package gslbparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GslbparameterDataSourceModel is the data-source-specific model, decoupled from
// GslbparameterResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (flags,
// builtin, feature, incarnation, overridepersistencyfororder). Every non-key
// attribute is Computed.
type GslbparameterDataSourceModel struct {
	Id                        types.String `tfsdk:"id"`
	Automaticconfigsync       types.String `tfsdk:"automaticconfigsync"`
	Dropldnsreq               types.String `tfsdk:"dropldnsreq"`
	Gslbconfigsyncmonitor     types.String `tfsdk:"gslbconfigsyncmonitor"`
	Gslbsvcstatedelaytime     types.Int64  `tfsdk:"gslbsvcstatedelaytime"`
	Gslbsyncinterval          types.Int64  `tfsdk:"gslbsyncinterval"`
	Gslbsynclocfiles          types.String `tfsdk:"gslbsynclocfiles"`
	Gslbsyncmode              types.String `tfsdk:"gslbsyncmode"`
	Gslbsyncsaveconfigcommand types.String `tfsdk:"gslbsyncsaveconfigcommand"`
	Ldnsentrytimeout          types.Int64  `tfsdk:"ldnsentrytimeout"`
	Ldnsmask                  types.String `tfsdk:"ldnsmask"`
	Ldnsprobeorder            types.List   `tfsdk:"ldnsprobeorder"`
	Mepkeepalivetimeout       types.Int64  `tfsdk:"mepkeepalivetimeout"`
	Rtttolerance              types.Int64  `tfsdk:"rtttolerance"`
	Sourceipwhitelisting      types.String `tfsdk:"sourceipwhitelisting"`
	Svcstatelearningtime      types.Int64  `tfsdk:"svcstatelearningtime"`
	Undefaction               types.String `tfsdk:"undefaction"`
	Usekrpcchannelforsync     types.String `tfsdk:"usekrpcchannelforsync"`
	V6ldnsmasklen             types.Int64  `tfsdk:"v6ldnsmasklen"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/gslbparameter.json). Never settable; populated from GET.
	Flags                       types.Int64  `tfsdk:"flags"`
	Builtin                     types.List   `tfsdk:"builtin"`
	Feature                     types.String `tfsdk:"feature"`
	Incarnation                 types.Int64  `tfsdk:"incarnation"`
	Overridepersistencyfororder types.String `tfsdk:"overridepersistencyfororder"`
}

func GslbparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"automaticconfigsync": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "GSLB configuration will be synced automatically to remote gslb sites if enabled.",
			},
			"dropldnsreq": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop LDNS requests if round-trip time (RTT) information is not available.",
			},
			"gslbconfigsyncmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, remote gslb site's rsync port will be monitored and site is considered for configuration sync only when the monitor is successful.",
			},
			"gslbsvcstatedelaytime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of delay in updating the state of GSLB service to DOWN when MEP goes down.\n			This parameter is applicable only if monitors are not bound to GSLB services",
			},
			"gslbsyncinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time duartion (in seconds) for which the gslb sync process will wait before checking for config changes.",
			},
			"gslbsynclocfiles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If disabled, Location files will not be synced to the remote sites as part of manual sync and automatic sync.",
			},
			"gslbsyncmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mode in which configuration will be synced from master site to remote sites.",
			},
			"gslbsyncsaveconfigcommand": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, 'save ns config' command will be treated as other GSLB commands and synced to GSLB nodes when auto gslb sync option is enabled.",
			},
			"ldnsentrytimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which an inactive LDNS entry is removed.",
			},
			"ldnsmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IPv4 network mask with which to create LDNS entries.",
			},
			"ldnsprobeorder": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Order in which monitors should be initiated to calculate RTT.",
			},
			"mepkeepalivetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time duartion (in seconds) during which if no new packets received by Local gslb site from Remote gslb site then mark the MEP connection DOWN",
			},
			"rtttolerance": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Tolerance, in milliseconds, for newly learned round-trip time (RTT) values. If the difference between the old RTT value and the newly computed RTT value is less than or equal to the specified tolerance value, the LDNS entry in the network metric table is not updated with the new RTT value. Prevents the exchange of metrics when variations in RTT values are negligible.",
			},
			"sourceipwhitelisting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, local gslb site private IP would be used as the source IP while initiating MEP/GSLB sync connection if srcIP is not configured for GSLB site.",
			},
			"svcstatelearningtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time (in seconds) within which local or child site services remain in learning phase. GSLB site will enter the learning phase after reboot, HA failover, Cluster GSLB owner node changes or MEP being enabled on local node.  Backup parent (if configured) will selectively move the adopted children's GSLB services to learning phase when primary parent goes down. While a service is in learning period, remote site will not honour the state and stats got through MEP for that service. State can be learnt from health monitor if bound explicitly.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform when policy evaluation creates an UNDEF condition. Available settings function as follows:\n* NOLBACTION - Does not consider LB action in making LB decision.\n* RESET - Reset the request and notify the user, so that the user can resend the request.\n* DROP - Drop the request without sending a response to the user.",
			},
			"usekrpcchannelforsync": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is to use Krpc channel for GSLB sync.",
			},
			"v6ldnsmasklen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Mask for creating LDNS entries for IPv6 source addresses. The mask is defined as the number of leading bits to consider, in the source IP address, when creating an LDNS entry.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "State of the GSLB parameter.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"incarnation": schema.Int64Attribute{
				Computed:    true,
				Description: "This is a counter to maintain the gslb sync incarnation number.",
			},
			"overridepersistencyfororder": schema.StringAttribute{
				Computed:    true,
				Description: "This option is used to override persistency when order is configured for services or servicegroups.",
			},
		},
	}
}

// gslbparameterDataSourceSetAttrFromGet projects a NITRO gslbparameter GET
// response onto the data-source model. Attributes are filled from the GET (or
// left Null when the GET omits them) via the shared utils.MapGet* helpers.
func gslbparameterDataSourceSetAttrFromGet(ctx context.Context, data *GslbparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In gslbparameterDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Automaticconfigsync = utils.MapGetString(g, "automaticconfigsync")
	data.Dropldnsreq = utils.MapGetString(g, "dropldnsreq")
	data.Gslbconfigsyncmonitor = utils.MapGetString(g, "gslbconfigsyncmonitor")
	data.Gslbsvcstatedelaytime = utils.MapGetInt64(g, "gslbsvcstatedelaytime")
	data.Gslbsyncinterval = utils.MapGetInt64(g, "gslbsyncinterval")
	data.Gslbsynclocfiles = utils.MapGetString(g, "gslbsynclocfiles")
	data.Gslbsyncmode = utils.MapGetString(g, "gslbsyncmode")
	data.Gslbsyncsaveconfigcommand = utils.MapGetString(g, "gslbsyncsaveconfigcommand")
	data.Ldnsentrytimeout = utils.MapGetInt64(g, "ldnsentrytimeout")
	data.Ldnsmask = utils.MapGetString(g, "ldnsmask")
	data.Ldnsprobeorder = utils.MapGetStringList(g, "ldnsprobeorder")
	data.Mepkeepalivetimeout = utils.MapGetInt64(g, "mepkeepalivetimeout")
	data.Rtttolerance = utils.MapGetInt64(g, "rtttolerance")
	data.Sourceipwhitelisting = utils.MapGetString(g, "sourceipwhitelisting")
	data.Svcstatelearningtime = utils.MapGetInt64(g, "svcstatelearningtime")
	data.Undefaction = utils.MapGetString(g, "undefaction")
	data.Usekrpcchannelforsync = utils.MapGetString(g, "usekrpcchannelforsync")
	data.V6ldnsmasklen = utils.MapGetInt64(g, "v6ldnsmasklen")

	// Read-only attributes.
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Incarnation = utils.MapGetInt64(g, "incarnation")
	data.Overridepersistencyfororder = utils.MapGetString(g, "overridepersistencyfororder")

	// Singleton (unnamed) resource - static ID matching the resource behavior.
	data.Id = types.StringValue("gslbparameter-config")
}
