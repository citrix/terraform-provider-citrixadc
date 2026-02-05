package nsip

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

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

// NsipResourceModel describes the resource data model.
type NsipResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Advertiseondefaultpartition types.String `tfsdk:"advertiseondefaultpartition"`
	Arp                         types.String `tfsdk:"arp"`
	Arpowner                    types.Int64  `tfsdk:"arpowner"`
	Arpresponse                 types.String `tfsdk:"arpresponse"`
	Bgp                         types.String `tfsdk:"bgp"`
	Decrementttl                types.String `tfsdk:"decrementttl"`
	Dynamicrouting              types.String `tfsdk:"dynamicrouting"`
	Ftp                         types.String `tfsdk:"ftp"`
	Gui                         types.String `tfsdk:"gui"`
	Hostroute                   types.String `tfsdk:"hostroute"`
	Hostrtgw                    types.String `tfsdk:"hostrtgw"`
	Icmp                        types.String `tfsdk:"icmp"`
	Icmpresponse                types.String `tfsdk:"icmpresponse"`
	Ipaddress                   types.String `tfsdk:"ipaddress"`
	Metric                      types.Int64  `tfsdk:"metric"`
	Mgmtaccess                  types.String `tfsdk:"mgmtaccess"`
	Mptcpadvertise              types.String `tfsdk:"mptcpadvertise"`
	Netmask                     types.String `tfsdk:"netmask"`
	Networkroute                types.String `tfsdk:"networkroute"`
	Ospf                        types.String `tfsdk:"ospf"`
	Ospfarea                    types.Int64  `tfsdk:"ospfarea"`
	Ospflsatype                 types.String `tfsdk:"ospflsatype"`
	Ownerdownresponse           types.String `tfsdk:"ownerdownresponse"`
	Ownernode                   types.Int64  `tfsdk:"ownernode"`
	Restrictaccess              types.String `tfsdk:"restrictaccess"`
	Rip                         types.String `tfsdk:"rip"`
	Snmp                        types.String `tfsdk:"snmp"`
	Ssh                         types.String `tfsdk:"ssh"`
	State                       types.String `tfsdk:"state"`
	Tag                         types.Int64  `tfsdk:"tag"`
	Td                          types.Int64  `tfsdk:"td"`
	Telnet                      types.String `tfsdk:"telnet"`
	Type                        types.String `tfsdk:"type"`
	Vrid                        types.Int64  `tfsdk:"vrid"`
	Vserver                     types.String `tfsdk:"vserver"`
	Vserverrhilevel             types.String `tfsdk:"vserverrhilevel"`
}

