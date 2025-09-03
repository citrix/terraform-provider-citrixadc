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
* Binding class showing the servicegroupmember that can be bound to servicegroup.
*/
type Servicegroupservicegroupmemberbinding struct {
	/**
	* IP Address.
	*/
	Ip string `json:"ip,omitempty"`
	/**
	* Server port number.
	*/
	Port int `json:"port,omitempty"`
	/**
	* The state of the service
	*/
	Svrstate string `json:"svrstate,omitempty"`
	/**
	* Time when last state change occurred. Seconds part.
	*/
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	/**
	* Time in 10 millisecond ticks since the last state change.
	*/
	Tickssincelaststatechange int `json:"tickssincelaststatechange,omitempty"`
	/**
	* Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
	*/
	Weight int `json:"weight,omitempty"`
	/**
	* Name of the server to which to bind the service group.
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* The identifier for this IP:Port pair. Used when the persistency type is set to Custom Server ID.
	*/
	Customserverid string `json:"customserverid,omitempty"`
	/**
	* The  identifier for the service. This is used when the persistency type is set to Custom Server ID.
	*/
	Serverid int `json:"serverid,omitempty"`
	/**
	* Initial state of the service group.
	*/
	State string `json:"state,omitempty"`
	/**
	* The hash identifier for the service. This must be unique for each service. This parameter is used by hash based load balancing methods.
	*/
	Hashid int `json:"hashid,omitempty"`
	/**
	* Wait for all existing connections to the service to terminate before shutting down the service.
	*/
	Graceful string `json:"graceful,omitempty"`
	/**
	* Time, in seconds, allocated for a shutdown of the services in the service group. During this period, new requests are sent to the service only for clients who already have persistent sessions on the appliance. Requests from new clients are load balanced among other available services. After the delay time expires, no requests are sent to the service, and the service is marked as unavailable (OUT OF SERVICE).
	*/
	Delay int `json:"delay,omitempty"`
	/**
	* Specify the nameserver to which the query for bound domain needs to be sent. If not specified, use the global nameserver
	*/
	Nameserver string `json:"nameserver,omitempty"`
	/**
	* Specify the TTL for DNS record for domain based service.The default value of ttl is 0 which indicates to use the TTL received in DNS response for monitors
	*/
	Dbsttl int `json:"dbsttl,omitempty"`
	/**
	* This gives the priority of the FQDN service items for SRV server binding
	*/
	Svcitmpriority int `json:"svcitmpriority,omitempty"`
	/**
	* Specify reason if service group member in TROFS
	*/
	Trofsreason string `json:"trofsreason,omitempty"`
	/**
	* Order number to be assigned to the servicegroup member
	*/
	Order int `json:"order,omitempty"`
	/**
	* Order number in string form to be assigned to the servicegroup member
	*/
	Orderstr string `json:"orderstr,omitempty"`
	/**
	* Delay before moving to TROFS
	*/
	Trofsdelay int `json:"trofsdelay,omitempty"`
	/**
	* Name of the service group.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`


}