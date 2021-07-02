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
* Binding class showing the crldistribution that can be bound to sslcertkey.
*/
type Sslcertkeycrldistributionbinding struct {
	/**
	* Issuer name.
	*/
	Issuer string `json:"issuer,omitempty"`
	/**
	* Name of the certificate-key pair.
	*/
	Certkey string `json:"certkey,omitempty"`
	/**
	* The certificate-key pair being unbound is a Certificate Authority (CA) certificate. If you choose this option, the certificate-key pair is unbound from the list of CA certificates that were bound to the specified SSL virtual server or SSL service.
	*/
	Ca bool `json:"ca,omitempty"`


}