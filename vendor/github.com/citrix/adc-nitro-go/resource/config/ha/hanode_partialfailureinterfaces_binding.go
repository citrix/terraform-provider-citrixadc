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

package ha

/**
* Binding class showing the partialfailureinterfaces that can be bound to hanode.
*/
type Hanodepartialfailureinterfacesbinding struct {
	/**
	* Interfaces causing Partial Failure.
	*/
	Pfifaces string `json:"pfifaces,omitempty"`
	/**
	* Number that uniquely identifies the local node. The ID of the local node is always 0.
	*/
	Id int `json:"id,omitempty"`
	/**
	* A route that you want the Citrix ADC to monitor in its internal routing table. You can specify an IPv4 address or network, or an IPv6 address or network prefix. If you specify an IPv4 network address or IPv6 network prefix, the appliance monitors any route that matches the network or prefix.
	*/
	Routemonitor string `json:"routemonitor,omitempty"`


}