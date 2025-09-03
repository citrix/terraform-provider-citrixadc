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

package lb

/**
* Configuration for SIP parameters resource.
*/
type Lbsipparameters struct {
	/**
	* Port number with which to match the source port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.
	*/
	Rnatsrcport int `json:"rnatsrcport,omitempty"`
	/**
	* Port number with which to match the destination port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.
	*/
	Rnatdstport int `json:"rnatdstport,omitempty"`
	/**
	* Time, in seconds, for which a client must wait before initiating a connection after receiving a 503 Service Unavailable response from the SIP server. The time value is sent in the "Retry-After" header in the 503 response.
	*/
	Retrydur int `json:"retrydur,omitempty"`
	/**
	* Add the rport parameter to the VIA headers of SIP requests that virtual servers receive from clients or servers.
	*/
	Addrportvip string `json:"addrportvip,omitempty"`
	/**
	* Maximum number of 503 Service Unavailable responses to generate, once every 10 milliseconds, when a SIP virtual server becomes unavailable.
	*/
	Sip503ratethreshold int `json:"sip503ratethreshold,omitempty"`
	/**
	* Port number with which to match the source port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.
	*/
	Rnatsecuresrcport int `json:"rnatsecuresrcport,omitempty"`
	/**
	* Port number with which to match the destination port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.
	*/
	Rnatsecuredstport int `json:"rnatsecuredstport,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
