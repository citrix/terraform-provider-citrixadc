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

package lsn

/**
* Binding class showing the controlchannel that can be bound to lsnsipalgcall.
*/
type Lsnsipalgcallcontrolchannelbinding struct {
	/**
	* IP address of the channel.
	*/
	Channelip string `json:"channelip,omitempty"`
	/**
	* Natted IP address of the channel.
	*/
	Channelnatip string `json:"channelnatip,omitempty"`
	/**
	* port for the channel.
	*/
	Channelport *int `json:"channelport,omitempty"`
	/**
	* Natted port for the channel.
	*/
	Channelnatport *int `json:"channelnatport,omitempty"`
	/**
	* Channel transport protocol.
	*/
	Channelprotocol string `json:"channelprotocol,omitempty"`
	/**
	* Flags for the call entry.
	*/
	Channelflags *int `json:"channelflags,omitempty"`
	/**
	* Timeout for the channel.
	*/
	Channeltimeout *int `json:"channeltimeout,omitempty"`
	/**
	* Call ID for the SIP call.
	*/
	Callid string `json:"callid,omitempty"`


}