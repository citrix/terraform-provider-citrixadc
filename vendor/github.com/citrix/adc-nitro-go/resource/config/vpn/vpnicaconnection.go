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

package vpn

/**
* Configuration for active ica connections resource.
*/
type Vpnicaconnection struct {
	/**
	* User name for which ica connections needs to be terminated.
	*/
	Username string `json:"username,omitempty"`
	/**
	* Transport type for the existing Existing ICA conenction.
	*/
	Transproto string `json:"transproto,omitempty"`
	/**
	* Unique number that identifies the cluster node.
	*/
	Nodeid *int `json:"nodeid,omitempty"`
	/**
	* Terminate all active icaconnections.
	*/
	All bool `json:"all,omitempty"`

	//------- Read only Parameter ---------;

	Domain string `json:"domain,omitempty"`
	Srcip string `json:"srcip,omitempty"`
	Srcport string `json:"srcport,omitempty"`
	Destip string `json:"destip,omitempty"`
	Destport string `json:"destport,omitempty"`
	Peid string `json:"peid,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
