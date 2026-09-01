---
subcategory: "Network"
---

# Data Source: channel

The channel data source allows you to retrieve information about Link Aggregation (LA) channels.


## Example usage

```terraform
data "citrixadc_channel" "tf_channel" {
  channelid = "LA/3"
}

output "tagall" {
  value = data.citrixadc_channel.tf_channel.tagall
}

output "speed" {
  value = data.citrixadc_channel.tf_channel.speed
}
```


## Argument Reference

* `channelid` - (Required) ID for the LA channel or cluster LA channel or LR channel to be created. Specify an LA channel in LA/x notation, where x can range from 1 to 8 or cluster LA channel in CLA/x notation or Link redundant channel in LR/x notation, where x can range from 1 to 4. Cannot be changed after the LA channel is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `bandwidthhigh` - High threshold value for the bandwidth usage of the LA channel, in Mbps. The Citrix ADC generates an SNMP trap message when the bandwidth usage of the LA channel is greater than or equal to the specified high threshold value.
* `bandwidthnormal` - Normal threshold value for the bandwidth usage of the LA channel, in Mbps. When the bandwidth usage of the LA channel returns to less than or equal to the specified normal threshold after exceeding the high threshold, the Citrix ADC generates an SNMP trap message to indicate that the bandwidth usage has returned to normal.
* `conndistr` - The 'connection' distribution mode for the LA channel.
* `flowctl` - Specifies the flow control type for this LA channel to manage the flow of frames. Flow control is a function as mentioned in clause 31 of the IEEE 802.3 standard. Flow control allows congested ports to pause traffic from the peer device. Flow control is achieved by sending PAUSE frames.
* `haheartbeat` - In a High Availability (HA) configuration, configure the LA channel for sending heartbeats. LA channel that has HA Heartbeat disabled should not send the heartbeats.
* `hamonitor` - In a High Availability (HA) configuration, monitor the LA channel for failure events. Failure of any LA channel that has HA MON enabled triggers HA failover.
* `ifalias` - Alias name for the LA channel. Used only to enhance readability. To perform any operations, you have to specify the LA channel ID.
* `ifnum` - Interfaces to be bound to the LA channel of a Citrix ADC or to the LA channel of a cluster configuration. For an LA channel of a Citrix ADC, specify an interface in C/U notation (for example, 1/3). For an LA channel of a cluster configuration, specify an interface in N/C/U notation (for example, 2/1/3). where C can take one of the following values: 0 - Indicates a management interface, 1 - Indicates a 1 Gbps port, 10 - Indicates a 10 Gbps port. U is a unique integer for representing an interface in a particular port group. N is the ID of the node to which an interface belongs in a cluster configuration. Use spaces to separate multiple entries.
* `lamac` - Specifies a MAC address for the LA channels configured in Citrix ADC virtual appliances (VPX). This MAC address is persistent after each reboot. If you don't specify this parameter, a MAC address is generated randomly for each LA channel. These MAC addresses change after each reboot.
* `linkredundancy` - Link Redundancy for Cluster LAG.
* `lrminthroughput` - Specifies the minimum throughput threshold (in Mbps) to be met by the active subchannel. Setting this parameter automatically divides an LACP channel into logical subchannels, with one subchannel active and the others in standby mode.  When the maximum supported throughput of the active channel falls below the lrMinThroughput value, link failover occurs and a standby subchannel becomes active.
* `macdistr` - The  'MAC' distribution mode for the LA channel.
* `mode` - The initital mode for the LA channel.
* `mtu` - The Maximum Transmission Unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes CRC, that can be transmitted and received by an interface. The default value of MTU is 1500 on all the interface of Citrix ADC, some Cloud Platforms will restrict Citrix ADC to use the lesser default value. Any MTU value more than 1500 is called Jumbo MTU and will make the interface as jumbo enabled. The Maximum Jumbo MTU in Citrix ADC is 9216, however, some Virtualized / Cloud Platforms will have lesser Maximum Jumbo MTU Value (9000). In the case of Cluster, the Backplane interface requires an MTU value of 78 bytes more than the Max MTU configured on any other Data-Plane Interface. When the Data plane interfaces are all at default 1500 MTU, Cluster Back Plane will be automatically set to 1578 (1500 + 78) MTU. If a Backplane interface is reset to Data Plane Interface, then the 1578 MTU will be automatically reset to the default MTU of 1500(or whatever lesser default value). If any data plane interface of a Cluster is configured with a Jumbo MTU ( > 1500), then all backplane interfaces require to be configured with a minimum MTU of 'Highest Data Plane MTU in the Cluster + 78'. That makes the maximum Jumbo MTU for any Data-Plane Interface in a Cluster System to be '9138 (9216 - 78)., where 9216 is the maximum Jumbo MTU. On certain Virtualized / Cloud Platforms, the maximum  possible MTU is restricted to a lesser value, Similar calculation can be applied, Maximum Data Plane MTU in Cluster = (Maximum possible MTU - 78).
* `speed` - Ethernet speed of the channel, in Mbps. If the speed of any bound interface is greater than or equal to the value set for this parameter, the state of the interface is UP. Otherwise, the state is INACTIVE. Bound Interfaces whose state is INACTIVE do not process any traffic.
* `state` - Enable or disable the LA channel.
* `tagall` - Adds a four-byte 802.1q tag to every packet sent on this channel.  The ON setting applies tags for all VLANs that are bound to this channel. OFF applies the tag for all VLANs other than the native VLAN.
* `throughput` - Low threshold value for the throughput of the LA channel, in Mbps. In an high availability (HA) configuration, failover is triggered when the LA channel has HA MON enabled and the throughput is below the specified threshold.
* `trunk` - This is deprecated by tagall.
* `id` - The id of the channel. It has the same value as the `channelid` attribute.

