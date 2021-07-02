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
* Binding class showing the srcip that can be bound to vxlan.
*/
type Vxlansrcipbinding struct {
	/**
	* The source IP address to use in outgoing vxlan packets.
	*/
	Srcip string `json:"srcip,omitempty"`
	/**
	* A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.
	*/
	Id uint32 `json:"id,omitempty"`


}