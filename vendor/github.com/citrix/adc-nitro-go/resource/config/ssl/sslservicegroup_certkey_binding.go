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
* Binding class showing the certkey that can be bound to sslservicegroup.
*/
type Sslservicegroupcertkeybinding struct {
	/**
	* The name of the certificate bound to the SSL service group.
	*/
	Certkeyname string `json:"certkeyname,omitempty"`
	/**
	* The state of the CRL check parameter. (Mandatory/Optional)
	*/
	Crlcheck string `json:"crlcheck,omitempty"`
	/**
	* The state of the OCSP check parameter. (Mandatory/Optional)
	*/
	Ocspcheck string `json:"ocspcheck,omitempty"`
	/**
	* CA certificate.
	*/
	Ca bool `json:"ca,omitempty"`
	/**
	* The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.
	*/
	Snicert bool `json:"snicert,omitempty"`
	/**
	* The name of the SSL service to which the SSL policy needs to be bound.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`


}