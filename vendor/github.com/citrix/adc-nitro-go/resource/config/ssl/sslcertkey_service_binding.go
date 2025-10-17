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
* Binding class showing the service that can be bound to sslcertkey.
*/
type Sslcertkeyservicebinding struct {
	/**
	* Service name to which the certificate key pair is bound.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* Vserver Id
	*/
	Data *int `json:"data,omitempty"`
	/**
	* Version.
	*/
	Version *int `json:"version,omitempty"`
	/**
	* Name of the certificate-key pair.
	*/
	Certkey string `json:"certkey,omitempty"`
	/**
	* Bind the certificate to the named SSL service or service group.
	*/
	Service bool `json:"service,omitempty"`
	/**
	* The name of the SSL service group to which the certificate-key pair needs to be bound. Use the "add servicegroup" command to create this service.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* The certificate-key pair being unbound is a Certificate Authority (CA) certificate. If you choose this option, the certificate-key pair is unbound from the list of CA certificates that were bound to the specified SSL virtual server or SSL service.
	*/
	Ca bool `json:"ca,omitempty"`


}