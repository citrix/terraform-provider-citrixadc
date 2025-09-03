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
* Binding class showing the auditsyslogpolicy that can be bound to csvserver.
*/
type Csvserverauditsyslogpolicybinding struct {
	/**
	* Policies bound to this vserver.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Priority for the policy.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Name of the content switching virtual server to which the content switching policy applies.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the Load Balancing virtual server to which the content is switched, if policy rule is evaluated to be TRUE.
		Example: bind cs vs cs1 -policyname pol1 -priority 101 -targetLBVserver lb1
		Note: Use this parameter only in case of Content Switching policy bind operations to a CS vserver
	*/
	Targetlbvserver string `json:"targetlbvserver,omitempty"`
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
		* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Bind point at which policy needs to be bound. Note: Content switching policies are evaluated only at request time.
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