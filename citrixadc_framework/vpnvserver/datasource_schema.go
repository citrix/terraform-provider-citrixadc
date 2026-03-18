package vpnvserver

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

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
			"windowsepapluginupgrade": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to set plugin upgrade behaviour for Win",
			},
		},
	}
}
