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
* Configuration for Azure Key Vault entity resource.
*/
type Authenticationazurekeyvault struct {
	/**
	* Name for the new Azure Key Vault profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the Azure vault account as configured in azure portal.
	*/
	Vaultname string `json:"vaultname,omitempty"`
	/**
	* Unique identity of the relying party requesting for authentication.
	*/
	Clientid string `json:"clientid,omitempty"`
	/**
	* Unique secret string to authorize relying party at authorization server.
	*/
	Clientsecret string `json:"clientsecret,omitempty"`
	/**
	* Friendly name of the Key to be used to compute signature.
	*/
	Servicekeyname string `json:"servicekeyname,omitempty"`
	/**
	* Algorithm to be used to sign/verify transactions
	*/
	Signaturealg string `json:"signaturealg,omitempty"`
	/**
	* URL endpoint on relying party to which the OAuth token is to be sent.
	*/
	Tokenendpoint string `json:"tokenendpoint,omitempty"`
	/**
	* Name of the service used to send push notifications
	*/
	Pushservice string `json:"pushservice,omitempty"`
	/**
	* This is the group that is added to user sessions that match current IdP policy. It can be used in policies to identify relying party trust.
	*/
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	/**
	* Interval at which access token in obtained.
	*/
	Refreshinterval int `json:"refreshinterval,omitempty"`
	/**
	* TenantID of the application. This is usually specific to providers such as Microsoft and usually refers to the deployment identifier.
	*/
	Tenantid string `json:"tenantid,omitempty"`
	/**
	* If authentication is disabled, otp checks are not performed after azure vault keys are obtained. This is useful to distinguish whether user has registered devices. 
	*/
	Authentication string `json:"authentication,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
