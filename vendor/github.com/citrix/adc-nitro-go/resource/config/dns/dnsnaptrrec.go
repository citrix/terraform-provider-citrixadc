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
* Configuration for NAPTR record resource.
*/
type Dnsnaptrrec struct {
	/**
	* Name of the domain for the NAPTR record.
	*/
	Domain string `json:"domain,omitempty"`
	/**
	* An integer specifying the order in which the NAPTR records MUST be processed in order to accurately represent the ordered list of Rules. The ordering is from lowest to highest
	*/
	Order *int `json:"order,omitempty"`
	/**
	* An integer specifying the preference of this NAPTR among NAPTR records having same order. lower the number, higher the preference.
	*/
	Preference *int `json:"preference,omitempty"`
	/**
	* flags for this NAPTR.
	*/
	Flags string `json:"flags,omitempty"`
	/**
	* Service Parameters applicable to this delegation path.
	*/
	Services string `json:"services,omitempty"`
	/**
	* The regular expression, that specifies the substitution expression for this NAPTR
	*/
	Regexp string `json:"regexp,omitempty"`
	/**
	* The replacement domain name for this NAPTR.
	*/
	Replacement string `json:"replacement,omitempty"`
	/**
	* Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
	*/
	Ttl *int `json:"ttl,omitempty"`
	/**
	* Unique, internally generated record ID. View the details of the naptr record to obtain its record ID. Records can be removed by either specifying the domain name and record id OR by specifying
		domain name and all other naptr record attributes as was supplied during the add command.
	*/
	Recordid *int `json:"recordid,omitempty"`
	/**
	* Subnet for which the cached NAPTR record need to be removed.
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
	Vservername string `json:"vservername,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
