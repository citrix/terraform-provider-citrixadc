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
* Configuration for proxy record resource.
*/
type Dnsproxyrecords struct {
	/**
	* Filter the DNS records to be flushed.e.g flush dns proxyRecords -type A   will flush only the A records from the cache. 
	*/
	Type string `json:"type,omitempty"`
	/**
	* Filter the Negative DNS records i.e NXDOMAIN and NODATA entries to be flushed. e.g flush dns proxyRecords NXDOMAIN will flush only the NXDOMAIN entries from the cache
	*/
	Negrectype string `json:"negrectype,omitempty"`

}
