---
subcategory: "Network"
---

# Data Source: interface

Use this data source to retrieve information about a specific network interface on the Citrix ADC.

## Example Usage

### Basic Example

```hcl
data "citrixadc_interface" "example" {
  interface_id = "1/1"
}

output "interface_mtu" {
  value = data.citrixadc_interface.example.mtu
}
```

## Argument Reference

* `interface_id` - (Required) Interface number, in C/U format, where C can take one of the following values:
  * `0` - Indicates a management interface
  * `1` - Indicates a 1 Gbps port
  * `10` - Indicates a 10 Gbps port
  * `LA` - Indicates a link aggregation port
  * `LO` - Indicates a loop back port
  
  U is a unique integer for representing an interface in a particular port group.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the interface resource.

* `autoneg` - Auto-negotiation state of the interface. With the ENABLED setting, the Citrix ADC auto-negotiates the speed and duplex settings with the peer network device on the link. The Citrix ADC appliance auto-negotiates the settings of only those parameters (speed or duplex mode) for which the value is set as AUTO.

* `bandwidthhigh` - High threshold value for the bandwidth usage of the interface, in Mbps. The Citrix ADC generates an SNMP trap message when the bandwidth usage of the interface is greater than or equal to the specified high threshold value.

* `bandwidthnormal` - Normal threshold value for the bandwidth usage of the interface, in Mbps. When the bandwidth usage of the interface becomes less than or equal to the specified normal threshold after exceeding the high threshold, the Citrix ADC generates an SNMP trap message to indicate that the bandwidth usage has returned to normal.

* `duplex` - The duplex mode for the interface. If you set the duplex mode to AUTO, the Citrix ADC attempts to auto-negotiate the duplex mode of the interface when it is UP. You must enable auto negotiation on the interface. If you set a duplex mode other than AUTO, you must specify the same duplex mode for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors.

* `flowctl` - 802.3x flow control setting for the interface. The 802.3x specification does not define flow control for 10 Mbps and 100 Mbps speeds, but if a Gigabit Ethernet interface operates at those speeds, the flow control settings can be applied. The flow control setting that is finally applied to an interface depends on auto-negotiation. With the ON option, the peer negotiates the flow control, but the appliance then forces two-way flow control for the interface.

* `haheartbeat` - In a High Availability (HA) or Cluster configuration, configure the interface for sending heartbeats. In an HA or Cluster configuration, an interface that has HA Heartbeat disabled should not send the heartbeats.

* `hamonitor` - In a High Availability (HA) configuration, monitor the interface for failure events. In an HA configuration, an interface that has HA MON enabled and is not bound to any Failover Interface Set (FIS), is a critical interface. Failure or disabling of any critical interface triggers HA failover.

* `ifalias` - Alias name for the interface. Used only to enhance readability. To perform any operations, you have to specify the interface ID.

* `lacpkey` - Integer identifying the LACP LA channel to which the interface is to be bound. For an LA channel of the Citrix ADC, this digit specifies the variable x of an LA channel in LA/x notation, where x can range from 1 to 8. For example, if you specify 3 as the LACP key for an LA channel, the interface is bound to the LA channel LA/3. For an LA channel of a cluster configuration, this digit specifies the variable y of a cluster LA channel in CLA/(y-4) notation, where y can range from 5 to 8. For example, if you specify 6 as the LACP key for a cluster LA channel, the interface is bound to the cluster LA channel CLA/2.

* `lacpmode` - Bind the interface to a LA channel created by the Link Aggregation control protocol (LACP). Available settings:
  * `Active` - The LA channel port of the Citrix ADC generates LACPDU messages on a regular basis, regardless of any need expressed by its peer device to receive them
  * `Passive` - The LA channel port of the Citrix ADC does not transmit LACPDU messages unless the peer device port is in the active mode. That is, the port does not speak unless spoken to
  * `Disabled` - Unbinds the interface from the LA channel. If this is the only interface in the LA channel, the LA channel is removed

* `lacppriority` - LACP port priority, expressed as an integer. The lower the number, the higher the priority. The Citrix ADC limits the number of interfaces in an LA channel to sixteen.

* `lacptimeout` - Interval at which the Citrix ADC sends LACPDU messages to the peer device on the LA channel. Available settings:
  * `LONG` - 30 seconds
  * `SHORT` - 1 second

