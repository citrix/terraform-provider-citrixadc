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
	Tcpthreshold uint32 `json:"tcpthreshold,omitempty"`
	/**
	* Number of UDP packets permitted per 10 milliseconds.
	*/
	Udpthreshold uint32 `json:"udpthreshold,omitempty"`
	/**
	* Number of ICMP packets permitted per 10 milliseconds.
	*/
	Icmpthreshold uint32 `json:"icmpthreshold,omitempty"`
	/**
	* The number of TCP RST packets permitted per 10 milli second. zero means rate control is disabled and 0xffffffff means every thing is rate controlled
	*/
	Tcprstthreshold uint32 `json:"tcprstthreshold,omitempty"`

}
