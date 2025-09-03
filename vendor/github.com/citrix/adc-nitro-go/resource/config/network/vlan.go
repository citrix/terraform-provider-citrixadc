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
* Configuration for "VLAN" resource.
*/
type Vlan struct {
	/**
	* A positive integer that uniquely identifies a VLAN.
	*/
	Id int `json:"id,omitempty"`
	/**
	* A name for the VLAN. Must begin with a letter, a number, or the underscore symbol, and can consist of from 1 to 31 letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a name that helps identify the VLAN. However, you cannot perform any VLAN operation by specifying this name instead of the VLAN ID.
	*/
	Aliasname string `json:"aliasname,omitempty"`
	/**
	* Enable dynamic routing on this VLAN.
	*/
	Dynamicrouting string `json:"dynamicrouting,omitempty"`
	/**
	* Enable all IPv6 dynamic routing protocols on this VLAN. Note: For the ENABLED setting to work, you must configure IPv6 dynamic routing protocols from the VTYSH command line.
	*/
	Ipv6dynamicrouting string `json:"ipv6dynamicrouting,omitempty"`
	/**
	* Specifies the maximum transmission unit (MTU), in bytes. The MTU is the largest packet size, excluding 14 bytes of ethernet header and 4 bytes of crc, that can be transmitted and received over this VLAN.
	*/
	Mtu int `json:"mtu,omitempty"`
	/**
	* If sharing is enabled, then this vlan can be shared across multiple partitions by binding it to all those partitions. If sharing is disabled, then this vlan can be bound to only one of the partitions.
	*/
	Sharing string `json:"sharing,omitempty"`

	//------- Read only Parameter ---------;

	Linklocalipv6addr string `json:"linklocalipv6addr,omitempty"`
	Rnat string `json:"rnat,omitempty"`
	Portbitmap string `json:"portbitmap,omitempty"`
	Lsbitmap string `json:"lsbitmap,omitempty"`
	Tagbitmap string `json:"tagbitmap,omitempty"`
	Lstagbitmap string `json:"lstagbitmap,omitempty"`
	Ifaces string `json:"ifaces,omitempty"`
	Tagifaces string `json:"tagifaces,omitempty"`
	Ifnum string `json:"ifnum,omitempty"`
	Tagged string `json:"tagged,omitempty"`
	Vlantd string `json:"vlantd,omitempty"`
	Sdxvlan string `json:"sdxvlan,omitempty"`
	Partitionname string `json:"partitionname,omitempty"`
	Vxlan string `json:"vxlan,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
