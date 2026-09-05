package nsip

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsipDataSourceModel is the data-source-specific model, decoupled from
// NsipResourceModel.
//
// A data source is a pure read surface, so it exposes the read/write attributes
// (as Computed outputs) AND the read-only runtime/status attributes the resource
// deliberately omits (flags, freeports, iptype, VIP vserver counters, ...), all
// of which are populated only from a GET.
type NsipDataSourceModel struct {
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

	// Read-only (GET-only) runtime/status attributes from the NITRO read-only
	// set. Never settable; populated from GET.
	Flags                    types.Int64  `tfsdk:"flags"`
	Hostrtgwact              types.String `tfsdk:"hostrtgwact"`
	Ospfareaval              types.Int64  `tfsdk:"ospfareaval"`
	Viprtadv2bsd             types.Bool   `tfsdk:"viprtadv2bsd"`
	Vipvsercount             types.Int64  `tfsdk:"vipvsercount"`
	Vipvserdowncount         types.Int64  `tfsdk:"vipvserdowncount"`
	Vipvsrvrrhiactivecount   types.Int64  `tfsdk:"vipvsrvrrhiactivecount"`
	Vipvsrvrrhiactiveupcount types.Int64  `tfsdk:"vipvsrvrrhiactiveupcount"`
	Freeports                types.Int64  `tfsdk:"freeports"`
	Iptype                   types.List   `tfsdk:"iptype"`
	Operationalarpowner      types.Int64  `tfsdk:"operationalarpowner"`
}

func NsipDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"advertiseondefaultpartition": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise VIPs from Shared VLAN on Default Partition.",
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to ARP requests for this IP address.",
			},
			"arpowner": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The arp owner in a Cluster for this IP address. It can vary from 0 to 31.",
			},
			"arpresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to ARP requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP. Available settings function as follows:\n\n* NONE - The Citrix ADC responds to any ARP request for the VIP address, irrespective of the states of the virtual servers associated with the address.\n* ONE VSERVER - The Citrix ADC responds to any ARP request for the VIP address if at least one of the associated virtual servers is in UP state.\n* ALL VSERVER - The Citrix ADC responds to any ARP request for the VIP address if all of the associated virtual servers are in UP state.",
			},
			"bgp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to enable or disable BGP on this IP address for the entity.",
			},
			"decrementttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Decrement TTL by 1 when ENABLED.This setting is applicable only for UDP traffic.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow dynamic routing on this IP address. Specific to Subnet IP (SNIP) address.",
			},
			"ftp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow File Transfer Protocol (FTP) access to this IP address.",
			},
			"gui": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow graphical user interface (GUI) access to this IP address.",
			},
			"hostroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to push the VIP to ZebOS routing table for Kernel route redistribution through dynamic routing protocols",
			},
			"hostrtgw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the gateway of the route for this VIP address.",
			},
			"icmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to ICMP requests for this IP address.",
			},
			"icmpresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Computed:    true,
				Description: "Allow access to management applications on this IP address.",
			},
			"mptcpadvertise": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, this IP will be advertised by Citrix ADC to MPTCP enabled clients as part of ADD_ADDR option.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet mask associated with the IP address.",
			},
			"networkroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to push the SNIP subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol.",
			},
			"ospf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to enable or disable OSPF on this IP address for the entity.",
			},
			"ospfarea": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the area in which the type1 link-state advertisements (LSAs) are to be advertised for this virtual IP (VIP)  address by the OSPF protocol running on the Citrix ADC.  When this parameter is not set, the VIP is advertised on all areas.",
			},
			"ospflsatype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of LSAs to be used by the OSPF protocol, running on the Citrix ADC, for advertising the route for this VIP address.",
			},
			"ownerdownresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "in cluster system, if the owner node is down, whether should it respond to icmp/arp",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node in a Cluster for this IP address. Owner node can vary from 0 to 31. If ownernode is not specified then the IP is treated as Striped IP.",
			},
			"restrictaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Block access to nonmanagement applications on this IP. This option is applicable for MIPs, SNIPs, and NSIP, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system.",
			},
			"rip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to enable or disable RIP on this IP address for the entity.",
			},
			"snmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Simple Network Management Protocol (SNMP) access to this IP address.",
			},
			"ssh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow secure shell (SSH) access to this IP address.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the IP address.",
			},
			"tag": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Tag value for the network/host route associated with this IP.",
			},
			"td": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. TD id 4095 is used reserved for  LSN use",
			},
			"telnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Telnet access to this IP address.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the IP address to create on the Citrix ADC. Cannot be changed after the IP address is created. The following are the different types of Citrix ADC owned IP addresses:\n* A Subnet IP (SNIP) address is used by the Citrix ADC to communicate with the servers. The Citrix ADC also uses the subnet IP address when generating its own packets, such as packets related to dynamic routing protocols, or to send monitor probes to check the health of the servers.\n* A Virtual IP (VIP) address is the IP address associated with a virtual server. It is the IP address to which clients connect. An appliance managing a wide range of traffic may have many VIPs configured. Some of the attributes of the VIP address are customized to meet the requirements of the virtual server.\n* A GSLB site IP (GSLBIP) address is associated with a GSLB site. It is not mandatory to specify a GSLBIP address when you initially configure the Citrix ADC. A GSLBIP address is used only when you create a GSLB site.\n* A Cluster IP (CLIP) address is the management address of the cluster. All cluster configurations must be performed by accessing the cluster through this IP address.",
			},
			"vrid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this option to set (enable or disable) the virtual server attribute for this IP address.",
			},
			"vserverrhilevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise the route for the Virtual IP (VIP) address on the basis of the state of the virtual servers associated with that VIP.\n* NONE - Advertise the route for the VIP address, regardless of the state of the virtual servers associated with the address.\n* ONE VSERVER - Advertise the route for the VIP address if at least one of the associated virtual servers is in UP state.\n* ALL VSERVER - Advertise the route for the VIP address if all of the associated virtual servers are in UP state.\n* VSVR_CNTRLD - Advertise the route for the VIP address according to the  RHIstate (RHI STATE) parameter setting on all the associated virtual servers of the VIP address along with their states.\n\nWhen Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:\n * If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address.\n * If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state.\n *If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state.",
			},

			// Read-only (GET-only) runtime/status attributes surfaced by the data
			// source (these are intentionally NOT modeled on the resource). All Computed.
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "The flags for this entry.",
			},
			"hostrtgwact": schema.StringAttribute{
				Computed:    true,
				Description: "Actual Gateway used for advertising host route.",
			},
			"ospfareaval": schema.Int64Attribute{
				Computed:    true,
				Description: "The area ID of the area in which OSPF Type1 LSAs are advertised.",
			},
			"viprtadv2bsd": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether this route is advertised to FreeBSD.",
			},
			"vipvsercount": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of vservers bound to this VIP.",
			},
			"vipvserdowncount": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of vservers bound to this VIP, which are down.",
			},
			"vipvsrvrrhiactivecount": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of vservers that have RHI state ACTIVE.",
			},
			"vipvsrvrrhiactiveupcount": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of vservers that have RHI state ACTIVE, which are UP.",
			},
			"freeports": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of free Ports available on this IP.",
			},
			"iptype": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Type(s) of this IP. Possible values: [ SNIP, VIP, NSIP, HostIP, GSLBsiteIP, CLIP ]",
			},
			"operationalarpowner": schema.Int64Attribute{
				Computed:    true,
				Description: "Run time operational Arp Owner.",
			},
		},
	}
}

