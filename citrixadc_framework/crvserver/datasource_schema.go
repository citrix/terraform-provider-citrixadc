package crvserver

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CrvserverDataSourceModel is the data-source-specific model, decoupled from
// CrvserverResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only status/binding metadata attributes that the
// resource deliberately omits (ip, value, ngname, type, curstate, status,
// authentication, homepage, rule, policyname, pipolicyhits, servicename,
// weight, targetvserver, priority, somethod, sopersistence, lbvserver,
// bindpoint, invoke, labeltype, labelname, gotopriorityexpression,
// nodefaultbindings). Every non-key attribute is Computed.
type CrvserverDataSourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Appflowlog               types.String `tfsdk:"appflowlog"`
	Arp                      types.String `tfsdk:"arp"`
	Backendssl               types.String `tfsdk:"backendssl"`
	Backupvserver            types.String `tfsdk:"backupvserver"`
	Cachetype                types.String `tfsdk:"cachetype"`
	Cachevserver             types.String `tfsdk:"cachevserver"`
	Clttimeout               types.Int64  `tfsdk:"clttimeout"`
	Comment                  types.String `tfsdk:"comment"`
	Destinationvserver       types.String `tfsdk:"destinationvserver"`
	Disableprimaryondown     types.String `tfsdk:"disableprimaryondown"`
	Disallowserviceaccess    types.String `tfsdk:"disallowserviceaccess"`
	Dnsvservername           types.String `tfsdk:"dnsvservername"`
	Domain                   types.String `tfsdk:"domain"`
	Downstateflush           types.String `tfsdk:"downstateflush"`
	Format                   types.String `tfsdk:"format"`
	Ghost                    types.String `tfsdk:"ghost"`
	Httpprofilename          types.String `tfsdk:"httpprofilename"`
	Icmpvsrresponse          types.String `tfsdk:"icmpvsrresponse"`
	Ipset                    types.String `tfsdk:"ipset"`
	Ipv46                    types.String `tfsdk:"ipv46"`
	L2conn                   types.String `tfsdk:"l2conn"`
	Listenpolicy             types.String `tfsdk:"listenpolicy"`
	Listenpriority           types.Int64  `tfsdk:"listenpriority"`
	Map                      types.String `tfsdk:"map"`
	Name                     types.String `tfsdk:"name"` // Required lookup key
	Netprofile               types.String `tfsdk:"netprofile"`
	Newname                  types.String `tfsdk:"newname"`
	Onpolicymatch            types.String `tfsdk:"onpolicymatch"`
	Originusip               types.String `tfsdk:"originusip"`
	Port                     types.Int64  `tfsdk:"port"`
	Precedence               types.String `tfsdk:"precedence"`
	Probeport                types.Int64  `tfsdk:"probeport"`
	Probeprotocol            types.String `tfsdk:"probeprotocol"`
	Probesuccessresponsecode types.String `tfsdk:"probesuccessresponsecode"`
	Range                    types.Int64  `tfsdk:"range"`
	Redirect                 types.String `tfsdk:"redirect"`
	Redirecturl              types.String `tfsdk:"redirecturl"`
	Reuse                    types.String `tfsdk:"reuse"`
	Rhistate                 types.String `tfsdk:"rhistate"`
	Servicetype              types.String `tfsdk:"servicetype"`
	Sopersistencetimeout     types.Int64  `tfsdk:"sopersistencetimeout"`
	Sothreshold              types.Int64  `tfsdk:"sothreshold"`
	Srcipexpr                types.String `tfsdk:"srcipexpr"`
	State                    types.String `tfsdk:"state"`
	Tcpprobeport             types.Int64  `tfsdk:"tcpprobeport"`
	Tcpprofilename           types.String `tfsdk:"tcpprofilename"`
	Td                       types.Int64  `tfsdk:"td"`
	Useoriginipportforcache  types.String `tfsdk:"useoriginipportforcache"`
	Useportrange             types.String `tfsdk:"useportrange"`
	Via                      types.String `tfsdk:"via"`
	Wasmmodule               types.String `tfsdk:"wasmmodule"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/crvserver.json). Never settable.
	Ip                     types.String `tfsdk:"ip"`
	Value                  types.String `tfsdk:"value"`
	Ngname                 types.String `tfsdk:"ngname"`
	Type                   types.String `tfsdk:"type"`
	Curstate               types.String `tfsdk:"curstate"`
	Status                 types.Int64  `tfsdk:"status"`
	Authentication         types.String `tfsdk:"authentication"`
	Homepage               types.String `tfsdk:"homepage"`
	Rule                   types.String `tfsdk:"rule"`
	Policyname             types.String `tfsdk:"policyname"`
	Pipolicyhits           types.Int64  `tfsdk:"pipolicyhits"`
	Servicename            types.String `tfsdk:"servicename"`
	Weight                 types.Int64  `tfsdk:"weight"`
	Targetvserver          types.String `tfsdk:"targetvserver"`
	Priority               types.Int64  `tfsdk:"priority"`
	Somethod               types.String `tfsdk:"somethod"`
	Sopersistence          types.String `tfsdk:"sopersistence"`
	Lbvserver              types.String `tfsdk:"lbvserver"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Labelname              types.String `tfsdk:"labelname"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Nodefaultbindings      types.String `tfsdk:"nodefaultbindings"`
}

func CrvserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging of AppFlow information.",
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use ARP to determine the destination MAC address.",
			},
			"backendssl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Decides whether the backend connection made by Citrix ADC to the origin server will be HTTP or SSL. Applicable only for SSL type CR Forward proxy vserver.",
			},
			"backupvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the backup virtual server to which traffic is forwarded if the active server becomes unavailable.",
			},
			"cachetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mode of operation for the cache redirection virtual server. Available settings function as follows:\n* TRANSPARENT - Intercept all traffic flowing to the appliance and apply cache redirection policies to determine whether content should be served from the cache or from the origin server.\n* FORWARD - Resolve the hostname of the incoming request, by using a DNS server, and forward requests for non-cacheable content to the resolved origin servers. Cacheable requests are sent to the configured cache servers.\n* REVERSE - Configure reverse proxy caches for specific origin servers. Incoming traffic directed to the reverse proxy can either be served from a cache server or be sent to the origin server with or without modification to the URL.\nThe default value for cache type is TRANSPARENT if service is HTTP or SSL whereas the default cache type is FORWARD if the service is HDX.",
			},
			"cachevserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the default cache virtual server to which to redirect requests (the default target of the cache redirection virtual server).",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time-out value, in seconds, after which to terminate an idle client connection.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this virtual server.",
			},
			"destinationvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination virtual server for a transparent or forward proxy cache redirection virtual server.",
			},
			"disableprimaryondown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Continue sending traffic to a backup virtual server even after the primary virtual server comes UP from the DOWN state.",
			},
			"disallowserviceaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is effective when a FORWARD type cr vserver is added. By default, this parameter is DISABLED. When it is ENABLED, backend services cannot be accessed through a FORWARD type cr vserver.",
			},
			"dnsvservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS virtual server that resolves domain names arriving at the forward proxy virtual server.\nNote: This parameter applies only to forward proxy virtual servers, not reverse or transparent.",
			},
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default domain for reverse proxies. Domains are configured to direct an incoming request from a specified source domain to a specified target domain. There can be several configured pairs of source and target domains. You can select one pair to be the default. If the host header or URL of an incoming request does not include a source domain, this option sends the request to the specified target domain.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform delayed cleanup of connections to this virtual server.",
			},
			"format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"ghost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the profile containing HTTP configuration information for cache redirection virtual server.",
			},
			"icmpvsrresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Criterion for responding to PING requests sent to this virtual server. If ACTIVE, respond only if the virtual server is available. If PASSIVE, respond even if the virtual server is not available.",
			},
			"ipset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current cr vserver",
			},
			"ipv46": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address of the cache redirection virtual server. Usually a public IP address. Clients send connection requests to this IP address.\nNote: For a transparent cache redirection virtual server, use an asterisk (*) to specify a wildcard virtual server address.",
			},
			"l2conn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use L2 parameters, such as MAC, VLAN, and channel to identify a connection.",
			},
			"listenpolicy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the listen policy for the cache redirection virtual server. Can be either an in-line expression or the name of a named expression.",
			},
			"listenpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of the listen policy specified by the Listen Policy parameter. The lower the number, higher the priority.",
			},
			"map": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Obsolete.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the cache redirection virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the cache redirection virtual server is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my server\" or 'my server').",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network profile containing network configurations for the cache redirection virtual server.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the cache redirection virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"onpolicymatch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Redirect requests that match the policy to either the cache or the origin server, as specified.\nNote: For this option to work, you must set the cache redirection type to POLICY.",
			},
			"originusip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the client's IP address as the source IP address in requests sent to the origin server.\nNote: You can enable this parameter to implement fully transparent CR deployment.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number of the virtual server.",
			},
			"precedence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy (URL or RULE) that takes precedence on the cache redirection virtual server. Applies only to cache redirection virtual servers that have both URL and RULE based policies. If you specify URL, URL based policies are applied first, in the following order:\n1.   Domain and exact URL\n2.   Domain, prefix and suffix\n3.   Domain and suffix\n4.   Domain and prefix\n5.   Domain only\n6.   Exact URL\n7.   Prefix and suffix\n8.   Suffix only\n9.   Prefix only\n10.  Default\nIf you specify RULE, the rule based policies are applied before URL based policies are applied.",
			},
			"probeport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC provides support for external health check of the vserver status. Select port for HTTP/TCP monitring",
			},
			"probeprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC provides support for external health check of the vserver status. Select HTTP or TCP probes for healthcheck",
			},
			"probesuccessresponsecode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP code to return in SUCCESS case.",
			},
			"range": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of consecutive IP addresses, starting with the address specified by the IPAddress parameter, to include in a range of addresses assigned to this virtual server.",
			},
			"redirect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of cache server to which to redirect HTTP requests. Available settings function as follows:\n* CACHE - Direct all requests to the cache.\n* POLICY - Apply the cache redirection policy to determine whether the request should be directed to the cache or to the origin.\n* ORIGIN - Direct all requests to the origin server.",
			},
			"redirecturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the server to which to redirect traffic if the cache redirection virtual server configured on the Citrix ADC becomes unavailable.",
			},
			"reuse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Reuse TCP connections to the origin server across client connections. Do not set this parameter unless the Service Type parameter is set to HTTP. If you set this parameter to OFF, the possible settings of the Redirect parameter function as follows:\n* CACHE - TCP connections to the cache servers are not reused.\n* ORIGIN - TCP connections to the origin servers are not reused.\n* POLICY - TCP connections to the origin servers are not reused.\nIf you set the Reuse parameter to ON, connections to origin servers and connections to cache servers are reused.",
			},
			"rhistate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A host route is injected according to the setting on the virtual servers\n            * If set to PASSIVE on all the virtual servers that share the IP address, the appliance always injects the hostroute.\n            * If set to ACTIVE on all the virtual servers that share the IP address, the appliance injects even if one virtual server is UP.\n            * If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance, injects even if one virtual server set to ACTIVE is UP.",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol (type of service) handled by the virtual server.",
			},
			"sopersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time-out, in minutes, for spillover persistence.",
			},
			"sothreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "For CONNECTION (or) DYNAMICCONNECTION spillover, the number of connections above which the virtual server enters spillover mode. For BANDWIDTH spillover, the amount of incoming and outgoing traffic (in Kbps) before spillover. For HEALTH spillover, the percentage of active services (by weight) below which spillover occurs.",
			},
			"srcipexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression used to extract the source IP addresses from the requests originating from the cache. Can be either an in-line expression or the name of a named expression.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the cache redirection virtual server.",
			},
			"tcpprobeport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for external TCP probe. NetScaler provides support for external TCP health check of the vserver status over the selected port. This option is only supported for vservers assigned with an IPAddress or ipset.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the profile containing TCP configuration information for the cache redirection virtual server.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"useoriginipportforcache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use origin ip/port while forwarding request to the cache. Change the destination IP, destination port of the request came to CR vserver to Origin IP and Origin Port and forward it to Cache",
			},
			"useportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use a port number from the port range (set by using the set ns param command, or in the Create Virtual Server (Cache Redirection) dialog box) as the source port in the requests sent to the origin server.",
			},
			"via": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert a via header in each HTTP request. In the case of a cache miss, the request is redirected from the cache server to the origin server. This header indicates whether the request is being sent from a cache server.",
			},
			"wasmmodule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the WASM module to assign to this virtual server.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "IP address of the cache redirection virtual server.",
			},
			"value": schema.StringAttribute{
				Computed:    true,
				Description: "The ssl card status for the transparent ssl cr vserver. Possible values = Certkey/Certkeybundle/Vault not bound/Cert-store not usable, SSL feature disabled",
			},
			"ngname": schema.StringAttribute{
				Computed:    true,
				Description: "Nodegroup devno to which this crvserver belongs to.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "Virtual server type. Possible values = CONTENT, ADDRESS",
			},
			"curstate": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the cr vserver. Possible values = UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED",
			},
			"status": schema.Int64Attribute{
				Computed:    true,
				Description: "Status.",
			},
			"authentication": schema.StringAttribute{
				Computed:    true,
				Description: "Authentication. Possible values = ON, OFF",
			},
			"homepage": schema.StringAttribute{
				Computed:    true,
				Description: "Home page.",
			},
			"rule": schema.StringAttribute{
				Computed:    true,
				Description: "Rule.",
			},
			"policyname": schema.StringAttribute{
				Computed:    true,
				Description: "Policies bound to this vserver.",
			},
			"pipolicyhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"servicename": schema.StringAttribute{
				Computed:    true,
				Description: "Service name.",
			},
			"weight": schema.Int64Attribute{
				Computed:    true,
				Description: "Weight for this service.",
			},
			"targetvserver": schema.StringAttribute{
				Computed:    true,
				Description: "The CSW target server names.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "The priority for the policy.",
			},
			"somethod": schema.StringAttribute{
				Computed:    true,
				Description: "The spillover factor. When the main virtual server reaches this spillover threshold, it will give further traffic to the backupvserver. Possible values = CONNECTION, DYNAMICCONNECTION, BANDWIDTH, HEALTH, LLMQUOTA, NONE",
			},
			"sopersistence": schema.StringAttribute{
				Computed:    true,
				Description: "The state of spillover persistence. Possible values = ENABLED, DISABLED",
			},
			"lbvserver": schema.StringAttribute{
				Computed:    true,
				Description: "The Default target server name.",
			},
			"bindpoint": schema.StringAttribute{
				Computed:    true,
				Description: "The bindpoint to which the policy is bound. Possible values = REQUEST, RESPONSE, ICA_REQUEST",
			},
			"invoke": schema.BoolAttribute{
				Computed:    true,
				Description: "Invoke flag.",
			},
			"labeltype": schema.StringAttribute{
				Computed:    true,
				Description: "The invocation type. Possible values = reqvserver, resvserver, policylabel",
			},
			"labelname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the label invoked.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"nodefaultbindings": schema.StringAttribute{
				Computed:    true,
				Description: "to determine if the configuration will have default ssl CIPHER and ECC curve bindings. Possible values = YES, NO",
			},
		},
	}
}

// crvserverDataSourceSetAttrFromGet projects a NITRO crvserver GET response onto
// the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func crvserverDataSourceSetAttrFromGet(ctx context.Context, data *CrvserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In crvserverDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Appflowlog = utils.MapGetString(g, "appflowlog")
	data.Arp = utils.MapGetString(g, "arp")
	data.Backendssl = utils.MapGetString(g, "backendssl")
	data.Backupvserver = utils.MapGetString(g, "backupvserver")
	data.Cachetype = utils.MapGetString(g, "cachetype")
	data.Cachevserver = utils.MapGetString(g, "cachevserver")
	data.Clttimeout = utils.MapGetInt64(g, "clttimeout")
	data.Comment = utils.MapGetString(g, "comment")
	data.Destinationvserver = utils.MapGetString(g, "destinationvserver")
	data.Disableprimaryondown = utils.MapGetString(g, "disableprimaryondown")
	data.Disallowserviceaccess = utils.MapGetString(g, "disallowserviceaccess")
	data.Dnsvservername = utils.MapGetString(g, "dnsvservername")
	data.Domain = utils.MapGetString(g, "domain")
	data.Downstateflush = utils.MapGetString(g, "downstateflush")
	data.Format = utils.MapGetString(g, "format")
	data.Ghost = utils.MapGetString(g, "ghost")
	data.Httpprofilename = utils.MapGetString(g, "httpprofilename")
	data.Icmpvsrresponse = utils.MapGetString(g, "icmpvsrresponse")
	data.Ipset = utils.MapGetString(g, "ipset")
	data.Ipv46 = utils.MapGetString(g, "ipv46")
	data.L2conn = utils.MapGetString(g, "l2conn")
	data.Listenpolicy = utils.MapGetString(g, "listenpolicy")
	data.Listenpriority = utils.MapGetInt64(g, "listenpriority")
	data.Map = utils.MapGetString(g, "map")
	data.Netprofile = utils.MapGetString(g, "netprofile")
	data.Onpolicymatch = utils.MapGetString(g, "onpolicymatch")
	data.Originusip = utils.MapGetString(g, "originusip")
	data.Port = utils.MapGetInt64(g, "port")
	data.Precedence = utils.MapGetString(g, "precedence")
	data.Probeport = utils.MapGetInt64(g, "probeport")
	data.Probeprotocol = utils.MapGetString(g, "probeprotocol")
	data.Probesuccessresponsecode = utils.MapGetString(g, "probesuccessresponsecode")
	data.Range = utils.MapGetInt64(g, "range")
	data.Redirect = utils.MapGetString(g, "redirect")
	data.Redirecturl = utils.MapGetString(g, "redirecturl")
	data.Reuse = utils.MapGetString(g, "reuse")
	data.Rhistate = utils.MapGetString(g, "rhistate")
	data.Servicetype = utils.MapGetString(g, "servicetype")
	data.Sopersistencetimeout = utils.MapGetInt64(g, "sopersistencetimeout")
	data.Sothreshold = utils.MapGetInt64(g, "sothreshold")
	data.Srcipexpr = utils.MapGetString(g, "srcipexpr")
	data.State = utils.MapGetString(g, "state")
	data.Tcpprobeport = utils.MapGetInt64(g, "tcpprobeport")
	data.Tcpprofilename = utils.MapGetString(g, "tcpprofilename")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Useoriginipportforcache = utils.MapGetString(g, "useoriginipportforcache")
	data.Useportrange = utils.MapGetString(g, "useportrange")
	data.Via = utils.MapGetString(g, "via")
	data.Wasmmodule = utils.MapGetString(g, "wasmmodule")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Ip = utils.MapGetString(g, "ip")
	data.Value = utils.MapGetString(g, "value")
	data.Ngname = utils.MapGetString(g, "ngname")
	data.Type = utils.MapGetString(g, "type")
	data.Curstate = utils.MapGetString(g, "curstate")
	data.Status = utils.MapGetInt64(g, "status")
	data.Authentication = utils.MapGetString(g, "authentication")
	data.Homepage = utils.MapGetString(g, "homepage")
	data.Rule = utils.MapGetString(g, "rule")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Pipolicyhits = utils.MapGetInt64(g, "pipolicyhits")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Weight = utils.MapGetInt64(g, "weight")
	data.Targetvserver = utils.MapGetString(g, "targetvserver")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Somethod = utils.MapGetString(g, "somethod")
	data.Sopersistence = utils.MapGetString(g, "sopersistence")
	data.Lbvserver = utils.MapGetString(g, "lbvserver")
	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Nodefaultbindings = utils.MapGetString(g, "nodefaultbindings")
}
