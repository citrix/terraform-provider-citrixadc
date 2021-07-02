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

package ns

/**
* Configuration for default encryption parameters resource.
*/
type Nsencryptionparams struct {
	/**
	* Cipher method (and key length) to be used to encrypt and decrypt content. The default value is AES256.
	*/
	Method string `json:"method,omitempty"`
	/**
	* The base64-encoded key generation number, method, and key value.
		Note:
		* Do not include this argument if you are changing the encryption method.
		* To generate a new key value for the current encryption method, specify an empty string \(""\) as the value of this parameter. The parameter is passed implicitly, with its automatically generated value, to the Citrix ADC packet engines even when it is not included in the command. Passing the parameter to the packet engines enables the appliance to save the key value to the configuration file and to propagate the key value to the secondary appliance in a high availability setup.
	*/
	Keyvalue string `json:"keyvalue,omitempty"`

}
