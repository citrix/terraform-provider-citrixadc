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
* Configuration for capacity resource.
*/
type Nscapacity struct {
	/**
	* System bandwidth limit.
	*/
	Bandwidth int `json:"bandwidth,omitempty"`
	/**
	* appliance platform type.
	*/
	Platform string `json:"platform,omitempty"`
	/**
	* licensed using vcpu pool.
	*/
	Vcpu bool `json:"vcpu,omitempty"`
	/**
	* Product edition.
	*/
	Edition string `json:"edition,omitempty"`
	/**
	* Bandwidth unit.
	*/
	Unit string `json:"unit,omitempty"`
	/**
	* Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
	*/
	Username string `json:"username,omitempty"`
	/**
	* Password to use when authenticating with ADM Agent for LAS licensing.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid int `json:"nodeid,omitempty"`

	//------- Read only Parameter ---------;

	Actualbandwidth string `json:"actualbandwidth,omitempty"`
	Vcpucount string `json:"vcpucount,omitempty"`
	Maxvcpucount string `json:"maxvcpucount,omitempty"`
	Maxbandwidth string `json:"maxbandwidth,omitempty"`
	Minbandwidth string `json:"minbandwidth,omitempty"`
	Instancecount string `json:"instancecount,omitempty"`
	Daystoexpiration string `json:"daystoexpiration,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
