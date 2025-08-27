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

package ssl

/**
* Binding class showing the sslcacertbundle that can be bound to sslservice.
*/
type Sslservicesslcacertbundlebinding struct {
	/**
	* CA certbundle name bound to the service.
	*/
	Cacertbundlename string `json:"cacertbundlename,omitempty"`
	/**
	* The flag is used to indicate whether all CA_names in this particular CA certificate bundle needs to be sent to the SSL client while requesting for client certificate in a SSL handshake
	*/
	Skipcacertbundle bool `json:"skipcacertbundle,omitempty"`
	/**
	* Name of the SSL service for which to set advanced configuration.
	*/
	Servicename string `json:"servicename,omitempty"`


}