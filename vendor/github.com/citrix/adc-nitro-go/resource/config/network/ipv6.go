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
* Configuration for ip v6 resource.
*/
type Ipv6 struct {
	/**
	* Enable the Citrix ADC to learn about various routes from Router Advertisement (RA) and Router Solicitation (RS) messages sent by the routers.
	*/
	Ralearning string `json:"ralearning,omitempty"`
	/**
	* Enable the Citrix ADC to do Router Redirection.
	*/
	Routerredirection string `json:"routerredirection,omitempty"`
	/**
	* Base reachable time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, that the Citrix ADC assumes an adjacent device is reachable after receiving a reachability confirmation.
	*/
	Ndbasereachtime int `json:"ndbasereachtime,omitempty"`
	/**
	* Retransmission time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, between retransmitted Neighbor Solicitation (NS) messages, to an adjacent device.
	*/
	Ndretransmissiontime int `json:"ndretransmissiontime,omitempty"`
	/**
	* Prefix used for translating packets from private IPv6 servers to IPv4 packets. This prefix has a length of 96 bits (128-32 = 96). The IPv6 servers embed the destination IP address of the IPv4 servers or hosts in the last 32 bits of the destination IP address field of the IPv6 packets. The first 96 bits of the destination IP address field are set as the IPv6 NAT prefix. IPv6 packets addressed to this prefix have to be routed to the Citrix ADC to ensure that the IPv6-IPv4 translation is done by the appliance.
	*/
	Natprefix string `json:"natprefix,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td"` // Zero is a valid value
	/**
	* Enable the Citrix ADC to do Duplicate Address
		Detection (DAD) for all the Citrix ADC owned IPv6 addresses regardless of whether they are obtained through stateless auto configuration, DHCPv6, or manual configuration.
	*/
	Dodad string `json:"dodad,omitempty"`
	/**
	* IPV6 NATPREFIX used in NAT46 scenario when USIP is turned on
	*/
	Usipnatprefix string `json:"usipnatprefix,omitempty"`

	//------- Read only Parameter ---------;

	Basereachtime string `json:"basereachtime,omitempty"`
	Reachtime string `json:"reachtime,omitempty"`
	Ndreachtime string `json:"ndreachtime,omitempty"`
	Retransmissiontime string `json:"retransmissiontime,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
