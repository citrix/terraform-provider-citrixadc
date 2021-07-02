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
* Binding class showing the vserver that can be bound to gslbdomain.
*/
type Gslbdomainvserverbinding struct {
	Vservername string `json:"vservername,omitempty"`
	/**
	* The type GSLB service
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* The state of the vserver
	*/
	State string `json:"state,omitempty"`
	/**
	* The load balancing method set for the virtual server
	*/
	Lbmethod string `json:"lbmethod,omitempty"`
	/**
	* The IP type for this GSLB vserver.
	*/
	Dnsrecordtype string `json:"dnsrecordtype,omitempty"`
	/**
	* Indicates the backup method in case the primary fails
	*/
	Backuplbmethod string `json:"backuplbmethod,omitempty"`
	/**
	* Indicates if persistence is set on the gslb vserver
	*/
	Persistencetype string `json:"persistencetype,omitempty"`
	/**
	* Send clients an empty DNS response when the GSLB virtual server is DOWN.
	*/
	Edr string `json:"edr,omitempty"`
	Mir string `json:"mir,omitempty"`
	/**
	* Dynamic weight method of the vserver
	*/
	Dynamicweight string `json:"dynamicweight,omitempty"`
	/**
	* Time since last state change
	*/
	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	/**
	* Indicates if Client IP option is enabled
	*/
	Cip string `json:"cip,omitempty"`
	/**
	* Persistence id of the gslb vserver
	*/
	Persistenceid uint32 `json:"persistenceid,omitempty"`
	/**
	* Netmask
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* Number of bits to consider, in an IPv6 source IP address, for creating the hash that is required by the SOURCEIPHASH load balancing method.
	*/
	V6netmasklen uint32 `json:"v6netmasklen,omitempty"`
	/**
	* Name of the site to which the service belongs.
	*/
	Sitename string `json:"sitename,omitempty"`
	/**
	* Indicates the type of cookie persistence set
	*/
	Sitepersistence string `json:"sitepersistence,omitempty"`
	/**
	* The site prefix string.
	*/
	Siteprefix string `json:"siteprefix,omitempty"`
	/**
	* The string that is sent to the service. Applicable to HTTP ,HTTP-ECV and RTSP monitor types.
	*/
	Customheaders string `json:"customheaders,omitempty"`
	/**
	* The optional IPv4 network mask applied to IPv4 addresses to establish source IP address based persistence.
	*/
	Persistmask string `json:"persistmask,omitempty"`
	/**
	* Number of bits to consider in an IPv6 source IP address when creating source IP address based persistence sessions.
	*/
	V6persistmasklen uint32 `json:"v6persistmasklen,omitempty"`
	/**
	* Name of the Domain
	*/
	Name string `json:"name,omitempty"`


}