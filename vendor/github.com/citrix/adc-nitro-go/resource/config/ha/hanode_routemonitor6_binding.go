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
* Binding class showing the routemonitor6 that can be bound to hanode.
*/
type Hanoderoutemonitor6binding struct {
	/**
	* The IP address (IPv4 or IPv6).
	*/
	Routemonitor string `json:"routemonitor,omitempty"`
	/**
	* The netmask.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* The flags for this entry.
	*/
	Flags int `json:"flags,omitempty"`
	/**
	* State for route monitor
	*/
	Routemonitorstate string `json:"routemonitorstate,omitempty"`
	/**
	* Number that uniquely identifies the local node. The ID of the local node is always 0.
	*/
	Id int `json:"id"` // Zero is a valid value


}