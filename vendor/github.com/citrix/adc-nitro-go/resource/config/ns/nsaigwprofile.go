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
* Configuration for AI GW profile resource.
*/
type Nsaigwprofile struct {
	/**
	* Name of the AIGW Profile.
	*/
	Name string `json:"name,omitempty"`
	/**
	* The type of AI GW endpoint type
	*/
	Endpointtype string `json:"endpointtype,omitempty"`
	/**
	* The binding entity for the aigw profile
	*/
	Profiletype string `json:"profiletype,omitempty"`
	/**
	* Token capacity of the backend server.
	*/
	Tokenquota *int `json:"tokenquota,omitempty"`
	/**
	* Quota refresh rate, in minutes.
	*/
	Quotarefreshfrequency *int `json:"quotarefreshfrequency,omitempty"`
	/**
	* Authentication token/API Key for the AI GW Endpoint.
	*/
	Authtoken string `json:"authtoken,omitempty"`

	//------- Read only Parameter ---------;

	Refcnt string `json:"refcnt,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
