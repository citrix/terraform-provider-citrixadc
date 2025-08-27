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

package spillover

/**
* Configuration for Spillover policy resource.
*/
type Spilloverpolicy struct {
	/**
	* Name of the spillover policy.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Expression to be used by the spillover policy.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Action for the spillover policy. Action is created using add spillover action command
	*/
	Action string `json:"action,omitempty"`
	/**
	* Any comments that you might want to associate with the spillover policy.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the spillover policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
		Choose a name that reflects the function that the policy performs. 
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
