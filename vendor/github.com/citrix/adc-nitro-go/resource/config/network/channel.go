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
* Configuration for channel resource.
*/
type Channel struct {
	/**
	* ID for the LA channel or cluster LA channel or LR channel to be created. Specify an LA channel in LA/x notation, where x can range from 1 to 8 or cluster LA channel in CLA/x notation or Link redundant channel in LR/x notation, where x can range from 1 to 4. Cannot be changed after the LA channel is created.
	*/
	Id string `json:"id,omitempty"`
	/**
	* Interfaces to be bound to the LA channel of a Citrix ADC or to the LA channel of a cluster configuration.
		For an LA channel of a Citrix ADC, specify an interface in C/U notation (for example, 1/3). 
		For an LA channel of a cluster configuration, specify an interface in N/C/U notation (for example, 2/1/3).
		where C can take one of the following values:
		* 0 - Indicates a management interface.
		* 1 - Indicates a 1 Gbps port.
		* 10 - Indicates a 10 Gbps port.
		U is a unique integer for representing an interface in a particular port group.
		N is the ID of the node to which an interface belongs in a cluster configuration.
		Use spaces to separate multiple entries.
	*/
	Ifnum []string `json:"ifnum,omitempty"`
	/**
	* Enable or disable the LA channel.
	*/
	State string `json:"state,omitempty"`
	/**
	* The initital mode for the LA channel.
	*/
	Mode string `json:"mode,omitempty"`
	/**
	* The 'connection' distribution mode for the LA channel.
	*/
	Conndistr string `json:"conndistr,omitempty"`
	/**
	* The  'MAC' distribution mode for the LA channel.
	*/
	Macdistr string `json:"macdistr,omitempty"`
	/**
	* Specifies a MAC address for the LA channels configured in Citrix ADC virtual appliances (VPX). This MAC address is persistent after each reboot. 
		If you don't specify this parameter, a MAC address is generated randomly for each LA channel. These MAC addresses change after each reboot.
	*/
	Lamac string `json:"lamac,omitempty"`
	/**
	* Ethernet speed of the channel, in Mbps. If the speed of any bound interface is greater than or equal to the value set for this parameter, the state of the interface is UP. Otherwise, the state is INACTIVE. Bound Interfaces whose state is INACTIVE do not process any traffic.
	*/
	Speed string `json:"speed,omitempty"`
	/**
	* Specifies the flow control type for this LA channel to manage the flow of frames. Flow control is a function as mentioned in clause 31 of the IEEE 802.3 standard. Flow control allows congested ports to pause traffic from the peer device. Flow control is achieved by sending PAUSE frames.
	*/
	Flowctl string `json:"flowctl,omitempty"`
	/**
	* In a High Availability (HA) configuration, monitor the LA channel for failure events. Failure of any LA channel that has HA MON enabled triggers HA failover.
	*/
	Hamonitor string `json:"hamonitor,omitempty"`
	/**
	* In a High Availability (HA) configuration, configure the LA channel for sending heartbeats. LA channel that has HA Heartbeat disabled should not send the heartbeats.
	*/
	Haheartbeat string `json:"haheartbeat,omitempty"`
	/**
	* Adds a four-byte 802.1q tag to every packet sent on this channel.  The ON setting applies tags for all VLANs that are bound to this channel. OFF applies the tag for all VLANs other than the native VLAN.
	*/
	Tagall string `json:"tagall,omitempty"`
	/**
	* This is deprecated by tagall
	*/
	Trunk string `json:"trunk,omitempty"`
	/**
	* Alias name for the LA channel. Used only to enhance readability. To perform any operations, you have to specify the LA channel ID.
	*/
	Ifalias string `json:"ifalias,omitempty"`
	/**
	* Low threshold value for the throughput of the LA channel, in Mbps. In an high availability (HA) configuration, failover is triggered when the LA channel has HA MON enabled and the throughput is below the specified threshold.
	*/
	Throughput int `json:"throughput,omitempty"`
	/**
	* High threshold value for the bandwidth usage of the LA channel, in Mbps. The Citrix ADC generates an SNMP trap message when the bandwidth usage of the LA channel is greater than or equal to the specified high threshold value.
	*/
	Bandwidthhigh int `json:"bandwidthhigh,omitempty"`
	/**
	* Normal threshold value for the bandwidth usage of the LA channel, in Mbps. When the bandwidth usage of the LA channel returns to less than or equal to the specified normal threshold after exceeding the high threshold, the Citrix ADC generates an SNMP trap message to indicate that the bandwidth usage has returned to normal.
	*/
	Bandwidthnormal int `json:"bandwidthnormal,omitempty"`
	/**
	* The Maximum Transmission Unit (MTU) is the largest packet size, measured in bytes excluding 14 bytes ethernet header and 4 bytes CRC, that can be transmitted and received by an interface. The default value of MTU is 1500 on all the interface of Citrix ADC, some Cloud Platforms will restrict Citrix ADC to use the lesser default value. Any MTU value more than 1500 is called Jumbo MTU and will make the interface as jumbo enabled. The Maximum Jumbo MTU in Citrix ADC is 9216, however, some Virtualized / Cloud Platforms will have lesser Maximum Jumbo MTU Value (9000). In the case of Cluster, the Backplane interface requires an MTU value of 78 bytes more than the Max MTU configured on any other Data-Plane Interface. When the Data plane interfaces are all at default 1500 MTU, Cluster Back Plane will be automatically set to 1578 (1500 + 78) MTU. If a Backplane interface is reset to Data Plane Interface, then the 1578 MTU will be automatically reset to the default MTU of 1500(or whatever lesser default value). If any data plane interface of a Cluster is configured with a Jumbo MTU ( > 1500), then all backplane interfaces require to be configured with a minimum MTU of 'Highest Data Plane MTU in the Cluster + 78'. That makes the maximum Jumbo MTU for any Data-Plane Interface in a Cluster System to be '9138 (9216 - 78)., where 9216 is the maximum Jumbo MTU. On certain Virtualized / Cloud Platforms, the maximum  possible MTU is restricted to a lesser value, Similar calculation can be applied, Maximum Data Plane MTU in Cluster = (Maximum possible MTU - 78). 
	*/
	Mtu int `json:"mtu,omitempty"`
	/**
	* Specifies the minimum throughput threshold (in Mbps) to be met by the active subchannel. Setting this parameter automatically divides an LACP channel into logical subchannels, with one subchannel active and the others in standby mode.  When the maximum supported throughput of the active channel falls below the lrMinThroughput value, link failover occurs and a standby subchannel becomes active.
	*/
	Lrminthroughput int `json:"lrminthroughput,omitempty"`
	/**
	* Link Redundancy for Cluster LAG.
	*/
	Linkredundancy string `json:"linkredundancy,omitempty"`

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
	Reqmedia string `json:"reqmedia,omitempty"`
	Reqspeed string `json:"reqspeed,omitempty"`
	Reqduplex string `json:"reqduplex,omitempty"`
	Reqflowcontrol string `json:"reqflowcontrol,omitempty"`
	Media string `json:"media,omitempty"`
	Actspeed string `json:"actspeed,omitempty"`
	Duplex string `json:"duplex,omitempty"`
	Actflowctl string `json:"actflowctl,omitempty"`
	Lamode string `json:"lamode,omitempty"`
	Autoneg string `json:"autoneg,omitempty"`
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
	Bdgmuted string `json:"bdgmuted,omitempty"`
	Vmac string `json:"vmac,omitempty"`
	Vmac6 string `json:"vmac6,omitempty"`
	Reqthroughput string `json:"reqthroughput,omitempty"`
	Actthroughput string `json:"actthroughput,omitempty"`
	Backplane string `json:"backplane,omitempty"`
	Cleartime string `json:"cleartime,omitempty"`
	Lacpmode string `json:"lacpmode,omitempty"`
	Lacptimeout string `json:"lacptimeout,omitempty"`
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
	Lldpmode string `json:"lldpmode,omitempty"`

}
