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
* Configuration for ech config resource.
*/
type Sslechconfig struct {
	/**
	* The ECH config name configured.
	*/
	Echconfigname string `json:"echconfigname,omitempty"`
	/**
	* The supported cipher suite that encrypts the client Hello Message.
	*/
	Echcipher string `json:"echcipher,omitempty"`
	/**
	* The name of the configured HPKE key
	*/
	Hpkekeyname string `json:"hpkekeyname,omitempty"`
	/**
	* The public name of ech config means FQDN or any string
	*/
	Echpublicname string `json:"echpublicname,omitempty"`
	/**
	* The config id of the ech config.
	*/
	Echconfigid *int `json:"echconfigid,omitempty"`
	/**
	* The version of ECH for which this configuration is used.
	*/
	Version *int `json:"version,omitempty"`

	//------- Read only Parameter ---------;

	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
