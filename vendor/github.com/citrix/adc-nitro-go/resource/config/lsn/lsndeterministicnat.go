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
* Configuration for deterministic NAT resource.
*/
type Lsndeterministicnat struct {
	/**
	* The name of the LSN Client.
	*/
	Clientname string `json:"clientname,omitempty"`
	/**
	* IPv6 address of the LSN subscriber or B4 device.
	*/
	Network6 string `json:"network6,omitempty"`
	/**
	* The Client IP address.
	*/
	Subscrip string `json:"subscrip,omitempty"`
	/**
	* The LSN client TD.
	*/
	Td int `json:"td,omitempty"`
	/**
	* The NAT IP address.
	*/
	Natip string `json:"natip,omitempty"`

	//------- Read only Parameter ---------;

	Natprefix string `json:"natprefix,omitempty"`
	Subscrip2 string `json:"subscrip2,omitempty"`
	Natip2 string `json:"natip2,omitempty"`
	Firstport string `json:"firstport,omitempty"`
	Lastport string `json:"lastport,omitempty"`
	Srctd string `json:"srctd,omitempty"`
	Nattype string `json:"nattype,omitempty"`

}
