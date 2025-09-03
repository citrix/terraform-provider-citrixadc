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

package ns

/**
* Configuration for Migration operation resource.
*/
type Nsmigration struct {
	/**
	* Displays the current active migrated session details, if DUMPSESSION option is YES.
	*/
	Dumpsession string `json:"dumpsession,omitempty"`

	//------- Read only Parameter ---------;

	Migrationstatus string `json:"migrationstatus,omitempty"`
	Migrationstarttime string `json:"migrationstarttime,omitempty"`
	Migrationendtime string `json:"migrationendtime,omitempty"`
	Migrationrollbackstarttime string `json:"migrationrollbackstarttime,omitempty"`
	Srcip string `json:"srcip,omitempty"`
	Srcport string `json:"srcport,omitempty"`
	Destip string `json:"destip,omitempty"`
	Destport string `json:"destport,omitempty"`
	Timeout string `json:"timeout,omitempty"`
	Migdfdsessionsallocated string `json:"migdfdsessionsallocated,omitempty"`
	Migdfdsessionsactive string `json:"migdfdsessionsactive,omitempty"`
	Migl4sessionsallocated string `json:"migl4sessionsallocated,omitempty"`
	Migl4sessionsactive string `json:"migl4sessionsactive,omitempty"`
	Migdfdsessionsallocatedrollback string `json:"migdfdsessionsallocatedrollback,omitempty"`
	Migdfdsessionsactiverollback string `json:"migdfdsessionsactiverollback,omitempty"`
	Migl4sessionsallocatedrollback string `json:"migl4sessionsallocatedrollback,omitempty"`
	Migl4sessionsactiverollback string `json:"migl4sessionsactiverollback,omitempty"`
	Mighastateflag string `json:"mighastateflag,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
