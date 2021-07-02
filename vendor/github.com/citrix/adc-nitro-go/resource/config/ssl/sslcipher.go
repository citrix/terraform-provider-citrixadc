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

package ssl

/**
* Configuration for cipher resource.
*/
type Sslcipher struct {
	/**
	* Name for the user-defined cipher group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the cipher group is created.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my ciphergroup" or 'my ciphergroup').
	*/
	Ciphergroupname string `json:"ciphergroupname,omitempty"`
	/**
	* The individual cipher name(s), a user-defined cipher group, or a system predefined cipher alias that will be added to the  predefined cipher alias that will be added to the group cipherGroupName.
		If a cipher alias or a cipher group is specified, all the individual ciphers in the cipher alias or group will be added to the user-defined cipher group.
	*/
	Ciphgrpalias string `json:"ciphgrpalias,omitempty"`
	/**
	* Cipher name.
	*/
	Ciphername string `json:"ciphername,omitempty"`
	/**
	* This indicates priority assigned to the particular cipher
	*/
	Cipherpriority uint32 `json:"cipherpriority,omitempty"`
	/**
	* Name of the profile to which cipher is attached.
	*/
	Sslprofile string `json:"sslprofile,omitempty"`

}
