package channel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ChannelDataSourceModel describes the datasource data model.
// It mirrors ChannelResourceModel but exposes the lookup key as "channelid"
// (no underscore) to preserve the datasource's original, backward-compatible
// attribute name. The resource keeps its own "channel_id" attribute.
type ChannelDataSourceModel struct {
	Id              types.String `tfsdk:"id"`
	Bandwidthhigh   types.Int64  `tfsdk:"bandwidthhigh"`
	Bandwidthnormal types.Int64  `tfsdk:"bandwidthnormal"`
	Conndistr       types.String `tfsdk:"conndistr"`
	Flowctl         types.String `tfsdk:"flowctl"`
	Haheartbeat     types.String `tfsdk:"haheartbeat"`
	Hamonitor       types.String `tfsdk:"hamonitor"`
	Channelid       types.String `tfsdk:"channelid"`
	Ifalias         types.String `tfsdk:"ifalias"`
	Ifnum           types.List   `tfsdk:"ifnum"`
	Lamac           types.String `tfsdk:"lamac"`
	Linkredundancy  types.String `tfsdk:"linkredundancy"`
	Lrminthroughput types.Int64  `tfsdk:"lrminthroughput"`
	Macdistr        types.String `tfsdk:"macdistr"`
	Mode            types.String `tfsdk:"mode"`
	Mtu             types.Int64  `tfsdk:"mtu"`
	Speed           types.String `tfsdk:"speed"`
	State           types.String `tfsdk:"state"`
	Tagall          types.String `tfsdk:"tagall"`
	Throughput      types.Int64  `tfsdk:"throughput"`
	Trunk           types.String `tfsdk:"trunk"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/channel.json). Never settable; populated from GET.
	Devicename                types.String `tfsdk:"devicename"`
	Unit                      types.Int64  `tfsdk:"unit"`
	Description               types.String `tfsdk:"description"`
	Flags                     types.Int64  `tfsdk:"flags"`
	Actualmtu                 types.Int64  `tfsdk:"actualmtu"`
	Vlan                      types.Int64  `tfsdk:"vlan"`
	Mac                       types.String `tfsdk:"mac"`
	Uptime                    types.Int64  `tfsdk:"uptime"`
	Downtime                  types.Int64  `tfsdk:"downtime"`
	Reqmedia                  types.String `tfsdk:"reqmedia"`
	Reqspeed                  types.String `tfsdk:"reqspeed"`
	Reqduplex                 types.String `tfsdk:"reqduplex"`
	Reqflowcontrol            types.String `tfsdk:"reqflowcontrol"`
	Media                     types.String `tfsdk:"media"`
	Actspeed                  types.String `tfsdk:"actspeed"`
	Duplex                    types.String `tfsdk:"duplex"`
	Actflowctl                types.String `tfsdk:"actflowctl"`
	Lamode                    types.String `tfsdk:"lamode"`
	Autoneg                   types.Int64  `tfsdk:"autoneg"`
	Autonegresult             types.Int64  `tfsdk:"autonegresult"`
	Tagged                    types.Int64  `tfsdk:"tagged"`
	Taggedany                 types.Int64  `tfsdk:"taggedany"`
	Taggedautolearn           types.Int64  `tfsdk:"taggedautolearn"`
	Hangdetect                types.Int64  `tfsdk:"hangdetect"`
	Hangreset                 types.Int64  `tfsdk:"hangreset"`
	Linkstate                 types.Int64  `tfsdk:"linkstate"`
	Intfstate                 types.Int64  `tfsdk:"intfstate"`
	Rxpackets                 types.Int64  `tfsdk:"rxpackets"`
	Rxbytes                   types.Int64  `tfsdk:"rxbytes"`
	Rxerrors                  types.Int64  `tfsdk:"rxerrors"`
	Rxdrops                   types.Int64  `tfsdk:"rxdrops"`
	Txpackets                 types.Int64  `tfsdk:"txpackets"`
	Txbytes                   types.Int64  `tfsdk:"txbytes"`
	Txerrors                  types.Int64  `tfsdk:"txerrors"`
	Txdrops                   types.Int64  `tfsdk:"txdrops"`
	Indisc                    types.Int64  `tfsdk:"indisc"`
	Outdisc                   types.Int64  `tfsdk:"outdisc"`
	Fctls                     types.Int64  `tfsdk:"fctls"`
	Hangs                     types.Int64  `tfsdk:"hangs"`
	Stsstalls                 types.Int64  `tfsdk:"stsstalls"`
	Txstalls                  types.Int64  `tfsdk:"txstalls"`
	Rxstalls                  types.Int64  `tfsdk:"rxstalls"`
	Bdgmuted                  types.Int64  `tfsdk:"bdgmuted"`
	Vmac                      types.String `tfsdk:"vmac"`
	Vmac6                     types.String `tfsdk:"vmac6"`
	Reqthroughput             types.Int64  `tfsdk:"reqthroughput"`
	Actthroughput             types.Int64  `tfsdk:"actthroughput"`
	Backplane                 types.String `tfsdk:"backplane"`
	Cleartime                 types.Int64  `tfsdk:"cleartime"`
	Lacpmode                  types.String `tfsdk:"lacpmode"`
	Lacptimeout               types.String `tfsdk:"lacptimeout"`
	Lacpactorpriority         types.Int64  `tfsdk:"lacpactorpriority"`
	Lacpactorportno           types.Int64  `tfsdk:"lacpactorportno"`
	Lacppartnerstate          types.String `tfsdk:"lacppartnerstate"`
	Lacppartnertimeout        types.String `tfsdk:"lacppartnertimeout"`
	Lacppartneraggregation    types.String `tfsdk:"lacppartneraggregation"`
	Lacppartnerinsync         types.String `tfsdk:"lacppartnerinsync"`
	Lacppartnercollecting     types.String `tfsdk:"lacppartnercollecting"`
	Lacppartnerdistributing   types.String `tfsdk:"lacppartnerdistributing"`
	Lacppartnerdefaulted      types.String `tfsdk:"lacppartnerdefaulted"`
	Lacppartnerexpired        types.String `tfsdk:"lacppartnerexpired"`
	Lacppartnerpriority       types.Int64  `tfsdk:"lacppartnerpriority"`
	Lacppartnersystemmac      types.String `tfsdk:"lacppartnersystemmac"`
	Lacppartnersystempriority types.Int64  `tfsdk:"lacppartnersystempriority"`
	Lacppartnerportno         types.Int64  `tfsdk:"lacppartnerportno"`
	Lacppartnerkey            types.Int64  `tfsdk:"lacppartnerkey"`
	Lacpactoraggregation      types.String `tfsdk:"lacpactoraggregation"`
	Lacpactorinsync           types.String `tfsdk:"lacpactorinsync"`
	Lacpactorcollecting       types.String `tfsdk:"lacpactorcollecting"`
	Lacpactordistributing     types.String `tfsdk:"lacpactordistributing"`
	Lacpportmuxstate          types.String `tfsdk:"lacpportmuxstate"`
	Lacpportrxstat            types.String `tfsdk:"lacpportrxstat"`
	Lacpportselectstate       types.String `tfsdk:"lacpportselectstate"`
	Lldpmode                  types.String `tfsdk:"lldpmode"`
}

func ChannelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bandwidthhigh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "High threshold value for the bandwidth usage of the LA channel, in Mbps. The Citrix ADC generates an SNMP trap message when the bandwidth usage of the LA channel is greater than or equal to the specified high threshold value.",
			},
			"bandwidthnormal": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Normal threshold value for the bandwidth usage of the LA channel, in Mbps. When the bandwidth usage of the LA channel returns to less than or equal to the specified normal threshold after exceeding the high threshold, the Citrix ADC generates an SNMP trap message to indicate that the bandwidth usage has returned to normal.",
			},
			"conndistr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The 'connection' distribution mode for the LA channel.",
			},
			"flowctl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the flow control type for this LA channel to manage the flow of frames. Flow control is a function as mentioned in clause 31 of the IEEE 802.3 standard. Flow control allows congested ports to pause traffic from the peer device. Flow control is achieved by sending PAUSE frames.",
			},
			"haheartbeat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In a High Availability (HA) configuration, configure the LA channel for sending heartbeats. LA channel that has HA Heartbeat disabled should not send the heartbeats.",
			},
			"hamonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "In a High Availability (HA) configuration, monitor the LA channel for failure events. Failure of any LA channel that has HA MON enabled triggers HA failover.",
			},
			"channelid": schema.StringAttribute{
				Required:    true,
				Description: "ID for the LA channel or cluster LA channel or LR channel to be created. Specify an LA channel in LA/x notation, where x can range from 1 to 8 or cluster LA channel in CLA/x notation or Link redundant channel in LR/x notation, where x can range from 1 to 4. Cannot be changed after the LA channel is created.",
			},
			"ifalias": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Alias name for the LA channel. Used only to enhance readability. To perform any operations, you have to specify the LA channel ID.",
			},
			"ifnum": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Interfaces to be bound to the LA channel of a Citrix ADC or to the LA channel of a cluster configuration.\nFor an LA channel of a Citrix ADC, specify an interface in C/U notation (for example, 1/3).\nFor an LA channel of a cluster configuration, specify an interface in N/C/U notation (for example, 2/1/3).\nwhere C can take one of the following values:\n* 0 - Indicates a management interface.\n* 1 - Indicates a 1 Gbps port.\n* 10 - Indicates a 10 Gbps port.\nU is a unique integer for representing an interface in a particular port group.\nN is the ID of the node to which an interface belongs in a cluster configuration.\nUse spaces to separate multiple entries.",
			},
			"lamac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies a MAC address for the LA channels configured in Citrix ADC virtual appliances (VPX). This MAC address is persistent after each reboot.\nIf you don't specify this parameter, a MAC address is generated randomly for each LA channel. These MAC addresses change after each reboot.",
			},
			"linkredundancy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Link Redundancy for Cluster LAG.",
			},
			"lrminthroughput": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the minimum throughput threshold (in Mbps) to be met by the active subchannel. Setting this parameter automatically divides an LACP channel into logical subchannels, with one subchannel active and the others in standby mode.  When the maximum supported throughput of the active channel falls below the lrMinThroughput value, link failover occurs and a standby subchannel becomes active.",
			},
			"macdistr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The  'MAC' distribution mode for the LA channel.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The initital mode for the LA channel.",
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The Maximum Transmission Unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes CRC, that can be transmitted and received by an interface. The default value of MTU is 1500 on all the interface of Citrix ADC, some Cloud Platforms will restrict Citrix ADC to use the lesser default value. Any MTU value more than 1500 is called Jumbo MTU and will make the interface as jumbo enabled. The Maximum Jumbo MTU in Citrix ADC is 9216, however, some Virtualized / Cloud Platforms will have lesser Maximum Jumbo MTU Value (9000). In the case of Cluster, the Backplane interface requires an MTU value of 78 bytes more than the Max MTU configured on any other Data-Plane Interface. When the Data plane interfaces are all at default 1500 MTU, Cluster Back Plane will be automatically set to 1578 (1500 + 78) MTU. If a Backplane interface is reset to Data Plane Interface, then the 1578 MTU will be automatically reset to the default MTU of 1500(or whatever lesser default value). If any data plane interface of a Cluster is configured with a Jumbo MTU ( > 1500), then all backplane interfaces require to be configured with a minimum MTU of 'Highest Data Plane MTU in the Cluster + 78'. That makes the maximum Jumbo MTU for any Data-Plane Interface in a Cluster System to be '9138 (9216 - 78)., where 9216 is the maximum Jumbo MTU. On certain Virtualized / Cloud Platforms, the maximum  possible MTU is restricted to a lesser value, Similar calculation can be applied, Maximum Data Plane MTU in Cluster = (Maximum possible MTU - 78).",
			},
			"speed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ethernet speed of the channel, in Mbps. If the speed of any bound interface is greater than or equal to the value set for this parameter, the state of the interface is UP. Otherwise, the state is INACTIVE. Bound Interfaces whose state is INACTIVE do not process any traffic.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable the LA channel.",
			},
			"tagall": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Adds a four-byte 802.1q tag to every packet sent on this channel.  The ON setting applies tags for all VLANs that are bound to this channel. OFF applies the tag for all VLANs other than the native VLAN.",
			},
			"throughput": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Low threshold value for the throughput of the LA channel, in Mbps. In an high availability (HA) configuration, failover is triggered when the LA channel has HA MON enabled and the throughput is below the specified threshold.",
			},
			"trunk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is deprecated by tagall",
			},

			"devicename": schema.StringAttribute{
				Computed:    true,
				Description: "LA channel name in form LA/x, where x is channel ID, which ranges from 1 to 8 or LR channel name in form LR/x, where x is channel ID, which ranges from 1 to 4.",
			},
			"unit": schema.Int64Attribute{
				Computed:    true,
				Description: "Unit number of the channel. This is an internal reference number that the Citrix ADC uses to identify the channel.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "The IEEE standard that the channel is based on.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags of this channel.",
			},
			"actualmtu": schema.Int64Attribute{
				Computed:    true,
				Description: "MTU of the channel. This is the maximum frame size that the channel can process.",
			},
			"vlan": schema.Int64Attribute{
				Computed:    true,
				Description: "Native VLAN of the channel.",
			},
			"mac": schema.StringAttribute{
				Computed:    true,
				Description: "MAC address of the channel.",
			},
			"uptime": schema.Int64Attribute{
				Computed:    true,
				Description: "Duration for which the channel is UP. (Example: 3 hours 1 minute 1 second). This value is reset when the channel state changes to DOWN.",
			},
			"downtime": schema.Int64Attribute{
				Computed:    true,
				Description: "Duration for which the channel is DOWN. (Example: 3 hours 1 minute 1 second). This value is reset when the channel state changes to UP.",
			},
			"reqmedia": schema.StringAttribute{
				Computed:    true,
				Description: "Requested media setting for this channel. Since there is no media associated with LA, the displayed values carry no significance.",
			},
			"reqspeed": schema.StringAttribute{
				Computed:    true,
				Description: "Requested speed setting for this channel. Since no media are associated with LA, this speed is used to determine the threshold for the slave interfaces. If the speed of the member interface is less than the requested speed, that interface is considered inactive.",
			},
			"reqduplex": schema.StringAttribute{
				Computed:    true,
				Description: "Requested duplex setting for this channel. Since no media are associated with LA, the displayed values carry no significance.",
			},
			"reqflowcontrol": schema.StringAttribute{
				Computed:    true,
				Description: "Requested flow control setting for this channel. Since no media are associated with LA, the displayed values carry no significance.",
			},
			"media": schema.StringAttribute{
				Computed:    true,
				Description: "Requested media setting for this interface.",
			},
			"actspeed": schema.StringAttribute{
				Computed:    true,
				Description: "Actual speed setting for this channel.",
			},
			"duplex": schema.StringAttribute{
				Computed:    true,
				Description: "Actualduplex setting for this interface.",
			},
			"actflowctl": schema.StringAttribute{
				Computed:    true,
				Description: "Actual flow control setting for this channel.",
			},
			"lamode": schema.StringAttribute{
				Computed:    true,
				Description: "The  mode(AUTO/MANNUAL) for the LA channel.",
			},
			"autoneg": schema.Int64Attribute{
				Computed:    true,
				Description: "Requested auto negotiation setting for this channel. Since no media are associated with LA, this setting has no effect.",
			},
			"autonegresult": schema.Int64Attribute{
				Computed:    true,
				Description: "Actual  auto negotiation setting for this channel.",
			},
			"tagged": schema.Int64Attribute{
				Computed:    true,
				Description: "VLAN tags setting on this channel.",
			},
			"taggedany": schema.Int64Attribute{
				Computed:    true,
				Description: "Channel setting to accept/drop all tagged packets.",
			},
			"taggedautolearn": schema.Int64Attribute{
				Computed:    true,
				Description: "Dynaminc vlan membership on this channel.",
			},
			"hangdetect": schema.Int64Attribute{
				Computed:    true,
				Description: "Hang detect for this channel.",
			},
			"hangreset": schema.Int64Attribute{
				Computed:    true,
				Description: "Hang reset for this channel.",
			},
			"linkstate": schema.Int64Attribute{
				Computed:    true,
				Description: "The current state of the link associated with the interface. For logical interfaces (LA), the state of the link is dependent on the state of the slave interfaces. For the link to be UP at least one of the slave interfaces needs to be UP.",
			},
			"intfstate": schema.Int64Attribute{
				Computed:    true,
				Description: "Current state of the specified interface.  The interface state set to UP only if the link state is UP and administrative state is ENABLED.",
			},
			"rxpackets": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of bytes received by all the slave interfaces of the channel since the Citrix ADC was started or the interface statistics were cleared.",
			},
			"rxbytes": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of packets received by all member interfaces since the Citrix ADC was started or the interface statistics were cleared.",
			},
			"rxerrors": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of inbound packets dropped by the hardware of the slave interfaces since the Citrix ADC was started or the interface statistics were cleared. Possible causes of dropped packets are CRC, length (undersize or oversize), and alignment errors.",
			},
			"rxdrops": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of inbound packets dropped by the channel's slave interfaces. Commonly dropped packets are multicast frames, spanning tree BPDUs, packets destined to a MAC not owned by the Citrix ADC when L2 mode is disabled, or packets tagged for a VLAN that is not bound to the interface.  In most healthy networks, this statistic increments at a steady rate regardless of traffic load.  A sharp spike in dropped packets generally indicates an issue with connected L2 switches, such as a forwarding database overflow resulting in packets being broadcast on all ports.",
			},
			"txpackets": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of packets transmitted by slave interfaces of a channel since the Citrix ADC was started or the interface statistics were cleared.",
			},
			"txbytes": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of bytes transmitted by slave interfaces of a channel since the Citrix ADC was started or the interface statistics were cleared.",
			},
			"txerrors": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of outbound packets dropped by the hardware of a channel's slave interfaces since the Citrix ADC was started or the interface statistics were cleared. Possible causes of dropped packets are length (undersize or oversize) errors and lack of resources.",
			},
			"txdrops": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of packets dropped in transmission by a channel's slave interfaces for one of the following reasons:",
			},
			"indisc": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of error-free inbound packets discarded by a channel's slave interfaces because of a lack of resources (for example, insufficient receive buffers).",
			},
			"outdisc": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of error-free outbound packets discarded by a channel's slave interfaces because of a lack of resources. This statistic is not available on:",
			},
			"fctls": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times flow control is performed on a channel's slave interfaces because of pause frames.",
			},
			"hangs": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hangs that occurred on the channel's slave interfaces.",
			},
			"stsstalls": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of status stalls that occurred on the channel's slave interfaces.",
			},
			"txstalls": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of Tx stalls happened that occurred on the channel's slave interfaces.",
			},
			"rxstalls": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of Rx stalls that occurred on the channel's slave interfaces.",
			},
			"bdgmuted": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times a channel's slave interfaces stopped transmitting and receiving packets because of MAC moves between ports.",
			},
			"vmac": schema.StringAttribute{
				Computed:    true,
				Description: "Virtual MAC of this channel.",
			},
			"vmac6": schema.StringAttribute{
				Computed:    true,
				Description: "Virtual MAC for IPv6 on this interface.",
			},
			"reqthroughput": schema.Int64Attribute{
				Computed:    true,
				Description: "Minimum required throughput for an interface. Failover is triggered if the operating throughput of a Link Aggregation (LA) channel for which HAMON is ON falls below this value.",
			},
			"actthroughput": schema.Int64Attribute{
				Computed:    true,
				Description: "Actual throughput for the interface.",
			},
			"backplane": schema.StringAttribute{
				Computed:    true,
				Description: "The cluster backplane status of the LA. If the status is enabled, the LA is part of the cluster backplane. By default, the backplane status is disabled.",
			},
			"cleartime": schema.Int64Attribute{
				Computed:    true,
				Description: "Time since the interface stats are cleared last time.",
			},
			"lacpmode": schema.StringAttribute{
				Computed:    true,
				Description: "The LACP mode of the specified interface. The possible values are:",
			},
			"lacptimeout": schema.StringAttribute{
				Computed:    true,
				Description: "Time to wait for the LACPDU.  If a LACPDU is not received within this interval, the Citrix ADC markes the link partner port as DOWN. Possible values: Long and Short. Long lacptimeout is 90 sec and Short LACP timeout is 3 sec.",
			},
			"lacpactorpriority": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Actor Priority. A LACP port priority is configured on each port using LACP. LACP uses the port priority with the port number to form the port identifier. The port priority determines which ports should be put in standby mode when there is a hardware limitation that prevents all compatible ports from aggregating.",
			},
			"lacpactorportno": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Actor port number. LACP uses the port priority with the port number to form the port identifier.",
			},
			"lacppartnerstate": schema.StringAttribute{
				Computed:    true,
				Description: "LACP Partner State. Whether the port is in Active or Passive negotiating state.",
			},
			"lacppartnertimeout": schema.StringAttribute{
				Computed:    true,
				Description: "The timeout value for the information revieved in LACPDUs. It can have values as SHORT or LONG. The SHORT timeout is 3s and the LONG timeout is 90s.",
			},
			"lacppartneraggregation": schema.StringAttribute{
				Computed:    true,
				Description: "The Aggregation flag indicates that the participant will allow the link to be used as part of an aggregate. Otherwise the link is to be used as an individual link, i.e. not aggregated with any other.",
			},
			"lacppartnerinsync": schema.StringAttribute{
				Computed:    true,
				Description: "The Synchronization flag indicates that the transmitting participant.s mux component is in sync with the system id and key information transmitted.",
			},
			"lacppartnercollecting": schema.StringAttribute{
				Computed:    true,
				Description: "The Collecting flag indicates that the participant.s collector, i.e. the reception component of the mux, is definitely on. If set the flag communicates collecting.",
			},
			"lacppartnerdistributing": schema.StringAttribute{
				Computed:    true,
				Description: "The Distributing flag indicates that the participant.s distributor is not definitely off. If reset the flag indicates not distributing.",
			},
			"lacppartnerdefaulted": schema.StringAttribute{
				Computed:    true,
				Description: "If the timer expires in the Expired state, the Receive Machine enters the Defaulted state.",
			},
			"lacppartnerexpired": schema.StringAttribute{
				Computed:    true,
				Description: "If the LACPDUs are received for timeout period, the Receive Machine enters the Expired state and the timer is restarted with the timeout value of SHORT timeout.",
			},
			"lacppartnerpriority": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Partner Priority. A LACP port priority is configured on each port using LACP. LACP uses the port priority with the port number to form the port identifier.",
			},
			"lacppartnersystemmac": schema.StringAttribute{
				Computed:    true,
				Description: "LACP Partner System MAC.",
			},
			"lacppartnersystempriority": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Partner System Priority. The LACP partner's system priority. The values for the priority range from 0 to 65535. The lower the value, the higher the system priority. The switch with the lower system priority value determines which links between LACP partner are active and which are in the standby for each LACP Channel.",
			},
			"lacppartnerportno": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Partner Port number. LACP uses the port priority with the port number to form the port identifier.",
			},
			"lacppartnerkey": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Partner Key. The LACP key used by the partner port.",
			},
			"lacpactoraggregation": schema.StringAttribute{
				Computed:    true,
				Description: "The Aggregation flag indicates that the participant will allow the link to be used as part of an aggregate. Otherwise the link is to be used as an individual link, i.e. not aggregated with any other.",
			},
			"lacpactorinsync": schema.StringAttribute{
				Computed:    true,
				Description: "The Synchronization flag indicates that the transmitting participant.s mux component is in sync with the system id and key information transmitted.",
			},
			"lacpactorcollecting": schema.StringAttribute{
				Computed:    true,
				Description: "The Collecting flag indicates that the participant.s collector, i.e. the reception component of the mux, is definitely on. If set the flag communicates collecting.",
			},
			"lacpactordistributing": schema.StringAttribute{
				Computed:    true,
				Description: "The Distributing flag indicates that the participant.s distributor is not definitely off. If reset the flag indicates not distributing.",
			},
			"lacpportmuxstate": schema.StringAttribute{
				Computed:    true,
				Description: "LACP Port MUX state. The state of the MUX control machine. The  Mux Control Machine attaches the physical port to an aggregate port, using the Selection Logic to choose an appropriate port, and turns the distributor and collector for the physical port on or off as required by protocol	information.",
			},
			"lacpportrxstat": schema.StringAttribute{
				Computed:    true,
				Description: "LACP Port RX state. The state of the Receive machine. The Receive Machine maintains partner information, recording protocol information from LACPDUs sent by remote partner(s). Received information is subject to a timeout, and if sufficient time elapses the receive machine will revert to using default partner information.",
			},
			"lacpportselectstate": schema.StringAttribute{
				Computed:    true,
				Description: "LACP Port SELECT state. The state of the SELECT state machine, It could be SELECTED or UNSELECTED.",
			},
			"lldpmode": schema.StringAttribute{
				Computed:    true,
				Description: "Link Layer Discovery Protocol (LLDP) mode for an interface. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels.",
			},
		},
	}
}

// channelDataSourceSetAttrFromGet projects a NITRO channel GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func channelDataSourceSetAttrFromGet(ctx context.Context, data *ChannelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In channelDataSourceSetAttrFromGet Function")

	// The NITRO channel primary key is returned under "id"; expose it as both the
	// datasource "channelid" lookup key and the synthetic "id".
	if v, ok := g["id"]; ok && v != nil {
		data.Channelid = types.StringValue(utils.AnyToString(v))
	}
	data.Id = types.StringValue(data.Channelid.ValueString())

	// Existing read/write attributes as read-back outputs.
	data.Bandwidthhigh = utils.MapGetInt64(g, "bandwidthhigh")
	data.Bandwidthnormal = utils.MapGetInt64(g, "bandwidthnormal")
	data.Conndistr = utils.MapGetString(g, "conndistr")
	data.Flowctl = utils.MapGetString(g, "flowctl")
	data.Haheartbeat = utils.MapGetString(g, "haheartbeat")
	data.Hamonitor = utils.MapGetString(g, "hamonitor")
	data.Ifalias = utils.MapGetString(g, "ifalias")
	data.Ifnum = utils.MapGetStringList(g, "ifnum")
	data.Lamac = utils.MapGetString(g, "lamac")
	data.Linkredundancy = utils.MapGetString(g, "linkredundancy")
	data.Lrminthroughput = utils.MapGetInt64(g, "lrminthroughput")
	data.Macdistr = utils.MapGetString(g, "macdistr")
	data.Mode = utils.MapGetString(g, "mode")
	data.Mtu = utils.MapGetInt64(g, "mtu")
	data.Speed = utils.MapGetString(g, "speed")
	data.State = utils.MapGetString(g, "state")
	data.Tagall = utils.MapGetString(g, "tagall")
	data.Throughput = utils.MapGetInt64(g, "throughput")
	data.Trunk = utils.MapGetString(g, "trunk")

	// Read-only metadata.
	data.Devicename = utils.MapGetString(g, "devicename")
	data.Unit = utils.MapGetInt64(g, "unit")
	data.Description = utils.MapGetString(g, "description")
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Actualmtu = utils.MapGetInt64(g, "actualmtu")
	data.Vlan = utils.MapGetInt64(g, "vlan")
	data.Mac = utils.MapGetString(g, "mac")
	data.Uptime = utils.MapGetInt64(g, "uptime")
	data.Downtime = utils.MapGetInt64(g, "downtime")
	data.Reqmedia = utils.MapGetString(g, "reqmedia")
	data.Reqspeed = utils.MapGetString(g, "reqspeed")
	data.Reqduplex = utils.MapGetString(g, "reqduplex")
	data.Reqflowcontrol = utils.MapGetString(g, "reqflowcontrol")
	data.Media = utils.MapGetString(g, "media")
	data.Actspeed = utils.MapGetString(g, "actspeed")
	data.Duplex = utils.MapGetString(g, "duplex")
	data.Actflowctl = utils.MapGetString(g, "actflowctl")
	data.Lamode = utils.MapGetString(g, "lamode")
	data.Autoneg = utils.MapGetInt64(g, "autoneg")
	data.Autonegresult = utils.MapGetInt64(g, "autonegresult")
	data.Tagged = utils.MapGetInt64(g, "tagged")
	data.Taggedany = utils.MapGetInt64(g, "taggedany")
	data.Taggedautolearn = utils.MapGetInt64(g, "taggedautolearn")
	data.Hangdetect = utils.MapGetInt64(g, "hangdetect")
	data.Hangreset = utils.MapGetInt64(g, "hangreset")
	data.Linkstate = utils.MapGetInt64(g, "linkstate")
	data.Intfstate = utils.MapGetInt64(g, "intfstate")
	data.Rxpackets = utils.MapGetInt64(g, "rxpackets")
	data.Rxbytes = utils.MapGetInt64(g, "rxbytes")
	data.Rxerrors = utils.MapGetInt64(g, "rxerrors")
	data.Rxdrops = utils.MapGetInt64(g, "rxdrops")
	data.Txpackets = utils.MapGetInt64(g, "txpackets")
	data.Txbytes = utils.MapGetInt64(g, "txbytes")
	data.Txerrors = utils.MapGetInt64(g, "txerrors")
	data.Txdrops = utils.MapGetInt64(g, "txdrops")
	data.Indisc = utils.MapGetInt64(g, "indisc")
	data.Outdisc = utils.MapGetInt64(g, "outdisc")
	data.Fctls = utils.MapGetInt64(g, "fctls")
	data.Hangs = utils.MapGetInt64(g, "hangs")
	data.Stsstalls = utils.MapGetInt64(g, "stsstalls")
	data.Txstalls = utils.MapGetInt64(g, "txstalls")
	data.Rxstalls = utils.MapGetInt64(g, "rxstalls")
	data.Bdgmuted = utils.MapGetInt64(g, "bdgmuted")
	data.Vmac = utils.MapGetString(g, "vmac")
	data.Vmac6 = utils.MapGetString(g, "vmac6")
	data.Reqthroughput = utils.MapGetInt64(g, "reqthroughput")
	data.Actthroughput = utils.MapGetInt64(g, "actthroughput")
	data.Backplane = utils.MapGetString(g, "backplane")
	data.Cleartime = utils.MapGetInt64(g, "cleartime")
	data.Lacpmode = utils.MapGetString(g, "lacpmode")
	data.Lacptimeout = utils.MapGetString(g, "lacptimeout")
	data.Lacpactorpriority = utils.MapGetInt64(g, "lacpactorpriority")
	data.Lacpactorportno = utils.MapGetInt64(g, "lacpactorportno")
	data.Lacppartnerstate = utils.MapGetString(g, "lacppartnerstate")
	data.Lacppartnertimeout = utils.MapGetString(g, "lacppartnertimeout")
	data.Lacppartneraggregation = utils.MapGetString(g, "lacppartneraggregation")
	data.Lacppartnerinsync = utils.MapGetString(g, "lacppartnerinsync")
	data.Lacppartnercollecting = utils.MapGetString(g, "lacppartnercollecting")
	data.Lacppartnerdistributing = utils.MapGetString(g, "lacppartnerdistributing")
	data.Lacppartnerdefaulted = utils.MapGetString(g, "lacppartnerdefaulted")
	data.Lacppartnerexpired = utils.MapGetString(g, "lacppartnerexpired")
	data.Lacppartnerpriority = utils.MapGetInt64(g, "lacppartnerpriority")
	data.Lacppartnersystemmac = utils.MapGetString(g, "lacppartnersystemmac")
	data.Lacppartnersystempriority = utils.MapGetInt64(g, "lacppartnersystempriority")
	data.Lacppartnerportno = utils.MapGetInt64(g, "lacppartnerportno")
	data.Lacppartnerkey = utils.MapGetInt64(g, "lacppartnerkey")
	data.Lacpactoraggregation = utils.MapGetString(g, "lacpactoraggregation")
	data.Lacpactorinsync = utils.MapGetString(g, "lacpactorinsync")
	data.Lacpactorcollecting = utils.MapGetString(g, "lacpactorcollecting")
	data.Lacpactordistributing = utils.MapGetString(g, "lacpactordistributing")
	data.Lacpportmuxstate = utils.MapGetString(g, "lacpportmuxstate")
	data.Lacpportrxstat = utils.MapGetString(g, "lacpportrxstat")
	data.Lacpportselectstate = utils.MapGetString(g, "lacpportselectstate")
	data.Lldpmode = utils.MapGetString(g, "lldpmode")
}
