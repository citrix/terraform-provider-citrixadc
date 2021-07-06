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
* Binding class showing the lbmonitor that can be bound to gslbservicegroup.
*/
type Gslbservicegrouplbmonitorbinding struct {
	/**
	* Monitor name.
	*/
	Monitorname string `json:"monitor_name,omitempty"`
	/**
	* weight of the monitor that is bound to GSLB servicegroup.
	*/
	Monweight int `json:"monweight,omitempty"`
	/**
	* Monitor state.
	*/
	Monstate string `json:"monstate,omitempty"`
	/**
	* Weight to assign to the servers in the service group. Specifies the capacity of the servers relative to the other servers in the load balancing configuration. The higher the weight, the higher the percentage of requests sent to the service.
	*/
	Weight int `json:"weight,omitempty"`
	/**
	* Indicates if load monitor is passive. A passive load monitor does not remove service from LB decision when threshold is breached.
	*/
	Passive bool `json:"passive,omitempty"`
	/**
	* Name of the GSLB service group.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* Port number of the GSLB service. Each service must have a unique port number.
	*/
	Port int `json:"port,omitempty"`
	/**
	* Initial state of the service after binding.
	*/
	State string `json:"state,omitempty"`
	/**
	* Unique numerical identifier used by hash based load balancing methods to identify a service.
	*/
	Hashid int `json:"hashid,omitempty"`
	/**
	* The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.
	*/
	Publicip string `json:"publicip,omitempty"`
	/**
	* The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.
	*/
	Publicport int `json:"publicport,omitempty"`
	/**
	* The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
	*/
	Siteprefix string `json:"siteprefix,omitempty"`


}