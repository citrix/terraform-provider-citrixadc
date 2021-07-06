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

package ns

/**
* Configuration for ICAP profile resource.
*/
type Nsicapprofile struct {
	/**
	* Name for an ICAP profile. Must begin with a letter, number, or the underscore \(_\) character. Other characters allowed, after the first character, are the hyphen \(-\), period \(.\), hash \(\#\), space \( \), at \(@\), colon \(:\), and equal \(=\) characters. The name of a ICAP profile cannot be changed after it is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my icap profile" or 'my icap profile'\).
	*/
	Name string `json:"name,omitempty"`
	/**
	* Enable or Disable preview header with ICAP request. This feature allows an ICAP server to see the beginning of a transaction, then decide if it wants to opt-out of the transaction early instead of receiving the remainder of the request message.
	*/
	Preview string `json:"preview,omitempty"`
	/**
	* Value of Preview Header field. Citrix ADC uses the minimum of this set value and the preview size received on OPTIONS response.
	*/
	Previewlength int `json:"previewlength,omitempty"`
	/**
	* URI representing icap service. It is a mandatory argument while creating an icapprofile.
	*/
	Uri string `json:"uri,omitempty"`
	/**
	* ICAP Host Header
	*/
	Hostheader string `json:"hostheader,omitempty"`
	/**
	* ICAP User Agent Header String
	*/
	Useragent string `json:"useragent,omitempty"`
	/**
	* ICAP Mode of operation. It is a mandatory argument while creating an icapprofile.
	*/
	Mode string `json:"mode,omitempty"`
	/**
	* Query parameters to be included with ICAP request URI. Entered values should be in arg=value format. For more than one parameters, add & separated values. e.g.: arg1=val1&arg2=val2.
	*/
	Queryparams string `json:"queryparams,omitempty"`
	/**
	* If enabled, Citrix ADC keeps the ICAP connection alive after a transaction to reuse it to send next ICAP request.
	*/
	Connectionkeepalive string `json:"connectionkeepalive,omitempty"`
	/**
	* Enable or Disable sending Allow: 204 header in ICAP request.
	*/
	Allow204 string `json:"allow204,omitempty"`
	/**
	* Insert custom ICAP headers in the ICAP request to send to ICAP server. The headers can be static or can be dynamically constructed using PI Policy Expression. For example, to send static user agent and Client's IP address, the expression can be specified as "User-Agent: NS-ICAP-Client/V1.0\r\nX-Client-IP: "+CLIENT.IP.SRC+"\r\n".
		The Citrix ADC does not check the validity of the specified header name-value. You must manually validate the specified header syntax.
	*/
	Inserticapheaders string `json:"inserticapheaders,omitempty"`
	/**
	* Exact HTTP request, in the form of an expression, which the Citrix ADC encapsulates and sends to the ICAP server. If you set this parameter, the ICAP request is sent using only this header. This can be used when the HTTP header is not available to send or ICAP server only needs part of the incoming HTTP request. The request expression is constrained by the feature for which it is used.
		The Citrix ADC does not check the validity of this request. You must manually validate the request.
	*/
	Inserthttprequest string `json:"inserthttprequest,omitempty"`
	/**
	* Time, in seconds, within which the remote server should respond to the ICAP-request. If the Netscaler does not receive full response with this time, the specified request timeout action is performed. Zero value disables this timeout functionality.
	*/
	Reqtimeout int `json:"reqtimeout,omitempty"`
	/**
	* Name of the action to perform if the Vserver/Server representing the remote service does not respond with any response within the timeout value configured. The Supported actions are
		* BYPASS - This Ignores the remote server response and sends the request/response to Client/Server.
		* If the ICAP response with Encapsulated headers is not received within the request-timeout value configured, this Ignores the remote ICAP server response and sends the Full request/response to Server/Client.
		* RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.
		* DROP - Drop the request without sending a response to the user.
	*/
	Reqtimeoutaction string `json:"reqtimeoutaction,omitempty"`
	/**
	* Name of the audit message action which would be evaluated on receiving the ICAP response to emit the logs.
	*/
	Logaction string `json:"logaction,omitempty"`

}
