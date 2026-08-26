package vpnvserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnvserverResourceModel describes the resource data model.
type VpnvserverResourceModel struct {
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
	Name                         types.String `tfsdk:"name"`
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
}

func (r *VpnvserverResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnvserver resource.",
			},
			// SDK v2: Optional+Computed, no default (read value from ADC).
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
				Default:     stringdefault.StaticString("ENABLED"),
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
				Default:     stringdefault.StaticString("ENABLED"),
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
				Default:     stringdefault.StaticString("DISABLED"),
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
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Close existing connections when the virtual server is marked DOWN, which means the server might have timed out. Disconnecting existing connections frees resources and in certain cases speeds recovery of overloaded load balancing setups. Enable this setting on servers in which the connections can safely be closed when they are marked DOWN.  Do not enable DOWN state flush on servers that must complete their transactions.",
			},
			"dtls": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ON"),
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
				Default:     stringdefault.StaticString("PASSIVE"),
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
				Default:     stringdefault.StaticString("OFF"),
				Description: "This option enables/disables seamless SSO for this Vserver.",
			},
			"logoutonsmartcardremoval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
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
			// SDK v2: Required + ForceNew.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Citrix Gateway virtual server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the virtual server is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my server\" or 'my server').",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the network profile.",
			},
			// Rename-only attribute (NITRO ?action=rename). Optional user input, never
			// echoed by GET; excluded from the add payload and driven in-place in Update.
			"newname": schema.StringAttribute{
				Optional:    true,
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
				Default:     stringdefault.StaticString("PASSIVE"),
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Configure secure private access",
			},
			// SDK v2: Required + ForceNew.
			"servicetype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
		},
	}
}

