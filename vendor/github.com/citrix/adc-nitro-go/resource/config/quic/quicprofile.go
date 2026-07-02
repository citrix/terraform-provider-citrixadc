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
	* Name for the QUIC profile.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Maximum idle timeout, in seconds, for a QUIC connection.
	*/
	Maxidletimeout *int `json:"maxidletimeout,omitempty"`
	/**
	* Size of the largest UDP datagram payload, in bytes, that the Citrix ADC is willing to receive.
	*/
	Maxudppayloadsize *int `json:"maxudppayloadsize,omitempty"`
	/**
	* Initial value, in bytes, for the maximum amount of data that can be sent on a QUIC connection.
	*/
	Initialmaxdata *int `json:"initialmaxdata,omitempty"`
	/**
	* Initial flow control limit, in bytes, for locally-initiated bidirectional QUIC streams.
	*/
	Initialmaxstreamdatabidilocal *int `json:"initialmaxstreamdatabidilocal,omitempty"`
	/**
	* Initial flow control limit, in bytes, for remotely-initiated bidirectional QUIC streams.
	*/
	Initialmaxstreamdatabidiremote *int `json:"initialmaxstreamdatabidiremote,omitempty"`
	/**
	* Initial flow control limit, in bytes, for unidirectional QUIC streams.
	*/
	Initialmaxstreamdatauni *int `json:"initialmaxstreamdatauni,omitempty"`
	/**
	* Initial maximum number of bidirectional streams the remote QUIC endpoint may initiate.
	*/
	Initialmaxstreamsbidi *int `json:"initialmaxstreamsbidi,omitempty"`
	/**
	* Initial maximum number of unidirectional streams the remote QUIC endpoint may initiate.
	*/
	Initialmaxstreamsuni *int `json:"initialmaxstreamsuni,omitempty"`
	/**
	* Exponent used to decode the ACK Delay field in QUIC ACK frames.
	*/
	Ackdelayexponent *int `json:"ackdelayexponent,omitempty"`
	/**
	* Maximum amount of time, in milliseconds, by which the Citrix ADC will delay sending acknowledgments.
	*/
	Maxackdelay *int `json:"maxackdelay,omitempty"`
	/**
	* Maximum number of QUIC connection IDs from the remote QUIC endpoint that the Citrix ADC is willing to store.
	*/
	Activeconnectionidlimit *int `json:"activeconnectionidlimit,omitempty"`
	/**
	* Whether the Citrix ADC should allow the remote QUIC endpoint to perform active QUIC connection migration.
	*/
	Activeconnectionmigration string `json:"activeconnectionmigration,omitempty"`
	/**
	* Congestion control algorithm to be used for QUIC connections.
	*/
	Congestionctrlalgorithm string `json:"congestionctrlalgorithm,omitempty"`
	/**
	* Maximum number of UDP datagrams that can be transmitted in a single transmission burst.
	*/
	Maxudpdatagramsperburst *int `json:"maxudpdatagramsperburst,omitempty"`
	/**
	* Whether the Citrix ADC should perform stateless address validation for QUIC clients.
	*/
	Statelessaddressvalidation string `json:"statelessaddressvalidation,omitempty"`
	/**
	* Validity period, in seconds, of address validation tokens issued through QUIC NEW_TOKEN frames.
	*/
	Newtokenvalidityperiod *int `json:"newtokenvalidityperiod,omitempty"`
	/**
	* Validity period, in seconds, of address validation tokens issued through QUIC Retry packets.
	*/
	Retrytokenvalidityperiod *int `json:"retrytokenvalidityperiod,omitempty"`
	//------- Read only Parameter ---------;
	Refcnt             string   `json:"refcnt,omitempty"`
	Builtin            []string `json:"builtin,omitempty"`
	Feature            string   `json:"feature,omitempty"`
	Nextgenapiresource string   `json:"_nextgenapiresource,omitempty"`
}
