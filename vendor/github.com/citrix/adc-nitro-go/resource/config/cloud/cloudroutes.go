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
* Configuration for cloud routes resource.
*/
type Cloudroutes struct {
	/**
	* Name for the route.
	*/
	Name string `json:"name,omitempty"`
	/**
	* client vpc network name
	*/
	Routesvpcnetwork string `json:"routesvpcnetwork,omitempty"`
	/**
	* vip subnet in CIDR format
	*/
	Vipsubnet string `json:"vipsubnet,omitempty"`
	/**
	* vip vpc network name
	*/
	Vipvpcnetwork string `json:"vipvpcnetwork,omitempty"`
	/**
	* IPv4 or IPv6 address attached to the  nic interface towards vpc mentiond in vpcnetwork
	*/
	Clientipaddress string `json:"clientipaddress,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
