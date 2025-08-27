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

package cr

/**
* Binding class showing the policymap that can be bound to crvserver.
*/
type Crvserverpolicymapbinding struct {
	/**
	* Policies bound to this vserver.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* The CSW target server names.
	*/
	Targetvserver string `json:"targetvserver,omitempty"`
	/**
	* Name of the cache redirection virtual server to which to bind the cache redirection policy.
	*/
	Name string `json:"name,omitempty"`
	/**
	* An unsigned integer that determines the priority of the policy relative to other policies bound to this cache redirection virtual server. The lower the value, higher the priority. Note: This option is available only when binding content switching, filtering, and compression policies to a cache redirection virtual server.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:
		* NEXT - Evaluate the policy with the next higher priority number.
		* END - End policy evaluation.
		* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.
		* An expression that evaluates to a number.
		If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:
		* If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.
		* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.
		* If the expression evaluates to a priority number that is numerically higher than the highest numbered priority, policy evaluation ends.
		An UNDEF event is triggered if:
		* The expression is invalid.
		* The expression evaluates to a priority number that is numerically lower than the current policy's priority.
		* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), b
		ut does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number incr
		ements by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite
		policies, because content switching policies are evaluated only at request time.
	*/
	Bindpoint string `json:"bindpoint,omitempty"`
	/**
	* Invoke a policy label if this policy's rule evaluates to TRUE.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of label to be invoked.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the label to be invoked.
	*/
	Labelname string `json:"labelname,omitempty"`


}