package Interface

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// InterfaceDataSourceModel is the data-source-specific model, decoupled from
// InterfaceResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only interface metadata/statistics that the resource
// deliberately omits. Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type InterfaceDataSourceModel struct {
	Id types.String `tfsdk:"id"`

	// Read/write attributes surfaced here as Computed outputs.
	Autoneg          types.String `tfsdk:"autoneg"`
	Bandwidthhigh    types.Int64  `tfsdk:"bandwidthhigh"`
	Bandwidthnormal  types.Int64  `tfsdk:"bandwidthnormal"`
	Duplex           types.String `tfsdk:"duplex"`
	Flowctl          types.String `tfsdk:"flowctl"`
	Haheartbeat      types.String `tfsdk:"haheartbeat"`
	Hamonitor        types.String `tfsdk:"hamonitor"`
	Interfaceid      types.String `tfsdk:"interface_id"` // Required lookup key
	Ifalias          types.String `tfsdk:"ifalias"`
	Lacpkey          types.Int64  `tfsdk:"lacpkey"`
	Lacpmode         types.String `tfsdk:"lacpmode"`
	Lacppriority     types.Int64  `tfsdk:"lacppriority"`
	Lacptimeout      types.String `tfsdk:"lacptimeout"`
	Lagtype          types.String `tfsdk:"lagtype"`
	Linkredundancy   types.String `tfsdk:"linkredundancy"`
	Lldpmode         types.String `tfsdk:"lldpmode"`
	Lrsetpriority    types.Int64  `tfsdk:"lrsetpriority"`
	Mtu              types.Int64  `tfsdk:"mtu"`
	Ringsize         types.Int64  `tfsdk:"ringsize"`
	Ringtype         types.String `tfsdk:"ringtype"`
	Speed            types.String `tfsdk:"speed"`
	State            types.String `tfsdk:"state"`
	Tagall           types.String `tfsdk:"tagall"`
	Throughput       types.Int64  `tfsdk:"throughput"`
	Trunk            types.String `tfsdk:"trunk"`
	Trunkallowedvlan types.List   `tfsdk:"trunkallowedvlan"`
	Trunkmode        types.String `tfsdk:"trunkmode"`

	// Read-only (GET-only) interface metadata/statistics from the NITRO doc
	// read-only set (zion73x_readonly/interface.json). Never settable; populated
	// from GET. Null when the appliance omits them.
	Devicename                types.String `tfsdk:"devicename"`
	Unit                      types.Int64  `tfsdk:"unit"`
	Description               types.String `tfsdk:"description"`
	Flags                     types.Int64  `tfsdk:"flags"`
	Actualmtu                 types.Int64  `tfsdk:"actualmtu"`
	Vlan                      types.Int64  `tfsdk:"vlan"`
	Mac                       types.String `tfsdk:"mac"`
	Uptime                    types.Int64  `tfsdk:"uptime"`
	Downtime                  types.Int64  `tfsdk:"downtime"`
	Actualringsize            types.Int64  `tfsdk:"actualringsize"`
	Reqmedia                  types.String `tfsdk:"reqmedia"`
	Reqspeed                  types.String `tfsdk:"reqspeed"`
	Reqduplex                 types.String `tfsdk:"reqduplex"`
	Reqflowcontrol            types.String `tfsdk:"reqflowcontrol"`
	Actmedia                  types.String `tfsdk:"actmedia"`
	Actspeed                  types.String `tfsdk:"actspeed"`
	Actduplex                 types.String `tfsdk:"actduplex"`
	Actflowctl                types.String `tfsdk:"actflowctl"`
	Mode                      types.String `tfsdk:"mode"`
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
	Bdgmacmoved               types.Int64  `tfsdk:"bdgmacmoved"`
	Bdgmuted                  types.Int64  `tfsdk:"bdgmuted"`
	Vmac                      types.String `tfsdk:"vmac"`
	Vmac6                     types.String `tfsdk:"vmac6"`
	Reqthroughput             types.Int64  `tfsdk:"reqthroughput"`
	Actthroughput             types.Int64  `tfsdk:"actthroughput"`
	Backplane                 types.String `tfsdk:"backplane"`
	Ifnum                     types.List   `tfsdk:"ifnum"`
	Cleartime                 types.Int64  `tfsdk:"cleartime"`
	Slavestate                types.Int64  `tfsdk:"slavestate"`
	Slavemedia                types.Int64  `tfsdk:"slavemedia"`
	Slavespeed                types.Int64  `tfsdk:"slavespeed"`
	Slaveduplex               types.Int64  `tfsdk:"slaveduplex"`
	Slaveflowctl              types.Int64  `tfsdk:"slaveflowctl"`
	Slavetime                 types.Int64  `tfsdk:"slavetime"`
	Intftype                  types.String `tfsdk:"intftype"`
	Svmcmd                    types.Int64  `tfsdk:"svmcmd"`
	Lacpactormode             types.String `tfsdk:"lacpactormode"`
	Lacpactortimeout          types.String `tfsdk:"lacpactortimeout"`
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
	Lractiveintf              types.Int64  `tfsdk:"lractiveintf"`
}

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
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Link state of the interface (ENABLED/DISABLED). Configuring this attribute enables or disables the interface via the NITRO enable/disable actions.",
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

			// Read-only (GET-only) interface metadata/statistics surfaced by the
			// data source (these are intentionally NOT modeled on the resource).
			// All Computed; null when the appliance omits them.
			"devicename": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the interface.",
			},
			"unit": schema.Int64Attribute{
				Computed:    true,
				Description: "Unit number for this interface, signifying the sequence number in which this interface is discovered on this Citrix ADC.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Display the type of interface, the speeds at which this interface can operate, and, if applicable, the type of SFP.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags for this interface. Used for communicating the device states.",
			},
			"actualmtu": schema.Int64Attribute{
				Computed:    true,
				Description: "MTU for this interface (the largest frame that can transit this interface).",
			},
			"vlan": schema.Int64Attribute{
				Computed:    true,
				Description: "Native VLAN for this interface.",
			},
			"mac": schema.StringAttribute{
				Computed:    true,
				Description: "MAC address for this interface.",
			},
			"uptime": schema.Int64Attribute{
				Computed:    true,
				Description: "Duration for which the interface has been UP. This value is reset when the interface state changes to DOWN.",
			},
			"downtime": schema.Int64Attribute{
				Computed:    true,
				Description: "Duration for which the interface has been DOWN. This value is reset when the interface state changes to UP.",
			},
			"actualringsize": schema.Int64Attribute{
				Computed:    true,
				Description: "Actual receive ringsize of the interface.",
			},
			"reqmedia": schema.StringAttribute{
				Computed:    true,
				Description: "Requested media setting for this interface.",
			},
			"reqspeed": schema.StringAttribute{
				Computed:    true,
				Description: "Requested speed setting for this interface.",
			},
			"reqduplex": schema.StringAttribute{
				Computed:    true,
				Description: "Requested duplex setting for this interface.",
			},
			"reqflowcontrol": schema.StringAttribute{
				Computed:    true,
				Description: "Requested flow control setting for this interface.",
			},
			"actmedia": schema.StringAttribute{
				Computed:    true,
				Description: "Actual media setting for this interface.",
			},
			"actspeed": schema.StringAttribute{
				Computed:    true,
				Description: "Actual speed setting for this interface.",
			},
			"actduplex": schema.StringAttribute{
				Computed:    true,
				Description: "Actual duplex setting for this interface.",
			},
			"actflowctl": schema.StringAttribute{
				Computed:    true,
				Description: "Actual flow control setting for this interface.",
			},
			"mode": schema.StringAttribute{
				Computed:    true,
				Description: "Interface Link Aggregation mode (Auto/Manual) setting.",
			},
			"autonegresult": schema.Int64Attribute{
				Computed:    true,
				Description: "Actual auto-negotiation setting for this interface.",
			},
			"tagged": schema.Int64Attribute{
				Computed:    true,
				Description: "VLAN tags setting on this channel.",
			},
			"taggedany": schema.Int64Attribute{
				Computed:    true,
				Description: "Interface setting to accept/drop all tagged packets.",
			},
			"taggedautolearn": schema.Int64Attribute{
				Computed:    true,
				Description: "Dynamic VLAN membership autolearning enabled or disabled on this interface.",
			},
			"hangdetect": schema.Int64Attribute{
				Computed:    true,
				Description: "Hang detection enabled or disabled for this interface.",
			},
			"hangreset": schema.Int64Attribute{
				Computed:    true,
				Description: "Hang reset enabled or disabled for this interface.",
			},
			"linkstate": schema.Int64Attribute{
				Computed:    true,
				Description: "The current state of the link associated with the interface.",
			},
			"intfstate": schema.Int64Attribute{
				Computed:    true,
				Description: "Current state of the specified interface. The interface state is set to UP only if the link state is UP and administrative state is ENABLED.",
			},
			"rxpackets": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of packets received by an interface since the Citrix ADC was started or the interface statistics were cleared.",
			},
			"rxbytes": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of bytes received by an interface since the Citrix ADC was started or the interface statistics were cleared.",
			},
			"rxerrors": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of inbound packets dropped by the hardware on a specified interface.",
			},
			"rxdrops": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of inbound packets dropped by the specified interface.",
			},
			"txpackets": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of packets transmitted by an interface since the Citrix ADC was started or the interface statistics were cleared.",
			},
			"txbytes": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of bytes transmitted by an interface since the Citrix ADC was started or the interface statistics were cleared.",
			},
			"txerrors": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of outbound packets dropped by the hardware on a specified interface.",
			},
			"txdrops": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of packets dropped in transmission by the specified interface.",
			},
			"indisc": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of error-free inbound packets discarded by the specified interface because of a lack of resources.",
			},
			"outdisc": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of error-free outbound packets discarded by the specified interface because of a lack of resources.",
			},
			"fctls": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times flow control is performed on the specified interface because of received pause frames.",
			},
			"hangs": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times the specified interface detected hangs in the transmit and receive paths.",
			},
			"stsstalls": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times the status updates for a specified interface were stalled.",
			},
			"txstalls": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times the interface stalled when transmitting packets.",
			},
			"rxstalls": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times the interface stalled when receiving packets.",
			},
			"bdgmacmoved": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of MAC moves between ports.",
			},
			"bdgmuted": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times the specified interface stopped transmitting and receiving packets because of MAC moves between ports.",
			},
			"vmac": schema.StringAttribute{
				Computed:    true,
				Description: "Virtual MAC of this interface.",
			},
			"vmac6": schema.StringAttribute{
				Computed:    true,
				Description: "Virtual MAC for IPv6 of this interface.",
			},
			"reqthroughput": schema.Int64Attribute{
				Computed:    true,
				Description: "Minimum required throughput for an interface.",
			},
			"actthroughput": schema.Int64Attribute{
				Computed:    true,
				Description: "Actual throughput for the interface.",
			},
			"backplane": schema.StringAttribute{
				Computed:    true,
				Description: "The cluster backplane status of the interface.",
			},
			"ifnum": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Contains the LA Master, if the interface is part of LA channel.",
			},
			"cleartime": schema.Int64Attribute{
				Computed:    true,
				Description: "Time since the interface stats are cleared last time.",
			},
			"slavestate": schema.Int64Attribute{
				Computed:    true,
				Description: "State of the member interfaces.",
			},
			"slavemedia": schema.Int64Attribute{
				Computed:    true,
				Description: "Media type of the member interfaces.",
			},
			"slavespeed": schema.Int64Attribute{
				Computed:    true,
				Description: "Speed of the member interfaces.",
			},
			"slaveduplex": schema.Int64Attribute{
				Computed:    true,
				Description: "Duplex of the member interfaces.",
			},
			"slaveflowctl": schema.Int64Attribute{
				Computed:    true,
				Description: "Flowcontrol of the member interfaces.",
			},
			"slavetime": schema.Int64Attribute{
				Computed:    true,
				Description: "UP time of the member interfaces.",
			},
			"intftype": schema.StringAttribute{
				Computed:    true,
				Description: "Interface Type (virtual, physical or loopback).",
			},
			"svmcmd": schema.Int64Attribute{
				Computed:    true,
				Description: "Attribute to identify the source of cmd; set to 1 when SVM fires the nitro cmd.",
			},
			"lacpactormode": schema.StringAttribute{
				Computed:    true,
				Description: "LACP actor mode (DISABLED, ACTIVE, PASSIVE).",
			},
			"lacpactortimeout": schema.StringAttribute{
				Computed:    true,
				Description: "Interval at which the Citrix ADC sends LACPDU messages to the peer device on the LA channel (LONG, SHORT).",
			},
			"lacpactorpriority": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Actor Priority.",
			},
			"lacpactorportno": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Actor port number.",
			},
			"lacppartnerstate": schema.StringAttribute{
				Computed:    true,
				Description: "LACP Partner State. Whether the port is in Active or Passive negotiating state.",
			},
			"lacppartnertimeout": schema.StringAttribute{
				Computed:    true,
				Description: "The timeout value for the information received in LACPDUs (LONG, SHORT).",
			},
			"lacppartneraggregation": schema.StringAttribute{
				Computed:    true,
				Description: "The Aggregation flag of the partner.",
			},
			"lacppartnerinsync": schema.StringAttribute{
				Computed:    true,
				Description: "The Synchronization flag of the partner.",
			},
			"lacppartnercollecting": schema.StringAttribute{
				Computed:    true,
				Description: "The Collecting flag of the partner.",
			},
			"lacppartnerdistributing": schema.StringAttribute{
				Computed:    true,
				Description: "The Distributing flag of the partner.",
			},
			"lacppartnerdefaulted": schema.StringAttribute{
				Computed:    true,
				Description: "Whether the Receive Machine of the partner entered the Defaulted state.",
			},
			"lacppartnerexpired": schema.StringAttribute{
				Computed:    true,
				Description: "Whether the Receive Machine of the partner entered the Expired state.",
			},
			"lacppartnerpriority": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Partner Priority.",
			},
			"lacppartnersystemmac": schema.StringAttribute{
				Computed:    true,
				Description: "LACP Partner System MAC.",
			},
			"lacppartnersystempriority": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Partner System Priority.",
			},
			"lacppartnerportno": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Partner Port number.",
			},
			"lacppartnerkey": schema.Int64Attribute{
				Computed:    true,
				Description: "LACP Partner Key. The LACP key used by the partner port.",
			},
			"lacpactoraggregation": schema.StringAttribute{
				Computed:    true,
				Description: "The Aggregation flag of the actor.",
			},
			"lacpactorinsync": schema.StringAttribute{
				Computed:    true,
				Description: "The Synchronization flag of the actor.",
			},
			"lacpactorcollecting": schema.StringAttribute{
				Computed:    true,
				Description: "The Collecting flag of the actor.",
			},
			"lacpactordistributing": schema.StringAttribute{
				Computed:    true,
				Description: "The Distributing flag of the actor.",
			},
			"lacpportmuxstate": schema.StringAttribute{
				Computed:    true,
				Description: "LACP Port MUX state.",
			},
			"lacpportrxstat": schema.StringAttribute{
				Computed:    true,
				Description: "LACP Port RX state.",
			},
			"lacpportselectstate": schema.StringAttribute{
				Computed:    true,
				Description: "LACP Port SELECT state (SELECTED, UNSELECTED, STANDBY).",
			},
			"lractiveintf": schema.Int64Attribute{
				Computed:    true,
				Description: "LR set member interface state (active/inactive).",
			},
		},
	}
}

