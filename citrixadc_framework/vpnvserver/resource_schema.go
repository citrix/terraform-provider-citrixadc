package vpnvserver

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

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
			"accessrestrictedpageredirect": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("CDN"),
				Description: "By default, an access restricted page hosted on secure private access CDN is displayed when a restricted app is accessed. The setting can be changed to NS to display the access restricted page hosted on the gateway or OFF to not display any access restricted page.",
			},
			"advancedepa": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This option tells whether advanced EPA is enabled on this virtual server",
			},
			"appflowlog": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Log AppFlow records that contain standard NetFlow or IPFIX information, such as time stamps for the beginning and end of a flow, packet count, and byte count. Also log records that contain application-level information, such as HTTP web addresses, HTTP request methods and response status codes, server response time, and latency.",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
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
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "When client requests ShareFile resources and Citrix Gateway detects that the user is unauthenticated or the user session has expired, disabling this option takes the user to the originally requested ShareFile resource after authentication (instead of taking the user to the default VPN home page)",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the virtual server.",
			},
			"deploymenttype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("5"),
				Description: "0",
			},
			"devicecert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates whether device certificate check as a part of EPA is on or off.",
			},
			"deviceposture": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable device posture",
			},
			"doublehop": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Use the Citrix Gateway appliance in a double-hop configuration. A double-hop deployment provides an extra layer of security for the internal network by using three firewalls to divide the DMZ into two stages. Such a deployment can have one appliance in the DMZ and one appliance in the secure network.",
			},
			"downstateflush": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Close existing connections when the virtual server is marked DOWN, which means the server might have timed out. Disconnecting existing connections frees resources and in certain cases speeds recovery of overloaded load balancing setups. Enable this setting on servers in which the connections can safely be closed when they are marked DOWN.  Do not enable DOWN state flush on servers that must complete their transactions.",
			},
			"dtls": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "This option starts/stops the turn service on the vserver",
			},
			"failedlogintimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of minutes an account will be locked if user exceeds maximum permissible attempts",
			},
			"httpprofilename": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("nshttp_default_strict_validation"),
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
				Default:     stringdefault.StaticString("none"),
				Description: "String specifying the listen policy for the Citrix Gateway virtual server. Can be either a named expression or an expression. The Citrix Gateway virtual server processes only the traffic for which the expression evaluates to true.",
			},
			"listenpriority": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(101),
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
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the Citrix Gateway virtual server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my server\" or 'my server').",
			},
			"pcoipvserverprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the PCoIP vserver profile associated with the vserver.",
			},
			"port": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "TCP port on which the virtual server listens.",
			},
			"quicprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the QUIC profile to assign to this virtual server.",
			},
			"range": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Range of Citrix Gateway virtual server IP addresses. The consecutively numbered range of IP addresses begins with the address specified by the IP Address parameter.\nIn the configuration utility, select Network VServer to enter a range.",
			},
			"rdpserverprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the RDP server profile associated with the vserver.",
			},
			"rhistate": schema.StringAttribute{
				Optional:    true,
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
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Configure secure private access",
			},
			"servicetype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("SSL"),
				Description: "Protocol used by the Citrix Gateway virtual server.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
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

