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

package network

/**
* Configuration for bridge group resource.
*/
type Bridgegroup struct {
	/**
	* An integer that uniquely identifies the bridge group.
	*/
	Id int `json:"id,omitempty"`
	/**
	* Enable dynamic routing for this bridgegroup.
	*/
	Dynamicrouting string `json:"dynamicrouting,omitempty"`
	/**
	* Enable all IPv6 dynamic routing protocols on all VLANs bound to this bridgegroup. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.
	*/
	Ipv6dynamicrouting string `json:"ipv6dynamicrouting,omitempty"`

	//------- Read only Parameter ---------;

	Flags string `json:"flags,omitempty"`
	Portbitmap string `json:"portbitmap,omitempty"`
	Tagbitmap string `json:"tagbitmap,omitempty"`
	Ifaces string `json:"ifaces,omitempty"`
	Tagifaces string `json:"tagifaces,omitempty"`
	Rnat string `json:"rnat,omitempty"`
	Partitionname string `json:"partitionname,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
