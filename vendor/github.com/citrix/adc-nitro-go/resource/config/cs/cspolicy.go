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

package cs

/**
* Configuration for content-switching policy resource.
*/
type Cspolicy struct {
	/**
	* Name for the content switching policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after a policy is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* URL string that is matched with the URL of a request. Can contain a wildcard character. Specify the string value in the following format: [[prefix] [*]] [.suffix].
	*/
	Url string `json:"url,omitempty"`
	/**
	* Expression, or name of a named expression, against which traffic is evaluated.
		The following requirements apply only to the Citrix ADC CLI:
		*  If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
		*  If the expression itself includes double quotation marks, escape the quotations by using the  character. 
		*  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* The domain name. The string value can range to 63 characters.
	*/
	Domain string `json:"domain,omitempty"`
	/**
	* Content switching action that names the target load balancing virtual server to which the traffic is switched.
	*/
	Action string `json:"action,omitempty"`
	/**
	* The log action associated with the content switching policy
	*/
	Logaction string `json:"logaction,omitempty"`
	/**
	* The new name of the content switching policy.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Vstype string `json:"vstype,omitempty"`
	Hits string `json:"hits,omitempty"`
	Bindhits string `json:"bindhits,omitempty"`
	Labelname string `json:"labelname,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Priority string `json:"priority,omitempty"`
	Activepolicy string `json:"activepolicy,omitempty"`
	Cspolicytype string `json:"cspolicytype,omitempty"`

}