func vpnvserverGetThePayloadFromtheConfig(ctx context.Context, data *VpnvserverResourceModel) vpn.Vpnvserver {
	tflog.Debug(ctx, "In vpnvserverGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnvserver := vpn.Vpnvserver{}
	if !data.Accessrestrictedpageredirect.IsNull() {
		vpnvserver.Accessrestrictedpageredirect = data.Accessrestrictedpageredirect.ValueString()
	}
	if !data.Advancedepa.IsNull() {
		vpnvserver.Advancedepa = data.Advancedepa.ValueString()
	}
	if !data.Appflowlog.IsNull() {
		vpnvserver.Appflowlog = data.Appflowlog.ValueString()
	}
	if !data.Authentication.IsNull() {
		vpnvserver.Authentication = data.Authentication.ValueString()
	}
	if !data.Authnprofile.IsNull() {
		vpnvserver.Authnprofile = data.Authnprofile.ValueString()
	}
	if !data.Certkeynames.IsNull() {
		vpnvserver.Certkeynames = data.Certkeynames.ValueString()
	}
	if !data.Cginfrahomepageredirect.IsNull() {
		vpnvserver.Cginfrahomepageredirect = data.Cginfrahomepageredirect.ValueString()
	}
	if !data.Comment.IsNull() {
		vpnvserver.Comment = data.Comment.ValueString()
	}
	if !data.Deploymenttype.IsNull() {
		vpnvserver.Deploymenttype = data.Deploymenttype.ValueString()
	}
	if !data.Devicecert.IsNull() {
		vpnvserver.Devicecert = data.Devicecert.ValueString()
	}
	if !data.Deviceposture.IsNull() {
		vpnvserver.Deviceposture = data.Deviceposture.ValueString()
	}
	if !data.Doublehop.IsNull() {
		vpnvserver.Doublehop = data.Doublehop.ValueString()
	}
	if !data.Downstateflush.IsNull() {
		vpnvserver.Downstateflush = data.Downstateflush.ValueString()
	}
	if !data.Dtls.IsNull() {
		vpnvserver.Dtls = data.Dtls.ValueString()
	}
	if !data.Failedlogintimeout.IsNull() {
		vpnvserver.Failedlogintimeout = utils.IntPtr(int(data.Failedlogintimeout.ValueInt64()))
	}
	if !data.Httpprofilename.IsNull() {
		vpnvserver.Httpprofilename = data.Httpprofilename.ValueString()
	}
	if !data.Icaonly.IsNull() {
		vpnvserver.Icaonly = data.Icaonly.ValueString()
	}
	if !data.Icaproxysessionmigration.IsNull() {
		vpnvserver.Icaproxysessionmigration = data.Icaproxysessionmigration.ValueString()
	}
	if !data.Icmpvsrresponse.IsNull() {
		vpnvserver.Icmpvsrresponse = data.Icmpvsrresponse.ValueString()
	}
	if !data.Ipset.IsNull() {
		vpnvserver.Ipset = data.Ipset.ValueString()
	}
	if !data.Ipv46.IsNull() {
		vpnvserver.Ipv46 = data.Ipv46.ValueString()
	}
	if !data.L2conn.IsNull() {
		vpnvserver.L2conn = data.L2conn.ValueString()
	}
	if !data.Linuxepapluginupgrade.IsNull() {
		vpnvserver.Linuxepapluginupgrade = data.Linuxepapluginupgrade.ValueString()
	}
	if !data.Listenpolicy.IsNull() {
		vpnvserver.Listenpolicy = data.Listenpolicy.ValueString()
	}
	if !data.Listenpriority.IsNull() {
		vpnvserver.Listenpriority = utils.IntPtr(int(data.Listenpriority.ValueInt64()))
	}
	if !data.Loginonce.IsNull() {
		vpnvserver.Loginonce = data.Loginonce.ValueString()
	}
	if !data.Logoutonsmartcardremoval.IsNull() {
		vpnvserver.Logoutonsmartcardremoval = data.Logoutonsmartcardremoval.ValueString()
	}
	if !data.Macepapluginupgrade.IsNull() {
		vpnvserver.Macepapluginupgrade = data.Macepapluginupgrade.ValueString()
	}
	if !data.Maxaaausers.IsNull() {
		vpnvserver.Maxaaausers = utils.IntPtr(int(data.Maxaaausers.ValueInt64()))
	}
	if !data.Maxloginattempts.IsNull() {
		vpnvserver.Maxloginattempts = utils.IntPtr(int(data.Maxloginattempts.ValueInt64()))
	}
	if !data.Name.IsNull() {
		vpnvserver.Name = data.Name.ValueString()
	}
	if !data.Netprofile.IsNull() {
		vpnvserver.Netprofile = data.Netprofile.ValueString()
	}
	if !data.Newname.IsNull() {
		vpnvserver.Newname = data.Newname.ValueString()
	}
	if !data.Pcoipvserverprofilename.IsNull() {
		vpnvserver.Pcoipvserverprofilename = data.Pcoipvserverprofilename.ValueString()
	}
	if !data.Port.IsNull() {
		vpnvserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}
	if !data.Quicprofilename.IsNull() {
		vpnvserver.Quicprofilename = data.Quicprofilename.ValueString()
	}
	if !data.Range.IsNull() {
		vpnvserver.Range = utils.IntPtr(int(data.Range.ValueInt64()))
	}
	if !data.Rdpserverprofilename.IsNull() {
		vpnvserver.Rdpserverprofilename = data.Rdpserverprofilename.ValueString()
	}
	if !data.Rhistate.IsNull() {
		vpnvserver.Rhistate = data.Rhistate.ValueString()
	}
	if !data.Samesite.IsNull() {
		vpnvserver.Samesite = data.Samesite.ValueString()
	}
	if !data.Secureprivateaccess.IsNull() {
		vpnvserver.Secureprivateaccess = data.Secureprivateaccess.ValueString()
	}
	if !data.Servicetype.IsNull() {
		vpnvserver.Servicetype = data.Servicetype.ValueString()
	}
	if !data.State.IsNull() {
		vpnvserver.State = data.State.ValueString()
	}
	if !data.Tcpprofilename.IsNull() {
		vpnvserver.Tcpprofilename = data.Tcpprofilename.ValueString()
	}
	if !data.Userdomains.IsNull() {
		vpnvserver.Userdomains = data.Userdomains.ValueString()
	}
	if !data.Vserverfqdn.IsNull() {
		vpnvserver.Vserverfqdn = data.Vserverfqdn.ValueString()
	}
	if !data.Windowsepapluginupgrade.IsNull() {
		vpnvserver.Windowsepapluginupgrade = data.Windowsepapluginupgrade.ValueString()
	}

	return vpnvserver
}

func vpnvserverSetAttrFromGet(ctx context.Context, data *VpnvserverResourceModel, getResponseData map[string]interface{}) *VpnvserverResourceModel {
	tflog.Debug(ctx, "In vpnvserverSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
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
	if val, ok := getResponseData["windowsepapluginupgrade"]; ok && val != nil {
		data.Windowsepapluginupgrade = types.StringValue(val.(string))
	} else {
		data.Windowsepapluginupgrade = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
