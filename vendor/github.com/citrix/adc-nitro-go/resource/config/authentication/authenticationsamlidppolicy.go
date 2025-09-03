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
* Configuration for AAA Saml IdentityProvider (IdP) policy resource.
*/
type Authenticationsamlidppolicy struct {
	/**
	* Name for the SAML Identity Provider (IdP) authentication policy. This is used for configuring Citrix ADC as SAML Identity Provider. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Expression which is evaluated to choose a profile for authentication.
		The following requirements apply only to the Citrix ADC CLI:
		* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
		* If the expression itself includes double quotation marks, escape the quotations by using the \ character.
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Name of the profile to apply to requests or connections that match this policy.
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
	* New name for the SAML IdentityProvider policy.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my samlidppolicy policy" or 'my samlidppolicy policy').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
