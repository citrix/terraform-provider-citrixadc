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

package autoscale

/**
* Configuration for Autoscale policy resource.
*/
type Autoscalepolicy struct {
	/**
	* The name of the autoscale policy.
	*/
	Name string `json:"name,omitempty"`
	/**
	* The rule associated with the policy.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* The autoscale profile associated with the policy.
	*/
	Action string `json:"action,omitempty"`
	/**
	* Comments associated with this autoscale policy.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* The log action associated with the autoscale policy
	*/
	Logaction string `json:"logaction,omitempty"`
	/**
	* The new name of the autoscale policy.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Hits string `json:"hits,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Priority string `json:"priority,omitempty"`

}
