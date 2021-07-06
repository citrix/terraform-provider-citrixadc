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

package ssl

/**
* Binding class showing the policy that can be bound to sslpolicylabel.
*/
type Sslpolicylabelpolicybinding struct {
	/**
	* Name of the SSL policy to bind to the policy label.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Specifies the priority of the policy.
	*/
	Priority uint32 `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Type of policy label invocation.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the label to invoke if the current policy rule evaluates to TRUE.
	*/
	Invokelabelname string `json:"invoke_labelname,omitempty"`
	/**
	* Name of the SSL policy label to which to bind policies.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Invoke policies bound to a policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.
	*/
	Invoke bool `json:"invoke,omitempty"`


}