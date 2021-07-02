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

package cs

/**
* Configuration for CS policy label resource.
*/
type Cspolicylabel struct {
	/**
	* Name for the policy label. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. 
		The label name must be unique within the list of policy labels for content switching.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policylabel" or 'my policylabel').
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Protocol supported by the policy label. All policies bound to the policy label must either match the specified protocol or be a subtype of that protocol. Available settings function as follows:
		* HTTP - Supports policies that process HTTP traffic. Used to access unencrypted Web sites. (The default.)
		* SSL - Supports policies that process HTTPS/SSL encrypted traffic. Used to access encrypted Web sites.
		* TCP - Supports policies that process any type of TCP traffic, including HTTP.
		* SSL_TCP - Supports policies that process SSL-encrypted TCP traffic, including SSL.
		* UDP - Supports policies that process any type of UDP-based traffic, including DNS.
		* DNS - Supports policies that process DNS traffic.
		* ANY - Supports all types of policies except HTTP, SSL, and TCP.             
		* SIP_UDP - Supports policies that process UDP based Session Initiation Protocol (SIP) traffic. SIP initiates, manages, and terminates multimedia communications sessions, and has emerged as the standard for Internet telephony (VoIP).
		* RTSP - Supports policies that process Real Time Streaming Protocol (RTSP) traffic. RTSP provides delivery of multimedia and other streaming data, such as audio, video, and other types of streamed media.
		* RADIUS - Supports policies that process Remote Authentication Dial In User Service (RADIUS) traffic. RADIUS supports combined authentication, authorization, and auditing services for network management.
		* MYSQL - Supports policies that process MYSQL traffic.
		* MSSQL - Supports policies that process Microsoft SQL traffic.
	*/
	Cspolicylabeltype string `json:"cspolicylabeltype,omitempty"`
	/**
	* The new name of the content switching policylabel.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Numpol string `json:"numpol,omitempty"`
	Hits string `json:"hits,omitempty"`
	Policyname string `json:"policyname,omitempty"`
	Priority string `json:"priority,omitempty"`
	Targetvserver string `json:"targetvserver,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Invokelabelname string `json:"invoke_labelname,omitempty"`

}
