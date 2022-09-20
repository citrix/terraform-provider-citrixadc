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

package smpp

/**
* Configuration for SMPP configuration parameters resource.
*/
type Smppparam struct {
	/**
	* Mode in which the client binds to the ADC. Applicable settings function as follows:
		* TRANSCEIVER - Client can send and receive messages to and from the message center.
		* TRANSMITTERONLY - Client can only send messages.
		* RECEIVERONLY - Client can only receive messages.
	*/
	Clientmode string `json:"clientmode,omitempty"`
	/**
	* Queue SMPP messages if a client that is capable of receiving the destination address messages is not available.
	*/
	Msgqueue string `json:"msgqueue,omitempty"`
	/**
	* Maximum number of SMPP messages that can be queued. After the limit is reached, the Citrix ADC sends a deliver_sm_resp PDU, with an appropriate error message, to the message center.
	*/
	Msgqueuesize int `json:"msgqueuesize,omitempty"`
	/**
	* Type of Number, such as an international number or a national number, used in the ESME address sent in the bind request.
	*/
	Addrton int `json:"addrton,omitempty"`
	/**
	* Numbering Plan Indicator, such as landline, data, or WAP client, used in the ESME address sent in the bind request.
	*/
	Addrnpi int `json:"addrnpi,omitempty"`
	/**
	* Set of SME addresses, sent in the bind request, serviced by the ESME.
	*/
	Addrrange string `json:"addrrange,omitempty"`

}
