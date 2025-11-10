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
* Configuration for portallocation resource.
*/
type Portallocation struct {
	/**
	* IP address for which Port allocation needs to be seen. IP address can be SNIP/VIP/NSIP
	*/
	Srcip string `json:"srcip,omitempty"`
	/**
	* Destination IP address or Server IP for which Port allocation needs to be seen
	*/
	Destip string `json:"destip,omitempty"`
	/**
	* Destination Port or Server port configuration
	*/
	Destport *int `json:"destport,omitempty"`
	/**
	* Protocol for the traffic. TCP traffic: 1, All other protocol traffic: 0
	*/
	Protocol *int `json:"protocol,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid *int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Freeports string `json:"freeports,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
