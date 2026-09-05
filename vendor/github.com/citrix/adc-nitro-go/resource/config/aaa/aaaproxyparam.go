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

package aaa

/**
* Configuration for AAA proxy parameter resource.
*/
type Aaaproxyparam struct {
	/**
	* IP address and Port of the proxy server to be used for HTTP access for this request. This can be configured in the following manner: First way is to configure in ipaddress:port format like a.b.c.d:e or this can be specified via the alternate format which is to specify like http://a.b.c.d without port or http://a.b.c.d:8080 with port
	*/
	Proxy string `json:"proxy,omitempty"`
	/**
	* this indicates whether Proxy-Authorization header will be sent or not
	*/
	Proxyauthorization string `json:"proxyauthorization,omitempty"`
	/**
	* username that will be sent as part of Basic Proxy-Authorization header
	*/
	Proxyusername string `json:"proxyusername,omitempty"`
	/**
	* password that will be sent as part of Basic Proxy-Authorization header
	*/
	Proxypassword string `json:"proxypassword,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
