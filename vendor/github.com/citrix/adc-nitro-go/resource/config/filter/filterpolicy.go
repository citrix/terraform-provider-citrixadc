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

package filter

/**
* Configuration for filter policy resource.
*/
type Filterpolicy struct {
	/**
	* Name for the filtering action. Must begin with a letter, number, or the underscore character (_). Other characters allowed, after the first character, are the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), and colon (:) characters. Choose a name that helps identify the type of action. The name cannot be updated after the policy is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Citrix ADC classic expression specifying the type of connections that match this policy.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Name of the action to be performed on requests that match the policy. Cannot be specified if the rule includes condition to be evaluated for responses.
	*/
	Reqaction string `json:"reqaction,omitempty"`
	/**
	* The action to be performed on the response. The string value can be a filter action created filter action or a built-in action.
	*/
	Resaction string `json:"resaction,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`

}
