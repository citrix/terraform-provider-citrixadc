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
* Binding class showing the service that can be bound to lbmonbindings.
*/
type Lbmonbindingsservicebinding struct {
	/**
	* The name of the service.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* The IPAddress of the service.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* The port of the service.
	*/
	Port *int `json:"port,omitempty"`
	/**
	* The type of service
	*/
	Servicetype string `json:"servicetype,omitempty"`
	/**
	* The state of the service
	*/
	Svrstate string `json:"svrstate,omitempty"`
	/**
	* The configured state (enable/disable) of Monitor on this service.
	*/
	Monsvcstate string `json:"monsvcstate,omitempty"`
	/**
	* The name of the monitor.
	*/
	Monitorname string `json:"monitorname,omitempty"`


}