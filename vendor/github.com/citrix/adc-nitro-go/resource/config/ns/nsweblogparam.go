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
* Configuration for Web log parameters resource.
*/
type Nsweblogparam struct {
	/**
	* Buffer size, in MB, allocated for log transaction data on the system. The maximum value is limited to the memory available on the system.
	*/
	Buffersizemb *int `json:"buffersizemb,omitempty"`
	/**
	* Name(s) of HTTP request headers whose values should be exported by the Web Logging feature.
	*/
	Customreqhdrs []string `json:"customreqhdrs,omitempty"`
	/**
	* Name(s) of HTTP response headers whose values should be exported by the Web Logging feature.
	*/
	Customrsphdrs []string `json:"customrsphdrs,omitempty"`

	//------- Read only Parameter ---------;

	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
