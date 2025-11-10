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

package ipsec

/**
* Configuration for IPSEC profile resource.
*/
type Ipsecprofile struct {
	/**
	* The name of the ipsec profile
	*/
	Name string `json:"name,omitempty"`
	/**
	* IKE Protocol Version
	*/
	Ikeversion string `json:"ikeversion,omitempty"`
	/**
	* Type of encryption algorithm (Note: Selection of AES enables AES128)
	*/
	Encalgo []string `json:"encalgo,omitempty"`
	/**
	* Type of hashing algorithm
	*/
	Hashalgo []string `json:"hashalgo,omitempty"`
	/**
	* Lifetime of IKE SA in seconds. Lifetime of IPSec SA will be (lifetime of IKE SA/8)
	*/
	Lifetime *int `json:"lifetime,omitempty"`
	/**
	* Pre shared key value
	*/
	Psk string `json:"psk,omitempty"`
	/**
	* Public key file path
	*/
	Publickey string `json:"publickey,omitempty"`
	/**
	* Private key file path
	*/
	Privatekey string `json:"privatekey,omitempty"`
	/**
	* Peer public key file path
	*/
	Peerpublickey string `json:"peerpublickey,omitempty"`
	/**
	* Number of seconds after which a notify payload is sent to check the liveliness of the peer. Additional retries are done as per retransmit interval setting. Zero value disables liveliness checks.
	*/
	Livenesscheckinterval *int `json:"livenesscheckinterval,omitempty"`
	/**
	* IPSec Replay window size for the data traffic
	*/
	Replaywindowsize *int `json:"replaywindowsize,omitempty"`
	/**
	* IKE retry interval for bringing up the connection
	*/
	Ikeretryinterval *int `json:"ikeretryinterval,omitempty"`
	/**
	* The interval in seconds to retry sending the IKE messages to peer, three consecutive attempts are done with doubled interval after every failure.
	*/
	Retransmissiontime *int `json:"retransmissiontime,omitempty"`
	/**
	* Enable/Disable PFS.
	*/
	Perfectforwardsecrecy string `json:"perfectforwardsecrecy,omitempty"`

	//------- Read only Parameter ---------;

	Responderonly string `json:"responderonly,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
