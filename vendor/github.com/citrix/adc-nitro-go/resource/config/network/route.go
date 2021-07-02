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
* Configuration for route resource.
*/
type Route struct {
	/**
	* IPv4 network address for which to add a route entry in the routing table of the Citrix ADC.
	*/
	Network string `json:"network,omitempty"`
	/**
	* The subnet mask associated with the network address.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* IP address of the gateway for this route. Can be either the IP address of the gateway, or can be null to specify a null interface route.
	*/
	Gateway string `json:"gateway,omitempty"`
	/**
	* VLAN as the gateway for this route.
	*/
	Vlan uint32 `json:"vlan,omitempty"`
	/**
	* Positive integer used by the routing algorithms to determine preference for using this route. The lower the cost, the higher the preference.
	*/
	Cost uint32 `json:"cost,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* Administrative distance of this route, which determines the preference of this route over other routes, with same destination, from different routing protocols. A lower value is preferred.
	*/
	Distance uint32 `json:"distance,omitempty"`
	/**
	* The cost of a route is used to compare routes of the same type. The route having the lowest cost is the most preferred route. Possible values: 0 through 65535. Default: 0.
	*/
	Cost1 uint32 `json:"cost1,omitempty"`
	/**
	* Positive integer used by the routing algorithms to determine preference for this route over others of equal cost. The lower the weight, the higher the preference.
	*/
	Weight uint32 `json:"weight,omitempty"`
	/**
	* Advertise this route.
	*/
	Advertise string `json:"advertise,omitempty"`
	/**
	* Routing protocol used for advertising this route.
	*/
	Protocol []string `json:"protocol,omitempty"`
	/**
	* Monitor this route using a monitor of type ARP or PING.
	*/
	Msr string `json:"msr,omitempty"`
	/**
	* Name of the monitor, of type ARP or PING, configured on the Citrix ADC to monitor this route.
	*/
	Monitor string `json:"monitor,omitempty"`
	/**
	* The owner node group in a Cluster for this route. If owner node group is not specified then the route is treated as Striped route.
	*/
	Ownergroup string `json:"ownergroup,omitempty"`
	/**
	* Protocol used by routes that you want to remove from the routing table of the Citrix ADC.
	*/
	Routetype string `json:"routetype,omitempty"`
	/**
	* Display a detailed view.
	*/
	Detail bool `json:"detail,omitempty"`

	//------- Read only Parameter ---------;

	Gatewayname string `json:"gatewayname,omitempty"`
	Type string `json:"type,omitempty"`
	Dynamic string `json:"dynamic,omitempty"`
	Static string `json:"Static,omitempty"`
	Permanent string `json:"permanent,omitempty"`
	Direct string `json:"direct,omitempty"`
	Nat string `json:"nat,omitempty"`
	Lbroute string `json:"lbroute,omitempty"`
	Adv string `json:"adv,omitempty"`
	Tunnel string `json:"tunnel,omitempty"`
	Data string `json:"data,omitempty"`
	Data0 string `json:"data0,omitempty"`
	Flags string `json:"flags,omitempty"`
	Routeowners string `json:"routeowners,omitempty"`
	Retain string `json:"retain,omitempty"`
	Ospf string `json:"ospf,omitempty"`
	Isis string `json:"isis,omitempty"`
	Rip string `json:"rip,omitempty"`
	Bgp string `json:"bgp,omitempty"`
	Dhcp string `json:"dhcp,omitempty"`
	Advospf string `json:"advospf,omitempty"`
	Advisis string `json:"advisis,omitempty"`
	Advrip string `json:"advrip,omitempty"`
	Advbgp string `json:"advbgp,omitempty"`
	State string `json:"state,omitempty"`
	Totalprobes string `json:"totalprobes,omitempty"`
	Totalfailedprobes string `json:"totalfailedprobes,omitempty"`
	Failedprobes string `json:"failedprobes,omitempty"`
	Monstatcode string `json:"monstatcode,omitempty"`
	Monstatparam1 string `json:"monstatparam1,omitempty"`
	Monstatparam2 string `json:"monstatparam2,omitempty"`
	Monstatparam3 string `json:"monstatparam3,omitempty"`

}
