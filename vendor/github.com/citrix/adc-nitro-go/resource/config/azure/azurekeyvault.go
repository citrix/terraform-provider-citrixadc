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
* Configuration for Azure Key Vault entity resource.
*/
type Azurekeyvault struct {
	/**
	* Name for the Key Vault. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the Key Vault is created.
		CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my keyvault" or 'my keyvault').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the Key Vault configured in Azure cloud using either the Azure CLI or the Azure portal (GUI) with complete domain name. Example: Test.vault.azure.net.
	*/
	Azurevaultname string `json:"azurevaultname,omitempty"`
	/**
	* Name of the Azure Application object created on the ADC appliance. This object will be used for authentication with Azure Active Directory
	*/
	Azureapplication string `json:"azureapplication,omitempty"`

	//------- Read only Parameter ---------;

	State string `json:"state,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
