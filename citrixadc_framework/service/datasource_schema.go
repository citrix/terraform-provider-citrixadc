package service

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ServiceDataSourceModel is the data-source-specific model, decoupled from
// ServiceResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only operational/status attributes that the resource
// deliberately omits (svrstate, monitor_state, numofconnections, ...). The
// Framework's per-attribute model <-> schema reflection requires this model to
// have exactly the attributes the data-source schema declares, which is why it
// cannot reuse the resource model.
type ServiceDataSourceModel struct {
	Id                           types.String `tfsdk:"id"`
	Internal                     types.Bool   `tfsdk:"internal"`
	Accessdown                   types.String `tfsdk:"accessdown"`
	Aigwprofilename              types.String `tfsdk:"aigwprofilename"`
	All                          types.Bool   `tfsdk:"all"`
	Appflowlog                   types.String `tfsdk:"appflowlog"`
	Cacheable                    types.String `tfsdk:"cacheable"`
	Cachetype                    types.String `tfsdk:"cachetype"`
	Cip                          types.String `tfsdk:"cip"`
	Cipheader                    types.String `tfsdk:"cipheader"`
	Cka                          types.String `tfsdk:"cka"`
	Cleartextport                types.Int64  `tfsdk:"cleartextport"`
	Clttimeout                   types.Int64  `tfsdk:"clttimeout"`
	Cmp                          types.String `tfsdk:"cmp"`
	Comment                      types.String `tfsdk:"comment"`
	Contentinspectionprofilename types.String `tfsdk:"contentinspectionprofilename"`
	Customserverid               types.String `tfsdk:"customserverid"`
	Delay                        types.Int64  `tfsdk:"delay"`
	Dnsprofilename               types.String `tfsdk:"dnsprofilename"`
	Downstateflush               types.String `tfsdk:"downstateflush"`
	Graceful                     types.String `tfsdk:"graceful"`
	Hashid                       types.Int64  `tfsdk:"hashid"`
	Healthmonitor                types.String `tfsdk:"healthmonitor"`
	Httpprofilename              types.String `tfsdk:"httpprofilename"`
	Ip                           types.String `tfsdk:"ip"`
	Ipaddress                    types.String `tfsdk:"ipaddress"`
	Maxbandwidth                 types.Int64  `tfsdk:"maxbandwidth"`
	Maxclient                    types.Int64  `tfsdk:"maxclient"`
	Maxreq                       types.Int64  `tfsdk:"maxreq"`
	Mcpprofilename               types.String `tfsdk:"mcpprofilename"`
	Monconnectionclose           types.String `tfsdk:"monconnectionclose"`
	Monitornamesvc               types.String `tfsdk:"monitornamesvc"`
	Monthreshold                 types.Int64  `tfsdk:"monthreshold"`
	Name                         types.String `tfsdk:"name"`
	Netprofile                   types.String `tfsdk:"netprofile"`
	Riseapbrstatsmsgcode         types.Int64  `tfsdk:"riseapbrstatsmsgcode"`
	Pathmonitor                  types.String `tfsdk:"pathmonitor"`
	Pathmonitorindv              types.String `tfsdk:"pathmonitorindv"`
	Port                         types.Int64  `tfsdk:"port"`
	Processlocal                 types.String `tfsdk:"processlocal"`
	Quicprofilename              types.String `tfsdk:"quicprofilename"`
	Rtspsessionidremap           types.String `tfsdk:"rtspsessionidremap"`
	Serverid                     types.Int64  `tfsdk:"serverid"`
	Servername                   types.String `tfsdk:"servername"`
	Servicetype                  types.String `tfsdk:"servicetype"`
	Sp                           types.String `tfsdk:"sp"`
	State                        types.String `tfsdk:"state"`
	Svrtimeout                   types.Int64  `tfsdk:"svrtimeout"`
	Tcpb                         types.String `tfsdk:"tcpb"`
	Tcpprofilename               types.String `tfsdk:"tcpprofilename"`
	Td                           types.Int64  `tfsdk:"td"`
	Useproxyport                 types.String `tfsdk:"useproxyport"`
	Usip                         types.String `tfsdk:"usip"`
	Wasmmodule                   types.String `tfsdk:"wasmmodule"`
	Weight                       types.Int64  `tfsdk:"weight"`

	// Convenience blocks (shared with the resource model).
	Lbvserver            types.String `tfsdk:"lbvserver"`
	Lbmonitor            types.String `tfsdk:"lbmonitor"`
	Snienable            types.String `tfsdk:"snienable"`
	Commonname           types.String `tfsdk:"commonname"`
	WaitUntilDisabled    types.Bool   `tfsdk:"wait_until_disabled"`
	DisabledTimeout      types.String `tfsdk:"disabled_timeout"`
	DisabledPollDelay    types.String `tfsdk:"disabled_poll_delay"`
	DisabledPollInterval types.String `tfsdk:"disabled_poll_interval"`

	// Read-only (GET-only) operational/status attributes from the NITRO doc
	// read-only set (zion73x_readonly/service.json). Never settable; populated
	// from GET, and null when the appliance omits them.
	Numofconnections          types.Int64  `tfsdk:"numofconnections"`
	Policyname                types.String `tfsdk:"policyname"`
	Serviceconftype           types.Bool   `tfsdk:"serviceconftype"`
	Serviceconftype2          types.String `tfsdk:"serviceconftype2"`
	Value                     types.String `tfsdk:"value"`
	Gslb                      types.String `tfsdk:"gslb"`
	DupState                  types.String `tfsdk:"dup_state"`
	Publicip                  types.String `tfsdk:"publicip"`
	Publicport                types.Int64  `tfsdk:"publicport"`
	Svrstate                  types.String `tfsdk:"svrstate"`
	MonitorState              types.String `tfsdk:"monitor_state"`
	Monstatcode               types.Int64  `tfsdk:"monstatcode"`
	Lastresponse              types.String `tfsdk:"lastresponse"`
	Responsetime              types.Int64  `tfsdk:"responsetime"`
	Monstatparam1             types.Int64  `tfsdk:"monstatparam1"`
	Monstatparam2             types.Int64  `tfsdk:"monstatparam2"`
	Monstatparam3             types.Int64  `tfsdk:"monstatparam3"`
	Statechangetimesec        types.String `tfsdk:"statechangetimesec"`
	Statechangetimemsec       types.Int64  `tfsdk:"statechangetimemsec"`
	Tickssincelaststatechange types.Int64  `tfsdk:"tickssincelaststatechange"`
	Stateupdatereason         types.Int64  `tfsdk:"stateupdatereason"`
	Clmonowner                types.Int64  `tfsdk:"clmonowner"`
	Clmonview                 types.Int64  `tfsdk:"clmonview"`
	Serviceipstr              types.String `tfsdk:"serviceipstr"`
	Oracleserverversion       types.String `tfsdk:"oracleserverversion"`
	Nodefaultbindings         types.String `tfsdk:"nodefaultbindings"`
	Monuserstatusmesg         types.String `tfsdk:"monuserstatusmesg"`
	Builtin                   types.List   `tfsdk:"builtin"`
	Feature                   types.String `tfsdk:"feature"`
}

func ServiceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"internal": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display only dynamically learned services.",
			},
			"accessdown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use Layer 2 mode to bridge the packets sent to this service if it is marked as DOWN. If the service is DOWN, and this parameter is disabled, the packets are dropped.",
			},
			"aigwprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the AIGW Profile that contains AIGW Endpoint setting for the service.",
			},
			"all": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Display both user-configured and dynamically learned services.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging of AppFlow information.",
			},
			"cacheable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the transparent cache redirection virtual server to forward requests to the cache server.\nNote: Do not specify this parameter if you set the Cache Type parameter.",
			},
			"cachetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache type supported by the cache server.",
			},
			"cip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Before forwarding a request to the service, insert an HTTP header with the client's IPv4 or IPv6 address as its value. Used if the server needs the client's IP address for security, accounting, or other purposes, and setting the Use Source IP parameter is not a viable option.",
			},
			"cipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the HTTP header whose value must be set to the IP address of the client. Used with the Client IP parameter. If you set the Client IP parameter, and you do not specify a name for the header, the appliance uses the header name specified for the global Client IP Header parameter (the cipHeader parameter in the set ns param CLI command or the Client IP Header parameter in the Configure HTTP Parameters dialog box at System > Settings > Change HTTP parameters). If the global Client IP Header parameter is not specified, the appliance inserts a header with the name \"client-ip.\"",
			},
			"cka": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable client keep-alive for the service.",
			},
			"cleartextport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port to which clear text data must be sent after the appliance decrypts incoming SSL traffic. Applicable to transparent SSL services.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle client connection.",
			},
			"cmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable compression for the service.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the service.",
			},
			"contentinspectionprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ContentInspection profile that contains IPS/IDS communication related setting for the service",
			},
			"customserverid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique identifier for the service. Used when the persistency type for the virtual server is set to Custom Server ID.",
			},
			"delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, allocated to the NetScaler for a graceful shutdown of the service. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the service. DNS profile properties will applied to the transactions processed by a service. This parameter is valid only for ADNS, ADNS-TCP and ADNS-DOT services.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with a service whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"graceful": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Shut down gracefully, not accepting any new connections, and disabling the service when all of its connections are closed.",
			},
			"hashid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "A numerical identifier that can be used by hash based load balancing methods. Must be unique for each service.",
			},
			"healthmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Monitor the health of this service. Available settings function as follows:\nYES - Send probes to check the health of the service.\nNO - Do not send probes to check the health of the service. With the NO option, the appliance shows the service as UP at all times.",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP profile that contains HTTP configuration settings for the service.",
			},
			"ip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP to assign to the service.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new IP address of the service.",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, allocated to the service.",
			},
			"maxclient": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of simultaneous open connections to the service.",
			},
			"maxreq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of requests that can be sent on a persistent connection to the service.\nNote: Connection requests beyond this value are rejected.",
			},
			"mcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of MCP profile which will be attached to the service.",
			},
			"monconnectionclose": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Close monitoring connections by sending the service a connection termination message with the specified bit set.",
			},
			"monitornamesvc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the monitor bound to the specified service.",
			},
			"monthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum sum of weights of the monitors that are bound to this service. Used to determine whether to mark a service as UP or DOWN.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the service. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the service has been created.",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Network profile to use for the service.",
			},
			"riseapbrstatsmsgcode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The code indicating the rise apbr status.",
			},
			"pathmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path monitoring for clustering",
			},
			"pathmonitorindv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Individual Path monitoring decisions",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number of the service.",
			},
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By turning on this option packets destined to a service in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution.",
			},
			"quicprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of QUIC profile which will be attached to the service.",
			},
			"rtspsessionidremap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable RTSP session ID mapping for the service.",
			},
			"serverid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The  identifier for the service. This is used when the persistency type is set to Custom Server ID.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the server that hosts the service.",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol in which data is exchanged with the service.",
			},
			"sp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable surge protection for the service.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the service.",
			},
			"svrtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which to terminate an idle server connection.",
			},
			"tcpb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable TCP buffering for the service.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile that contains TCP configuration settings for the service.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the proxy port as the source port when initiating connections with the server. With the NO setting, the client-side connection port is used as the source port for the server-side connection.\nNote: This parameter is available only when the Use Source IP (USIP) parameter is set to YES.",
			},
			"usip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the client's IP address as the source IP address when initiating a connection to the server. When creating a service, if you do not set this parameter, the service inherits the global Use Source IP setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the service.",
			},
			"wasmmodule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the WASM module to bind to this service.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the monitor-service binding. When a monitor is UP, the weight assigned to its binding with the service determines how much the monitor contributes toward keeping the health of the service above the value configured for the Monitor Threshold parameter.",
			},

			// Convenience blocks (shared with the resource model).
			"lbvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the lb vserver to which the service is bound.",
			},
			"lbmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the lb monitor bound to the service.",
			},
			"snienable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the Server Name Indication (SNI) feature on the service (SSL services only).",
			},
			"commonname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL service.",
			},
			"wait_until_disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set, the provider waits until the service reaches the DISABLED state before returning.",
			},
			"disabled_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum duration to wait for the service to reach the DISABLED state.",
			},
			"disabled_poll_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Delay before the first poll while waiting for the DISABLED state.",
			},
			"disabled_poll_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval between polls while waiting for the DISABLED state.",
			},

			// Read-only (GET-only) operational/status attributes surfaced by the
			// data source (intentionally NOT modeled on the resource). All Computed.
			"numofconnections": schema.Int64Attribute{
				Computed:    true,
				Description: "This will tell the number of client side connections are still open.",
			},
			"policyname": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the policyname for which this service is bound.",
			},
			"serviceconftype": schema.BoolAttribute{
				Computed:    true,
				Description: "The configuration type of the service.",
			},
			"serviceconftype2": schema.StringAttribute{
				Computed:    true,
				Description: "The configuration type of the service (Internal/Dynamic/Configured).",
			},
			"value": schema.StringAttribute{
				Computed:    true,
				Description: "SSL status.",
			},
			"gslb": schema.StringAttribute{
				Computed:    true,
				Description: "The GSLB option for the corresponding virtual server.",
			},
			"dup_state": schema.StringAttribute{
				Computed:    true,
				Description: "Added this field for getting state value from table.",
			},
			"publicip": schema.StringAttribute{
				Computed:    true,
				Description: "public ip.",
			},
			"publicport": schema.Int64Attribute{
				Computed:    true,
				Description: "public port.",
			},
			"svrstate": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the service.",
			},
			"monitor_state": schema.StringAttribute{
				Computed:    true,
				Description: "The running state of the monitor on this service.",
			},
			"monstatcode": schema.Int64Attribute{
				Computed:    true,
				Description: "The code indicating the monitor response.",
			},
			"lastresponse": schema.StringAttribute{
				Computed:    true,
				Description: "The string form of monstatcode.",
			},
			"responsetime": schema.Int64Attribute{
				Computed:    true,
				Description: "Response time of this monitor.",
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
			"statechangetimesec": schema.StringAttribute{
				Computed:    true,
				Description: "Time when last state change happened. Seconds part.",
			},
			"statechangetimemsec": schema.Int64Attribute{
				Computed:    true,
				Description: "Time at which last state change happened. Milliseconds part.",
			},
			"tickssincelaststatechange": schema.Int64Attribute{
				Computed:    true,
				Description: "Time in 10 millisecond ticks since the last state change.",
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
			"serviceipstr": schema.StringAttribute{
				Computed:    true,
				Description: "This field has been intorduced to show the dbs services ip.",
			},
			"oracleserverversion": schema.StringAttribute{
				Computed:    true,
				Description: "Oracle server version.",
			},
			"nodefaultbindings": schema.StringAttribute{
				Computed:    true,
				Description: "to determine if the configuration will have default ssl CIPHER and ECC curve bindings.",
			},
			"monuserstatusmesg": schema.StringAttribute{
				Computed:    true,
				Description: "This field has been introduced to show user monitor failure reasons.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether the service is built-in (MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL).",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// serviceDataSourceSetAttrFromGet projects a NITRO service GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func serviceDataSourceSetAttrFromGet(ctx context.Context, data *ServiceDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In serviceDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes surfaced as read-back outputs.
	data.Internal = utils.MapGetBool(g, "Internal")
	data.Accessdown = utils.MapGetString(g, "accessdown")
	data.Aigwprofilename = utils.MapGetString(g, "aigwprofilename")
	data.All = utils.MapGetBool(g, "all")
	data.Appflowlog = utils.MapGetString(g, "appflowlog")
	data.Cacheable = utils.MapGetString(g, "cacheable")
	data.Cachetype = utils.MapGetString(g, "cachetype")
	data.Cip = utils.MapGetString(g, "cip")
	data.Cipheader = utils.MapGetString(g, "cipheader")
	data.Cka = utils.MapGetString(g, "cka")
	data.Cleartextport = utils.MapGetInt64(g, "cleartextport")
	data.Clttimeout = utils.MapGetInt64(g, "clttimeout")
	data.Cmp = utils.MapGetString(g, "cmp")
	data.Comment = utils.MapGetString(g, "comment")
	data.Contentinspectionprofilename = utils.MapGetString(g, "contentinspectionprofilename")
	data.Customserverid = utils.MapGetString(g, "customserverid")
	data.Dnsprofilename = utils.MapGetString(g, "dnsprofilename")
	data.Downstateflush = utils.MapGetString(g, "downstateflush")
	data.Hashid = utils.MapGetInt64(g, "hashid")
	data.Healthmonitor = utils.MapGetString(g, "healthmonitor")
	data.Httpprofilename = utils.MapGetString(g, "httpprofilename")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Maxbandwidth = utils.MapGetInt64(g, "maxbandwidth")
	data.Maxclient = utils.MapGetInt64(g, "maxclient")
	data.Maxreq = utils.MapGetInt64(g, "maxreq")
	data.Mcpprofilename = utils.MapGetString(g, "mcpprofilename")
	data.Monconnectionclose = utils.MapGetString(g, "monconnectionclose")
	data.Monitornamesvc = utils.MapGetString(g, "monitor_name_svc")
	data.Monthreshold = utils.MapGetInt64(g, "monthreshold")
	data.Netprofile = utils.MapGetString(g, "netprofile")
	data.Riseapbrstatsmsgcode = utils.MapGetInt64(g, "riseapbrstatsmsgcode")
	data.Pathmonitor = utils.MapGetString(g, "pathmonitor")
	data.Pathmonitorindv = utils.MapGetString(g, "pathmonitorindv")
	data.Port = utils.MapGetInt64(g, "port")
	data.Processlocal = utils.MapGetString(g, "processlocal")
	data.Quicprofilename = utils.MapGetString(g, "quicprofilename")
	data.Rtspsessionidremap = utils.MapGetString(g, "rtspsessionidremap")
	data.Serverid = utils.MapGetInt64(g, "serverid")
	data.Servername = utils.MapGetString(g, "servername")
	data.Servicetype = utils.MapGetString(g, "servicetype")
	data.Svrtimeout = utils.MapGetInt64(g, "svrtimeout")
	data.Tcpb = utils.MapGetString(g, "tcpb")
	data.Tcpprofilename = utils.MapGetString(g, "tcpprofilename")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Useproxyport = utils.MapGetString(g, "useproxyport")
	data.Usip = utils.MapGetString(g, "usip")
	data.Wasmmodule = utils.MapGetString(g, "wasmmodule")
	data.Weight = utils.MapGetInt64(g, "weight")

	// sp: the ADC reports "ON (but effectively OFF)"; normalise to "ON" (SDK v2 parity).
	if val, ok := g["sp"]; ok && val != nil {
		if s, ok2 := val.(string); ok2 && s == "ON (but effectively OFF)" {
			data.Sp = types.StringValue("ON")
		} else {
			data.Sp = types.StringValue(utils.AnyToString(val))
		}
	} else {
		data.Sp = types.StringNull()
	}

	// state: derived from svrstate (SDK v2 parity).
	if val, ok := g["svrstate"]; ok && val != nil {
		if s, ok2 := val.(string); ok2 && s == "OUT OF SERVICE" {
			data.State = types.StringValue("DISABLED")
		} else {
			data.State = types.StringValue("ENABLED")
		}
	} else {
		data.State = types.StringNull()
	}

	// Action-only / write-only inputs the base service GET never returns -> Null.
	data.Ip = types.StringNull()
	data.Delay = types.Int64Null()
	data.Graceful = types.StringNull()

	// Convenience blocks are not part of the plain service GET -> Null.
	data.Lbvserver = types.StringNull()
	data.Lbmonitor = types.StringNull()
	data.Snienable = types.StringNull()
	data.Commonname = types.StringNull()
	data.WaitUntilDisabled = types.BoolNull()
	data.DisabledTimeout = types.StringNull()
	data.DisabledPollDelay = types.StringNull()
	data.DisabledPollInterval = types.StringNull()

	// Read-only (GET-only) operational/status metadata.
	data.Numofconnections = utils.MapGetInt64(g, "numofconnections")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Serviceconftype = utils.MapGetBool(g, "serviceconftype")
	data.Serviceconftype2 = utils.MapGetString(g, "serviceconftype2")
	data.Value = utils.MapGetString(g, "value")
	data.Gslb = utils.MapGetString(g, "gslb")
	data.DupState = utils.MapGetString(g, "dup_state")
	data.Publicip = utils.MapGetString(g, "publicip")
	data.Publicport = utils.MapGetInt64(g, "publicport")
	data.Svrstate = utils.MapGetString(g, "svrstate")
	data.MonitorState = utils.MapGetString(g, "monitor_state")
	data.Monstatcode = utils.MapGetInt64(g, "monstatcode")
	data.Lastresponse = utils.MapGetString(g, "lastresponse")
	data.Responsetime = utils.MapGetInt64(g, "responsetime")
	data.Monstatparam1 = utils.MapGetInt64(g, "monstatparam1")
	data.Monstatparam2 = utils.MapGetInt64(g, "monstatparam2")
	data.Monstatparam3 = utils.MapGetInt64(g, "monstatparam3")
	data.Statechangetimesec = utils.MapGetString(g, "statechangetimesec")
	data.Statechangetimemsec = utils.MapGetInt64(g, "statechangetimemsec")
	data.Tickssincelaststatechange = utils.MapGetInt64(g, "tickssincelaststatechange")
	data.Stateupdatereason = utils.MapGetInt64(g, "stateupdatereason")
	data.Clmonowner = utils.MapGetInt64(g, "clmonowner")
	data.Clmonview = utils.MapGetInt64(g, "clmonview")
	data.Serviceipstr = utils.MapGetString(g, "serviceipstr")
	data.Oracleserverversion = utils.MapGetString(g, "oracleserverversion")
	data.Nodefaultbindings = utils.MapGetString(g, "nodefaultbindings")
	data.Monuserstatusmesg = utils.MapGetString(g, "monuserstatusmesg")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
