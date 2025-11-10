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
* Configuration for server record resource.
*/
type Dnssrvrec struct {
	/**
	* Domain name, which, by convention, is prefixed by the symbolic name of the desired service and the symbolic name of the desired protocol, each with an underscore (_) prepended. For example, if an SRV-aware client wants to discover a SIP service that is provided over UDP, in the domain example.com, the client performs a lookup for _sip._udp.example.com.
	*/
	Domain string `json:"domain,omitempty"`
	/**
	* Target host for the specified service.
	*/
	Target string `json:"target,omitempty"`
	/**
	* Integer specifying the priority of the target host. The lower the number, the higher the priority. If multiple target hosts have the same priority, selection is based on the Weight parameter.
	*/
	Priority *int `json:"priority,omitempty"`
	/**
	* Weight for the target host. Aids host selection when two or more hosts have the same priority. A larger number indicates greater weight.
	*/
	Weight *int `json:"weight,omitempty"`
	/**
	* Port on which the target host listens for client requests.
	*/
	Port *int `json:"port,omitempty"`
	/**
	* Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
	*/
	Ttl *int `json:"ttl,omitempty"`
	/**
	* Subnet for which the cached SRV record need to be removed.
	*/
	Ecssubnet string `json:"ecssubnet,omitempty"`
	/**
	* Type of records to display. Available settings function as follows:
		* ADNS - Display all authoritative address records.
		* PROXY - Display all proxy address records.
		* ALL - Display all address records.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid *int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Authtype string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
