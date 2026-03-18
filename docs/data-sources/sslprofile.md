---
page_title: "citrixadc_sslprofile Data Source - terraform-provider-citrixadc"
subcategory: ""
---

# citrixadc_sslprofile (Data Source)

Data source to retrieve information about a specific SSL profile.

## Example Usage

```terraform
data "citrixadc_sslprofile" "example" {
  name = "my_sslprofile"
}
```

## Schema

### Required

- `name` (String) Name for the SSL profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created.

### Optional

- `allowextendedmastersecret` (String) When set to YES, attempt to use the TLS Extended Master Secret (EMS, as described in RFC 7627) when negotiating TLS 1.0, TLS 1.1 and TLS 1.2 connection parameters. EMS must be supported by both the TLS client and server in order to be enabled during a handshake. This setting applies to both frontend and backend SSL profiles.
- `allowunknownsni` (String) Controls how the handshake is handled when the server name extension does not match any of the bound certificates. These checks are performed only if the session is SNI enabled (i.e. when profile bound to vserver has SNIEnable and Client Hello arrived with SNI extension). Available settings function as follows: ENABLED - handshakes with an unknown SNI are allowed to continue, if a default cert is bound. DISLABLED - handshakes with an unknown SNI are not allowed to continue.
- `alpnprotocol` (String) Application protocol supported by the server and used in negotiation of the protocol with the client. Possible values are HTTP1.1, HTTP2 and NONE. Default value is NONE which implies application protocol is not enabled hence remain unknown to the TLS layer. This parameter is relevant only if SSL connection is handled by the virtual server of the type SSL_TCP.
- `ciphername` (String) The cipher group/alias/individual cipher configuration
- `cipherpriority` (Number) cipher priority
- `cipherredirect` (String) State of Cipher Redirect. If this parameter is set to ENABLED, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a cipher mismatch between the virtual server or service and the client. This parameter is not applicable when configuring a backend profile.
- `cipherurl` (String) The redirect URL to be used with the Cipher Redirect feature.
- `cleartextport` (Number) Port on which clear-text data is sent by the appliance to the server. Do not specify this parameter for SSL offloading with end-to-end encryption.
- `clientauth` (String) State of client authentication. In service-based SSL offload, the service terminates the SSL handshake if the SSL client does not provide a valid certificate. This parameter is not applicable when configuring a backend profile.
- `clientauthuseboundcachain` (String) Certficates bound on the VIP are used for validating the client cert. Certficates came along with client cert are not used for validating the client cert
- `clientcert` (String) The rule for client certificate requirement in client authentication.
- `commonname` (String) Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server.
- `defaultsni` (String) Default domain name supported by the SSL virtual server. The parameter is effective, when zero touch certificate management is active for the SSL virtual server i.e. no manual SNI cert or default server cert is bound to the v-server. For SSL transactions, when SNI is not presented by the client, server-certificate corresponding to the default SNI, if available in the cert-store, is selected else connection is terminated.
- `denysslreneg` (String) Deny renegotiation in specified circumstances. Available settings function as follows: NO - Allow SSL renegotiation. FRONTEND_CLIENT - Deny secure and nonsecure SSL renegotiation initiated by the client. FRONTEND_CLIENTSERVER - Deny secure and nonsecure SSL renegotiation initiated by the client or the Citrix ADC during policy-based client authentication. ALL - Deny all secure and nonsecure SSL renegotiation. NONSECURE - Deny nonsecure SSL renegotiation. Allows only clients that support RFC 5746.
- `dh` (String) State of Diffie-Hellman (DH) key exchange. This parameter is not applicable when configuring a backend profile.
- `dhcount` (Number) Number of interactions, between the client and the Citrix ADC, after which the DH private-public pair is regenerated. A value of zero (0) specifies refresh every time. This parameter is not applicable when configuring a backend profile. Allowed DH count values are 0 and >= 500.
- `dhekeyexchangewithpsk` (String) Whether or not the SSL Virtual Server will require a DHE key exchange to occur when a PSK is accepted during a TLS 1.3 resumption handshake. A DHE key exchange ensures forward secrecy even in the event that ticket keys are compromised, at the expense of an additional round trip and resources required to carry out the DHE key exchange. If disabled, a DHE key exchange will be performed when a PSK is accepted but only if requested by the client. If enabled, the server will require a DHE key exchange when a PSK is accepted regardless of whether the client supports combined PSK-DHE key exchange. This setting only has an effect when resumption is enabled.
- `dhfile` (String) The file name and path for the DH parameter.
- `dhkeyexpsizelimit` (String) This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits.
- `dropreqwithnohostheader` (String) Host header check for SNI enabled sessions. If this check is enabled and the HTTP request does not contain the host header for SNI enabled sessions(i.e vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension), the request is dropped.
- `encryptedclienthello` (String) State of TLS 1.3 Encrypted Client Hello Support
- `encrypttriggerpktcount` (Number) Maximum number of queued packets after which encryption is triggered. Use this setting for SSL transactions that send small packets from server to Citrix ADC.
- `ersa` (String) State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts. This parameter is not applicable when configuring a backend profile.
- `ersacount` (Number) The refresh count for the re-generation of RSA public-key and private-key pair.
- `hsts` (String) State of HSTS protocol support for the SSL profile. Using HSTS, a server can enforce the use of an HTTPS connection for all communication with a client
- `includesubdomains` (String) Enable HSTS for subdomains. If set to Yes, a client must send only HTTPS requests for subdomains.
- `insertionencoding` (String) Encoding method used to insert the subject or issuer's name in HTTP requests to servers.
- `maxage` (Number) Set the maximum time, in seconds, in the strict transport security (STS) header during which the client must send only HTTPS requests to the server
- `maxrenegrate` (Number) Maximum number of renegotiation requests allowed, in one second, to each SSL entity to which this profile is bound. When set to 0, an unlimited number of renegotiation requests are allowed. Applicable only when Deny SSL renegotiation is set to a value other than ALL.
- `ocspstapling` (String) State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values: ENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake. DISABLED: The appliance does not check the status of the server certificate.
- `preload` (String) Flag indicates the consent of the site owner to have their domain preloaded.
- `prevsessionkeylifetime` (Number) This option sets the life time of symm key used to generate session tickets issued by NS in secs
- `pushenctrigger` (String) Trigger encryption on the basis of the PUSH flag value. Available settings function as follows: ALWAYS - Any PUSH packet triggers encryption. IGNORE - Ignore PUSH packet for triggering encryption. MERGE - For a consecutive sequence of PUSH packets, the last PUSH packet triggers encryption. TIMER - PUSH packet triggering encryption is delayed by the time defined in the set ssl parameter command or in the Change Advanced SSL Settings dialog box.
- `pushenctriggertimeout` (Number) PUSH encryption trigger timeout value. The timeout value is applied only if you set the Push Encryption Trigger parameter to Timer in the SSL virtual server settings.
- `pushflag` (Number) Insert PUSH flag into decrypted, encrypted, or all records. If the PUSH flag is set to a value other than 0, the buffered records are forwarded on the basis of the value of the PUSH flag. Available settings function as follows: 0 - Auto (PUSH flag is not set.) 1 - Insert PUSH flag into every decrypted record. 2 -Insert PUSH flag into every encrypted record. 3 - Insert PUSH flag into every decrypted and encrypted record.
- `quantumsize` (String) Amount of data to collect before the data is pushed to the crypto hardware for encryption. For large downloads, a larger quantum size better utilizes the crypto resources.
- `redirectportrewrite` (String) State of the port rewrite while performing HTTPS redirect. If this parameter is set to ENABLED, and the URL from the server does not contain the standard port, the port is rewritten to the standard.
- `sendclosenotify` (String) Enable sending SSL Close-Notify at the end of a transaction.
- `serverauth` (String) State of server authentication support for the SSL Backend profile.
- `sessionkeylifetime` (Number) This option sets the life time of symm key used to generate session tickets issued by NS in secs
- `sessionticket` (String) This option enables the use of session tickets, as per the RFC 5077
- `sessionticketkeydata` (String) Session ticket enc/dec key , admin can set it
- `sessionticketkeyrefresh` (String) This option enables the use of session tickets, as per the RFC 5077
- `sessionticketlifetime` (Number) This option sets the life time of session tickets issued by NS in secs
- `sessreuse` (String) State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.
- `sesstimeout` (Number) The Session timeout value in seconds.
- `skipclientcertpolicycheck` (String) This flag controls the processing of X509 certificate policies. If this option is Enabled, then the policy check in Client authentication will be skipped. This option can be used only when Client Authentication is Enabled and ClientCert is set to Mandatory
- `snienable` (String) State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.
- `snihttphostmatch` (String) Controls how the HTTP 'Host' header value is validated. These checks are performed only if the session is SNI enabled (i.e when vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension) and HTTP request contains 'Host' header. Available settings function as follows: CERT - Request is forwarded if the 'Host' value is covered by the certificate used to establish this SSL session. Note: 'CERT' matching mode cannot be applied in TLS 1.3 connections established by resuming from a previous TLS 1.3 session. On these connections, 'STRICT' matching mode will be used instead. STRICT - Request is forwarded only if value of 'Host' header in HTTP is identical to the 'Server name' value passed in 'Client Hello' of the SSL connection. NO - No validation is performed on the HTTP 'Host' header value.
- `ssl3` (String) State of SSLv3 protocol support for the SSL profile. Note: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.
- `sslclientlogs` (String) When enabled, NetScaler will log the session ID and SNI name during SSL handshakes on both the external and internal interfaces.
- `sslimaxsessperserver` (Number) Maximum ssl session to be cached per dynamic origin server. A unique ssl session is created for each SNI received from the client on ClientHello and the matching session is used for server session reuse.
- `sslinterception` (String) Enable or disable transparent interception of SSL sessions.
- `ssliocspcheck` (String) Enable or disable OCSP check for origin server certificate.
- `sslireneg` (String) Enable or disable triggering the client renegotiation when renegotiation request is received from the origin server.
- `ssllogprofile` (String) The name of the ssllogprofile.
- `sslprofiletype` (String) Type of profile. Front end profiles apply to the entity that receives requests from a client. Backend profiles apply to the entity that sends client requests to a server.
- `sslredirect` (String) State of HTTPS redirects for the SSL service. For an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect. If SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break. This parameter is not applicable when configuring a backend profile.
- `ssltriggertimeout` (Number) Time, in milliseconds, after which encryption is triggered for transactions that are not tracked on the Citrix ADC because their length is not known. There can be a delay of up to 10ms from the specified timeout value before the packet is pushed into the queue.
- `strictcachecks` (String) Enable strict CA certificate checks on the appliance.
- `strictsigdigestcheck` (String) Parameter indicating to check whether peer entity certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Citrix ADC.
- `tls1` (String) State of TLSv1.0 protocol support for the SSL profile.
- `tls11` (String) State of TLSv1.1 protocol support for the SSL profile.
- `tls12` (String) State of TLSv1.2 protocol support for the SSL profile.
- `tls13` (String) State of TLSv1.3 protocol support for the SSL profile.
- `tls13sessionticketsperauthcontext` (Number) Number of tickets the SSL Virtual Server will issue anytime TLS 1.3 is negotiated, ticket-based resumption is enabled, and either (1) a handshake completes or (2) post-handhsake client auth completes. This value can be increased to enable clients to open multiple parallel connections using a fresh ticket for each connection. No tickets are sent if resumption is disabled.
- `zerorttearlydata` (String) State of TLS 1.3 0-RTT early data support for the SSL Virtual Server. This setting only has an effect if resumption is enabled, as early data cannot be sent along with an initial handshake. Early application data has significantly different security properties - in particular there is no guarantee that the data cannot be replayed.

### Read-Only

- `id` (String) The ID of this resource.
