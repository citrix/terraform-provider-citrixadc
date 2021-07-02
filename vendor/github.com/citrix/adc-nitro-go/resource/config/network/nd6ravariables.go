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
* Configuration for nd6 Router Advertisment configuration variables resource.
*/
type Nd6ravariables struct {
	/**
	* The VLAN number.
	*/
	Vlan uint32 `json:"vlan,omitempty"`
	/**
	* Cease router advertisements on this vlan.
	*/
	Ceaserouteradv string `json:"ceaserouteradv,omitempty"`
	/**
	* whether the router sends periodic RAs and responds to Router Solicitations.
	*/
	Sendrouteradv string `json:"sendrouteradv,omitempty"`
	/**
	* Include source link layer address option in RA messages.
	*/
	Srclinklayeraddroption string `json:"srclinklayeraddroption,omitempty"`
	/**
	* Send only Unicast Router Advertisements in respond to Router Solicitations.
	*/
	Onlyunicastrtadvresponse string `json:"onlyunicastrtadvresponse,omitempty"`
	/**
	* Value to be placed in the Managed address configuration flag field.
	*/
	Managedaddrconfig string `json:"managedaddrconfig,omitempty"`
	/**
	* Value to be placed in the Other configuration flag field.
	*/
	Otheraddrconfig string `json:"otheraddrconfig,omitempty"`
	/**
	* Current Hop limit.
	*/
	Currhoplimit uint32 `json:"currhoplimit,omitempty"`
	/**
	* Maximum time allowed between unsolicited multicast RAs, in seconds.
	*/
	Maxrtadvinterval uint32 `json:"maxrtadvinterval,omitempty"`
	/**
	* Minimum time interval between RA messages, in seconds.
	*/
	Minrtadvinterval uint32 `json:"minrtadvinterval,omitempty"`
	/**
	* The Link MTU.
	*/
	Linkmtu uint32 `json:"linkmtu,omitempty"`
	/**
	* Reachable time, in milliseconds.
	*/
	Reachabletime uint32 `json:"reachabletime,omitempty"`
	/**
	* Retransmission time, in milliseconds.
	*/
	Retranstime uint32 `json:"retranstime,omitempty"`
	/**
	* Default life time, in seconds.
	*/
	Defaultlifetime int32 `json:"defaultlifetime,omitempty"`

	//------- Read only Parameter ---------;

	Lastrtadvtime string `json:"lastrtadvtime,omitempty"`
	Nextrtadvdelay string `json:"nextrtadvdelay,omitempty"`

}
