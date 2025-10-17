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
* Configuration for DNS action resource.
*/
type Dnsaction struct {
	/**
	* Name of the dns action.
	*/
	Actionname string `json:"actionname,omitempty"`
	/**
	* The type of DNS action that is being configured.
	*/
	Actiontype string `json:"actiontype,omitempty"`
	/**
	* List of IP address to be returned in case of rewrite_response actiontype. They can be of IPV4 or IPV6 type.
		In case of set command We will remove all the IP address previously present in the action and will add new once given in set dns action command.
	*/
	Ipaddress []string `json:"ipaddress,omitempty"`
	/**
	* Time to live, in seconds.
	*/
	Ttl *int `json:"ttl,omitempty"`
	/**
	* The view name that must be used for the given action.
	*/
	Viewname string `json:"viewname,omitempty"`
	/**
	* The location list in priority order used for the given action.
	*/
	Preferredloclist []string `json:"preferredloclist,omitempty"`
	/**
	* Name of the DNS profile to be associated with the transaction for which the action is chosen
	*/
	Dnsprofilename string `json:"dnsprofilename,omitempty"`

	//------- Read only Parameter ---------;

	Drop string `json:"drop,omitempty"`
	Cachebypass string `json:"cachebypass,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
