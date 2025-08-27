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
* Configuration for dns64 policy resource.
*/
type Dnspolicy64 struct {
	/**
	* Name for the DNS64 policy.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Expression against which DNS traffic is evaluated.
		Note:
		* On the command line interface, if the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.
		* If the expression itself includes double quotation marks, you must escape the quotations by using the  character.
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
		Example: CLIENT.IP.SRC.IN_SUBENT(23.34.0.0/16)
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Name of the DNS64 action to perform when the rule evaluates to TRUE. The built in actions function as follows:
		* A default dns64 action with prefix <default prefix> and mapped and exclude are any
		You can create custom actions by using the add dns action command in the CLI or the DNS64 > Actions > Create DNS64 Action dialog box in the Citrix ADC configuration utility.
	*/
	Action string `json:"action,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Description string `json:"description,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
