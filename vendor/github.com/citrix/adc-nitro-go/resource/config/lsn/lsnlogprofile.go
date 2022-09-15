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

package lsn

/**
* Configuration for LSN logging Profile resource.
*/
type Lsnlogprofile struct {
	/**
	* The name of the logging Profile.
	*/
	Logprofilename string `json:"logprofilename,omitempty"`
	/**
	* Subscriber ID information is logged if option is enabled.
	*/
	Logsubscrinfo string `json:"logsubscrinfo,omitempty"`
	/**
	* Logs in Compact Logging format if option is enabled.
	*/
	Logcompact string `json:"logcompact,omitempty"`
	/**
	* Logs in IPFIX  format if option is enabled.
	*/
	Logipfix string `json:"logipfix,omitempty"`
	/**
	* Name of the Analytics Profile attached to this lsn profile.
	*/
	Analyticsprofile string `json:"analyticsprofile,omitempty"`
	/**
	* LSN Session deletion will not be logged if disabled.
	*/
	Logsessdeletion string `json:"logsessdeletion,omitempty"`

}
