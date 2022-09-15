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

package rdp

/**
* Configuration for RDP serverprofile resource.
*/
type Rdpserverprofile struct {
	/**
	* The name of the rdp server profile
	*/
	Name string `json:"name,omitempty"`
	/**
	* IPv4 or IPv6 address of RDP listener. This terminates client RDP connections.
	*/
	Rdpip string `json:"rdpip,omitempty"`
	/**
	* TCP port on which the RDP connection is established.
	*/
	Rdpport int `json:"rdpport,omitempty"`
	/**
	* Pre shared key value
	*/
	Psk string `json:"psk,omitempty"`
	/**
	* Enable/Disable RDP redirection support. This needs to be enabled in presence of connection broker or session directory with IP cookie(msts cookie) based redirection support
	*/
	Rdpredirection string `json:"rdpredirection,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
