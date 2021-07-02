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
* Binding class showing the servicegroup that can be bound to lbmonitor.
*/
type Lbmonitorservicegroupbinding struct {
	/**
	* Name of the monitor.
	*/
	Monitorname string `json:"monitorname,omitempty"`
	/**
	* Name of the service or service group.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* State of the monitor. The state setting for a monitor of a given type affects all monitors of that type. For example, if an HTTP monitor is enabled, all HTTP monitors on the appliance are (or remain) enabled. If an HTTP monitor is disabled, all HTTP monitors on the appliance are disabled.
	*/
	Dupstate string `json:"dup_state,omitempty"`
	/**
	* Weight to assign to the binding between the monitor and service.
	*/
	Dupweight uint32 `json:"dup_weight,omitempty"`
	/**
	* Name of the service group.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* State of the monitor. The state setting for a monitor of a given type affects all monitors of that type. For example, if an HTTP monitor is enabled, all HTTP monitors on the appliance are (or remain) enabled. If an HTTP monitor is disabled, all HTTP monitors on the appliance are disabled.
	*/
	State string `json:"state,omitempty"`
	/**
	* Weight to assign to the binding between the monitor and service.
	*/
	Weight uint32 `json:"weight,omitempty"`


}