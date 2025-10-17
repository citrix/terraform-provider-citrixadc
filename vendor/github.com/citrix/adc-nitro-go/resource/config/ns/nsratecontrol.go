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

package ns

/**
* Configuration for rate control resource.
*/
type Nsratecontrol struct {
	/**
	* Number of SYNs permitted per 10 milliseconds.
	*/
	Tcpthreshold *int `json:"tcpthreshold"` // Zero is acceptable value
	/**
	* Number of UDP packets permitted per 10 milliseconds.
	*/
	Udpthreshold *int `json:"udpthreshold"` // Zero is acceptable value
	/**
	* Number of ICMP packets permitted per 10 milliseconds.
	*/
	Icmpthreshold *int `json:"icmpthreshold"` // Zero is acceptable value
	/**
	* The number of TCP RST packets permitted per 10 milliseconds. zero means rate control is disabled and 0xffffffff means every thing is rate controlled
	*/
	Tcprstthreshold *int `json:"tcprstthreshold"` // Zero is acceptable value

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`
}
