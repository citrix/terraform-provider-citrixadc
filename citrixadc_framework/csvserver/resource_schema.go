package csvserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

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

// CsvserverResourceModel describes the resource data model.
type CsvserverResourceModel struct {
	Id                       types.String `tfsdk:"id"`
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
}

func (r *CsvserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the csvserver resource.",
			},
			"apiprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The API profile where one or more API specs are bounded to.",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
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
				Default:     int64default.StaticInt64(2),
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
				Default:     stringdefault.StaticString("True"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Continue forwarding the traffic to backup virtual server even after the primary server comes UP from the DOWN state.",
			},
			"dnsoverhttps": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "This option is used to enable/disable DNS over HTTPS (DoH) processing.",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the VServer. DNS profile properties will applied to the transactions processed by a VServer. This parameter is valid only for DNS and DNS-TCP VServers.",
			},
			"dnsrecordtype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NSGSLB_IPV4"),
				Description: "0",
			},
			"domainname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain name for which to change the time to live (TTL) and/or backup service IP address.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
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
				Default:     stringdefault.StaticString("PASSIVE"),
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
				Description: "IP address pattern, in dotted decimal notation, for identifying packets to be accepted by the virtual server. The IP Mask parameter specifies which part of the destination IP address is matched against the pattern. Mutually exclusive with the IP Address parameter.\nFor example, if the IP pattern assigned to the virtual server is 198.51.100.0 and the IP mask is 255.255.240.0 (a forward mask), the first 20 bits in the destination IP addresses are matched with the first 20 bits in the pattern. The virtual server accepts requests with IP addresses that range from 198.51.96.1 to 198.51.111.254. You can also use a pattern such as 0.0.2.2 and a mask such as 0.0.255.255 (a reverse mask).\nIf a destination IP address matches more than one IP pattern, the pattern with the longest match is selected, and the associated virtual server processes the request. For example, if the virtual servers, vs1 and vs2, have the same IP pattern, 0.0.100.128, but different IP masks of 0.0.255.255 and 0.0.224.255, a destination IP address of 198.51.100.128 has the longest match with the IP pattern of vs1. If a destination IP address matches two or more virtual servers to the same extent, the request is processed by the virtual server whose port number matches the port number in the request.",
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
				Default:     stringdefault.StaticString("NONE"),
				Description: "String specifying the listen policy for the content switching virtual server. Can be either the name of an existing expression or an in-line expression.",
			},
			"listenpriority": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(101),
				Description: "Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.",
			},
			"mssqlserverversion": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("2008R2"),
				Description: "The version of the MSSQL server",
			},
			"mysqlcharacterset": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(8),
				Description: "The character set returned by the mysql vserver.",
			},
			"mysqlprotocolversion": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(10),
				Description: "The protocol version returned by the mysql vserver.",
			},
			"mysqlservercapabilities": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(41613),
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"oracleserverversion": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("10G"),
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Port number for content switching virtual server.",
			},
			"precedence": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("RULE"),
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
				Default:     stringdefault.StaticString("200 OK"),
				Description: "HTTP code to return in SUCCESS case.",
			},
			"push": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Process traffic with the push virtual server that is bound to this content switching virtual server (specified by the Push VServer parameter). The service type of the push virtual server should be either HTTP or SSL.",
			},
			"pushlabel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("none"),
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
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Number of consecutive IP addresses, starting with the address specified by the IP Address parameter, to include in a range of addresses assigned to this virtual server.",
			},
			"redirectfromport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for the virtual server, from which we absorb the traffic for http redirect",
			},
			"redirectportrewrite": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "State of port rewrite while performing HTTP redirect.",
			},
			"redirecturl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL to which traffic is redirected if the virtual server becomes unavailable. The service type of the virtual server should be either HTTP or SSL.\nCaution: Make sure that the domain in the URL does not match the domain specified for a content switching policy. If it does, requests are continuously redirected to the unavailable virtual server.",
			},
			"rhistate": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("PASSIVE"),
				Description: "A host route is injected according to the setting on the virtual servers\n            * If set to PASSIVE on all the virtual servers that share the IP address, the appliance always injects the hostroute.\n            * If set to ACTIVE on all the virtual servers that share the IP address, the appliance injects even if one virtual server is UP.\n            * If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance, injects even if one virtual server set to ACTIVE is UP.",
			},
			"rtspnat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable network address translation (NAT) for real-time streaming protocol (RTSP) connections.",
			},
			"servicetype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Maintain source-IP based persistence on primary and backup virtual servers.",
			},
			"sopersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2),
				Description: "Time-out value, in minutes, for spillover persistence.",
			},
			"sothreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Depending on the spillover method, the maximum number of connections or the maximum total bandwidth (Kbps) that a virtual server can handle before spillover occurs.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Initial state of the load balancing virtual server.",
			},
			"stateupdate": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable state updates for a specific content switching virtual server. By default, the Content Switching virtual server is always UP, regardless of the state of the Load Balancing virtual servers bound to it. This parameter interacts with the global setting as follows:\nGlobal Level | Vserver Level | Result\nENABLED      ENABLED        ENABLED\nENABLED      DISABLED       ENABLED\nDISABLED     ENABLED        ENABLED\nDISABLED     DISABLED       DISABLED\nIf you want to enable state updates for only some content switching virtual servers, be sure to disable the state update parameter.",
			},
			"targettype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"v6persistmasklen": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(128),
				Description: "Persistence mask for IP based persistence types, for IPv6 virtual servers.",
			},
			"vipheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of virtual server IP and port header, for use with the VServer IP Port Insertion parameter.",
			},
		},
	}
}

