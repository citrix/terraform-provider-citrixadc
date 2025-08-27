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

package lb

/**
* Configuration for lb policy label resource.
*/
type Lbpolicylabel struct {
	/**
	* Name for the LB policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my lb policy label" or 'my lb policy label').
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Protocols supported by the policylabel. Available Types are :
		* HTTP - HTTP requests.
		* DNS - DNS request.
		* OTHERTCP - OTHERTCP request.
		* SIP_UDP - SIP_UDP request.
		* SIP_TCP - SIP_TCP request.
		* MYSQL - MYSQL request.
		* MSSQL - MSSQL request.
		* ORACLE - ORACLE request.
		* NAT - NAT request.
		* DIAMETER - DIAMETER request.
		* RADIUS - RADIUS request.
		* MQTT - MQTT request.
		* QUIC_BRIDGE - QUIC_BRIDGE request.
		* HTTP_QUIC - HTTP_QUIC request.
	*/
	Policylabeltype string `json:"policylabeltype,omitempty"`
	/**
	* Any comments to preserve information about this LB policy label.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the LB policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Numpol string `json:"numpol,omitempty"`
	Hits string `json:"hits,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Invokelabelname string `json:"invoke_labelname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
