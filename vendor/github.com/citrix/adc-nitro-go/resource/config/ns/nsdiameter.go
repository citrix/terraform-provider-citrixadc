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
* Configuration for Diameter Parameters resource.
*/
type Nsdiameter struct {
	/**
	* DiameterIdentity to be used by NS. DiameterIdentity is used to identify a Diameter node uniquely. Before setting up diameter configuration, Citrix ADC (as a Diameter node) MUST be assigned a unique DiameterIdentity.
		example =>
		set ns diameter -identity netscaler.com
		Now whenever Citrix ADC needs to use identity in diameter messages. It will use 'netscaler.com' as Origin-Host AVP as defined in RFC3588
	*/
	Identity string `json:"identity,omitempty"`
	/**
	* Diameter Realm to be used by NS.
		example =>
		set ns diameter -realm com
		Now whenever Citrix ADC system needs to use realm in diameter messages. It will use 'com' as Origin-Realm AVP as defined in RFC3588
	*/
	Realm string `json:"realm,omitempty"`
	/**
	* when a Server connection goes down, whether to close the corresponding client connection if there were requests pending on the server.
	*/
	Serverclosepropagation string `json:"serverclosepropagation,omitempty"`
	/**
	* ID of the cluster node for which the diameter id is set, can be configured only through CLIP
	*/
	Ownernode uint32 `json:"ownernode,omitempty"`

}
