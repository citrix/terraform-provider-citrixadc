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

package ha

/**
* Configuration for sync resource.
*/
type Hasync struct {
	/**
	* Force synchronization regardless of the state of HA propagation and HA synchronization on either node.
	*/
	Force bool `json:"force,omitempty"`
	/**
	* After synchronization, automatically save the configuration in the secondary node configuration file (ns.conf) without prompting for confirmation.
	*/
	Save string `json:"save,omitempty"`

}
