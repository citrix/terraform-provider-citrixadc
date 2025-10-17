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
	Vlan *int `json:"vlan,omitempty"`
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
	Currhoplimit *int `json:"currhoplimit,omitempty"`
	/**
	* Maximum time allowed between unsolicited multicast RAs, in seconds.
	*/
	Maxrtadvinterval *int `json:"maxrtadvinterval,omitempty"`
	/**
	* Minimum time interval between RA messages, in seconds.
	*/
	Minrtadvinterval *int `json:"minrtadvinterval,omitempty"`
	/**
	* The Link MTU.
	*/
	Linkmtu *int `json:"linkmtu,omitempty"`
	/**
	* Reachable time, in milliseconds.
	*/
	Reachabletime *int `json:"reachabletime,omitempty"`
	/**
	* Retransmission time, in milliseconds.
	*/
	Retranstime *int `json:"retranstime,omitempty"`
	/**
	* Default life time, in seconds.
	*/
	Defaultlifetime *int `json:"defaultlifetime,omitempty"`

	//------- Read only Parameter ---------;

	Lastrtadvtime string `json:"lastrtadvtime,omitempty"`
	Nextrtadvdelay string `json:"nextrtadvdelay,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
