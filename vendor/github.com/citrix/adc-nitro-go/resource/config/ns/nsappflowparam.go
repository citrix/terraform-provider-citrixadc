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
* Configuration for appflowParam resource.
*/
type Nsappflowparam struct {
	/**
	* IPFIX template refresh interval (in seconds).
	*/
	Templaterefresh *int `json:"templaterefresh,omitempty"`
	/**
	* MTU to be used for IPFIX UDP packets.
	*/
	Udppmtu *int `json:"udppmtu,omitempty"`
	/**
	* Enable AppFlow HTTP URL logging.
	*/
	Httpurl string `json:"httpurl,omitempty"`
	/**
	* Enable AppFlow HTTP cookie logging.
	*/
	Httpcookie string `json:"httpcookie,omitempty"`
	/**
	* Enable AppFlow HTTP referer logging.
	*/
	Httpreferer string `json:"httpreferer,omitempty"`
	/**
	* Enable AppFlow HTTP method logging.
	*/
	Httpmethod string `json:"httpmethod,omitempty"`
	/**
	* Enable AppFlow HTTP host logging.
	*/
	Httphost string `json:"httphost,omitempty"`
	/**
	* Enable AppFlow HTTP user-agent logging.
	*/
	Httpuseragent string `json:"httpuseragent,omitempty"`
	/**
	* Control whether AppFlow records should be generated only for client-side traffic.
	*/
	Clienttrafficonly string `json:"clienttrafficonly,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
