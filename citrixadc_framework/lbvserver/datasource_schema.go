package lbvserver

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbvserverDataSourceModel is the data-source-specific model, decoupled from
// LbvserverResourceModel. A data source is a pure read surface, so it can expose
// the full GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes that the resource deliberately omits.
type LbvserverDataSourceModel struct {
	Id                                 types.String `tfsdk:"id"`
	Sslcertkey                         types.String `tfsdk:"sslcertkey"`
	Snisslcertkeys                     types.Set    `tfsdk:"snisslcertkeys"`
	Sslprofile                         types.String `tfsdk:"sslprofile"`
	Ciphers                            types.List   `tfsdk:"ciphers"`
	Ciphersuites                       types.List   `tfsdk:"ciphersuites"`
	Sslpolicybinding                   types.Set    `tfsdk:"sslpolicybinding"`
	Adfsproxyprofile                   types.String `tfsdk:"adfsproxyprofile"`
	Aigwprofilename                    types.String `tfsdk:"aigwprofilename"`
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
	Mcpprofilename                     types.String `tfsdk:"mcpprofilename"`
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
	Wasmmodule                         types.String `tfsdk:"wasmmodule"`
	Weight                             types.Int64  `tfsdk:"weight"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/lbvserver.json). Never settable; populated from GET.
	Value                     types.String `tfsdk:"value"`
	Ipmapping                 types.String `tfsdk:"ipmapping"`
	Ngname                    types.String `tfsdk:"ngname"`
	Type                      types.String `tfsdk:"type"`
	Curstate                  types.String `tfsdk:"curstate"`
	Effectivestate            types.String `tfsdk:"effectivestate"`
	Status                    types.Int64  `tfsdk:"status"`
	Lbrrreason                types.Int64  `tfsdk:"lbrrreason"`
	Redirect                  types.String `tfsdk:"redirect"`
	Precedence                types.String `tfsdk:"precedence"`
	Homepage                  types.String `tfsdk:"homepage"`
	Dnsvservername            types.String `tfsdk:"dnsvservername"`
	Domain                    types.String `tfsdk:"domain"`
	Cachevserver              types.String `tfsdk:"cachevserver"`
	Health                    types.Int64  `tfsdk:"health"`
	Ruletype                  types.Int64  `tfsdk:"ruletype"`
	Groupname                 types.String `tfsdk:"groupname"`
	Cookiedomain              types.String `tfsdk:"cookiedomain"`
	Map                       types.String `tfsdk:"map"`
	Gt2gb                     types.String `tfsdk:"gt2gb"`
	Consolidatedlconn         types.String `tfsdk:"consolidatedlconn"`
	Consolidatedlconngbl      types.String `tfsdk:"consolidatedlconngbl"`
	Thresholdvalue            types.Int64  `tfsdk:"thresholdvalue"`
	Bindpoint                 types.String `tfsdk:"bindpoint"`
	Version                   types.Int64  `tfsdk:"version"`
	Totalservices             types.Int64  `tfsdk:"totalservices"`
	Activeservices            types.Int64  `tfsdk:"activeservices"`
	Statechangetimesec        types.String `tfsdk:"statechangetimesec"`
	Statechangetimeseconds    types.Int64  `tfsdk:"statechangetimeseconds"`
	Statechangetimemsec       types.Int64  `tfsdk:"statechangetimemsec"`
	Tickssincelaststatechange types.Int64  `tfsdk:"tickssincelaststatechange"`
	Isgslb                    types.Bool   `tfsdk:"isgslb"`
	Vsvrdynconnsothreshold    types.Int64  `tfsdk:"vsvrdynconnsothreshold"`
	Backupvserverstatus       types.String `tfsdk:"backupvserverstatus"`
	Nodefaultbindings         types.String `tfsdk:"nodefaultbindings"`
	Currentactiveorder        types.String `tfsdk:"currentactiveorder"`
}

func LbvserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"adfsproxyprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the adfsProxy profile to be used to support ADFSPIP protocol for ADFS servers.",
			},
			"aigwprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the AIGW frontend profile. For the LB vserver to function as AI gateway, this parameter must be set. Once this parameter is set using add lb vserver, it cannot be unset.",
			},
			"apiprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The API profile where one or more API specs are bounded to.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "Backup load balancing method. Becomes operational if the primary load balancing me\nthod fails or cannot be used.\n                       Valid only if the primary method is based on static proximity.",
			},
			"backuppersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
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
				Computed:    true,
				Description: "Enable database specific load balancing for MySQL and MSSQL service types.",
			},
			"disableprimaryondown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the primary virtual server goes down, do not allow it to return to primary status until manually enabled.",
			},
			"dns64": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This argument is for enabling/disabling the dns64 on lbvserver",
			},
			"dnsoverhttps": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is used to enable/disable DNS over HTTPS (DoH) processing.",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the VServer. DNS profile properties will be applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"hashlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "How the Citrix ADC responds to ping requests received for an IP address that is common to one or more virtual servers.",
			},
			"insertvserveripport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert an HTTP header, whose value is the IP address and port number of the virtual server, before forwarding a request to the server.",
			},
			"ipmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP mask, in dotted decimal notation, for the IP Pattern parameter.",
			},
			"ippattern": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server.",
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
				Description: "Use Layer 2 parameters (channel number, MAC address, and VLAN ID) in addition to the 4-tuple to identify a connection.",
			},
			"lbmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Load balancing method.",
			},
			"lbprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LB profile which is associated to the vserver",
			},
			"listenpolicy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression identifying traffic accepted by the virtual server.",
			},
			"listenpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the listen policy. A higher number specifies a lower priority.",
			},
			"m": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Redirection mode for load balancing.",
			},
			"macmoderetainvlan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is used to retain vlan information of incoming packet when macmode is enabled",
			},
			"maxautoscalemembers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of members expected to be present when vserver is used in Autoscale.",
			},
			"mcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the MCP profile to attach to this lb vserver. Enables MCP protocol processing.",
			},
			"minautoscalemembers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum number of members expected to be present when vserver is used in Autoscale.",
			},
			"mssqlserverversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "For a load balancing virtual server of type MSSQL, the Microsoft SQL Server version.",
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
				Description: "Name for the virtual server.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 subnet mask to apply to the destination IP address or source IP address when the load balancing method is DESTINATIONIPHASH or SOURCEIPHASH.",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the network profile to associate with the virtual server.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the virtual server.",
			},
			"newservicerequest": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of requests, or percentage of the load on existing services, by which to increase the load on a new service at each interval in slow-start mode.",
			},
			"newservicerequestincrementinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval, in seconds, between successive increments in the load on a new service or a service whose state has just changed from DOWN to UP.",
			},
			"newservicerequestunit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Units in which to increment load at each interval in slow-start mode.",
			},
			"oracleserverversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Description: "Persist AVP number for Diameter Persistency.",
			},
			"persistencebackup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Backup persistence type for the virtual server. Becomes operational if the primary persistence mechanism fails.",
			},
			"persistencetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of persistence for the virtual server.",
			},
			"persistmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask for IP based persistence types, for IPv4 virtual servers.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "HTTP code to return in SUCCESS case.",
			},
			"processlocal": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By turning on this option packets destined to a vserver in a cluster will not under go any steering.",
			},
			"push": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Process traffic with the push virtual server that is bound to this load balancing virtual server.",
			},
			"pushlabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression for extracting a label from the server's response.",
			},
			"pushmulticlients": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates.",
			},
			"pushvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates.",
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
				Optional:    true,
				Computed:    true,
				Description: "Number of IP addresses that the appliance must generate and assign to the virtual server.",
			},
			"recursionavailable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set to YES, this option causes the DNS replies from this vserver to have the RA bit turned on.",
			},
			"redirectfromport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for the virtual server, from which we absorb the traffic for http redirect",
			},
			"redirectportrewrite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rewrite the port and change the protocol to ensure successful HTTP redirects from services.",
			},
			"redirurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which to redirect traffic if the virtual server becomes unavailable.",
			},
			"redirurlflags": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The redirect URL to be unset.",
			},
			"resrule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying which part of a server's response to use for creating rule based persistence sessions (persistence type RULE).",
			},
			"retainconnectionsoncluster": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout.",
			},
			"rhistate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Route Health Injection (RHI) functionality of the NetScaler appliance for advertising the route of the VIP address associated with the virtual server.",
			},
			"rtspnat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use network address translation (NAT) for RTSP data connections.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, against which traffic is evaluated.",
			},
			"servicename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service to bind to the virtual server.",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by the service (also called the service type).",
			},
			"sessionless": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform load balancing on a per-packet basis, without establishing sessions.",
			},
			"skippersistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Description: "Type of threshold that, when exceeded, triggers spillover.",
			},
			"sopersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If spillover occurs, maintain source IP address based persistence for both primary and backup virtual servers.",
			},
			"sopersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout for spillover persistence, in minutes.",
			},
			"sothreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold at which spillover occurs.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the load balancing virtual server.",
			},
			"tcpprobeport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for external TCP probe.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile whose settings are to be applied to the virtual server.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period for which a persistence session is in effect.",
			},
			"toggleorder": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure this option to toggle order preference",
			},
			"tosid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TOS ID of the virtual server. Applicable only when the load balancing redirection mode is set to TOS.",
			},
			"trofspersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When value is ENABLED, Trofs persistence is honored. When value is DISABLED, Trofs persistence is not honored.",
			},
			"v6netmasklen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of bits to consider in an IPv6 destination or source IP address, for creating the hash that is required by the DESTINATIONIPHASH and SOURCEIPHASH load balancing methods.",
			},
			"v6persistmasklen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask for IP based persistence types, for IPv6 virtual servers.",
			},
			"vipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the inserted header. The default name is vip-header.",
			},
			"wasmmodule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the WASM module to assign to this virtual server.",
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight to assign to the specified service.",
			},
			"sslcertkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL certificate-key pair bound to the (SSL) load balancing virtual server.",
			},
			"snisslcertkeys": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "Names of the SNI SSL certificate-key pairs bound to the (SSL) load balancing virtual server.",
			},
			"sslprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL profile bound to the (SSL) load balancing virtual server.",
			},
			"ciphers": schema.ListAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "Cipher alias names bound to the (SSL) load balancing virtual server.",
			},
			"ciphersuites": schema.ListAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "Individual cipher suite names bound to the (SSL) load balancing virtual server.",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"value": schema.StringAttribute{
				Computed:    true,
				Description: "SSL status. Possible values = Certkey/Certkeybundle/Vault not bound/Cert-store not usable, SSL feature disabled.",
			},
			"ipmapping": schema.StringAttribute{
				Computed:    true,
				Description: "The permanent mapping for the V6 Address.",
			},
			"ngname": schema.StringAttribute{
				Computed:    true,
				Description: "Nodegroup name to which this lbvsever belongs to.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "Type of LB vserver. Possible values = CONTENT, ADDRESS.",
			},
			"curstate": schema.StringAttribute{
				Computed:    true,
				Description: "Current LB vserver state.",
			},
			"effectivestate": schema.StringAttribute{
				Computed:    true,
				Description: "Effective state of the LB vserver, based on the state of backup vservers.",
			},
			"status": schema.Int64Attribute{
				Computed:    true,
				Description: "Current status of the lb vserver.",
			},
			"lbrrreason": schema.Int64Attribute{
				Computed:    true,
				Description: "Reason why a vserver is in RR.",
			},
			"redirect": schema.StringAttribute{
				Computed:    true,
				Description: "Cache redirect type. Possible values = CACHE, POLICY, ORIGIN.",
			},
			"precedence": schema.StringAttribute{
				Computed:    true,
				Description: "Precedence. Possible values = RULE, URL.",
			},
			"homepage": schema.StringAttribute{
				Computed:    true,
				Description: "Home page.",
			},
			"dnsvservername": schema.StringAttribute{
				Computed:    true,
				Description: "DNS vserver name.",
			},
			"domain": schema.StringAttribute{
				Computed:    true,
				Description: "Domain.",
			},
			"cachevserver": schema.StringAttribute{
				Computed:    true,
				Description: "Cache virtual server.",
			},
			"health": schema.Int64Attribute{
				Computed:    true,
				Description: "Health of vserver based on percentage of weights of active svcs/all svcs.",
			},
			"ruletype": schema.Int64Attribute{
				Computed:    true,
				Description: "Rule type.",
			},
			"groupname": schema.StringAttribute{
				Computed:    true,
				Description: "LB group to which the lb vserver is to be bound.",
			},
			"cookiedomain": schema.StringAttribute{
				Computed:    true,
				Description: "Domain name to be used in the set cookie header in case of cookie persistence.",
			},
			"map": schema.StringAttribute{
				Computed:    true,
				Description: "Map. Possible values = ON, OFF.",
			},
			"gt2gb": schema.StringAttribute{
				Computed:    true,
				Description: "Allow for greater than 2 GB transactions on this vserver.",
			},
			"consolidatedlconn": schema.StringAttribute{
				Computed:    true,
				Description: "Use consolidated stats for LeastConnection.",
			},
			"consolidatedlconngbl": schema.StringAttribute{
				Computed:    true,
				Description: "Fetches Global setting.",
			},
			"thresholdvalue": schema.Int64Attribute{
				Computed:    true,
				Description: "Tells whether threshold exceeded for this service participating in CUSTOMLB.",
			},
			"bindpoint": schema.StringAttribute{
				Computed:    true,
				Description: "The bindpoint to which the policy is bound.",
			},
			"version": schema.Int64Attribute{
				Computed:    true,
				Description: "Cookie version.",
			},
			"totalservices": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of services bound to the vserver.",
			},
			"activeservices": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of active services bound to the vserver.",
			},
			"statechangetimesec": schema.StringAttribute{
				Computed:    true,
				Description: "Time when last state change happened. Seconds part.",
			},
			"statechangetimeseconds": schema.Int64Attribute{
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
			"isgslb": schema.BoolAttribute{
				Computed:    true,
				Description: "This field is set to true if it is a GSLBVserver.",
			},
			"vsvrdynconnsothreshold": schema.Int64Attribute{
				Computed:    true,
				Description: "Spillover threshold for dynamic connection.",
			},
			"backupvserverstatus": schema.StringAttribute{
				Computed:    true,
				Description: "Status of BackUp Vserver.",
			},
			"nodefaultbindings": schema.StringAttribute{
				Computed:    true,
				Description: "To determine if the configuration will have default ssl CIPHER and ECC curve bindings.",
			},
			"currentactiveorder": schema.StringAttribute{
				Computed:    true,
				Description: "Current order that takes the traffic in case service or servicegroup is bound with order.",
			},
		},
		Blocks: map[string]schema.Block{
			"sslpolicybinding": schema.SetNestedBlock{
				Description: "SSL policies bound to the (SSL) load balancing virtual server.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"gotopriorityexpression": schema.StringAttribute{Optional: true, Computed: true},
						"invoke":                 schema.BoolAttribute{Optional: true, Computed: true},
						"labelname":              schema.StringAttribute{Optional: true, Computed: true},
						"labeltype":              schema.StringAttribute{Optional: true, Computed: true},
						"policyname":             schema.StringAttribute{Optional: true, Computed: true},
						"priority":               schema.Int64Attribute{Optional: true, Computed: true},
						"type":                   schema.StringAttribute{Optional: true, Computed: true},
					},
				},
			},
		},
	}
}

// lbvserverDataSourceSetAttrFromGet projects a NITRO lbvserver GET response onto
// the data-source model. The shared utils.MapGet* helpers fill each attribute
// from the GET (or leave it Null when the GET omits it). Binding-sourced and
// action-only inputs the base GET never returns are set to their typed Null.
func lbvserverDataSourceSetAttrFromGet(ctx context.Context, data *LbvserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In lbvserverDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Adfsproxyprofile = utils.MapGetString(g, "adfsproxyprofile")
	data.Aigwprofilename = utils.MapGetString(g, "aigwprofilename")
	data.Apiprofile = utils.MapGetString(g, "apiprofile")
	data.Appflowlog = utils.MapGetString(g, "appflowlog")
	data.Authentication = utils.MapGetString(g, "authentication")
	data.Authenticationhost = utils.MapGetString(g, "authenticationhost")
	data.Authn401 = utils.MapGetString(g, "authn401")
	data.Authnprofile = utils.MapGetString(g, "authnprofile")
	data.Authnvsname = utils.MapGetString(g, "authnvsname")
	data.Backuplbmethod = utils.MapGetString(g, "backuplbmethod")
	data.Backuppersistencetimeout = utils.MapGetInt64(g, "backuppersistencetimeout")
	data.Backupvserver = utils.MapGetString(g, "backupvserver")
	data.Bypassaaaa = utils.MapGetString(g, "bypassaaaa")
	data.Cacheable = utils.MapGetString(g, "cacheable")
	data.Clttimeout = utils.MapGetInt64(g, "clttimeout")
	data.Comment = utils.MapGetString(g, "comment")
	data.Connfailover = utils.MapGetString(g, "connfailover")
	data.Cookiename = utils.MapGetString(g, "cookiename")
	data.Datalength = utils.MapGetInt64(g, "datalength")
	data.Dataoffset = utils.MapGetInt64(g, "dataoffset")
	data.Dbprofilename = utils.MapGetString(g, "dbprofilename")
	data.Dbslb = utils.MapGetString(g, "dbslb")
	data.Disableprimaryondown = utils.MapGetString(g, "disableprimaryondown")
	data.Dns64 = utils.MapGetString(g, "dns64")
	data.Dnsoverhttps = utils.MapGetString(g, "dnsoverhttps")
	data.Dnsprofilename = utils.MapGetString(g, "dnsprofilename")
	data.Downstateflush = utils.MapGetString(g, "downstateflush")
	data.Hashlength = utils.MapGetInt64(g, "hashlength")
	data.Healththreshold = utils.MapGetInt64(g, "healththreshold")
	data.Httpprofilename = utils.MapGetString(g, "httpprofilename")
	data.Httpsredirecturl = utils.MapGetString(g, "httpsredirecturl")
	data.Icmpvsrresponse = utils.MapGetString(g, "icmpvsrresponse")
	data.Insertvserveripport = utils.MapGetString(g, "insertvserveripport")
	data.Ipmask = utils.MapGetString(g, "ipmask")
	data.Ippattern = utils.MapGetString(g, "ippattern")
	data.Ipset = utils.MapGetString(g, "ipset")
	data.Ipv46 = utils.MapGetString(g, "ipv46")
	data.L2conn = utils.MapGetString(g, "l2conn")
	data.Lbmethod = utils.MapGetString(g, "lbmethod")
	data.Lbprofilename = utils.MapGetString(g, "lbprofilename")
	data.Listenpolicy = utils.MapGetString(g, "listenpolicy")
	data.Listenpriority = utils.MapGetInt64(g, "listenpriority")
	data.M = utils.MapGetString(g, "m")
	data.Macmoderetainvlan = utils.MapGetString(g, "macmoderetainvlan")
	data.Maxautoscalemembers = utils.MapGetInt64(g, "maxautoscalemembers")
	data.Mcpprofilename = utils.MapGetString(g, "mcpprofilename")
	data.Minautoscalemembers = utils.MapGetInt64(g, "minautoscalemembers")
	data.Mssqlserverversion = utils.MapGetString(g, "mssqlserverversion")
	data.Mysqlcharacterset = utils.MapGetInt64(g, "mysqlcharacterset")
	data.Mysqlprotocolversion = utils.MapGetInt64(g, "mysqlprotocolversion")
	data.Mysqlservercapabilities = utils.MapGetInt64(g, "mysqlservercapabilities")
	data.Mysqlserverversion = utils.MapGetString(g, "mysqlserverversion")
	data.Netmask = utils.MapGetString(g, "netmask")
	data.Netprofile = utils.MapGetString(g, "netprofile")
	data.Newservicerequest = utils.MapGetInt64(g, "newservicerequest")
	data.Newservicerequestincrementinterval = utils.MapGetInt64(g, "newservicerequestincrementinterval")
	data.Newservicerequestunit = utils.MapGetString(g, "newservicerequestunit")
	data.Oracleserverversion = utils.MapGetString(g, "oracleserverversion")
	data.Order = utils.MapGetInt64(g, "order")
	data.Orderthreshold = utils.MapGetInt64(g, "orderthreshold")
	data.Persistencebackup = utils.MapGetString(g, "persistencebackup")
	data.Persistencetype = utils.MapGetString(g, "persistencetype")
	data.Persistmask = utils.MapGetString(g, "persistmask")
	data.Port = utils.MapGetInt64(g, "port")
	data.Probeport = utils.MapGetInt64(g, "probeport")
	data.Probeprotocol = utils.MapGetString(g, "probeprotocol")
	data.Probesuccessresponsecode = utils.MapGetString(g, "probesuccessresponsecode")
	data.Processlocal = utils.MapGetString(g, "processlocal")
	data.Push = utils.MapGetString(g, "push")
	data.Pushlabel = utils.MapGetString(g, "pushlabel")
	data.Pushmulticlients = utils.MapGetString(g, "pushmulticlients")
	data.Pushvserver = utils.MapGetString(g, "pushvserver")
	data.Quicbridgeprofilename = utils.MapGetString(g, "quicbridgeprofilename")
	data.Quicprofilename = utils.MapGetString(g, "quicprofilename")
	data.Range = utils.MapGetInt64(g, "range")
	data.Recursionavailable = utils.MapGetString(g, "recursionavailable")
	data.Redirectfromport = utils.MapGetInt64(g, "redirectfromport")
	data.Redirectportrewrite = utils.MapGetString(g, "redirectportrewrite")
	data.Redirurl = utils.MapGetString(g, "redirurl")
	data.Resrule = utils.MapGetString(g, "resrule")
	data.Retainconnectionsoncluster = utils.MapGetString(g, "retainconnectionsoncluster")
	data.Rhistate = utils.MapGetString(g, "rhistate")
	data.Rtspnat = utils.MapGetString(g, "rtspnat")
	data.Rule = utils.MapGetString(g, "rule")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Servicetype = utils.MapGetString(g, "servicetype")
	data.Sessionless = utils.MapGetString(g, "sessionless")
	data.Skippersistency = utils.MapGetString(g, "skippersistency")
	data.Sobackupaction = utils.MapGetString(g, "sobackupaction")
	data.Somethod = utils.MapGetString(g, "somethod")
	data.Sopersistence = utils.MapGetString(g, "sopersistence")
	data.Sopersistencetimeout = utils.MapGetInt64(g, "sopersistencetimeout")
	data.Sothreshold = utils.MapGetInt64(g, "sothreshold")
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
	data.Timeout = utils.MapGetInt64(g, "timeout")
	data.Toggleorder = utils.MapGetString(g, "toggleorder")
	data.Tosid = utils.MapGetInt64(g, "tosid")
	data.Trofspersistence = utils.MapGetString(g, "trofspersistence")
	data.V6netmasklen = utils.MapGetInt64(g, "v6netmasklen")
	data.V6persistmasklen = utils.MapGetInt64(g, "v6persistmasklen")
	data.Vipheader = utils.MapGetString(g, "vipheader")
	data.Wasmmodule = utils.MapGetString(g, "wasmmodule")
	data.Weight = utils.MapGetInt64(g, "weight")

	// persistavpno is an Int64-typed list; MapGetStringList would yield the wrong
	// element type, so convert inline.
	if v, ok := g["persistavpno"]; ok && v != nil {
		if rawList, lok := v.([]interface{}); lok {
			intList := make([]int64, 0, len(rawList))
			for _, e := range rawList {
				if iv, err := utils.ConvertToInt64(e); err == nil {
					intList = append(intList, iv)
				}
			}
			if lv, d := types.ListValueFrom(ctx, types.Int64Type, intList); !d.HasError() {
				data.Persistavpno = lv
			} else {
				data.Persistavpno = types.ListNull(types.Int64Type)
			}
		} else {
			data.Persistavpno = types.ListNull(types.Int64Type)
		}
	} else {
		data.Persistavpno = types.ListNull(types.Int64Type)
	}

	// newname / redirurlflags are action-only inputs the GET never returns -> Null.
	data.Newname = types.StringNull()
	data.Redirurlflags = types.BoolNull()

	// SSL bindings are sourced from separate binding endpoints, not the base
	// lbvserver GET, so they are Null on this data source projection.
	data.Sslcertkey = types.StringNull()
	data.Sslprofile = types.StringNull()
	data.Snisslcertkeys = types.SetNull(types.StringType)
	data.Ciphers = types.ListNull(types.StringType)
	data.Ciphersuites = types.ListNull(types.StringType)
	data.Sslpolicybinding = types.SetNull(types.ObjectType{AttrTypes: sslpolicybindingAttrTypes})

	// Read-only attributes.
	data.Value = utils.MapGetString(g, "value")
	data.Ipmapping = utils.MapGetString(g, "ipmapping")
	data.Ngname = utils.MapGetString(g, "ngname")
	data.Type = utils.MapGetString(g, "type")
	data.Curstate = utils.MapGetString(g, "curstate")
	data.Effectivestate = utils.MapGetString(g, "effectivestate")
	data.Status = utils.MapGetInt64(g, "status")
	data.Lbrrreason = utils.MapGetInt64(g, "lbrrreason")
	data.Redirect = utils.MapGetString(g, "redirect")
	data.Precedence = utils.MapGetString(g, "precedence")
	data.Homepage = utils.MapGetString(g, "homepage")
	data.Dnsvservername = utils.MapGetString(g, "dnsvservername")
	data.Domain = utils.MapGetString(g, "domain")
	data.Cachevserver = utils.MapGetString(g, "cachevserver")
	data.Health = utils.MapGetInt64(g, "health")
	data.Ruletype = utils.MapGetInt64(g, "ruletype")
	data.Groupname = utils.MapGetString(g, "groupname")
	data.Cookiedomain = utils.MapGetString(g, "cookiedomain")
	data.Map = utils.MapGetString(g, "map")
	data.Gt2gb = utils.MapGetString(g, "gt2gb")
	data.Consolidatedlconn = utils.MapGetString(g, "consolidatedlconn")
	data.Consolidatedlconngbl = utils.MapGetString(g, "consolidatedlconngbl")
	data.Thresholdvalue = utils.MapGetInt64(g, "thresholdvalue")
	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Version = utils.MapGetInt64(g, "version")
	data.Totalservices = utils.MapGetInt64(g, "totalservices")
	data.Activeservices = utils.MapGetInt64(g, "activeservices")
	data.Statechangetimesec = utils.MapGetString(g, "statechangetimesec")
	data.Statechangetimeseconds = utils.MapGetInt64(g, "statechangetimeseconds")
	data.Statechangetimemsec = utils.MapGetInt64(g, "statechangetimemsec")
	data.Tickssincelaststatechange = utils.MapGetInt64(g, "tickssincelaststatechange")
	data.Isgslb = utils.MapGetBool(g, "isgslb")
	data.Vsvrdynconnsothreshold = utils.MapGetInt64(g, "vsvrdynconnsothreshold")
	data.Backupvserverstatus = utils.MapGetString(g, "backupvserverstatus")
	data.Nodefaultbindings = utils.MapGetString(g, "nodefaultbindings")
	data.Currentactiveorder = utils.MapGetString(g, "currentactiveorder")
}
