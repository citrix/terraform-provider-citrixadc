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
* Configuration for Backup Data for ns backup and restore resource.
*/
type Systembackup struct {
	/**
	* Name of the backup file(*.tgz) to be restored.
	*/
	Filename string `json:"filename,omitempty"`
	/**
	* This option will create backup file with local timezone timestamp
	*/
	Uselocaltimezone bool `json:"uselocaltimezone,omitempty"`
	/**
	* Level of data to be backed up.
	*/
	Level string `json:"level,omitempty"`
	/**
	* Use this option to add kernel in the backup file
	*/
	Includekernel string `json:"includekernel,omitempty"`
	/**
	* Comment specified at the time of creation of the backup file(*.tgz).
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* Use this option to skip taking backup during restore operation
	*/
	Skipbackup bool `json:"skipbackup,omitempty"`

	//------- Read only Parameter ---------;

	Size string `json:"size,omitempty"`
	Creationtime string `json:"creationtime,omitempty"`
	Version string `json:"version,omitempty"`
	Createdby string `json:"createdby,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
