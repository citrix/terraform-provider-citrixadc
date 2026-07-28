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

package cloudtunnel

/**
* Configuration for Cloud Tunnel virtual server resource.
*/
type Cloudtunnelvserver struct {
	/**
	* Name for the Cloud Tunnel virtual server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space,colon (:), at (@), equals (=), and hyphen (-) characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example,
		"my server" or 'my server').
	*/
	Name string `json:"name,omitempty"`
	/**
	* ServiceType of Listener using which traffic will be tunneled through cloud tunnel server.
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* String specifying the listen policy for the Cloud Tunnel virtual server. Can be either a named expression or an expression. The Cloud Tunnel virtual server processes only the traffic for which the expression evaluates to true.
	*/
	Listenpolicy string `json:"listenpolicy,omitempty"`
	/**
	* Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server, the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request.
	*/
	Listenpriority *int `json:"listenpriority,omitempty"`

	//------- Read only Parameter ---------;

	State string `json:"state,omitempty"`
	Effectivestate string `json:"effectivestate,omitempty"`
	Type string `json:"type,omitempty"`
	Ip string `json:"ip,omitempty"`
	Ipv46 string `json:"ipv46,omitempty"`
	Ippattern string `json:"ippattern,omitempty"`
	Port string `json:"port,omitempty"`
	Range string `json:"range,omitempty"`
	Cachetype string `json:"cachetype,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
