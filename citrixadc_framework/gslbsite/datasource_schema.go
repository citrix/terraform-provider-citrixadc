package gslbsite

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GslbsiteDataSourceModel is the data-source-specific model, decoupled from
// GslbsiteResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only NITRO attributes the resource deliberately omits
// (status, sitestate, persistencemepstatus, ...). Every non-key attribute is
// Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares,
// which is why it cannot reuse the resource model.
type GslbsiteDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Backupparentlist       types.List   `tfsdk:"backupparentlist"`
	Clip                   types.String `tfsdk:"clip"`
	Krpcnodesrcip          types.String `tfsdk:"krpcnodesrcip"`
	Metricexchange         types.String `tfsdk:"metricexchange"`
	Naptrreplacementsuffix types.String `tfsdk:"naptrreplacementsuffix"`
	Newname                types.String `tfsdk:"newname"`
	Nwmetricexchange       types.String `tfsdk:"nwmetricexchange"`
	Parentsite             types.String `tfsdk:"parentsite"`
	Publicclip             types.String `tfsdk:"publicclip"`
	Publicip               types.String `tfsdk:"publicip"`
	Sessionexchange        types.String `tfsdk:"sessionexchange"`
	Siteipaddress          types.String `tfsdk:"siteipaddress"`
	Sitename               types.String `tfsdk:"sitename"` // Required lookup key
	Sitepassword           types.String `tfsdk:"sitepassword"`
	SitepasswordWo         types.String `tfsdk:"sitepassword_wo"`
	SitepasswordWoVersion  types.Int64  `tfsdk:"sitepassword_wo_version"`
	Sitetype               types.String `tfsdk:"sitetype"`
	Triggermonitor         types.String `tfsdk:"triggermonitor"`

	// Read-only (GET-only) NITRO attributes from the read-only set
	// (zion73x_readonly/gslbsite.json). Never settable; populated from GET.
	Status               types.String `tfsdk:"status"`
	Persistencemepstatus types.String `tfsdk:"persistencemepstatus"`
	Version              types.Int64  `tfsdk:"version"`
	Curbackupparentip    types.String `tfsdk:"curbackupparentip"`
	Sitestate            types.String `tfsdk:"sitestate"`
	Oldname              types.String `tfsdk:"oldname"`
}

func GslbsiteDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"backupparentlist": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The list of backup gslb sites configured in preferred order. Need to be parent gsb sites.",
			},
			"clip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cluster IP address. Specify this parameter to connect to the remote cluster site for GSLB auto-sync. Note: The cluster IP address is defined when creating the cluster.",
			},
			"krpcnodesrcip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IP address to be used to communicate with this GSLB site. Minimum length =  1",
			},
			"metricexchange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exchange metrics with other sites. Metrics are exchanged by using Metric Exchange Protocol (MEP). The appliances in the GSLB setup exchange health information once every second.\n\nIf you disable metrics exchange, you can use only static load balancing methods (such as round robin, static proximity, or the hash-based methods), and if you disable metrics exchange when a dynamic load balancing method (such as least connection) is in operation, the appliance falls back to round robin. Also, if you disable metrics exchange, you must use a monitor to determine the state of GSLB services. Otherwise, the service is marked as DOWN.",
			},
			"naptrreplacementsuffix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The naptr replacement suffix configured here will be used to construct the naptr replacement field in NAPTR record.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the GSLB site.",
			},
			"nwmetricexchange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exchange, with other GSLB sites, network metrics such as round-trip time (RTT), learned from communications with various local DNS (LDNS) servers used by clients. RTT information is used in the dynamic RTT load balancing method, and is exchanged every 5 seconds.",
			},
			"parentsite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parent site of the GSLB site, in a parent-child topology.",
			},
			"publicclip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address to be used to globally access the remote cluster when it is deployed behind a NAT. It can be same as the normal cluster IP address.",
			},
			"publicip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Public IP address for the local site. Required only if the appliance is deployed in a private address space and the site has a public IP address hosted on an external firewall or a NAT device.",
			},
			"sessionexchange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exchange persistent session entries with other GSLB sites every five seconds.",
			},
			"siteipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address for the GSLB site. The GSLB site uses this IP address to communicate with other GSLB sites. For a local site, use any IP address that is owned by the appliance (for example, a SNIP or MIP address, or the IP address of the ADNS service).",
			},
			"sitename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the GSLB site. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my gslbsite\" or 'my gslbsite').",
			},
			"sitepassword": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password to be used for mep communication between gslb site nodes.",
			},
			"sitepassword_wo": schema.StringAttribute{
				Optional:    true,
				Description: "Password to be used for mep communication between gslb site nodes.",
			},
			"sitepassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a sitepassword_wo update.",
			},
			"sitetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of site to create. If the type is not specified, the appliance automatically detects and sets the type on the basis of the IP address being assigned to the site. If the specified site IP address is owned by the appliance (for example, a MIP address or SNIP address), the site is a local site. Otherwise, it is a remote site.",
			},
			"triggermonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the conditions under which the GSLB service must be monitored by a monitor, if one is bound. Available settings function as follows:\n* ALWAYS - Monitor the GSLB service at all times.\n* MEPDOWN - Monitor the GSLB service only when the exchange of metrics through the Metrics Exchange Protocol (MEP) is disabled.\nMEPDOWN_SVCDOWN - Monitor the service in either of the following situations:\n* The exchange of metrics through MEP is disabled.\n* The exchange of metrics through MEP is enabled but the status of the service, learned through metrics exchange, is DOWN.",
			},

			// Read-only (GET-only) NITRO attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "Current metric exchange status (ACTIVE, INACTIVE, DOWN).",
			},
			"persistencemepstatus": schema.StringAttribute{
				Computed:    true,
				Description: "Network metric and persistence exchange MEP connection status (ACTIVE, INACTIVE, DOWN).",
			},
			"version": schema.Int64Attribute{
				Computed:    true,
				Description: "Will be true if the remote site's version is ncore compatible with the local site (>= 9.2).",
			},
			"curbackupparentip": schema.StringAttribute{
				Computed:    true,
				Description: "Current active backup parent IP address since the configured is DOWN.",
			},
			"sitestate": schema.StringAttribute{
				Computed:    true,
				Description: "Site state (for example UP, DOWN, OUT OF SERVICE, DISABLED).",
			},
			"oldname": schema.StringAttribute{
				Computed:    true,
				Description: "Old name for the GSLB site.",
			},
		},
	}
}

// gslbsiteDataSourceSetAttrFromGet projects a NITRO gslbsite GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them) — no unknown->null resolution or plan preservation is required. The
// shared utils.MapGet* helpers implement that projection.
func gslbsiteDataSourceSetAttrFromGet(ctx context.Context, data *GslbsiteDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In gslbsiteDataSourceSetAttrFromGet Function")

	if v, ok := g["sitename"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Sitename = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Backupparentlist = utils.MapGetStringList(g, "backupparentlist")
	data.Clip = utils.MapGetString(g, "clip")
	data.Krpcnodesrcip = utils.MapGetString(g, "krpcnodesrcip")
	data.Metricexchange = utils.MapGetString(g, "metricexchange")
	data.Naptrreplacementsuffix = utils.MapGetString(g, "naptrreplacementsuffix")
	data.Newname = utils.MapGetString(g, "newname")
	data.Nwmetricexchange = utils.MapGetString(g, "nwmetricexchange")
	data.Parentsite = utils.MapGetString(g, "parentsite")
	data.Publicclip = utils.MapGetString(g, "publicclip")
	data.Publicip = utils.MapGetString(g, "publicip")
	data.Sessionexchange = utils.MapGetString(g, "sessionexchange")
	data.Siteipaddress = utils.MapGetString(g, "siteipaddress")
	data.Sitetype = utils.MapGetString(g, "sitetype")
	data.Triggermonitor = utils.MapGetString(g, "triggermonitor")

	// sitepassword / sitepassword_wo(+version) are secret/write-only inputs the
	// GET never returns -> Null.
	data.Sitepassword = types.StringNull()
	data.SitepasswordWo = types.StringNull()
	data.SitepasswordWoVersion = types.Int64Null()

	// Read-only NITRO attributes.
	data.Status = utils.MapGetString(g, "status")
	data.Persistencemepstatus = utils.MapGetString(g, "persistencemepstatus")
	data.Version = utils.MapGetInt64(g, "version")
	data.Curbackupparentip = utils.MapGetString(g, "curbackupparentip")
	data.Sitestate = utils.MapGetString(g, "sitestate")
	data.Oldname = utils.MapGetString(g, "oldname")
}
