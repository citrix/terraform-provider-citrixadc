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

package authentication

/**
* Binding class showing the vserver that can be bound to authenticationdfapolicy.
*/
type Authenticationdfapolicyvserverbinding struct {
	/**
	* The entity name to which policy is bound
	*/
	Boundto string `json:"boundto,omitempty"`
	Priority uint32 `json:"priority,omitempty"`
	Activepolicy uint32 `json:"activepolicy,omitempty"`
	/**
	* Name of the WebAuth policy.
	*/
	Name string `json:"name,omitempty"`


}