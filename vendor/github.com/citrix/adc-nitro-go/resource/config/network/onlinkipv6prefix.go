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
* Configuration for on-link IPv6 global prefixes for Router Advertisment resource.
*/
type Onlinkipv6prefix struct {
	/**
	* Onlink prefixes for RA messages.
	*/
	Ipv6prefix string `json:"ipv6prefix,omitempty"`
	/**
	* RA Prefix onlink flag.
	*/
	Onlinkprefix string `json:"onlinkprefix,omitempty"`
	/**
	* RA Prefix Autonomus flag.
	*/
	Autonomusprefix string `json:"autonomusprefix,omitempty"`
	/**
	* Depricate the prefix.
	*/
	Depricateprefix string `json:"depricateprefix,omitempty"`
	/**
	* RA Prefix Autonomus flag.
	*/
	Decrementprefixlifetimes string `json:"decrementprefixlifetimes,omitempty"`
	/**
	* Valide life time of the prefix, in seconds.
	*/
	Prefixvalidelifetime int `json:"prefixvalidelifetime,omitempty"`
	/**
	* Preferred life time of the prefix, in seconds.
	*/
	Prefixpreferredlifetime int `json:"prefixpreferredlifetime,omitempty"`

	//------- Read only Parameter ---------;

	Prefixcurrvalidelft string `json:"prefixcurrvalidelft,omitempty"`
	Prefixcurrpreferredlft string `json:"prefixcurrpreferredlft,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
