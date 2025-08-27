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
* Configuration for DNS zone resource.
*/
type Dnszone struct {
	/**
	* Name of the zone to create.
	*/
	Zonename string `json:"zonename,omitempty"`
	/**
	* Deploy the zone in proxy mode. Enable in the following scenarios:
		* The load balanced DNS servers are authoritative for the zone and all resource records that are part of the zone.
		* The load balanced DNS servers are authoritative for the zone, but the Citrix ADC owns a subset of the resource records that belong to the zone (partial zone ownership configuration). Typically seen in global server load balancing (GSLB) configurations, in which the appliance responds authoritatively to queries for GSLB domain names but forwards queries for other domain names in the zone to the load balanced servers.
		In either scenario, do not create the zone's Start of Authority (SOA) and name server (NS) resource records on the appliance.
		Disable if the appliance is authoritative for the zone, but make sure that you have created the SOA and NS records on the appliance before you create the zone.
	*/
	Proxymode string `json:"proxymode,omitempty"`
	/**
	* Enable dnssec offload for this zone.
	*/
	Dnssecoffload string `json:"dnssecoffload,omitempty"`
	/**
	* Enable nsec generation for dnssec offload.
	*/
	Nsec string `json:"nsec,omitempty"`
	/**
	* Name of the public/private DNS key pair with which to sign the zone. You can sign a zone with up to four keys.
	*/
	Keyname []string `json:"keyname,omitempty"`
	/**
	* Type of zone to display. Mutually exclusive with the DNS Zone (zoneName) parameter. Available settings function as follows:
		* ADNS - Display all the zones for which the Citrix ADC is authoritative.
		* PROXY - Display all the zones for which the Citrix ADC is functioning as a proxy server.
		* ALL - Display all the zones configured on the appliance.
	*/
	Type string `json:"type,omitempty"`

	//------- Read only Parameter ---------;

	Flags string `json:"flags,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