* `lagtype` - Type of entity (Citrix ADC or cluster configuration) for which to create the channel.

* `linkredundancy` - Link Redundancy for Cluster LAG.

* `lldpmode` - Link Layer Discovery Protocol (LLDP) mode for an interface. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels.

* `lrsetpriority` - LRSET port priority, expressed as an integer ranging from 1 to 1024. The highest priority is 1. The Citrix ADC limits the number of interfaces in an LRSET to 8. Within a LRSET the highest LR Priority Interface is considered as the first candidate for the Active interface, if the interface is UP.

* `mtu` - The Maximum Transmission Unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes CRC, that can be transmitted and received by an interface. The default value of MTU is 1500 on all the interface of Citrix ADC, some Cloud Platforms will restrict Citrix ADC to use the lesser default value. Any MTU value more than 1500 is called Jumbo MTU and will make the interface as jumbo enabled. The Maximum Jumbo MTU in Citrix ADC is 9216, however, some Virtualized / Cloud Platforms will have lesser Maximum Jumbo MTU Value (9000). In the case of Cluster, the Backplane interface requires an MTU value of 78 bytes more than the Max MTU configured on any other Data-Plane Interface.

* `ringsize` - The receive ringsize of the interface. A higher number provides more number of buffers in handling incoming traffic.

* `ringtype` - The receive ringtype of the interface (Fixed or Elastic). A fixed ring type pre-allocates configured number of buffers irrespective of traffic rate. In contrast, an elastic ring, expands and shrinks based on incoming traffic rate.

* `speed` - Ethernet speed of the interface, in Mbps. If you set the speed as AUTO, the Citrix ADC attempts to auto-negotiate or auto-sense the link speed of the interface when it is UP. You must enable auto negotiation on the interface. If you set a speed other than AUTO, you must specify the same speed for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors. Some interfaces do not support certain speeds.

* `tagall` - Add a four-byte 802.1q tag to every packet sent on this interface. The ON setting applies the tag for this interface's native VLAN. OFF applies the tag for all VLANs other than the native VLAN.

* `throughput` - Low threshold value for the throughput of the interface, in Mbps. In an HA configuration, failover is triggered if the interface has HA MON enabled and the throughput is below the specified the threshold.

* `trunk` - This argument is deprecated by tagall.

* `trunkallowedvlan` - VLAN ID or range of VLAN IDs will be allowed on this trunk interface. In the command line interface, separate the range with a hyphen. For example: 40-90.

* `trunkmode` - Accept and send 802.1q VLAN tagged packets, based on Allowed Vlan List of this interface.

### Read-only interface metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_interface` resource). They are Computed / GET-only. Any attribute the appliance does not return is `null`.

