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

package snmp

/**
* Configuration for SNMP user resource.
*/
type Snmpuser struct {
	/**
	* Name for the SNMPv3 user. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose it in double or single quotation marks (for example, "my user" or 'my user').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Name of the configured SNMPv3 group to which to bind this SNMPv3 user. The access rights (bound SNMPv3 views) and security level set for this group are assigned to this user.
	*/
	Group string `json:"group,omitempty"`
	/**
	* Authentication algorithm used by the Citrix ADC and the SNMPv3 user for authenticating the communication between them. You must specify the same authentication algorithm when you configure the SNMPv3 user in the SNMP manager.
	*/
	Authtype string `json:"authtype,omitempty"`
	/**
	* Plain-text pass phrase to be used by the authentication algorithm specified by the authType (Authentication Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the pass phrase includes one or more spaces, enclose it in double or single quotation marks (for example, "my phrase" or 'my phrase').
	*/
	Authpasswd string `json:"authpasswd,omitempty"`
	/**
	* Encryption algorithm used by the Citrix ADC and the SNMPv3 user for encrypting the communication between them. You must specify the same encryption algorithm when you configure the SNMPv3 user in the SNMP manager.
	*/
	Privtype string `json:"privtype,omitempty"`
	/**
	* Encryption key to be used by the encryption algorithm specified by the privType (Encryption Type) parameter. Can consist of 8 to 63 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.
		The following requirement applies only to the Citrix ADC CLI:
		If the key includes one or more spaces, enclose it in double or single quotation marks (for example, "my key" or 'my key').
	*/
	Privpasswd string `json:"privpasswd,omitempty"`

	//------- Read only Parameter ---------;

	Engineid string `json:"engineid,omitempty"`
	Storagetype string `json:"storagetype,omitempty"`
	Status string `json:"status,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
