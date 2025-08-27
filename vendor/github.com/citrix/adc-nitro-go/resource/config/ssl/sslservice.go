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
* Configuration for SSL service resource.
*/
type Sslservice struct {
	/**
	* Name of the SSL service.
	*/
	Servicename string `json:"servicename,omitempty"`
	/**
	* State of Diffie-Hellman (DH) key exchange. This parameter is not applicable when configuring a backend service.
	*/
	Dh string `json:"dh,omitempty"`
	/**
	* Name for and, optionally, path to the PEM-format DH parameter file to be installed. /nsconfig/ssl/ is the default path. This parameter is not applicable when configuring a backend service.
	*/
	Dhfile string `json:"dhfile,omitempty"`
	/**
	* Number of interactions, between the client and the Citrix ADC, after which the DH private-public pair is regenerated. A value of zero (0) specifies refresh every time. This parameter is not applicable when configuring a backend service. Allowed DH count values are 0 and >= 500.
	*/
	Dhcount int `json:"dhcount,omitempty"`
	/**
	* This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits.
	*/
	Dhkeyexpsizelimit string `json:"dhkeyexpsizelimit,omitempty"`
	/**
	* State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts.
		This parameter is not applicable when configuring a backend service.
	*/
	Ersa string `json:"ersa,omitempty"`
	/**
	* Refresh count for regeneration of RSA public-key and private-key pair. Zero (0) specifies infinite usage (no refresh).
		This parameter is not applicable when configuring a backend service.
	*/
	Ersacount int `json:"ersacount,omitempty"`
	/**
	* State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.
	*/
	Sessreuse string `json:"sessreuse,omitempty"`
	/**
	* Time, in seconds, for which to keep the session active. Any session resumption request received after the timeout period will require a fresh SSL handshake and establishment of a new SSL session.
	*/
	Sesstimeout int `json:"sesstimeout,omitempty"`
	/**
	* State of Cipher Redirect. If this parameter is set to ENABLED, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a cipher mismatch between the virtual server or service and the client.
		This parameter is not applicable when configuring a backend service.
	*/
	Cipherredirect string `json:"cipherredirect,omitempty"`
	/**
	* URL of the page to which to redirect the client in case of a cipher mismatch. Typically, this page has a clear explanation of the error or an alternative location that the transaction can continue from.
		This parameter is not applicable when configuring a backend service.
	*/
	Cipherurl string `json:"cipherurl,omitempty"`
	/**
	* State of SSLv2 Redirect. If this parameter is set to ENABLED, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a protocol version mismatch between the virtual server or service and the client.
		This parameter is not applicable when configuring a backend service.
	*/
	Sslv2redirect string `json:"sslv2redirect,omitempty"`
	/**
	* URL of the page to which to redirect the client in case of a protocol version mismatch. Typically, this page has a clear explanation of the error or an alternative location that the transaction can continue from.
		This parameter is not applicable when configuring a backend service.
	*/
	Sslv2url string `json:"sslv2url,omitempty"`
	/**
	* State of client authentication. In service-based SSL offload, the service terminates the SSL handshake if the SSL client does not provide a valid certificate.
		This parameter is not applicable when configuring a backend service.
	*/
	Clientauth string `json:"clientauth,omitempty"`
	/**
	* Type of client authentication. If this parameter is set to MANDATORY, the appliance terminates the SSL handshake if the SSL client does not provide a valid certificate. With the OPTIONAL setting, the appliance requests a certificate from the SSL clients but proceeds with the SSL transaction even if the client presents an invalid certificate.
		This parameter is not applicable when configuring a backend SSL service.
		Caution: Define proper access control policies before changing this setting to Optional.
	*/
	Clientcert string `json:"clientcert,omitempty"`
	/**
	* State of HTTPS redirects for the SSL service.
		For an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect.
		If SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break.
		This parameter is not applicable when configuring a backend service.
	*/
	Sslredirect string `json:"sslredirect,omitempty"`
	/**
	* State of the port rewrite while performing HTTPS redirect. If this parameter is set to ENABLED, and the URL from the server does not contain the standard port, the port is rewritten to the standard.
	*/
	Redirectportrewrite string `json:"redirectportrewrite,omitempty"`
	/**
	* State of SSLv2 protocol support for the SSL service.
		This parameter is not applicable when configuring a backend service.
	*/
	Ssl2 string `json:"ssl2,omitempty"`
	/**
	* State of SSLv3 protocol support for the SSL service.
		Note: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.
	*/
	Ssl3 string `json:"ssl3,omitempty"`
	/**
	* State of TLSv1.0 protocol support for the SSL service.
	*/
	Tls1 string `json:"tls1,omitempty"`
	/**
	* State of TLSv1.1 protocol support for the SSL service.
	*/
	Tls11 string `json:"tls11,omitempty"`
	/**
	* State of TLSv1.2 protocol support for the SSL service.
	*/
	Tls12 string `json:"tls12,omitempty"`
	/**
	* State of TLSv1.3 protocol support for the SSL service.
	*/
	Tls13 string `json:"tls13,omitempty"`
	/**
	* State of DTLSv1.0 protocol support for the SSL service.
	*/
	Dtls1 string `json:"dtls1,omitempty"`
	/**
	* State of DTLSv1.2 protocol support for the SSL service.
	*/
	Dtls12 string `json:"dtls12,omitempty"`
	/**
	* State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.
	*/
	Snienable string `json:"snienable,omitempty"`
	/**
	* State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values:
		ENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake.
		DISABLED: The appliance does not check the status of the server certificate. 
	*/
	Ocspstapling string `json:"ocspstapling,omitempty"`
	/**
	* State of server authentication support for the SSL service.
	*/
	Serverauth string `json:"serverauth,omitempty"`
	/**
	* Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server
	*/
	Commonname string `json:"commonname,omitempty"`
	/**
	* Trigger encryption on the basis of the PUSH flag value. Available settings function as follows:
		* ALWAYS - Any PUSH packet triggers encryption.
		* IGNORE - Ignore PUSH packet for triggering encryption.
		* MERGE - For a consecutive sequence of PUSH packets, the last PUSH packet triggers encryption.
		* TIMER - PUSH packet triggering encryption is delayed by the time defined in the set ssl parameter command or in the Change Advanced SSL Settings dialog box.
	*/
	Pushenctrigger string `json:"pushenctrigger,omitempty"`
	/**
	* Enable sending SSL Close-Notify at the end of a transaction
	*/
	Sendclosenotify string `json:"sendclosenotify,omitempty"`
	/**
	* Name of the DTLS profile that contains DTLS settings for the service.
	*/
	Dtlsprofilename string `json:"dtlsprofilename,omitempty"`
	/**
	* Name of the SSL profile that contains SSL settings for the service.
	*/
	Sslprofile string `json:"sslprofile,omitempty"`
	/**
	* Parameter indicating to check whether peer's certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Citrix ADC
	*/
	Strictsigdigestcheck string `json:"strictsigdigestcheck,omitempty"`
	/**
	* This parameter is used to enable or disable the logging of additional information, such as the Session ID and SNI name, from SSL handshakes to the audit logs.
	*/
	Sslclientlogs string `json:"sslclientlogs,omitempty"`

	//------- Read only Parameter ---------;

	Nonfipsciphers string `json:"nonfipsciphers,omitempty"`
	Service string `json:"service,omitempty"`
	Skipcaname string `json:"skipcaname,omitempty"`
	Dtlsflag string `json:"dtlsflag,omitempty"`
	Quicflag string `json:"quicflag,omitempty"`
	Skipcacertbundle string `json:"skipcacertbundle,omitempty"`
	Nextgenapiresource string `json:"_nextgenapiresource,omitempty"`

}