### Read-only channel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_channel` resource). They are GET-only / Computed, and any attribute the appliance does not return is `null`.

* `devicename` - LA channel name in form LA/x, where x is channel ID, which ranges from 1 to 8 or LR channel name in form LR/x, where x is channel ID, which ranges from 1 to 4.
* `unit` - Unit number of the channel. This is an internal reference number that the Citrix ADC uses to identify the channel.
* `description` - The IEEE standard that the channel is based on.
* `flags` - Flags of this channel.
* `actualmtu` - MTU of the channel. This is the maximum frame size that the channel can process.
* `vlan` - Native VLAN of the channel.
* `mac` - MAC address of the channel.
* `uptime` - Duration for which the channel is UP. (Example: 3 hours 1 minute 1 second). This value is reset when the channel state changes to DOWN.
* `downtime` - Duration for which the channel is DOWN. (Example: 3 hours 1 minute 1 second). This value is reset when the channel state changes to UP.
* `reqmedia` - Requested media setting for this channel. Since there is no media associated with LA, the displayed values carry no significance.
* `reqspeed` - Requested speed setting for this channel. Since no media are associated with LA, this speed is used to determine the threshold for the slave interfaces. If the speed of the member interface is less than the requested speed, that interface is considered inactive.
* `reqduplex` - Requested duplex setting for this channel. Since no media are associated with LA, the displayed values carry no significance.
* `reqflowcontrol` - Requested flow control setting for this channel. Since no media are associated with LA, the displayed values carry no significance.
* `media` - Requested media setting for this interface.
* `actspeed` - Actual speed setting for this channel.
* `duplex` - Actualduplex setting for this interface.
* `actflowctl` - Actual flow control setting for this channel.
* `lamode` - The  mode(AUTO/MANNUAL) for the LA channel.
* `autoneg` - Requested auto negotiation setting for this channel. Since no media are associated with LA, this setting has no effect.
* `autonegresult` - Actual  auto negotiation setting for this channel.
* `tagged` - VLAN tags setting on this channel.
* `taggedany` - Channel setting to accept/drop all tagged packets.
* `taggedautolearn` - Dynaminc vlan membership on this channel.
* `hangdetect` - Hang detect for this channel.
* `hangreset` - Hang reset for this channel.
* `linkstate` - The current state of the link associated with the interface. For logical interfaces (LA), the state of the link is dependent on the state of the slave interfaces. For the link to be UP at least one of the slave interfaces needs to be UP.
* `intfstate` - Current state of the specified interface.  The interface state set to UP only if the link state is UP and administrative state is ENABLED.
* `rxpackets` - Number of bytes received by all the slave interfaces of the channel since the Citrix ADC was started or the interface statistics were cleared.
* `rxbytes` - Number of packets received by all member interfaces since the Citrix ADC was started or the interface statistics were cleared.
* `rxerrors` - Number of inbound packets dropped by the hardware of the slave interfaces since the Citrix ADC was started or the interface statistics were cleared. Possible causes of dropped packets are CRC, length (undersize or oversize), and alignment errors.
* `rxdrops` - Number of inbound packets dropped by the channel's slave interfaces. Commonly dropped packets are multicast frames, spanning tree BPDUs, packets destined to a MAC not owned by the Citrix ADC when L2 mode is disabled, or packets tagged for a VLAN that is not bound to the interface.  In most healthy networks, this statistic increments at a steady rate regardless of traffic load.  A sharp spike in dropped packets generally indicates an issue with connected L2 switches, such as a forwarding database overflow resulting in packets being broadcast on all ports.
* `txpackets` - Number of packets transmitted by slave interfaces of a channel since the Citrix ADC was started or the interface statistics were cleared.
* `txbytes` - Number of bytes transmitted by slave interfaces of a channel since the Citrix ADC was started or the interface statistics were cleared.
* `txerrors` - Number of outbound packets dropped by the hardware of a channel's slave interfaces since the Citrix ADC was started or the interface statistics were cleared. Possible causes of dropped packets are length (undersize or oversize) errors and lack of resources.
* `txdrops` - Number of packets dropped in transmission by a channel's slave interfaces for one of the following reasons:
* `indisc` - Number of error-free inbound packets discarded by a channel's slave interfaces because of a lack of resources (for example, insufficient receive buffers).
* `outdisc` - Number of error-free outbound packets discarded by a channel's slave interfaces because of a lack of resources. This statistic is not available on:
* `fctls` - Number of times flow control is performed on a channel's slave interfaces because of pause frames.
* `hangs` - Number of hangs that occurred on the channel's slave interfaces.
* `stsstalls` - Number of status stalls that occurred on the channel's slave interfaces.
* `txstalls` - Number of Tx stalls happened that occurred on the channel's slave interfaces.
* `rxstalls` - Number of Rx stalls that occurred on the channel's slave interfaces.
* `bdgmuted` - Number of times a channel's slave interfaces stopped transmitting and receiving packets because of MAC moves between ports.
* `vmac` - Virtual MAC of this channel.
* `vmac6` - Virtual MAC for IPv6 on this interface.
* `reqthroughput` - Minimum required throughput for an interface. Failover is triggered if the operating throughput of a Link Aggregation (LA) channel for which HAMON is ON falls below this value.
* `actthroughput` - Actual throughput for the interface.
* `backplane` - The cluster backplane status of the LA. If the status is enabled, the LA is part of the cluster backplane. By default, the backplane status is disabled.
* `cleartime` - Time since the interface stats are cleared last time.
* `lacpmode` - The LACP mode of the specified interface. The possible values are:
* `lacptimeout` - Time to wait for the LACPDU.  If a LACPDU is not received within this interval, the Citrix ADC markes the link partner port as DOWN. Possible values: Long and Short. Long lacptimeout is 90 sec and Short LACP timeout is 3 sec.
* `lacpactorpriority` - LACP Actor Priority. A LACP port priority is configured on each port using LACP. LACP uses the port priority with the port number to form the port identifier. The port priority determines which ports should be put in standby mode when there is a hardware limitation that prevents all compatible ports from aggregating.
* `lacpactorportno` - LACP Actor port number. LACP uses the port priority with the port number to form the port identifier.
* `lacppartnerstate` - LACP Partner State. Whether the port is in Active or Passive negotiating state.
* `lacppartnertimeout` - The timeout value for the information revieved in LACPDUs. It can have values as SHORT or LONG. The SHORT timeout is 3s and the LONG timeout is 90s.
* `lacppartneraggregation` - The Aggregation flag indicates that the participant will allow the link to be used as part of an aggregate. Otherwise the link is to be used as an individual link, i.e. not aggregated with any other.
* `lacppartnerinsync` - The Synchronization flag indicates that the transmitting participant.s mux component is in sync with the system id and key information transmitted.
* `lacppartnercollecting` - The Collecting flag indicates that the participant.s collector, i.e. the reception component of the mux, is definitely on. If set the flag communicates collecting.
* `lacppartnerdistributing` - The Distributing flag indicates that the participant.s distributor is not definitely off. If reset the flag indicates not distributing.
* `lacppartnerdefaulted` - If the timer expires in the Expired state, the Receive Machine enters the Defaulted state.
* `lacppartnerexpired` - If the LACPDUs are received for timeout period, the Receive Machine enters the Expired state and the timer is restarted with the timeout value of SHORT timeout.
* `lacppartnerpriority` - LACP Partner Priority. A LACP port priority is configured on each port using LACP. LACP uses the port priority with the port number to form the port identifier.
* `lacppartnersystemmac` - LACP Partner System MAC.
* `lacppartnersystempriority` - LACP Partner System Priority. The LACP partner's system priority. The values for the priority range from 0 to 65535. The lower the value, the higher the system priority. The switch with the lower system priority value determines which links between LACP partner are active and which are in the standby for each LACP Channel.
* `lacppartnerportno` - LACP Partner Port number. LACP uses the port priority with the port number to form the port identifier.
* `lacppartnerkey` - LACP Partner Key. The LACP key used by the partner port.
* `lacpactoraggregation` - The Aggregation flag indicates that the participant will allow the link to be used as part of an aggregate. Otherwise the link is to be used as an individual link, i.e. not aggregated with any other.
* `lacpactorinsync` - The Synchronization flag indicates that the transmitting participant.s mux component is in sync with the system id and key information transmitted.
* `lacpactorcollecting` - The Collecting flag indicates that the participant.s collector, i.e. the reception component of the mux, is definitely on. If set the flag communicates collecting.
* `lacpactordistributing` - The Distributing flag indicates that the participant.s distributor is not definitely off. If reset the flag indicates not distributing.
* `lacpportmuxstate` - LACP Port MUX state. The state of the MUX control machine. The  Mux Control Machine attaches the physical port to an aggregate port, using the Selection Logic to choose an appropriate port, and turns the distributor and collector for the physical port on or off as required by protocol information.
* `lacpportrxstat` - LACP Port RX state. The state of the Receive machine. The Receive Machine maintains partner information, recording protocol information from LACPDUs sent by remote partner(s). Received information is subject to a timeout, and if sufficient time elapses the receive machine will revert to using default partner information.
* `lacpportselectstate` - LACP Port SELECT state. The state of the SELECT state machine, It could be SELECTED or UNSELECTED.
* `lldpmode` - Link Layer Discovery Protocol (LLDP) mode for an interface. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels.
