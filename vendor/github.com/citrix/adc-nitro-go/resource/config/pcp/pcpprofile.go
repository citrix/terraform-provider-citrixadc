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

package pcp

/**
* Configuration for PCP Profile resource.
*/
type Pcpprofile struct {
	/**
	* Name for the PCP Profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my pcpProfile" or my pcpProfile).
	*/
	Name string `json:"name,omitempty"`
	/**
	* This argument is for enabling/disabling the MAP opcode  of current PCP Profile
	*/
	Mapping string `json:"mapping,omitempty"`
	/**
	* This argument is for enabling/disabling the PEER opcode of current PCP Profile
	*/
	Peer string `json:"peer,omitempty"`
	/**
	* Integer value that identify the minimum mapping lifetime (in seconds) for a pcp profile. default(120s)
	*/
	Minmaplife *int `json:"minmaplife,omitempty"`
	/**
	* Integer value that identify the maximum mapping lifetime (in seconds) for a pcp profile. default(86400s = 24Hours).
	*/
	Maxmaplife *int `json:"maxmaplife,omitempty"`
	/**
	* Integer value that identify the number announce message to be send.
	*/
	Announcemulticount *int `json:"announcemulticount"` // Zero is a valid value
	/**
	* This argument is for enabling/disabling the THIRD PARTY opcode of current PCP Profile
	*/
	Thirdparty string `json:"thirdparty,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
