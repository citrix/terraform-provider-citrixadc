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
* Configuration for view resource.
*/
type Snmpview struct {
	/**
	* Name for the SNMPv3 view. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a name that helps identify the SNMPv3 view.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose it in double or single quotation marks (for example, "my view" or 'my view').
	*/
	Name string `json:"name,omitempty"`
	/**
	* A particular branch (subtree) of the MIB tree that you want to associate with this SNMPv3 view. You must specify the subtree as an SNMP OID.
	*/
	Subtree string `json:"subtree,omitempty"`
	/**
	* Include or exclude the subtree, specified by the subtree parameter, in or from this view. This setting can be useful when you have included a subtree, such as A, in an SNMPv3 view and you want to exclude a specific subtree of A, such as B, from the SNMPv3 view.
	*/
	Type string `json:"type,omitempty"`

	//------- Read only Parameter ---------;

	Storagetype string `json:"storagetype,omitempty"`
	Status string `json:"status,omitempty"`

}
