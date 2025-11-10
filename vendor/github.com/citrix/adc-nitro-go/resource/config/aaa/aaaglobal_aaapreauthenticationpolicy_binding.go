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

package aaa

/**
* Binding class showing the aaapreauthenticationpolicy that can be bound to aaaglobal.
 */
type Aaaglobalaaapreauthenticationpolicybinding struct {
	/**
	* Name of the policy to be unbound.
	 */
	Policy string `json:"policy,omitempty"`
	/**
	* Priority of the bound policy
	 */
	Priority *int `json:"priority,omitempty"`
	/**
	* Bound policy type
	 */
	Bindpolicytype *int `json:"bindpolicytype,omitempty"`
	/**
	* Indicates that a variable is a built-in (SYSTEM INTERNAL) type.
	 */
	Builtin []string `json:"builtin,omitempty"`
}
