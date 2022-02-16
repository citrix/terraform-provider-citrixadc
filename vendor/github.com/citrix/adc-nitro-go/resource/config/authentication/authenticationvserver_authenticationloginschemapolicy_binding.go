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
* Binding class showing the authenticationloginschemapolicy that can be bound to authenticationvserver.
*/
type Authenticationvserverauthenticationloginschemapolicybinding struct {
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
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Name of the authentication virtual server to which to bind the policy.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
	*/
	Secondary bool `json:"secondary,omitempty"`
	/**
	* Applicable only while bindind classic authentication policy as advance authentication policy use nFactor
	*/
	Groupextraction bool `json:"groupextraction,omitempty"`
	/**
	* Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor
	*/
	Nextfactor string `json:"nextfactor,omitempty"`
	/**
	* Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.
	*/
	Bindpoint string `json:"bindpoint,omitempty"`


}