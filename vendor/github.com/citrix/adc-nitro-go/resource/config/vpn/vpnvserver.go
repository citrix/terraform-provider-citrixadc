/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package vpn

/**
* Configuration for VPN virtual server resource.
*/
type Vpnvserver struct {
	/**
	* Name for the Citrix Gateway virtual server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my server" or 'my server').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Protocol used by the Citrix Gateway virtual server.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* IPv4 or IPv6 address of the Citrix Gateway virtual server. Usually a public IP address. User devices send connection requests to this IP address.
	*/
	Ipv46 string `json:"ipv46,omitempty"`
	/**
	* Range of Citrix Gateway virtual server IP addresses. The consecutively numbered range of IP addresses begins with the address specified by the IP Address parameter.
		In the configuration utility, select Network VServer to enter a range.
	*/
	Range *int `json:"range,omitempty"`
	/**
	* TCP port on which the virtual server listens.
	*/
	Port *int `json:"port,omitempty"`
	/**
	* The list of IPv4/IPv6 addresses bound to ipset would form a part of listening service on the current vpn vserver
	*/
	Ipset string `json:"ipset,omitempty"`
	/**
	* State of the virtual server. If the virtual server is disabled, requests are not processed.
	*/
	State string `json:"state,omitempty"`
	/**
	* Require authentication for users connecting to Citrix Gateway.
	*/
	Authentication string `json:"authentication,omitempty"`
	/**
	* Use the Citrix Gateway appliance in a double-hop configuration. A double-hop deployment provides an extra layer of security for the internal network by using three firewalls to divide the DMZ into two stages. Such a deployment can have one appliance in the DMZ and one appliance in the secure network.
	*/
	Doublehop string `json:"doublehop,omitempty"`
	/**
	* Maximum number of concurrent user sessions allowed on this virtual server. The actual number of users allowed to log on to this virtual server depends on the total number of user licenses.
	*/
	Maxaaausers *int `json:"maxaaausers,omitempty"`
	/**
	* - When set to ON, it implies Basic mode where the user can log on using either Citrix Receiver or a browser and get access to the published apps configured at the XenApp/XenDEsktop environment pointed out by the WIHome parameter. Users are not allowed to connect using the Citrix Gateway Plug-in and end point scans cannot be configured. Number of users that can log in and access the apps are not limited by the license in this mode.
		- When set to OFF, it implies Smart Access mode where the user can log on using either Citrix Receiver or a browser or a Citrix Gateway Plug-in. The admin can configure end point scans to be run on the client systems and then use the results to control access to the published apps. In this mode, the client can connect to the gateway in other client modes namely VPN and CVPN. Number of users that can log in and access the resources are limited by the CCU licenses in this mode.
	*/
	Icaonly string `json:"icaonly,omitempty"`
	/**
	* This option determines if an existing ICA Proxy session is transferred when the user logs on from another device.
	*/
	Icaproxysessionmigration string `json:"icaproxysessionmigration,omitempty"`
	/**
	* This option starts/stops the turn service on the vserver
	*/
	Dtls string `json:"dtls,omitempty"`
	/**
	* This option enables/disables seamless SSO for this Vserver.
	*/
	Loginonce string `json:"loginonce,omitempty"`
	/**
	* This option tells whether advanced EPA is enabled on this virtual server
	*/
	Advancedepa string `json:"advancedepa,omitempty"`
	/**
	* Indicates whether device certificate check as a part of EPA is on or off.
	*/
	Devicecert string `json:"devicecert,omitempty"`
	/**
	* Name of the certificate key that was bound to the corresponding SSL virtual server as the Certificate Authority for the device certificate
	*/
	Certkeynames string `json:"certkeynames,omitempty"`
	/**
	* Close existing connections when the virtual server is marked DOWN, which means the server might have timed out. Disconnecting existing connections frees resources and in certain cases speeds recovery of overloaded load balancing setups. Enable this setting on servers in which the connections can safely be closed when they are marked DOWN.  Do not enable DOWN state flush on servers that must complete their transactions.
	*/
	Downstateflush string `json:"downstateflush,omitempty"`
	/**
	* String specifying the listen policy for the Citrix Gateway virtual server. Can be either a named expression or an expression. The Citrix Gateway virtual server processes only the traffic for which the expression evaluates to true.
	*/
	Listenpolicy string `json:"listenpolicy,omitempty"`
	/**
	* Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server, the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.
	*/
	Listenpriority *int `json:"listenpriority,omitempty"`
	/**
	* Name of the TCP profile to assign to this virtual server.
	*/
	Tcpprofilename string `json:"tcpprofilename,omitempty"`
	/**
	* Name of the HTTP profile to assign to this virtual server.
	*/
	Httpprofilename string `json:"httpprofilename,omitempty"`
	/**
	* Any comments associated with the virtual server.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Log AppFlow records that contain standard NetFlow or IPFIX information, such as time stamps for the beginning and end of a flow, packet count, and byte count. Also log records that contain application-level information, such as HTTP web addresses, HTTP request methods and response status codes, server response time, and latency.
	*/
	Appflowlog string `json:"appflowlog,omitempty"`
	/**
	* Criterion for responding to PING requests sent to this virtual server. If this parameter is set to ACTIVE, respond only if the virtual server is available. With the PASSIVE setting, respond even if the virtual server is not available.
	*/
	Icmpvsrresponse string `json:"icmpvsrresponse,omitempty"`
	/**
	* A host route is injected according to the setting on the virtual servers.
		* If set to PASSIVE on all the virtual servers that share the IP address, the appliance always injects the hostroute.
		* If set to ACTIVE on all the virtual servers that share the IP address, the appliance injects even if one virtual server is UP.
		* If set to ACTIVE on some virtual servers and PASSIVE on the others, the appliance injects even if one virtual server set to ACTIVE is UP.
	*/
	Rhistate string `json:"rhistate,omitempty"`
	/**
	* The name of the network profile.
	*/
	Netprofile string `json:"netprofile,omitempty"`
	/**
	* When client requests ShareFile resources and Citrix Gateway detects that the user is unauthenticated or the user session has expired, disabling this option takes the user to the originally requested ShareFile resource after authentication (instead of taking the user to the default VPN home page)
	*/
	Cginfrahomepageredirect string `json:"cginfrahomepageredirect,omitempty"`
	/**
	* Configure secure private access
	*/
	Secureprivateaccess string `json:"secureprivateaccess,omitempty"`
	/**
	* By default, an access restricted page hosted on secure private access CDN is displayed when a restricted app is accessed. The setting can be changed to NS to display the access restricted page hosted on the gateway or OFF to not display any access restricted page.
	*/
	Accessrestrictedpageredirect string `json:"accessrestrictedpageredirect,omitempty"`
	/**
	* Maximum number of logon attempts
	*/
	Maxloginattempts *int `json:"maxloginattempts,omitempty"`
	/**
	* Number of minutes an account will be locked if user exceeds maximum permissible attempts
	*/
	Failedlogintimeout *int `json:"failedlogintimeout,omitempty"`
	/**
	* Use Layer 2 parameters (channel number, MAC address, and VLAN ID) in addition to the 4-tuple (<source IP>:<source port>::<destination IP>:<destination port>) that is used to identify a connection. Allows multiple TCP and non-TCP connections with the same 4-tuple to coexist on the Citrix ADC.
	*/
	L2conn string `json:"l2conn,omitempty"`
	Deploymenttype string `json:"deploymenttype,omitempty"`
	/**
	* Name of the RDP server profile associated with the vserver.
	*/
	Rdpserverprofilename string `json:"rdpserverprofilename,omitempty"`
	/**
	* Option to set plugin upgrade behaviour for Win
	*/
	Windowsepapluginupgrade string `json:"windowsepapluginupgrade,omitempty"`
	/**
	* Option to set plugin upgrade behaviour for Linux
	*/
	Linuxepapluginupgrade string `json:"linuxepapluginupgrade,omitempty"`
	/**
	* Option to set plugin upgrade behaviour for Mac
	*/
	Macepapluginupgrade string `json:"macepapluginupgrade,omitempty"`
	/**
	* Option to VPN plugin behavior when smartcard or its reader is removed
	*/
	Logoutonsmartcardremoval string `json:"logoutonsmartcardremoval,omitempty"`
	/**
	* List of user domains specified as comma seperated value
	*/
	Userdomains string `json:"userdomains,omitempty"`
	/**
	* Authentication Profile entity on virtual server. This entity can be used to offload authentication to AAA vserver for multi-factor(nFactor) authentication
	*/
	Authnprofile string `json:"authnprofile,omitempty"`
	/**
	* Fully qualified domain name for a VPN virtual server. This is used during StoreFront configuration generation.
	*/
	Vserverfqdn string `json:"vserverfqdn,omitempty"`
	/**
	* Name of the PCoIP vserver profile associated with the vserver.
	*/
	Pcoipvserverprofilename string `json:"pcoipvserverprofilename,omitempty"`
	/**
	* SameSite attribute value for Cookies generated in VPN context. This attribute value will be appended only for the cookies which are specified in the builtin patset ns_cookies_samesite
	*/
	Samesite string `json:"samesite,omitempty"`
	/**
	* Name of the QUIC profile to assign to this virtual server.
	*/
	Quicprofilename string `json:"quicprofilename,omitempty"`
	/**
	* Enable device posture
	*/
	Deviceposture string `json:"deviceposture,omitempty"`
	/**
	* New name for the Citrix Gateway virtual server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my server" or 'my server').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Ip string `json:"ip,omitempty"`
	Value string `json:"value,omitempty"`
	Type string `json:"type,omitempty"`
	Curstate string `json:"curstate,omitempty"`
	Status string `json:"status,omitempty"`
	Cachetype string `json:"cachetype,omitempty"`
	Redirect string `json:"redirect,omitempty"`
	Precedence string `json:"precedence,omitempty"`
	Redirecturl string `json:"redirecturl,omitempty"`
	Curaaausers string `json:"curaaausers,omitempty"`
	Curtotalusers string `json:"curtotalusers,omitempty"`
	Domain string `json:"domain,omitempty"`
	Rule string `json:"rule,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Weight string `json:"weight,omitempty"`
	Cachevserver string `json:"cachevserver,omitempty"`
	Backupvserver string `json:"backupvserver,omitempty"`
	Clttimeout string `json:"clttimeout,omitempty"`
	Somethod string `json:"somethod,omitempty"`
	Sothreshold string `json:"sothreshold,omitempty"`
	Sopersistence string `json:"sopersistence,omitempty"`
	Sopersistencetimeout string `json:"sopersistencetimeout,omitempty"`
	Usemip string `json:"usemip,omitempty"`
	Map string `json:"map,omitempty"`
	Bindpoint string `json:"bindpoint,omitempty"`
	Disableprimaryondown string `json:"disableprimaryondown,omitempty"`
	Secondary string `json:"secondary,omitempty"`
	Groupextraction string `json:"groupextraction,omitempty"`
	Epaprofileoptional string `json:"epaprofileoptional,omitempty"`
	Ngname string `json:"ngname,omitempty"`
	Csvserver string `json:"csvserver,omitempty"`
	Nodefaultbindings string `json:"nodefaultbindings,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
	Response string `json:"response,omitempty"`

}
