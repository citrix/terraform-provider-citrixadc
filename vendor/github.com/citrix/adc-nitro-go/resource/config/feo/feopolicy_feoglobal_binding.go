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

package feo

/**
* Binding class showing the feoglobal that can be bound to feopolicy.
*/
type Feopolicyfeoglobalbinding struct {
	/**
	* Location where the policy is bound to.
	*/
	Boundto string `json:"boundto,omitempty"`
	/**
	* Priority of the policy.
	*/
	Priority *int `json:"priority,omitempty"`
	/**
	* Indicates whether a policy is bound or not.
	*/
	Activepolicy *int `json:"activepolicy,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* The name of the front end optimization policy.
	*/
	Name string `json:"name,omitempty"`


}