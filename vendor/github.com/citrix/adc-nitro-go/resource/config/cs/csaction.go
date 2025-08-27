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

package cs

/**
* Configuration for Content Switching action resource.
*/
type Csaction struct {
	/**
	* Name for the content switching action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the content switching action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the load balancing virtual server to which the content is switched.
	*/
	Targetlbvserver string `json:"targetlbvserver,omitempty"`
	/**
	* Name of the VPN, GSLB or Authentication virtual server to which the content is switched.
	*/
	Targetvserver string `json:"targetvserver,omitempty"`
	/**
	* Information about this content switching action.
	*/
	Targetvserverexpr string `json:"targetvserverexpr,omitempty"`
	/**
	* Comments associated with this cs action.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the content switching action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my name" or 'my name').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
