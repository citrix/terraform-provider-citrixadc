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
* Configuration for SNMP group resource.
*/
type Snmpgroup struct {
	/**
	* Name for the SNMPv3 group. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  You should choose a name that helps identify the SNMPv3 group. 
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose it in double or single quotation marks (for example, "my name" or 'my name').
	*/
	Name string `json:"name,omitempty"`
	/**
	* Security level required for communication between the Citrix ADC and the SNMPv3 users who belong to the group. Specify one of the following options:
		noAuthNoPriv. Require neither authentication nor encryption.
		authNoPriv. Require authentication but no encryption.
		authPriv. Require authentication and encryption.
		Note: If you specify authentication, you must specify an encryption algorithm when you assign an SNMPv3 user to the group. If you also specify encryption, you must assign both an authentication and an encryption algorithm for each group member.
	*/
	Securitylevel string `json:"securitylevel,omitempty"`
	/**
	* Name of the configured SNMPv3 view that you want to bind to this SNMPv3 group. An SNMPv3 user bound to this group can access the subtrees that are bound to this SNMPv3 view as type INCLUDED, but cannot access the ones that are type EXCLUDED. If the Citrix ADC has multiple SNMPv3 view entries with the same name, all such entries are associated with the SNMPv3 group.
	*/
	Readviewname string `json:"readviewname,omitempty"`

	//------- Read only Parameter ---------;

	Storagetype string `json:"storagetype,omitempty"`
	Status string `json:"status,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
