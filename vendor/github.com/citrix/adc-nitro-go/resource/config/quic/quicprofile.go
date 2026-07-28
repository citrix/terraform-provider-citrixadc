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

package quic

/**
* Configuration for QUIC profile resource.
*/
type Quicprofile struct {
	/**
	* Name for the QUIC profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.
	*/
	Name string `json:"name,omitempty"`
	/**
	* An integer value advertised by the Citrix ADC to the remote QUIC endpoint, indicating an exponent that the remote QUIC endpoint should use, to decode the ACK Delay field in QUIC ACK frames sent by the Citrix ADC.
	*/
	Ackdelayexponent *int `json:"ackdelayexponent,omitempty"`
	/**
	* An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum number of QUIC connection IDs from the remote QUIC endpoint, that the Citrix ADC is willing to store.
	*/
	Activeconnectionidlimit *int `json:"activeconnectionidlimit,omitempty"`
	/**
	* Specify whether the Citrix ADC should allow the remote QUIC endpoint to perform active QUIC connection migration.
	*/
	Activeconnectionmigration string `json:"activeconnectionmigration,omitempty"`
	/**
	* Specify the congestion control algorithm to be used for QUIC connections. The default congestion control algorithm is CUBIC.
	*/
	Congestionctrlalgorithm string `json:"congestionctrlalgorithm,omitempty"`
	/**
	* An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial value, in bytes, for the maximum amount of data that can be sent on a QUIC connection.
	*/
	Initialmaxdata *int `json:"initialmaxdata,omitempty"`
	/**
	* An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the Citrix ADC.
	*/
	Initialmaxstreamdatabidilocal *int `json:"initialmaxstreamdatabidilocal,omitempty"`
	/**
	* An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the remote QUIC endpoint.
	*/
	Initialmaxstreamdatabidiremote *int `json:"initialmaxstreamdatabidiremote,omitempty"`
	/**
	* An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for unidirectional streams initiated by the remote QUIC endpoint.
	*/
	Initialmaxstreamdatauni *int `json:"initialmaxstreamdatauni,omitempty"`
	/**
	* An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial maximum number of bidirectional streams the remote QUIC endpoint may initiate.
	*/
	Initialmaxstreamsbidi *int `json:"initialmaxstreamsbidi,omitempty"`
	/**
	* An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial maximum number of unidirectional streams the remote QUIC endpoint may initiate.
	*/
	Initialmaxstreamsuni *int `json:"initialmaxstreamsuni,omitempty"`
	/**
	* An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum amount of time, in milliseconds, by which the Citrix ADC will delay sending acknowledgments.
	*/
	Maxackdelay *int `json:"maxackdelay,omitempty"`
	/**
	* An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum idle timeout, in seconds, for a QUIC connection. A QUIC connection will be silently discarded by the Citrix ADC if it remains idle for longer than the minimum of the idle timeout values advertised by the Citrix ADC and the remote QUIC endpoint, and three times the current Probe Timeout (PTO).
	*/
	Maxidletimeout *int `json:"maxidletimeout,omitempty"`
	/**
	* An integer value, specifying the maximum number of UDP datagrams that can be transmitted by the Citrix ADC in a single transmission burst on a QUIC connection.
	*/
	Maxudpdatagramsperburst *int `json:"maxudpdatagramsperburst,omitempty"`
	/**
	* An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the size of the largest UDP datagram payload, in bytes, that the Citrix ADC is willing to receive on a QUIC connection.
	*/
	Maxudppayloadsize *int `json:"maxudppayloadsize,omitempty"`
	/**
	* An integer value, specifying the validity period, in seconds, of address validation tokens issued through QUIC NEW_TOKEN frames sent by the Citrix ADC.
	*/
	Newtokenvalidityperiod *int `json:"newtokenvalidityperiod,omitempty"`
	/**
	* An integer value, specifying the validity period, in seconds, of address validation tokens issued through QUIC Retry packets sent by the Citrix ADC.
	*/
	Retrytokenvalidityperiod *int `json:"retrytokenvalidityperiod,omitempty"`
	/**
	* Specify whether the Citrix ADC should perform stateless address validation for QUIC clients, by sending tokens in QUIC Retry packets during QUIC connection establishment, and by sending tokens in QUIC NEW_TOKEN frames after QUIC connection establishment.
	*/
	Statelessaddressvalidation string `json:"statelessaddressvalidation,omitempty"`

	//------- Read only Parameter ---------;

	Refcnt string `json:"refcnt,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
