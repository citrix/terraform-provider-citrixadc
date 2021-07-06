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

package appfw

/**
* Binding class showing the appfwpolicy that can be bound to appfwpolicylabel.
*/
type Appfwpolicylabelappfwpolicybinding struct {
	/**
	* Name of the application firewall policy to bind to the policy label.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Positive integer specifying the priority of the policy. A lower number specifies a higher priority. Must be unique within a group of policies that are bound to the same bind point or label. Policies are evaluated in the order of their priority numbers.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
	*/
	Invoke bool `json:"invoke,omitempty"`
	/**
	* Type of policy label to invoke if the current policy evaluates to TRUE and the invoke parameter is set. Available settings function as follows:
		* reqvserver. Invoke the unnamed policy label associated with the specified request virtual server.
		* policylabel. Invoke the specified user-defined policy label.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
	*/
	Invokelabelname string `json:"invoke_labelname,omitempty"`
	/**
	* Name of the application firewall policy label.
	*/
	Labelname string `json:"labelname,omitempty"`


}