// nsipDataSourceSetAttrFromGet projects a NITRO nsip GET response onto the
// data-source model. There is no prior plan/state to preserve, so the shared
// utils.MapGet* helpers fill each attribute (or leave it Null when the GET omits
// it). The ID is the plain ipaddress value, matching the resource ID scheme.
func nsipDataSourceSetAttrFromGet(ctx context.Context, data *NsipDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nsipDataSourceSetAttrFromGet Function")

	if v, ok := g["ipaddress"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Ipaddress = types.StringValue(utils.AnyToString(v))
	}

	data.Advertiseondefaultpartition = utils.MapGetString(g, "advertiseondefaultpartition")
	data.Arp = utils.MapGetString(g, "arp")
	data.Arpowner = utils.MapGetInt64(g, "arpowner")
	data.Arpresponse = utils.MapGetString(g, "arpresponse")
	data.Bgp = utils.MapGetString(g, "bgp")
	data.Decrementttl = utils.MapGetString(g, "decrementttl")
	data.Dynamicrouting = utils.MapGetString(g, "dynamicrouting")
	data.Ftp = utils.MapGetString(g, "ftp")
	data.Gui = utils.MapGetString(g, "gui")
	data.Hostroute = utils.MapGetString(g, "hostroute")
	data.Hostrtgw = utils.MapGetString(g, "hostrtgw")
	data.Icmp = utils.MapGetString(g, "icmp")
	data.Icmpresponse = utils.MapGetString(g, "icmpresponse")
	data.Metric = utils.MapGetInt64(g, "metric")
	data.Mgmtaccess = utils.MapGetString(g, "mgmtaccess")
	data.Mptcpadvertise = utils.MapGetString(g, "mptcpadvertise")
	data.Netmask = utils.MapGetString(g, "netmask")
	data.Networkroute = utils.MapGetString(g, "networkroute")
	data.Ospf = utils.MapGetString(g, "ospf")
	data.Ospfarea = utils.MapGetInt64(g, "ospfarea")
	data.Ospflsatype = utils.MapGetString(g, "ospflsatype")
	data.Ownerdownresponse = utils.MapGetString(g, "ownerdownresponse")
	data.Ownernode = utils.MapGetInt64(g, "ownernode")
	data.Restrictaccess = utils.MapGetString(g, "restrictaccess")
	data.Rip = utils.MapGetString(g, "rip")
	data.Snmp = utils.MapGetString(g, "snmp")
	data.Ssh = utils.MapGetString(g, "ssh")
	data.State = utils.MapGetString(g, "state")
	data.Tag = utils.MapGetInt64(g, "tag")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}
	data.Telnet = utils.MapGetString(g, "telnet")
	data.Type = utils.MapGetString(g, "type")
	data.Vrid = utils.MapGetInt64(g, "vrid")
	data.Vserver = utils.MapGetString(g, "vserver")
	data.Vserverrhilevel = utils.MapGetString(g, "vserverrhilevel")

	// Read-only runtime/status attributes.
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Hostrtgwact = utils.MapGetString(g, "hostrtgwact")
	data.Ospfareaval = utils.MapGetInt64(g, "ospfareaval")
	data.Viprtadv2bsd = utils.MapGetBool(g, "viprtadv2bsd")
	data.Vipvsercount = utils.MapGetInt64(g, "vipvsercount")
	data.Vipvserdowncount = utils.MapGetInt64(g, "vipvserdowncount")
	data.Vipvsrvrrhiactivecount = utils.MapGetInt64(g, "vipvsrvrrhiactivecount")
	data.Vipvsrvrrhiactiveupcount = utils.MapGetInt64(g, "vipvsrvrrhiactiveupcount")
	data.Freeports = utils.MapGetInt64(g, "freeports")
	data.Iptype = utils.MapGetStringList(g, "iptype")
	data.Operationalarpowner = utils.MapGetInt64(g, "operationalarpowner")
}
