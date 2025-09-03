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

package lsn

/**
* Configuration for LSN parameter resource.
*/
type Lsnparameter struct {
	/**
	* Amount of Citrix ADC memory to reserve for the LSN feature, in multiples of 2MB.
		Note: If you later reduce the value of this parameter, the amount of active memory is not reduced. Changing the configured memory limit can only increase the amount of active memory.
		This command is deprecated, use 'set extendedmemoryparam -memlimit' instead.
	*/
	Memlimit int `json:"memlimit,omitempty"`
	/**
	* Synchronize all LSN sessions with the secondary node in a high availability (HA) deployment (global synchronization). After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary node (new primary).
		The global session synchronization parameter and session synchronization parameters (group level) of all LSN groups are enabled by default.
		For a group, when both the global level and the group level LSN session synchronization parameters are enabled, the primary node synchronizes information of all LSN sessions related to this LSN group with the secondary node.
	*/
	Sessionsync string `json:"sessionsync,omitempty"`
	/**
	* LSN global setting for controlling subscriber aware session removal, when this is enabled, when ever the subscriber info is deleted from subscriber database, sessions corresponding to that subscriber will be removed. if this setting is disabled, subscriber sessions will be timed out as per the idle time out settings.
	*/
	Subscrsessionremoval string `json:"subscrsessionremoval,omitempty"`

	//------- Read only Parameter ---------;

	Memlimitactive string `json:"memlimitactive,omitempty"`
	Maxmemlimit string `json:"maxmemlimit,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
