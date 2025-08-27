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
* Binding class showing the cspolicylabel that can be bound to cspolicy.
*/
type Cspolicycspolicylabelbinding struct {
	/**
	* Location where policy is bound
	*/
	Boundto string `json:"boundto,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* priority of bound policy
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Total number of hits.
	*/
	Hits int `json:"hits,omitempty"`
	/**
	* The invocation type.
	*/
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the label invoked.
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Name of the content switching policy to display. If this parameter is omitted, details of all the policies are displayed.
	*/
	Policyname string `json:"policyname,omitempty"`


}