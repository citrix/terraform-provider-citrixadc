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

package utility


type Traceroute struct {
	/**
	* Print a summary of how many probes were not answered for each hop.
	*/
	S bool `json:"S,omitempty"`
	/**
	* Print hop addresses numerically instead of symbolically and numerically.
	*/
	N bool `json:"n,omitempty"`
	/**
	* Bypass normal routing tables and send directly to a host on an attached network. If the host is not on a directly attached network, an error is returned.
	*/
	R bool `json:"r,omitempty"`
	/**
	* Verbose output. List received ICMP packets other than TIME_EXCEEDED and UNREACHABLE.
	*/
	V bool `json:"v,omitempty"`
	/**
	* Minimum TTL value used in outgoing probe packets.
	*/
	M uint32 `json:"M,omitempty"`
	/**
	* Maximum TTL value used in outgoing probe packets. For Nitro API, default value is taken as 10.
	*/
	M_ int32 `json:"m,omitempty"`
	/**
	* Send packets of specified IP protocol. The currently supported protocols are UDP and ICMP.
	*/
	P string `json:"P,omitempty"`
	/**
	* Base port number used in probes.
	*/
	P_ int32 `json:"p,omitempty"`
	/**
	* Number of queries per hop. For Nitro API, defalut value is taken as 1.
	*/
	Q int32 `json:"q,omitempty"`
	/**
	* Source IP address to use in the outgoing query packets. If the IP address does not belong to this appliance,  an error is returned and nothing is sent.
	*/
	S_ string `json:"s,omitempty"`
	/**
	* Traffic Domain Id
	*/
	T uint32 `json:"T,omitempty"`
	/**
	* Type-of-service in query packets.
	*/
	T_ int32 `json:"t,omitempty"`
	/**
	* Time (in seconds) to wait for a response to a query. For Nitro API, defalut value is set to 3.
	*/
	W int32 `json:"w,omitempty"`
	/**
	* Destination host IP address or name.
	*/
	Host string `json:"host,omitempty"`
	/**
	* Length (in bytes) of the query packets.
	*/
	Packetlen int32 `json:"packetlen,omitempty"`

	//------- Read only Parameter ---------;

	Response string `json:"response,omitempty"`

}
