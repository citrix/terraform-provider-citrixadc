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

package authentication

/**
* Configuration for RADIUS action resource.
*/
type Authenticationradiusaction struct {
	/**
	* Name for the RADIUS action. 
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the RADIUS action is added.
	*/
	Name string `json:"name,omitempty"`
	/**
	* IP address assigned to the RADIUS server.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* RADIUS server name as a FQDN.  Mutually exclusive with RADIUS IP address.
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* Port number on which the RADIUS server listens for connections.
	*/
	Serverport int `json:"serverport,omitempty"`
	/**
	* Number of seconds the Citrix ADC waits for a response from the RADIUS server.
	*/
	Authtimeout int `json:"authtimeout,omitempty"`
	/**
	* Key shared between the RADIUS server and the Citrix ADC. 
		Required to allow the Citrix ADC to communicate with the RADIUS server.
	*/
	Radkey string `json:"radkey,omitempty"`
	/**
	* If enabled, the Citrix ADC IP address (NSIP) is sent to the RADIUS server as the  Network Access Server IP (NASIP) address. 
		The RADIUS protocol defines the meaning and use of the NASIP address.
	*/
	Radnasip string `json:"radnasip,omitempty"`
	/**
	* If configured, this string is sent to the RADIUS server as the Network Access Server ID (NASID).
	*/
	Radnasid string `json:"radnasid,omitempty"`
	/**
	* RADIUS vendor ID attribute, used for RADIUS group extraction.
	*/
	Radvendorid int `json:"radvendorid,omitempty"`
	/**
	* RADIUS attribute type, used for RADIUS group extraction.
	*/
	Radattributetype int `json:"radattributetype,omitempty"`
	/**
	* RADIUS groups prefix string. 
		This groups prefix precedes the group names within a RADIUS attribute for RADIUS group extraction.
	*/
	Radgroupsprefix string `json:"radgroupsprefix,omitempty"`
	/**
	* RADIUS group separator string
		The group separator delimits group names within a RADIUS attribute for RADIUS group extraction.
	*/
	Radgroupseparator string `json:"radgroupseparator,omitempty"`
	/**
	* Encoding type for passwords in RADIUS packets that the Citrix ADC sends to the RADIUS server.
	*/
	Passencoding string `json:"passencoding,omitempty"`
	/**
	* Vendor ID of the intranet IP attribute in the RADIUS response.
		NOTE: A value of 0 indicates that the attribute is not vendor encoded.
	*/
	Ipvendorid int `json:"ipvendorid,omitempty"`
	/**
	* Remote IP address attribute type in a RADIUS response.
	*/
	Ipattributetype int `json:"ipattributetype,omitempty"`
	/**
	* Whether the RADIUS server is currently accepting accounting messages.
	*/
	Accounting string `json:"accounting,omitempty"`
	/**
	* Vendor ID of the attribute, in the RADIUS response, used to extract the user password.
	*/
	Pwdvendorid int `json:"pwdvendorid,omitempty"`
	/**
	* Vendor-specific password attribute type in a RADIUS response.
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

	//------- Read only Parameter ---------;

	Ipaddress string `json:"ipaddress,omitempty"`
	Success string `json:"success,omitempty"`
	Failure string `json:"failure,omitempty"`

}
