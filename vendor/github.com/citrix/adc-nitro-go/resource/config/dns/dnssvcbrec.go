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

package dns

/**
* Configuration for SVCB/HTTPS service binding record resource.
*/
type Dnssvcbrec struct {
	/**
	* Domain name for the SVCB/HTTPS record.
	*/
	Domain string `json:"domain,omitempty"`
	/**
	* Service priority (0 for AliasMode, >0 for ServiceMode).
	*/
	Priority *int `json:"priority,omitempty"`
	/**
	* Target domain name.
	*/
	Targetname string `json:"targetname,omitempty"`
	/**
	* Service type: SVCB or HTTPS.
	*/
	Svcbtype string `json:"svcbtype,omitempty"`
	/**
	* Comma-separated list of ALPN protocol identifiers.
	*/
	Alpn string `json:"alpn,omitempty"`
	/**
	* Port number for the service.
	*/
	Port *int `json:"port,omitempty"`
	/**
	* Comma-separated list of IPv4 hint addresses.
	*/
	Ipv4hint string `json:"ipv4hint,omitempty"`
	/**
	* Comma-separated list of IPv6 hint addresses.
	*/
	Ipv6hint string `json:"ipv6hint,omitempty"`
	/**
	* Base64-encoded ECH configuration.
	*/
	Encryptedclienthello string `json:"encryptedclienthello,omitempty"`
	/**
	* Indicates no default ALPN protocols.
	*/
	Nodefaultalpn bool `json:"nodefaultalpn,omitempty"`
	/**
	* Comma-separated list of mandatory SvcParam keys.
	*/
	Mandatory string `json:"mandatory,omitempty"`
	/**
	* Time to Live (TTL) in seconds.
	*/
	Ttl *int `json:"ttl,omitempty"`
	/**
	* Type of records: ADNS, PROXY, or ALL.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid *int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Ecssubnet string `json:"ecssubnet,omitempty"`
	Authtype string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
