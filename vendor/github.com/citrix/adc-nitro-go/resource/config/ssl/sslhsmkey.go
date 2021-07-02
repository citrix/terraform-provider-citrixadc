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
* Configuration for HSM key resource.
*/
type Sslhsmkey struct {
	Hsmkeyname string `json:"hsmkeyname,omitempty"`
	/**
	* Type of HSM.
	*/
	Hsmtype string `json:"hsmtype,omitempty"`
	/**
	* Name of the key. optionally, for Thales, path to the HSM key file; /var/opt/nfast/kmdata/local/ is the default path. Applies when HSMTYPE is THALES or KEYVAULT.
	*/
	Key string `json:"key,omitempty"`
	/**
	* Serial number of the partition on which the key is present. Applies only to SafeNet HSM.
	*/
	Serialnum string `json:"serialnum,omitempty"`
	/**
	* Password for a partition. Applies only to SafeNet HSM.
	*/
	Password string `json:"password,omitempty"`
	/**
	* Name of keystore object representing HSM where key is stored. For example, name of keyvault object or azurekeyvault authentication object. Applies only to KEYVAULT type HSM.
	*/
	Keystore string `json:"keystore,omitempty"`

	//------- Read only Parameter ---------;

	State string `json:"state,omitempty"`

}
