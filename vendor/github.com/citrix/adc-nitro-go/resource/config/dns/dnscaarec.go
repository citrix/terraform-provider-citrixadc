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
* Configuration for CAA record resource.
*/
type Dnscaarec struct {
	/**
	* Domain name of the CAA record.
	*/
	Domain string `json:"domain,omitempty"`
	/**
	* Value associated with the chosen property tag in the CAA resource record. Enclose the string in single or double quotation marks.
	*/
	Valuestring string `json:"valuestring,omitempty"`
	/**
	* String that represents the identifier of the property represented by the CAA record. The RFC currently defines three available tags - issue, issuwild and iodef.
	*/
	Tag string `json:"tag,omitempty"`
	/**
	* Flag associated with the CAA record.
	*/
	Flag string `json:"flag,omitempty"`
	/**
	* Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
	*/
	Ttl int `json:"ttl,omitempty"`
	/**
	* Unique, internally generated record ID. View the details of the CAA record to obtain its record ID. Records can be removedby either specifying the domain name and record id OR by specifying domain name and all other CAA record attributes as was supplied during the add command.
	*/
	Recordid int `json:"recordid,omitempty"`
	/**
	* Subnet for which the cached CAA record need to be removed.
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
	Nodeid int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Authtype string `json:"authtype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
