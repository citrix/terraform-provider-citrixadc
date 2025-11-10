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
* Binding class showing the servicegroupentitymonbindings that can be bound to servicegroup.
*/
type Servicegroupservicegroupentitymonbindingsbinding struct {
	Servicegroupentname2 string `json:"servicegroupentname2,omitempty"`
	/**
	* Monitor name.
	*/
	Monitorname string `json:"monitor_name,omitempty"`
	/**
	* The running state of the monitor on this service.
	*/
	Monitorstate string `json:"monitor_state,omitempty"`
	/**
	* Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
	*/
	Passive bool `json:"passive,omitempty"`
	/**
	* Total number of probes sent to monitor this service.
	*/
	Monitortotalprobes *int `json:"monitortotalprobes,omitempty"`
	/**
	* Total number of failed probes
	*/
	Monitortotalfailedprobes *int `json:"monitortotalfailedprobes,omitempty"`
	/**
	* Total number of currently failed probes
	*/
	Monitorcurrentfailedprobes *int `json:"monitorcurrentfailedprobes,omitempty"`
	/**
	* The string form of monstatcode.
	*/
	Lastresponse string `json:"lastresponse,omitempty"`
	/**
	* Response time of this monitor in milli sec.
	*/
	Responsetime *int `json:"responsetime,omitempty"`
	/**
	* Name of the service group.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* Port number of the service. Each service must have a unique port number.
	*/
	Port *int `json:"port,omitempty"`
	Weight *int `json:"weight,omitempty"`
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