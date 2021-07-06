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

package lb

/**
* Binding class showing the auditnslogpolicy that can be bound to lbvserver.
*/
type Lbvserverauditnslogpolicybinding struct {
	/**
	* Name of the policy bound to the LB vserver.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Priority.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Use SureConnect on the virtual server.
	*/
	Sc string `json:"sc,omitempty"`
	/**
	* Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). 
	*/
	Name string `json:"name,omitempty"`
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
	* Bind point to which to bind the policy.
	*/
	Bindpoint string `json:"bindpoint,omitempty"`
	/**
	* Invoke policies bound to a virtual server or policy label.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of policy label to invoke. Applicable only to rewrite, videooptimization and cache policies. Available settings function as follows:
		* reqvserver - Evaluate the request against the request-based policies bound to the specified virtual server.
		* resvserver - Evaluate the response against the response-based policies bound to the specified virtual server.
		* policylabel - invoke the request or response against the specified user-defined policy label.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the virtual server or user-defined policy label to invoke if the policy evaluates to TRUE.
	*/
	Labelname string `json:"labelname,omitempty"`


}