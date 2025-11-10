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
* Configuration for appflowCollector resource.
*/
type Nsappflowcollector struct {
	/**
	* Name of the AppFlow collector.
	*/
	Name string `json:"name,omitempty"`
	/**
	* The IPv4 address of the AppFlow collector.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* The UDP port on which the AppFlow collector is listening.
	*/
	Port *int `json:"port,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
