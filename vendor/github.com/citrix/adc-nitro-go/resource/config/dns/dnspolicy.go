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
* Configuration for DNS policy resource.
*/
type Dnspolicy struct {
	/**
	* Name for the DNS policy.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Expression against which DNS traffic is evaluated.
		Note:
		* On the command line interface, if the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.
		* If the expression itself includes double quotation marks, you must escape the quotations by using the  character.
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
		Example: CLIENT.UDP.DNS.DOMAIN.EQ("domainname")
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* The view name that must be used for the given policy.
	*/
	Viewname string `json:"viewname,omitempty"`
	/**
	* The location used for the given policy. This is deprecated attribute. Please use -prefLocList
	*/
	Preferredlocation string `json:"preferredlocation,omitempty"`
	/**
	* The location list in priority order used for the given policy.
	*/
	Preferredloclist []string `json:"preferredloclist,omitempty"`
	/**
	* The dns packet must be dropped.
	*/
	Drop string `json:"drop,omitempty"`
	/**
	* By pass dns cache for this.
	*/
	Cachebypass string `json:"cachebypass,omitempty"`
	/**
	* Name of the DNS action to perform when the rule evaluates to TRUE. The built in actions function as follows:
		* dns_default_act_Drop. Drop the DNS request.
		* dns_default_act_Cachebypass. Bypass the DNS cache and forward the request to the name server.
		You can create custom actions by using the add dns action command in the CLI or the DNS > Actions > Create DNS Action dialog box in the Citrix ADC configuration utility.
	*/
	Actionname string `json:"actionname,omitempty"`
	/**
	* Name of the messagelog action to use for requests that match this policy.
	*/
	Logaction string `json:"logaction,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Description string `json:"description,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
