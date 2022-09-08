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

package appqoe

/**
* Configuration for AppQoS policy resource.
*/
type Appqoepolicy struct {
	Name string `json:"name,omitempty"`
	/**
	* Expression or name of a named expression, against which the request is evaluated. The policy is applied if the rule evaluates to true.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Configured AppQoE action to trigger
	*/
	Action string `json:"action,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`

}
