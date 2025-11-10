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
* Binding class showing the feopolicy that can be bound to feoglobal.
*/
type Feoglobalfeopolicybinding struct {
	/**
	* The name of the globally bound front end optimization policy.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* The priority assigned to the policy binding.
	*/
	Priority *int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* Bindpoint to which the policy is bound.
	*/
	Type string `json:"type,omitempty"`
	/**
	* The number of policies bound to the bindpoint.
	*/
	Numpol *int `json:"numpol,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`


}