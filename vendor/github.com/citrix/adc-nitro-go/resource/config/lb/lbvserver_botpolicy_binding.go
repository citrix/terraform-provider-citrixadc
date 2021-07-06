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
* Binding class showing the botpolicy that can be bound to lbvserver.
*/
type Lbvserverbotpolicybinding struct {
	/**
	* Name of the policy bound to the LB vserver.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Priority.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Invoke policies bound to a virtual server or policy label.
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
	* Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the virtual server is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vserver" or 'my vserver'). 
	*/
	Name string `json:"name,omitempty"`
	/**
	* Bind point to which to bind the policy.
	*/
	Bindpoint string `json:"bindpoint,omitempty"`


}