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

package aaa

/**
* Configuration for RADIUS parameter resource.
*/
type Aaaradiusparams struct {
	/**
	* IP address of your RADIUS server.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* Port number on which the RADIUS server listens for connections.
	*/
	Serverport int `json:"serverport,omitempty"`
	/**
	* Maximum number of seconds that the Citrix ADC waits for a response from the RADIUS server.
	*/
	Authtimeout int `json:"authtimeout,omitempty"`
	/**
	* The key shared between the RADIUS server and clients.
		Required for allowing the Citrix ADC to communicate with the RADIUS server.
	*/
	Radkey string `json:"radkey,omitempty"`
	/**
	* Send the Citrix ADC IP (NSIP) address to the RADIUS server as the Network Access Server IP (NASIP) part of the Radius protocol.
	*/
	Radnasip string `json:"radnasip,omitempty"`
	/**
	* Send the Network Access Server ID (NASID) for your Citrix ADC to the RADIUS server as the nasid part of the Radius protocol.
	*/
	Radnasid string `json:"radnasid,omitempty"`
	/**
	* Vendor ID for RADIUS group extraction.
	*/
	Radvendorid int `json:"radvendorid,omitempty"`
	/**
	* Attribute type for RADIUS group extraction.
	*/
	Radattributetype int `json:"radattributetype,omitempty"`
	/**
	* Prefix string that precedes group names within a RADIUS attribute for RADIUS group extraction.
	*/
	Radgroupsprefix string `json:"radgroupsprefix,omitempty"`
	/**
	* Group separator string that delimits group names within a RADIUS attribute for RADIUS group extraction.
	*/
	Radgroupseparator string `json:"radgroupseparator,omitempty"`
	/**
	* Enable password encoding in RADIUS packets that the Citrix ADC sends to the RADIUS server.
	*/
	Passencoding string `json:"passencoding,omitempty"`
	/**
	* Vendor ID attribute in the RADIUS response.
		If the attribute is not vendor-encoded, it is set to 0.
	*/
	Ipvendorid int `json:"ipvendorid,omitempty"`
	/**
	* IP attribute type in the RADIUS response.
	*/
	Ipattributetype int `json:"ipattributetype,omitempty"`
	/**
	* Configure the RADIUS server state to accept or refuse accounting messages.
	*/
	Accounting string `json:"accounting,omitempty"`
	/**
	* Vendor ID of the password in the RADIUS response. Used to extract the user password.
	*/
	Pwdvendorid int `json:"pwdvendorid,omitempty"`
	/**
	* Attribute type of the Vendor ID in the RADIUS response.
	*/
	Pwdattributetype int `json:"pwdattributetype,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	/**
	* Send Calling-Station-ID of the client to the RADIUS server. IP Address of the client is sent as its Calling-Station-ID.
	*/
	Callingstationid string `json:"callingstationid,omitempty"`
	/**
	* Number of retry by the Citrix ADC before getting response from the RADIUS server.
	*/
	Authservretry int `json:"authservretry,omitempty"`
	/**
	* Configure the RADIUS server state to accept or refuse authentication messages.
	*/
	Authentication string `json:"authentication,omitempty"`
	/**
	* Send Tunnel Endpoint Client IP address to the RADIUS server.
	*/
	Tunnelendpointclientip string `json:"tunnelendpointclientip,omitempty"`
	/**
	* Control whether the Message-Authenticator attribute is included in a RADIUS Access-Request packet.
	*/
	Messageauthenticator string `json:"messageauthenticator,omitempty"`

	//------- Read only Parameter ---------;

	Groupauthname string `json:"groupauthname,omitempty"`
	Ipaddress string `json:"ipaddress,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
