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
* Binding class showing the certkey that can be bound to sslservice.
*/
type Sslservicecertkeybinding struct {
	/**
	* The certificate key pair binding.
	*/
	Certkeyname string `json:"certkeyname,omitempty"`
	/**
	* The clearTextPort settings.
	*/
	Cleartextport int32 `json:"cleartextport,omitempty"`
	/**
	* The state of the CRL check parameter. (Mandatory/Optional)
	*/
	Crlcheck string `json:"crlcheck,omitempty"`
	/**
	* Rule to use for the OCSP responder associated with the CA certificate during client authentication. If MANDATORY is specified, deny all SSL clients if the OCSP check fails because of connectivity issues with the remote OCSP server, or any other reason that prevents the OCSP check. With the OPTIONAL setting, allow SSL clients even if the OCSP check fails except when the client certificate is revoked.
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
	* The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting      for client certificate in a SSL handshake
	*/
	Skipcaname bool `json:"skipcaname,omitempty"`
	/**
	* Name of the SSL service for which to set advanced configuration.
	*/
	Servicename string `json:"servicename,omitempty"`


}