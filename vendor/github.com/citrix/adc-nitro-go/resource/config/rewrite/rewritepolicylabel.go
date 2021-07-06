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

package rewrite

/**
* Configuration for rewrite policy label resource.
*/
type Rewritepolicylabel struct {
	/**
	* Name for the rewrite policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rewrite policy label is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my rewrite policy label" or 'my rewrite policy label').
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Types of transformations allowed by the policies bound to the label. For Rewrite, the following types are supported:
		* http_req - HTTP requests
		* http_res - HTTP responses
		* othertcp_req - Non-HTTP TCP requests
		* othertcp_res - Non-HTTP TCP responses
		* url - URLs
		* text - Text strings
		* clientless_vpn_req - Citrix ADC clientless VPN requests
		* clientless_vpn_res - Citrix ADC clientless VPN responses
		* sipudp_req - SIP requests
		* sipudp_res - SIP responses
		* diameter_req - DIAMETER requests
		* diameter_res - DIAMETER responses
		* radius_req - RADIUS requests
		* radius_res - RADIUS responses
		* dns_req - DNS requests
		* dns_res - DNS responses
	*/
	Transform string `json:"transform,omitempty"`
	/**
	* Any comments to preserve information about this rewrite policy label.
	*/
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the rewrite policy label. 
		Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy label" or 'my policy label').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Numpol string `json:"numpol,omitempty"`
	Hits string `json:"hits,omitempty"`
	Priority string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Invokelabelname string `json:"invoke_labelname,omitempty"`
	Flowtype string `json:"flowtype,omitempty"`
	Description string `json:"description,omitempty"`
	Isdefault string `json:"isdefault,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`

}
