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
* Configuration for route 6 resource.
*/
type Route6 struct {
	/**
	* IPv6 network address for which to add a route entry to the routing table of the Citrix ADC.
	*/
	Network string `json:"network,omitempty"`
	/**
	* The gateway for this route. The value for this parameter is either an IPv6 address or null.
	*/
	Gateway string `json:"gateway,omitempty"`
	/**
	* Integer value that uniquely identifies a VLAN through which the Citrix ADC forwards the packets for this route.
	*/
	Vlan int `json:"vlan,omitempty"`
	/**
	* Integer value that uniquely identifies a VXLAN through which the Citrix ADC forwards the packets for this route.
	*/
	Vxlan int `json:"vxlan,omitempty"`
	/**
	* Positive integer used by the routing algorithms to determine preference for this route over others of equal cost. The lower the weight, the higher the preference.
	*/
	Weight int `json:"weight,omitempty"`
	/**
	* Administrative distance of this route from the appliance.
	*/
	Distance int `json:"distance,omitempty"`
	/**
	* Positive integer used by the routing algorithms to determine preference for this route. The lower the cost, the higher the preference.
	*/
	Cost int `json:"cost,omitempty"`
	/**
	* Advertise this route.
	*/
	Advertise string `json:"advertise,omitempty"`
	/**
	* Monitor this route with a monitor of type ND6 or PING.
	*/
	Msr string `json:"msr,omitempty"`
	/**
	* Name of the monitor, of type ND6 or PING, configured on the Citrix ADC to monitor this route.
	*/
	Monitor string `json:"monitor,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td,omitempty"`
	/**
	* The owner node group in a Cluster for this route6. If owner node group is not specified then the route is treated as Striped route.
	*/
	Ownergroup string `json:"ownergroup,omitempty"`
	/**
	* Route in management plane.
	*/
	Mgmt bool `json:"mgmt,omitempty"`
	/**
	* Type of IPv6 routes to remove from the routing table of the Citrix ADC.
	*/
	Routetype string `json:"routetype,omitempty"`
	/**
	* To get a detailed view.
	*/
	Detail bool `json:"detail,omitempty"`

	//------- Read only Parameter ---------;

	Gatewayname string `json:"gatewayname,omitempty"`
	Type string `json:"type,omitempty"`
	Dynamic string `json:"dynamic,omitempty"`
	Data string `json:"data,omitempty"`
	Flags string `json:"flags,omitempty"`
	State string `json:"state,omitempty"`
	Totalprobes string `json:"totalprobes,omitempty"`
	Totalfailedprobes string `json:"totalfailedprobes,omitempty"`
	Failedprobes string `json:"failedprobes,omitempty"`
	Monstatcode string `json:"monstatcode,omitempty"`
	Monstatparam1 string `json:"monstatparam1,omitempty"`
	Monstatparam2 string `json:"monstatparam2,omitempty"`
	Monstatparam3 string `json:"monstatparam3,omitempty"`
	Data1 string `json:"data1,omitempty"`
	Routeowners string `json:"routeowners,omitempty"`
	Retain string `json:"retain,omitempty"`
	Static string `json:"Static,omitempty"`
	Permanent string `json:"permanent,omitempty"`
	Connected string `json:"connected,omitempty"`
	Ospfv3 string `json:"ospfv3,omitempty"`
	Isis string `json:"isis,omitempty"`
	Active string `json:"active,omitempty"`
	Bgp string `json:"bgp,omitempty"`
	Rip string `json:"rip,omitempty"`
	Raroute string `json:"raroute,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
