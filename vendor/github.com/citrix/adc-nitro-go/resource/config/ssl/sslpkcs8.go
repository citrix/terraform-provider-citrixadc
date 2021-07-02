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
* Configuration for pkcs8 resource.
*/
type Sslpkcs8 struct {
	/**
	* Name for and, optionally, path to, the output file where the PKCS#8 format key file is stored. /nsconfig/ssl/ is the default path.
	*/
	Pkcs8file string `json:"pkcs8file,omitempty"`
	/**
	* Name of and, optionally, path to the input key file to be converted from PEM or DER format to PKCS#8 format. /nsconfig/ssl/ is the default path.
	*/
	Keyfile string `json:"keyfile,omitempty"`
	/**
	* Format in which the key file is stored on the appliance.
	*/
	Keyform string `json:"keyform,omitempty"`
	/**
	* Password to assign to the file if the key is encrypted. Applies only for PEM format files.
	*/
	Password string `json:"password,omitempty"`

}
