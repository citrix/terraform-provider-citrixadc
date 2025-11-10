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
* Configuration for VR ID parameter resource.
*/
type Vridparam struct {
	/**
	* Forward packets to the master node, in an active-active mode configuration, if the virtual server is in the backup state and sharing is disabled.
	*/
	Sendtomaster string `json:"sendtomaster,omitempty"`
	/**
	* Interval, in milliseconds, between vrrp advertisement messages sent to the peer node in active-active mode.
	*/
	Hellointerval *int `json:"hellointerval,omitempty"`
	/**
	* Number of seconds after which a peer node in active-active mode is marked down if vrrp advertisements are not received from the peer node.
	*/
	Deadinterval *int `json:"deadinterval,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
