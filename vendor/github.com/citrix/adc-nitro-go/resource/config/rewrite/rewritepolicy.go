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

package rewrite

/**
* Configuration for rewrite policy resource.
*/
type Rewritepolicy struct {
	/**
	* Name for the rewrite policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my rewrite policy" or 'my rewrite policy').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Expression against which traffic is evaluated.
		The following requirements apply only to the Citrix ADC CLI:
		* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.
		* If the expression itself includes double quotation marks, escape the quotations by using the \ character. 
		* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Name of the rewrite action to perform if the request or response matches this rewrite policy.
		There are also some built-in actions which can be used. These are:
		* NOREWRITE - Send the request from the client to the server or response from the server to the client without making any changes in the message.
		* RESET - Resets the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.
		* DROP - Drop the request without sending a response to the user.
	*/
	Action string `json:"action,omitempty"`
	/**
	* Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.
	*/
	Undefaction string `json:"undefaction,omitempty"`
	/**
	* Any comments to preserve information about this rewrite policy.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Name of messagelog action to use when a request matches this policy.
	*/
	Logaction string `json:"logaction,omitempty"`
	/**
	* New name for the rewrite policy. 
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my rewrite policy" or 'my rewrite policy').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Description string `json:"description,omitempty"`
	Isdefault string `json:"isdefault,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
