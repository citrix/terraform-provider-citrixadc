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

package reputation

/**
* Configuration for Reputation service settings resource.
*/
type Reputationsettings struct {
	/**
	* Proxy server IP to get Reputation data.
	*/
	Proxyserver string `json:"proxyserver,omitempty"`
	/**
	* Proxy server port.
	*/
	Proxyport *int `json:"proxyport,omitempty"`
	/**
	* Proxy Username
	*/
	Proxyusername string `json:"proxyusername,omitempty"`
	/**
	* Password with which user logs on.
	*/
	Proxypassword string `json:"proxypassword,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
