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


type Ping6 struct {
	/**
	* Set socket buffer size. If used, should be used with roughly +100 then the datalen (-s option). The default value is 8192.
	*/
	B *int `json:"b,omitempty"`
	/**
	* Number of packets to send. The default value is infinite. For Nitro API, defalut value is taken as 5.
	*/
	C *int `json:"c,omitempty"`
	/**
	* Waiting time, in seconds. The default value is 1 second.
	*/
	I *int `json:"i,omitempty"`
	/**
	* Network interface on which to ping, if you have multiple interfaces.
	*/
	I_ string `json:"I,omitempty"`
	/**
	* By default, ping6 asks the kernel to fragment packets to fit into the minimum IPv6 MTU.The -m option will suppress the behavior for unicast packets.
	*/
	M bool `json:"m,omitempty"`
	/**
	* Numeric output only. No name resolution.
	*/
	N bool `json:"n,omitempty"`
	/**
	* Pattern to fill in packets. Can be up to 16 bytes, useful for diagnosing data-dependent problems.
	*/
	P string `json:"p,omitempty"`
	/**
	* Quiet output. Only summary is printed. For Nitro API, this flag is set by default
	*/
	Q bool `json:"q,omitempty"`
	/**
	* Data size, in bytes. The default value is 32.
	*/
	S *int `json:"s,omitempty"`
	/**
	* VLAN ID for link local address.
	*/
	V *int `json:"V,omitempty"`
	/**
	* Source IP address to be used in the outgoing query packets.
	*/
	S_ string `json:"S,omitempty"`
	/**
	* Traffic Domain Id
	*/
	T *int `json:"T,omitempty"`
	/**
	* Timeout in seconds before ping6 exits
	*/
	T_ *int `json:"t,omitempty"`
	/**
	* Address of host to ping.
	*/
	HostName string `json:"hostName,omitempty"`

	//------- Read only Parameter ---------;

	Response string `json:"response,omitempty"`

}
