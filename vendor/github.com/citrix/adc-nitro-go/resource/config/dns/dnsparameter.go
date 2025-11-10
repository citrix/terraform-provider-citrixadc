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
* Configuration for DNS parameter resource.
*/
type Dnsparameter struct {
	/**
	* Maximum number of retry attempts when no response is received for a query sent to a name server. Applies to end resolver and forwarder configurations.
	*/
	Retries *int `json:"retries,omitempty"`
	/**
	* Minimum permissible time to live (TTL) for all records cached in the DNS cache by DNS proxy, end resolver, and forwarder configurations. If the TTL of a record that is to be cached is lower than the value configured for minTTL, the TTL of the record is set to the value of minTTL before caching. When you modify this setting, the new value is applied only to those records that are cached after the modification. The TTL values of existing records are not changed.
	*/
	Minttl *int `json:"minttl"` // Zero is a valid value
	/**
	* Maximum time to live (TTL) for all records cached in the DNS cache by DNS proxy, end resolver, and forwarder configurations. If the TTL of a record that is to be cached is higher than the value configured for maxTTL, the TTL of the record is set to the value of maxTTL before caching. When you modify this setting, the new value is applied only to those records that are cached after the modification. The TTL values of existing records are not changed.
	*/
	Maxttl *int `json:"maxttl,omitempty"`
	/**
	* Cache resource records in the DNS cache. Applies to resource records obtained through proxy configurations only. End resolver and forwarder configurations always cache records in the DNS cache, and you cannot disable this behavior. When you disable record caching, the appliance stops caching server responses. However, cached records are not flushed. The appliance does not serve requests from the cache until record caching is enabled again.
	*/
	Cacherecords string `json:"cacherecords,omitempty"`
	/**
	* Type of lookup (DNS or WINS) to attempt first. If the first-priority lookup fails, the second-priority lookup is attempted. Used only by the SSL VPN feature.
	*/
	Namelookuppriority string `json:"namelookuppriority,omitempty"`
	/**
	* Function as an end resolver and recursively resolve queries for domains that are not hosted on the Citrix ADC. Also resolve queries recursively when the external name servers configured on the appliance (for a forwarder configuration) are unavailable. When external name servers are unavailable, the appliance queries a root server and resolves the request recursively, as it does for an end resolver configuration.
	*/
	Recursion string `json:"recursion,omitempty"`
	/**
	* Type of DNS queries (A, AAAA, or both) to generate during the routine functioning of certain Citrix ADC features, such as SSL VPN, cache redirection, and the integrated cache. The queries are sent to the external name servers that are configured for the forwarder function. If you specify both query types, you can also specify the order. Available settings function as follows:
		* OnlyAQuery. Send queries for IPv4 address records (A records) only.
		* OnlyAAAAQuery. Send queries for IPv6 address records (AAAA records) instead of queries for IPv4 address records (A records).
		* AThenAAAAQuery. Send a query for an A record, and then send a query for an AAAA record if the query for the A record results in a NODATA response from the name server.
		* AAAAThenAQuery. Send a query for an AAAA record, and then send a query for an A record if the query for the AAAA record results in a NODATA response from the name server.
	*/
	Resolutionorder string `json:"resolutionorder,omitempty"`
	/**
	* Enable or disable the Domain Name System Security Extensions (DNSSEC) feature on the appliance. Note: Even when the DNSSEC feature is enabled, forwarder configurations (used by internal Citrix ADC features such as SSL VPN and Cache Redirection for name resolution) do not support the DNSSEC OK (DO) bit in the EDNS0 OPT header.
	*/
	Dnssec string `json:"dnssec,omitempty"`
	/**
	* Maximum number of concurrent DNS requests to allow on a single client connection, which is identified by the <clientip:port>-<vserver ip:port> tuple. A value of 0 (zero) applies no limit to the number of concurrent DNS requests allowed on a single client connection.
	*/
	Maxpipeline *int `json:"maxpipeline,omitempty"`
	/**
	* Send a root referral if a client queries a domain name that is unrelated to the domains configured/cached on the Citrix ADC. If the setting is disabled, the appliance sends a blank response instead of a root referral. Applicable to domains for which the appliance is authoritative. Disable the parameter when the appliance is under attack from a client that is sending a flood of queries for unrelated domains.
	*/
	Dnsrootreferral string `json:"dnsrootreferral,omitempty"`
	/**
	* While doing DNS64 resolution, this parameter specifies the time to wait before sending an A query if no response is received from backend DNS server for AAAA query.
	*/
	Dns64timeout *int `json:"dns64timeout"` // Zero is a valid value
	/**
	* Maximum number of subnets that can be cached corresponding to a single domain. Subnet caching will occur for responses with EDNS Client Subnet (ECS) option. Caching of such responses can be disabled using DNS profile settings. A value of zero indicates that the number of subnets cached is limited only by existing memory constraints. The default value is zero.
	*/
	Ecsmaxsubnets *int `json:"ecsmaxsubnets"` // Zero is a valid value
	/**
	* Maximum time to live (TTL) for all negative records ( NXDONAIN and NODATA ) cached in the DNS cache by DNS proxy, end resolver, and forwarder configurations. If the TTL of a record that is to be cached is higher than the value configured for maxnegcacheTTL, the TTL of the record is set to the value of maxnegcacheTTL before caching. When you modify this setting, the new value is applied only to those records that are cached after the modification. The TTL values of existing records are not changed.
	*/
	Maxnegcachettl *int `json:"maxnegcachettl,omitempty"`
	/**
	* This parameter is applicable only in proxy mode and if this parameter is enabled  we will forward all the client requests to the backend DNS server and the response served will be cached on Citrix ADC
	*/
	Cachehitbypass string `json:"cachehitbypass,omitempty"`
	/**
	* Maximum memory, in megabytes, that can be used for dns caching per Packet Engine.
	*/
	Maxcachesize *int `json:"maxcachesize"` // Zero is a valid value
	/**
	* Maximum number of active concurrent DNS resolutions per Packet Engine
	*/
	Resolvermaxactiveresolutions *int `json:"resolvermaxactiveresolutions,omitempty"`
	/**
	* Maximum DNS-TCP connections opened for recursive resolution per Packet Engine
	*/
	Resolvermaxtcpconnections *int `json:"resolvermaxtcpconnections,omitempty"`
	/**
	* Maximum wait time in seconds for the response on DNS-TCP connection for recursive resolution per Packet Engine
	*/
	Resolvermaxtcptimeout *int `json:"resolvermaxtcptimeout,omitempty"`
	/**
	* Maximum memory, in megabytes, that can be used for caching of negative DNS responses per packet engine.
	*/
	Maxnegativecachesize *int `json:"maxnegativecachesize"` // Zero is a valid value
	/**
	* If this flag is set to YES, the existing entries in cache do not age out. On reaching the max limit the cache records are frozen
	*/
	Cachenoexpire string `json:"cachenoexpire,omitempty"`
	/**
	* Processing requests split across multiple packets
	*/
	Splitpktqueryprocessing string `json:"splitpktqueryprocessing,omitempty"`
	/**
	* Cache ECS responses with a Scope Prefix length of zero. Such a cached response will be used for all queries with this domain name and any subnet. When disabled, ECS responses with Scope Prefix length of zero will be cached, but not tied to any subnet. This option has no effect if caching of ECS responses is disabled in the corresponding DNS profile.
	*/
	Cacheecszeroprefix string `json:"cacheecszeroprefix,omitempty"`
	/**
	* Maximum UDP packet size that can be handled by Citrix ADC. This is the value advertised by Citrix ADC when responding as an authoritative server and it is also used when Citrix ADC queries other name servers as a forwarder. When acting as a proxy, requests from clients are limited by this parameter - if a request contains a size greater than this value in the OPT record, it will be replaced.
	*/
	Maxudppacketsize *int `json:"maxudppacketsize,omitempty"`
	/**
	* Flag to enable/disable DNS zones configuration transfer to remote GSLB site nodes
	*/
	Zonetransfer string `json:"zonetransfer,omitempty"`
	/**
	* Flag to enable/disable saving of rollover operations executed automatically to avoid config loss.
		Applicable only when autorollover option is enabled on a key. Note: when you enable this, full configuration will be saved
	*/
	Autosavekeyops string `json:"autosavekeyops,omitempty"`
	/**
	* Rate limit threshold for Non-Existant domain (NXDOMAIN) responses generated from Citrix ADC. Once the threshold is breached , DNS queries leading to NXDOMAIN response will be dropped. This threshold will not be applied for NXDOMAIN responses got from the backend. The threshold will be applied per packet engine and per second.
	*/
	Nxdomainratelimitthreshold *int `json:"nxdomainratelimitthreshold"` // Zero is a valid value

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nxdomainthresholdcrossed string `json:"nxdomainthresholdcrossed,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
