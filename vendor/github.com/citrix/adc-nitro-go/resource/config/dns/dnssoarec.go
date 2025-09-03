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
* Configuration for SOA record resource.
*/
type Dnssoarec struct {
	/**
	* Domain name for which to add the SOA record.
	*/
	Domain string `json:"domain,omitempty"`
	/**
	* Domain name of the name server that responds authoritatively for the domain.
	*/
	Originserver string `json:"originserver,omitempty"`
	/**
	* Email address of the contact to whom domain issues can be addressed. In the email address, replace the @ sign with a period (.). For example, enter domainadmin.example.com instead of domainadmin@example.com.
	*/
	Contact string `json:"contact,omitempty"`
	/**
	* The secondary server uses this parameter to determine whether it requires a zone transfer from the primary server.
	*/
	Serial int `json:"serial,omitempty"`
	/**
	* Time, in seconds, for which a secondary server must wait between successive checks on the value of the serial number.
	*/
	Refresh int `json:"refresh,omitempty"`
	/**
	* Time, in seconds, between retries if a secondary server's attempt to contact the primary server for a zone refresh fails.
	*/
	Retry int `json:"retry,omitempty"`
	/**
	* Time, in seconds, after which the zone data on a secondary name server can no longer be considered authoritative because all refresh and retry attempts made during the period have failed. After the expiry period, the secondary server stops serving the zone. Typically one week. Not used by the primary server.
	*/
	Expire int `json:"expire,omitempty"`
	/**
	* Default time to live (TTL) for all records in the zone. Can be overridden for individual records.
	*/
	Minimum int `json:"minimum,omitempty"`
	/**
	* Time to Live (TTL), in seconds, for the record. TTL is the time for which the record must be cached by DNS proxies. The specified TTL is applied to all the resource records that are of the same record type and belong to the specified domain name. For example, if you add an address record, with a TTL of 36000, to the domain name example.com, the TTLs of all the address records of example.com are changed to 36000. If the TTL is not specified, the Citrix ADC uses either the DNS zone's minimum TTL or, if the SOA record is not available on the appliance, the default value of 3600.
	*/
	Ttl int `json:"ttl,omitempty"`
	/**
	* Subnet for which the cached SOA record need to be removed.
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
