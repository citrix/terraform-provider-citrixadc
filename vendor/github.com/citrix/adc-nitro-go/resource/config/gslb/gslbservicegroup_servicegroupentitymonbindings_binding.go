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
* Binding class showing the servicegroupentitymonbindings that can be bound to gslbservicegroup.
*/
type Gslbservicegroupservicegroupentitymonbindingsbinding struct {
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
	Monitortotalprobes uint32 `json:"monitortotalprobes,omitempty"`
	/**
	* Total number of failed probes
	*/
	Monitortotalfailedprobes uint32 `json:"monitortotalfailedprobes,omitempty"`
	/**
	* Total number of currently failed probes
	*/
	Monitorcurrentfailedprobes uint32 `json:"monitorcurrentfailedprobes,omitempty"`
	/**
	* The string form of monstatcode.
	*/
	Lastresponse string `json:"lastresponse,omitempty"`
	/**
	* Name of the GSLB service group.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* Port number of the GSLB service. Each service must have a unique port number.
	*/
	Port int32 `json:"port,omitempty"`
	/**
	* When used along with monitor name, it specifies the weight of the monitor binding. When used along with servername & port pair, it specifies the weight of this GSLB service .
	*/
	Weight uint32 `json:"weight,omitempty"`
	/**
	* Initial state of the service after binding.
	*/
	State string `json:"state,omitempty"`
	/**
	* Unique numerical identifier used by hash based load balancing methods to identify a service.
	*/
	Hashid uint32 `json:"hashid,omitempty"`
	/**
	* The public IP address that a NAT device translates to the GSLB service's private IP address. Optional.
	*/
	Publicip string `json:"publicip,omitempty"`
	/**
	* The public port associated with the GSLB service's public IP address. The port is mapped to the service's private port number. Applicable to the local GSLB service. Optional.
	*/
	Publicport int32 `json:"publicport,omitempty"`
	/**
	* The site's prefix string. When the GSLB service group is bound to a GSLB virtual server, a GSLB site domain is generated internally for each bound serviceitem-domain pair by concatenating the site prefix of the service item and the name of the domain. If the special string NONE is specified, the site-prefix string is unset. When implementing HTTP redirect site persistence, the Citrix ADC redirects GSLB requests to GSLB services by using their site domains.
	*/
	Siteprefix string `json:"siteprefix,omitempty"`


}