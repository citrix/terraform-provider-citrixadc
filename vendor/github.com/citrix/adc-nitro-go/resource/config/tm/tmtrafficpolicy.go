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

package tm

/**
* Configuration for TM traffic policy resource.
*/
type Tmtrafficpolicy struct {
	/**
	* Name for the traffic policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the Citrix ADC named expression, or an expression, that the policy uses to determine whether to apply certain action on the current traffic.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Name of the action to apply to requests or connections that match this policy.
	*/
	Action string `json:"action,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
