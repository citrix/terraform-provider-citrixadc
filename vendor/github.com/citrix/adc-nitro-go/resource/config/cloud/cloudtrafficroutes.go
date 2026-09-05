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

package cloud

/**
* Configuration for cloud trafficroutes resource.
*/
type Cloudtrafficroutes struct {
	/**
	* Name for the traffic cloud route.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Target VPC network name
	*/
	Targetvpcnetwork string `json:"targetvpcnetwork,omitempty"`
	/**
	* Destination IP range in CIDR format
	*/
	Destrange string `json:"destrange,omitempty"`
	/**
	* Next hop IP address
	*/
	Nexthopip string `json:"nexthopip,omitempty"`
	/**
	* cluster owner node id for the nexthopipaddress
	*/
	Ownernode *int `json:"ownernode,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
