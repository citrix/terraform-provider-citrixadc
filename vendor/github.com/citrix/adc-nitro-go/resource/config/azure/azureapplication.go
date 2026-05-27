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

package azure

/**
* Configuration for Azure Application resource.
*/
type Azureapplication struct {
	/**
	* Name for the application. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the application is created.',
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my application" or 'my application').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Application ID that is generated when an application is created in Azure Active Directory using either the Azure CLI or the Azure portal (GUI)
	*/
	Clientid string `json:"clientid,omitempty"`
	/**
	* Password for the application configured in Azure Active Directory. The password is specified in the Azure CLI or generated in the Azure portal (GUI).
	*/
	Clientsecret string `json:"clientsecret,omitempty"`
	/**
	* ID of the directory inside Azure Active Directory in which the application was created
	*/
	Tenantid string `json:"tenantid,omitempty"`
	/**
	* Vault resource for which access token is granted. Example : vault.azure.net
	*/
	Vaultresource string `json:"vaultresource,omitempty"`
	/**
	* URL from where access token can be obtained. If the token end point is not specified, the default value is https://login.microsoftonline.com/<tenant id>.
	*/
	Tokenendpoint string `json:"tokenendpoint,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
