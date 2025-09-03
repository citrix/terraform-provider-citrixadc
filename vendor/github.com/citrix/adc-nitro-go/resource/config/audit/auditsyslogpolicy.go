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

package audit

/**
* Configuration for system log policy resource.
*/
type Auditsyslogpolicy struct {
	/**
	* Name for the policy.
		Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the syslog policy is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my syslog policy" or 'my syslog policy').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the Citrix ADC named rule, or an expression, that defines the messages to be logged to the syslog server.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Syslog server action to perform when this policy matches traffic.
		NOTE: A syslog server action must be associated with a syslog audit policy.
	*/
	Action string `json:"action,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Expressiontype string `json:"expressiontype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
