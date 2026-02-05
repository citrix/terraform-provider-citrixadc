package Interface

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func InterfaceDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"autoneg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Auto-negotiation state of the interface. With the ENABLED setting, the Citrix ADC auto-negotiates the speed and duplex settings with the peer network device on the link. The Citrix ADC appliance auto-negotiates the settings of only those parameters (speed or duplex mode) for which the value is set as AUTO.",
			},
			"bandwidthhigh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "High threshold value for the bandwidth usage of the interface, in Mbps. The Citrix ADC generates an SNMP trap message when the bandwidth usage of the interface is greater than or equal to the specified high threshold value.",
			},
			"bandwidthnormal": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Normal threshold value for the bandwidth usage of the interface, in Mbps. When the bandwidth usage of the interface becomes less than or equal to the specified normal threshold after exceeding the high threshold, the Citrix ADC generates an SNMP trap message to indicate that the bandwidth usage has returned to normal.",
			},
			"duplex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The duplex mode for the interface. Notes:* If you set the duplex mode to AUTO, the Citrix ADC attempts to auto-negotiate the duplex mode of the interface when it is UP. You must enable auto negotiation on the interface. If you set a duplex mode other than AUTO, you must specify the same duplex mode for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors.",
			},
			"flowctl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "802.3x flow control setting for the interface.  The 802.3x specification does not define flow control for 10 Mbps and 100 Mbps speeds, but if a Gigabit Ethernet interface operates at those speeds, the flow control settings can be applied. The flow control setting that is finally applied to an interface depends on auto-negotiation. With the ON option, the peer negotiates the flow control, but the appliance then forces two-way flow control for the interface.",
			},
			"haheartbeat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In a High Availability (HA) or Cluster configuration, configure the interface for sending heartbeats. In an HA or Cluster configuration, an interface that has HA Heartbeat disabled should not send the heartbeats.",
			},
			"hamonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In a High Availability (HA) configuration, monitor the interface for failure events. In an HA configuration, an interface that has HA MON enabled and is not bound to any Failover Interface Set (FIS), is a critical interface. Failure or disabling of any critical interface triggers HA failover.",
			},
			"interface_id": schema.StringAttribute{
				Required:    true,
				Description: "Interface number, in C/U format, where C can take one of the following values:\n* 0 - Indicates a management interface.\n* 1 - Indicates a 1 Gbps port.\n* 10 - Indicates a 10 Gbps port.\n* LA - Indicates a link aggregation port.\n* LO - Indicates a loop back port.\nU is a unique integer for representing an interface in a particular port group.",
			},
			"ifalias": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Alias name for the interface. Used only to enhance readability. To perform any operations, you have to specify the interface ID.",
			},
			"lacpkey": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer identifying the LACP LA channel to which the interface is to be bound.\nFor an LA channel of the Citrix ADC, this digit specifies the variable x of an LA channel in LA/x notation, where x can range from 1 to 8. For example, if you specify 3 as the LACP key for an LA channel, the interface is bound to the LA channel LA/3.\nFor an LA channel of a cluster configuration, this digit specifies the variable y of a cluster LA channel in CLA/(y-4) notation, where y can range from 5 to 8. For example, if you specify 6 as the LACP key for a cluster LA channel, the interface is bound to the cluster LA channel CLA/2.",
			},
			"lacpmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bind the interface to a LA channel created by the Link Aggregation control protocol (LACP).\nAvailable settings function as follows:\n* Active - The LA channel port of the Citrix ADC generates LACPDU messages on a regular basis, regardless of any need expressed by its peer device to receive them.\n* Passive - The LA channel port of the Citrix ADC does not transmit LACPDU messages unless the peer device port is in the active mode. That is, the port does not speak unless spoken to.\n* Disabled - Unbinds the interface from the LA channel. If this is the only interface in the LA channel, the LA channel is removed.",
			},
			"lacppriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "LACP port priority, expressed as an integer. The lower the number, the higher the priority. The Citrix ADC limits the number of interfaces in an LA channel to sixteen.",
			},
			"lacptimeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval at which the Citrix ADC sends LACPDU messages to the peer device on the LA channel.\nAvailable settings function as follows:\nLONG - 30 seconds.\nSHORT - 1 second.",
			},
			"lagtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of entity (Citrix ADC or cluster configuration) for which to create the channel.",
			},
			"linkredundancy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Link Redundancy for Cluster LAG.",
			},
			"lldpmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Link Layer Discovery Protocol (LLDP) mode for an interface. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels.",
			},
			"lrsetpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "LRSET port priority, expressed as an integer ranging from 1 to 1024. The highest priority is 1. The Citrix ADC limits the number of interfaces in an LRSET to 8. Within a LRSET the highest LR Priority Interface is considered as the first candidate for the Active interface, if the interface is UP.",
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The Maximum Transmission Unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes CRC, that can be transmitted and received by an interface. The default value of MTU is 1500 on all the interface of Citrix ADC, some Cloud Platforms will restrict Citrix ADC to use the lesser default value. Any MTU value more than 1500 is called Jumbo MTU and will make the interface as jumbo enabled. The Maximum Jumbo MTU in Citrix ADC is 9216, however, some Virtualized / Cloud Platforms will have lesser Maximum Jumbo MTU Value (9000). In the case of Cluster, the Backplane interface requires an MTU value of 78 bytes more than the Max MTU configured on any other Data-Plane Interface. When the Data plane interfaces are all at default 1500 MTU, Cluster Back Plane will be automatically set to 1578 (1500 + 78) MTU. If a Backplane interface is reset to Data Plane Interface, then the 1578 MTU will be automatically reset to the default MTU of 1500(or whatever lesser default value). If any data plane interface of a Cluster is configured with a Jumbo MTU ( > 1500), then all backplane interfaces require to be configured with a minimum MTU of 'Highest Data Plane MTU in the Cluster + 78'. That makes the maximum Jumbo MTU for any Data-Plane Interface in a Cluster System to be '9138 (9216 - 78)., where 9216 is the maximum Jumbo MTU. On certain Virtualized / Cloud Platforms, the maximum  possible MTU is restricted to a lesser value, Similar calculation can be applied, Maximum Data Plane MTU in Cluster = (Maximum possible MTU - 78).",
			},
			"ringsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The receive ringsize of the interface. A higher number provides more number of buffers in handling incoming traffic.",
			},
			"ringtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The receive ringtype of the interface (Fixed or Elastic). A fixed ring type pre-allocates configured number of buffers irrespective of traffic rate. In contrast, an elastic ring, expands and shrinks based on incoming traffic rate.",
			},
			"speed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ethernet speed of the interface, in Mbps.\nNotes:\n* If you set the speed as AUTO, the Citrix ADC attempts to auto-negotiate or auto-sense the link speed of the interface when it is UP. You must enable auto negotiation on the interface.\n* If you set a speed other than AUTO, you must specify the same speed for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors.\nSome interfaces do not support certain speeds. If you specify an unsupported speed, an error message appears.",
			},
			"tagall": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Add a four-byte 802.1q tag to every packet sent on this interface.  The ON setting applies the tag for this interface's native VLAN. OFF applies the tag for all VLANs other than the native VLAN.",
			},
			"throughput": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Low threshold value for the throughput of the interface, in Mbps. In an HA configuration, failover is triggered if the interface has HA MON enabled and the throughput is below the specified the threshold.",
			},
			"trunk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This argument is deprecated by tagall.",
			},
			"trunkallowedvlan": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "VLAN ID or range of VLAN IDs will be allowed on this trunk interface. In the command line interface, separate the range with a hyphen. For example: 40-90.",
			},
			"trunkmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Accept and send 802.1q VLAN tagged packets, based on Allowed Vlan List of this interface.",
			},
		},
	}
}
