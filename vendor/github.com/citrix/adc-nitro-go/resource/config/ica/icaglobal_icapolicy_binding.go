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

package ica

/**
* Binding class showing the icapolicy that can be bound to icaglobal.
*/
type Icaglobalicapolicybinding struct {
	/**
	* Name of the ICA policy.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* Global bind point for which to show detailed information about the policies bound to the bind point.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Specifies the priority of the policy.
	*/
	Priority *int `json:"priority,omitempty"`
	/**
	* Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	/**
	* The number of policies bound to the bindpoint.
	*/
	Numpol *int `json:"numpol,omitempty"`
	/**
	* Flow type of the bound ICA policy.
	*/
	Flowtype *int `json:"flowtype,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`


}