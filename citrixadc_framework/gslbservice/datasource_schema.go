package gslbservice

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GslbserviceDataSourceModel is the data-source-specific model, decoupled from
// GslbserviceResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (gslb,
// svrstate, svreffgslbstate, health/state metadata, ...). Every non-key
// attribute is Computed.
type GslbserviceDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Appflowlog       types.String `tfsdk:"appflowlog"`
	Cip              types.String `tfsdk:"cip"`
	Cipheader        types.String `tfsdk:"cipheader"`
	Clttimeout       types.Int64  `tfsdk:"clttimeout"`
	Cnameentry       types.String `tfsdk:"cnameentry"`
	Comment          types.String `tfsdk:"comment"`
	Cookietimeout    types.Int64  `tfsdk:"cookietimeout"`
	Delay            types.Int64  `tfsdk:"delay"`
	Downstateflush   types.String `tfsdk:"downstateflush"`
	Hashid           types.Int64  `tfsdk:"hashid"`
	Healthmonitor    types.String `tfsdk:"healthmonitor"`
	Ip               types.String `tfsdk:"ip"`
	Ipaddress        types.String `tfsdk:"ipaddress"`
	Maxaaausers      types.Int64  `tfsdk:"maxaaausers"`
	Maxbandwidth     types.Int64  `tfsdk:"maxbandwidth"`
	Maxclient        types.Int64  `tfsdk:"maxclient"`
	Monitornamesvc   types.String `tfsdk:"monitornamesvc"`
	Monthreshold     types.Int64  `tfsdk:"monthreshold"`
	Naptrdomainttl   types.Int64  `tfsdk:"naptrdomainttl"`
	Naptrorder       types.Int64  `tfsdk:"naptrorder"`
	Naptrpreference  types.Int64  `tfsdk:"naptrpreference"`
	Naptrreplacement types.String `tfsdk:"naptrreplacement"`
	Naptrservices    types.String `tfsdk:"naptrservices"`
	Port             types.Int64  `tfsdk:"port"`
	Publicip         types.String `tfsdk:"publicip"`
	Publicport       types.Int64  `tfsdk:"publicport"`
	Servername       types.String `tfsdk:"servername"`
	Servicename      types.String `tfsdk:"servicename"` // Required lookup key
	Servicetype      types.String `tfsdk:"servicetype"`
	Sitename         types.String `tfsdk:"sitename"`
	Sitepersistence  types.String `tfsdk:"sitepersistence"`
	Siteprefix       types.String `tfsdk:"siteprefix"`
	State            types.String `tfsdk:"state"`
	Svrtimeout       types.Int64  `tfsdk:"svrtimeout"`
	Viewip           types.String `tfsdk:"viewip"`
	Viewname         types.String `tfsdk:"viewname"`
	Weight           types.Int64  `tfsdk:"weight"`
	Lbmonitorbinding types.Set    `tfsdk:"lbmonitorbinding"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/gslbservice.json). Never settable; populated from GET.
	Gslb                      types.String `tfsdk:"gslb"`
	Svrstate                  types.String `tfsdk:"svrstate"`
	Svreffgslbstate           types.String `tfsdk:"svreffgslbstate"`
	Gslbthreshold             types.Int64  `tfsdk:"gslbthreshold"`
	Gslbsvcstats              types.Int64  `tfsdk:"gslbsvcstats"`
	Monstate                  types.String `tfsdk:"monstate"`
	Preferredlocation         types.String `tfsdk:"preferredlocation"`
	MonitorState              types.String `tfsdk:"monitor_state"`
	Statechangetimesec        types.String `tfsdk:"statechangetimesec"`
	Tickssincelaststatechange types.Int64  `tfsdk:"tickssincelaststatechange"`
	Threshold                 types.String `tfsdk:"threshold"`
	Clmonowner                types.Int64  `tfsdk:"clmonowner"`
	Clmonview                 types.Int64  `tfsdk:"clmonview"`
	Gslbsvchealth             types.Int64  `tfsdk:"gslbsvchealth"`
	Glsbsvchealthdescr        types.String `tfsdk:"glsbsvchealthdescr"`
	Nodefaultbindings         types.String `tfsdk:"nodefaultbindings"`
}

func GslbserviceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging appflow flow information",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In the request that is forwarded to the GSLB service, insert a header that stores the client's IP address. Client IP header insertion is used in connection-proxy based site persistence.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the HTTP header that stores the client's IP address. Used with the Client IP option. If client IP header insertion is enabled on the service and a name is not specified for the header, the Citrix ADC uses the name specified by the cipHeader parameter in the set ns param command or, in the GUI, the Client IP Header parameter in the Configure HTTP Parameters dialog box.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle time, in seconds, after which a client connection is terminated. Applicable if connection proxy based site persistence is used.",
			},
			"cnameentry": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Canonical name of the GSLB service. Used in CNAME-based GSLB.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments that you might want to associate with the GSLB service.",
			},
			"cookietimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout value, in minutes, for the cookie, when cookie based site persistence is enabled.",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The time, in seconds, after which the GSLB service is disabled when disabling with -delay.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with the GSLB service when its state transitions from UP to DOWN. Do not enable this option for services that must complete their transactions. Applicable if connection proxy based site persistence is used.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique hash identifier for the GSLB service, used by hash based load balancing methods.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor the health of the GSLB service.",
			},
			"ip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address for the GSLB service. Should represent a load balancing, content switching, or VPN virtual server on the Citrix ADC, or the IP address of another load balancing device.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new IP address of the service.",
			},
			"maxaaausers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of SSL VPN users that can be logged on concurrently to the VPN virtual server that is represented by this GSLB service. A GSLB service whose user count reaches the maximum is not considered when a GSLB decision is made, until the count drops below the maximum.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the maximum bandwidth allowed for the service. A GSLB service whose bandwidth reaches the maximum is not considered when a GSLB decision is made, until its bandwidth consumption drops below the maximum.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The maximum number of open connections that the service can support at any given time. A GSLB service whose connection count reaches the maximum is not considered when a GSLB decision is made, until the connection count drops below the maximum.",
			},
			"monitornamesvc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor to bind to the service.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitoring threshold value for the GSLB service. If the sum of the weights of the monitors that are bound to this GSLB service and are in the UP state is not equal to or greater than this threshold value, the service is marked as DOWN.",
			},
			"naptrdomainttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Modify the TTL of the internally created naptr domain",
			},
			"naptrorder": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest",
			},
			"naptrpreference": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.",
			},
			"naptrreplacement": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The replacement domain name for this NAPTR.",
			},
			"naptrservices": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service Parameters applicable to this delegation path.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which the load balancing entity represented by this GSLB service listens.",
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
				Description: "Name of the server hosting the GSLB service.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the GSLB service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the GSLB service is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my gslbsvc\" or 'my gslbsvc').",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of service to create.",
			},
			"sitename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the GSLB site to which the service belongs.",
			},
			"sitepersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use cookie-based site persistence. Applicable only to HTTP and SSL GSLB services.",
			},
			"siteprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The site's prefix string. When the service is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound service-domain pair by concatenating the site prefix of the service and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the service.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle time, in seconds, after which a server connection is terminated. Applicable if connection proxy based site persistence is used.",
			},
			"viewip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address to be used for the given view",
			},
			"viewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. A larger number specifies a greater weight. Contributes to the monitoring threshold, which determines the state of the service.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"gslb": schema.StringAttribute{
				Computed:    true,
				Description: "GSLB service scope. Possible values: REMOTE, LOCAL.",
			},
			"svrstate": schema.StringAttribute{
				Computed:    true,
				Description: "Server state.",
			},
			"svreffgslbstate": schema.StringAttribute{
				Computed:    true,
				Description: "Effective state of the gslb svc.",
			},
			"gslbthreshold": schema.Int64Attribute{
				Computed:    true,
				Description: "Indicates if gslb svc has reached threshold.",
			},
			"gslbsvcstats": schema.Int64Attribute{
				Computed:    true,
				Description: "Indicates if gslb svc has stats of the primary or the whole chain.",
			},
			"monstate": schema.StringAttribute{
				Computed:    true,
				Description: "State of the monitor bound to gslb service. Possible values: ENABLED, DISABLED.",
			},
			"preferredlocation": schema.StringAttribute{
				Computed:    true,
				Description: "Prefered location.",
			},
			"monitor_state": schema.StringAttribute{
				Computed:    true,
				Description: "The running state of the monitor on this service.",
			},
			"statechangetimesec": schema.StringAttribute{
				Computed:    true,
				Description: "Time when last state change happened. Seconds part.",
			},
			"tickssincelaststatechange": schema.Int64Attribute{
				Computed:    true,
				Description: "Time in 10 millisecond ticks since the last state change.",
			},
			"threshold": schema.StringAttribute{
				Computed:    true,
				Description: "Threshold state. Possible values: ABOVE, BELOW.",
			},
			"clmonowner": schema.Int64Attribute{
				Computed:    true,
				Description: "Tells the mon owner of the gslb service.",
			},
			"clmonview": schema.Int64Attribute{
				Computed:    true,
				Description: "Tells the view id of the monitoring owner.",
			},
			"gslbsvchealth": schema.Int64Attribute{
				Computed:    true,
				Description: "This parameter displays effective health of a GSLB service.",
			},
			"glsbsvchealthdescr": schema.StringAttribute{
				Computed:    true,
				Description: "Displays the warning message related to health a of GSLB service.",
			},
			"nodefaultbindings": schema.StringAttribute{
				Computed:    true,
				Description: "To determine if the configuration will have default ssl CIPHER and ECC curve bindings. Possible values: YES, NO.",
			},
		},
		Blocks: map[string]schema.Block{
			// The nested attributes are defined in gslbserviceLbmonitorbindingDSAttrs
			// so this schema stays a clean top-level attribute/block map.
			"lbmonitorbinding": schema.SetNestedBlock{
				Description: "Monitors bound to the GSLB service.",
				NestedObject: schema.NestedBlockObject{
					Attributes: gslbserviceLbmonitorbindingDSAttrs(),
				},
			},
		},
	}
}

// gslbserviceDataSourceSetAttrFromGet projects a NITRO gslbservice GET response
// onto the data-source model. Attributes are filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers. The
// lbmonitorbinding block is left as read from config (the gslbservice GET does
// not return it), preserving the prior data-source behavior.
func gslbserviceDataSourceSetAttrFromGet(ctx context.Context, data *GslbserviceDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In gslbserviceDataSourceSetAttrFromGet Function")

	if v, ok := g["servicename"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Servicename = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Appflowlog = utils.MapGetString(g, "appflowlog")
	data.Cip = utils.MapGetString(g, "cip")
	data.Cipheader = utils.MapGetString(g, "cipheader")
	data.Clttimeout = utils.MapGetInt64(g, "clttimeout")
	data.Cnameentry = utils.MapGetString(g, "cnameentry")
	data.Comment = utils.MapGetString(g, "comment")
	data.Cookietimeout = utils.MapGetInt64(g, "cookietimeout")
	// delay is config-only; NITRO never returns it -> Null.
	data.Delay = utils.MapGetInt64(g, "delay")
	data.Downstateflush = utils.MapGetString(g, "downstateflush")
	data.Hashid = utils.MapGetInt64(g, "hashid")
	data.Healthmonitor = utils.MapGetString(g, "healthmonitor")
	// ip is not returned by NITRO; it is mapped from ipaddress.
	data.Ip = utils.MapGetString(g, "ipaddress")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Maxaaausers = utils.MapGetInt64(g, "maxaaausers")
	data.Maxbandwidth = utils.MapGetInt64(g, "maxbandwidth")
	data.Maxclient = utils.MapGetInt64(g, "maxclient")
	data.Monitornamesvc = utils.MapGetString(g, "monitornamesvc")
	data.Monthreshold = utils.MapGetInt64(g, "monthreshold")
	data.Naptrdomainttl = utils.MapGetInt64(g, "naptrdomainttl")
	data.Naptrorder = utils.MapGetInt64(g, "naptrorder")
	data.Naptrpreference = utils.MapGetInt64(g, "naptrpreference")
	data.Naptrreplacement = utils.MapGetString(g, "naptrreplacement")
	data.Naptrservices = utils.MapGetString(g, "naptrservices")
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
	data.Viewip = utils.MapGetString(g, "viewip")
	data.Viewname = utils.MapGetString(g, "viewname")
	data.Weight = utils.MapGetInt64(g, "weight")

	// Read-only attributes.
	data.Gslb = utils.MapGetString(g, "gslb")
	data.Svrstate = utils.MapGetString(g, "svrstate")
	data.Svreffgslbstate = utils.MapGetString(g, "svreffgslbstate")
	data.Gslbthreshold = utils.MapGetInt64(g, "gslbthreshold")
	data.Gslbsvcstats = utils.MapGetInt64(g, "gslbsvcstats")
	data.Monstate = utils.MapGetString(g, "monstate")
	data.Preferredlocation = utils.MapGetString(g, "preferredlocation")
	data.MonitorState = utils.MapGetString(g, "monitor_state")
	data.Statechangetimesec = utils.MapGetString(g, "statechangetimesec")
	data.Tickssincelaststatechange = utils.MapGetInt64(g, "tickssincelaststatechange")
	data.Threshold = utils.MapGetString(g, "threshold")
	data.Clmonowner = utils.MapGetInt64(g, "clmonowner")
	data.Clmonview = utils.MapGetInt64(g, "clmonview")
	data.Gslbsvchealth = utils.MapGetInt64(g, "gslbsvchealth")
	data.Glsbsvchealthdescr = utils.MapGetString(g, "glsbsvchealthdescr")
	data.Nodefaultbindings = utils.MapGetString(g, "nodefaultbindings")

	// lbmonitorbinding block: left as read from config (GET does not return it).
}
