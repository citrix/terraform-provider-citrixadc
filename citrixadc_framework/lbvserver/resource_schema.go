package lbvserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbvserverResourceModel describes the resource data model.
type LbvserverResourceModel struct {
	Id                                 types.String `tfsdk:"id"`
	Adfsproxyprofile                   types.String `tfsdk:"adfsproxyprofile"`
	Apiprofile                         types.String `tfsdk:"apiprofile"`
	Appflowlog                         types.String `tfsdk:"appflowlog"`
	Authentication                     types.String `tfsdk:"authentication"`
	Authenticationhost                 types.String `tfsdk:"authenticationhost"`
	Authn401                           types.String `tfsdk:"authn401"`
	Authnprofile                       types.String `tfsdk:"authnprofile"`
	Authnvsname                        types.String `tfsdk:"authnvsname"`
	Backuplbmethod                     types.String `tfsdk:"backuplbmethod"`
	Backuppersistencetimeout           types.Int64  `tfsdk:"backuppersistencetimeout"`
	Backupvserver                      types.String `tfsdk:"backupvserver"`
	Bypassaaaa                         types.String `tfsdk:"bypassaaaa"`
	Cacheable                          types.String `tfsdk:"cacheable"`
	Clttimeout                         types.Int64  `tfsdk:"clttimeout"`
	Comment                            types.String `tfsdk:"comment"`
	Connfailover                       types.String `tfsdk:"connfailover"`
	Cookiename                         types.String `tfsdk:"cookiename"`
	Datalength                         types.Int64  `tfsdk:"datalength"`
	Dataoffset                         types.Int64  `tfsdk:"dataoffset"`
	Dbprofilename                      types.String `tfsdk:"dbprofilename"`
	Dbslb                              types.String `tfsdk:"dbslb"`
	Disableprimaryondown               types.String `tfsdk:"disableprimaryondown"`
	Dns64                              types.String `tfsdk:"dns64"`
	Dnsoverhttps                       types.String `tfsdk:"dnsoverhttps"`
	Dnsprofilename                     types.String `tfsdk:"dnsprofilename"`
	Downstateflush                     types.String `tfsdk:"downstateflush"`
	Hashlength                         types.Int64  `tfsdk:"hashlength"`
	Healththreshold                    types.Int64  `tfsdk:"healththreshold"`
	Httpprofilename                    types.String `tfsdk:"httpprofilename"`
	Httpsredirecturl                   types.String `tfsdk:"httpsredirecturl"`
	Icmpvsrresponse                    types.String `tfsdk:"icmpvsrresponse"`
	Insertvserveripport                types.String `tfsdk:"insertvserveripport"`
	Ipmask                             types.String `tfsdk:"ipmask"`
	Ippattern                          types.String `tfsdk:"ippattern"`
	Ipset                              types.String `tfsdk:"ipset"`
	Ipv46                              types.String `tfsdk:"ipv46"`
	L2conn                             types.String `tfsdk:"l2conn"`
	Lbmethod                           types.String `tfsdk:"lbmethod"`
	Lbprofilename                      types.String `tfsdk:"lbprofilename"`
	Listenpolicy                       types.String `tfsdk:"listenpolicy"`
	Listenpriority                     types.Int64  `tfsdk:"listenpriority"`
	M                                  types.String `tfsdk:"m"`
	Macmoderetainvlan                  types.String `tfsdk:"macmoderetainvlan"`
	Maxautoscalemembers                types.Int64  `tfsdk:"maxautoscalemembers"`
	Minautoscalemembers                types.Int64  `tfsdk:"minautoscalemembers"`
	Mssqlserverversion                 types.String `tfsdk:"mssqlserverversion"`
	Mysqlcharacterset                  types.Int64  `tfsdk:"mysqlcharacterset"`
	Mysqlprotocolversion               types.Int64  `tfsdk:"mysqlprotocolversion"`
	Mysqlservercapabilities            types.Int64  `tfsdk:"mysqlservercapabilities"`
	Mysqlserverversion                 types.String `tfsdk:"mysqlserverversion"`
	Name                               types.String `tfsdk:"name"`
	Netmask                            types.String `tfsdk:"netmask"`
	Netprofile                         types.String `tfsdk:"netprofile"`
	Newname                            types.String `tfsdk:"newname"`
	Newservicerequest                  types.Int64  `tfsdk:"newservicerequest"`
	Newservicerequestincrementinterval types.Int64  `tfsdk:"newservicerequestincrementinterval"`
	Newservicerequestunit              types.String `tfsdk:"newservicerequestunit"`
	Oracleserverversion                types.String `tfsdk:"oracleserverversion"`
	Order                              types.Int64  `tfsdk:"order"`
	Orderthreshold                     types.Int64  `tfsdk:"orderthreshold"`
	Persistavpno                       types.List   `tfsdk:"persistavpno"`
	Persistencebackup                  types.String `tfsdk:"persistencebackup"`
	Persistencetype                    types.String `tfsdk:"persistencetype"`
	Persistmask                        types.String `tfsdk:"persistmask"`
	Port                               types.Int64  `tfsdk:"port"`
	Probeport                          types.Int64  `tfsdk:"probeport"`
	Probeprotocol                      types.String `tfsdk:"probeprotocol"`
	Probesuccessresponsecode           types.String `tfsdk:"probesuccessresponsecode"`
	Processlocal                       types.String `tfsdk:"processlocal"`
	Push                               types.String `tfsdk:"push"`
	Pushlabel                          types.String `tfsdk:"pushlabel"`
	Pushmulticlients                   types.String `tfsdk:"pushmulticlients"`
	Pushvserver                        types.String `tfsdk:"pushvserver"`
	Quicbridgeprofilename              types.String `tfsdk:"quicbridgeprofilename"`
	Quicprofilename                    types.String `tfsdk:"quicprofilename"`
	Range                              types.Int64  `tfsdk:"range"`
	Recursionavailable                 types.String `tfsdk:"recursionavailable"`
	Redirectfromport                   types.Int64  `tfsdk:"redirectfromport"`
	Redirectportrewrite                types.String `tfsdk:"redirectportrewrite"`
	Redirurl                           types.String `tfsdk:"redirurl"`
	Redirurlflags                      types.Bool   `tfsdk:"redirurlflags"`
	Resrule                            types.String `tfsdk:"resrule"`
	Retainconnectionsoncluster         types.String `tfsdk:"retainconnectionsoncluster"`
	Rhistate                           types.String `tfsdk:"rhistate"`
	Rtspnat                            types.String `tfsdk:"rtspnat"`
	Rule                               types.String `tfsdk:"rule"`
	Servicename                        types.String `tfsdk:"servicename"`
	Servicetype                        types.String `tfsdk:"servicetype"`
	Sessionless                        types.String `tfsdk:"sessionless"`
	Skippersistency                    types.String `tfsdk:"skippersistency"`
	Sobackupaction                     types.String `tfsdk:"sobackupaction"`
	Somethod                           types.String `tfsdk:"somethod"`
	Sopersistence                      types.String `tfsdk:"sopersistence"`
	Sopersistencetimeout               types.Int64  `tfsdk:"sopersistencetimeout"`
	Sothreshold                        types.Int64  `tfsdk:"sothreshold"`
	State                              types.String `tfsdk:"state"`
	Tcpprobeport                       types.Int64  `tfsdk:"tcpprobeport"`
	Tcpprofilename                     types.String `tfsdk:"tcpprofilename"`
	Td                                 types.Int64  `tfsdk:"td"`
	Timeout                            types.Int64  `tfsdk:"timeout"`
	Toggleorder                        types.String `tfsdk:"toggleorder"`
	Tosid                              types.Int64  `tfsdk:"tosid"`
	Trofspersistence                   types.String `tfsdk:"trofspersistence"`
	V6netmasklen                       types.Int64  `tfsdk:"v6netmasklen"`
	V6persistmasklen                   types.Int64  `tfsdk:"v6persistmasklen"`
	Vipheader                          types.String `tfsdk:"vipheader"`
	Weight                             types.Int64  `tfsdk:"weight"`
}

