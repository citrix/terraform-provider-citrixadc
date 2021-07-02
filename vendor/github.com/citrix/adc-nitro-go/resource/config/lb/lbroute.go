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

package lb

/**
* Configuration for LB route resource.
*/
type Lbroute struct {
	/**
	* The IP address of the network to which the route belongs.
	*/
	Network string `json:"network,omitempty"`
	/**
	* The netmask to which the route belongs.
	*/
	Netmask string `json:"netmask,omitempty"`
	/**
	* The name of the route.
	*/
	Gatewayname string `json:"gatewayname,omitempty"`
	/**
	* Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
	*/
	Td uint32 `json:"td,omitempty"`

	//------- Read only Parameter ---------;

	Flags string `json:"flags,omitempty"`

}