// interfaceDataSourceSetAttrFromGet projects a NITRO interface GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func interfaceDataSourceSetAttrFromGet(ctx context.Context, data *InterfaceDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In interfaceDataSourceSetAttrFromGet Function")

	if v, ok := g["id"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Interfaceid = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Autoneg = utils.MapGetString(g, "autoneg")
	data.Bandwidthhigh = utils.MapGetInt64(g, "bandwidthhigh")
	data.Bandwidthnormal = utils.MapGetInt64(g, "bandwidthnormal")
	data.Duplex = utils.MapGetString(g, "duplex")
	data.Flowctl = utils.MapGetString(g, "flowctl")
	data.Haheartbeat = utils.MapGetString(g, "haheartbeat")
	data.Hamonitor = utils.MapGetString(g, "hamonitor")
	data.Ifalias = utils.MapGetString(g, "ifalias")
	data.Lacpkey = utils.MapGetInt64(g, "lacpkey")
	data.Lacpmode = utils.MapGetString(g, "lacpmode")
	data.Lacppriority = utils.MapGetInt64(g, "lacppriority")
	data.Lacptimeout = utils.MapGetString(g, "lacptimeout")
	data.Lagtype = utils.MapGetString(g, "lagtype")
	data.Linkredundancy = utils.MapGetString(g, "linkredundancy")
	data.Lldpmode = utils.MapGetString(g, "lldpmode")
	data.Lrsetpriority = utils.MapGetInt64(g, "lrsetpriority")
	data.Mtu = utils.MapGetInt64(g, "mtu")
	data.Ringsize = utils.MapGetInt64(g, "ringsize")
	data.Ringtype = utils.MapGetString(g, "ringtype")
	data.Speed = utils.MapGetString(g, "speed")
	data.State = utils.MapGetString(g, "state")
	data.Tagall = utils.MapGetString(g, "tagall")
	data.Throughput = utils.MapGetInt64(g, "throughput")
	data.Trunk = utils.MapGetString(g, "trunk")
	data.Trunkallowedvlan = utils.MapGetStringList(g, "trunkallowedvlan")
	data.Trunkmode = utils.MapGetString(g, "trunkmode")

	// Read-only interface metadata/statistics.
	data.Devicename = utils.MapGetString(g, "devicename")
	data.Unit = utils.MapGetInt64(g, "unit")
	data.Description = utils.MapGetString(g, "description")
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Actualmtu = utils.MapGetInt64(g, "actualmtu")
	data.Vlan = utils.MapGetInt64(g, "vlan")
	data.Mac = utils.MapGetString(g, "mac")
	data.Uptime = utils.MapGetInt64(g, "uptime")
	data.Downtime = utils.MapGetInt64(g, "downtime")
	data.Actualringsize = utils.MapGetInt64(g, "actualringsize")
	data.Reqmedia = utils.MapGetString(g, "reqmedia")
	data.Reqspeed = utils.MapGetString(g, "reqspeed")
	data.Reqduplex = utils.MapGetString(g, "reqduplex")
	data.Reqflowcontrol = utils.MapGetString(g, "reqflowcontrol")
	data.Actmedia = utils.MapGetString(g, "actmedia")
	data.Actspeed = utils.MapGetString(g, "actspeed")
	data.Actduplex = utils.MapGetString(g, "actduplex")
	data.Actflowctl = utils.MapGetString(g, "actflowctl")
	data.Mode = utils.MapGetString(g, "mode")
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
	data.Bdgmacmoved = utils.MapGetInt64(g, "bdgmacmoved")
	data.Bdgmuted = utils.MapGetInt64(g, "bdgmuted")
	data.Vmac = utils.MapGetString(g, "vmac")
	data.Vmac6 = utils.MapGetString(g, "vmac6")
	data.Reqthroughput = utils.MapGetInt64(g, "reqthroughput")
	data.Actthroughput = utils.MapGetInt64(g, "actthroughput")
	data.Backplane = utils.MapGetString(g, "backplane")
	data.Ifnum = utils.MapGetStringList(g, "ifnum")
	data.Cleartime = utils.MapGetInt64(g, "cleartime")
	data.Slavestate = utils.MapGetInt64(g, "slavestate")
	data.Slavemedia = utils.MapGetInt64(g, "slavemedia")
	data.Slavespeed = utils.MapGetInt64(g, "slavespeed")
	data.Slaveduplex = utils.MapGetInt64(g, "slaveduplex")
	data.Slaveflowctl = utils.MapGetInt64(g, "slaveflowctl")
	data.Slavetime = utils.MapGetInt64(g, "slavetime")
	data.Intftype = utils.MapGetString(g, "intftype")
	data.Svmcmd = utils.MapGetInt64(g, "svmcmd")
	data.Lacpactormode = utils.MapGetString(g, "lacpactormode")
	data.Lacpactortimeout = utils.MapGetString(g, "lacpactortimeout")
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
	data.Lractiveintf = utils.MapGetInt64(g, "lractiveintf")
}
