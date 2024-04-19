/*
* Copyright (c) 2024 Citrix Systems, Inc.
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

package lb

/**
* Configuration for lb policy resource.
 */
 type Lbpolicy struct {
	/**
	* Name of the LB policy.
	* Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the LB policy is added.
	* The following requirement applies only to the Citrix ADC CLI:
	* If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my lb policy" or 'my lb policy').
	 */
	Name string `json:"name,omitempty"`
	/**
	* Expression against which traffic is evaluated.
	 */
	Rule string `json:"rule,omitempty"`
	/**
	* Name of action to use if the request matches this LB policy.
	 */
	Action string `json:"action,omitempty"`
	/**
	* Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Available settings function as follows:
	* NOLBACTION - Does not consider LB actions in making LB decision.
	* RESET - Reset the request and notify the user, so that the user can resend the request.
	* DROP - Drop the request without sending a response to the user.
	 */
	Undefaction string `json:"undefaction,omitempty"`
	/**
	* Name of the messagelog action to use for requests that match this policy.
	 */
	Logaction string `json:"logaction,omitempty"`
	/**
	* Any type of information about this LB policy.
	 */
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the LB policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
	* The following requirement applies only to the Citrix ADC CLI:
	* If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my lb policy" or 'my lb policy').
	* Minimum length = 1
	 */
	Newname string `json:"newname,omitempty"`
}
