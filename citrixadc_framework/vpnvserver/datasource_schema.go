package vpnvserver

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnvserverDataSourceModel is the data-source-specific model, decoupled from
// VpnvserverResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime metadata attributes that the resource
// deliberately omits (ip, curstate, status, curaaausers, ...). Every non-key
// attribute is Computed.
type VpnvserverDataSourceModel struct {
	Id                           types.String `tfsdk:"id"`
	Accessrestrictedpageredirect types.String `tfsdk:"accessrestrictedpageredirect"`
	Advancedepa                  types.String `tfsdk:"advancedepa"`
	Appflowlog                   types.String `tfsdk:"appflowlog"`
	Authentication               types.String `tfsdk:"authentication"`
	Authnprofile                 types.String `tfsdk:"authnprofile"`
	Certkeynames                 types.String `tfsdk:"certkeynames"`
	Cginfrahomepageredirect      types.String `tfsdk:"cginfrahomepageredirect"`
	Comment                      types.String `tfsdk:"comment"`
	Deploymenttype               types.String `tfsdk:"deploymenttype"`
	Devicecert                   types.String `tfsdk:"devicecert"`
	Deviceposture                types.String `tfsdk:"deviceposture"`
	Doublehop                    types.String `tfsdk:"doublehop"`
	Downstateflush               types.String `tfsdk:"downstateflush"`
	Dtls                         types.String `tfsdk:"dtls"`
	Failedlogintimeout           types.Int64  `tfsdk:"failedlogintimeout"`
	Gslbsitefqdn                 types.String `tfsdk:"gslbsitefqdn"`
	Httpprofilename              types.String `tfsdk:"httpprofilename"`
	Icaonly                      types.String `tfsdk:"icaonly"`
	Icaproxysessionmigration     types.String `tfsdk:"icaproxysessionmigration"`
	Icmpvsrresponse              types.String `tfsdk:"icmpvsrresponse"`
	Ipset                        types.String `tfsdk:"ipset"`
	Ipv46                        types.String `tfsdk:"ipv46"`
	L2conn                       types.String `tfsdk:"l2conn"`
	Linuxepapluginupgrade        types.String `tfsdk:"linuxepapluginupgrade"`
	Listenpolicy                 types.String `tfsdk:"listenpolicy"`
	Listenpriority               types.Int64  `tfsdk:"listenpriority"`
	Loginonce                    types.String `tfsdk:"loginonce"`
	Logoutonsmartcardremoval     types.String `tfsdk:"logoutonsmartcardremoval"`
	Macepapluginupgrade          types.String `tfsdk:"macepapluginupgrade"`
	Maxaaausers                  types.Int64  `tfsdk:"maxaaausers"`
	Maxloginattempts             types.Int64  `tfsdk:"maxloginattempts"`
	Name                         types.String `tfsdk:"name"` // Required lookup key
	Netprofile                   types.String `tfsdk:"netprofile"`
	Newname                      types.String `tfsdk:"newname"`
	Pcoipvserverprofilename      types.String `tfsdk:"pcoipvserverprofilename"`
	Port                         types.Int64  `tfsdk:"port"`
	Quicprofilename              types.String `tfsdk:"quicprofilename"`
	Range                        types.Int64  `tfsdk:"range"`
	Rdpserverprofilename         types.String `tfsdk:"rdpserverprofilename"`
	Rhistate                     types.String `tfsdk:"rhistate"`
	Samesite                     types.String `tfsdk:"samesite"`
	Secureprivateaccess          types.String `tfsdk:"secureprivateaccess"`
	Servicetype                  types.String `tfsdk:"servicetype"`
	State                        types.String `tfsdk:"state"`
	Tcpprofilename               types.String `tfsdk:"tcpprofilename"`
	Userdomains                  types.String `tfsdk:"userdomains"`
	Vserverfqdn                  types.String `tfsdk:"vserverfqdn"`
	Wasmmodule                   types.String `tfsdk:"wasmmodule"`
	Windowsepapluginupgrade      types.String `tfsdk:"windowsepapluginupgrade"`

	// Read-only (GET-only) runtime/metadata from the NITRO doc read-only set
	// (zion73x_readonly/vpnvserver.json). Never settable; populated from GET.
	Ip                   types.String `tfsdk:"ip"`
	Value                types.String `tfsdk:"value"`
	Type                 types.String `tfsdk:"type"`
	Curstate             types.String `tfsdk:"curstate"`
	Status               types.Int64  `tfsdk:"status"`
	Cachetype            types.String `tfsdk:"cachetype"`
	Redirect             types.String `tfsdk:"redirect"`
	Precedence           types.String `tfsdk:"precedence"`
	Redirecturl          types.String `tfsdk:"redirecturl"`
	Curaaausers          types.Int64  `tfsdk:"curaaausers"`
	Curtotalusers        types.Int64  `tfsdk:"curtotalusers"`
	Domain               types.String `tfsdk:"domain"`
	Rule                 types.String `tfsdk:"rule"`
	Servicename          types.String `tfsdk:"servicename"`
	Weight               types.Int64  `tfsdk:"weight"`
	Cachevserver         types.String `tfsdk:"cachevserver"`
	Backupvserver        types.String `tfsdk:"backupvserver"`
	Clttimeout           types.Int64  `tfsdk:"clttimeout"`
	Somethod             types.String `tfsdk:"somethod"`
	Sothreshold          types.Int64  `tfsdk:"sothreshold"`
	Sopersistence        types.String `tfsdk:"sopersistence"`
	Sopersistencetimeout types.Int64  `tfsdk:"sopersistencetimeout"`
	Usemip               types.String `tfsdk:"usemip"`
	Map                  types.String `tfsdk:"map"`
	Bindpoint            types.String `tfsdk:"bindpoint"`
	Disableprimaryondown types.String `tfsdk:"disableprimaryondown"`
	Secondary            types.Bool   `tfsdk:"secondary"`
	Groupextraction      types.Bool   `tfsdk:"groupextraction"`
	Epaprofileoptional   types.Bool   `tfsdk:"epaprofileoptional"`
	Ngname               types.String `tfsdk:"ngname"`
	Csvserver            types.String `tfsdk:"csvserver"`
	Nodefaultbindings    types.String `tfsdk:"nodefaultbindings"`
	Response             types.String `tfsdk:"response"`
}

func VpnvserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"accessrestrictedpageredirect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By default, an access restricted page hosted on secure private access CDN is displayed when a restricted app is accessed. The setting can be changed to NS to display the access restricted page hosted on the gateway or OFF to not display any access restricted page.",
			},
			"advancedepa": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option tells whether advanced EPA is enabled on this virtual server",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log AppFlow records that contain standard NetFlow or IPFIX information, such as time stamps for the beginning and end of a flow, packet count, and byte count. Also log records that contain application-level information, such as HTTP web addresses, HTTP request methods and response status codes, server response time, and latency.",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Require authentication for users connecting to Citrix Gateway.",
			},
			"authnprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication Profile entity on virtual server. This entity can be used to offload authentication to AAA vserver for multi-factor(nFactor) authentication",
			},
			"certkeynames": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the certificate key that was bound to the corresponding SSL virtual server as the Certificate Authority for the device certificate",
			},
			"cginfrahomepageredirect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When client requests ShareFile resources and Citrix Gateway detects that the user is unauthenticated or the user session has expired, disabling this option takes the user to the originally requested ShareFile resource after authentication (instead of taking the user to the default VPN home page)",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the virtual server.",
			},
			"deploymenttype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"devicecert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates whether device certificate check as a part of EPA is on or off.",
			},
			"deviceposture": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable device posture",
			},
			"doublehop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the Citrix Gateway appliance in a double-hop configuration. A double-hop deployment provides an extra layer of security for the internal network by using three firewalls to divide the DMZ into two stages. Such a deployment can have one appliance in the DMZ and one appliance in the secure network.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Close existing connections when the virtual server is marked DOWN, which means the server might have timed out. Disconnecting existing connections frees resources and in certain cases speeds recovery of overloaded load balancing setups. Enable this setting on servers in which the connections can safely be closed when they are marked DOWN.  Do not enable DOWN state flush on servers that must complete their transactions.",
			},
			"dtls": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option starts/stops the turn service on the vserver",
			},
			"failedlogintimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of minutes an account will be locked if user exceeds maximum permissible attempts",
			},
			"gslbsitefqdn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name of the SPA site. This is used for Secure Private Access configuration.",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the HTTP profile to assign to this virtual server.",
			},
			"icaonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "- When set to ON, it implies Basic mode where the user can log on using either Citrix Receiver or a browser and get access to the published apps configured at the XenApp/XenDEsktop environment pointed out by the WIHome parameter. Users are not allowed to connect using the Citrix Gateway Plug-in and end point scans cannot be configured. Number of users that can log in and access the apps are not limited by the license in this mode.\n\n- When set to OFF, it implies Smart Access mode where the user can log on using either Citrix Receiver or a browser or a Citrix Gateway Plug-in. The admin can configure end point scans to be run on the client systems and then use the results to control access to the published apps. In this mode, the client can connect to the gateway in other client modes namely VPN and CVPN. Number of users that can log in and access the resources are limited by the CCU licenses in this mode.",
			},
			"icaproxysessionmigration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option determines if an existing ICA Proxy session is transferred when the user logs on from another device.",
			},
			"icmpvsrresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Criterion for responding to PING requests sent to this virtual server. If this parameter is set to ACTIVE, respond only if the virtual server is available. With the PASSIVE setting, respond even if the virtual server is not available.",
			},
			"ipset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current vpn vserver",
			},
			"ipv46": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address of the Citrix Gateway virtual server. Usually a public IP address. User devices send connection requests to this IP address.",
			},
			"l2conn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use Layer 2 parameters (channel number, MAC address, and VLAN ID) in addition to the 4-tuple (<source IP>:<source port>::<destination IP>:<destination port>) that is used to identify a connection. Allows multiple TCP and non-TCP connections with the same 4-tuple to coexist on the Citrix ADC.",
			},
			"linuxepapluginupgrade": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to set plugin upgrade behaviour for Linux",
			},
			"listenpolicy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the listen policy for the Citrix Gateway virtual server. Can be either a named expression or an expression. The Citrix Gateway virtual server processes only the traffic for which the expression evaluates to true.",
			},
			"listenpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server, the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.",
			},
			"loginonce": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option enables/disables seamless SSO for this Vserver.",
			},
			"logoutonsmartcardremoval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to VPN plugin behavior when smartcard or its reader is removed",
			},
			"macepapluginupgrade": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to set plugin upgrade behaviour for Mac",
			},
			"maxaaausers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent user sessions allowed on this virtual server. The actual number of users allowed to log on to this virtual server depends on the total number of user licenses.",
			},
			"maxloginattempts": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of logon attempts",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Citrix Gateway virtual server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my server\" or 'my server').",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the network profile.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the Citrix Gateway virtual server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my server\" or 'my server').",
			},
			"pcoipvserverprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the PCoIP vserver profile associated with the vserver.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP port on which the virtual server listens.",
			},
			"quicprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the QUIC profile to assign to this virtual server.",
			},
			"range": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Range of Citrix Gateway virtual server IP addresses. The consecutively numbered range of IP addresses begins with the address specified by the IP Address parameter.\nIn the configuration utility, select Network VServer to enter a range.",
			},
			"rdpserverprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the RDP server profile associated with the vserver.",
			},
			"rhistate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A host route is injected according to the setting on the virtual servers.\n            * If set to PASSIVE on all the virtual servers that share the IP address, the appliance always injects the hostroute.\n            * If set to ACTIVE on all the virtual servers that share the IP address, the appliance injects even if one virtual server is UP.\n            * If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance injects even if one virtual server set to ACTIVE is UP.",
			},
			"samesite": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SameSite attribute value for Cookies generated in VPN context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite",
			},
			"secureprivateaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Configure secure private access",
			},
			"servicetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Protocol used by the Citrix Gateway virtual server.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "State of the virtual server. If the virtual server is disabled, requests are not processed.",
			},
			"tcpprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the TCP profile to assign to this virtual server.",
			},
			"userdomains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of user domains specified as comma seperated value",
			},
			"vserverfqdn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name for a VPN virtual server. This is used during StoreFront configuration generation.",
			},
			"wasmmodule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the WASM module to assign to this virtual server.",
			},
			"windowsepapluginupgrade": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to set plugin upgrade behaviour for Win",
			},

			// Read-only (GET-only) runtime/metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "The Virtual IP address of the VPN virtual server.",
			},
			"value": schema.StringAttribute{
				Computed:    true,
				Description: "Indicates whether or not the certificate is bound or if SSL offload is disabled.",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "The type of virtual server; for example, CONTENT based or ADDRESS based.",
			},
			"curstate": schema.StringAttribute{
				Computed:    true,
				Description: "The current state of the virtual server, as UP, DOWN, BUSY, and so on.",
			},
			"status": schema.Int64Attribute{
				Computed:    true,
				Description: "Whether or not this virtual server responds to ARPs and whether or not round-robin selection is temporarily in effect.",
			},
			"cachetype": schema.StringAttribute{
				Computed:    true,
				Description: "Virtual server cache type. The options are: TRANSPARENT, REVERSE, and FORWARD.",
			},
			"redirect": schema.StringAttribute{
				Computed:    true,
				Description: "The cache redirect policy.",
			},
			"precedence": schema.StringAttribute{
				Computed:    true,
				Description: "The type of policy (URL or RULE) that takes precedence on the content switching virtual server.",
			},
			"redirecturl": schema.StringAttribute{
				Computed:    true,
				Description: "The URL where traffic is redirected if the virtual server in system becomes unavailable.",
			},
			"curaaausers": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of current users logged on to this virtual server.",
			},
			"curtotalusers": schema.Int64Attribute{
				Computed:    true,
				Description: "The total number of current users connected through this virtual server.",
			},
			"domain": schema.StringAttribute{
				Computed:    true,
				Description: "The domain name of the server for which a service needs to be added.",
			},
			"rule": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the rule, or expression, if any, that policy for the VPN server is to use.",
			},
			"servicename": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the service, if any, to which the virtual server policy is bound.",
			},
			"weight": schema.Int64Attribute{
				Computed:    true,
				Description: "Weight for this service, if any. This weight is used when the system performs load balancing.",
			},
			"cachevserver": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the default target cache virtual server, if any, to which requests are redirected.",
			},
			"backupvserver": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the backup VPN virtual server for this VPN virtual server.",
			},
			"clttimeout": schema.Int64Attribute{
				Computed:    true,
				Description: "The idle time, if any, in seconds after which the client connection is terminated.",
			},
			"somethod": schema.StringAttribute{
				Computed:    true,
				Description: "The method used to determine whether or not a new connection will spill over the allocated block of intranet IP addresses.",
			},
			"sothreshold": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of client connections after which the mapped IP address is used as the client source IP address instead of an address from the allocated block of intranet IP addresses.",
			},
			"sopersistence": schema.StringAttribute{
				Computed:    true,
				Description: "Whether or not cookie-based site persistance is enabled for this VPN vserver.",
			},
			"sopersistencetimeout": schema.Int64Attribute{
				Computed:    true,
				Description: "The timeout, if any, for cookie-based site persistance of this VPN vserver.",
			},
			"usemip": schema.StringAttribute{
				Computed:    true,
				Description: "Deprecated. See 'map' below.",
			},
			"map": schema.StringAttribute{
				Computed:    true,
				Description: "Whether or not mapped IP addresses are ON or OFF.",
			},
			"bindpoint": schema.StringAttribute{
				Computed:    true,
				Description: "Bindpoint to which the policy is bound.",
			},
			"disableprimaryondown": schema.StringAttribute{
				Computed:    true,
				Description: "Tells whether traffic will continue reaching backup virtual servers even after the primary virtual server comes UP from DOWN state.",
			},
			"secondary": schema.BoolAttribute{
				Computed:    true,
				Description: "Binds the authentication policy as the secondary policy to use in a two-factor configuration.",
			},
			"groupextraction": schema.BoolAttribute{
				Computed:    true,
				Description: "Binds the authentication policy to a tertiary chain which will be used only for group extraction.",
			},
			"epaprofileoptional": schema.BoolAttribute{
				Computed:    true,
				Description: "Mark the EPA profile optional for preauthentication EPA profile. User would be shown a logon page even if the EPA profile fails to evaluate.",
			},
			"ngname": schema.StringAttribute{
				Computed:    true,
				Description: "Node group devno to which this authentication virtual sever belongs.",
			},
			"csvserver": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the CS vserver to which the VPN vserver is bound.",
			},
			"nodefaultbindings": schema.StringAttribute{
				Computed:    true,
				Description: "Whether the configuration will have default ssl CIPHER and ECC curve bindings.",
			},
			"response": schema.StringAttribute{
				Computed:    true,
				Description: "Response.",
			},
		},
	}
}

