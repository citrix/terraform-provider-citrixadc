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
* Configuration for Secure Private Access Profile resource.
*/
type Vpnsecureprivateaccessprofile struct {
	/**
	* name of Secure Private Access profile
	*/
	Name string `json:"name,omitempty"`
	/**
	* Public URL for your Secure Private Access site or load balancing server.
	*/
	Url string `json:"url,omitempty"`
	/**
	* Customer ID of the citrix cloud customer.
	*/
	Customerid string `json:"customerid,omitempty"`
	/**
	* Secure Private Access Chrome Enterprise Premium mode of operation.
	*/
	Chromeenterprisepremiummode string `json:"chromeenterprisepremiummode,omitempty"`
	/**
	* Your organization's unique ID on Google's Admin console Profile settings.
	*/
	Googlecustomerid string `json:"googlecustomerid,omitempty"`
	/**
	* The ID of the Google Secure Gateway.
	*/
	Googlesecuritygatewayid string `json:"googlesecuritygatewayid,omitempty"`
	/**
	* Automatically configures the session for Citrix Secure Access client connectivity.
	*/
	Forceclienttype string `json:"forceclienttype,omitempty"`
	/**
	* Secure Private Access Shared Secret.
	*/
	Sharedsecret string `json:"sharedsecret,omitempty"`

	//------- Read only Parameter ---------;

	Clouddeployment string `json:"clouddeployment,omitempty"`
	Accessrestrictedpageredirect string `json:"accessrestrictedpageredirect,omitempty"`
	Serverstatus string `json:"serverstatus,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
