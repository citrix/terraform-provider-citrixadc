---
subcategory: "Network"
---

# Resource: channel

The channel resource is used to create channel.


## Example usage

```hcl
resource "citrixadc_channel" "tf_channel" {
  channel_id = "LA/3"
  tagall     = "ON"
  speed      = "1000"
}
```


## Argument Reference

* `channel_id` - (Required) ID for the LA channel or cluster LA channel or LR channel to be created. Specify an LA channel in LA/x notation, where x can range from 1 to 8 or cluster LA channel in CLA/x notation or Link redundant channel in LR/x notation, where x can range from 1 to 4. Cannot be changed after the LA channel is created.
* `ifnum` - (Optional) Interfaces to be bound to the LA channel of a Citrix ADC or to the LA channel of a cluster configuration. For an LA channel of a Citrix ADC, specify an interface in C/U notation (for example, 1/3). For an LA channel of a cluster configuration, specify an interface in N/C/U notation (for example, 2/1/3). where C can take one of the following values: * 0 - Indicates a management interface. * 1 - Indicates a 1 Gbps port. * 10 - Indicates a 10 Gbps port. U is a unique integer for representing an interface in a particular port group. N is the ID of the node to which an interface belongs in a cluster configuration. Use spaces to separate multiple entries.
* `state` - (Optional) Enable or disable the LA channel. Possible values: [ ENABLED, DISABLED ]
* `mode` - (Optional) The initital mode for the LA channel. Possible values: [ MANUAL, AUTO ]
* `conndistr` - (Optional) The 'connection' distribution mode for the LA channel. Possible values: [ DISABLED, ENABLED ]
* `macdistr` - (Optional) The  'MAC' distribution mode for the LA channel. Possible values: [ SOURCE, DESTINATION, BOTH ]
* `lamac` - (Optional) Specifies a MAC address for the LA channels configured in Citrix ADC virtual appliances (VPX). This MAC address is persistent after each reboot. If you don't specify this parameter, a MAC address is generated randomly for each LA channel. These MAC addresses change after each reboot.
* `speed` - (Optional) Ethernet speed of the channel, in Mbps. If the speed of any bound interface is greater than or equal to the value set for this parameter, the state of the interface is UP. Otherwise, the state is INACTIVE. Bound Interfaces whose state is INACTIVE do not process any traffic. Possible values: [ AUTO, 10, 100, 1000, 10000, 25000, 40000, 50000, 100000 ]
* `flowctl` - (Optional) Specifies the flow control type for this LA channel to manage the flow of frames. Flow control is a function as mentioned in clause 31 of the IEEE 802.3 standard. Flow control allows congested ports to pause traffic from the peer device. Flow control is achieved by sending PAUSE frames. Possible values: [ OFF, RX, TX, RXTX, ON ]
* `hamonitor` - (Optional) In a High Availability (HA) configuration, monitor the LA channel for failure events. Failure of any LA channel that has HA MON enabled triggers HA failover. Possible values: [ on, off ]
* `haheartbeat` - (Optional) In a High Availability (HA) configuration, configure the LA channel for sending heartbeats. LA channel that has HA Heartbeat disabled should not send the heartbeats. Possible values: [ on, off ]
* `tagall` - (Optional) Adds a four-byte 802.1q tag to every packet sent on this channel.  The ON setting applies tags for all VLANs that are bound to this channel. OFF applies the tag for all VLANs other than the native VLAN. Possible values: [ on, off ]
* `trunk` - (Optional) This is deprecated by tagall. Possible values: [ on, off ]
* `ifalias` - (Optional) Alias name for the LA channel. Used only to enhance readability. To perform any operations, you have to specify the LA channel ID. Maximum length =  31
* `throughput` - (Optional) Low threshold value for the throughput of the LA channel, in Mbps. In an high availability (HA) configuration, failover is triggered when the LA channel has HA MON enabled and the throughput is below the specified threshold. Minimum value =  0 Maximum value =  160000
* `bandwidthhigh` - (Optional) High threshold value for the bandwidth usage of the LA channel, in Mbps. The Citrix ADC generates an SNMP trap message when the bandwidth usage of the LA channel is greater than or equal to the specified high threshold value. Minimum value =  0 Maximum value =  160000
* `bandwidthnormal` - (Optional) Normal threshold value for the bandwidth usage of the LA channel, in Mbps. When the bandwidth usage of the LA channel returns to less than or equal to the specified normal threshold after exceeding the high threshold, the Citrix ADC generates an SNMP trap message to indicate that the bandwidth usage has returned to normal. Minimum value =  0 Maximum value =  160000
* `mtu` - (Optional) The Maximum Transmission Unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes CRC, that can be transmitted and received by an interface. The default value of MTU is 1500 on all the interface of Citrix ADC, some Cloud Platforms will restrict Citrix ADC to use the lesser default value. Any MTU value more than 1500 is called Jumbo MTU and will make the interface as jumbo enabled. The Maximum Jumbo MTU in Citrix ADC is 9216, however, some Virtualized / Cloud Platforms will have lesser Maximum Jumbo MTU Value (9000). In the case of Cluster, the Backplane interface requires an MTU value of 78 bytes more than the Max MTU configured on any other Data-Plane Interface. When the Data plane interfaces are all at default 1500 MTU, Cluster Back Plane will be automatically set to 1578 (1500 + 78) MTU. If a Backplane interface is reset to Data Plane Interface, then the 1578 MTU will be automatically reset to the default MTU of 1500(or whatever lesser default value). If any data plane interface of a Cluster is configured with a Jumbo MTU ( > 1500), then all backplane interfaces require to be configured with a minimum MTU of 'Highest Data Plane MTU in the Cluster + 78'. That makes the maximum Jumbo MTU for any Data-Plane Interface in a Cluster System to be '9138 (9216 - 78)., where 9216 is the maximum Jumbo MTU. On certain Virtualized / Cloud Platforms, the maximum  possible MTU is restricted to a lesser value, Similar calculation can be applied, Maximum Data Plane MTU in Cluster = (Maximum possible MTU - 78). . Minimum value =  1500 Maximum value =  9216
* `lrminthroughput` - (Optional) Specifies the minimum throughput threshold (in Mbps) to be met by the active subchannel. Setting this parameter automatically divides an LACP channel into logical subchannels, with one subchannel active and the others in standby mode.  When the maximum supported throughput of the active channel falls below the lrMinThroughput value, link failover occurs and a standby subchannel becomes active. Minimum value =  0 Maximum value =  80000
* `linkredundancy` - (Optional) Link Redundancy for Cluster LAG. Possible values: [ on, off ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the channel. It has the same value as the `channel_id` attribute.


## Import

A channel can be imported using its name, e.g.

```shell
terraform import citrixadc_channel.tf_channel LA/3
```
