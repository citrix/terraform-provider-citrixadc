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
* Configuration for ContentInspection profile resource.
*/
type Contentinspectionprofile struct {
	/**
	* Name of a ContentInspection profile. Must begin with a letter, number, or the underscore \(_\) character. Other characters allowed, after the first character, are the hyphen \(-\), period \(.\), hash \(\#\), space \( \), at \(@\), colon \(:\), and equal \(=\) characters. The name of a IPS profile cannot be changed after it is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my ips profile" or 'my ips profile'\).
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of ContentInspection profile. Following types are available to configure:
		INLINEINSPECTION : To inspect the packets/requests using IPS.
		MIRROR : To forward cloned packets.
	*/
	Type string `json:"type,omitempty"`
	/**
	* Ingress interface for CI profile.It is a mandatory argument while creating an ContentInspection profile of IPS type.
	*/
	Ingressinterface string `json:"ingressinterface,omitempty"`
	/**
	* Ingress Vlan for CI
	*/
	Ingressvlan *int `json:"ingressvlan,omitempty"`
	/**
	* Egress interface for CI profile.It is a mandatory argument while creating an ContentInspection profile of type INLINEINSPECTION or MIRROR.
	*/
	Egressinterface string `json:"egressinterface,omitempty"`
	/**
	* IP Tunnel for CI profile. It is used while creating a ContentInspection profile of type MIRROR when the IDS device is in a different network
	*/
	Iptunnel string `json:"iptunnel,omitempty"`
	/**
	* Egress Vlan for CI
	*/
	Egressvlan *int `json:"egressvlan,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
