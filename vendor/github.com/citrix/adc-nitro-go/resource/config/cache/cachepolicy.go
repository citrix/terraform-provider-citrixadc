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

package cache

/**
* Configuration for Integrated Cache policy resource.
*/
type Cachepolicy struct {
	/**
	* Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the policy is created.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Expression against which the traffic is evaluated.
		The following requirements apply only to the Citrix ADC CLI:
		* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
		* If the expression itself includes double quotation marks, escape the quotations by using the \ character.
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Action to apply to content that matches the policy.
		* CACHE or MAY_CACHE action - positive cachability policy
		* NOCACHE or MAY_NOCACHE action - negative cachability policy
		* INVAL action - Dynamic Invalidation Policy
	*/
	Action string `json:"action,omitempty"`
	/**
	* Name of the content group in which to store the object when the final result of policy evaluation is CACHE. The content group must exist before being mentioned here. Use the "show cache contentgroup" command to view the list of existing content groups.
	*/
	Storeingroup string `json:"storeingroup,omitempty"`
	/**
	* Content group(s) to be invalidated when the INVAL action is applied. Maximum number of content groups that can be specified is 16.
	*/
	Invalgroups []string `json:"invalgroups,omitempty"`
	/**
	* Content groups(s) in which the objects will be invalidated if the action is INVAL.
	*/
	Invalobjects []string `json:"invalobjects,omitempty"`
	/**
	* Action to be performed when the result of rule evaluation is undefined.
	*/
	Undefaction string `json:"undefaction,omitempty"`
	/**
	* New name for the cache policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Flags string `json:"flags,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
