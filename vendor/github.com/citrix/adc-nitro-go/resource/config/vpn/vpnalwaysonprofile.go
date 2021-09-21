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

package vpn

/**
* Configuration for AlwyasON profile resource.
*/
type Vpnalwaysonprofile struct {
	/**
	* name of AlwaysON profile
	*/
	Name string `json:"name,omitempty"`
	/**
	* Option to block network traffic when tunnel is not established(and the config requires that tunnel be established). When set to onlyToGateway, the network traffic to and from the client (except Gateway IP) is blocked. When set to fullAccess, the network traffic is not blocked
	*/
	Networkaccessonvpnfailure string `json:"networkaccessonvpnfailure,omitempty"`
	/**
	* Allow/Deny user to log off and connect to another Gateway
	*/
	Clientcontrol string `json:"clientcontrol,omitempty"`
	/**
	* Option to decide if tunnel should be established when in enterprise network. When locationBasedVPN is remote, client tries to detect if it is located in enterprise network or not and establishes the tunnel if not in enterprise network. Dns suffixes configured using -add dns suffix- are used to decide if the client is in the enterprise network or not. If the resolution of the DNS suffix results in private IP, client is said to be in enterprise network. When set to EveryWhere, the client skips the check to detect if it is on the enterprise network and tries to establish the tunnel
	*/
	Locationbasedvpn string `json:"locationbasedvpn,omitempty"`

}
