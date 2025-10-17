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

package appflow

/**
* Configuration for AppFlow collector resource.
*/
type Appflowcollector struct {
	/**
	* Name for the collector. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at
		(@), equals (=), and hyphen (-) characters.
		Only four collectors can be configured.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow collector" or 'my appflow collector').
	*/
	Name string `json:"name,omitempty"`
	/**
	* IPv4 address of the collector.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Port on which the collector listens.
	*/
	Port *int `json:"port,omitempty"`
	/**
	* Netprofile to associate with the collector. The IP address defined in the profile is used as the source IP address for AppFlow traffic for this collector.  If you do not set this parameter, the Citrix ADC IP (NSIP) address is used as the source IP address.
	*/
	Netprofile string `json:"netprofile,omitempty"`
	/**
	* Type of collector: either logstream or ipfix or rest.
	*/
	Transport string `json:"transport,omitempty"`
	/**
	* New name for the collector. Must begin with an ASCII alphabetic or underscore (_) character, and must
		contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at(@), equals (=), and hyphen (-) characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my appflow coll" or 'my appflow coll').
	*/
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	State string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
