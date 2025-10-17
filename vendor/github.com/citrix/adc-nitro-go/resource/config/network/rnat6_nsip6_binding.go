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
* Binding class showing the nsip6 that can be bound to rnat6.
*/
type Rnat6nsip6binding struct {
	/**
	* Nat IP Address.
	*/
	Natip6 string `json:"natip6,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td *int `json:"td,omitempty"`
	/**
	* The owner node group in a Cluster for this rnat rule.
	*/
	Ownergroup string `json:"ownergroup,omitempty"`
	/**
	* Name of the RNAT6 rule to which to bind NAT IPs.
	*/
	Name string `json:"name,omitempty"`


}