func (r *LbvserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbvserver resource.",
			},
			"adfsproxyprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the adfsProxy profile to be used to support ADFSPIP protocol for ADFS servers.",
			},
			"apiprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The API profile where one or more API specs are bounded to.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Apply AppFlow logging to the virtual server.",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable user authentication.",
			},
			"authenticationhost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name (FQDN) of the authentication virtual server to which the user must be redirected for authentication. Make sure that the Authentication parameter is set to ENABLED.",
			},
			"authn401": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable user authentication with HTTP 401 responses.",
			},
			"authnprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the authentication profile to be used when authentication is turned on.",
			},
			"authnvsname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of an authentication virtual server with which to authenticate users.",
			},
			"backuplbmethod": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ROUNDROBIN"),
				Description: "Backup load balancing method. Becomes operational if the primary load balancing me\nthod fails or cannot be used.\n                       Valid only if the primary method is based on static proximity.",
			},
			"backuppersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2),
				Description: "Time period for which backup persistence is in effect.",
			},
			"backupvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the backup virtual server to which to forward requests if the primary virtual server goes DOWN or reaches its spillover threshold.",
			},
			"bypassaaaa": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If this option is enabled while resolving DNS64 query AAAA queries are not sent to back end dns server",
			},
			"cacheable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Route cacheable requests to a cache redirection virtual server. The load balancing virtual server can forward requests only to a transparent cache redirection virtual server that has an IP address and port combination of *:80, so such a cache redirection virtual server must be configured on the appliance.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle time, in seconds, after which a client connection is terminated.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments that you might want to associate with the virtual server.",
			},
			"connfailover": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Mode in which the connection failover feature must operate for the virtual server. After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary appliance. Clients remain connected to the same servers. Available settings function as follows:\n* STATEFUL - The primary appliance shares state information with the secondary appliance, in real time, resulting in some runtime processing overhead.\n* STATELESS - State information is not shared, and the new primary appliance tries to re-create the packet flow on the basis of the information contained in the packets it receives.\n* DISABLED - Connection failover does not occur.",
			},
			"cookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.",
			},
			"datalength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Length of the token to be extracted from the data segment of an incoming packet, for use in the token method of load balancing. The length of the token, specified in bytes, must not be greater than 24 KB. Applicable to virtual servers of type TCP.",
			},
			"dataoffset": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Offset to be considered when extracting a token from the TCP payload. Applicable to virtual servers, of type TCP, using the token method of load balancing. Must be within the first 24 KB of the TCP payload.",
			},
			"dbprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DB profile whose settings are to be applied to the virtual server.",
			},
			"dbslb": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable database specific load balancing for MySQL and MSSQL service types.",
			},
			"disableprimaryondown": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If the primary virtual server goes down, do not allow it to return to primary status until manually enabled.",
			},
			"dns64": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This argument is for enabling/disabling the dns64 on lbvserver",
			},
			"dnsoverhttps": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option is used to enable/disable DNS over HTTPS (DoH) processing.",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the VServer. DNS profile properties will be applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"hashlength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(80),
				Description: "Number of bytes to consider for the hash value used in the URLHASH and DOMAINHASH load balancing methods.",
			},
			"healththreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold in percent of active services below which vserver state is made down. If this threshold is 0, vserver state will be up even if one bound service is up.",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP profile whose settings are to be applied to the virtual server.",
			},
			"httpsredirecturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which all HTTP traffic received on the port specified in the -redirectFromPort parameter is redirected.",
			},
			"icmpvsrresponse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("PASSIVE"),
				Description: "How the Citrix ADC responds to ping requests received for an IP address that is common to one or more virtual servers. Available settings function as follows:\n* If set to PASSIVE on all the virtual servers that share the IP address, the appliance always responds to the ping requests.\n* If set to ACTIVE on all the virtual servers that share the IP address, the appliance responds to the ping requests if at least one of the virtual servers is UP. Otherwise, the appliance does not respond.\n* If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance responds if at least one virtual server with the ACTIVE setting is UP. Otherwise, the appliance does not respond.\nNote: This parameter is available at the virtual server level. A similar parameter, ICMP Response, is available at the IP address level, for IPv4 addresses of type VIP. To set that parameter, use the add ip command in the CLI or the Create IP dialog box in the GUI.",
			},
			"insertvserveripport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert an HTTP header, whose value is the IP address and port number of the virtual server, before forwarding a request to the server. The format of the header is <vipHeader>: <virtual server IP address>_<port number >, where vipHeader is the name that you specify for the header. If the virtual server has an IPv6 address, the address in the header is enclosed in brackets ([ and ]) to separate it from the port number. If you have mapped an IPv4 address to a virtual server's IPv6 address, the value of this parameter determines which IP address is inserted in the header, as follows:\n* VIPADDR - Insert the IP address of the virtual server in the HTTP header regardless of whether the virtual server has an IPv4 address or an IPv6 address. A mapped IPv4 address, if configured, is ignored.\n* V6TOV4MAPPING - Insert the IPv4 address that is mapped to the virtual server's IPv6 address. If a mapped IPv4 address is not configured, insert the IPv6 address.\n* OFF - Disable header insertion.",
			},
			"ipmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP mask, in dotted decimal notation, for the IP Pattern parameter. Can have leading or trailing non-zero octets (for example, 255.255.240.0 or 0.0.255.255). Accordingly, the mask specifies whether the first n bits or the last n bits of the destination IP address in a client request are to be matched with the corresponding bits in the IP pattern. The former is called a forward mask. The latter is called a reverse mask.",
			},
			"ippattern": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server. The IP Mask parameter specifies which part of the destination IP address is matched against the pattern.  Mutually exclusive with the IP Address parameter.\nFor example, if the IP pattern assigned to the virtual server is 198.51.100.0 and the IP mask is 255.255.240.0 (a forward mask), the first 20 bits in the destination IP addresses are matched with the first 20 bits in the pattern. The virtual server accepts requests with IP addresses that range from 198.51.96.1 to 198.51.111.254.  You can also use a pattern such as 0.0.2.2 and a mask such as 0.0.255.255 (a reverse mask).\nIf a destination IP address matches more than one IP pattern, the pattern with the longest match is selected, and the associated virtual server processes the request. For example, if virtual servers vs1 and vs2 have the same IP pattern, 0.0.100.128, but different IP masks of 0.0.255.255 and 0.0.224.255, a destination IP address of 198.51.100.128 has the longest match with the IP pattern of vs1. If a destination IP address matches two or more virtual servers to the same extent, the request is processed by the virtual server whose port number matches the port number in the request.",
			},
			"ipset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current lb vserver",
			},
			"ipv46": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address to assign to the virtual server.",
			},
			"l2conn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use Layer 2 parameters (channel number, MAC address, and VLAN ID) in addition to the 4-tuple (<source IP>:<source port>::<destination IP>:<destination port>) that is used to identify a connection. Allows multiple TCP and non-TCP connections with the same 4-tuple to co-exist on the Citrix ADC.",
			},
			"lbmethod": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("LEASTCONNECTION"),
				Description: "Load balancing method.  The available settings function as follows:\n* ROUNDROBIN - Distribute requests in rotation, regardless of the load. Weights can be assigned to services to enforce weighted round robin distribution.\n* LEASTCONNECTION (default) - Select the service with the fewest connections.\n* LEASTRESPONSETIME - Select the service with the lowest average response time.\n* LEASTBANDWIDTH - Select the service currently handling the least traffic.\n* LEASTPACKETS - Select the service currently serving the lowest number of packets per second.\n* CUSTOMLOAD - Base service selection on the SNMP metrics obtained by custom load monitors.\n* LRTM - Select the service with the lowest response time. Response times are learned through monitoring probes. This method also takes the number of active connections into account.\nAlso available are a number of hashing methods, in which the appliance extracts a predetermined portion of the request, creates a hash of the portion, and then checks whether any previous requests had the same hash value. If it finds a match, it forwards the request to the service that served those previous requests. Following are the hashing methods:\n* URLHASH - Create a hash of the request URL (or part of the URL).\n* DOMAINHASH - Create a hash of the domain name in the request (or part of the domain name). The domain name is taken from either the URL or the Host header. If the domain name appears in both locations, the URL is preferred. If the request does not contain a domain name, the load balancing method defaults to LEASTCONNECTION.\n* DESTINATIONIPHASH - Create a hash of the destination IP address in the IP header.\n* SOURCEIPHASH - Create a hash of the source IP address in the IP header.\n* TOKEN - Extract a token from the request, create a hash of the token, and then select the service to which any previous requests with the same token hash value were sent.\n* SRCIPDESTIPHASH - Create a hash of the string obtained by concatenating the source IP address and destination IP address in the IP header.\n* SRCIPSRCPORTHASH - Create a hash of the source IP address and source port in the IP header.\n* CALLIDHASH - Create a hash of the SIP Call-ID header.\n* USER_TOKEN - Same as TOKEN LB method but token needs to be provided from an extension.",
			},
			"lbprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LB profile which is associated to the vserver",
			},
			"listenpolicy": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NONE"),
				Description: "Expression identifying traffic accepted by the virtual server. Can be either an expression (for example, CLIENT.IP.DST.IN_SUBNET(192.0.2.0/24) or the name of a named expression. In the above example, the virtual server accepts all requests whose destination IP address is in the 192.0.2.0/24 subnet.",
			},
			"listenpriority": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(101),
				Description: "Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.",
			},
			"m": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("IP"),
				Description: "Redirection mode for load balancing. Available settings function as follows:\n* IP - Before forwarding a request to a server, change the destination IP address to the server's IP address.\n* MAC - Before forwarding a request to a server, change the destination MAC address to the server's MAC address.  The destination IP address is not changed. MAC-based redirection mode is used mostly in firewall load balancing deployments.\n* IPTUNNEL - Perform IP-in-IP encapsulation for client IP packets. In the outer IP headers, set the destination IP address to the IP address of the server and the source IP address to the subnet IP (SNIP). The client IP packets are not modified. Applicable to both IPv4 and IPv6 packets.\n* TOS - Encode the virtual server's TOS ID in the TOS field of the IP header.\nYou can use either the IPTUNNEL or the TOS option to implement Direct Server Return (DSR).",
			},
			"macmoderetainvlan": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option is used to retain vlan information of incoming packet when macmode is enabled",
			},
			"maxautoscalemembers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of members expected to be present when vserver is used in Autoscale.",
			},
			"minautoscalemembers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of members expected to be present when vserver is used in Autoscale.",
			},
			"mssqlserverversion": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("2008R2"),
				Description: "For a load balancing virtual server of type MSSQL, the Microsoft SQL Server version. Set this parameter if you expect some clients to run a version different from the version of the database. This setting provides compatibility between the client-side and server-side connections by ensuring that all communication conforms to the server's version.",
			},
			"mysqlcharacterset": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Character set that the virtual server advertises to clients.",
			},
			"mysqlprotocolversion": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MySQL protocol version that the virtual server advertises to clients.",
			},
			"mysqlservercapabilities": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Server capabilities that the virtual server advertises to clients.",
			},
			"mysqlserverversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MySQL server version string that the virtual server advertises to clients.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 subnet mask to apply to the destination IP address or source IP address when the load balancing method is DESTINATIONIPHASH or SOURCEIPHASH.",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network profile to associate with the virtual server. If you set this parameter, the virtual server uses only the IP addresses in the network profile as source IP addresses when initiating connections with servers.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the virtual server.",
			},
			"newservicerequest": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of requests, or percentage of the load on existing services, by which to increase the load on a new service at each interval in slow-start mode. A non-zero value indicates that slow-start is applicable. A zero value indicates that the global RR startup parameter is applied. Changing the value to zero will cause services currently in slow start to take the full traffic as determined by the LB method. Subsequently, any new services added will use the global RR factor.",
			},
			"newservicerequestincrementinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in seconds, between successive increments in the load on a new service or a service whose state has just changed from DOWN to UP. A value of 0 (zero) specifies manual slow start.",
			},
			"newservicerequestunit": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("PER_SECOND"),
				Description: "Units in which to increment load at each interval in slow-start mode.",
			},
			"oracleserverversion": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("10G"),
				Description: "Oracle server version",
			},
			"order": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order number to be assigned to the service when it is bound to the lb vserver.",
			},
			"orderthreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is used to to specify the threshold of minimum number of services to be UP in an order, for it to be considered in Lb decision.",
			},
			"persistavpno": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Persist AVP number for Diameter Persistency.\n            In case this AVP is not defined in Base RFC 3588 and it is nested inside a Grouped AVP,\n            define a sequence of AVP numbers (max 3) in order of parent to child. So say persist AVP number X\n            is nested inside AVP Y which is nested in Z, then define the list as  Z Y X",
			},
			"persistencebackup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Backup persistence type for the virtual server. Becomes operational if the primary persistence mechanism fails.",
			},
			"persistencetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of persistence for the virtual server. Available settings function as follows:\n* SOURCEIP - Connections from the same client IP address belong to the same persistence session.\n* COOKIEINSERT - Connections that have the same HTTP Cookie, inserted by a Set-Cookie directive from a server, belong to the same persistence session.\n* SSLSESSION - Connections that have the same SSL Session ID belong to the same persistence session.\n* CUSTOMSERVERID - Connections with the same server ID form part of the same session. For this persistence type, set the Server ID (CustomServerID) parameter for each service and configure the Rule parameter to identify the server ID in a request.\n* RULE - All connections that match a user defined rule belong to the same persistence session.\n* URLPASSIVE - Requests that have the same server ID in the URL query belong to the same persistence session. The server ID is the hexadecimal representation of the IP address and port of the service to which the request must be forwarded. This persistence type requires a rule to identify the server ID in the request.\n* DESTIP - Connections to the same destination IP address belong to the same persistence session.\n* SRCIPDESTIP - Connections that have the same source IP address and destination IP address belong to the same persistence session.\n* CALLID - Connections that have the same CALL-ID SIP header belong to the same persistence session.\n* RTSPSID - Connections that have the same RTSP Session ID belong to the same persistence session.\n* FIXSESSION - Connections that have the same SenderCompID and TargetCompID values belong to the same persistence session.\n* USERSESSION - Persistence session is created based on the persistence parameter value provided from an extension.",
			},
			"persistmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask for IP based persistence types, for IPv4 virtual servers.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port number for the virtual server.",
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
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "By turning on this option packets destined to a vserver in a cluster will not under go any steering. Turn this option for single packet request response mode or when the upstream device is performing a proper RSS for connection based distribution.",
			},
			"push": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Process traffic with the push virtual server that is bound to this load balancing virtual server.",
			},
			"pushlabel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "Expression for extracting a label from the server's response. Can be either an expression or the name of a named expression.",
			},
			"pushmulticlients": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates.",
			},
			"pushvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates received on the load balancing virtual server that you are configuring.",
			},
			"quicbridgeprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the QUIC Bridge profile whose settings are to be applied to the virtual server.",
			},
			"quicprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of QUIC profile which will be attached to the VServer.",
			},
			"range": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Number of IP addresses that the appliance must generate and assign to the virtual server. The virtual server then functions as a network virtual server, accepting traffic on any of the generated IP addresses. The IP addresses are generated automatically, as follows:\n* For a range of n, the last octet of the address specified by the IP Address parameter increments n-1 times.\n* If the last octet exceeds 255, it rolls over to 0 and the third octet increments by 1.\nNote: The Range parameter assigns multiple IP addresses to one virtual server. To generate an array of virtual servers, each of which owns only one IP address, use brackets in the IP Address and Name parameters to specify the range. For example:\nadd lb vserver my_vserver[1-3] HTTP 192.0.2.[1-3] 80",
			},
			"recursionavailable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set to YES, this option causes the DNS replies from this vserver to have the RA bit turned on. Typically one would set this option to YES, when the vserver is load balancing a set of DNS servers thatsupport recursive queries.",
			},
			"redirectfromport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for the virtual server, from which we absorb the traffic for http redirect",
			},
			"redirectportrewrite": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Rewrite the port and change the protocol to ensure successful HTTP redirects from services.",
			},
			"redirurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which to redirect traffic if the virtual server becomes unavailable.\nWARNING! Make sure that the domain in the URL does not match the domain specified for a content switching policy. If it does, requests are continuously redirected to the unavailable virtual server.",
			},
			"redirurlflags": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "The redirect URL to be unset.",
			},
			"resrule": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "Expression specifying which part of a server's response to use for creating rule based persistence sessions (persistence type RULE). Can be either an expression or the name of a named expression.\nExample:\nHTTP.RES.HEADER(\"setcookie\").VALUE(0).TYPECAST_NVLIST_T('=',';').VALUE(\"server1\").",
			},
			"retainconnectionsoncluster": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled.",
			},
			"rhistate": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("PASSIVE"),
				Description: "Route Health Injection (RHI) functionality of the NetSaler appliance for advertising the route of the VIP address associated with the virtual server. When Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:\n* If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address.\n* If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state.\n* If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state.",
			},
			"rtspnat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use network address translation (NAT) for RTSP data connections.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "Expression, or name of a named expression, against which traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"servicename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service to bind to the virtual server.",
			},
			"servicetype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Protocol used by the service (also called the service type).",
			},
			"sessionless": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Perform load balancing on a per-packet basis, without establishing sessions. Recommended for load balancing of intrusion detection system (IDS) servers and scenarios involving direct server return (DSR), where session information is unnecessary.",
			},
			"skippersistency": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("None"),
				Description: "This argument decides the behavior incase the service which is selected from an existing persistence session has reached threshold.",
			},
			"sobackupaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists",
			},
			"somethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of threshold that, when exceeded, triggers spillover. Available settings function as follows:\n* CONNECTION - Spillover occurs when the number of client connections exceeds the threshold.\n* DYNAMICCONNECTION - Spillover occurs when the number of client connections at the virtual server exceeds the sum of the maximum client (Max Clients) settings for bound services. Do not specify a spillover threshold for this setting, because the threshold is implied by the Max Clients settings of bound services.\n* BANDWIDTH - Spillover occurs when the bandwidth consumed by the virtual server's incoming and outgoing traffic exceeds the threshold.\n* HEALTH - Spillover occurs when the percentage of weights of the services that are UP drops below the threshold. For example, if services svc1, svc2, and svc3 are bound to a virtual server, with weights 1, 2, and 3, and the spillover threshold is 50%, spillover occurs if svc1 and svc3 or svc2 and svc3 transition to DOWN.\n* NONE - Spillover does not occur.",
			},
			"sopersistence": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "If spillover occurs, maintain source IP address based persistence for both primary and backup virtual servers.",
			},
			"sopersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2),
				Description: "Timeout for spillover persistence, in minutes.",
			},
			"sothreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold at which spillover occurs. Specify an integer for the CONNECTION spillover method, a bandwidth value in kilobits per second for the BANDWIDTH method (do not enter the units), or a percentage for the HEALTH method (do not enter the percentage symbol).",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "State of the load balancing virtual server.",
			},
			"tcpprobeport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for external TCP probe. NetScaler provides support for external TCP health check of the vserver status over the selected port. This option is only supported for vservers assigned with an IPAddress or ipset.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile whose settings are to be applied to the virtual server.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2),
				Description: "Time period for which a persistence session is in effect.",
			},
			"toggleorder": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ASCENDING"),
				Description: "Configure this option to toggle order preference",
			},
			"tosid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TOS ID of the virtual server. Applicable only when the load balancing redirection mode is set to TOS.",
			},
			"trofspersistence": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "When value is ENABLED, Trofs persistence is honored. When value is DISABLED, Trofs persistence is not honored.",
			},
			"v6netmasklen": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(128),
				Description: "Number of bits to consider in an IPv6 destination or source IP address, for creating the hash that is required by the DESTINATIONIPHASH and SOURCEIPHASH load balancing methods.",
			},
			"v6persistmasklen": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(128),
				Description: "Persistence mask for IP based persistence types, for IPv6 virtual servers.",
			},
			"vipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the inserted header. The default name is vip-header.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the specified service.",
			},
		},
	}
}

