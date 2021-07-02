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
* Configuration for SSL service group resource.
*/
type Sslservicegroup struct {
	/**
	* Name of the SSL service group for which to set advanced configuration.
	*/
	Servicegroupname string `json:"servicegroupname,omitempty"`
	/**
	* Name of the SSL profile that contains SSL settings for the Service Group.
	*/
	Sslprofile string `json:"sslprofile,omitempty"`
	/**
	* State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.
	*/
	Sessreuse string `json:"sessreuse,omitempty"`
	/**
	* Time, in seconds, for which to keep the session active. Any session resumption request received after the timeout period will require a fresh SSL handshake and establishment of a new SSL session.
	*/
	Sesstimeout uint32 `json:"sesstimeout,omitempty"`
	/**
	* State of SSLv3 protocol support for the SSL service group.
		Note: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.
	*/
	Ssl3 string `json:"ssl3,omitempty"`
	/**
	* State of TLSv1.0 protocol support for the SSL service group.
	*/
	Tls1 string `json:"tls1,omitempty"`
	/**
	* State of TLSv1.1 protocol support for the SSL service group.
	*/
	Tls11 string `json:"tls11,omitempty"`
	/**
	* State of TLSv1.2 protocol support for the SSL service group.
	*/
	Tls12 string `json:"tls12,omitempty"`
	/**
	* State of TLSv1.3 protocol support for the SSL service group.
	*/
	Tls13 string `json:"tls13,omitempty"`
	/**
	* State of the Server Name Indication (SNI) feature on the service. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.
	*/
	Snienable string `json:"snienable,omitempty"`
	/**
	* State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values:
		ENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake.
		DISABLED: The appliance does not check the status of the server certificate.
	*/
	Ocspstapling string `json:"ocspstapling,omitempty"`
	/**
	* State of server authentication support for the SSL service group.
	*/
	Serverauth string `json:"serverauth,omitempty"`
	/**
	* Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server
	*/
	Commonname string `json:"commonname,omitempty"`
	/**
	* Enable sending SSL Close-Notify at the end of a transaction
	*/
	Sendclosenotify string `json:"sendclosenotify,omitempty"`
	/**
	* Parameter indicating to check whether peer's certificate is signed with one of signature-hash combination supported by Citrix ADC
	*/
	Strictsigdigestcheck string `json:"strictsigdigestcheck,omitempty"`

	//------- Read only Parameter ---------;

	Dh string `json:"dh,omitempty"`
	Dhfile string `json:"dhfile,omitempty"`
	Dhcount string `json:"dhcount,omitempty"`
	Dhkeyexpsizelimit string `json:"dhkeyexpsizelimit,omitempty"`
	Ersa string `json:"ersa,omitempty"`
	Ersacount string `json:"ersacount,omitempty"`
	Cipherredirect string `json:"cipherredirect,omitempty"`
	Cipherurl string `json:"cipherurl,omitempty"`
	Sslv2redirect string `json:"sslv2redirect,omitempty"`
	Sslv2url string `json:"sslv2url,omitempty"`
	Clientauth string `json:"clientauth,omitempty"`
	Clientcert string `json:"clientcert,omitempty"`
	Sslredirect string `json:"sslredirect,omitempty"`
	Redirectportrewrite string `json:"redirectportrewrite,omitempty"`
	Nonfipsciphers string `json:"nonfipsciphers,omitempty"`
	Ssl2 string `json:"ssl2,omitempty"`
	Ocspcheck string `json:"ocspcheck,omitempty"`
	Crlcheck string `json:"crlcheck,omitempty"`
	Cleartextport string `json:"cleartextport,omitempty"`
	Servicename string `json:"servicename,omitempty"`
	Ca string `json:"ca,omitempty"`
	Snicert string `json:"snicert,omitempty"`

}
