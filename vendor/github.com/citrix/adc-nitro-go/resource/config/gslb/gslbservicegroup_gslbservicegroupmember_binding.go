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

package gslb

/**
* Binding class showing the gslbservicegroupmember that can be bound to gslbservicegroup.
*/
type Gslbservicegroupgslbservicegroupmemberbinding struct {
	/**
	* IP Address.
	*/
	Ip string `json:"ip,omitempty"`
	/**
	* Server port number.
	*/
	Port int `json:"port,omitempty"`
	/**
	* The state of the GSLB service
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
	* Initial state of the GSLB service group.
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
	* The time allowed (in seconds) for a graceful shutdown. During this period, new connections or requests will continue to be sent to this service for clients who already have a persistent session on the system. Connections or requests from fresh or new clients who do not yet have a persistence sessions on the system will not be sent to the service. Instead, they will be load balanced among other available services. After the delay time expires, no new requests or connections will be sent to the service.
	*/
	Delay int `json:"delay,omitempty"`
	/**
	* The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.
	*/
	Publicip string `json:"publicip,omitempty"`
	/**
	* The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.
	*/
	Publicport int `json:"publicport,omitempty"`
	/**
	* Indicates if gslb svc has reached threshold
	*/
	Gslbthreshold int `json:"gslbthreshold,omitempty"`
	Threshold string `json:"threshold,omitempty"`
	/**
	* Prefered location.
	*/
	Preferredlocation string `json:"preferredlocation,omitempty"`
	/**
	* The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
	*/
	Siteprefix string `json:"siteprefix,omitempty"`
	/**
	* Order number to be assigned to the gslb servicegroup member
	*/
	Order int `json:"order,omitempty"`
	/**
	* Order number in string form assigned to the gslb servicegroup member
	*/
	Orderstr string `json:"orderstr,omitempty"`
	/**
	* Delay before moving to TROFS
	*/
	Trofsdelay int `json:"trofsdelay,omitempty"`
	/**
	* Name of the GSLB service group.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`


}