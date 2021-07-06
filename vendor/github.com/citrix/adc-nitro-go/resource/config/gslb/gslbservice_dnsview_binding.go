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

package gslb

/**
* Binding class showing the dnsview that can be bound to gslbservice.
*/
type Gslbservicednsviewbinding struct {
	/**
	* Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.
	*/
	Viewname string `json:"viewname,omitempty"`
	/**
	* IP address to be used for the given view
	*/
	Viewip string `json:"viewip,omitempty"`
	/**
	* Name of the GSLB service.
	*/
	Servicename string `json:"servicename,omitempty"`


}