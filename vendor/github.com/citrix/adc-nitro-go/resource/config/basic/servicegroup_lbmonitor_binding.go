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

package basic

/**
* Binding class showing the lbmonitor that can be bound to servicegroup.
*/
type Servicegrouplbmonitorbinding struct {
	/**
	* Monitor name.
	*/
	Monitorname string `json:"monitor_name,omitempty"`
	/**
	* weight of the monitor that is bound to servicegroup.
	*/
	Monweight *int `json:"monweight,omitempty"`
	/**
	* Monitor state.
	*/
	Monstate string `json:"monstate,omitempty"`
	/**
	* Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
	*/
	Weight *int `json:"weight,omitempty"`
	/**
	* Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
	*/
	Passive bool `json:"passive,omitempty"`
	/**
	* Name of the service group.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* Port number of the service. Each service must have a unique port number.
	*/
	Port *int `json:"port,omitempty"`
	/**
	* Unique service identifier. Used when the persistency type for the virtual server is set to Custom Server ID.
	*/
	Customserverid string `json:"customserverid,omitempty"`
	/**
	* The  identifier for the service. This is used when the persistency type is set to Custom Server ID.
	*/
	Serverid *int `json:"serverid,omitempty"`
	/**
	* Initial state of the service after binding.
	*/
	State string `json:"state,omitempty"`
	/**
	* Unique numerical identifier used by hash based load balancing methods to identify a service.
	*/
	Hashid *int `json:"hashid,omitempty"`
	/**
	* Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver
	*/
	Nameserver string `json:"nameserver,omitempty"`
	/**
	* Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors
	*/
	Dbsttl *int `json:"dbsttl,omitempty"`
	/**
	* Order number to be assigned to the servicegroup member
	*/
	Order *int `json:"order,omitempty"`


}