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

package network

/**
* Configuration for Layer 3 related parameter resource.
*/
type L3param struct {
	/**
	* Perform NAT if only the source is in the private network
	*/
	Srcnat string `json:"srcnat,omitempty"`
	/**
	* NS generated ICMP pkts per 10ms rate threshold
	*/
	Icmpgenratethreshold uint32 `json:"icmpgenratethreshold,omitempty"`
	/**
	* USNIP/USIP settings override RNAT settings for configured
		service/virtual server traffic.. 
	*/
	Overridernat string `json:"overridernat,omitempty"`
	/**
	* Enable dropping the IP DF flag.
	*/
	Dropdfflag string `json:"dropdfflag,omitempty"`
	/**
	* Enable round robin usage of mapped IPs.
	*/
	Miproundrobin string `json:"miproundrobin,omitempty"`
	/**
	* Enable external loopback.
	*/
	Externalloopback string `json:"externalloopback,omitempty"`
	/**
	* Enable/Disable learning PMTU of IP tunnel when ICMP error does not contain connection information.
	*/
	Tnlpmtuwoconn string `json:"tnlpmtuwoconn,omitempty"`
	/**
	* Enable detection of stray server side pkts in USIP mode.
	*/
	Usipserverstraypkt string `json:"usipserverstraypkt,omitempty"`
	/**
	* Enable forwarding of ICMP fragments.
	*/
	Forwardicmpfragments string `json:"forwardicmpfragments,omitempty"`
	/**
	* Enable dropping of IP fragments.
	*/
	Dropipfragments string `json:"dropipfragments,omitempty"`
	/**
	* Parameter to tune acl logging time
	*/
	Acllogtime uint32 `json:"acllogtime,omitempty"`
	/**
	* Do not apply ACLs for internal ports
	*/
	Implicitaclallow string `json:"implicitaclallow,omitempty"`
	/**
	* Enable/Disable Dynamic routing on partition. This configuration is not applicable to default partition
	*/
	Dynamicrouting string `json:"dynamicrouting,omitempty"`
	/**
	* Enable/Disable IPv6 Dynamic routing
	*/
	Ipv6dynamicrouting string `json:"ipv6dynamicrouting,omitempty"`
	/**
	* Enable/Disable IPv4 Class E address clients
	*/
	Allowclasseipv4 string `json:"allowclasseipv4,omitempty"`

}
