package crvserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cr"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// CrvserverResourceModel describes the resource data model.
type CrvserverResourceModel struct {
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
	Name                     types.String `tfsdk:"name"`
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
}

func (r *CrvserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the crvserver resource.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable logging of AppFlow information.",
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use ARP to determine the destination MAC address.",
			},
			"backendssl": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Decides whether the backend connection made by Citrix ADC to the origin server will be HTTP or SSL. Applicable only for SSL type CR Forward proxy vserver.",
			},
			"backupvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the backup virtual server to which traffic is forwarded if the active server becomes unavailable.",
			},
			"cachetype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Continue sending traffic to a backup virtual server even after the primary virtual server comes UP from the DOWN state.",
			},
			"disallowserviceaccess": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Perform delayed cleanup of connections to this virtual server.",
			},
			"format": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"ghost": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the profile containing HTTP configuration information for cache redirection virtual server.",
			},
			"icmpvsrresponse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("PASSIVE"),
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
				Default:     stringdefault.StaticString("NONE"),
				Description: "String specifying the listen policy for the cache redirection virtual server. Can be either an in-line expression or the name of a named expression.",
			},
			"listenpriority": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(101),
				Description: "Priority of the listen policy specified by the Listen Policy parameter. The lower the number, higher the priority.",
			},
			"map": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the cache redirection virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"onpolicymatch": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ORIGIN"),
				Description: "Redirect requests that match the policy to either the cache or the origin server, as specified.\nNote: For this option to work, you must set the cache redirection type to POLICY.",
			},
			"originusip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the client's IP address as the source IP address in requests sent to the origin server.\nNote: You can enable this parameter to implement fully transparent CR deployment.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(80),
				Description: "Port number of the virtual server.",
			},
			"precedence": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("RULE"),
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
				Default:     stringdefault.StaticString("200 OK"),
				Description: "HTTP code to return in SUCCESS case.",
			},
			"range": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Number of consecutive IP addresses, starting with the address specified by the IPAddress parameter, to include in a range of addresses assigned to this virtual server.",
			},
			"redirect": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("POLICY"),
				Description: "Type of cache server to which to redirect HTTP requests. Available settings function as follows:\n* CACHE - Direct all requests to the cache.\n* POLICY - Apply the cache redirection policy to determine whether the request should be directed to the cache or to the origin.\n* ORIGIN - Direct all requests to the origin server.",
			},
			"redirecturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL of the server to which to redirect traffic if the cache redirection virtual server configured on the Citrix ADC becomes unavailable.",
			},
			"reuse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "Reuse TCP connections to the origin server across client connections. Do not set this parameter unless the Service Type parameter is set to HTTP. If you set this parameter to OFF, the possible settings of the Redirect parameter function as follows:\n* CACHE - TCP connections to the cache servers are not reused.\n* ORIGIN - TCP connections to the origin servers are not reused.\n* POLICY - TCP connections to the origin servers are not reused.\nIf you set the Reuse parameter to ON, connections to origin servers and connections to cache servers are reused.",
			},
			"rhistate": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("PASSIVE"),
				Description: "A host route is injected according to the setting on the virtual servers\n            * If set to PASSIVE on all the virtual servers that share the IP address, the appliance always injects the hostroute.\n            * If set to ACTIVE on all the virtual servers that share the IP address, the appliance injects even if one virtual server is UP.\n            * If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance, injects even if one virtual server set to ACTIVE is UP.",
			},
			"servicetype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol (type of service) handled by the virtual server.",
			},
			"sopersistencetimeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Time-out, in minutes, for spillover persistence.",
			},
			"sothreshold": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "For CONNECTION (or) DYNAMICCONNECTION spillover, the number of connections above which the virtual server enters spillover mode. For BANDWIDTH spillover, the amount of incoming and outgoing traffic (in Kbps) before spillover. For HEALTH spillover, the percentage of active services (by weight) below which spillover occurs.",
			},
			"srcipexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression used to extract the source IP addresses from the requests originating from the cache. Can be either an in-line expression or the name of a named expression.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
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
				Default:     stringdefault.StaticString("True"),
				Description: "Insert a via header in each HTTP request. In the case of a cache miss, the request is redirected from the cache server to the origin server. This header indicates whether the request is being sent from a cache server.",
			},
		},
	}
}

