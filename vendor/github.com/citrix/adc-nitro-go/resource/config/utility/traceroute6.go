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


type Traceroute6 struct {
	/**
	* Print hop addresses numerically rather than symbolically and numerically.
	*/
	N bool `json:"n,omitempty"`
	/**
	* Use ICMP ECHO for probes.
	*/
	I bool `json:"I,omitempty"`
	/**
	* Bypass normal routing tables and send directly to a host on an attached network. If the host is not on a directly attached network, an error is returned.
	*/
	R bool `json:"r,omitempty"`
	/**
	* Verbose output. List received ICMP packets other than TIME_EXCEEDED and UNREACHABLE.
	*/
	V bool `json:"v,omitempty"`
	/**
	* Maximum hop value for outgoing probe packets. For Nitro API, default value is taken as 10.
	*/
	M int `json:"m,omitempty"`
	/**
	* Base port number used in probes.
	*/
	P int `json:"p,omitempty"`
	/**
	* Number of probes per hop. For Nitro API, default value is taken as 1.
	*/
	Q int `json:"q,omitempty"`
	/**
	* Source IP address to use in the outgoing query packets. If the IP address does not belong to this appliance,  an error is returned and nothing is sent.
	*/
	S string `json:"s,omitempty"`
	/**
	* Traffic Domain Id
	*/
	T int `json:"T,omitempty"`
	/**
	* Time (in seconds) to wait for a response to a query. For Nitro API, defalut value is set to 3.
	*/
	W int `json:"w,omitempty"`
	/**
	* Destination host IP address or name.
	*/
	Host string `json:"host,omitempty"`
	/**
	* Length (in bytes) of the query packets.
	*/
	Packetlen int `json:"packetlen,omitempty"`

	//------- Read only Parameter ---------;

	Response string `json:"response,omitempty"`

}
