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
* Binding class showing the nspartition that can be bound to systemuser.
*/
type Systemusernspartitionbinding struct {
	/**
	* Name of the Partition to bind to the system user.
	*/
	Partitionname string `json:"partitionname,omitempty"`
	/**
	* Name of the system-user entry to which to bind the command policy.
	*/
	Username string `json:"username,omitempty"`


}