// vpnvserverGetThePayloadFromthePlan builds the NITRO add/update body from the model.
// newname is excluded (it is rename-only, handled via ?action=rename in Update).
func vpnvserverGetThePayloadFromthePlan(ctx context.Context, data *VpnvserverResourceModel) vpn.Vpnvserver {
	tflog.Debug(ctx, "In vpnvserverGetThePayloadFromthePlan Function")

	vpnvserver := vpn.Vpnvserver{}
	if !data.Accessrestrictedpageredirect.IsNull() && !data.Accessrestrictedpageredirect.IsUnknown() {
		vpnvserver.Accessrestrictedpageredirect = data.Accessrestrictedpageredirect.ValueString()
	}
	if !data.Advancedepa.IsNull() && !data.Advancedepa.IsUnknown() {
		vpnvserver.Advancedepa = data.Advancedepa.ValueString()
	}
	if !data.Appflowlog.IsNull() && !data.Appflowlog.IsUnknown() {
		vpnvserver.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Authentication.IsNull() && !data.Authentication.IsUnknown() {
		vpnvserver.Authentication = data.Authentication.ValueString()
	}
	if !data.Authnprofile.IsNull() && !data.Authnprofile.IsUnknown() {
		vpnvserver.Authnprofile = data.Authnprofile.ValueString()
	}
	if !data.Certkeynames.IsNull() && !data.Certkeynames.IsUnknown() {
		vpnvserver.Certkeynames = data.Certkeynames.ValueString()
	}
	if !data.Cginfrahomepageredirect.IsNull() && !data.Cginfrahomepageredirect.IsUnknown() {
		vpnvserver.Cginfrahomepageredirect = data.Cginfrahomepageredirect.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		vpnvserver.Comment = data.Comment.ValueString()
	}
	if !data.Deploymenttype.IsNull() && !data.Deploymenttype.IsUnknown() {
		vpnvserver.Deploymenttype = data.Deploymenttype.ValueString()
	}
	if !data.Devicecert.IsNull() && !data.Devicecert.IsUnknown() {
		vpnvserver.Devicecert = data.Devicecert.ValueString()
	}
	if !data.Deviceposture.IsNull() && !data.Deviceposture.IsUnknown() {
		vpnvserver.Deviceposture = data.Deviceposture.ValueString()
	}
	if !data.Doublehop.IsNull() && !data.Doublehop.IsUnknown() {
		vpnvserver.Doublehop = data.Doublehop.ValueString()
	}
	if !data.Downstateflush.IsNull() && !data.Downstateflush.IsUnknown() {
		vpnvserver.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Dtls.IsNull() && !data.Dtls.IsUnknown() {
		vpnvserver.Dtls = data.Dtls.ValueString()
	}
	if !data.Failedlogintimeout.IsNull() && !data.Failedlogintimeout.IsUnknown() {
		vpnvserver.Failedlogintimeout = utils.IntPtr(int(data.Failedlogintimeout.ValueInt64()))
	}
	if !data.Gslbsitefqdn.IsNull() && !data.Gslbsitefqdn.IsUnknown() {
		vpnvserver.Gslbsitefqdn = data.Gslbsitefqdn.ValueString()
	}
	if !data.Httpprofilename.IsNull() && !data.Httpprofilename.IsUnknown() {
		vpnvserver.Httpprofilename = data.Httpprofilename.ValueString()
	}
	if !data.Icaonly.IsNull() && !data.Icaonly.IsUnknown() {
		vpnvserver.Icaonly = data.Icaonly.ValueString()
	}
	if !data.Icaproxysessionmigration.IsNull() && !data.Icaproxysessionmigration.IsUnknown() {
		vpnvserver.Icaproxysessionmigration = data.Icaproxysessionmigration.ValueString()
	}
	if !data.Icmpvsrresponse.IsNull() && !data.Icmpvsrresponse.IsUnknown() {
		vpnvserver.Icmpvsrresponse = data.Icmpvsrresponse.ValueString()
	}
	if !data.Ipset.IsNull() && !data.Ipset.IsUnknown() {
		vpnvserver.Ipset = data.Ipset.ValueString()
	}
	if !data.Ipv46.IsNull() && !data.Ipv46.IsUnknown() {
		vpnvserver.Ipv46 = data.Ipv46.ValueString()
	}
	if !data.L2conn.IsNull() && !data.L2conn.IsUnknown() {
		vpnvserver.L2conn = data.L2conn.ValueString()
	}
	if !data.Linuxepapluginupgrade.IsNull() && !data.Linuxepapluginupgrade.IsUnknown() {
		vpnvserver.Linuxepapluginupgrade = data.Linuxepapluginupgrade.ValueString()
	}
	if !data.Listenpolicy.IsNull() && !data.Listenpolicy.IsUnknown() {
		vpnvserver.Listenpolicy = data.Listenpolicy.ValueString()
	}
	if !data.Listenpriority.IsNull() && !data.Listenpriority.IsUnknown() {
		vpnvserver.Listenpriority = utils.IntPtr(int(data.Listenpriority.ValueInt64()))
	}
	if !data.Loginonce.IsNull() && !data.Loginonce.IsUnknown() {
		vpnvserver.Loginonce = data.Loginonce.ValueString()
	}
	if !data.Logoutonsmartcardremoval.IsNull() && !data.Logoutonsmartcardremoval.IsUnknown() {
		vpnvserver.Logoutonsmartcardremoval = data.Logoutonsmartcardremoval.ValueString()
	}
	if !data.Macepapluginupgrade.IsNull() && !data.Macepapluginupgrade.IsUnknown() {
		vpnvserver.Macepapluginupgrade = data.Macepapluginupgrade.ValueString()
	}
	if !data.Maxaaausers.IsNull() && !data.Maxaaausers.IsUnknown() {
		vpnvserver.Maxaaausers = utils.IntPtr(int(data.Maxaaausers.ValueInt64()))
	}
	if !data.Maxloginattempts.IsNull() && !data.Maxloginattempts.IsUnknown() {
		vpnvserver.Maxloginattempts = utils.IntPtr(int(data.Maxloginattempts.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnvserver.Name = data.Name.ValueString()
	}
	if !data.Netprofile.IsNull() && !data.Netprofile.IsUnknown() {
		vpnvserver.Netprofile = data.Netprofile.ValueString()
	}
	// newname is rename-only; excluded from the add/update payload.
	if !data.Pcoipvserverprofilename.IsNull() && !data.Pcoipvserverprofilename.IsUnknown() {
		vpnvserver.Pcoipvserverprofilename = data.Pcoipvserverprofilename.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		vpnvserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Quicprofilename.IsNull() && !data.Quicprofilename.IsUnknown() {
		vpnvserver.Quicprofilename = data.Quicprofilename.ValueString()
	}
	if !data.Range.IsNull() && !data.Range.IsUnknown() {
		vpnvserver.Range = utils.IntPtr(int(data.Range.ValueInt64()))
	}
	if !data.Rdpserverprofilename.IsNull() && !data.Rdpserverprofilename.IsUnknown() {
		vpnvserver.Rdpserverprofilename = data.Rdpserverprofilename.ValueString()
	}
	if !data.Rhistate.IsNull() && !data.Rhistate.IsUnknown() {
		vpnvserver.Rhistate = data.Rhistate.ValueString()
	}
	if !data.Samesite.IsNull() && !data.Samesite.IsUnknown() {
		vpnvserver.Samesite = data.Samesite.ValueString()
	}
	if !data.Secureprivateaccess.IsNull() && !data.Secureprivateaccess.IsUnknown() {
		vpnvserver.Secureprivateaccess = data.Secureprivateaccess.ValueString()
	}
	if !data.Servicetype.IsNull() && !data.Servicetype.IsUnknown() {
		vpnvserver.Servicetype = data.Servicetype.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		vpnvserver.State = data.State.ValueString()
	}
	if !data.Tcpprofilename.IsNull() && !data.Tcpprofilename.IsUnknown() {
		vpnvserver.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Userdomains.IsNull() && !data.Userdomains.IsUnknown() {
		vpnvserver.Userdomains = data.Userdomains.ValueString()
	}
	if !data.Vserverfqdn.IsNull() && !data.Vserverfqdn.IsUnknown() {
		vpnvserver.Vserverfqdn = data.Vserverfqdn.ValueString()
	}
	if !data.Wasmmodule.IsNull() && !data.Wasmmodule.IsUnknown() {
		vpnvserver.Wasmmodule = data.Wasmmodule.ValueString()
	}
	if !data.Windowsepapluginupgrade.IsNull() && !data.Windowsepapluginupgrade.IsUnknown() {
		vpnvserver.Windowsepapluginupgrade = data.Windowsepapluginupgrade.ValueString()
	}

	return vpnvserver
}

// vpnvserverGetTheUpdatablePayloadFromThePlan builds the NITRO set/update body,
// mirroring the SDK v2 update contract: it includes ONLY the attributes that
// actually changed (and are known & non-null). It EXCLUDES every create-only
// param — the ones present in the NITRO add payload but absent from the update
// payload: servicetype, port, range, state and deploymenttype (plus the name key,
// also RequiresReplace/ForceNew). NITRO rejects any of these on set with errorcode
// 278 ("Invalid argument"). newname is rename-only (?action=rename). The bool
// return reports whether any updateable attribute changed.
func vpnvserverGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *VpnvserverResourceModel, state *VpnvserverResourceModel) (vpn.Vpnvserver, bool) {
	tflog.Debug(ctx, "In vpnvserverGetTheUpdatablePayloadFromThePlan Function")

	vpnvserver := vpn.Vpnvserver{}
	hasChange := false

	if !data.Accessrestrictedpageredirect.Equal(state.Accessrestrictedpageredirect) && !data.Accessrestrictedpageredirect.IsUnknown() && !data.Accessrestrictedpageredirect.IsNull() {
		vpnvserver.Accessrestrictedpageredirect = data.Accessrestrictedpageredirect.ValueString()
		hasChange = true
	}
	if !data.Advancedepa.Equal(state.Advancedepa) && !data.Advancedepa.IsUnknown() && !data.Advancedepa.IsNull() {
		vpnvserver.Advancedepa = data.Advancedepa.ValueString()
		hasChange = true
	}
	if !data.Appflowlog.Equal(state.Appflowlog) && !data.Appflowlog.IsUnknown() && !data.Appflowlog.IsNull() {
		vpnvserver.Appflowlog = data.Appflowlog.ValueString()
		hasChange = true
	}
	if !data.Authentication.Equal(state.Authentication) && !data.Authentication.IsUnknown() && !data.Authentication.IsNull() {
		vpnvserver.Authentication = data.Authentication.ValueString()
		hasChange = true
	}
	if !data.Authnprofile.Equal(state.Authnprofile) && !data.Authnprofile.IsUnknown() && !data.Authnprofile.IsNull() {
		vpnvserver.Authnprofile = data.Authnprofile.ValueString()
		hasChange = true
	}
	if !data.Certkeynames.Equal(state.Certkeynames) && !data.Certkeynames.IsUnknown() && !data.Certkeynames.IsNull() {
		vpnvserver.Certkeynames = data.Certkeynames.ValueString()
		hasChange = true
	}
	if !data.Cginfrahomepageredirect.Equal(state.Cginfrahomepageredirect) && !data.Cginfrahomepageredirect.IsUnknown() && !data.Cginfrahomepageredirect.IsNull() {
		vpnvserver.Cginfrahomepageredirect = data.Cginfrahomepageredirect.ValueString()
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) && !data.Comment.IsUnknown() && !data.Comment.IsNull() {
		vpnvserver.Comment = data.Comment.ValueString()
		hasChange = true
	}
	// deploymenttype is create-only (present in the NITRO add payload, absent from
	// the update payload); NITRO rejects it on set (errorcode 278). Excluded here.
	if !data.Devicecert.Equal(state.Devicecert) && !data.Devicecert.IsUnknown() && !data.Devicecert.IsNull() {
		vpnvserver.Devicecert = data.Devicecert.ValueString()
		hasChange = true
	}
	if !data.Deviceposture.Equal(state.Deviceposture) && !data.Deviceposture.IsUnknown() && !data.Deviceposture.IsNull() {
		vpnvserver.Deviceposture = data.Deviceposture.ValueString()
		hasChange = true
	}
	if !data.Doublehop.Equal(state.Doublehop) && !data.Doublehop.IsUnknown() && !data.Doublehop.IsNull() {
		vpnvserver.Doublehop = data.Doublehop.ValueString()
		hasChange = true
	}
	if !data.Downstateflush.Equal(state.Downstateflush) && !data.Downstateflush.IsUnknown() && !data.Downstateflush.IsNull() {
		vpnvserver.Downstateflush = data.Downstateflush.ValueString()
		hasChange = true
	}
	if !data.Dtls.Equal(state.Dtls) && !data.Dtls.IsUnknown() && !data.Dtls.IsNull() {
		vpnvserver.Dtls = data.Dtls.ValueString()
		hasChange = true
	}
	if !data.Failedlogintimeout.Equal(state.Failedlogintimeout) && !data.Failedlogintimeout.IsUnknown() && !data.Failedlogintimeout.IsNull() {
		vpnvserver.Failedlogintimeout = utils.IntPtr(int(data.Failedlogintimeout.ValueInt64()))
		hasChange = true
	}
	if !data.Gslbsitefqdn.Equal(state.Gslbsitefqdn) && !data.Gslbsitefqdn.IsUnknown() && !data.Gslbsitefqdn.IsNull() {
		vpnvserver.Gslbsitefqdn = data.Gslbsitefqdn.ValueString()
		hasChange = true
	}
	if !data.Httpprofilename.Equal(state.Httpprofilename) && !data.Httpprofilename.IsUnknown() && !data.Httpprofilename.IsNull() {
		vpnvserver.Httpprofilename = data.Httpprofilename.ValueString()
		hasChange = true
	}
	if !data.Icaonly.Equal(state.Icaonly) && !data.Icaonly.IsUnknown() && !data.Icaonly.IsNull() {
		vpnvserver.Icaonly = data.Icaonly.ValueString()
		hasChange = true
	}
	if !data.Icaproxysessionmigration.Equal(state.Icaproxysessionmigration) && !data.Icaproxysessionmigration.IsUnknown() && !data.Icaproxysessionmigration.IsNull() {
		vpnvserver.Icaproxysessionmigration = data.Icaproxysessionmigration.ValueString()
		hasChange = true
	}
	if !data.Icmpvsrresponse.Equal(state.Icmpvsrresponse) && !data.Icmpvsrresponse.IsUnknown() && !data.Icmpvsrresponse.IsNull() {
		vpnvserver.Icmpvsrresponse = data.Icmpvsrresponse.ValueString()
		hasChange = true
	}
	if !data.Ipset.Equal(state.Ipset) && !data.Ipset.IsUnknown() && !data.Ipset.IsNull() {
		vpnvserver.Ipset = data.Ipset.ValueString()
		hasChange = true
	}
	if !data.Ipv46.Equal(state.Ipv46) && !data.Ipv46.IsUnknown() && !data.Ipv46.IsNull() {
		vpnvserver.Ipv46 = data.Ipv46.ValueString()
		hasChange = true
	}
	if !data.L2conn.Equal(state.L2conn) && !data.L2conn.IsUnknown() && !data.L2conn.IsNull() {
		vpnvserver.L2conn = data.L2conn.ValueString()
		hasChange = true
	}
	if !data.Linuxepapluginupgrade.Equal(state.Linuxepapluginupgrade) && !data.Linuxepapluginupgrade.IsUnknown() && !data.Linuxepapluginupgrade.IsNull() {
		vpnvserver.Linuxepapluginupgrade = data.Linuxepapluginupgrade.ValueString()
		hasChange = true
	}
	if !data.Listenpolicy.Equal(state.Listenpolicy) && !data.Listenpolicy.IsUnknown() && !data.Listenpolicy.IsNull() {
		vpnvserver.Listenpolicy = data.Listenpolicy.ValueString()
		hasChange = true
	}
	if !data.Listenpriority.Equal(state.Listenpriority) && !data.Listenpriority.IsUnknown() && !data.Listenpriority.IsNull() {
		vpnvserver.Listenpriority = utils.IntPtr(int(data.Listenpriority.ValueInt64()))
		hasChange = true
	}
	if !data.Loginonce.Equal(state.Loginonce) && !data.Loginonce.IsUnknown() && !data.Loginonce.IsNull() {
		vpnvserver.Loginonce = data.Loginonce.ValueString()
		hasChange = true
	}
	if !data.Logoutonsmartcardremoval.Equal(state.Logoutonsmartcardremoval) && !data.Logoutonsmartcardremoval.IsUnknown() && !data.Logoutonsmartcardremoval.IsNull() {
		vpnvserver.Logoutonsmartcardremoval = data.Logoutonsmartcardremoval.ValueString()
		hasChange = true
	}
	if !data.Macepapluginupgrade.Equal(state.Macepapluginupgrade) && !data.Macepapluginupgrade.IsUnknown() && !data.Macepapluginupgrade.IsNull() {
		vpnvserver.Macepapluginupgrade = data.Macepapluginupgrade.ValueString()
		hasChange = true
	}
	if !data.Maxaaausers.Equal(state.Maxaaausers) && !data.Maxaaausers.IsUnknown() && !data.Maxaaausers.IsNull() {
		vpnvserver.Maxaaausers = utils.IntPtr(int(data.Maxaaausers.ValueInt64()))
		hasChange = true
	}
	if !data.Maxloginattempts.Equal(state.Maxloginattempts) && !data.Maxloginattempts.IsUnknown() && !data.Maxloginattempts.IsNull() {
		vpnvserver.Maxloginattempts = utils.IntPtr(int(data.Maxloginattempts.ValueInt64()))
		hasChange = true
	}
	if !data.Netprofile.Equal(state.Netprofile) && !data.Netprofile.IsUnknown() && !data.Netprofile.IsNull() {
		vpnvserver.Netprofile = data.Netprofile.ValueString()
		hasChange = true
	}
	if !data.Pcoipvserverprofilename.Equal(state.Pcoipvserverprofilename) && !data.Pcoipvserverprofilename.IsUnknown() && !data.Pcoipvserverprofilename.IsNull() {
		vpnvserver.Pcoipvserverprofilename = data.Pcoipvserverprofilename.ValueString()
		hasChange = true
	}
	// port is create-only (present in the NITRO add payload, absent from the update
	// payload); NITRO rejects it on set (errorcode 278). Excluded here.
	if !data.Quicprofilename.Equal(state.Quicprofilename) && !data.Quicprofilename.IsUnknown() && !data.Quicprofilename.IsNull() {
		vpnvserver.Quicprofilename = data.Quicprofilename.ValueString()
		hasChange = true
	}
	// range is create-only (present in the NITRO add payload, absent from the update
	// payload); NITRO rejects it on set (errorcode 278). Excluded here.
	if !data.Rdpserverprofilename.Equal(state.Rdpserverprofilename) && !data.Rdpserverprofilename.IsUnknown() && !data.Rdpserverprofilename.IsNull() {
		vpnvserver.Rdpserverprofilename = data.Rdpserverprofilename.ValueString()
		hasChange = true
	}
	if !data.Rhistate.Equal(state.Rhistate) && !data.Rhistate.IsUnknown() && !data.Rhistate.IsNull() {
		vpnvserver.Rhistate = data.Rhistate.ValueString()
		hasChange = true
	}
	if !data.Samesite.Equal(state.Samesite) && !data.Samesite.IsUnknown() && !data.Samesite.IsNull() {
		vpnvserver.Samesite = data.Samesite.ValueString()
		hasChange = true
	}
	if !data.Secureprivateaccess.Equal(state.Secureprivateaccess) && !data.Secureprivateaccess.IsUnknown() && !data.Secureprivateaccess.IsNull() {
		vpnvserver.Secureprivateaccess = data.Secureprivateaccess.ValueString()
		hasChange = true
	}
	// state is create-only via add (toggled through enable/disable, not accepted in
	// the NITRO update payload); NITRO rejects it on set (errorcode 278). Excluded here.
	if !data.Tcpprofilename.Equal(state.Tcpprofilename) && !data.Tcpprofilename.IsUnknown() && !data.Tcpprofilename.IsNull() {
		vpnvserver.Tcpprofilename = data.Tcpprofilename.ValueString()
		hasChange = true
	}
	if !data.Userdomains.Equal(state.Userdomains) && !data.Userdomains.IsUnknown() && !data.Userdomains.IsNull() {
		vpnvserver.Userdomains = data.Userdomains.ValueString()
		hasChange = true
	}
	if !data.Vserverfqdn.Equal(state.Vserverfqdn) && !data.Vserverfqdn.IsUnknown() && !data.Vserverfqdn.IsNull() {
		vpnvserver.Vserverfqdn = data.Vserverfqdn.ValueString()
		hasChange = true
	}
	if !data.Wasmmodule.Equal(state.Wasmmodule) && !data.Wasmmodule.IsUnknown() && !data.Wasmmodule.IsNull() {
		vpnvserver.Wasmmodule = data.Wasmmodule.ValueString()
		hasChange = true
	}
	if !data.Windowsepapluginupgrade.Equal(state.Windowsepapluginupgrade) && !data.Windowsepapluginupgrade.IsUnknown() && !data.Windowsepapluginupgrade.IsNull() {
		vpnvserver.Windowsepapluginupgrade = data.Windowsepapluginupgrade.ValueString()
		hasChange = true
	}

	return vpnvserver, hasChange
}

// vpnvserverSetAttrFromGet maps a NITRO GET response onto the resource model.
// It guards the else-branches so a configured value that NITRO omits from GET
// (omit-on-default) is preserved rather than clobbered to null; it only nulls
// an attribute that is still Unknown (a Computed attr the user never set).
func vpnvserverSetAttrFromGet(ctx context.Context, data *VpnvserverResourceModel, getResponseData map[string]interface{}) *VpnvserverResourceModel {
	tflog.Debug(ctx, "In vpnvserverSetAttrFromGet Function")

	if val, ok := getResponseData["accessrestrictedpageredirect"]; ok && val != nil {
		data.Accessrestrictedpageredirect = types.StringValue(val.(string))
	} else if data.Accessrestrictedpageredirect.IsUnknown() {
		data.Accessrestrictedpageredirect = types.StringNull()
	}
	if val, ok := getResponseData["advancedepa"]; ok && val != nil {
		data.Advancedepa = types.StringValue(val.(string))
	} else if data.Advancedepa.IsUnknown() {
		data.Advancedepa = types.StringNull()
	}
	if val, ok := getResponseData["appflowlog"]; ok && val != nil {
		data.Appflowlog = types.StringValue(val.(string))
	} else if data.Appflowlog.IsUnknown() {
		data.Appflowlog = types.StringNull()
	}
	if val, ok := getResponseData["authentication"]; ok && val != nil {
		data.Authentication = types.StringValue(val.(string))
	} else if data.Authentication.IsUnknown() {
		data.Authentication = types.StringNull()
	}
	if val, ok := getResponseData["authnprofile"]; ok && val != nil {
		data.Authnprofile = types.StringValue(val.(string))
	} else if data.Authnprofile.IsUnknown() {
		data.Authnprofile = types.StringNull()
	}
	if val, ok := getResponseData["certkeynames"]; ok && val != nil {
		data.Certkeynames = types.StringValue(val.(string))
	} else if data.Certkeynames.IsUnknown() {
		data.Certkeynames = types.StringNull()
	}
	if val, ok := getResponseData["cginfrahomepageredirect"]; ok && val != nil {
		data.Cginfrahomepageredirect = types.StringValue(val.(string))
	} else if data.Cginfrahomepageredirect.IsUnknown() {
		data.Cginfrahomepageredirect = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["deploymenttype"]; ok && val != nil {
		data.Deploymenttype = types.StringValue(val.(string))
	} else if data.Deploymenttype.IsUnknown() {
		data.Deploymenttype = types.StringNull()
	}
	if val, ok := getResponseData["devicecert"]; ok && val != nil {
		data.Devicecert = types.StringValue(val.(string))
	} else if data.Devicecert.IsUnknown() {
		data.Devicecert = types.StringNull()
	}
	if val, ok := getResponseData["deviceposture"]; ok && val != nil {
		data.Deviceposture = types.StringValue(val.(string))
	} else if data.Deviceposture.IsUnknown() {
		data.Deviceposture = types.StringNull()
	}
	if val, ok := getResponseData["doublehop"]; ok && val != nil {
		data.Doublehop = types.StringValue(val.(string))
	} else if data.Doublehop.IsUnknown() {
		data.Doublehop = types.StringNull()
	}
	if val, ok := getResponseData["downstateflush"]; ok && val != nil {
		data.Downstateflush = types.StringValue(val.(string))
	} else if data.Downstateflush.IsUnknown() {
		data.Downstateflush = types.StringNull()
	}
	if val, ok := getResponseData["dtls"]; ok && val != nil {
		data.Dtls = types.StringValue(val.(string))
	} else if data.Dtls.IsUnknown() {
		data.Dtls = types.StringNull()
	}
	if val, ok := getResponseData["failedlogintimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Failedlogintimeout = types.Int64Value(intVal)
		}
	} else if data.Failedlogintimeout.IsUnknown() {
		data.Failedlogintimeout = types.Int64Null()
	}
	if val, ok := getResponseData["gslbsitefqdn"]; ok && val != nil {
		data.Gslbsitefqdn = types.StringValue(val.(string))
	} else if data.Gslbsitefqdn.IsUnknown() {
		data.Gslbsitefqdn = types.StringNull()
	}
	if val, ok := getResponseData["httpprofilename"]; ok && val != nil {
		data.Httpprofilename = types.StringValue(val.(string))
	} else if data.Httpprofilename.IsUnknown() {
		data.Httpprofilename = types.StringNull()
	}
	if val, ok := getResponseData["icaonly"]; ok && val != nil {
		data.Icaonly = types.StringValue(val.(string))
	} else if data.Icaonly.IsUnknown() {
		data.Icaonly = types.StringNull()
	}
	if val, ok := getResponseData["icaproxysessionmigration"]; ok && val != nil {
		data.Icaproxysessionmigration = types.StringValue(val.(string))
	} else if data.Icaproxysessionmigration.IsUnknown() {
		data.Icaproxysessionmigration = types.StringNull()
	}
	if val, ok := getResponseData["icmpvsrresponse"]; ok && val != nil {
		data.Icmpvsrresponse = types.StringValue(val.(string))
	} else if data.Icmpvsrresponse.IsUnknown() {
		data.Icmpvsrresponse = types.StringNull()
	}
	if val, ok := getResponseData["ipset"]; ok && val != nil {
		data.Ipset = types.StringValue(val.(string))
	} else if data.Ipset.IsUnknown() {
		data.Ipset = types.StringNull()
	}
	if val, ok := getResponseData["ipv46"]; ok && val != nil {
		data.Ipv46 = types.StringValue(val.(string))
	} else if data.Ipv46.IsUnknown() {
		data.Ipv46 = types.StringNull()
	}
	if val, ok := getResponseData["l2conn"]; ok && val != nil {
		data.L2conn = types.StringValue(val.(string))
	} else if data.L2conn.IsUnknown() {
		data.L2conn = types.StringNull()
	}
	if val, ok := getResponseData["linuxepapluginupgrade"]; ok && val != nil {
		data.Linuxepapluginupgrade = types.StringValue(val.(string))
	} else if data.Linuxepapluginupgrade.IsUnknown() {
		data.Linuxepapluginupgrade = types.StringNull()
	}
	if val, ok := getResponseData["listenpolicy"]; ok && val != nil {
		data.Listenpolicy = types.StringValue(val.(string))
	} else if data.Listenpolicy.IsUnknown() {
		data.Listenpolicy = types.StringNull()
	}
	if val, ok := getResponseData["listenpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Listenpriority = types.Int64Value(intVal)
		}
	} else if data.Listenpriority.IsUnknown() {
		data.Listenpriority = types.Int64Null()
	}
	if val, ok := getResponseData["loginonce"]; ok && val != nil {
		data.Loginonce = types.StringValue(val.(string))
	} else if data.Loginonce.IsUnknown() {
		data.Loginonce = types.StringNull()
	}
	if val, ok := getResponseData["logoutonsmartcardremoval"]; ok && val != nil {
		data.Logoutonsmartcardremoval = types.StringValue(val.(string))
	} else if data.Logoutonsmartcardremoval.IsUnknown() {
		data.Logoutonsmartcardremoval = types.StringNull()
	}
	if val, ok := getResponseData["macepapluginupgrade"]; ok && val != nil {
		data.Macepapluginupgrade = types.StringValue(val.(string))
	} else if data.Macepapluginupgrade.IsUnknown() {
		data.Macepapluginupgrade = types.StringNull()
	}
	if val, ok := getResponseData["maxaaausers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxaaausers = types.Int64Value(intVal)
		}
	} else if data.Maxaaausers.IsUnknown() {
		data.Maxaaausers = types.Int64Null()
	}
	if val, ok := getResponseData["maxloginattempts"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxloginattempts = types.Int64Value(intVal)
		}
	} else if data.Maxloginattempts.IsUnknown() {
		data.Maxloginattempts = types.Int64Null()
	}
	// name is the key. Only adopt it from GET when the model has none (import);
	// otherwise preserve the configured/state value so a rename does not clobber it.
	if val, ok := getResponseData["name"]; ok && val != nil {
		if data.Name.IsNull() || data.Name.IsUnknown() {
			data.Name = types.StringValue(val.(string))
		}
	}
	if val, ok := getResponseData["netprofile"]; ok && val != nil {
		data.Netprofile = types.StringValue(val.(string))
	} else if data.Netprofile.IsUnknown() {
		data.Netprofile = types.StringNull()
	}
	// newname is rename-only and never returned by GET; leave it as-is.
	if val, ok := getResponseData["pcoipvserverprofilename"]; ok && val != nil {
		data.Pcoipvserverprofilename = types.StringValue(val.(string))
	} else if data.Pcoipvserverprofilename.IsUnknown() {
		data.Pcoipvserverprofilename = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else if data.Port.IsUnknown() {
		data.Port = types.Int64Null()
	}
	if val, ok := getResponseData["quicprofilename"]; ok && val != nil {
		data.Quicprofilename = types.StringValue(val.(string))
	} else if data.Quicprofilename.IsUnknown() {
		data.Quicprofilename = types.StringNull()
	}
	if val, ok := getResponseData["range"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Range = types.Int64Value(intVal)
		}
	} else if data.Range.IsUnknown() {
		data.Range = types.Int64Null()
	}
	if val, ok := getResponseData["rdpserverprofilename"]; ok && val != nil {
		data.Rdpserverprofilename = types.StringValue(val.(string))
	} else if data.Rdpserverprofilename.IsUnknown() {
		data.Rdpserverprofilename = types.StringNull()
	}
	if val, ok := getResponseData["rhistate"]; ok && val != nil {
		data.Rhistate = types.StringValue(val.(string))
	} else if data.Rhistate.IsUnknown() {
		data.Rhistate = types.StringNull()
	}
	if val, ok := getResponseData["samesite"]; ok && val != nil {
		data.Samesite = types.StringValue(val.(string))
	} else if data.Samesite.IsUnknown() {
		data.Samesite = types.StringNull()
	}
	if val, ok := getResponseData["secureprivateaccess"]; ok && val != nil {
		data.Secureprivateaccess = types.StringValue(val.(string))
	} else if data.Secureprivateaccess.IsUnknown() {
		data.Secureprivateaccess = types.StringNull()
	}
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	} else if data.Servicetype.IsUnknown() {
		data.Servicetype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else if data.State.IsUnknown() {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["tcpprofilename"]; ok && val != nil {
		data.Tcpprofilename = types.StringValue(val.(string))
	} else if data.Tcpprofilename.IsUnknown() {
		data.Tcpprofilename = types.StringNull()
	}
	if val, ok := getResponseData["userdomains"]; ok && val != nil {
		data.Userdomains = types.StringValue(val.(string))
	} else if data.Userdomains.IsUnknown() {
		data.Userdomains = types.StringNull()
	}
	if val, ok := getResponseData["vserverfqdn"]; ok && val != nil {
		data.Vserverfqdn = types.StringValue(val.(string))
	} else if data.Vserverfqdn.IsUnknown() {
		data.Vserverfqdn = types.StringNull()
	}
	if val, ok := getResponseData["wasmmodule"]; ok && val != nil {
		data.Wasmmodule = types.StringValue(val.(string))
	} else if data.Wasmmodule.IsUnknown() {
		data.Wasmmodule = types.StringNull()
	}
	if val, ok := getResponseData["windowsepapluginupgrade"]; ok && val != nil {
		data.Windowsepapluginupgrade = types.StringValue(val.(string))
	} else if data.Windowsepapluginupgrade.IsUnknown() {
		data.Windowsepapluginupgrade = types.StringNull()
	}

	// The ID tracks the live object name (needed for rename + import).
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Id = types.StringValue(val.(string))
	} else {
		data.Id = types.StringValue(data.Name.ValueString())
	}

	return data
}

// vpnvserverSetAttrFromGetForDatasource maps a NITRO GET response onto the model
// for the datasource path: it copies every returned value (nulling absent ones),
// since a datasource must reflect the live ADC configuration exactly.
func vpnvserverSetAttrFromGetForDatasource(ctx context.Context, data *VpnvserverResourceModel, getResponseData map[string]interface{}) *VpnvserverResourceModel {
	tflog.Debug(ctx, "In vpnvserverSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["accessrestrictedpageredirect"]; ok && val != nil {
		data.Accessrestrictedpageredirect = types.StringValue(val.(string))
	} else {
		data.Accessrestrictedpageredirect = types.StringNull()
	}
	if val, ok := getResponseData["advancedepa"]; ok && val != nil {
		data.Advancedepa = types.StringValue(val.(string))
	} else {
		data.Advancedepa = types.StringNull()
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
	if val, ok := getResponseData["authnprofile"]; ok && val != nil {
		data.Authnprofile = types.StringValue(val.(string))
	} else {
		data.Authnprofile = types.StringNull()
	}
	if val, ok := getResponseData["certkeynames"]; ok && val != nil {
		data.Certkeynames = types.StringValue(val.(string))
	} else {
		data.Certkeynames = types.StringNull()
	}
	if val, ok := getResponseData["cginfrahomepageredirect"]; ok && val != nil {
		data.Cginfrahomepageredirect = types.StringValue(val.(string))
	} else {
		data.Cginfrahomepageredirect = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["deploymenttype"]; ok && val != nil {
		data.Deploymenttype = types.StringValue(val.(string))
	} else {
		data.Deploymenttype = types.StringNull()
	}
	if val, ok := getResponseData["devicecert"]; ok && val != nil {
		data.Devicecert = types.StringValue(val.(string))
	} else {
		data.Devicecert = types.StringNull()
	}
	if val, ok := getResponseData["deviceposture"]; ok && val != nil {
		data.Deviceposture = types.StringValue(val.(string))
	} else {
		data.Deviceposture = types.StringNull()
	}
	if val, ok := getResponseData["doublehop"]; ok && val != nil {
		data.Doublehop = types.StringValue(val.(string))
	} else {
		data.Doublehop = types.StringNull()
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
	if val, ok := getResponseData["failedlogintimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Failedlogintimeout = types.Int64Value(intVal)
		}
	} else {
		data.Failedlogintimeout = types.Int64Null()
	}
	if val, ok := getResponseData["gslbsitefqdn"]; ok && val != nil {
		data.Gslbsitefqdn = types.StringValue(val.(string))
	} else {
		data.Gslbsitefqdn = types.StringNull()
	}
	if val, ok := getResponseData["httpprofilename"]; ok && val != nil {
		data.Httpprofilename = types.StringValue(val.(string))
	} else {
		data.Httpprofilename = types.StringNull()
	}
	if val, ok := getResponseData["icaonly"]; ok && val != nil {
		data.Icaonly = types.StringValue(val.(string))
	} else {
		data.Icaonly = types.StringNull()
	}
	if val, ok := getResponseData["icaproxysessionmigration"]; ok && val != nil {
		data.Icaproxysessionmigration = types.StringValue(val.(string))
	} else {
		data.Icaproxysessionmigration = types.StringNull()
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
	if val, ok := getResponseData["linuxepapluginupgrade"]; ok && val != nil {
		data.Linuxepapluginupgrade = types.StringValue(val.(string))
	} else {
		data.Linuxepapluginupgrade = types.StringNull()
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
	if val, ok := getResponseData["loginonce"]; ok && val != nil {
		data.Loginonce = types.StringValue(val.(string))
	} else {
		data.Loginonce = types.StringNull()
	}
	if val, ok := getResponseData["logoutonsmartcardremoval"]; ok && val != nil {
		data.Logoutonsmartcardremoval = types.StringValue(val.(string))
	} else {
		data.Logoutonsmartcardremoval = types.StringNull()
	}
	if val, ok := getResponseData["macepapluginupgrade"]; ok && val != nil {
		data.Macepapluginupgrade = types.StringValue(val.(string))
	} else {
		data.Macepapluginupgrade = types.StringNull()
	}
	if val, ok := getResponseData["maxaaausers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxaaausers = types.Int64Value(intVal)
		}
	} else {
		data.Maxaaausers = types.Int64Null()
	}
	if val, ok := getResponseData["maxloginattempts"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxloginattempts = types.Int64Value(intVal)
		}
	} else {
		data.Maxloginattempts = types.Int64Null()
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
	// newname is rename-only and never returned by GET.
	data.Newname = types.StringNull()
	if val, ok := getResponseData["pcoipvserverprofilename"]; ok && val != nil {
		data.Pcoipvserverprofilename = types.StringValue(val.(string))
	} else {
		data.Pcoipvserverprofilename = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Port = types.Int64Value(intVal)
		}
	} else {
		data.Port = types.Int64Null()
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
	if val, ok := getResponseData["rdpserverprofilename"]; ok && val != nil {
		data.Rdpserverprofilename = types.StringValue(val.(string))
	} else {
		data.Rdpserverprofilename = types.StringNull()
	}
	if val, ok := getResponseData["rhistate"]; ok && val != nil {
		data.Rhistate = types.StringValue(val.(string))
	} else {
		data.Rhistate = types.StringNull()
	}
	if val, ok := getResponseData["samesite"]; ok && val != nil {
		data.Samesite = types.StringValue(val.(string))
	} else {
		data.Samesite = types.StringNull()
	}
	if val, ok := getResponseData["secureprivateaccess"]; ok && val != nil {
		data.Secureprivateaccess = types.StringValue(val.(string))
	} else {
		data.Secureprivateaccess = types.StringNull()
	}
	if val, ok := getResponseData["servicetype"]; ok && val != nil {
		data.Servicetype = types.StringValue(val.(string))
	} else {
		data.Servicetype = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["tcpprofilename"]; ok && val != nil {
		data.Tcpprofilename = types.StringValue(val.(string))
	} else {
		data.Tcpprofilename = types.StringNull()
	}
	if val, ok := getResponseData["userdomains"]; ok && val != nil {
		data.Userdomains = types.StringValue(val.(string))
	} else {
		data.Userdomains = types.StringNull()
	}
	if val, ok := getResponseData["vserverfqdn"]; ok && val != nil {
		data.Vserverfqdn = types.StringValue(val.(string))
	} else {
		data.Vserverfqdn = types.StringNull()
	}
	if val, ok := getResponseData["wasmmodule"]; ok && val != nil {
		data.Wasmmodule = types.StringValue(val.(string))
	} else {
		data.Wasmmodule = types.StringNull()
	}
	if val, ok := getResponseData["windowsepapluginupgrade"]; ok && val != nil {
		data.Windowsepapluginupgrade = types.StringValue(val.(string))
	} else {
		data.Windowsepapluginupgrade = types.StringNull()
	}

	// Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
