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
* Configuration for INAT parameter resource.
*/
type Inatparam struct {
	/**
	* The prefix used for translating packets received from private IPv6 servers into IPv4 packets. This prefix has a length of 96 bits (128-32 = 96). The IPv6 servers embed the destination IP address of the IPv4 servers or hosts in the last 32 bits of the destination IP address field of the IPv6 packets. The first 96 bits of the destination IP address field are set as the IPv6 NAT prefix. IPv6 packets addressed to this prefix have to be routed to the Citrix ADC to ensure that the IPv6-IPv4 translation is done by the appliance.
	*/
	Nat46v6prefix string `json:"nat46v6prefix,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td,omitempty"`
	/**
	* Ignore TOS.
	*/
	Nat46ignoretos string `json:"nat46ignoretos,omitempty"`
	/**
	* Calculate checksum for UDP packets with zero checksum
	*/
	Nat46zerochecksum string `json:"nat46zerochecksum,omitempty"`
	/**
	* MTU setting for the IPv6 side. If the incoming IPv4 packet greater than this, either fragment or send icmp need fragmentation error.
	*/
	Nat46v6mtu int `json:"nat46v6mtu,omitempty"`
	/**
	* When disabled, translator will not insert IPv6 fragmentation header for non fragmented IPv4 packets
	*/
	Nat46fragheader string `json:"nat46fragheader,omitempty"`

}
