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
* Binding class showing the gslbservicegroup that can be bound to lbmonbindings.
*/
type Lbmonbindingsgslbservicegroupbinding struct {
	/**
	* The name of the service group.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* The type of service
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* The state of the servicegroup.
	*/
	Boundservicegroupsvrstate string `json:"boundservicegroupsvrstate,omitempty"`
	/**
	* The configured state (enable/disable) of Monitor on this service.
	*/
	Monstate string `json:"monstate,omitempty"`
	/**
	* The name of the monitor.
	*/
	Monitorname string `json:"monitorname,omitempty"`


}