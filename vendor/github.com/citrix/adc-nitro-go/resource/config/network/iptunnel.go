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
* Configuration for ip Tunnel resource.
*/
type Iptunnel struct {
	/**
	* Name for the IP tunnel. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).
	*/
	Name string `json:"name,omitempty"`
	/**
	* Public IPv4 address, of the remote device, used to set up the tunnel. For this parameter, you can alternatively specify a network address.
	*/
	Remote string `json:"remote,omitempty"`
	/**
	* Subnet mask of the remote IP address of the tunnel.
	*/
	Remotesubnetmask string `json:"remotesubnetmask,omitempty"`
	/**
	* Type of Citrix ADC owned public IPv4 address, configured on the local Citrix ADC and used to set up the tunnel.
	*/
	Local string `json:"local,omitempty"`
	/**
	* Name of the protocol to be used on this tunnel.
	*/
	Protocol string `json:"protocol,omitempty"`
	/**
	* Virtual network identifier (VNID) is the value that identifies a specific virtual network in the data plane.
	*/
	Vnid *int `json:"vnid,omitempty"`
	/**
	* Option to select Vlan Tagging.
	*/
	Vlantagging string `json:"vlantagging,omitempty"`
	/**
	* Specifies UDP destination port for Geneve packets. Default port is 6081.
	*/
	Destport *int `json:"destport,omitempty"`
	/**
	* Default behavior is to copy the ToS field of the internal IP Packet (Payload) to the outer IP packet (Transport packet). But the user can configure a new ToS field using this option.
	*/
	Tosinherit string `json:"tosinherit,omitempty"`
	/**
	* The payload GRE will carry
	*/
	Grepayload string `json:"grepayload,omitempty"`
	/**
	* Name of IPSec profile to be associated.
	*/
	Ipsecprofilename string `json:"ipsecprofilename,omitempty"`
	/**
	* The vlan for mulicast packets
	*/
	Vlan *int `json:"vlan,omitempty"`
	/**
	* The owner node group in a Cluster for the iptunnel.
	*/
	Ownergroup string `json:"ownergroup,omitempty"`

	//------- Read only Parameter ---------;

	Sysname string `json:"sysname,omitempty"`
	Type string `json:"type,omitempty"`
	Encapip string `json:"encapip,omitempty"`
	Channel string `json:"channel,omitempty"`
	Tunneltype string `json:"tunneltype,omitempty"`
	Ipsectunnelstatus string `json:"ipsectunnelstatus,omitempty"`
	Refcnt string `json:"refcnt,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
