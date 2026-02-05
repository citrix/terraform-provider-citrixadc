package nsip6

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func Nsip6DataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"advertiseondefaultpartition": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise VIPs from Shared VLAN on Default Partition",
			},
			"decrementhoplimit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Decrement Hop Limit by 1 when ENABLED.This setting is applicable only for UDP traffic.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow dynamic routing on this IP address. Specific to Subnet IPv6 (SNIP6) address.",
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
				Description: "Option to push the VIP6 to ZebOS routing table for Kernel route redistribution through dynamic routing protocols.",
			},
			"icmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to ICMP requests for this IP address.",
			},
			"icmpresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to ICMPv6 requests for a Virtual IP (VIP) address on the basis of the states of the virtual servers associated with that VIP",
			},
			"ip6hostrtgw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 address of the gateway for the route. If Gateway is not set, VIP uses :: as the gateway.",
			},
			"ipv6address": schema.StringAttribute{
				Required:    true,
				Description: "IPv6 address to create on the Citrix ADC.",
			},
			"map": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mapped IPV4 address for the IPV6 address.",
			},
			"metric": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value to add to or subtract from the cost of the route advertised for the VIP6 address.",
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
			"nd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Respond to Neighbor Discovery (ND) requests for this IP address.",
			},
			"ndowner": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NdOwner in Cluster for VIPS and Striped SNIPS",
			},
			"networkroute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option to push the SNIP6 subnet to ZebOS routing table for Kernel route redistribution through dynamic routing protocol.",
			},
			"ospf6lsatype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of LSAs to be used by the IPv6 OSPF protocol, running on the Citrix ADC, for advertising the route for the VIP6 address.",
			},
			"ospfarea": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the area in which the Intra-Area-Prefix LSAs are to be advertised for the VIP6 address by the IPv6 OSPF protocol running on the Citrix ADC. When ospfArea is not set, VIP6 is advertised on all areas.",
			},
			"ownerdownresponse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "in cluster system, if the owner node is down, whether should it respond to icmp/arp",
			},
			"ownernode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the cluster node for which you are adding the IP address. Must be used if you want the IP address to be active only on the specific node. Can be configured only through the cluster IP address. Cannot be changed after the IP address is created.",
			},
			"restrictaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Block access to nonmanagement applications on this IP address. This option is applicable forMIP6s, SNIP6s, and NSIP6s, and is disabled by default. Nonmanagement applications can run on the underlying Citrix ADC Free BSD operating system.",
			},
			"scope": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Scope of the IPv6 address to be created. Cannot be changed after the IP address is created.",
			},
			"snmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Simple Network Management Protocol (SNMP) access to this IP address.",
			},
			"ssh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow secure Shell (SSH) access to this IP address.",
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
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"telnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow Telnet access to this IP address.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of IP address to be created on the Citrix ADC. Cannot be changed after the IP address is created.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The VLAN number.",
			},
			"vrid6": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "A positive integer that uniquely identifies a VMAC address for binding to this VIP address. This binding is used to set up Citrix ADCs in an active-active configuration using VRRP.",
			},
			"vserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the state of all the virtual servers associated with this VIP6 address.",
			},
			"vserverrhilevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Advertise or do not advertise the route for the Virtual IP (VIP6) address on the basis of the state of the virtual servers associated with that VIP6.\n* NONE - Advertise the route for the VIP6 address, irrespective of the state of the virtual servers associated with the address.\n* ONE VSERVER - Advertise the route for the VIP6 address if at least one of the associated virtual servers is in UP state.\n* ALL VSERVER - Advertise the route for the VIP6 address if all of the associated virtual servers are in UP state.\n* VSVR_CNTRLD.   Advertise the route for the VIP address according to the  RHIstate (RHI STATE) parameter setting on all the associated virtual servers of the VIP address along with their states.\n\nWhen Vserver RHI Level (RHI) parameter is set to VSVR_CNTRLD, the following are different RHI behaviors for the VIP address on the basis of RHIstate (RHI STATE) settings on the virtual servers associated with the VIP address:\n * If you set RHI STATE to PASSIVE on all virtual servers, the Citrix ADC always advertises the route for the VIP address.\n * If you set RHI STATE to ACTIVE on all virtual servers, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers is in UP state.\n *If you set RHI STATE to ACTIVE on some and PASSIVE on others, the Citrix ADC advertises the route for the VIP address if at least one of the associated virtual servers, whose RHI STATE set to ACTIVE, is in UP state.",
			},
		},
	}
}
