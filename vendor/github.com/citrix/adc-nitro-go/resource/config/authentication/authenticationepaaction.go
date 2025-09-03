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
* Configuration for epa action resource.
*/
type Authenticationepaaction struct {
	/**
	* Name for the epa action. Must begin with a
		letter, number, or the underscore character (_), and must consist
		only of letters, numbers, and the hyphen (-), period (.) pound
		(#), space ( ), at (@), equals (=), colon (:), and underscore
		characters. Cannot be changed after epa action is created.The following requirement applies only to the Citrix ADC CLI:If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my aaa action" or 'my aaa action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* it holds the ClientSecurityExpression to be sent to the client
	*/
	Csecexpr string `json:"csecexpr,omitempty"`
	/**
	* String specifying the name of a process to be terminated by the endpoint analysis (EPA) tool. Multiple processes to be delimited by comma
	*/
	Killprocess string `json:"killprocess,omitempty"`
	/**
	* String specifying the path(s) and name(s) of the files to be deleted by the endpoint analysis (EPA) tool. Multiple files to be delimited by comma
	*/
	Deletefiles string `json:"deletefiles,omitempty"`
	/**
	* This is the default group that is chosen when the EPA check succeeds.
	*/
	Defaultepagroup string `json:"defaultepagroup,omitempty"`
	/**
	* This is the quarantine group that is chosen when the EPA check fails
		if configured.
	*/
	Quarantinegroup string `json:"quarantinegroup,omitempty"`
	/**
	* Parameter to enable/disable device posture service scan
	*/
	Deviceposture string `json:"deviceposture,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
