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
* Configuration for DNS profile resource.
*/
type Dnsprofile struct {
	/**
	* Name of the DNS profile
	*/
	Dnsprofilename string `json:"dnsprofilename,omitempty"`
	/**
	* DNS recursive resolution; if enabled, will do recursive resolution for DNS query when the profile is associated with ADNS service, CS Vserver and DNS action
	*/
	Recursiveresolution string `json:"recursiveresolution,omitempty"`
	/**
	* DNS query logging; if enabled, DNS query information such as DNS query id, DNS query flags , DNS domain name and DNS query type will be logged
	*/
	Dnsquerylogging string `json:"dnsquerylogging,omitempty"`
	/**
	* DNS answer section; if enabled, answer section in the response will be logged.
	*/
	Dnsanswerseclogging string `json:"dnsanswerseclogging,omitempty"`
	/**
	* DNS extended logging; if enabled, authority and additional section in the response will be logged.
	*/
	Dnsextendedlogging string `json:"dnsextendedlogging,omitempty"`
	/**
	* DNS error logging; if enabled, whenever error is encountered in DNS module reason for the error will be logged.
	*/
	Dnserrorlogging string `json:"dnserrorlogging,omitempty"`
	/**
	* Cache resource records in the DNS cache. Applies to resource records obtained through proxy configurations only. End resolver and forwarder configurations always cache records in the DNS cache, and you cannot disable this behavior. When you disable record caching, the appliance stops caching server responses. However, cached records are not flushed. The appliance does not serve requests from the cache until record caching is enabled again.
	*/
	Cacherecords string `json:"cacherecords,omitempty"`
	/**
	* Cache negative responses in the DNS cache. When disabled, the appliance stops caching negative responses except referral records. This applies to all configurations - proxy, end resolver, and forwarder. However, cached responses are not flushed. The appliance does not serve negative responses from the cache until this parameter is enabled again.
	*/
	Cachenegativeresponses string `json:"cachenegativeresponses,omitempty"`
	/**
	* Drop the DNS requests containing multiple queries. When enabled, DNS requests containing multiple queries will be dropped. In case of proxy configuration by default the DNS request containing multiple queries is forwarded to the backend and in case of ADNS and Resolver configuration NOCODE error response will be sent to the client.
	*/
	Dropmultiqueryrequest string `json:"dropmultiqueryrequest,omitempty"`
	/**
	* Cache DNS responses with EDNS Client Subnet(ECS) option in the DNS cache. When disabled, the appliance stops caching responses with ECS option. This is relevant to proxy configuration. Enabling/disabling support of ECS option when Citrix ADC is authoritative for a GSLB domain is supported using a knob in GSLB vserver. In all other modes, ECS option is ignored.
	*/
	Cacheecsresponses string `json:"cacheecsresponses,omitempty"`
	/**
	* Insert ECS Option on DNS query
	*/
	Insertecs string `json:"insertecs,omitempty"`
	/**
	* Replace ECS Option on DNS query
	*/
	Replaceecs string `json:"replaceecs,omitempty"`
	/**
	* The maximum ecs prefix length that will be cached
	*/
	Maxcacheableecsprefixlength int `json:"maxcacheableecsprefixlength,omitempty"`
	/**
	* The maximum ecs prefix length that will be cached for IPv6 subnets
	*/
	Maxcacheableecsprefixlength6 int `json:"maxcacheableecsprefixlength6,omitempty"`

	//------- Read only Parameter ---------;

	Referencecount string `json:"referencecount,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
