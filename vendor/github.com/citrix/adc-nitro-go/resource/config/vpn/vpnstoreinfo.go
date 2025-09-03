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
* Configuration for Store information for a URL resource.
*/
type Vpnstoreinfo struct {
	/**
	* StoreFront URL to be scanned
	*/
	Url string `json:"url,omitempty"`

	//------- Read only Parameter ---------;

	Storeserverstatus string `json:"storeserverstatus,omitempty"`
	Storeserverissf string `json:"storeserverissf,omitempty"`
	Storeapisupport string `json:"storeapisupport,omitempty"`
	Storelist string `json:"storelist,omitempty"`
	Storestatus string `json:"storestatus,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