func lbvserverGetThePayloadFromtheConfig(ctx context.Context, data *LbvserverResourceModel) lb.Lbvserver {
	tflog.Debug(ctx, "In lbvserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbvserver := lb.Lbvserver{}
	if !data.Adfsproxyprofile.IsNull() {
		lbvserver.Adfsproxyprofile = data.Adfsproxyprofile.ValueString()
	}
	if !data.Apiprofile.IsNull() {
		lbvserver.Apiprofile = data.Apiprofile.ValueString()
	}
	if !data.Appflowlog.IsNull() {
		lbvserver.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Authentication.IsNull() {
		lbvserver.Authentication = data.Authentication.ValueString()
	}
	if !data.Authenticationhost.IsNull() {
		lbvserver.Authenticationhost = data.Authenticationhost.ValueString()
	}
	if !data.Authn401.IsNull() {
		lbvserver.Authn401 = data.Authn401.ValueString()
	}
	if !data.Authnprofile.IsNull() {
		lbvserver.Authnprofile = data.Authnprofile.ValueString()
	}
	if !data.Authnvsname.IsNull() {
		lbvserver.Authnvsname = data.Authnvsname.ValueString()
	}
	if !data.Backuplbmethod.IsNull() {
		lbvserver.Backuplbmethod = data.Backuplbmethod.ValueString()
	}
	if !data.Backuppersistencetimeout.IsNull() {
		lbvserver.Backuppersistencetimeout = utils.IntPtr(int(data.Backuppersistencetimeout.ValueInt64()))
	}
	if !data.Backupvserver.IsNull() {
		lbvserver.Backupvserver = data.Backupvserver.ValueString()
	}
	if !data.Bypassaaaa.IsNull() {
		lbvserver.Bypassaaaa = data.Bypassaaaa.ValueString()
	}
	if !data.Cacheable.IsNull() {
		lbvserver.Cacheable = data.Cacheable.ValueString()
	}
	if !data.Clttimeout.IsNull() {
		lbvserver.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Comment.IsNull() {
		lbvserver.Comment = data.Comment.ValueString()
	}
	if !data.Connfailover.IsNull() {
		lbvserver.Connfailover = data.Connfailover.ValueString()
	}
	if !data.Cookiename.IsNull() {
		lbvserver.Cookiename = data.Cookiename.ValueString()
	}
	if !data.Datalength.IsNull() {
		lbvserver.Datalength = utils.IntPtr(int(data.Datalength.ValueInt64()))
	}
	if !data.Dataoffset.IsNull() {
		lbvserver.Dataoffset = utils.IntPtr(int(data.Dataoffset.ValueInt64()))
	}
	if !data.Dbprofilename.IsNull() {
		lbvserver.Dbprofilename = data.Dbprofilename.ValueString()
	}
	if !data.Dbslb.IsNull() {
		lbvserver.Dbslb = data.Dbslb.ValueString()
	}
	if !data.Disableprimaryondown.IsNull() {
		lbvserver.Disableprimaryondown = data.Disableprimaryondown.ValueString()
	}
	if !data.Dns64.IsNull() {
		lbvserver.Dns64 = data.Dns64.ValueString()
	}
	if !data.Dnsoverhttps.IsNull() {
		lbvserver.Dnsoverhttps = data.Dnsoverhttps.ValueString()
	}
	if !data.Dnsprofilename.IsNull() {
		lbvserver.Dnsprofilename = data.Dnsprofilename.ValueString()
	}
	if !data.Downstateflush.IsNull() {
		lbvserver.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Hashlength.IsNull() {
		lbvserver.Hashlength = utils.IntPtr(int(data.Hashlength.ValueInt64()))
	}
	if !data.Healththreshold.IsNull() {
		lbvserver.Healththreshold = utils.IntPtr(int(data.Healththreshold.ValueInt64()))
	}
	if !data.Httpprofilename.IsNull() {
		lbvserver.Httpprofilename = data.Httpprofilename.ValueString()
	}
	if !data.Httpsredirecturl.IsNull() {
		lbvserver.Httpsredirecturl = data.Httpsredirecturl.ValueString()
	}
	if !data.Icmpvsrresponse.IsNull() {
		lbvserver.Icmpvsrresponse = data.Icmpvsrresponse.ValueString()
	}
	if !data.Insertvserveripport.IsNull() {
		lbvserver.Insertvserveripport = data.Insertvserveripport.ValueString()
	}
	if !data.Ipmask.IsNull() {
		lbvserver.Ipmask = data.Ipmask.ValueString()
	}
	if !data.Ippattern.IsNull() {
		lbvserver.Ippattern = data.Ippattern.ValueString()
	}
	if !data.Ipset.IsNull() {
		lbvserver.Ipset = data.Ipset.ValueString()
	}
	if !data.Ipv46.IsNull() {
		lbvserver.Ipv46 = data.Ipv46.ValueString()
	}
	if !data.L2conn.IsNull() {
		lbvserver.L2conn = data.L2conn.ValueString()
	}
	if !data.Lbmethod.IsNull() {
		lbvserver.Lbmethod = data.Lbmethod.ValueString()
	}
	if !data.Lbprofilename.IsNull() {
		lbvserver.Lbprofilename = data.Lbprofilename.ValueString()
	}
	if !data.Listenpolicy.IsNull() {
		lbvserver.Listenpolicy = data.Listenpolicy.ValueString()
	}
	if !data.Listenpriority.IsNull() {
		lbvserver.Listenpriority = utils.IntPtr(int(data.Listenpriority.ValueInt64()))
	}
	if !data.M.IsNull() {
		lbvserver.M = data.M.ValueString()
	}
	if !data.Macmoderetainvlan.IsNull() {
		lbvserver.Macmoderetainvlan = data.Macmoderetainvlan.ValueString()
	}
	if !data.Maxautoscalemembers.IsNull() {
		lbvserver.Maxautoscalemembers = utils.IntPtr(int(data.Maxautoscalemembers.ValueInt64()))
	}
	if !data.Minautoscalemembers.IsNull() {
		lbvserver.Minautoscalemembers = utils.IntPtr(int(data.Minautoscalemembers.ValueInt64()))
	}
	if !data.Mssqlserverversion.IsNull() {
		lbvserver.Mssqlserverversion = data.Mssqlserverversion.ValueString()
	}
	if !data.Mysqlcharacterset.IsNull() {
		lbvserver.Mysqlcharacterset = utils.IntPtr(int(data.Mysqlcharacterset.ValueInt64()))
	}
	if !data.Mysqlprotocolversion.IsNull() {
		lbvserver.Mysqlprotocolversion = utils.IntPtr(int(data.Mysqlprotocolversion.ValueInt64()))
	}
	if !data.Mysqlservercapabilities.IsNull() {
		lbvserver.Mysqlservercapabilities = utils.IntPtr(int(data.Mysqlservercapabilities.ValueInt64()))
	}
	if !data.Mysqlserverversion.IsNull() {
		lbvserver.Mysqlserverversion = data.Mysqlserverversion.ValueString()
	}
	if !data.Name.IsNull() {
		lbvserver.Name = data.Name.ValueString()
	}
	if !data.Netmask.IsNull() {
		lbvserver.Netmask = data.Netmask.ValueString()
	}
	if !data.Netprofile.IsNull() {
		lbvserver.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Newname.IsNull() {
		lbvserver.Newname = data.Newname.ValueString()
	}
	if !data.Newservicerequest.IsNull() {
		lbvserver.Newservicerequest = utils.IntPtr(int(data.Newservicerequest.ValueInt64()))
	}
	if !data.Newservicerequestincrementinterval.IsNull() {
		lbvserver.Newservicerequestincrementinterval = utils.IntPtr(int(data.Newservicerequestincrementinterval.ValueInt64()))
	}
	if !data.Newservicerequestunit.IsNull() {
		lbvserver.Newservicerequestunit = data.Newservicerequestunit.ValueString()
	}
	if !data.Oracleserverversion.IsNull() {
		lbvserver.Oracleserverversion = data.Oracleserverversion.ValueString()
	}
	if !data.Order.IsNull() {
		lbvserver.Order = utils.IntPtr(int(data.Order.ValueInt64()))
	}
	if !data.Orderthreshold.IsNull() {
		lbvserver.Orderthreshold = utils.IntPtr(int(data.Orderthreshold.ValueInt64()))
	}
	if !data.Persistencebackup.IsNull() {
		lbvserver.Persistencebackup = data.Persistencebackup.ValueString()
	}
	if !data.Persistencetype.IsNull() {
		lbvserver.Persistencetype = data.Persistencetype.ValueString()
	}
	if !data.Persistmask.IsNull() {
		lbvserver.Persistmask = data.Persistmask.ValueString()
	}
	if !data.Port.IsNull() {
		lbvserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Probeport.IsNull() {
		lbvserver.Probeport = utils.IntPtr(int(data.Probeport.ValueInt64()))
	}
	if !data.Probeprotocol.IsNull() {
		lbvserver.Probeprotocol = data.Probeprotocol.ValueString()
	}
	if !data.Probesuccessresponsecode.IsNull() {
		lbvserver.Probesuccessresponsecode = data.Probesuccessresponsecode.ValueString()
	}
	if !data.Processlocal.IsNull() {
		lbvserver.Processlocal = data.Processlocal.ValueString()
	}
	if !data.Push.IsNull() {
		lbvserver.Push = data.Push.ValueString()
	}
	if !data.Pushlabel.IsNull() {
		lbvserver.Pushlabel = data.Pushlabel.ValueString()
	}
	if !data.Pushmulticlients.IsNull() {
		lbvserver.Pushmulticlients = data.Pushmulticlients.ValueString()
	}
	if !data.Pushvserver.IsNull() {
		lbvserver.Pushvserver = data.Pushvserver.ValueString()
	}
	if !data.Quicbridgeprofilename.IsNull() {
		lbvserver.Quicbridgeprofilename = data.Quicbridgeprofilename.ValueString()
	}
	if !data.Quicprofilename.IsNull() {
		lbvserver.Quicprofilename = data.Quicprofilename.ValueString()
	}
	if !data.Range.IsNull() {
		lbvserver.Range = utils.IntPtr(int(data.Range.ValueInt64()))
	}
	if !data.Recursionavailable.IsNull() {
		lbvserver.Recursionavailable = data.Recursionavailable.ValueString()
	}
	if !data.Redirectfromport.IsNull() {
		lbvserver.Redirectfromport = utils.IntPtr(int(data.Redirectfromport.ValueInt64()))
	}
	if !data.Redirectportrewrite.IsNull() {
		lbvserver.Redirectportrewrite = data.Redirectportrewrite.ValueString()
	}
	if !data.Redirurl.IsNull() {
		lbvserver.Redirurl = data.Redirurl.ValueString()
	}
	if !data.Redirurlflags.IsNull() {
		lbvserver.Redirurlflags = data.Redirurlflags.ValueBool()
	}
	if !data.Resrule.IsNull() {
		lbvserver.Resrule = data.Resrule.ValueString()
	}
	if !data.Retainconnectionsoncluster.IsNull() {
		lbvserver.Retainconnectionsoncluster = data.Retainconnectionsoncluster.ValueString()
	}
	if !data.Rhistate.IsNull() {
		lbvserver.Rhistate = data.Rhistate.ValueString()
	}
	if !data.Rtspnat.IsNull() {
		lbvserver.Rtspnat = data.Rtspnat.ValueString()
	}
	if !data.Rule.IsNull() {
		lbvserver.Rule = data.Rule.ValueString()
	}
	if !data.Servicename.IsNull() {
		lbvserver.Servicename = data.Servicename.ValueString()
	}
	if !data.Servicetype.IsNull() {
		lbvserver.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sessionless.IsNull() {
		lbvserver.Sessionless = data.Sessionless.ValueString()
	}
	if !data.Skippersistency.IsNull() {
		lbvserver.Skippersistency = data.Skippersistency.ValueString()
	}
	if !data.Sobackupaction.IsNull() {
		lbvserver.Sobackupaction = data.Sobackupaction.ValueString()
	}
	if !data.Somethod.IsNull() {
		lbvserver.Somethod = data.Somethod.ValueString()
	}
	if !data.Sopersistence.IsNull() {
		lbvserver.Sopersistence = data.Sopersistence.ValueString()
	}
	if !data.Sopersistencetimeout.IsNull() {
		lbvserver.Sopersistencetimeout = utils.IntPtr(int(data.Sopersistencetimeout.ValueInt64()))
	}
	if !data.Sothreshold.IsNull() {
		lbvserver.Sothreshold = utils.IntPtr(int(data.Sothreshold.ValueInt64()))
	}
	if !data.State.IsNull() {
		lbvserver.State = data.State.ValueString()
	}
	if !data.Tcpprobeport.IsNull() {
		lbvserver.Tcpprobeport = utils.IntPtr(int(data.Tcpprobeport.ValueInt64()))
	}
	if !data.Tcpprofilename.IsNull() {
		lbvserver.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Td.IsNull() {
		lbvserver.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Timeout.IsNull() {
		lbvserver.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Toggleorder.IsNull() {
		lbvserver.Toggleorder = data.Toggleorder.ValueString()
	}
	if !data.Tosid.IsNull() {
		lbvserver.Tosid = utils.IntPtr(int(data.Tosid.ValueInt64()))
	}
	if !data.Trofspersistence.IsNull() {
		lbvserver.Trofspersistence = data.Trofspersistence.ValueString()
	}
	if !data.V6netmasklen.IsNull() {
		lbvserver.V6netmasklen = utils.IntPtr(int(data.V6netmasklen.ValueInt64()))
	}
	if !data.V6persistmasklen.IsNull() {
		lbvserver.V6persistmasklen = utils.IntPtr(int(data.V6persistmasklen.ValueInt64()))
	}
	if !data.Vipheader.IsNull() {
		lbvserver.Vipheader = data.Vipheader.ValueString()
	}
	if !data.Weight.IsNull() {
		lbvserver.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return lbvserver
}

func lbvserverSetAttrFromGet(ctx context.Context, data *LbvserverResourceModel, getResponseData map[string]interface{}) *LbvserverResourceModel {
	tflog.Debug(ctx, "In lbvserverSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["adfsproxyprofile"]; ok && val != nil {
		data.Adfsproxyprofile = types.StringValue(val.(string))
	} else {
		data.Adfsproxyprofile = types.StringNull()
	}
	if val, ok := getResponseData["apiprofile"]; ok && val != nil {
		data.Apiprofile = types.StringValue(val.(string))
	} else {
		data.Apiprofile = types.StringNull()
	}
	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else {
		data.Appflowlog = types.StringNull()
	}
	if val, ok := getResponseData["authentication"]; ok && val != nil {
		data.Authentication = types.StringValue(val.(string))
	} else {
		data.Authentication = types.StringNull()
	}
	if val, ok := getResponseData["authenticationhost"]; ok && val != nil {
		data.Authenticationhost = types.StringValue(val.(string))
	} else {
		data.Authenticationhost = types.StringNull()
	}
	if val, ok := getResponseData["authn401"]; ok && val != nil {
		data.Authn401 = types.StringValue(val.(string))
	} else {
		data.Authn401 = types.StringNull()
	}
	if val, ok := getResponseData["authnprofile"]; ok && val != nil {
		data.Authnprofile = types.StringValue(val.(string))
	} else {
		data.Authnprofile = types.StringNull()
	}
	if val, ok := getResponseData["authnvsname"]; ok && val != nil {
		data.Authnvsname = types.StringValue(val.(string))
	} else {
		data.Authnvsname = types.StringNull()
	}
	if val, ok := getResponseData["backuplbmethod"]; ok && val != nil {
		data.Backuplbmethod = types.StringValue(val.(string))
	} else {
		data.Backuplbmethod = types.StringNull()
	}
	if val, ok := getResponseData["backuppersistencetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Backuppersistencetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Backuppersistencetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["backupvserver"]; ok && val != nil {
		data.Backupvserver = types.StringValue(val.(string))
	} else {
		data.Backupvserver = types.StringNull()
	}
	if val, ok := getResponseData["bypassaaaa"]; ok && val != nil {
		data.Bypassaaaa = types.StringValue(val.(string))
	} else {
		data.Bypassaaaa = types.StringNull()
	}
	if val, ok := getResponseData["cacheable"]; ok && val != nil {
		data.Cacheable = types.StringValue(val.(string))
	} else {
		data.Cacheable = types.StringNull()
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
	if val, ok := getResponseData["connfailover"]; ok && val != nil {
		data.Connfailover = types.StringValue(val.(string))
	} else {
		data.Connfailover = types.StringNull()
	}
	if val, ok := getResponseData["cookiename"]; ok && val != nil {
		data.Cookiename = types.StringValue(val.(string))
	} else {
		data.Cookiename = types.StringNull()
	}
	if val, ok := getResponseData["datalength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Datalength = types.Int64Value(intVal)
		}
	} else {
		data.Datalength = types.Int64Null()
	}
	if val, ok := getResponseData["dataoffset"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Dataoffset = types.Int64Value(intVal)
		}
	} else {
		data.Dataoffset = types.Int64Null()
	}
	if val, ok := getResponseData["dbprofilename"]; ok && val != nil {
		data.Dbprofilename = types.StringValue(val.(string))
	} else {
		data.Dbprofilename = types.StringNull()
	}
	if val, ok := getResponseData["dbslb"]; ok && val != nil {
		data.Dbslb = types.StringValue(val.(string))
	} else {
		data.Dbslb = types.StringNull()
	}
	if val, ok := getResponseData["disableprimaryondown"]; ok && val != nil {
		data.Disableprimaryondown = types.StringValue(val.(string))
	} else {
		data.Disableprimaryondown = types.StringNull()
	}
	if val, ok := getResponseData["dns64"]; ok && val != nil {
		data.Dns64 = types.StringValue(val.(string))
	} else {
		data.Dns64 = types.StringNull()
	}
	if val, ok := getResponseData["dnsoverhttps"]; ok && val != nil {
		data.Dnsoverhttps = types.StringValue(val.(string))
	} else {
		data.Dnsoverhttps = types.StringNull()
	}
	if val, ok := getResponseData["dnsprofilename"]; ok && val != nil {
		data.Dnsprofilename = types.StringValue(val.(string))
	} else {
		data.Dnsprofilename = types.StringNull()
	}
	if val, ok := getResponseData["downstateflush"]; ok && val != nil {
		data.Downstateflush = types.StringValue(val.(string))
	} else {
		data.Downstateflush = types.StringNull()
	}
	if val, ok := getResponseData["hashlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Hashlength = types.Int64Value(intVal)
		}
	} else {
		data.Hashlength = types.Int64Null()
	}
	if val, ok := getResponseData["healththreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Healththreshold = types.Int64Value(intVal)
		}
	} else {
		data.Healththreshold = types.Int64Null()
	}
	if val, ok := getResponseData["httpprofilename"]; ok && val != nil {
		data.Httpprofilename = types.StringValue(val.(string))
	} else {
		data.Httpprofilename = types.StringNull()
	}
	if val, ok := getResponseData["httpsredirecturl"]; ok && val != nil {
		data.Httpsredirecturl = types.StringValue(val.(string))
	} else {
		data.Httpsredirecturl = types.StringNull()
	}
	if val, ok := getResponseData["icmpvsrresponse"]; ok && val != nil {
		data.Icmpvsrresponse = types.StringValue(val.(string))
	} else {
		data.Icmpvsrresponse = types.StringNull()
	}
	if val, ok := getResponseData["insertvserveripport"]; ok && val != nil {
		data.Insertvserveripport = types.StringValue(val.(string))
	} else {
		data.Insertvserveripport = types.StringNull()
	}
	if val, ok := getResponseData["ipmask"]; ok && val != nil {
		data.Ipmask = types.StringValue(val.(string))
	} else {
		data.Ipmask = types.StringNull()
	}
	if val, ok := getResponseData["ippattern"]; ok && val != nil {
		data.Ippattern = types.StringValue(val.(string))
	} else {
		data.Ippattern = types.StringNull()
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
	if val, ok := getResponseData["lbmethod"]; ok && val != nil {
		data.Lbmethod = types.StringValue(val.(string))
	} else {
		data.Lbmethod = types.StringNull()
	}
	if val, ok := getResponseData["lbprofilename"]; ok && val != nil {
		data.Lbprofilename = types.StringValue(val.(string))
	} else {
		data.Lbprofilename = types.StringNull()
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
	if val, ok := getResponseData["m"]; ok && val != nil {
		data.M = types.StringValue(val.(string))
	} else {
		data.M = types.StringNull()
	}
	if val, ok := getResponseData["macmoderetainvlan"]; ok && val != nil {
		data.Macmoderetainvlan = types.StringValue(val.(string))
	} else {
		data.Macmoderetainvlan = types.StringNull()
	}
	if val, ok := getResponseData["maxautoscalemembers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxautoscalemembers = types.Int64Value(intVal)
		}
	} else {
		data.Maxautoscalemembers = types.Int64Null()
	}
	if val, ok := getResponseData["minautoscalemembers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Minautoscalemembers = types.Int64Value(intVal)
		}
	} else {
		data.Minautoscalemembers = types.Int64Null()
	}
	if val, ok := getResponseData["mssqlserverversion"]; ok && val != nil {
		data.Mssqlserverversion = types.StringValue(val.(string))
	} else {
		data.Mssqlserverversion = types.StringNull()
	}
	if val, ok := getResponseData["mysqlcharacterset"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mysqlcharacterset = types.Int64Value(intVal)
		}
	} else {
		data.Mysqlcharacterset = types.Int64Null()
	}
	if val, ok := getResponseData["mysqlprotocolversion"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mysqlprotocolversion = types.Int64Value(intVal)
		}
	} else {
		data.Mysqlprotocolversion = types.Int64Null()
	}
	if val, ok := getResponseData["mysqlservercapabilities"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mysqlservercapabilities = types.Int64Value(intVal)
		}
	} else {
		data.Mysqlservercapabilities = types.Int64Null()
	}
	if val, ok := getResponseData["mysqlserverversion"]; ok && val != nil {
		data.Mysqlserverversion = types.StringValue(val.(string))
	} else {
		data.Mysqlserverversion = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
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
	if val, ok := getResponseData["newservicerequest"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Newservicerequest = types.Int64Value(intVal)
		}
	} else {
		data.Newservicerequest = types.Int64Null()
	}
	if val, ok := getResponseData["newservicerequestincrementinterval"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Newservicerequestincrementinterval = types.Int64Value(intVal)
		}
	} else {
		data.Newservicerequestincrementinterval = types.Int64Null()
	}
	if val, ok := getResponseData["newservicerequestunit"]; ok && val != nil {
		data.Newservicerequestunit = types.StringValue(val.(string))
	} else {
		data.Newservicerequestunit = types.StringNull()
	}
	if val, ok := getResponseData["oracleserverversion"]; ok && val != nil {
		data.Oracleserverversion = types.StringValue(val.(string))
	} else {
		data.Oracleserverversion = types.StringNull()
	}
	if val, ok := getResponseData["order"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Order = types.Int64Value(intVal)
		}
	} else {
		data.Order = types.Int64Null()
	}
	if val, ok := getResponseData["orderthreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Orderthreshold = types.Int64Value(intVal)
		}
	} else {
		data.Orderthreshold = types.Int64Null()
	}
	if val, ok := getResponseData["persistencebackup"]; ok && val != nil {
		data.Persistencebackup = types.StringValue(val.(string))
	} else {
		data.Persistencebackup = types.StringNull()
	}
	if val, ok := getResponseData["persistencetype"]; ok && val != nil {
		data.Persistencetype = types.StringValue(val.(string))
	} else {
		data.Persistencetype = types.StringNull()
	}
	if val, ok := getResponseData["persistmask"]; ok && val != nil {
		data.Persistmask = types.StringValue(val.(string))
	} else {
		data.Persistmask = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
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
	if val, ok := getResponseData["processlocal"]; ok && val != nil {
		data.Processlocal = types.StringValue(val.(string))
	} else {
		data.Processlocal = types.StringNull()
	}
	if val, ok := getResponseData["push"]; ok && val != nil {
		data.Push = types.StringValue(val.(string))
	} else {
		data.Push = types.StringNull()
	}
	if val, ok := getResponseData["pushlabel"]; ok && val != nil {
		data.Pushlabel = types.StringValue(val.(string))
	} else {
		data.Pushlabel = types.StringNull()
	}
	if val, ok := getResponseData["pushmulticlients"]; ok && val != nil {
		data.Pushmulticlients = types.StringValue(val.(string))
	} else {
		data.Pushmulticlients = types.StringNull()
	}
	if val, ok := getResponseData["pushvserver"]; ok && val != nil {
		data.Pushvserver = types.StringValue(val.(string))
	} else {
		data.Pushvserver = types.StringNull()
	}
	if val, ok := getResponseData["quicbridgeprofilename"]; ok && val != nil {
		data.Quicbridgeprofilename = types.StringValue(val.(string))
	} else {
		data.Quicbridgeprofilename = types.StringNull()
	}
	if val, ok := getResponseData["quicprofilename"]; ok && val != nil {
		data.Quicprofilename = types.StringValue(val.(string))
	} else {
		data.Quicprofilename = types.StringNull()
	}
	if val, ok := getResponseData["range"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Range = types.Int64Value(intVal)
		}
	} else {
		data.Range = types.Int64Null()
	}
	if val, ok := getResponseData["recursionavailable"]; ok && val != nil {
		data.Recursionavailable = types.StringValue(val.(string))
	} else {
		data.Recursionavailable = types.StringNull()
	}
	if val, ok := getResponseData["redirectfromport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Redirectfromport = types.Int64Value(intVal)
		}
	} else {
		data.Redirectfromport = types.Int64Null()
	}
	if val, ok := getResponseData["redirectportrewrite"]; ok && val != nil {
		data.Redirectportrewrite = types.StringValue(val.(string))
	} else {
		data.Redirectportrewrite = types.StringNull()
	}
	if val, ok := getResponseData["redirurl"]; ok && val != nil {
		data.Redirurl = types.StringValue(val.(string))
	} else {
		data.Redirurl = types.StringNull()
	}
	if val, ok := getResponseData["redirurlflags"]; ok && val != nil {
		data.Redirurlflags = types.BoolValue(val.(bool))
	} else {
		data.Redirurlflags = types.BoolNull()
	}
	if val, ok := getResponseData["resrule"]; ok && val != nil {
		data.Resrule = types.StringValue(val.(string))
	} else {
		data.Resrule = types.StringNull()
	}
	if val, ok := getResponseData["retainconnectionsoncluster"]; ok && val != nil {
		data.Retainconnectionsoncluster = types.StringValue(val.(string))
	} else {
		data.Retainconnectionsoncluster = types.StringNull()
	}
	if val, ok := getResponseData["rhistate"]; ok && val != nil {
		data.Rhistate = types.StringValue(val.(string))
	} else {
		data.Rhistate = types.StringNull()
	}
	if val, ok := getResponseData["rtspnat"]; ok && val != nil {
		data.Rtspnat = types.StringValue(val.(string))
	} else {
		data.Rtspnat = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
	}
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	} else {
		data.Servicetype = types.StringNull()
	}
	if val, ok := getResponseData["sessionless"]; ok && val != nil {
		data.Sessionless = types.StringValue(val.(string))
	} else {
		data.Sessionless = types.StringNull()
	}
	if val, ok := getResponseData["skippersistency"]; ok && val != nil {
		data.Skippersistency = types.StringValue(val.(string))
	} else {
		data.Skippersistency = types.StringNull()
	}
	if val, ok := getResponseData["sobackupaction"]; ok && val != nil {
		data.Sobackupaction = types.StringValue(val.(string))
	} else {
		data.Sobackupaction = types.StringNull()
	}
	if val, ok := getResponseData["somethod"]; ok && val != nil {
		data.Somethod = types.StringValue(val.(string))
	} else {
		data.Somethod = types.StringNull()
	}
	if val, ok := getResponseData["sopersistence"]; ok && val != nil {
		data.Sopersistence = types.StringValue(val.(string))
	} else {
		data.Sopersistence = types.StringNull()
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
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else {
		data.Timeout = types.Int64Null()
	}
	if val, ok := getResponseData["toggleorder"]; ok && val != nil {
		data.Toggleorder = types.StringValue(val.(string))
	} else {
		data.Toggleorder = types.StringNull()
	}
	if val, ok := getResponseData["tosid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tosid = types.Int64Value(intVal)
		}
	} else {
		data.Tosid = types.Int64Null()
	}
	if val, ok := getResponseData["trofspersistence"]; ok && val != nil {
		data.Trofspersistence = types.StringValue(val.(string))
	} else {
		data.Trofspersistence = types.StringNull()
	}
	if val, ok := getResponseData["v6netmasklen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.V6netmasklen = types.Int64Value(intVal)
		}
	} else {
		data.V6netmasklen = types.Int64Null()
	}
	if val, ok := getResponseData["v6persistmasklen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.V6persistmasklen = types.Int64Value(intVal)
		}
	} else {
		data.V6persistmasklen = types.Int64Null()
	}
	if val, ok := getResponseData["vipheader"]; ok && val != nil {
		data.Vipheader = types.StringValue(val.(string))
	} else {
		data.Vipheader = types.StringNull()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
