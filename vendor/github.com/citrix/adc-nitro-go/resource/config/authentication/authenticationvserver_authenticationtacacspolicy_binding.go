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

package authentication

/**
* Binding class showing the authenticationtacacspolicy that can be bound to authenticationvserver.
*/
type Authenticationvserverauthenticationtacacspolicybinding struct {
	/**
	* The name of the policy, if any, bound to the authentication vserver.
	*/
	Policy string `json:"policy,omitempty"`
	/**
	* The priority, if any, of the vpn vserver policy.
	*/
	Priority int `json:"priority,omitempty"`
	Acttype int `json:"acttype,omitempty"`
	/**
	* Bind the authentication policy to the secondary chain.
		Provides for multifactor authentication in which a user must authenticate via both a primary authentication method and, afterward, via a secondary authentication method.
		Because user groups are aggregated across authentication systems, usernames must be the same on all authentication servers. Passwords can be different.
	*/
	Secondary bool `json:"secondary,omitempty"`
	/**
	* Name of the authentication virtual server to which to bind the policy.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
	*/
	Groupextraction bool `json:"groupextraction,omitempty"`
	/**
	* Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor
	*/
	Nextfactor string `json:"nextfactor,omitempty"`
	/**
	* Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:
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
	* Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.
	*/
	Bindpoint string `json:"bindpoint,omitempty"`


}