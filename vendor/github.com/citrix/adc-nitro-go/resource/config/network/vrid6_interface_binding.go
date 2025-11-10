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
* Binding class showing the interface that can be bound to vrid6.
*/
type Vrid6interfacebinding struct {
	/**
	* Interfaces to bind to the VMAC6, specified in (slot/port) notation (for example, 1/2).Use spaces to separate multiple entries.
	*/
	Ifnum string `json:"ifnum,omitempty"`
	/**
	* The VLAN in which this VRID resides.
	*/
	Vlan *int `json:"vlan,omitempty"`
	/**
	* Flags.
	*/
	Flags *int `json:"flags,omitempty"`
	/**
	* Integer value that uniquely identifies a VMAC6 address.
	*/
	Id *int `json:"id,omitempty"`


}