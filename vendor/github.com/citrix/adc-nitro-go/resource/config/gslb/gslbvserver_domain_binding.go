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
* Binding class showing the domain that can be bound to gslbvserver.
*/
type Gslbvserverdomainbinding struct {
	/**
	* Domain name for which to change the time to live (TTL) and/or backup service IP address.
	*/
	Domainname string `json:"domainname,omitempty"`
	/**
	* Time to live (TTL) for the domain.
	*/
	Ttl int `json:"ttl,omitempty"`
	/**
	* The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.
	*/
	Backupip string `json:"backupip,omitempty"`
	/**
	* The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.
	*/
	Cookiedomain string `json:"cookie_domain,omitempty"`
	/**
	* Timeout, in minutes, for the GSLB site cookie.
	*/
	Cookietimeout int `json:"cookietimeout,omitempty"`
	/**
	* TTL, in seconds, for all internally created site domains (created when a site prefix is configured on a GSLB service) that are associated with this virtual server.
	*/
	Sitedomainttl int `json:"sitedomainttl,omitempty"`
	/**
	* Name of the virtual server on which to perform the binding operation.
	*/
	Name string `json:"name,omitempty"`
	/**
	* The IP address of the backup service for the specified domain name. Used when all the services bound to the domain are down, or when the backup chain of virtual servers is down.
	*/
	Backupipflag bool `json:"backupipflag,omitempty"`
	/**
	* The cookie domain for the GSLB site. Used when inserting the GSLB site cookie in the HTTP response.
	*/
	Cookiedomainflag bool `json:"cookie_domainflag,omitempty"`
	/**
	* Order number to be assigned to the service when it is bound to the lb vserver.
	*/
	Order int `json:"order,omitempty"`


}