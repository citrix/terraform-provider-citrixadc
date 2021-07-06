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

package network

/**
* Binding class showing the nsip6 that can be bound to ipset.
*/
type Ipsetnsip6binding struct {
	/**
	* One or more IP addresses bound to the IP set.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* Name of the IP set to which to bind IP addresses.
	*/
	Name string `json:"name,omitempty"`


}