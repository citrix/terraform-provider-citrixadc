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

package appqoe

/**
* Configuration for QOS parameter resource.
*/
type Appqoeparameter struct {
	/**
	* Time, in seconds, between the first time and the next time the AppQoE alternative content window is displayed. The alternative content window is displayed only once during a session for the same browser accessing a configured URL, so this parameter determines the length of a session.
	*/
	Sessionlife *int `json:"sessionlife,omitempty"`
	/**
	* average number of client connections, that can sit in service waiting queue
	*/
	Avgwaitingclient *int `json:"avgwaitingclient"` // Zero is a valid value
	/**
	* maximum bandwidth which will determine whether to send alternate content response
	*/
	Maxaltrespbandwidth *int `json:"maxaltrespbandwidth,omitempty"`
	/**
	* average number of client connection that can queue up on vserver level without triggering DoS mitigation module
	*/
	Dosattackthresh *int `json:"dosattackthresh"` // Zero is a valid value

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
