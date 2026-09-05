package gslbservicegroup

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GslbservicegroupDataSourceModel is the data-source-specific model, decoupled
// from GslbservicegroupResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only NITRO attributes the resource deliberately omits
// (svrstate, numofconnections, monitor stats, effective-state, ...). Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type GslbservicegroupDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Appflowlog       types.String `tfsdk:"appflowlog"`
	Autodelayedtrofs types.String `tfsdk:"autodelayedtrofs"`
	Autoscale        types.String `tfsdk:"autoscale"`
	Cip              types.String `tfsdk:"cip"`
	Cipheader        types.String `tfsdk:"cipheader"`
	Clttimeout       types.Int64  `tfsdk:"clttimeout"`
	Comment          types.String `tfsdk:"comment"`
	Delay            types.Int64  `tfsdk:"delay"`
	Downstateflush   types.String `tfsdk:"downstateflush"`
	DupWeight        types.Int64  `tfsdk:"dup_weight"`
	Graceful         types.String `tfsdk:"graceful"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	Healthmonitor    types.String `tfsdk:"healthmonitor"`
	Includemembers   types.Bool   `tfsdk:"includemembers"`
	Maxbandwidth     types.Int64  `tfsdk:"maxbandwidth"`
	Maxclient        types.Int64  `tfsdk:"maxclient"`
	MonitorNameSvc   types.String `tfsdk:"monitor_name_svc"`
	Monthreshold     types.Int64  `tfsdk:"monthreshold"`
	Newname          types.String `tfsdk:"newname"`
	Order            types.Int64  `tfsdk:"order"`
	Port             types.Int64  `tfsdk:"port"`
	Publicip         types.String `tfsdk:"publicip"`
	Publicport       types.Int64  `tfsdk:"publicport"`
	Servername       types.String `tfsdk:"servername"`
	Servicegroupname types.String `tfsdk:"servicegroupname"` // Required lookup key
	Servicetype      types.String `tfsdk:"servicetype"`
	Sitename         types.String `tfsdk:"sitename"`
	Sitepersistence  types.String `tfsdk:"sitepersistence"`
	Siteprefix       types.String `tfsdk:"siteprefix"`
	State            types.String `tfsdk:"state"`
	Svrtimeout       types.Int64  `tfsdk:"svrtimeout"`
	Weight           types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) NITRO attributes from the read-only set
	// (zion73x_readonly/gslbservicegroup.json). Never settable; populated from GET.
	Numofconnections           types.Int64  `tfsdk:"numofconnections"`
	Serviceconftype            types.Bool   `tfsdk:"serviceconftype"`
	Value                      types.String `tfsdk:"value"`
	Svrstate                   types.String `tfsdk:"svrstate"`
	Ip                         types.String `tfsdk:"ip"`
	Monstatcode                types.Int64  `tfsdk:"monstatcode"`
	Monstatparam1              types.Int64  `tfsdk:"monstatparam1"`
	Monstatparam2              types.Int64  `tfsdk:"monstatparam2"`
	Monstatparam3              types.Int64  `tfsdk:"monstatparam3"`
	Statechangetimemsec        types.Int64  `tfsdk:"statechangetimemsec"`
	Stateupdatereason          types.Int64  `tfsdk:"stateupdatereason"`
	Clmonowner                 types.Int64  `tfsdk:"clmonowner"`
	Clmonview                  types.Int64  `tfsdk:"clmonview"`
	Groupcount                 types.Int64  `tfsdk:"groupcount"`
	Serviceipstr               types.String `tfsdk:"serviceipstr"`
	Servicegroupeffectivestate types.String `tfsdk:"servicegroupeffectivestate"`
	Gslb                       types.String `tfsdk:"gslb"`
	Svreffgslbstate            types.String `tfsdk:"svreffgslbstate"`
	Nodefaultbindings          types.String `tfsdk:"nodefaultbindings"`
}

func GslbservicegroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging of AppFlow information for the specified GSLB service group.",
			},
			"autodelayedtrofs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates graceful movement of the service to TROFS. System will wait for monitor response time out before moving to TROFS",
			},
			"autoscale": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Auto scale option for a GSLB servicegroup",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the Client IP header in requests forwarded to the GSLB service.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If client IP insertion is enabled, and the client IP header is not specified, the value of Client IP Header parameter or the value set by the set ns config command is used as client's IP header name.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle client connection.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the GSLB service group.",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with all the services in the GSLB service group whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"dup_weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "weight of the monitor that is bound to GSLB servicegroup.",
			},
			"graceful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Wait for all existing connections to the service to terminate before shutting down the service.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor the health of this GSLB service.Available settings function are as follows:\nYES - Send probes to check the health of the GSLB service.\nNO - Do not send probes to check the health of the GSLB service. With the NO option, the appliance shows the service as UP at all times.",
			},
			"includemembers": schema.BoolAttribute{
				Optional:    true,
				Description: "Display the members of the listed GSLB service groups in addition to their settings. Can be specified when no service group name is provided in the command. In that case, the details displayed for each service group are identical to the details displayed when a service group name is provided, except that bound monitors are not displayed.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, allocated for all the services in the GSLB service group.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of simultaneous open connections for the GSLB service group.",
			},
			"monitor_name_svc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor bound to the GSLB service group. Used to assign a weight to the monitor.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum sum of weights of the monitors that are bound to this GSLB service. Used to determine whether to mark a GSLB service as UP or DOWN.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the GSLB service group.",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the gslb servicegroup member",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Server port number.",
			},
			"publicip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.",
			},
			"publicport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the server to which to bind the service group.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB service group. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the name is created.",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used to exchange data with the GSLB service.",
			},
			"sitename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the GSLB site to which the service group belongs.",
			},
			"sitepersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use cookie-based site persistence. Applicable only to HTTP and SSL non-autoscale enabled GSLB servicegroups.",
			},
			"siteprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the GSLB service group.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle server connection.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.",
			},

			// Read-only (GET-only) NITRO attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"numofconnections": schema.Int64Attribute{
				Computed:    true,
				Description: "This will tell the number of client side connections are still open.",
			},
			"serviceconftype": schema.BoolAttribute{
				Computed:    true,
				Description: "The configuration type of the GSLB service group.",
			},
			"value": schema.StringAttribute{
				Computed:    true,
				Description: "SSL Status. Possible values: Certkey/Certkeybundle/Vault not bound/Cert-store not usable, SSL feature disabled.",
			},
			"svrstate": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the GSLB service (for example UP, DOWN, OUT OF SERVICE, DISABLED).",
			},
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "IP Address.",
			},
			"monstatcode": schema.Int64Attribute{
				Computed:    true,
				Description: "The code indicating the monitor response.",
			},
			"monstatparam1": schema.Int64Attribute{
				Computed:    true,
				Description: "First parameter for use with message code.",
			},
			"monstatparam2": schema.Int64Attribute{
				Computed:    true,
				Description: "Second parameter for use with message code.",
			},
			"monstatparam3": schema.Int64Attribute{
				Computed:    true,
				Description: "Third parameter for use with message code.",
			},
			"statechangetimemsec": schema.Int64Attribute{
				Computed:    true,
				Description: "Time when last state change occurred. Milliseconds part.",
			},
			"stateupdatereason": schema.Int64Attribute{
				Computed:    true,
				Description: "Checks state update reason on the secondary node.",
			},
			"clmonowner": schema.Int64Attribute{
				Computed:    true,
				Description: "Tells the mon owner of the service.",
			},
			"clmonview": schema.Int64Attribute{
				Computed:    true,
				Description: "Tells the view id of the monitoring owner.",
			},
			"groupcount": schema.Int64Attribute{
				Computed:    true,
				Description: "Servicegroup Count.",
			},
			"serviceipstr": schema.StringAttribute{
				Computed:    true,
				Description: "This field has been introduced to show the dbs services ip.",
			},
			"servicegroupeffectivestate": schema.StringAttribute{
				Computed:    true,
				Description: "Indicates the effective GSLB servicegroup state based on the state of the bound service items (UP, DOWN, OUT OF SERVICE, PARTIAL-UP, PARTIAL-DOWN).",
			},
			"gslb": schema.StringAttribute{
				Computed:    true,
				Description: "GSLB service scope (REMOTE, LOCAL).",
			},
			"svreffgslbstate": schema.StringAttribute{
				Computed:    true,
				Description: "Effective state of the gslb svc.",
			},
			"nodefaultbindings": schema.StringAttribute{
				Computed:    true,
				Description: "To determine if the configuration is from stylebooks (YES, NO).",
			},
		},
	}
}

// gslbservicegroupDataSourceSetAttrFromGet projects a NITRO gslbservicegroup GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) — no unknown->null resolution or plan preservation is
// required. The shared utils.MapGet* helpers implement that projection.
func gslbservicegroupDataSourceSetAttrFromGet(ctx context.Context, data *GslbservicegroupDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In gslbservicegroupDataSourceSetAttrFromGet Function")

	if v, ok := g["servicegroupname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Servicegroupname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Appflowlog = utils.MapGetString(g, "appflowlog")
	data.Autodelayedtrofs = utils.MapGetString(g, "autodelayedtrofs")
	data.Autoscale = utils.MapGetString(g, "autoscale")
	data.Cip = utils.MapGetString(g, "cip")
	data.Cipheader = utils.MapGetString(g, "cipheader")
	data.Clttimeout = utils.MapGetInt64(g, "clttimeout")
	data.Comment = utils.MapGetString(g, "comment")
	data.Delay = utils.MapGetInt64(g, "delay")
	data.Downstateflush = utils.MapGetString(g, "downstateflush")
	data.DupWeight = utils.MapGetInt64(g, "dup_weight")
	data.Graceful = utils.MapGetString(g, "graceful")
	data.Hashid = utils.MapGetInt64(g, "hashid")
	data.Healthmonitor = utils.MapGetString(g, "healthmonitor")
	data.Includemembers = utils.MapGetBool(g, "includemembers")
	data.Maxbandwidth = utils.MapGetInt64(g, "maxbandwidth")
	data.Maxclient = utils.MapGetInt64(g, "maxclient")
	data.MonitorNameSvc = utils.MapGetString(g, "monitor_name_svc")
	data.Monthreshold = utils.MapGetInt64(g, "monthreshold")
	data.Newname = utils.MapGetString(g, "newname")
	data.Order = utils.MapGetInt64(g, "order")
	data.Port = utils.MapGetInt64(g, "port")
	data.Publicip = utils.MapGetString(g, "publicip")
	data.Publicport = utils.MapGetInt64(g, "publicport")
	data.Servername = utils.MapGetString(g, "servername")
	data.Servicetype = utils.MapGetString(g, "servicetype")
	data.Sitename = utils.MapGetString(g, "sitename")
	data.Sitepersistence = utils.MapGetString(g, "sitepersistence")
	data.Siteprefix = utils.MapGetString(g, "siteprefix")
	data.State = utils.MapGetString(g, "state")
	data.Svrtimeout = utils.MapGetInt64(g, "svrtimeout")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Read-only NITRO attributes.
	data.Numofconnections = utils.MapGetInt64(g, "numofconnections")
	data.Serviceconftype = utils.MapGetBool(g, "serviceconftype")
	data.Value = utils.MapGetString(g, "value")
	data.Svrstate = utils.MapGetString(g, "svrstate")
	data.Ip = utils.MapGetString(g, "ip")
	data.Monstatcode = utils.MapGetInt64(g, "monstatcode")
	data.Monstatparam1 = utils.MapGetInt64(g, "monstatparam1")
	data.Monstatparam2 = utils.MapGetInt64(g, "monstatparam2")
	data.Monstatparam3 = utils.MapGetInt64(g, "monstatparam3")
	data.Statechangetimemsec = utils.MapGetInt64(g, "statechangetimemsec")
	data.Stateupdatereason = utils.MapGetInt64(g, "stateupdatereason")
	data.Clmonowner = utils.MapGetInt64(g, "clmonowner")
	data.Clmonview = utils.MapGetInt64(g, "clmonview")
	data.Groupcount = utils.MapGetInt64(g, "groupcount")
	data.Serviceipstr = utils.MapGetString(g, "serviceipstr")
	data.Servicegroupeffectivestate = utils.MapGetString(g, "servicegroupeffectivestate")
	data.Gslb = utils.MapGetString(g, "gslb")
	data.Svreffgslbstate = utils.MapGetString(g, "svreffgslbstate")
	data.Nodefaultbindings = utils.MapGetString(g, "nodefaultbindings")
}
