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

package ntp

/**
* Configuration for NTP parameter resource.
*/
type Ntpparam struct {
	/**
	* Apply NTP authentication, which enables the NTP client (Citrix ADC) to verify that the server is in fact known and trusted.
	*/
	Authentication string `json:"authentication,omitempty"`
	/**
	* Key identifiers that are trusted for server authentication with symmetric key cryptography in the keys file.
	*/
	Trustedkey []int `json:"trustedkey,omitempty"`
	/**
	* Autokey protocol requires the keys to be refreshed periodically. This parameter specifies the interval between regenerations of new session keys. In seconds, expressed as a power of 2.
	*/
	Autokeylogsec int `json:"autokeylogsec"` // Zero is a valid value
	/**
	* Interval between re-randomizations of the autokey seeds to prevent brute-force attacks on the autokey algorithms.
	*/
	Revokelogsec int `json:"revokelogsec"` // Zero is a valid value

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
