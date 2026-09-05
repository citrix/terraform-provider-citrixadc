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
* Binding class showing the sslpolicy that can be bound to sslservice.
*/
type Sslservicesslpolicybinding struct {
	/**
	* The SSL policy binding.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* The priority of the policies bound to this SSL service
	*/
	Priority *int `json:"priority,omitempty"`
	/**
	* The phase of the SSL connection in which the policy rule is evaluated. Possible Value: SERVER_AUTH_VAL_REQ. The bindpoint means::
		1. SERVER_AUTH_VAL_REQ: Policy evaluation will be performed during the verification of the server certificate. Action allowed with this type is: OCSPCERTVALIDATION.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Whether the bound policy is a inherited policy or not
	*/
	Polinherit *int `json:"polinherit,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Invoke flag. This attribute is relevant only for ADVANCED policies
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of policy label invocation.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the label to invoke if the current policy rule evaluates to TRUE.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Name of the SSL service for which to set advanced configuration.
	*/
	Servicename string `json:"servicename,omitempty"`


}