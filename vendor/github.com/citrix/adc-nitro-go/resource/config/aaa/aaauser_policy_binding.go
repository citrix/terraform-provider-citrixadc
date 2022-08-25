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

package aaa

/**
* Binding class showing the policy that can be bound to aaauser.
*/
type Aaauserpolicybinding struct {
	/**
	* The policy Name.
	*/
	Policy string `json:"policy,omitempty"`
	/**
	* Integer specifying the priority of the policy.  A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies max priority is 64000. 
	*/
	Priority uint32 `json:"priority,omitempty"`
	Acttype uint32 `json:"acttype,omitempty"`
	/**
	* Bindpoint to which the policy is bound.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* User account to which to bind the policy.
	*/
	Username string `json:"username,omitempty"`


}