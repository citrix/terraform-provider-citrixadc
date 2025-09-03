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

package vpn

/**
* Configuration for Clientless VPN rewrite policy resource.
*/
type Vpnclientlessaccesspolicy struct {
	/**
	* Name of the new clientless access policy.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Expression, or name of a named expression, specifying the traffic that matches the policy.
		The following requirements apply only to the Citrix ADC CLI:
		* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
		* If the expression itself includes double quotation marks, escape the quotations by using the \ character.
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Name of the profile to invoke for the clientless access.
	*/
	Profilename string `json:"profilename,omitempty"`

	//------- Read only Parameter ---------;

	Undefaction string `json:"undefaction,omitempty"`
	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Description string `json:"description,omitempty"`
	Isdefault string `json:"isdefault,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
