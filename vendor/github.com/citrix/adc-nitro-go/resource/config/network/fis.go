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

package network

/**
* Configuration for "FIS" resource.
*/
type Fis struct {
	/**
	* Name for the FIS to be created. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ). Note: In a cluster setup, the FIS name on each node must be unique.
	*/
	Name string `json:"name,omitempty"`
	/**
	* ID of the cluster node for which you are creating the FIS. Can be configured only through the cluster IP address.
	*/
	Ownernode int `json:"ownernode,omitempty"`

	//------- Read only Parameter ---------;

	Ifaces string `json:"ifaces,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
