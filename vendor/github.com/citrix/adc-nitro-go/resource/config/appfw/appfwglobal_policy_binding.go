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

package appfw

/**
* Binding class showing the policy that can be bound to appfwglobal.
*/
type Appfwglobalpolicybinding struct {
	/**
	* Name of the policy.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* The priority of the policy.
	*/
	Priority uint32 `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Enable or disable the binding to activate or deactivate the policy. This is applicable to classic policies only.
	*/
	State string `json:"state,omitempty"`
	/**
	* Type of policy label invocation.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* The number of policies bound to the bindpoint.
	*/
	Numpol uint32 `json:"numpol,omitempty"`
	/**
	* flowtype of the bound application firewall policy.
	*/
	Flowtype uint32 `json:"flowtype,omitempty"`
	/**
	* Bind point to which to policy is bound.
	*/
	Type string `json:"type,omitempty"`
	Policytype string `json:"policytype,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`


}