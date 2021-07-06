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

package basic

/**
* Configuration for virtual server resource.
*/
type Vserver struct {
	/**
	* The name of the virtual server to be removed.
	*/
	Name string `json:"name,omitempty"`
	/**
	* The name of the backup virtual server for this virtual server.
	*/
	Backupvserver string `json:"backupvserver,omitempty"`
	/**
	* The URL where traffic is redirected if the virtual server in the system becomes unavailable.
	*/
	Redirecturl string `json:"redirecturl,omitempty"`
	/**
	* Use this option to specify whether a virtual server (used for load balancing or content switching) routes requests to the cache redirection virtual server before sending it to the configured servers.
	*/
	Cacheable string `json:"cacheable,omitempty"`
	/**
	* The timeout value in seconds for idle client connection
	*/
	Clttimeout int `json:"clttimeout,omitempty"`
	/**
	* The spillover factor. The system will use this value to determine if it should send traffic to the backupvserver when the main virtual server reaches the spillover threshold.
	*/
	Somethod string `json:"somethod,omitempty"`
	/**
	* The state of the spillover persistence.
	*/
	Sopersistence string `json:"sopersistence,omitempty"`
	/**
	* The spillover persistence entry timeout.
	*/
	Sopersistencetimeout int `json:"sopersistencetimeout,omitempty"`
	/**
	* The spillver threshold value.
	*/
	Sothreshold int `json:"sothreshold,omitempty"`
	/**
	* The lb vserver of type PUSH/SSL_PUSH to which server pushes the updates received on the client facing non-push lb vserver.
	*/
	Pushvserver string `json:"pushvserver,omitempty"`

}
