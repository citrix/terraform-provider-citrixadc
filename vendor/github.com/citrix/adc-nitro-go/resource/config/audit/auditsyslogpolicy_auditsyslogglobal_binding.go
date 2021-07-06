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

package audit

/**
* Binding class showing the auditsyslogglobal that can be bound to auditsyslogpolicy.
*/
type Auditsyslogpolicyauditsyslogglobalbinding struct {
	/**
	* The entity name to which policy is bound
	*/
	Boundto string `json:"boundto,omitempty"`
	Priority int `json:"priority,omitempty"`
	Activepolicy int `json:"activepolicy,omitempty"`
	/**
	* Name of the policy.
	*/
	Name string `json:"name,omitempty"`


}