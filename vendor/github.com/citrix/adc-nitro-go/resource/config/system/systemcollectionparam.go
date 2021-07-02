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
* Configuration for collection parameter resource.
*/
type Systemcollectionparam struct {
	/**
	* SNMPv1 community name for authentication.
	*/
	Communityname string `json:"communityname,omitempty"`
	/**
	* specify the log level. Possible values CRITICAL,WARNING,INFO,DEBUG1,DEBUG2
	*/
	Loglevel string `json:"loglevel,omitempty"`
	/**
	* specify the data path to the database.
	*/
	Datapath string `json:"datapath,omitempty"`

}
