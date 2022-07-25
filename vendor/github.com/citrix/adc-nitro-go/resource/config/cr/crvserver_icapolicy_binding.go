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
* Binding class showing the icapolicy that can be bound to crvserver.
*/
type Crvservericapolicybinding struct {
	/**
	* Policies bound to this vserver.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* The priority for the policy.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Name of the cache redirection virtual server to which to bind the cache redirection policy.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.
	*/
	Targetvserver string `json:"targetvserver,omitempty"`
	/**
	* For a rewrite policy, the bind point to which to bind the policy. Note: This parameter applies only to rewrite
		policies, because content switching policies are evaluated only at request time.
	*/
	Bindpoint string `json:"bindpoint,omitempty"`
	/**
	* Invoke a policy label if this policy's rule evaluates to TRUE (valid only for default-syntax policies such as
		application firewall, transform, integrated cache, rewrite, responder, and content switching).
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