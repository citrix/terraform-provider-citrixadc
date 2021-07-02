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
* Binding class showing the certkey that can be bound to sslvserver.
*/
type Sslvservercertkeybinding struct {
	/**
	* The name of the certificate key pair binding.
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
	* Port on which clear-text data is sent by the appliance to the server. Do not specify this parameter for SSL offloading with end-to-end encryption.
	*/
	Cleartextport int32 `json:"cleartextport,omitempty"`
	/**
	* CA certificate.
	*/
	Ca bool `json:"ca,omitempty"`
	/**
	* The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.
	*/
	Snicert bool `json:"snicert,omitempty"`
	/**
	* The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake
	*/
	Skipcaname bool `json:"skipcaname,omitempty"`
	/**
	* Name of the SSL virtual server.
	*/
	Vservername string `json:"vservername,omitempty"`


}