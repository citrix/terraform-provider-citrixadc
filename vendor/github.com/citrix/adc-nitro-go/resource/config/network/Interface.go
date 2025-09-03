/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package network

/**
* Configuration for interface resource.
*/
type Interface struct {
	/**
	* Interface number, in C/U format, where C can take one of the following values:
		* 0 - Indicates a management interface.
		* 1 - Indicates a 1 Gbps port.
		* 10 - Indicates a 10 Gbps port.
		* LA - Indicates a link aggregation port.
		* LO - Indicates a loop back port.
		U is a unique integer for representing an interface in a particular port group.
	*/
	Id string `json:"id,omitempty"`
	/**
	* Ethernet speed of the interface, in Mbps.
		Notes:
		* If you set the speed as AUTO, the Citrix ADC attempts to auto-negotiate or auto-sense the link speed of the interface when it is UP. You must enable auto negotiation on the interface.
		* If you set a speed other than AUTO, you must specify the same speed for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors.
		Some interfaces do not support certain speeds. If you specify an unsupported speed, an error message appears.
	*/
	Speed string `json:"speed,omitempty"`
	/**
	* The duplex mode for the interface. Notes:* If you set the duplex mode to AUTO, the Citrix ADC attempts to auto-negotiate the duplex mode of the interface when it is UP. You must enable auto negotiation on the interface. If you set a duplex mode other than AUTO, you must specify the same duplex mode for the peer network device. Mismatched speed and duplex settings between the peer devices of a link lead to link errors, packet loss, and other errors.
	*/
	Duplex string `json:"duplex,omitempty"`
	/**
	* 802.3x flow control setting for the interface.  The 802.3x specification does not define flow control for 10 Mbps and 100 Mbps speeds, but if a Gigabit Ethernet interface operates at those speeds, the flow control settings can be applied. The flow control setting that is finally applied to an interface depends on auto-negotiation. With the ON option, the peer negotiates the flow control, but the appliance then forces two-way flow control for the interface.
	*/
	Flowctl string `json:"flowctl,omitempty"`
	/**
	* Auto-negotiation state of the interface. With the ENABLED setting, the Citrix ADC auto-negotiates the speed and duplex settings with the peer network device on the link. The Citrix ADC appliance auto-negotiates the settings of only those parameters (speed or duplex mode) for which the value is set as AUTO.
	*/
	Autoneg string `json:"autoneg,omitempty"`
	/**
	* In a High Availability (HA) configuration, monitor the interface for failure events. In an HA configuration, an interface that has HA MON enabled and is not bound to any Failover Interface Set (FIS), is a critical interface. Failure or disabling of any critical interface triggers HA failover.
	*/
	Hamonitor string `json:"hamonitor,omitempty"`
	/**
	* In a High Availability (HA) or Cluster configuration, configure the interface for sending heartbeats. In an HA or Cluster configuration, an interface that has HA Heartbeat disabled should not send the heartbeats.
	*/
	Haheartbeat string `json:"haheartbeat,omitempty"`
	/**
	* The Maximum Transmission Unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes CRC, that can be transmitted and received by an interface. The default value of MTU is 1500 on all the interface of Citrix ADC, some Cloud Platforms will restrict Citrix ADC to use the lesser default value. Any MTU value more than 1500 is called Jumbo MTU and will make the interface as jumbo enabled. The Maximum Jumbo MTU in Citrix ADC is 9216, however, some Virtualized / Cloud Platforms will have lesser Maximum Jumbo MTU Value (9000). In the case of Cluster, the Backplane interface requires an MTU value of 78 bytes more than the Max MTU configured on any other Data-Plane Interface. When the Data plane interfaces are all at default 1500 MTU, Cluster Back Plane will be automatically set to 1578 (1500 + 78) MTU. If a Backplane interface is reset to Data Plane Interface, then the 1578 MTU will be automatically reset to the default MTU of 1500(or whatever lesser default value). If any data plane interface of a Cluster is configured with a Jumbo MTU ( > 1500), then all backplane interfaces require to be configured with a minimum MTU of 'Highest Data Plane MTU in the Cluster + 78'. That makes the maximum Jumbo MTU for any Data-Plane Interface in a Cluster System to be '9138 (9216 - 78)., where 9216 is the maximum Jumbo MTU. On certain Virtualized / Cloud Platforms, the maximum  possible MTU is restricted to a lesser value, Similar calculation can be applied, Maximum Data Plane MTU in Cluster = (Maximum possible MTU - 78).
	*/
	Mtu int `json:"mtu,omitempty"`
	/**
	* The receive ringsize of the interface. A higher number provides more number of buffers in handling incoming traffic.
	*/
	Ringsize int `json:"ringsize,omitempty"`
	/**
	* The receive ringtype of the interface (Fixed or Elastic). A fixed ring type pre-allocates configured number of buffers irrespective of traffic rate. In contrast, an elastic ring, expands and shrinks based on incoming traffic rate.
	*/
	Ringtype string `json:"ringtype,omitempty"`
	/**
	* Add a four-byte 802.1q tag to every packet sent on this interface.  The ON setting applies the tag for this interface's native VLAN. OFF applies the tag for all VLANs other than the native VLAN.
	*/
	Tagall string `json:"tagall,omitempty"`
	/**
	* This argument is deprecated by tagall.
	*/
	Trunk string `json:"trunk,omitempty"`
	/**
	* Accept and send 802.1q VLAN tagged packets, based on Allowed Vlan List of this interface.
	*/
	Trunkmode string `json:"trunkmode,omitempty"`
	/**
	* VLAN ID or range of VLAN IDs will be allowed on this trunk interface. In the command line interface, separate the range with a hyphen. For example: 40-90.
	*/
	Trunkallowedvlan []string `json:"trunkallowedvlan,omitempty"`
	/**
	* Bind the interface to a LA channel created by the Link Aggregation control protocol (LACP).
		Available settings function as follows:
		* Active - The LA channel port of the Citrix ADC generates LACPDU messages on a regular basis, regardless of any need expressed by its peer device to receive them.
		* Passive - The LA channel port of the Citrix ADC does not transmit LACPDU messages unless the peer device port is in the active mode. That is, the port does not speak unless spoken to.
		* Disabled - Unbinds the interface from the LA channel. If this is the only interface in the LA channel, the LA channel is removed.
	*/
	Lacpmode string `json:"lacpmode,omitempty"`
	/**
	* Integer identifying the LACP LA channel to which the interface is to be bound.
		For an LA channel of the Citrix ADC, this digit specifies the variable x of an LA channel in LA/x notation, where x can range from 1 to 8. For example, if you specify 3 as the LACP key for an LA channel, the interface is bound to the LA channel LA/3.
		For an LA channel of a cluster configuration, this digit specifies the variable y of a cluster LA channel in CLA/(y-4) notation, where y can range from 5 to 8. For example, if you specify 6 as the LACP key for a cluster LA channel, the interface is bound to the cluster LA channel CLA/2.
	*/
	Lacpkey int `json:"lacpkey,omitempty"`
	/**
	* Type of entity (Citrix ADC or cluster configuration) for which to create the channel.
	*/
	Lagtype string `json:"lagtype,omitempty"`
	/**
	* LACP port priority, expressed as an integer. The lower the number, the higher the priority. The Citrix ADC limits the number of interfaces in an LA channel to sixteen.
	*/
	Lacppriority int `json:"lacppriority,omitempty"`
	/**
	* Interval at which the Citrix ADC sends LACPDU messages to the peer device on the LA channel.
		Available settings function as follows:
		LONG - 30 seconds.
		SHORT - 1 second.
	*/
	Lacptimeout string `json:"lacptimeout,omitempty"`
	/**
	* Alias name for the interface. Used only to enhance readability. To perform any operations, you have to specify the interface ID.
	*/
	Ifalias string `json:"ifalias,omitempty"`
	/**
	* Low threshold value for the throughput of the interface, in Mbps. In an HA configuration, failover is triggered if the interface has HA MON enabled and the throughput is below the specified the threshold.
	*/
	Throughput int `json:"throughput,omitempty"`
	/**
	* Link Redundancy for Cluster LAG.
	*/
	Linkredundancy string `json:"linkredundancy,omitempty"`
	/**
	* High threshold value for the bandwidth usage of the interface, in Mbps. The Citrix ADC generates an SNMP trap message when the bandwidth usage of the interface is greater than or equal to the specified high threshold value.
	*/
	Bandwidthhigh int `json:"bandwidthhigh,omitempty"`
	/**
	* Normal threshold value for the bandwidth usage of the interface, in Mbps. When the bandwidth usage of the interface becomes less than or equal to the specified normal threshold after exceeding the high threshold, the Citrix ADC generates an SNMP trap message to indicate that the bandwidth usage has returned to normal.
	*/
	Bandwidthnormal int `json:"bandwidthnormal,omitempty"`
	/**
	* Link Layer Discovery Protocol (LLDP) mode for an interface. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels.
	*/
	Lldpmode string `json:"lldpmode,omitempty"`
	/**
	* LRSET port priority, expressed as an integer ranging from 1 to 1024. The highest priority is 1. The Citrix ADC limits the number of interfaces in an LRSET to 8. Within a LRSET the highest LR Priority Interface is considered as the first candidate for the Active interface, if the interface is UP.
	*/
	Lrsetpriority int `json:"lrsetpriority,omitempty"`

	//------- Read only Parameter ---------;

	Devicename string `json:"devicename,omitempty"`
	Unit string `json:"unit,omitempty"`
	Description string `json:"description,omitempty"`
	Flags string `json:"flags,omitempty"`
	Actualmtu string `json:"actualmtu,omitempty"`
	Vlan string `json:"vlan,omitempty"`
	Mac string `json:"mac,omitempty"`
	Uptime string `json:"uptime,omitempty"`
	Downtime string `json:"downtime,omitempty"`
	Actualringsize string `json:"actualringsize,omitempty"`
	Reqmedia string `json:"reqmedia,omitempty"`
	Reqspeed string `json:"reqspeed,omitempty"`
	Reqduplex string `json:"reqduplex,omitempty"`
	Reqflowcontrol string `json:"reqflowcontrol,omitempty"`
	Actmedia string `json:"actmedia,omitempty"`
	Actspeed string `json:"actspeed,omitempty"`
	Actduplex string `json:"actduplex,omitempty"`
	Actflowctl string `json:"actflowctl,omitempty"`
	Mode string `json:"mode,omitempty"`
	State string `json:"state,omitempty"`
	Autonegresult string `json:"autonegresult,omitempty"`
	Tagged string `json:"tagged,omitempty"`
	Taggedany string `json:"taggedany,omitempty"`
	Taggedautolearn string `json:"taggedautolearn,omitempty"`
	Hangdetect string `json:"hangdetect,omitempty"`
	Hangreset string `json:"hangreset,omitempty"`
	Linkstate string `json:"linkstate,omitempty"`
	Intfstate string `json:"intfstate,omitempty"`
	Rxpackets string `json:"rxpackets,omitempty"`
	Rxbytes string `json:"rxbytes,omitempty"`
	Rxerrors string `json:"rxerrors,omitempty"`
	Rxdrops string `json:"rxdrops,omitempty"`
	Txpackets string `json:"txpackets,omitempty"`
	Txbytes string `json:"txbytes,omitempty"`
	Txerrors string `json:"txerrors,omitempty"`
	Txdrops string `json:"txdrops,omitempty"`
	Indisc string `json:"indisc,omitempty"`
	Outdisc string `json:"outdisc,omitempty"`
	Fctls string `json:"fctls,omitempty"`
	Hangs string `json:"hangs,omitempty"`
	Stsstalls string `json:"stsstalls,omitempty"`
	Txstalls string `json:"txstalls,omitempty"`
	Rxstalls string `json:"rxstalls,omitempty"`
	Bdgmacmoved string `json:"bdgmacmoved,omitempty"`
	Bdgmuted string `json:"bdgmuted,omitempty"`
	Vmac string `json:"vmac,omitempty"`
	Vmac6 string `json:"vmac6,omitempty"`
	Reqthroughput string `json:"reqthroughput,omitempty"`
	Actthroughput string `json:"actthroughput,omitempty"`
	Backplane string `json:"backplane,omitempty"`
	Ifnum string `json:"ifnum,omitempty"`
	Cleartime string `json:"cleartime,omitempty"`
	Slavestate string `json:"slavestate,omitempty"`
	Slavemedia string `json:"slavemedia,omitempty"`
	Slavespeed string `json:"slavespeed,omitempty"`
	Slaveduplex string `json:"slaveduplex,omitempty"`
	Slaveflowctl string `json:"slaveflowctl,omitempty"`
	Slavetime string `json:"slavetime,omitempty"`
	Intftype string `json:"intftype,omitempty"`
	Svmcmd string `json:"svmcmd,omitempty"`
	Lacpactormode string `json:"lacpactormode,omitempty"`
	Lacpactortimeout string `json:"lacpactortimeout,omitempty"`
	Lacpactorpriority string `json:"lacpactorpriority,omitempty"`
	Lacpactorportno string `json:"lacpactorportno,omitempty"`
	Lacppartnerstate string `json:"lacppartnerstate,omitempty"`
	Lacppartnertimeout string `json:"lacppartnertimeout,omitempty"`
	Lacppartneraggregation string `json:"lacppartneraggregation,omitempty"`
	Lacppartnerinsync string `json:"lacppartnerinsync,omitempty"`
	Lacppartnercollecting string `json:"lacppartnercollecting,omitempty"`
	Lacppartnerdistributing string `json:"lacppartnerdistributing,omitempty"`
	Lacppartnerdefaulted string `json:"lacppartnerdefaulted,omitempty"`
	Lacppartnerexpired string `json:"lacppartnerexpired,omitempty"`
	Lacppartnerpriority string `json:"lacppartnerpriority,omitempty"`
	Lacppartnersystemmac string `json:"lacppartnersystemmac,omitempty"`
	Lacppartnersystempriority string `json:"lacppartnersystempriority,omitempty"`
	Lacppartnerportno string `json:"lacppartnerportno,omitempty"`
	Lacppartnerkey string `json:"lacppartnerkey,omitempty"`
	Lacpactoraggregation string `json:"lacpactoraggregation,omitempty"`
	Lacpactorinsync string `json:"lacpactorinsync,omitempty"`
	Lacpactorcollecting string `json:"lacpactorcollecting,omitempty"`
	Lacpactordistributing string `json:"lacpactordistributing,omitempty"`
	Lacpportmuxstate string `json:"lacpportmuxstate,omitempty"`
	Lacpportrxstat string `json:"lacpportrxstat,omitempty"`
	Lacpportselectstate string `json:"lacpportselectstate,omitempty"`
	Lractiveintf string `json:"lractiveintf,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
