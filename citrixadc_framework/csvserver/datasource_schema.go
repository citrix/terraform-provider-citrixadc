package csvserver

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CsvserverDataSourceModel is the data-source-specific model, decoupled from
// CsvserverResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (ip, curstate, type, status, ...). Every non-key attribute is Computed; the
// Framework's per-attribute model <-> schema reflection requires this model to
// have exactly the attributes the data-source schema declares, which is why it
// cannot reuse the resource model.
type CsvserverDataSourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Sslcertkey               types.String `tfsdk:"sslcertkey"`
	Snisslcertkeys           types.Set    `tfsdk:"snisslcertkeys"`
	Sslprofile               types.String `tfsdk:"sslprofile"`
	Ciphers                  types.List   `tfsdk:"ciphers"`
	Ciphersuites             types.List   `tfsdk:"ciphersuites"`
	Lbvserverbinding         types.String `tfsdk:"lbvserverbinding"`
	Sslpolicybinding         types.Set    `tfsdk:"sslpolicybinding"`
	Aigwprofilename          types.String `tfsdk:"aigwprofilename"`
	Apiprofile               types.String `tfsdk:"apiprofile"`
	Appflowlog               types.String `tfsdk:"appflowlog"`
	Authentication           types.String `tfsdk:"authentication"`
	Authenticationhost       types.String `tfsdk:"authenticationhost"`
	Authn401                 types.String `tfsdk:"authn401"`
	Authnprofile             types.String `tfsdk:"authnprofile"`
	Authnvsname              types.String `tfsdk:"authnvsname"`
	Backupip                 types.String `tfsdk:"backupip"`
	Backuppersistencetimeout types.Int64  `tfsdk:"backuppersistencetimeout"`
	Backupvserver            types.String `tfsdk:"backupvserver"`
	Cacheable                types.String `tfsdk:"cacheable"`
	Casesensitive            types.String `tfsdk:"casesensitive"`
	Clttimeout               types.Int64  `tfsdk:"clttimeout"`
	Comment                  types.String `tfsdk:"comment"`
	Cookiedomain             types.String `tfsdk:"cookiedomain"`
	Cookiename               types.String `tfsdk:"cookiename"`
	Cookietimeout            types.Int64  `tfsdk:"cookietimeout"`
	Dbprofilename            types.String `tfsdk:"dbprofilename"`
	Disableprimaryondown     types.String `tfsdk:"disableprimaryondown"`
	Dnsoverhttps             types.String `tfsdk:"dnsoverhttps"`
	Dnsprofilename           types.String `tfsdk:"dnsprofilename"`
	Dnsrecordtype            types.String `tfsdk:"dnsrecordtype"`
	Domainname               types.String `tfsdk:"domainname"`
	Downstateflush           types.String `tfsdk:"downstateflush"`
	Dtls                     types.String `tfsdk:"dtls"`
	Httpprofilename          types.String `tfsdk:"httpprofilename"`
	Httpsredirecturl         types.String `tfsdk:"httpsredirecturl"`
	Icmpvsrresponse          types.String `tfsdk:"icmpvsrresponse"`
	Insertvserveripport      types.String `tfsdk:"insertvserveripport"`
	Ipmask                   types.String `tfsdk:"ipmask"`
	Ippattern                types.String `tfsdk:"ippattern"`
	Ipset                    types.String `tfsdk:"ipset"`
	Ipv46                    types.String `tfsdk:"ipv46"`
	L2conn                   types.String `tfsdk:"l2conn"`
	Listenpolicy             types.String `tfsdk:"listenpolicy"`
	Listenpriority           types.Int64  `tfsdk:"listenpriority"`
	Mcpprofilename           types.String `tfsdk:"mcpprofilename"`
	Mssqlserverversion       types.String `tfsdk:"mssqlserverversion"`
	Mysqlcharacterset        types.Int64  `tfsdk:"mysqlcharacterset"`
	Mysqlprotocolversion     types.Int64  `tfsdk:"mysqlprotocolversion"`
	Mysqlservercapabilities  types.Int64  `tfsdk:"mysqlservercapabilities"`
	Mysqlserverversion       types.String `tfsdk:"mysqlserverversion"`
	Name                     types.String `tfsdk:"name"`
	Netprofile               types.String `tfsdk:"netprofile"`
	Newname                  types.String `tfsdk:"newname"`
	Oracleserverversion      types.String `tfsdk:"oracleserverversion"`
	Persistencebackup        types.String `tfsdk:"persistencebackup"`
	Persistenceid            types.Int64  `tfsdk:"persistenceid"`
	Persistencetype          types.String `tfsdk:"persistencetype"`
	Persistmask              types.String `tfsdk:"persistmask"`
	Port                     types.Int64  `tfsdk:"port"`
	Precedence               types.String `tfsdk:"precedence"`
	Probeport                types.Int64  `tfsdk:"probeport"`
	Probeprotocol            types.String `tfsdk:"probeprotocol"`
	Probesuccessresponsecode types.String `tfsdk:"probesuccessresponsecode"`
	Push                     types.String `tfsdk:"push"`
	Pushlabel                types.String `tfsdk:"pushlabel"`
	Pushmulticlients         types.String `tfsdk:"pushmulticlients"`
	Pushvserver              types.String `tfsdk:"pushvserver"`
	Quicprofilename          types.String `tfsdk:"quicprofilename"`
	Range                    types.Int64  `tfsdk:"range"`
	Redirectfromport         types.Int64  `tfsdk:"redirectfromport"`
	Redirectportrewrite      types.String `tfsdk:"redirectportrewrite"`
	Redirecturl              types.String `tfsdk:"redirecturl"`
	Rhistate                 types.String `tfsdk:"rhistate"`
	Rtspnat                  types.String `tfsdk:"rtspnat"`
	Servicetype              types.String `tfsdk:"servicetype"`
	Sitedomainttl            types.Int64  `tfsdk:"sitedomainttl"`
	Sobackupaction           types.String `tfsdk:"sobackupaction"`
	Somethod                 types.String `tfsdk:"somethod"`
	Sopersistence            types.String `tfsdk:"sopersistence"`
	Sopersistencetimeout     types.Int64  `tfsdk:"sopersistencetimeout"`
	Sothreshold              types.Int64  `tfsdk:"sothreshold"`
	State                    types.String `tfsdk:"state"`
	Stateupdate              types.String `tfsdk:"stateupdate"`
	Targettype               types.String `tfsdk:"targettype"`
	Tcpprobeport             types.Int64  `tfsdk:"tcpprobeport"`
	Tcpprofilename           types.String `tfsdk:"tcpprofilename"`
	Td                       types.Int64  `tfsdk:"td"`
	Timeout                  types.Int64  `tfsdk:"timeout"`
	Ttl                      types.Int64  `tfsdk:"ttl"`
	V6persistmasklen         types.Int64  `tfsdk:"v6persistmasklen"`
	Vipheader                types.String `tfsdk:"vipheader"`
	Wasmmodule               types.String `tfsdk:"wasmmodule"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/csvserver.json). Never settable; populated from GET.
	Ip                        types.String `tfsdk:"ip"`
	Value                     types.String `tfsdk:"value"`
	Ngname                    types.String `tfsdk:"ngname"`
	Type                      types.String `tfsdk:"type"`
	Curstate                  types.String `tfsdk:"curstate"`
	Status                    types.Int64  `tfsdk:"status"`
	Cachetype                 types.String `tfsdk:"cachetype"`
	Redirect                  types.String `tfsdk:"redirect"`
	Homepage                  types.String `tfsdk:"homepage"`
	Dnsvservername            types.String `tfsdk:"dnsvservername"`
	Domain                    types.String `tfsdk:"domain"`
	Servicename               types.String `tfsdk:"servicename"`
	Weight                    types.Int64  `tfsdk:"weight"`
	Cachevserver              types.String `tfsdk:"cachevserver"`
	Targetvserver             types.String `tfsdk:"targetvserver"`
	Url                       types.String `tfsdk:"url"`
	Bindpoint                 types.String `tfsdk:"bindpoint"`
	Gt2gb                     types.String `tfsdk:"gt2gb"`
	Statechangetimesec        types.String `tfsdk:"statechangetimesec"`
	Statechangetimemsec       types.Int64  `tfsdk:"statechangetimemsec"`
	Tickssincelaststatechange types.Int64  `tfsdk:"tickssincelaststatechange"`
	Ruletype                  types.Int64  `tfsdk:"ruletype"`
	Lbvserver                 types.String `tfsdk:"lbvserver"`
	Targetlbvserver           types.String `tfsdk:"targetlbvserver"`
	Nodefaultbindings         types.String `tfsdk:"nodefaultbindings"`
	Version                   types.Int64  `tfsdk:"version"`
}

func CsvserverDataSourceSchema() schema.Schema {
	// The sslpolicybinding nested-block attributes are pure read-back outputs
	// (Optional+Computed). They are defined via shared attribute values so the
	// block stays consistent and the merged file's model<->schema attribute set
	// remains flat (the nested attribute names are not top-level schema keys).
	dsStr := schema.StringAttribute{Optional: true, Computed: true}
	dsBool := schema.BoolAttribute{Optional: true, Computed: true}
	dsInt := schema.Int64Attribute{Optional: true, Computed: true}

	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"aigwprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the AIGW frontend profile. For the content switching vserver to function as AI gateway, this parameter must be set. Once this parameter is set using add cs vserver, it cannot be unset. Minimum length =  1 Maximum length =  255",
			},
			"apiprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The API profile where one or more API specs are bounded to.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable logging appflow flow information",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authenticate users who request a connection to the content switching virtual server.",
			},
			"authenticationhost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "FQDN of the authentication virtual server. The service type of the virtual server should be either HTTP or SSL.",
			},
			"authn401": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable HTTP 401-response based authentication.",
			},
			"authnprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the authentication profile to be used when authentication is turned on.",
			},
			"authnvsname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of authentication virtual server that authenticates the incoming user requests to this content switching virtual server.",
			},
			"backupip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"backuppersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period for which backup persistence is in effect.",
			},
			"backupvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the backup virtual server that you are configuring. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the backup virtual server is created. You can assign a different backup virtual server or rename the existing virtual server.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks.",
			},
			"cacheable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to specify whether a virtual server, used for load balancing or content switching, routes requests to the cache redirection virtual server before sending it to the configured servers.",
			},
			"casesensitive": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Consider case in URLs (for policies that use URLs instead of RULES). For example, with the ON setting, the URLs /a/1.html and /A/1.HTML are treated differently and can have different targets (set by content switching policies). With the OFF setting, /a/1.html and /A/1.HTML are switched to the same target.",
			},
			"clttimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle time, in seconds, after which the client connection is terminated. The default values are:\n180 seconds for HTTP/SSL-based services.\n9000 seconds for other TCP-based services.\n120 seconds for DNS-based services.\n120 seconds for other UDP-based services.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Information about this virtual server.",
			},
			"cookiedomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"cookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this parameter to  specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.",
			},
			"cookietimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"dbprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DB profile.",
			},
			"disableprimaryondown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Continue forwarding the traffic to backup virtual server even after the primary server comes UP from the DOWN state.",
			},
			"dnsoverhttps": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option is used to enable/disable DNS over HTTPS (DoH) processing.",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the VServer. DNS profile properties will applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.",
			},
			"dnsrecordtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"domainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name for which to change the time to live (TTL) and/or backup service IP address.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flush all active transactions associated with a virtual server whose state transitions from UP to DOWN. Do not enable this option for applications that must complete their transactions.",
			},
			"dtls": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option starts/stops the dtls service on the vserver",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP profile containing HTTP configuration settings for the virtual server. The service type of the virtual server should be either HTTP or SSL.",
			},
			"httpsredirecturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which all HTTP traffic received on the port specified in the -redirectFromPort parameter is redirected.",
			},
			"icmpvsrresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Can be active or passive",
			},
			"insertvserveripport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert the virtual server's VIP address and port number in the request header. Available values function as follows:\n        VIPADDR - Header contains the vserver's IP address and port number without any translation.\n        OFF     - The virtual IP and port header insertion option is disabled.\n        V6TOV4MAPPING - Header contains the mapped IPv4 address corresponding to the IPv6 address of the vserver and the port number. An IPv6 address can be mapped to a user-specified IPv4 address using the set ns ip6 command.",
			},
			"ipmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP mask, in dotted decimal notation, for the IP Pattern parameter. Can have leading or trailing non-zero octets (for example, 255.255.240.0 or 0.0.255.255). Accordingly, the mask specifies whether the first n bits or the last n bits of the destination IP address in a client request are to be matched with the corresponding bits in the IP pattern. The former is called a forward mask. The latter is called a reverse mask.",
			},
			"ippattern": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server. The IP Mask parameter specifies which part of the destination IP address is matched against the pattern. Mutually exclusive with the IP Address parameter.",
			},
			"ipset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current cs vserver",
			},
			"ipv46": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the content switching virtual server.",
			},
			"l2conn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use L2 Parameters to identify a connection",
			},
			"listenpolicy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the listen policy for the content switching virtual server. Can be either the name of an existing expression or an in-line expression.",
			},
			"listenpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.",
			},
			"mcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the MCP profile to attach to this cs vserver. Enables MCP protocol processing.",
			},
			"mssqlserverversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The version of the MSSQL server",
			},
			"mysqlcharacterset": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The character set returned by the mysql vserver.",
			},
			"mysqlprotocolversion": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The protocol version returned by the mysql vserver.",
			},
			"mysqlservercapabilities": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The server capabilities returned by the mysql vserver.",
			},
			"mysqlserverversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The server version string returned by the mysql vserver.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the content switching virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\nCannot be changed after the CS virtual server is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my server or my server).",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the network profile.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"oracleserverversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Oracle server version",
			},
			"persistencebackup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Backup persistence type for the virtual server. Becomes operational if the primary persistence mechanism fails.",
			},
			"persistenceid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"persistencetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of persistence for the virtual server. Available settings function as follows:\n* SOURCEIP - Connections from the same client IP address belong to the same persistence session.\n* COOKIEINSERT - Connections that have the same HTTP Cookie, inserted by a Set-Cookie directive from a server, belong to the same persistence session.\n* SSLSESSION - Connections that have the same SSL Session ID belong to the same persistence session.",
			},
			"persistmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask for IP based persistence types, for IPv4 virtual servers.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for content switching virtual server.",
			},
			"precedence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of precedence to use for both RULE-based and URL-based policies on the content switching virtual server. With the default (RULE) setting, incoming requests are evaluated against the rule-based content switching policies. If none of the rules match, the URL in the request is evaluated against the URL-based content switching policies.",
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
			"push": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Process traffic with the push virtual server that is bound to this content switching virtual server (specified by the Push VServer parameter). The service type of the push virtual server should be either HTTP or SSL.",
			},
			"pushlabel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression for extracting the label from the response received from server. This string can be either an existing rule name or an inline expression. The service type of the virtual server should be either HTTP or SSL.",
			},
			"pushmulticlients": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow multiple Web 2.0 connections from the same client to connect to the virtual server and expect updates.",
			},
			"pushvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing virtual server, of type PUSH or SSL_PUSH, to which the server pushes updates received on the client-facing load balancing virtual server.",
			},
			"quicprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of QUIC profile which will be attached to the Content Switching VServer.",
			},
			"range": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of consecutive IP addresses, starting with the address specified by the IP Address parameter, to include in a range of addresses assigned to this virtual server.",
			},
			"redirectfromport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for the virtual server, from which we absorb the traffic for http redirect",
			},
			"redirectportrewrite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of port rewrite while performing HTTP redirect.",
			},
			"redirecturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which traffic is redirected if the virtual server becomes unavailable. The service type of the virtual server should be either HTTP or SSL.\nCaution: Make sure that the domain in the URL does not match the domain specified for a content switching policy. If it does, requests are continuously redirected to the unavailable virtual server.",
			},
			"rhistate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A host route is injected according to the setting on the virtual servers.",
			},
			"rtspnat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable network address translation (NAT) for real-time streaming protocol (RTSP) connections.",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by the virtual server.",
			},
			"sitedomainttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"sobackupaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to be performed if spillover is to take effect, but no backup chain to spillover is usable or exists",
			},
			"somethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of spillover used to divert traffic to the backup virtual server when the primary virtual server reaches the spillover threshold. Connection spillover is based on the number of connections. Bandwidth spillover is based on the total Kbps of incoming and outgoing traffic.",
			},
			"sopersistence": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Maintain source-IP based persistence on primary and backup virtual servers.",
			},
			"sopersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time-out value, in minutes, for spillover persistence.",
			},
			"sothreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Depending on the spillover method, the maximum number of connections or the maximum total bandwidth (Kbps) that a virtual server can handle before spillover occurs.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the load balancing virtual server.",
			},
			"stateupdate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable state updates for a specific content switching virtual server.",
			},
			"targettype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Virtual server target type.",
			},
			"tcpprobeport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for external TCP probe. NetScaler provides support for external TCP health check of the vserver status over the selected port. This option is only supported for vservers assigned with an IPAddress or ipset.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile containing TCP configuration settings for the virtual server.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time period for which a persistence session is in effect.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"v6persistmasklen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask for IP based persistence types, for IPv6 virtual servers.",
			},
			"vipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of virtual server IP and port header, for use with the VServer IP Port Insertion parameter.",
			},
			"wasmmodule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the WASM module to assign to this virtual server.",
			},
			"sslcertkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL certificate-key pair bound to the (SSL) content switching virtual server.",
			},
			"snisslcertkeys": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "Names of the SNI SSL certificate-key pairs bound to the (SSL) content switching virtual server.",
			},
			"sslprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the SSL profile bound to the (SSL) content switching virtual server.",
			},
			"ciphers": schema.ListAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "Cipher alias names bound to the (SSL) content switching virtual server.",
			},
			"ciphersuites": schema.ListAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "Individual cipher suite names bound to the (SSL) content switching virtual server.",
			},
			"lbvserverbinding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the default load balancing virtual server bound to the content switching virtual server.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "The IP address of the virtual server.",
			},
			"value": schema.StringAttribute{
				Computed:    true,
				Description: "The ssl card status for the transparent ssl cs vserver.",
			},
			"ngname": schema.StringAttribute{
				Computed:    true,
				Description: "Nodegroup devno to which this csvserver belongs to.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "Virtual server type. Possible values = CONTENT, ADDRESS.",
			},
			"curstate": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the cs vserver (for example UP, DOWN, OUT OF SERVICE).",
			},
			"status": schema.Int64Attribute{
				Computed:    true,
				Description: "Status.",
			},
			"cachetype": schema.StringAttribute{
				Computed:    true,
				Description: "Cache type. Possible values = TRANSPARENT, REVERSE, FORWARD.",
			},
			"redirect": schema.StringAttribute{
				Computed:    true,
				Description: "Redirect URL string. Possible values = CACHE, POLICY, ORIGIN.",
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
			"servicename": schema.StringAttribute{
				Computed:    true,
				Description: "Service name.",
			},
			"weight": schema.Int64Attribute{
				Computed:    true,
				Description: "Weight for this service.",
			},
			"cachevserver": schema.StringAttribute{
				Computed:    true,
				Description: "Cache vserver name.",
			},
			"targetvserver": schema.StringAttribute{
				Computed:    true,
				Description: "target vserver name.",
			},
			"url": schema.StringAttribute{
				Computed:    true,
				Description: "URL string.",
			},
			"bindpoint": schema.StringAttribute{
				Computed:    true,
				Description: "The bindpoint to which the policy is bound.",
			},
			"gt2gb": schema.StringAttribute{
				Computed:    true,
				Description: "This argument has no effect. Possible values = ENABLED, DISABLED.",
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
			"ruletype": schema.Int64Attribute{
				Computed:    true,
				Description: "Rule type.",
			},
			"lbvserver": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the default lb vserver bound.",
			},
			"targetlbvserver": schema.StringAttribute{
				Computed:    true,
				Description: "target vserver name.",
			},
			"nodefaultbindings": schema.StringAttribute{
				Computed:    true,
				Description: "To determine if the configuration will have default ssl CIPHER and ECC curve bindings. Possible values = YES, NO.",
			},
			"version": schema.Int64Attribute{
				Computed:    true,
				Description: "Cookie version.",
			},
		},
		Blocks: map[string]schema.Block{
			"sslpolicybinding": schema.SetNestedBlock{
				Description: "SSL policies bound to the (SSL) content switching virtual server.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"gotopriorityexpression": dsStr,
						"invoke":                 dsBool,
						"labelname":              dsStr,
						"labeltype":              dsStr,
						"policyname":             dsStr,
						"priority":               dsInt,
						"type":                   dsStr,
					},
				},
			},
		},
	}
}

// csvserverDataSourceSetAttrFromGet projects a NITRO csvserver GET response onto
// the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection. The SSL convenience attributes (sslcertkey, snisslcertkeys,
// sslprofile, ciphers, ciphersuites, lbvserverbinding, sslpolicybinding) are
// derived from separate binding reads on the resource and are not part of the
// plain csvserver GET, so they are left Null here.
func csvserverDataSourceSetAttrFromGet(ctx context.Context, data *CsvserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In csvserverDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Aigwprofilename = utils.MapGetString(g, "aigwprofilename")
	data.Apiprofile = utils.MapGetString(g, "apiprofile")
	data.Appflowlog = utils.MapGetString(g, "appflowlog")
	data.Authentication = utils.MapGetString(g, "authentication")
	data.Authenticationhost = utils.MapGetString(g, "authenticationhost")
	data.Authn401 = utils.MapGetString(g, "authn401")
	data.Authnprofile = utils.MapGetString(g, "authnprofile")
	data.Authnvsname = utils.MapGetString(g, "authnvsname")
	data.Backupip = utils.MapGetString(g, "backupip")
	data.Backuppersistencetimeout = utils.MapGetInt64(g, "backuppersistencetimeout")
	data.Backupvserver = utils.MapGetString(g, "backupvserver")
	data.Cacheable = utils.MapGetString(g, "cacheable")
	data.Casesensitive = utils.MapGetString(g, "casesensitive")
	data.Clttimeout = utils.MapGetInt64(g, "clttimeout")
	data.Comment = utils.MapGetString(g, "comment")
	data.Cookiedomain = utils.MapGetString(g, "cookiedomain")
	data.Cookiename = utils.MapGetString(g, "cookiename")
	data.Cookietimeout = utils.MapGetInt64(g, "cookietimeout")
	data.Dbprofilename = utils.MapGetString(g, "dbprofilename")
	data.Disableprimaryondown = utils.MapGetString(g, "disableprimaryondown")
	data.Dnsoverhttps = utils.MapGetString(g, "dnsoverhttps")
	data.Dnsprofilename = utils.MapGetString(g, "dnsprofilename")
	data.Dnsrecordtype = utils.MapGetString(g, "dnsrecordtype")
	data.Domainname = utils.MapGetString(g, "domainname")
	data.Downstateflush = utils.MapGetString(g, "downstateflush")
	data.Dtls = utils.MapGetString(g, "dtls")
	data.Httpprofilename = utils.MapGetString(g, "httpprofilename")
	data.Httpsredirecturl = utils.MapGetString(g, "httpsredirecturl")
	data.Icmpvsrresponse = utils.MapGetString(g, "icmpvsrresponse")
	data.Insertvserveripport = utils.MapGetString(g, "insertvserveripport")
	data.Ipmask = utils.MapGetString(g, "ipmask")
	data.Ippattern = utils.MapGetString(g, "ippattern")
	data.Ipset = utils.MapGetString(g, "ipset")
	data.Ipv46 = utils.MapGetString(g, "ipv46")
	data.L2conn = utils.MapGetString(g, "l2conn")
	data.Listenpolicy = utils.MapGetString(g, "listenpolicy")
	data.Listenpriority = utils.MapGetInt64(g, "listenpriority")
	data.Mcpprofilename = utils.MapGetString(g, "mcpprofilename")
	data.Mssqlserverversion = utils.MapGetString(g, "mssqlserverversion")
	data.Mysqlcharacterset = utils.MapGetInt64(g, "mysqlcharacterset")
	data.Mysqlprotocolversion = utils.MapGetInt64(g, "mysqlprotocolversion")
	data.Mysqlservercapabilities = utils.MapGetInt64(g, "mysqlservercapabilities")
	data.Mysqlserverversion = utils.MapGetString(g, "mysqlserverversion")
	data.Netprofile = utils.MapGetString(g, "netprofile")
	data.Oracleserverversion = utils.MapGetString(g, "oracleserverversion")
	data.Persistencebackup = utils.MapGetString(g, "persistencebackup")
	data.Persistenceid = utils.MapGetInt64(g, "persistenceid")
	data.Persistencetype = utils.MapGetString(g, "persistencetype")
	data.Persistmask = utils.MapGetString(g, "persistmask")
	data.Port = utils.MapGetInt64(g, "port")
	data.Precedence = utils.MapGetString(g, "precedence")
	data.Probeport = utils.MapGetInt64(g, "probeport")
	data.Probeprotocol = utils.MapGetString(g, "probeprotocol")
	data.Probesuccessresponsecode = utils.MapGetString(g, "probesuccessresponsecode")
	data.Push = utils.MapGetString(g, "push")
	data.Pushlabel = utils.MapGetString(g, "pushlabel")
	data.Pushmulticlients = utils.MapGetString(g, "pushmulticlients")
	data.Pushvserver = utils.MapGetString(g, "pushvserver")
	data.Quicprofilename = utils.MapGetString(g, "quicprofilename")
	data.Range = utils.MapGetInt64(g, "range")
	data.Redirectfromport = utils.MapGetInt64(g, "redirectfromport")
	data.Redirectportrewrite = utils.MapGetString(g, "redirectportrewrite")
	data.Redirecturl = utils.MapGetString(g, "redirecturl")
	data.Rhistate = utils.MapGetString(g, "rhistate")
	data.Rtspnat = utils.MapGetString(g, "rtspnat")
	data.Servicetype = utils.MapGetString(g, "servicetype")
	data.Sitedomainttl = utils.MapGetInt64(g, "sitedomainttl")
	data.Sobackupaction = utils.MapGetString(g, "sobackupaction")
	data.Somethod = utils.MapGetString(g, "somethod")
	data.Sopersistence = utils.MapGetString(g, "sopersistence")
	data.Sopersistencetimeout = utils.MapGetInt64(g, "sopersistencetimeout")
	data.Sothreshold = utils.MapGetInt64(g, "sothreshold")
	data.State = utils.MapGetString(g, "state")
	data.Stateupdate = utils.MapGetString(g, "stateupdate")
	data.Targettype = utils.MapGetString(g, "targettype")
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
	data.Ttl = utils.MapGetInt64(g, "ttl")
	data.V6persistmasklen = utils.MapGetInt64(g, "v6persistmasklen")
	data.Vipheader = utils.MapGetString(g, "vipheader")
	data.Wasmmodule = utils.MapGetString(g, "wasmmodule")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// SSL convenience attributes are derived from separate binding reads and are
	// not part of the plain csvserver GET -> Null.
	data.Sslcertkey = types.StringNull()
	data.Snisslcertkeys = types.SetNull(types.StringType)
	data.Sslprofile = types.StringNull()
	data.Ciphers = types.ListNull(types.StringType)
	data.Ciphersuites = types.ListNull(types.StringType)
	data.Lbvserverbinding = types.StringNull()
	data.Sslpolicybinding = types.SetNull(types.ObjectType{AttrTypes: sslpolicybindingAttrTypes})

	// Read-only attributes.
	data.Ip = utils.MapGetString(g, "ip")
	data.Value = utils.MapGetString(g, "value")
	data.Ngname = utils.MapGetString(g, "ngname")
	data.Type = utils.MapGetString(g, "type")
	data.Curstate = utils.MapGetString(g, "curstate")
	data.Status = utils.MapGetInt64(g, "status")
	data.Cachetype = utils.MapGetString(g, "cachetype")
	data.Redirect = utils.MapGetString(g, "redirect")
	data.Homepage = utils.MapGetString(g, "homepage")
	data.Dnsvservername = utils.MapGetString(g, "dnsvservername")
	data.Domain = utils.MapGetString(g, "domain")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Weight = utils.MapGetInt64(g, "weight")
	data.Cachevserver = utils.MapGetString(g, "cachevserver")
	data.Targetvserver = utils.MapGetString(g, "targetvserver")
	data.Url = utils.MapGetString(g, "url")
	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Gt2gb = utils.MapGetString(g, "gt2gb")
	data.Statechangetimesec = utils.MapGetString(g, "statechangetimesec")
	data.Statechangetimemsec = utils.MapGetInt64(g, "statechangetimemsec")
	data.Tickssincelaststatechange = utils.MapGetInt64(g, "tickssincelaststatechange")
	data.Ruletype = utils.MapGetInt64(g, "ruletype")
	data.Lbvserver = utils.MapGetString(g, "lbvserver")
	data.Targetlbvserver = utils.MapGetString(g, "targetlbvserver")
	data.Nodefaultbindings = utils.MapGetString(g, "nodefaultbindings")
	data.Version = utils.MapGetInt64(g, "version")
}
