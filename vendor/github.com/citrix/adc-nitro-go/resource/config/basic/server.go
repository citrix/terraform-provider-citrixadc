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
* Configuration for server resource.
*/
type Server struct {
	/**
	* Name for the server. 
		Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
		Can be changed after the name is created.
	*/
	Name string `json:"name,omitempty"`
	/**
	* IPv4 or IPv6 address of the server. If you create an IP address based server, you can specify the name of the server, instead of its IP address, when creating a service. Note: If you do not create a server entry, the server IP address that you enter when you create a service becomes the name of the server.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Domain name of the server. For a domain based configuration, you must create the server first.
	*/
	Domain string `json:"domain,omitempty"`
	/**
	* IP address used to transform the server's DNS-resolved IP address.
	*/
	Translationip string `json:"translationip,omitempty"`
	/**
	* The netmask of the translation ip
	*/
	Translationmask string `json:"translationmask,omitempty"`
	/**
	* Time, in seconds, for which the Citrix ADC must wait, after DNS resolution fails, before sending the next DNS query to resolve the domain name.
	*/
	Domainresolveretry int32 `json:"domainresolveretry,omitempty"`
	/**
	* Initial state of the server.
	*/
	State string `json:"state,omitempty"`
	/**
	* Support IPv6 addressing mode. If you configure a server with the IPv6 addressing mode, you cannot use the server in the IPv4 addressing mode.
	*/
	Ipv6address string `json:"ipv6address,omitempty"`
	/**
	* Any information about the server.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* Specify the type of DNS resolution to be done on the configured domain to get the backend services. Valid query types are A, AAAA and SRV with A being the default querytype. The type of DNS resolution done on the domains in SRV records is inherited from ipv6 argument.
	*/
	Querytype string `json:"querytype,omitempty"`
	/**
	* Immediately send a DNS query to resolve the server's domain name.
	*/
	Domainresolvenow bool `json:"domainresolvenow,omitempty"`
	/**
	* Time, in seconds, after which all the services configured on the server are disabled.
	*/
	Delay uint64 `json:"delay,omitempty"`
	/**
	* Shut down gracefully, without accepting any new connections, and disabling each service when all of its connections are closed.
	*/
	Graceful string `json:"graceful,omitempty"`
	/**
	* Display names of the servers that have been created for internal use.
	*/
	Internal bool `json:"Internal,omitempty"`
	/**
	* New name for the server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Statechangetimesec string `json:"statechangetimesec,omitempty"`
	Tickssincelaststatechange string `json:"tickssincelaststatechange,omitempty"`
	Autoscale string `json:"autoscale,omitempty"`
	Usip string `json:"usip,omitempty"`
	Cka string `json:"cka,omitempty"`
	Tcpb string `json:"tcpb,omitempty"`
	Cmp string `json:"cmp,omitempty"`
	Cacheable string `json:"cacheable,omitempty"`
	Sc string `json:"sc,omitempty"`
	Sp string `json:"sp,omitempty"`

}