func (r *NsipResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsip resource.",
			},
			"advertiseondefaultpartition": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Advertise VIPs from Shared VLAN on Default Partition.",
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Respond to ARP requests for this IP address.",
			},
			"arpowner": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(255),
				Description: "The arp owner in a Cluster for this IP address. It can vary from 0 to 31.",
			},
			"arpresponse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("5"),
				Description: "Respond to ARP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP. Available settings function as follows:\n\n* NONE - The Citrix ADC responds to any ARP request for the VIP address, irrespective of the states of the virtual servers associated with the address.\n* ONE VSERVER - The Citrix ADC responds to any ARP request for the VIP address if at least one of the associated virtual servers is in UP state.\n* ALL VSERVER - The Citrix ADC responds to any ARP request for the VIP address if all of the associated virtual servers are in UP state.",
			},
			"bgp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Use this option to enable or disable BGP on this IP address for the entity.",
			},
			"decrementttl": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Decrement TTL by 1 when ENABLED.This setting is applicable only for UDP traffic.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow dynamic routing on this IP address. Specific to Subnet IP (SNIP) address.",
			},
			"ftp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Allow File Transfer Protocol (FTP) access to this IP address.",
			},
			"gui": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Allow graphical user interface (GUI) access to this IP address.",
			},
			"hostroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to push the VIP to ZebOS routing table for Kernel route redistribution through dynamic routing protocols",
			},
			"hostrtgw": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("-1"),
				Description: "IP address of the gateway of the route for this VIP address.",
			},
			"icmp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Respond to ICMP requests for this IP address.",
			},
			"icmpresponse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("5"),
				Description: "Respond to ICMP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP. Available settings function as follows:\n* NONE - The Citrix ADC responds to any ICMP request for the VIP address, irrespective of the states of the virtual servers associated with the address.\n* ONE VSERVER - The Citrix ADC responds to any ICMP request for the VIP address if at least one of the associated virtual servers is in UP state.\n* ALL VSERVER - The Citrix ADC responds to any ICMP request for the VIP address if all of the associated virtual servers are in UP state.\n* VSVR_CNTRLD - The behavior depends on the ICMP VSERVER RESPONSE setting on all the associated virtual servers.\n\nThe following settings can be made for the ICMP VSERVER RESPONSE parameter on a virtual server:\n* If you set ICMP VSERVER RESPONSE to PASSIVE on all virtual servers, Citrix ADC always responds.\n* If you set ICMP VSERVER RESPONSE to ACTIVE on all virtual servers, Citrix ADC responds if even one virtual server is UP.\n* When you set ICMP VSERVER RESPONSE to ACTIVE on some and PASSIVE on others, Citrix ADC responds if even one virtual server set to ACTIVE is UP.",
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "IPv4 address to create on the Citrix ADC. Cannot be changed after the IP address is created.",
			},
			"metric": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value to add to or subtract from the cost of the route advertised for the VIP address.",
			},
			"mgmtaccess": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow access to management applications on this IP address.",
			},
			"mptcpadvertise": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, this IP will be advertised by Citrix ADC to MPTCP enabled clients as part of ADD_ADDR option.",
			},
			"netmask": schema.StringAttribute{
				Required:    true,
				Description: "Subnet mask associated with the IP address.",
			},
			"networkroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to push the SNIP subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol.",
			},
			"ospf": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Use this option to enable or disable OSPF on this IP address for the entity.",
			},
			"ospfarea": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(-1),
				Description: "ID of the area in which the type1 link-state advertisements (LSAs) are to be advertised for this virtual IP (VIP)  address by the OSPF protocol running on the Citrix ADC.  When this parameter is not set, the VIP is advertised on all areas.",
			},
			"ospflsatype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("TYPE5"),
				Description: "Type of LSAs to be used by the OSPF protocol, running on the Citrix ADC, for advertising the route for this VIP address.",
			},
			"ownerdownresponse": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "in cluster system, if the owner node is down, whether should it respond to icmp/arp",
			},
			"ownernode": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(255),
				Description: "The owner node in a Cluster for this IP address. Owner node can vary from 0 to 31. If ownernode is not specified then the IP is treated as Striped IP.",
			},
			"restrictaccess": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Block access to nonmanagement applications on this IP. This option is applicable for MIPs, SNIPs, and NSIP, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system.",
			},
			"rip": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Use this option to enable or disable RIP on this IP address for the entity.",
			},
			"snmp": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Allow Simple Network Management Protocol (SNMP) access to this IP address.",
			},
			"ssh": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Allow secure shell (SSH) access to this IP address.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable the IP address.",
			},
			"tag": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Tag value for the network/host route associated with this IP.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. TD id 4095 is used reserved for  LSN use",
			},
			"telnet": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Allow Telnet access to this IP address.",
			},
			"type": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("SNIP"),
				Description: "Type of the IP address to create on the Citrix ADC. Cannot be changed after the IP address is created. The following are the different types of Citrix ADC owned IP addresses:\n* A Subnet IP (SNIP) address is used by the Citrix ADC to communicate with the servers. The Citrix ADC also uses the subnet IP address when generating its own packets, such as packets related to dynamic routing protocols, or to send monitor probes to check the health of the servers.\n* A Virtual IP (VIP) address is the IP address associated with a virtual server. It is the IP address to which clients connect. An appliance managing a wide range of traffic may have many VIPs configured. Some of the attributes of the VIP address are customized to meet the requirements of the virtual server.\n* A GSLB site IP (GSLBIP) address is associated with a GSLB site. It is not mandatory to specify a GSLBIP address when you initially configure the Citrix ADC. A GSLBIP address is used only when you create a GSLB site.\n* A Cluster IP (CLIP) address is the management address of the cluster. All cluster configurations must be performed by accessing the cluster through this IP address.",
			},
			"vrid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Use this option to set (enable or disable) the virtual server attribute for this IP address.",
			},
			"vserverrhilevel": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ONE_VSERVER"),
				Description: "Advertise the route for the Virtual IP (VIP) address on the basis of the state of the virtual servers associated with that VIP.\n* NONE - Advertise the route for the VIP address, regardless of the state of the virtual servers associated with the address.\n* ONE VSERVER - Advertise the route for the VIP address if at least one of the associated virtual servers is in UP state.\n* ALL VSERVER - Advertise the route for the VIP address if all of the associated virtual servers are in UP state.\n* VSVR_CNTRLD - Advertise the route for the VIP address according to the  RHIstate (RHI STATE) parameter setting on all the associated virtual servers of the VIP address along with their states.\n\nWhen Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:\n * If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address.\n * If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state.\n *If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state.",
			},
		},
	}
}

