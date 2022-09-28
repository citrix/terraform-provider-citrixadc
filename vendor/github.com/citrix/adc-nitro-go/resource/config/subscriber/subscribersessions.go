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

package subscriber

/**
* Configuration for subscriber sesions resource.
*/
type Subscribersessions struct {
	/**
	* Subscriber IP Address.
	*/
	Ip string `json:"ip,omitempty"`
	/**
	* The vlan number on which the subscriber is located.
	*/
	Vlan int `json:"vlan,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Subscriptionidtype string `json:"subscriptionidtype,omitempty"`
	Subscriptionidvalue string `json:"subscriptionidvalue,omitempty"`
	Subscriberrules string `json:"subscriberrules,omitempty"`
	Flags string `json:"flags,omitempty"`
	Ttl string `json:"ttl,omitempty"`
	Idlettl string `json:"idlettl,omitempty"`
	Avpdisplaybuffer string `json:"avpdisplaybuffer,omitempty"`
	Servicepath string `json:"servicepath,omitempty"`

}
