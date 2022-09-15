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

package cache

/**
* Binding class showing the policy that can be bound to cacheglobal.
*/
type Cacheglobalpolicybinding struct {
	/**
	* Name of the cache policy.
	*/
	Policy string `json:"policy,omitempty"`
	/**
	* The bind point to which policy is bound. When you specify the type, detailed information about that bind point appears.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Specifies the priority of the policy.
	*/
	Priority uint32 `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority. Applicable only to default-syntax policies.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of policy label to invoke.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the label to invoke if the current policy rule evaluates to TRUE. (To invoke a label associated with a virtual server, specify the name of the virtual server.)
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* The number of policies bound to the bindpoint.
	*/
	Numpol uint32 `json:"numpol,omitempty"`
	/**
	* flowtype of the bound cache policy.
	*/
	Flowtype uint32 `json:"flowtype,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`
	/**
	* Specify whether this policy should be evaluated.
	*/
	Precededefrules string `json:"precededefrules,omitempty"`


}