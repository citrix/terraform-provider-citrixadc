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

package lsn

/**
* Configuration for lsn session resource.
*/
type Lsnsession struct {
	/**
	* Type of sessions to be displayed.
	*/
	Nattype string `json:"nattype,omitempty"`
	/**
	* Name of the LSN Client entity.
	*/
	Clientname string `json:"clientname,omitempty"`
	/**
	* IP address or network address of subscriber(s).
	*/
	Network string `json:"network,omitempty"`
	/**
	* Subnet mask for the IP address specified by the network parameter.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* IPv6 address of the LSN subscriber or B4 device.
	*/
	Network6 string `json:"network6,omitempty"`
	/**
	* Traffic domain ID of the LSN client entity.
	*/
	Td int `json:"td,omitempty"`
	/**
	* Mapped NAT IP address used in LSN sessions.
	*/
	Natip string `json:"natip,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`
	/**
	* Mapped NAT port used in the LSN sessions.
	*/
	Natport2 int `json:"natport2,omitempty"`

	//------- Read only Parameter ---------;

	Natprefix string `json:"natprefix,omitempty"`
	Subscrip string `json:"subscrip,omitempty"`
	Subscrport string `json:"subscrport,omitempty"`
	Destip string `json:"destip,omitempty"`
	Destport string `json:"destport,omitempty"`
	Natport string `json:"natport,omitempty"`
	Transportprotocol string `json:"transportprotocol,omitempty"`
	Sessionestdir string `json:"sessionestdir,omitempty"`
	Dsttd string `json:"dsttd,omitempty"`
	Srctd string `json:"srctd,omitempty"`
	Ipv6address string `json:"ipv6address,omitempty"`

}
