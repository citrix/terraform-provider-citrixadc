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
* Configuration for NAT64 parameter resource.
*/
type Nat64param struct {
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`
	/**
	* Ignore TOS.
	*/
	Nat64ignoretos string `json:"nat64ignoretos,omitempty"`
	/**
	* Calculate checksum for UDP packets with zero checksum
	*/
	Nat64zerochecksum string `json:"nat64zerochecksum,omitempty"`
	/**
	* MTU setting for the IPv6 side. If the incoming IPv4 packet greater than this, either fragment or send icmp need fragmentation error.
	*/
	Nat64v6mtu uint32 `json:"nat64v6mtu,omitempty"`
	/**
	* When disabled, translator will not insert IPv6 fragmentation header for non fragmented IPv4 packets
	*/
	Nat64fragheader string `json:"nat64fragheader,omitempty"`

}
