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

package ssl

/**
* Configuration for DTLS profile resource.
*/
type Ssldtlsprofile struct {
	/**
	* Name for the DTLS profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Source for the maximum record size value. If ENABLED, the value is taken from the PMTU table. If DISABLED, the value is taken from the profile.
	*/
	Pmtudiscovery string `json:"pmtudiscovery,omitempty"`
	/**
	* Maximum size of records that can be sent if PMTU is disabled.
	*/
	Maxrecordsize uint32 `json:"maxrecordsize,omitempty"`
	/**
	* Wait for the specified time, in seconds, before resending the request.
	*/
	Maxretrytime uint32 `json:"maxretrytime,omitempty"`
	/**
	* Send a Hello Verify request to validate the client.
	*/
	Helloverifyrequest string `json:"helloverifyrequest,omitempty"`
	/**
	* Terminate the session if the message authentication code (MAC) of the client and server do not match.
	*/
	Terminatesession string `json:"terminatesession,omitempty"`
	/**
	* Maximum number of packets to reassemble. This value helps protect against a fragmented packet attack.
	*/
	Maxpacketsize uint32 `json:"maxpacketsize,omitempty"`
	/**
	* Maximum number of datagrams that can be queued at DTLS layer for processing
	*/
	Maxholdqlen uint32 `json:"maxholdqlen,omitempty"`
	/**
	* Maximum number of bad MAC errors to ignore for a connection prior disconnect. Disabling parameter terminateSession terminates session immediately when bad MAC is detected in the connection.
	*/
	Maxbadmacignorecount uint32 `json:"maxbadmacignorecount,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
