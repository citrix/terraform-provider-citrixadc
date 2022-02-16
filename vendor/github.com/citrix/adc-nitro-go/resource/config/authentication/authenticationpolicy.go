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

package authentication

/**
* Configuration for Authentication Policy resource.
*/
type Authenticationpolicy struct {
	/**
	* Name for the advance AUTHENTICATION policy. 
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after AUTHENTICATION policy is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the Citrix ADC named rule, or an expression, that the policy uses to determine whether to attempt to authenticate the user with the AUTHENTICATION server.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Name of the authentication action to be performed if the policy matches.
	*/
	Action string `json:"action,omitempty"`
	/**
	* Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.
	*/
	Undefaction string `json:"undefaction,omitempty"`
	/**
	* Any comments to preserve information about this policy.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Name of messagelog action to use when a request matches this policy.
	*/
	Logaction string `json:"logaction,omitempty"`
	/**
	* New name for the authentication policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. 
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication policy" or 'my authentication policy').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Description string `json:"description,omitempty"`
	Policysubtype string `json:"policysubtype,omitempty"`

}
