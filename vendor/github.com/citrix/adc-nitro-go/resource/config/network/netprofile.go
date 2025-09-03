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

package network

/**
* Configuration for Network profile resource.
*/
type Netprofile struct {
	/**
	* Name for the net profile. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created. Choose a name that helps identify the net profile.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td int `json:"td,omitempty"`
	/**
	* IP address or the name of an IP set.
	*/
	Srcip string `json:"srcip,omitempty"`
	/**
	* When the net profile is associated with a virtual server or its bound services, this option enables the Citrix ADC to use the same  address, specified in the net profile, to communicate to servers for all sessions initiated from a particular client to the virtual server.
	*/
	Srcippersistency string `json:"srcippersistency,omitempty"`
	/**
	* USNIP/USIP settings override LSN settings for configured
		service/virtual server traffic.. 
	*/
	Overridelsn string `json:"overridelsn,omitempty"`
	/**
	* Response will be sent using learnt info if enabled. When creating a netprofile, if you do not set this parameter, the netprofile inherits the global MBF setting (available in the enable ns mode and disable ns mode CLI commands, or in the System > Settings > Configure modes > Configure Modes dialog box). However, you can override this setting after you create the netprofile
	*/
	Mbf string `json:"mbf,omitempty"`
	/**
	* Proxy Protocol Action (Enabled/Disabled)
	*/
	Proxyprotocol string `json:"proxyprotocol,omitempty"`
	/**
	* Proxy Protocol Version (V1/V2)
	*/
	Proxyprotocoltxversion string `json:"proxyprotocoltxversion,omitempty"`
	/**
	* ADC doesnt look for proxy header before TLS handshake, if enabled. Proxy protocol parsed after TLS handshake
	*/
	Proxyprotocolaftertlshandshake string `json:"proxyprotocolaftertlshandshake,omitempty"`

	//------- Read only Parameter ---------;

	Proxyprotocoltlvoptions string `json:"proxyprotocoltlvoptions,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
