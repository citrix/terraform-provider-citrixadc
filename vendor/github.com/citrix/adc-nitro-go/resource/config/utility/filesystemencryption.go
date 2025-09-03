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

package utility

/**
* Configuration for File System Encryption Information resource.
*/
type Filesystemencryption struct {
	/**
	* Number of times /flash directory has to be written with 0s.
	*/
	Ntimes0flash int `json:"ntimes0flash,omitempty"`
	/**
	* Number of times /var directory has to be written with 0s.
	*/
	Ntimes0var int `json:"ntimes0var,omitempty"`
	/**
	* Encryption Passphrase.
	*/
	Passphrase string `json:"passphrase,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Supportedstate string `json:"supportedstate,omitempty"`
	Effectivestate string `json:"effectivestate,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