// vpnvserverDataSourceSetAttrFromGet projects a NITRO vpnvserver GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func vpnvserverDataSourceSetAttrFromGet(ctx context.Context, data *VpnvserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnvserverDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Accessrestrictedpageredirect = utils.MapGetString(g, "accessrestrictedpageredirect")
	data.Advancedepa = utils.MapGetString(g, "advancedepa")
	data.Appflowlog = utils.MapGetString(g, "appflowlog")
	data.Authentication = utils.MapGetString(g, "authentication")
	data.Authnprofile = utils.MapGetString(g, "authnprofile")
	data.Certkeynames = utils.MapGetString(g, "certkeynames")
	data.Cginfrahomepageredirect = utils.MapGetString(g, "cginfrahomepageredirect")
	data.Comment = utils.MapGetString(g, "comment")
	data.Deploymenttype = utils.MapGetString(g, "deploymenttype")
	data.Devicecert = utils.MapGetString(g, "devicecert")
	data.Deviceposture = utils.MapGetString(g, "deviceposture")
	data.Doublehop = utils.MapGetString(g, "doublehop")
	data.Downstateflush = utils.MapGetString(g, "downstateflush")
	data.Dtls = utils.MapGetString(g, "dtls")
	data.Failedlogintimeout = utils.MapGetInt64(g, "failedlogintimeout")
	data.Gslbsitefqdn = utils.MapGetString(g, "gslbsitefqdn")
	data.Httpprofilename = utils.MapGetString(g, "httpprofilename")
	data.Icaonly = utils.MapGetString(g, "icaonly")
	data.Icaproxysessionmigration = utils.MapGetString(g, "icaproxysessionmigration")
	data.Icmpvsrresponse = utils.MapGetString(g, "icmpvsrresponse")
	data.Ipset = utils.MapGetString(g, "ipset")
	data.Ipv46 = utils.MapGetString(g, "ipv46")
	data.L2conn = utils.MapGetString(g, "l2conn")
	data.Linuxepapluginupgrade = utils.MapGetString(g, "linuxepapluginupgrade")
	data.Listenpolicy = utils.MapGetString(g, "listenpolicy")
	data.Listenpriority = utils.MapGetInt64(g, "listenpriority")
	data.Loginonce = utils.MapGetString(g, "loginonce")
	data.Logoutonsmartcardremoval = utils.MapGetString(g, "logoutonsmartcardremoval")
	data.Macepapluginupgrade = utils.MapGetString(g, "macepapluginupgrade")
	data.Maxaaausers = utils.MapGetInt64(g, "maxaaausers")
	data.Maxloginattempts = utils.MapGetInt64(g, "maxloginattempts")
	data.Netprofile = utils.MapGetString(g, "netprofile")
	data.Pcoipvserverprofilename = utils.MapGetString(g, "pcoipvserverprofilename")
	data.Port = utils.MapGetInt64(g, "port")
	data.Quicprofilename = utils.MapGetString(g, "quicprofilename")
	data.Range = utils.MapGetInt64(g, "range")
	data.Rdpserverprofilename = utils.MapGetString(g, "rdpserverprofilename")
	data.Rhistate = utils.MapGetString(g, "rhistate")
	data.Samesite = utils.MapGetString(g, "samesite")
	data.Secureprivateaccess = utils.MapGetString(g, "secureprivateaccess")
	data.Servicetype = utils.MapGetString(g, "servicetype")
	data.State = utils.MapGetString(g, "state")
	data.Tcpprofilename = utils.MapGetString(g, "tcpprofilename")
	data.Userdomains = utils.MapGetString(g, "userdomains")
	data.Vserverfqdn = utils.MapGetString(g, "vserverfqdn")
	data.Wasmmodule = utils.MapGetString(g, "wasmmodule")
	data.Windowsepapluginupgrade = utils.MapGetString(g, "windowsepapluginupgrade")

	// newname is a rename-only action input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only runtime/metadata.
	data.Ip = utils.MapGetString(g, "ip")
	data.Value = utils.MapGetString(g, "value")
	data.Type = utils.MapGetString(g, "type")
	data.Curstate = utils.MapGetString(g, "curstate")
	data.Status = utils.MapGetInt64(g, "status")
	data.Cachetype = utils.MapGetString(g, "cachetype")
	data.Redirect = utils.MapGetString(g, "redirect")
	data.Precedence = utils.MapGetString(g, "precedence")
	data.Redirecturl = utils.MapGetString(g, "redirecturl")
	data.Curaaausers = utils.MapGetInt64(g, "curaaausers")
	data.Curtotalusers = utils.MapGetInt64(g, "curtotalusers")
	data.Domain = utils.MapGetString(g, "domain")
	data.Rule = utils.MapGetString(g, "rule")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Weight = utils.MapGetInt64(g, "weight")
	data.Cachevserver = utils.MapGetString(g, "cachevserver")
	data.Backupvserver = utils.MapGetString(g, "backupvserver")
	data.Clttimeout = utils.MapGetInt64(g, "clttimeout")
	data.Somethod = utils.MapGetString(g, "somethod")
	data.Sothreshold = utils.MapGetInt64(g, "sothreshold")
	data.Sopersistence = utils.MapGetString(g, "sopersistence")
	data.Sopersistencetimeout = utils.MapGetInt64(g, "sopersistencetimeout")
	data.Usemip = utils.MapGetString(g, "usemip")
	data.Map = utils.MapGetString(g, "map")
	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Disableprimaryondown = utils.MapGetString(g, "disableprimaryondown")
	data.Secondary = utils.MapGetBool(g, "secondary")
	data.Groupextraction = utils.MapGetBool(g, "groupextraction")
	data.Epaprofileoptional = utils.MapGetBool(g, "epaprofileoptional")
	data.Ngname = utils.MapGetString(g, "ngname")
	data.Csvserver = utils.MapGetString(g, "csvserver")
	data.Nodefaultbindings = utils.MapGetString(g, "nodefaultbindings")
	data.Response = utils.MapGetString(g, "response")
}