func csvserverGetThePayloadFromtheConfig(ctx context.Context, data *CsvserverResourceModel) cs.Csvserver {
	tflog.Debug(ctx, "In csvserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	csvserver := cs.Csvserver{}
	if !data.Apiprofile.IsNull() {
		csvserver.Apiprofile = data.Apiprofile.ValueString()
	}
	if !data.Appflowlog.IsNull() {
		csvserver.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Authentication.IsNull() {
		csvserver.Authentication = data.Authentication.ValueString()
	}
	if !data.Authenticationhost.IsNull() {
		csvserver.Authenticationhost = data.Authenticationhost.ValueString()
	}
	if !data.Authn401.IsNull() {
		csvserver.Authn401 = data.Authn401.ValueString()
	}
	if !data.Authnprofile.IsNull() {
		csvserver.Authnprofile = data.Authnprofile.ValueString()
	}
	if !data.Authnvsname.IsNull() {
		csvserver.Authnvsname = data.Authnvsname.ValueString()
	}
	if !data.Backupip.IsNull() {
		csvserver.Backupip = data.Backupip.ValueString()
	}
	if !data.Backuppersistencetimeout.IsNull() {
		csvserver.Backuppersistencetimeout = utils.IntPtr(int(data.Backuppersistencetimeout.ValueInt64()))
	}
	if !data.Backupvserver.IsNull() {
		csvserver.Backupvserver = data.Backupvserver.ValueString()
	}
	if !data.Cacheable.IsNull() {
		csvserver.Cacheable = data.Cacheable.ValueString()
	}
	if !data.Casesensitive.IsNull() {
		csvserver.Casesensitive = data.Casesensitive.ValueString()
	}
	if !data.Clttimeout.IsNull() {
		csvserver.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
	}
	if !data.Comment.IsNull() {
		csvserver.Comment = data.Comment.ValueString()
	}
	if !data.Cookiedomain.IsNull() {
		csvserver.Cookiedomain = data.Cookiedomain.ValueString()
	}
	if !data.Cookiename.IsNull() {
		csvserver.Cookiename = data.Cookiename.ValueString()
	}
	if !data.Cookietimeout.IsNull() {
		csvserver.Cookietimeout = utils.IntPtr(int(data.Cookietimeout.ValueInt64()))
	}
	if !data.Dbprofilename.IsNull() {
		csvserver.Dbprofilename = data.Dbprofilename.ValueString()
	}
	if !data.Disableprimaryondown.IsNull() {
		csvserver.Disableprimaryondown = data.Disableprimaryondown.ValueString()
	}
	if !data.Dnsoverhttps.IsNull() {
		csvserver.Dnsoverhttps = data.Dnsoverhttps.ValueString()
	}
	if !data.Dnsprofilename.IsNull() {
		csvserver.Dnsprofilename = data.Dnsprofilename.ValueString()
	}
	if !data.Dnsrecordtype.IsNull() {
		csvserver.Dnsrecordtype = data.Dnsrecordtype.ValueString()
	}
	if !data.Domainname.IsNull() {
		csvserver.Domainname = data.Domainname.ValueString()
	}
	if !data.Downstateflush.IsNull() {
		csvserver.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Dtls.IsNull() {
		csvserver.Dtls = data.Dtls.ValueString()
	}
	if !data.Httpprofilename.IsNull() {
		csvserver.Httpprofilename = data.Httpprofilename.ValueString()
	}
	if !data.Httpsredirecturl.IsNull() {
		csvserver.Httpsredirecturl = data.Httpsredirecturl.ValueString()
	}
	if !data.Icmpvsrresponse.IsNull() {
		csvserver.Icmpvsrresponse = data.Icmpvsrresponse.ValueString()
	}
	if !data.Insertvserveripport.IsNull() {
		csvserver.Insertvserveripport = data.Insertvserveripport.ValueString()
	}
	if !data.Ipmask.IsNull() {
		csvserver.Ipmask = data.Ipmask.ValueString()
	}
	if !data.Ippattern.IsNull() {
		csvserver.Ippattern = data.Ippattern.ValueString()
	}
	if !data.Ipset.IsNull() {
		csvserver.Ipset = data.Ipset.ValueString()
	}
	if !data.Ipv46.IsNull() {
		csvserver.Ipv46 = data.Ipv46.ValueString()
	}
	if !data.L2conn.IsNull() {
		csvserver.L2conn = data.L2conn.ValueString()
	}
	if !data.Listenpolicy.IsNull() {
		csvserver.Listenpolicy = data.Listenpolicy.ValueString()
	}
	if !data.Listenpriority.IsNull() {
		csvserver.Listenpriority = utils.IntPtr(int(data.Listenpriority.ValueInt64()))
	}
	if !data.Mssqlserverversion.IsNull() {
		csvserver.Mssqlserverversion = data.Mssqlserverversion.ValueString()
	}
	if !data.Mysqlcharacterset.IsNull() {
		csvserver.Mysqlcharacterset = utils.IntPtr(int(data.Mysqlcharacterset.ValueInt64()))
	}
	if !data.Mysqlprotocolversion.IsNull() {
		csvserver.Mysqlprotocolversion = utils.IntPtr(int(data.Mysqlprotocolversion.ValueInt64()))
	}
	if !data.Mysqlservercapabilities.IsNull() {
		csvserver.Mysqlservercapabilities = utils.IntPtr(int(data.Mysqlservercapabilities.ValueInt64()))
	}
	if !data.Mysqlserverversion.IsNull() {
		csvserver.Mysqlserverversion = data.Mysqlserverversion.ValueString()
	}
	if !data.Name.IsNull() {
		csvserver.Name = data.Name.ValueString()
	}
	if !data.Netprofile.IsNull() {
		csvserver.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Newname.IsNull() {
		csvserver.Newname = data.Newname.ValueString()
	}
	if !data.Oracleserverversion.IsNull() {
		csvserver.Oracleserverversion = data.Oracleserverversion.ValueString()
	}
	if !data.Persistencebackup.IsNull() {
		csvserver.Persistencebackup = data.Persistencebackup.ValueString()
	}
	if !data.Persistenceid.IsNull() {
		csvserver.Persistenceid = utils.IntPtr(int(data.Persistenceid.ValueInt64()))
	}
	if !data.Persistencetype.IsNull() {
		csvserver.Persistencetype = data.Persistencetype.ValueString()
	}
	if !data.Persistmask.IsNull() {
		csvserver.Persistmask = data.Persistmask.ValueString()
	}
	if !data.Port.IsNull() {
		csvserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Precedence.IsNull() {
		csvserver.Precedence = data.Precedence.ValueString()
	}
	if !data.Probeport.IsNull() {
		csvserver.Probeport = utils.IntPtr(int(data.Probeport.ValueInt64()))
	}
	if !data.Probeprotocol.IsNull() {
		csvserver.Probeprotocol = data.Probeprotocol.ValueString()
	}
	if !data.Probesuccessresponsecode.IsNull() {
		csvserver.Probesuccessresponsecode = data.Probesuccessresponsecode.ValueString()
	}
	if !data.Push.IsNull() {
		csvserver.Push = data.Push.ValueString()
	}
	if !data.Pushlabel.IsNull() {
		csvserver.Pushlabel = data.Pushlabel.ValueString()
	}
	if !data.Pushmulticlients.IsNull() {
		csvserver.Pushmulticlients = data.Pushmulticlients.ValueString()
	}
	if !data.Pushvserver.IsNull() {
		csvserver.Pushvserver = data.Pushvserver.ValueString()
	}
	if !data.Quicprofilename.IsNull() {
		csvserver.Quicprofilename = data.Quicprofilename.ValueString()
	}
	if !data.Range.IsNull() {
		csvserver.Range = utils.IntPtr(int(data.Range.ValueInt64()))
	}
	if !data.Redirectfromport.IsNull() {
		csvserver.Redirectfromport = utils.IntPtr(int(data.Redirectfromport.ValueInt64()))
	}
	if !data.Redirectportrewrite.IsNull() {
		csvserver.Redirectportrewrite = data.Redirectportrewrite.ValueString()
	}
	if !data.Redirecturl.IsNull() {
		csvserver.Redirecturl = data.Redirecturl.ValueString()
	}
	if !data.Rhistate.IsNull() {
		csvserver.Rhistate = data.Rhistate.ValueString()
	}
	if !data.Rtspnat.IsNull() {
		csvserver.Rtspnat = data.Rtspnat.ValueString()
	}
	if !data.Servicetype.IsNull() {
		csvserver.Servicetype = data.Servicetype.ValueString()
	}
	if !data.Sitedomainttl.IsNull() {
		csvserver.Sitedomainttl = utils.IntPtr(int(data.Sitedomainttl.ValueInt64()))
	}
	if !data.Sobackupaction.IsNull() {
		csvserver.Sobackupaction = data.Sobackupaction.ValueString()
	}
	if !data.Somethod.IsNull() {
		csvserver.Somethod = data.Somethod.ValueString()
	}
	if !data.Sopersistence.IsNull() {
		csvserver.Sopersistence = data.Sopersistence.ValueString()
	}
	if !data.Sopersistencetimeout.IsNull() {
		csvserver.Sopersistencetimeout = utils.IntPtr(int(data.Sopersistencetimeout.ValueInt64()))
	}
	if !data.Sothreshold.IsNull() {
		csvserver.Sothreshold = utils.IntPtr(int(data.Sothreshold.ValueInt64()))
	}
	if !data.State.IsNull() {
		csvserver.State = data.State.ValueString()
	}
	if !data.Stateupdate.IsNull() {
		csvserver.Stateupdate = data.Stateupdate.ValueString()
	}
	if !data.Targettype.IsNull() {
		csvserver.Targettype = data.Targettype.ValueString()
	}
	if !data.Tcpprobeport.IsNull() {
		csvserver.Tcpprobeport = utils.IntPtr(int(data.Tcpprobeport.ValueInt64()))
	}
	if !data.Tcpprofilename.IsNull() {
		csvserver.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Td.IsNull() {
		csvserver.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Timeout.IsNull() {
		csvserver.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Ttl.IsNull() {
		csvserver.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}
	if !data.V6persistmasklen.IsNull() {
		csvserver.V6persistmasklen = utils.IntPtr(int(data.V6persistmasklen.ValueInt64()))
	}
	if !data.Vipheader.IsNull() {
		csvserver.Vipheader = data.Vipheader.ValueString()
	}

	return csvserver
}

func csvserverSetAttrFromGet(ctx context.Context, data *CsvserverResourceModel, getResponseData map[string]interface{}) *CsvserverResourceModel {
	tflog.Debug(ctx, "In csvserverSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["backupip"]; ok && val != nil {
		data.Backupip = types.StringValue(val.(string))
	} else {
		data.Backupip = types.StringNull()
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
	if val, ok := getResponseData["cacheable"]; ok && val != nil {
		data.Cacheable = types.StringValue(val.(string))
	} else {
		data.Cacheable = types.StringNull()
	}
	if val, ok := getResponseData["casesensitive"]; ok && val != nil {
		data.Casesensitive = types.StringValue(val.(string))
	} else {
		data.Casesensitive = types.StringNull()
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
	if val, ok := getResponseData["cookiedomain"]; ok && val != nil {
		data.Cookiedomain = types.StringValue(val.(string))
	} else {
		data.Cookiedomain = types.StringNull()
	}
	if val, ok := getResponseData["cookiename"]; ok && val != nil {
		data.Cookiename = types.StringValue(val.(string))
	} else {
		data.Cookiename = types.StringNull()
	}
	if val, ok := getResponseData["cookietimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cookietimeout = types.Int64Value(intVal)
		}
	} else {
		data.Cookietimeout = types.Int64Null()
	}
	if val, ok := getResponseData["dbprofilename"]; ok && val != nil {
		data.Dbprofilename = types.StringValue(val.(string))
	} else {
		data.Dbprofilename = types.StringNull()
	}
	if val, ok := getResponseData["disableprimaryondown"]; ok && val != nil {
		data.Disableprimaryondown = types.StringValue(val.(string))
	} else {
		data.Disableprimaryondown = types.StringNull()
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
	if val, ok := getResponseData["dnsrecordtype"]; ok && val != nil {
		data.Dnsrecordtype = types.StringValue(val.(string))
	} else {
		data.Dnsrecordtype = types.StringNull()
	}
	if val, ok := getResponseData["domainname"]; ok && val != nil {
		data.Domainname = types.StringValue(val.(string))
	} else {
		data.Domainname = types.StringNull()
	}
	if val, ok := getResponseData["downstateflush"]; ok && val != nil {
		data.Downstateflush = types.StringValue(val.(string))
	} else {
		data.Downstateflush = types.StringNull()
	}
	if val, ok := getResponseData["dtls"]; ok && val != nil {
		data.Dtls = types.StringValue(val.(string))
	} else {
		data.Dtls = types.StringNull()
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
	if val, ok := getResponseData["oracleserverversion"]; ok && val != nil {
		data.Oracleserverversion = types.StringValue(val.(string))
	} else {
		data.Oracleserverversion = types.StringNull()
	}
	if val, ok := getResponseData["persistencebackup"]; ok && val != nil {
		data.Persistencebackup = types.StringValue(val.(string))
	} else {
		data.Persistencebackup = types.StringNull()
	}
	if val, ok := getResponseData["persistenceid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Persistenceid = types.Int64Value(intVal)
		}
	} else {
		data.Persistenceid = types.Int64Null()
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
	if val, ok := getResponseData["redirecturl"]; ok && val != nil {
		data.Redirecturl = types.StringValue(val.(string))
	} else {
		data.Redirecturl = types.StringNull()
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
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	} else {
		data.Servicetype = types.StringNull()
	}
	if val, ok := getResponseData["sitedomainttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sitedomainttl = types.Int64Value(intVal)
		}
	} else {
		data.Sitedomainttl = types.Int64Null()
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
	if val, ok := getResponseData["stateupdate"]; ok && val != nil {
		data.Stateupdate = types.StringValue(val.(string))
	} else {
		data.Stateupdate = types.StringNull()
	}
	if val, ok := getResponseData["targettype"]; ok && val != nil {
		data.Targettype = types.StringValue(val.(string))
	} else {
		data.Targettype = types.StringNull()
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
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else {
		data.Ttl = types.Int64Null()
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
