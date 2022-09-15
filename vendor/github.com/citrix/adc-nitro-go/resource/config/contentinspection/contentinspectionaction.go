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

package contentinspection

/**
* Configuration for Content Inspection action resource.
*/
type Contentinspectionaction struct {
	/**
	* Name of the remote service action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of operation this action is going to perform. following actions are available to configure:
		* ICAP - forward the incoming request or response to an ICAP server for modification.
		* INLINEINSPECTION - forward the incoming or outgoing packets to IPS server for Intrusion Prevention.
		* MIRROR - Forwards cloned packets for Intrusion Detection.
		* NOINSPECTION - This does not forward incoming and outgoing packets to the Inspection device.
		* NSTRACE - capture current and further incoming packets on this transaction.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Name of the LB vserver or service
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* IP address of remoteService
	*/
	Serverip string `json:"serverip,omitempty"`
	/**
	* Port of remoteService
	*/
	Serverport int `json:"serverport,omitempty"`
	/**
	* Name of the ICAP profile to be attached to the contentInspection action.
	*/
	Icapprofilename string `json:"icapprofilename,omitempty"`
	/**
	* Name of the action to perform if the Vserver representing the remote service is not UP. This is not supported for NOINSPECTION Type. The Supported actions are:
		* RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.
		* DROP - Drop the request without sending a response to the user.
		* CONTINUE - It bypasses the ContentIsnpection and Continues/resumes the Traffic-Flow to Client/Server.
	*/
	Ifserverdown string `json:"ifserverdown,omitempty"`

	//------- Read only Parameter ---------;

	Reqtimeout string `json:"reqtimeout,omitempty"`
	Reqtimeoutaction string `json:"reqtimeoutaction,omitempty"`
	Hits string `json:"hits,omitempty"`
	Referencecount string `json:"referencecount,omitempty"`
	Undefhits string `json:"undefhits,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
