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
* Configuration for tacacs parameters resource.
*/
type Aaatacacsparams struct {
	/**
	* IP address of your TACACS+ server.
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* Port number on which the TACACS+ server listens for connections.
	*/
	Serverport *int `json:"serverport,omitempty"`
	/**
	* Maximum number of seconds that the Citrix ADC waits for a response from the TACACS+ server.
	*/
	Authtimeout *int `json:"authtimeout,omitempty"`
	/**
	* Key shared between the TACACS+ server and clients. Required for allowing the Citrix ADC to communicate with the TACACS+ server.
	*/
	Tacacssecret string `json:"tacacssecret,omitempty"`
	/**
	* Use streaming authorization on the TACACS+ server.
	*/
	Authorization string `json:"authorization,omitempty"`
	/**
	* Send accounting messages to the TACACS+ server.
	*/
	Accounting string `json:"accounting,omitempty"`
	/**
	* The option for sending accounting messages to the TACACS+ server.
	*/
	Auditfailedcmds string `json:"auditfailedcmds,omitempty"`
	/**
	* TACACS+ group attribute name.Used for group extraction on the TACACS+ server.
	*/
	Groupattrname string `json:"groupattrname,omitempty"`
	/**
	* This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
