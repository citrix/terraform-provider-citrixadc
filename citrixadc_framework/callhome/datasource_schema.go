package callhome

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CallhomeDataSourceModel is the data-source-specific model, decoupled from
// CallhomeResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type CallhomeDataSourceModel struct {
	Id types.String `tfsdk:"id"`

	// Existing read/write attributes, surfaced here as Computed outputs.
	Emailaddress     types.String `tfsdk:"emailaddress"`
	Hbcustominterval types.Int64  `tfsdk:"hbcustominterval"`
	Ipaddress        types.String `tfsdk:"ipaddress"`
	Mode             types.String `tfsdk:"mode"`
	Nodeid           types.Int64  `tfsdk:"nodeid"`
	Port             types.Int64  `tfsdk:"port"`
	Proxyauthservice types.String `tfsdk:"proxyauthservice"`
	Proxymode        types.String `tfsdk:"proxymode"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/callhome.json). Never settable; populated from GET.
	Sslcardfirstfailure  types.String `tfsdk:"sslcardfirstfailure"`
	Sslcardlatestfailure types.String `tfsdk:"sslcardlatestfailure"`
	Powfirstfail         types.String `tfsdk:"powfirstfail"`
	Powlatestfailure     types.String `tfsdk:"powlatestfailure"`
	Hddfirstfail         types.String `tfsdk:"hddfirstfail"`
	Hddlatestfailure     types.String `tfsdk:"hddlatestfailure"`
	Flashfirstfail       types.String `tfsdk:"flashfirstfail"`
	Flashlatestfailure   types.String `tfsdk:"flashlatestfailure"`
	Rlfirsthighdrop      types.String `tfsdk:"rlfirsthighdrop"`
	Rllatesthighdrop     types.String `tfsdk:"rllatesthighdrop"`
	Restartlatestfail    types.String `tfsdk:"restartlatestfail"`
	Memthrefirstanomaly  types.String `tfsdk:"memthrefirstanomaly"`
	Memthrelatestanomaly types.String `tfsdk:"memthrelatestanomaly"`
	Callhomestatus       types.List   `tfsdk:"callhomestatus"`
	Anomalydetection     types.String `tfsdk:"anomalydetection"`
}

func CallhomeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Data source to read Call Home configuration.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"emailaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Email address of the contact administrator.",
			},
			"hbcustominterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval (in days) between CallHome heartbeats",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the proxy server.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CallHome mode of operation",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP port on the Proxy server. This is a mandatory parameter for both IP address and service name based configuration.",
			},
			"proxyauthservice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the service that represents the proxy server.",
			},
			"proxymode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables or disables the proxy mode. The proxy server can be set by either specifying the IP address of the server or the name of the service representing the proxy server.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"sslcardfirstfailure": schema.StringAttribute{
				Computed:    true,
				Description: "First occurrence SSL card failure.",
			},
			"sslcardlatestfailure": schema.StringAttribute{
				Computed:    true,
				Description: "Latest occurrence SSL card failure.",
			},
			"powfirstfail": schema.StringAttribute{
				Computed:    true,
				Description: "First occurrence power supply unit failure.",
			},
			"powlatestfailure": schema.StringAttribute{
				Computed:    true,
				Description: "Latest occurrence power supply unit failure.",
			},
			"hddfirstfail": schema.StringAttribute{
				Computed:    true,
				Description: "First occurrence hard disk drive failure.",
			},
			"hddlatestfailure": schema.StringAttribute{
				Computed:    true,
				Description: "Latest occurrence hard disk drive failure.",
			},
			"flashfirstfail": schema.StringAttribute{
				Computed:    true,
				Description: "First occurrence compact flash failure.",
			},
			"flashlatestfailure": schema.StringAttribute{
				Computed:    true,
				Description: "Latest occurrence compact flush failure.",
			},
			"rlfirsthighdrop": schema.StringAttribute{
				Computed:    true,
				Description: "First occurence of high rate limit drops.",
			},
			"rllatesthighdrop": schema.StringAttribute{
				Computed:    true,
				Description: "Latest occurence of high rate limit drops.",
			},
			"restartlatestfail": schema.StringAttribute{
				Computed:    true,
				Description: "Latest occurrence warm restart failure.",
			},
			"memthrefirstanomaly": schema.StringAttribute{
				Computed:    true,
				Description: "First occurrence of memory anomaly.",
			},
			"memthrelatestanomaly": schema.StringAttribute{
				Computed:    true,
				Description: "Latest occurrence of memory anomaly.",
			},
			"callhomestatus": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Callhome feature enabled/disable, register with upload server successful/failed.",
			},
			"anomalydetection": schema.StringAttribute{
				Computed:    true,
				Description: "Enables or disables anomaly detection.",
			},
		},
	}
}

// callhomeDataSourceSetAttrFromGet projects a NITRO callhome GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection. callhome is
// a singleton, so the ID is a static string.
func callhomeDataSourceSetAttrFromGet(ctx context.Context, data *CallhomeDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In callhomeDataSourceSetAttrFromGet Function")

	// callhome is a singleton (no unique lookup key) -> static ID.
	data.Id = types.StringValue("callhome")

	// Existing read/write attributes as read-back outputs.
	data.Emailaddress = utils.MapGetString(g, "emailaddress")
	data.Hbcustominterval = utils.MapGetInt64(g, "hbcustominterval")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Mode = utils.MapGetString(g, "mode")
	data.Nodeid = utils.MapGetInt64(g, "nodeid")
	data.Port = utils.MapGetInt64(g, "port")
	data.Proxyauthservice = utils.MapGetString(g, "proxyauthservice")
	data.Proxymode = utils.MapGetString(g, "proxymode")

	// Read-only metadata.
	data.Sslcardfirstfailure = utils.MapGetString(g, "sslcardfirstfailure")
	data.Sslcardlatestfailure = utils.MapGetString(g, "sslcardlatestfailure")
	data.Powfirstfail = utils.MapGetString(g, "powfirstfail")
	data.Powlatestfailure = utils.MapGetString(g, "powlatestfailure")
	data.Hddfirstfail = utils.MapGetString(g, "hddfirstfail")
	data.Hddlatestfailure = utils.MapGetString(g, "hddlatestfailure")
	data.Flashfirstfail = utils.MapGetString(g, "flashfirstfail")
	data.Flashlatestfailure = utils.MapGetString(g, "flashlatestfailure")
	data.Rlfirsthighdrop = utils.MapGetString(g, "rlfirsthighdrop")
	data.Rllatesthighdrop = utils.MapGetString(g, "rllatesthighdrop")
	data.Restartlatestfail = utils.MapGetString(g, "restartlatestfail")
	data.Memthrefirstanomaly = utils.MapGetString(g, "memthrefirstanomaly")
	data.Memthrelatestanomaly = utils.MapGetString(g, "memthrelatestanomaly")
	data.Callhomestatus = utils.MapGetStringList(g, "callhomestatus")
	data.Anomalydetection = utils.MapGetString(g, "anomalydetection")
}
