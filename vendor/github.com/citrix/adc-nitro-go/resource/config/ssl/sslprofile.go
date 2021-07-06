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
* Configuration for SSL profile resource.
*/
type Sslprofile struct {
	/**
	* Name for the SSL profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created.
	*/
	Name string `json:"name,omitempty"`
	/**
	* Type of profile. Front end profiles apply to the entity that receives requests from a client. Backend profiles apply to the entity that sends client requests to a server.
	*/
	Sslprofiletype string `json:"sslprofiletype,omitempty"`
	/**
	* The name of the ssllogprofile.
	*/
	Ssllogprofile string `json:"ssllogprofile,omitempty"`
	/**
	* Number of interactions, between the client and the Citrix ADC, after which the DH private-public pair is regenerated. A value of zero (0) specifies refresh every time.
		This parameter is not applicable when configuring a backend profile. Allowed DH count values are 0 and >= 500.
	*/
	Dhcount int `json:"dhcount,omitempty"`
	/**
	* State of Diffie-Hellman (DH) key exchange.
		This parameter is not applicable when configuring a backend profile.
	*/
	Dh string `json:"dh,omitempty"`
	/**
	* The file name and path for the DH parameter.
	*/
	Dhfile string `json:"dhfile,omitempty"`
	/**
	* State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts.
		This parameter is not applicable when configuring a backend profile.
	*/
	Ersa string `json:"ersa,omitempty"`
	/**
	* The  refresh  count  for the re-generation of RSA public-key and private-key pair.
	*/
	Ersacount int `json:"ersacount,omitempty"`
	/**
	* State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.
	*/
	Sessreuse string `json:"sessreuse,omitempty"`
	/**
	* The Session timeout value in seconds.
	*/
	Sesstimeout int `json:"sesstimeout,omitempty"`
	/**
	* State of Cipher Redirect. If this parameter is set to ENABLED, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a cipher mismatch between the virtual server or service and the client.
		This parameter is not applicable when configuring a backend profile.
	*/
	Cipherredirect string `json:"cipherredirect,omitempty"`
	/**
	* The redirect URL to be used with the Cipher Redirect feature.
	*/
	Cipherurl string `json:"cipherurl,omitempty"`
	/**
	* State of client authentication. In service-based SSL offload, the service terminates the SSL handshake if the SSL client does not provide a valid certificate.
		This parameter is not applicable when configuring a backend profile.
	*/
	Clientauth string `json:"clientauth,omitempty"`
	/**
	* The rule for client certificate requirement in client authentication.
	*/
	Clientcert string `json:"clientcert,omitempty"`
	/**
	* This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits.
	*/
	Dhkeyexpsizelimit string `json:"dhkeyexpsizelimit,omitempty"`
	/**
	* State of HTTPS redirects for the SSL service.
		For an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect.
		If SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break.
		This parameter is not applicable when configuring a backend profile.
	*/
	Sslredirect string `json:"sslredirect,omitempty"`
	/**
	* State of the port rewrite while performing HTTPS redirect. If this parameter is set to ENABLED, and the URL from the server does not contain the standard port, the port is rewritten to the standard.
	*/
	Redirectportrewrite string `json:"redirectportrewrite,omitempty"`
	/**
	* State of SSLv3 protocol support for the SSL profile.
		Note: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.
	*/
	Ssl3 string `json:"ssl3,omitempty"`
	/**
	* State of TLSv1.0 protocol support for the SSL profile.
	*/
	Tls1 string `json:"tls1,omitempty"`
	/**
	* State of TLSv1.1 protocol support for the SSL profile.
	*/
	Tls11 string `json:"tls11,omitempty"`
	/**
	* State of TLSv1.2 protocol support for the SSL profile.
	*/
	Tls12 string `json:"tls12,omitempty"`
	/**
	* State of TLSv1.3 protocol support for the SSL profile.
	*/
	Tls13 string `json:"tls13,omitempty"`
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
	* State of server authentication support for the SSL Backend profile.
	*/
	Serverauth string `json:"serverauth,omitempty"`
	/**
	* Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server.
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
	* Enable sending SSL Close-Notify at the end of a transaction.
	*/
	Sendclosenotify string `json:"sendclosenotify,omitempty"`
	/**
	* Port on which clear-text data is sent by the appliance to the server. Do not specify this parameter for SSL offloading with end-to-end encryption.
	*/
	Cleartextport int `json:"cleartextport,omitempty"`
	/**
	* Encoding method used to insert the subject or issuer's name in HTTP requests to servers.
	*/
	Insertionencoding string `json:"insertionencoding,omitempty"`
	/**
	* Deny renegotiation in specified circumstances. Available settings function as follows:
		* NO - Allow SSL renegotiation.
		* FRONTEND_CLIENT - Deny secure and nonsecure SSL renegotiation initiated by the client.
		* FRONTEND_CLIENTSERVER - Deny secure and nonsecure SSL renegotiation initiated by the client or the Citrix ADC during policy-based client authentication.
		* ALL - Deny all secure and nonsecure SSL renegotiation.
		* NONSECURE - Deny nonsecure SSL renegotiation. Allows only clients that support RFC 5746.
	*/
	Denysslreneg string `json:"denysslreneg,omitempty"`
	/**
	* Amount of data to collect before the data is pushed to the crypto hardware for encryption. For large downloads, a larger quantum size better utilizes the crypto resources.
	*/
	Quantumsize string `json:"quantumsize,omitempty"`
	/**
	* Enable strict CA certificate checks on the appliance.
	*/
	Strictcachecks string `json:"strictcachecks,omitempty"`
	/**
	* Maximum number of queued packets after which encryption is triggered. Use this setting for SSL transactions that send small packets from server to Citrix ADC.
	*/
	Encrypttriggerpktcount int `json:"encrypttriggerpktcount,omitempty"`
	/**
	* Insert PUSH flag into decrypted, encrypted, or all records. If the PUSH flag is set to a value other than 0, the buffered records are forwarded on the basis of the value of the PUSH flag. Available settings function as follows:
		0 - Auto (PUSH flag is not set.)
		1 - Insert PUSH flag into every decrypted record.
		2 -Insert PUSH flag into every encrypted record.
		3 - Insert PUSH flag into every decrypted and encrypted record.
	*/
	Pushflag int `json:"pushflag,omitempty"`
	/**
	* Host header check for SNI enabled sessions. If this check is enabled and the HTTP request does not contain the host header for SNI enabled sessions(i.e vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension), the request is dropped.
	*/
	Dropreqwithnohostheader string `json:"dropreqwithnohostheader,omitempty"`
	/**
	* Controls how the HTTP 'Host' header value is validated. These checks are performed only if the session is SNI enabled (i.e when vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension) and HTTP request contains 'Host' header.
		Available settings function as follows:
		CERT   - Request is forwarded if the 'Host' value is covered
		by the certificate used to establish this SSL session.
		Note: 'CERT' matching mode cannot be applied in
		TLS 1.3 connections established by resuming from a
		previous TLS 1.3 session. On these connections, 'STRICT'
		matching mode will be used instead.
		STRICT - Request is forwarded only if value of 'Host' header
		in HTTP is identical to the 'Server name' value passed
		in 'Client Hello' of the SSL connection.
		NO     - No validation is performed on the HTTP 'Host'
		header value.
	*/
	Snihttphostmatch string `json:"snihttphostmatch,omitempty"`
	/**
	* PUSH encryption trigger timeout value. The timeout value is applied only if you set the Push Encryption Trigger parameter to Timer in the SSL virtual server settings.
	*/
	Pushenctriggertimeout int `json:"pushenctriggertimeout,omitempty"`
	/**
	* Time, in milliseconds, after which encryption is triggered for transactions that are not tracked on the Citrix ADC because their length is not known. There can be a delay of up to 10ms from the specified timeout value before the packet is pushed into the queue.
	*/
	Ssltriggertimeout int `json:"ssltriggertimeout,omitempty"`
	/**
	* Certficates bound on the VIP are used for validating the client cert. Certficates came along with client cert are not used for validating the client cert
	*/
	Clientauthuseboundcachain string `json:"clientauthuseboundcachain,omitempty"`
	/**
	* Enable or disable transparent interception of SSL sessions.
	*/
	Sslinterception string `json:"sslinterception,omitempty"`
	/**
	* Enable or disable triggering the client renegotiation when renegotiation request is received from the origin server.
	*/
	Sslireneg string `json:"sslireneg,omitempty"`
	/**
	* Enable or disable OCSP check for origin server certificate.
	*/
	Ssliocspcheck string `json:"ssliocspcheck,omitempty"`
	/**
	* Maximum ssl session to be cached per dynamic origin server. A unique ssl session is created for each SNI received from the client on ClientHello and the matching session is used for server session reuse.
	*/
	Sslimaxsessperserver int `json:"sslimaxsessperserver,omitempty"`
	/**
	* This option enables the use of session tickets, as per the RFC 5077
	*/
	Sessionticket string `json:"sessionticket,omitempty"`
	/**
	* This option sets the life time of session tickets issued by NS in secs
	*/
	Sessionticketlifetime int `json:"sessionticketlifetime,omitempty"`
	/**
	* This option enables the use of session tickets, as per the RFC 5077
	*/
	Sessionticketkeyrefresh string `json:"sessionticketkeyrefresh,omitempty"`
	/**
	* Session ticket enc/dec key , admin can set it
	*/
	Sessionticketkeydata string `json:"sessionticketkeydata,omitempty"`
	/**
	* This option sets the life time of symm key used to generate session tickets issued by NS in secs
	*/
	Sessionkeylifetime int `json:"sessionkeylifetime,omitempty"`
	/**
	* This option sets the life time of symm key used to generate session tickets issued by NS in secs
	*/
	Prevsessionkeylifetime int `json:"prevsessionkeylifetime,omitempty"`
	/**
	* State of HSTS protocol support for the SSL profile. Using HSTS, a server can enforce the use of an HTTPS connection for all communication with a client
	*/
	Hsts string `json:"hsts,omitempty"`
	/**
	* Set the maximum time, in seconds, in the strict transport security (STS) header during which the client must send only HTTPS requests to the server
	*/
	Maxage int `json:"maxage,omitempty"`
	/**
	* Enable HSTS for subdomains. If set to Yes, a client must send only HTTPS requests for subdomains.
	*/
	Includesubdomains string `json:"includesubdomains,omitempty"`
	/**
	* Flag indicates the consent of the site owner to have their domain preloaded.
	*/
	Preload string `json:"preload,omitempty"`
	/**
	* This flag controls the processing of X509 certificate policies. If this option is Enabled, then the policy check in Client authentication will be skipped. This option can be used only when Client Authentication is Enabled and ClientCert is set to Mandatory
	*/
	Skipclientcertpolicycheck string `json:"skipclientcertpolicycheck,omitempty"`
	/**
	* State of TLS 1.3 0-RTT early data support for the SSL Virtual Server. This setting only has an effect if resumption is enabled, as early data cannot be sent along with an initial handshake.
		Early application data has significantly different security properties - in particular there is no guarantee that the data cannot be replayed.
	*/
	Zerorttearlydata string `json:"zerorttearlydata,omitempty"`
	/**
	* Number of tickets the SSL Virtual Server will issue anytime TLS 1.3 is negotiated, ticket-based resumption is enabled, and either (1) a handshake completes or (2) post-handhsake client auth completes.
		This value can be increased to enable clients to open multiple parallel connections using a fresh ticket for each connection.
		No tickets are sent if resumption is disabled.
	*/
	Tls13sessionticketsperauthcontext int `json:"tls13sessionticketsperauthcontext,omitempty"`
	/**
	* Whether or not the SSL Virtual Server will require a DHE key exchange to occur when a PSK is accepted during a TLS 1.3 resumption handshake.
		A DHE key exchange ensures forward secrecy even in the event that ticket keys are compromised, at the expense of an additional round trip and resources required to carry out the DHE key exchange.
		If disabled, a DHE key exchange will be performed when a PSK is accepted but only if requested by the client.
		If enabled, the server will require a DHE key exchange when a PSK is accepted regardless of whether the client supports combined PSK-DHE key exchange. This setting only has an effect when resumption is enabled.
	*/
	Dhekeyexchangewithpsk string `json:"dhekeyexchangewithpsk,omitempty"`
	/**
	* When set to YES, attempt to use the TLS Extended Master Secret (EMS, as
		described in RFC 7627) when negotiating TLS 1.0, TLS 1.1 and TLS 1.2
		connection parameters. EMS must be supported by both the TLS client and server
		in order to be enabled during a handshake. This setting applies to both
		frontend and backend SSL profiles.
	*/
	Allowextendedmastersecret string `json:"allowextendedmastersecret,omitempty"`
	/**
	* Application protocol supported by the server and used in negotiation of the protocol with the client. Possible values are HTTP1.1, HTTP2 and NONE. Default value is NONE which implies application protocol is not enabled hence remain unknown to the TLS layer. This parameter is relevant only if SSL connection is handled by the virtual server of the type SSL_TCP.
	*/
	Alpnprotocol string `json:"alpnprotocol,omitempty"`
	/**
	* The cipher group/alias/individual cipher configuration
	*/
	Ciphername string `json:"ciphername,omitempty"`
	/**
	* cipher priority
	*/
	Cipherpriority int `json:"cipherpriority,omitempty"`
	/**
	* Parameter indicating to check whether peer entity certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Citrix ADC.
	*/
	Strictsigdigestcheck string `json:"strictsigdigestcheck,omitempty"`

	//------- Read only Parameter ---------;

	Nonfipsciphers string `json:"nonfipsciphers,omitempty"`
	Crlcheck string `json:"crlcheck,omitempty"`
	Ocspcheck string `json:"ocspcheck,omitempty"`
	Snicert string `json:"snicert,omitempty"`
	Skipcaname string `json:"skipcaname,omitempty"`
	Invoke string `json:"invoke,omitempty"`
	Labeltype string `json:"labeltype,omitempty"`
	Service string `json:"service,omitempty"`
	Builtin string `json:"builtin,omitempty"`
	Feature string `json:"feature,omitempty"`
	Sslpfobjecttype string `json:"sslpfobjecttype,omitempty"`
	Ssliverifyservercertforreuse string `json:"ssliverifyservercertforreuse,omitempty"`

}
