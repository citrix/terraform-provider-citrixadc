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
* Binding class showing the policy that can be bound to csvserver.
*/
type Csvserverpolicybinding struct {
	/**
	* Policies bound to this vserver.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Priority for the policy.
	*/
	Priority uint32 `json:"priority,omitempty"`
	/**
	* The state of SureConnect the specified virtual server.
	*/
	Sc string `json:"sc,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* The bindpoint to which the policy is bound
	*/
	Bindpoint string `json:"bindpoint,omitempty"`
	/**
	* Invoke flag.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* The invocation type.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the label invoked.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* target vserver name.
	*/
	Targetlbvserver string `json:"targetlbvserver,omitempty"`
	/**
	* Number of hits.
	*/
	Hits uint32 `json:"hits,omitempty"`
	/**
	* Number of hits.
	*/
	Pipolicyhits uint32 `json:"pipolicyhits,omitempty"`
	/**
	* Rule.
	*/
	Rule string `json:"rule,omitempty"`
	/**
	* Vserver Id of vserver
	*/
	Vserverid string `json:"vserverid,omitempty"`
	/**
	* Vserver id of the lb vserver that is inserted into the set-cookie HTTP header
	*/
	Cookieipport string `json:"cookieipport,omitempty"`
	/**
	* Name of the content switching virtual server to which the content switching policy applies.
	*/
	Name string `json:"name,omitempty"`


}