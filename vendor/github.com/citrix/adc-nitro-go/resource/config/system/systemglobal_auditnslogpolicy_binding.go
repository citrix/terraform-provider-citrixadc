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

package system

/**
* Binding class showing the auditnslogpolicy that can be bound to systemglobal.
*/
type Systemglobalauditnslogpolicybinding struct {
	/**
	* The name of the  command policy.
	*/
	Policyname string `json:"policyname,omitempty"`
	/**
	* The priority of the command policy.
	*/
	Priority int `json:"priority,omitempty"`
	/**
	* Indicates that a variable is a built-in (SYSTEM INTERNAL) type.
	*/
	Builtin []string `json:"builtin,omitempty"`
	/**
	* The feature to be checked while applying this config
	*/
	Feature string `json:"feature,omitempty"`
	Globalbindtype string `json:"globalbindtype,omitempty"`
	/**
	* On success invoke label. Applicable for advanced authentication policy binding
	*/
	Nextfactor string `json:"nextfactor,omitempty"`
	/**
	* Applicable only to advance authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE.  Specify one of the following values:
		* NEXT - Evaluate the policy with the next higher priority number.
		* END - End policy evaluation.
	*/
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`


}