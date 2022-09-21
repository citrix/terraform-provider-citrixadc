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
* Configuration for autoscale action resource.
*/
type Autoscaleaction struct {
	/**
	* ActionScale action name.
	*/
	Name string `json:"name,omitempty"`
	/**
	* The type of action.
	*/
	Type string `json:"type,omitempty"`
	/**
	* AutoScale profile name.
	*/
	Profilename string `json:"profilename,omitempty"`
	/**
	* Parameters to use in the action
	*/
	Parameters string `json:"parameters,omitempty"`
	/**
	* Time in minutes a VM is kept in inactive state before destroying
	*/
	Vmdestroygraceperiod int `json:"vmdestroygraceperiod,omitempty"`
	/**
	* Time in seconds no other policy is evaluated or action is taken
	*/
	Quiettime int `json:"quiettime,omitempty"`
	/**
	* Name of the vserver on which autoscale action has to be taken.
	*/
	Vserver string `json:"vserver,omitempty"`

}