func crvserverGetThePayloadFromtheConfig(ctx context.Context, data *CrvserverResourceModel) cr.Crvserver {
	tflog.Debug(ctx, "In crvserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	crvserver := cr.Crvserver{}
	if !data.Appflowlog.IsNull() {
		crvserver.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Arp.IsNull() {
		crvserver.Arp = data.Arp.ValueString()
	}
	if !data.Backendssl.IsNull() {
		crvserver.Backendssl = data.Backendssl.ValueString()
	}
	if !data.Backupvserver.IsNull() {
		crvserver.Backupvserver = data.Backupvserver.ValueString()
	}
	if !data.Cachetype.IsNull() {
		crvserver.Cachetype = data.Cachetype.ValueString()
	}
	if !data.Cachevserver.IsNull() {
		crvserver.Cachevserver = data.Cachevserver.ValueString()
	}
	if !data.Clttimeout.IsNull() {
		crvserver.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Comment.IsNull() {
		crvserver.Comment = data.Comment.ValueString()
	}
	if !data.Destinationvserver.IsNull() {
		crvserver.Destinationvserver = data.Destinationvserver.ValueString()
	}
	if !data.Disableprimaryondown.IsNull() {
		crvserver.Disableprimaryondown = data.Disableprimaryondown.ValueString()
	}
	if !data.Disallowserviceaccess.IsNull() {
		crvserver.Disallowserviceaccess = data.Disallowserviceaccess.ValueString()
	}
	if !data.Dnsvservername.IsNull() {
		crvserver.Dnsvservername = data.Dnsvservername.ValueString()
	}
	if !data.Domain.IsNull() {
		crvserver.Domain = data.Domain.ValueString()
	}
	if !data.Downstateflush.IsNull() {
		crvserver.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Format.IsNull() {
		crvserver.Format = data.Format.ValueString()
	}
	if !data.Ghost.IsNull() {
		crvserver.Ghost = data.Ghost.ValueString()
	}
	if !data.Httpprofilename.IsNull() {
		crvserver.Httpprofilename = data.Httpprofilename.ValueString()
	}
	if !data.Icmpvsrresponse.IsNull() {
		crvserver.Icmpvsrresponse = data.Icmpvsrresponse.ValueString()
	}
	if !data.Ipset.IsNull() {
		crvserver.Ipset = data.Ipset.ValueString()
	}
	if !data.Ipv46.IsNull() {
		crvserver.Ipv46 = data.Ipv46.ValueString()
	}
	if !data.L2conn.IsNull() {
		crvserver.L2conn = data.L2conn.ValueString()
	}
	if !data.Listenpolicy.IsNull() {
		crvserver.Listenpolicy = data.Listenpolicy.ValueString()
	}
	if !data.Listenpriority.IsNull() {
		crvserver.Listenpriority = utils.IntPtr(int(data.Listenpriority.ValueInt64()))
	}
	if !data.Map.IsNull() {
		crvserver.Map = data.Map.ValueString()
	}
	if !data.Name.IsNull() {
		crvserver.Name = data.Name.ValueString()
	}
	if !data.Netprofile.IsNull() {
		crvserver.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Newname.IsNull() {
		crvserver.Newname = data.Newname.ValueString()
	}
	if !data.Onpolicymatch.IsNull() {
		crvserver.Onpolicymatch = data.Onpolicymatch.ValueString()
	}
	if !data.Originusip.IsNull() {
		crvserver.Originusip = data.Originusip.ValueString()
	}
	if !data.Port.IsNull() {
		crvserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Precedence.IsNull() {
		crvserver.Precedence = data.Precedence.ValueString()
	}
	if !data.Probeport.IsNull() {
		crvserver.Probeport = utils.IntPtr(int(data.Probeport.ValueInt64()))
	}
	if !data.Probeprotocol.IsNull() {
		crvserver.Probeprotocol = data.Probeprotocol.ValueString()
	}
	if !data.Probesuccessresponsecode.IsNull() {
		crvserver.Probesuccessresponsecode = data.Probesuccessresponsecode.ValueString()
	}
	if !data.Range.IsNull() {
		crvserver.Range = utils.IntPtr(int(data.Range.ValueInt64()))
	}
	if !data.Redirect.IsNull() {
		crvserver.Redirect = data.Redirect.ValueString()
	}
	if !data.Redirecturl.IsNull() {
		crvserver.Redirecturl = data.Redirecturl.ValueString()
	}
	if !data.Reuse.IsNull() {
		crvserver.Reuse = data.Reuse.ValueString()
	}
	if !data.Rhistate.IsNull() {
		crvserver.Rhistate = data.Rhistate.ValueString()
	}
	if !data.Servicetype.IsNull() {
		crvserver.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sopersistencetimeout.IsNull() {
		crvserver.Sopersistencetimeout = utils.IntPtr(int(data.Sopersistencetimeout.ValueInt64()))
	}
	if !data.Sothreshold.IsNull() {
		crvserver.Sothreshold = utils.IntPtr(int(data.Sothreshold.ValueInt64()))
	}
	if !data.Srcipexpr.IsNull() {
		crvserver.Srcipexpr = data.Srcipexpr.ValueString()
	}
	if !data.State.IsNull() {
		crvserver.State = data.State.ValueString()
	}
	if !data.Tcpprobeport.IsNull() {
		crvserver.Tcpprobeport = utils.IntPtr(int(data.Tcpprobeport.ValueInt64()))
	}
	if !data.Tcpprofilename.IsNull() {
		crvserver.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Td.IsNull() {
		crvserver.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Useoriginipportforcache.IsNull() {
		crvserver.Useoriginipportforcache = data.Useoriginipportforcache.ValueString()
	}
	if !data.Useportrange.IsNull() {
		crvserver.Useportrange = data.Useportrange.ValueString()
	}
	if !data.Via.IsNull() {
		crvserver.Via = data.Via.ValueString()
	}

	return crvserver
}

func crvserverSetAttrFromGet(ctx context.Context, data *CrvserverResourceModel, getResponseData map[string]interface{}) *CrvserverResourceModel {
	tflog.Debug(ctx, "In crvserverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else {
		data.Appflowlog = types.StringNull()
	}
	if val, ok := getResponseData["arp"]; ok && val != nil {
		data.Arp = types.StringValue(val.(string))
	} else {
		data.Arp = types.StringNull()
	}
	if val, ok := getResponseData["backendssl"]; ok && val != nil {
		data.Backendssl = types.StringValue(val.(string))
	} else {
		data.Backendssl = types.StringNull()
	}
	if val, ok := getResponseData["backupvserver"]; ok && val != nil {
		data.Backupvserver = types.StringValue(val.(string))
	} else {
		data.Backupvserver = types.StringNull()
	}
	if val, ok := getResponseData["cachetype"]; ok && val != nil {
		data.Cachetype = types.StringValue(val.(string))
	} else {
		data.Cachetype = types.StringNull()
	}
	if val, ok := getResponseData["cachevserver"]; ok && val != nil {
		data.Cachevserver = types.StringValue(val.(string))
	} else {
		data.Cachevserver = types.StringNull()
	}
	if val, ok := getResponseData["clttimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Clttimeout = types.Int64Value(intVal)
		}
	} else {
		data.Clttimeout = types.Int64Null()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["destinationvserver"]; ok && val != nil {
		data.Destinationvserver = types.StringValue(val.(string))
	} else {
		data.Destinationvserver = types.StringNull()
	}
	if val, ok := getResponseData["disableprimaryondown"]; ok && val != nil {
		data.Disableprimaryondown = types.StringValue(val.(string))
	} else {
		data.Disableprimaryondown = types.StringNull()
	}
	if val, ok := getResponseData["disallowserviceaccess"]; ok && val != nil {
		data.Disallowserviceaccess = types.StringValue(val.(string))
	} else {
		data.Disallowserviceaccess = types.StringNull()
	}
	if val, ok := getResponseData["dnsvservername"]; ok && val != nil {
		data.Dnsvservername = types.StringValue(val.(string))
	} else {
		data.Dnsvservername = types.StringNull()
	}
	if val, ok := getResponseData["domain"]; ok && val != nil {
		data.Domain = types.StringValue(val.(string))
	} else {
		data.Domain = types.StringNull()
	}
	if val, ok := getResponseData["downstateflush"]; ok && val != nil {
		data.Downstateflush = types.StringValue(val.(string))
	} else {
		data.Downstateflush = types.StringNull()
	}
	if val, ok := getResponseData["format"]; ok && val != nil {
		data.Format = types.StringValue(val.(string))
	} else {
		data.Format = types.StringNull()
	}
	if val, ok := getResponseData["ghost"]; ok && val != nil {
		data.Ghost = types.StringValue(val.(string))
	} else {
		data.Ghost = types.StringNull()
	}
	if val, ok := getResponseData["httpprofilename"]; ok && val != nil {
		data.Httpprofilename = types.StringValue(val.(string))
	} else {
		data.Httpprofilename = types.StringNull()
	}
	if val, ok := getResponseData["icmpvsrresponse"]; ok && val != nil {
		data.Icmpvsrresponse = types.StringValue(val.(string))
	} else {
		data.Icmpvsrresponse = types.StringNull()
	}
	if val, ok := getResponseData["ipset"]; ok && val != nil {
		data.Ipset = types.StringValue(val.(string))
	} else {
		data.Ipset = types.StringNull()
	}
	if val, ok := getResponseData["ipv46"]; ok && val != nil {
		data.Ipv46 = types.StringValue(val.(string))
	} else {
		data.Ipv46 = types.StringNull()
	}
	if val, ok := getResponseData["l2conn"]; ok && val != nil {
		data.L2conn = types.StringValue(val.(string))
	} else {
		data.L2conn = types.StringNull()
	}
	if val, ok := getResponseData["listenpolicy"]; ok && val != nil {
		data.Listenpolicy = types.StringValue(val.(string))
	} else {
		data.Listenpolicy = types.StringNull()
	}
	if val, ok := getResponseData["listenpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Listenpriority = types.Int64Value(intVal)
		}
	} else {
		data.Listenpriority = types.Int64Null()
	}
	if val, ok := getResponseData["map"]; ok && val != nil {
		data.Map = types.StringValue(val.(string))
	} else {
		data.Map = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["netprofile"]; ok && val != nil {
		data.Netprofile = types.StringValue(val.(string))
	} else {
		data.Netprofile = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["onpolicymatch"]; ok && val != nil {
		data.Onpolicymatch = types.StringValue(val.(string))
	} else {
		data.Onpolicymatch = types.StringNull()
	}
	if val, ok := getResponseData["originusip"]; ok && val != nil {
		data.Originusip = types.StringValue(val.(string))
	} else {
		data.Originusip = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["precedence"]; ok && val != nil {
		data.Precedence = types.StringValue(val.(string))
	} else {
		data.Precedence = types.StringNull()
	}
	if val, ok := getResponseData["probeport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Probeport = types.Int64Value(intVal)
		}
	} else {
		data.Probeport = types.Int64Null()
	}
	if val, ok := getResponseData["probeprotocol"]; ok && val != nil {
		data.Probeprotocol = types.StringValue(val.(string))
	} else {
		data.Probeprotocol = types.StringNull()
	}
	if val, ok := getResponseData["probesuccessresponsecode"]; ok && val != nil {
		data.Probesuccessresponsecode = types.StringValue(val.(string))
	} else {
		data.Probesuccessresponsecode = types.StringNull()
	}
	if val, ok := getResponseData["range"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Range = types.Int64Value(intVal)
		}
	} else {
		data.Range = types.Int64Null()
	}
	if val, ok := getResponseData["redirect"]; ok && val != nil {
		data.Redirect = types.StringValue(val.(string))
	} else {
		data.Redirect = types.StringNull()
	}
	if val, ok := getResponseData["redirecturl"]; ok && val != nil {
		data.Redirecturl = types.StringValue(val.(string))
	} else {
		data.Redirecturl = types.StringNull()
	}
	if val, ok := getResponseData["reuse"]; ok && val != nil {
		data.Reuse = types.StringValue(val.(string))
	} else {
		data.Reuse = types.StringNull()
	}
	if val, ok := getResponseData["rhistate"]; ok && val != nil {
		data.Rhistate = types.StringValue(val.(string))
	} else {
		data.Rhistate = types.StringNull()
	}
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	} else {
		data.Servicetype = types.StringNull()
	}
	if val, ok := getResponseData["sopersistencetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sopersistencetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Sopersistencetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["sothreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sothreshold = types.Int64Value(intVal)
		}
	} else {
		data.Sothreshold = types.Int64Null()
	}
	if val, ok := getResponseData["srcipexpr"]; ok && val != nil {
		data.Srcipexpr = types.StringValue(val.(string))
	} else {
		data.Srcipexpr = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["tcpprobeport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tcpprobeport = types.Int64Value(intVal)
		}
	} else {
		data.Tcpprobeport = types.Int64Null()
	}
	if val, ok := getResponseData["tcpprofilename"]; ok && val != nil {
		data.Tcpprofilename = types.StringValue(val.(string))
	} else {
		data.Tcpprofilename = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["useoriginipportforcache"]; ok && val != nil {
		data.Useoriginipportforcache = types.StringValue(val.(string))
	} else {
		data.Useoriginipportforcache = types.StringNull()
	}
	if val, ok := getResponseData["useportrange"]; ok && val != nil {
		data.Useportrange = types.StringValue(val.(string))
	} else {
		data.Useportrange = types.StringNull()
	}
	if val, ok := getResponseData["via"]; ok && val != nil {
		data.Via = types.StringValue(val.(string))
	} else {
		data.Via = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
