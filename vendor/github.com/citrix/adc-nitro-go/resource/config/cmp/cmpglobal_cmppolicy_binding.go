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

package cmp

/**
* Binding class showing the cmppolicy that can be bound to cmpglobal.
*/
type Cmpglobalcmppolicybinding struct {
	/**
	* The name of the globally bound HTTP compression policy.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Positive integer specifying the priority of the policy. The lower the number, the higher the priority. By default, polices within a label are evaluated in the order of their priority numbers.
		In the configuration utility, you can click the Priority field and edit the priority level or drag the entry to a new position in the list. If you drag the entry to a new position, the priority level is updated automatically.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* The current state of the policy binding. This attribute is relevant only for CLASSIC policies.
	*/
	State string `json:"state,omitempty"`
	/**
	* Bind point to which the policy is bound.
	*/
	Type string `json:"type,omitempty"`
	/**
	* The number of policies bound to the bindpoint.
	*/
	Numpol int `json:"numpol,omitempty"`
	/**
	* Policy type (Classic/Advanced) to be bound.Used for display.
	*/
	Policytype string `json:"policytype,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`
	/**
	* Expression or other value specifying the priority of the next policy, within the policy label, to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:
		* NEXT - Evaluate the policy with the next higher numbered priority.
		* END - Stop evaluation.
		* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.
		* An expression that evaluates to a number.
		If you specify an expression, it's evaluation result determines the next policy to evaluate, as follows: 
		* If the expression evaluates to a higher numbered priority, that policy is evaluated next.
		* If the expression evaluates to the priority of the current policy, the policy with the next higher priority number is evaluated next.
		* If the expression evaluates to a priority number that is numerically higher than the highest priority number, policy evaluation ends.
		An UNDEF event is triggered if:
		* The expression is invalid.
		* The expression evaluates to a priority number that is numerically lower than the current policy's priority.
		* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Invoke policies bound to a virtual server or a policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority. Applicable only for default-syntax policies.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of policy label invocation. This argument is relevant only for advanced (default-syntax) policies.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the label to invoke if the current policy rule evaluates to TRUE. Applicable only to advanced (default-syntax) policies.
	*/
	Labelname string `json:"labelname,omitempty"`


}