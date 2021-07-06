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

package responder

/**
* Configuration for responder policy resource.
*/
type Responderpolicy struct {
	/**
	* Name for the responder policy.
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the responder policy is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my responder policy" or 'my responder policy').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Expression that the policy uses to determine whether to respond to the specified request.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Name of the responder action to perform if the request matches this responder policy. There are also some built-in actions which can be used. These are:
		* NOOP - Send the request to the protected server instead of responding to it.
		* RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.
		* DROP - Drop the request without sending a response to the user.
	*/
	Action string `json:"action,omitempty"`
	/**
	* Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition. Only the above built-in actions can be used.
	*/
	Undefaction string `json:"undefaction,omitempty"`
	/**
	* Any type of information about this responder policy.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Name of the messagelog action to use for requests that match this policy.
	*/
	Logaction string `json:"logaction,omitempty"`
	/**
	* AppFlow action to invoke for requests that match this policy.
	*/
	Appflowaction string `json:"appflowaction,omitempty"`
	/**
	* New name for the responder policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. 
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my responder policy" or 'my responder policy').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
