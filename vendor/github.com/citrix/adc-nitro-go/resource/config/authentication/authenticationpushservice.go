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

package authentication

/**
* Configuration for Service details for sending push notifications resource.
*/
type Authenticationpushservice struct {
	/**
	* Name for the push service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my push service" or 'my push service'). 
	*/
	Name string `json:"name,omitempty"`
	/**
	* Unique identity for communicating with Citrix Push server in cloud.
	*/
	Clientid string `json:"clientid,omitempty"`
	/**
	* Unique secret for communicating with Citrix Push server in cloud.
	*/
	Clientsecret string `json:"clientsecret,omitempty"`
	/**
	* Customer id/name of the account in cloud that is used to create clientid/secret pair.
	*/
	Customerid string `json:"customerid,omitempty"`
	/**
	* Interval at which certificates or idtoken is refreshed.
	*/
	Refreshinterval int `json:"refreshinterval,omitempty"`

	//------- Read only Parameter ---------;

	Namespace string `json:"Namespace,omitempty"`
	Hubname string `json:"hubname,omitempty"`
	Servicekey string `json:"servicekey,omitempty"`
	Servicekeyname string `json:"servicekeyname,omitempty"`
	Certendpoint string `json:"certendpoint,omitempty"`
	Pushservicestatus string `json:"pushservicestatus,omitempty"`
	Trustservice string `json:"trustservice,omitempty"`
	Pushcloudserverstatus string `json:"pushcloudserverstatus,omitempty"`
	Signingkeyname string `json:"signingkeyname,omitempty"`
	Signingkey string `json:"signingkey,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