* `devicename` - Name of the interface.
* `unit` - Unit number for this interface, signifying the sequence number in which this interface is discovered on this Citrix ADC.
* `description` - Display the type of interface, the speeds at which this interface can operate, and, if applicable, the type of SFP.
* `flags` - Flags for this interface. Used for communicating the device states.
* `actualmtu` - MTU for this interface (the largest frame that can transit this interface).
* `vlan` - Native VLAN for this interface.
* `mac` - MAC address for this interface.
* `uptime` - Duration for which the interface has been UP. Reset when the interface state changes to DOWN.
* `downtime` - Duration for which the interface has been DOWN. Reset when the interface state changes to UP.
* `actualringsize` - Actual receive ringsize of the interface.
* `reqmedia` - Requested media setting for this interface.
* `reqspeed` - Requested speed setting for this interface.
* `reqduplex` - Requested duplex setting for this interface.
* `reqflowcontrol` - Requested flow control setting for this interface.
* `actmedia` - Actual media setting for this interface.
* `actspeed` - Actual speed setting for this interface.
* `actduplex` - Actual duplex setting for this interface.
* `actflowctl` - Actual flow control setting for this interface.
* `mode` - Interface Link Aggregation mode (Auto/Manual) setting.
* `autonegresult` - Actual auto-negotiation setting for this interface.
* `tagged` - VLAN tags setting on this channel.
* `taggedany` - Interface setting to accept/drop all tagged packets.
* `taggedautolearn` - Dynamic VLAN membership autolearning enabled or disabled on this interface.
* `hangdetect` - Hang detection enabled or disabled for this interface.
* `hangreset` - Hang reset enabled or disabled for this interface.
* `linkstate` - The current state of the link associated with the interface.
* `intfstate` - Current state of the specified interface.
* `rxpackets` - Number of packets received by an interface.
* `rxbytes` - Number of bytes received by an interface.
* `rxerrors` - Number of inbound packets dropped by the hardware on a specified interface.
* `rxdrops` - Number of inbound packets dropped by the specified interface.
* `txpackets` - Number of packets transmitted by an interface.
* `txbytes` - Number of bytes transmitted by an interface.
* `txerrors` - Number of outbound packets dropped by the hardware on a specified interface.
* `txdrops` - Number of packets dropped in transmission by the specified interface.
* `indisc` - Number of error-free inbound packets discarded because of a lack of resources.
* `outdisc` - Number of error-free outbound packets discarded because of a lack of resources.
* `fctls` - Number of times flow control is performed because of received pause frames.
* `hangs` - Number of times the specified interface detected hangs in the transmit and receive paths.
* `stsstalls` - Number of times the status updates for a specified interface were stalled.
* `txstalls` - Number of times the interface stalled when transmitting packets.
* `rxstalls` - Number of times the interface stalled when receiving packets.
* `bdgmacmoved` - Number of MAC moves between ports.
* `bdgmuted` - Number of times the interface stopped transmitting and receiving packets because of MAC moves between ports.
* `vmac` - Virtual MAC of this interface.
* `vmac6` - Virtual MAC for IPv6 of this interface.
* `reqthroughput` - Minimum required throughput for an interface.
* `actthroughput` - Actual throughput for the interface.
* `backplane` - The cluster backplane status of the interface.
* `ifnum` - Contains the LA Master, if the interface is part of LA channel. A list of strings.
* `cleartime` - Time since the interface stats are cleared last time.
* `slavestate` - State of the member interfaces.
* `slavemedia` - Media type of the member interfaces.
* `slavespeed` - Speed of the member interfaces.
* `slaveduplex` - Duplex of the member interfaces.
* `slaveflowctl` - Flowcontrol of the member interfaces.
* `slavetime` - UP time of the member interfaces.
* `intftype` - Interface Type (virtual, physical or loopback).
* `svmcmd` - Attribute to identify the source of cmd; set to 1 when SVM fires the nitro cmd.
* `lacpactormode` - LACP actor mode (DISABLED, ACTIVE, PASSIVE).
* `lacpactortimeout` - Interval at which the Citrix ADC sends LACPDU messages (LONG, SHORT).
* `lacpactorpriority` - LACP Actor Priority.
* `lacpactorportno` - LACP Actor port number.
* `lacppartnerstate` - LACP Partner State.
* `lacppartnertimeout` - The timeout value for the information received in LACPDUs (LONG, SHORT).
* `lacppartneraggregation` - The Aggregation flag of the partner.
* `lacppartnerinsync` - The Synchronization flag of the partner.
* `lacppartnercollecting` - The Collecting flag of the partner.
* `lacppartnerdistributing` - The Distributing flag of the partner.
* `lacppartnerdefaulted` - Whether the Receive Machine of the partner entered the Defaulted state.
* `lacppartnerexpired` - Whether the Receive Machine of the partner entered the Expired state.
* `lacppartnerpriority` - LACP Partner Priority.
* `lacppartnersystemmac` - LACP Partner System MAC.
* `lacppartnersystempriority` - LACP Partner System Priority.
* `lacppartnerportno` - LACP Partner Port number.
* `lacppartnerkey` - LACP Partner Key.
* `lacpactoraggregation` - The Aggregation flag of the actor.
* `lacpactorinsync` - The Synchronization flag of the actor.
* `lacpactorcollecting` - The Collecting flag of the actor.
* `lacpactordistributing` - The Distributing flag of the actor.
* `lacpportmuxstate` - LACP Port MUX state.
* `lacpportrxstat` - LACP Port RX state.
* `lacpportselectstate` - LACP Port SELECT state (SELECTED, UNSELECTED, STANDBY).
* `lractiveintf` - LR set member interface state (active/inactive).
