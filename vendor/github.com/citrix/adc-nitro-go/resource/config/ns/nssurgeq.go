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
* Configuration for surge queue resource.
*/
type Nssurgeq struct {
	/**
	* Name of a virtual server, service or service group for which the SurgeQ must be flushed.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of a service group member. This argument is needed when you want to flush the SurgeQ of a service group.
	*/
	Servername string `json:"servername,omitempty"`
	/**
	* port on which server is bound to the entity(Servicegroup).
	*/
	Port int32 `json:"port,omitempty"`

}
