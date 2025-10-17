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

package tunnel

/**
* Binding class showing the tunneltrafficpolicy that can be bound to tunnelglobal.
*/
type Tunnelglobaltunneltrafficpolicybinding struct {
	/**
	* Policy name.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Priority.
	*/
	Priority *int `json:"priority,omitempty"`
	/**
	* Current state of the binding. If the binding is enabled, the policy is active.
	*/
	State string `json:"state,omitempty"`
	/**
	* Indicates that a variable is a built-in (SYSTEM INTERNAL) type.
	*/
	Builtin []string `json:"builtin,omitempty"`
	/**
	* The feature to be checked while applying this config
	*/
	Feature string `json:"feature,omitempty"`
	/**
	* Bind point to which the policy is bound.
	*/
	Type string `json:"type,omitempty"`
	/**
	* The number of policies bound to the bindpoint.
	*/
	Numpol *int `json:"numpol,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Policy type (Classic/Advanced) to be bound.Used for display.
	*/
	Policytype string `json:"policytype,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`


}