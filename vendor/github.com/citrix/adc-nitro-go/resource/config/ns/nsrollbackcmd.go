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
* Configuration for Generate rollback commands resource.
*/
type Nsrollbackcmd struct {
	/**
	* File that contains the commands for which the rollback commands must be generated. Specify the full path of the file name.
	*/
	Filename string `json:"filename,omitempty"`
	/**
	* Format in which the rollback commands must be generated.
	*/
	Outtype string `json:"outtype,omitempty"`

}