func nsipGetThePayloadFromtheConfig(ctx context.Context, data *NsipResourceModel) ns.Nsip {
	tflog.Debug(ctx, "In nsipGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsip := ns.Nsip{}
	if !data.Advertiseondefaultpartition.IsNull() {
		nsip.Advertiseondefaultpartition = data.Advertiseondefaultpartition.ValueString()
	}
	if !data.Arp.IsNull() {
		nsip.Arp = data.Arp.ValueString()
	}
	if !data.Arpowner.IsNull() {
		nsip.Arpowner = utils.IntPtr(int(data.Arpowner.ValueInt64()))
	}
	if !data.Arpresponse.IsNull() {
		nsip.Arpresponse = data.Arpresponse.ValueString()
	}
	if !data.Bgp.IsNull() {
		nsip.Bgp = data.Bgp.ValueString()
	}
	if !data.Decrementttl.IsNull() {
		nsip.Decrementttl = data.Decrementttl.ValueString()
	}
	if !data.Dynamicrouting.IsNull() {
		nsip.Dynamicrouting = data.Dynamicrouting.ValueString()
	}
	if !data.Ftp.IsNull() {
		nsip.Ftp = data.Ftp.ValueString()
	}
	if !data.Gui.IsNull() {
		nsip.Gui = data.Gui.ValueString()
	}
	if !data.Hostroute.IsNull() {
		nsip.Hostroute = data.Hostroute.ValueString()
	}
	if !data.Hostrtgw.IsNull() {
		nsip.Hostrtgw = data.Hostrtgw.ValueString()
	}
	if !data.Icmp.IsNull() {
		nsip.Icmp = data.Icmp.ValueString()
	}
	if !data.Icmpresponse.IsNull() {
		nsip.Icmpresponse = data.Icmpresponse.ValueString()
	}
	if !data.Ipaddress.IsNull() {
		nsip.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Metric.IsNull() {
		nsip.Metric = utils.IntPtr(int(data.Metric.ValueInt64()))
	}
	if !data.Mgmtaccess.IsNull() {
		nsip.Mgmtaccess = data.Mgmtaccess.ValueString()
	}
	if !data.Mptcpadvertise.IsNull() {
		nsip.Mptcpadvertise = data.Mptcpadvertise.ValueString()
	}
	if !data.Netmask.IsNull() {
		nsip.Netmask = data.Netmask.ValueString()
	}
	if !data.Networkroute.IsNull() {
		nsip.Networkroute = data.Networkroute.ValueString()
	}
	if !data.Ospf.IsNull() {
		nsip.Ospf = data.Ospf.ValueString()
	}
	if !data.Ospfarea.IsNull() {
		nsip.Ospfarea = utils.IntPtr(int(data.Ospfarea.ValueInt64()))
	}
	if !data.Ospflsatype.IsNull() {
		nsip.Ospflsatype = data.Ospflsatype.ValueString()
	}
	if !data.Ownerdownresponse.IsNull() {
		nsip.Ownerdownresponse = data.Ownerdownresponse.ValueString()
	}
	if !data.Ownernode.IsNull() {
		nsip.Ownernode = utils.IntPtr(int(data.Ownernode.ValueInt64()))
	}
	if !data.Restrictaccess.IsNull() {
		nsip.Restrictaccess = data.Restrictaccess.ValueString()
	}
	if !data.Rip.IsNull() {
		nsip.Rip = data.Rip.ValueString()
	}
	if !data.Snmp.IsNull() {
		nsip.Snmp = data.Snmp.ValueString()
	}
	if !data.Ssh.IsNull() {
		nsip.Ssh = data.Ssh.ValueString()
	}
	if !data.State.IsNull() {
		nsip.State = data.State.ValueString()
	}
	if !data.Tag.IsNull() {
		nsip.Tag = utils.IntPtr(int(data.Tag.ValueInt64()))
	}
	if !data.Td.IsNull() {
		nsip.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Telnet.IsNull() {
		nsip.Telnet = data.Telnet.ValueString()
	}
	if !data.Type.IsNull() {
		nsip.Type = data.Type.ValueString()
	}
	if !data.Vrid.IsNull() {
		nsip.Vrid = utils.IntPtr(int(data.Vrid.ValueInt64()))
	}
	if !data.Vserver.IsNull() {
		nsip.Vserver = data.Vserver.ValueString()
	}
	if !data.Vserverrhilevel.IsNull() {
		nsip.Vserverrhilevel = data.Vserverrhilevel.ValueString()
	}

	return nsip
}

func nsipSetAttrFromGet(ctx context.Context, data *NsipResourceModel, getResponseData map[string]interface{}) *NsipResourceModel {
	tflog.Debug(ctx, "In nsipSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["advertiseondefaultpartition"]; ok && val != nil {
		data.Advertiseondefaultpartition = types.StringValue(val.(string))
	} else {
		data.Advertiseondefaultpartition = types.StringNull()
	}
	if val, ok := getResponseData["arp"]; ok && val != nil {
		data.Arp = types.StringValue(val.(string))
	} else {
		data.Arp = types.StringNull()
	}
	if val, ok := getResponseData["arpowner"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Arpowner = types.Int64Value(intVal)
		}
	} else {
		data.Arpowner = types.Int64Null()
	}
	if val, ok := getResponseData["arpresponse"]; ok && val != nil {
		data.Arpresponse = types.StringValue(val.(string))
	} else {
		data.Arpresponse = types.StringNull()
	}
	if val, ok := getResponseData["bgp"]; ok && val != nil {
		data.Bgp = types.StringValue(val.(string))
	} else {
		data.Bgp = types.StringNull()
	}
	if val, ok := getResponseData["decrementttl"]; ok && val != nil {
		data.Decrementttl = types.StringValue(val.(string))
	} else {
		data.Decrementttl = types.StringNull()
	}
	if val, ok := getResponseData["dynamicrouting"]; ok && val != nil {
		data.Dynamicrouting = types.StringValue(val.(string))
	} else {
		data.Dynamicrouting = types.StringNull()
	}
	if val, ok := getResponseData["ftp"]; ok && val != nil {
		data.Ftp = types.StringValue(val.(string))
	} else {
		data.Ftp = types.StringNull()
	}
	if val, ok := getResponseData["gui"]; ok && val != nil {
		data.Gui = types.StringValue(val.(string))
	} else {
		data.Gui = types.StringNull()
	}
	if val, ok := getResponseData["hostroute"]; ok && val != nil {
		data.Hostroute = types.StringValue(val.(string))
	} else {
		data.Hostroute = types.StringNull()
	}
	if val, ok := getResponseData["hostrtgw"]; ok && val != nil {
		data.Hostrtgw = types.StringValue(val.(string))
	} else {
		data.Hostrtgw = types.StringNull()
	}
	if val, ok := getResponseData["icmp"]; ok && val != nil {
		data.Icmp = types.StringValue(val.(string))
	} else {
		data.Icmp = types.StringNull()
	}
	if val, ok := getResponseData["icmpresponse"]; ok && val != nil {
		data.Icmpresponse = types.StringValue(val.(string))
	} else {
		data.Icmpresponse = types.StringNull()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["metric"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Metric = types.Int64Value(intVal)
		}
	} else {
		data.Metric = types.Int64Null()
	}
	if val, ok := getResponseData["mgmtaccess"]; ok && val != nil {
		data.Mgmtaccess = types.StringValue(val.(string))
	} else {
		data.Mgmtaccess = types.StringNull()
	}
	if val, ok := getResponseData["mptcpadvertise"]; ok && val != nil {
		data.Mptcpadvertise = types.StringValue(val.(string))
	} else {
		data.Mptcpadvertise = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["networkroute"]; ok && val != nil {
		data.Networkroute = types.StringValue(val.(string))
	} else {
		data.Networkroute = types.StringNull()
	}
	if val, ok := getResponseData["ospf"]; ok && val != nil {
		data.Ospf = types.StringValue(val.(string))
	} else {
		data.Ospf = types.StringNull()
	}
	if val, ok := getResponseData["ospfarea"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ospfarea = types.Int64Value(intVal)
		}
	} else {
		data.Ospfarea = types.Int64Null()
	}
	if val, ok := getResponseData["ospflsatype"]; ok && val != nil {
		data.Ospflsatype = types.StringValue(val.(string))
	} else {
		data.Ospflsatype = types.StringNull()
	}
	if val, ok := getResponseData["ownerdownresponse"]; ok && val != nil {
		data.Ownerdownresponse = types.StringValue(val.(string))
	} else {
		data.Ownerdownresponse = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else {
		data.Ownernode = types.Int64Null()
	}
	if val, ok := getResponseData["restrictaccess"]; ok && val != nil {
		data.Restrictaccess = types.StringValue(val.(string))
	} else {
		data.Restrictaccess = types.StringNull()
	}
	if val, ok := getResponseData["rip"]; ok && val != nil {
		data.Rip = types.StringValue(val.(string))
	} else {
		data.Rip = types.StringNull()
	}
	if val, ok := getResponseData["snmp"]; ok && val != nil {
		data.Snmp = types.StringValue(val.(string))
	} else {
		data.Snmp = types.StringNull()
	}
	if val, ok := getResponseData["ssh"]; ok && val != nil {
		data.Ssh = types.StringValue(val.(string))
	} else {
		data.Ssh = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}
	if val, ok := getResponseData["tag"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Tag = types.Int64Value(intVal)
		}
	} else {
		data.Tag = types.Int64Null()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["telnet"]; ok && val != nil {
		data.Telnet = types.StringValue(val.(string))
	} else {
		data.Telnet = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["vrid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vrid = types.Int64Value(intVal)
		}
	} else {
		data.Vrid = types.Int64Null()
	}
	if val, ok := getResponseData["vserver"]; ok && val != nil {
		data.Vserver = types.StringValue(val.(string))
	} else {
		data.Vserver = types.StringNull()
	}
	if val, ok := getResponseData["vserverrhilevel"]; ok && val != nil {
		data.Vserverrhilevel = types.StringValue(val.(string))
	} else {
		data.Vserverrhilevel = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s,%d", data.Ipaddress.ValueString(), data.Td.ValueInt64()))

	return data
}
