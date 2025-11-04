---
subcategory: "SSL"
---

# Resource: sslvserver

The sslvserver resource is used to set advanced SSL configuration for an SSL virtual server.


## Example usage

```hcl
resource "citrixadc_sslvserver" "tf_sslvserver" {
	cipherredirect = "ENABLED"
	cipherurl = "http://www.citrix.com"
	cleartextport = "80"
	clientauth = "ENABLED"
	clientcert = "Optional"
	hsts = "ENABLED"
	includesubdomains = "YES"
	maxage = "100"
	ocspstapling = "ENABLED"
	preload = "YES"
	sendclosenotify = "YES"
	sessreuse = "ENABLED"
	sesstimeout = "180"
	snienable = "ENABLED"
	sslredirect = "ENABLED"
	strictsigdigestcheck = "ENABLED"
	tls1 = "ENABLED"
	tls11 = "ENABLED"
	tls12 = "ENABLED"
	tls13 = "ENABLED"
	tls13sessionticketsperauthcontext = "7"
	zerorttearlydata = "ENABLED"
	vservername = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name        = "tf_vserver"
	servicetype = "SSL"
}
```


## Argument Reference

* `vservername` - (Required) Name of the SSL virtual server for which to set advanced configuration.
* `cleartextport` - (Optional) Port on which clear-text data is sent by the appliance to the server. Do not specify this parameter for SSL offloading with end-to-end encryption.
* `dh` - (Optional) State of Diffie-Hellman (DH) key exchange. Possible values: [ ENABLED, DISABLED ]
* `dhfile` - (Optional) Name of and, optionally, path to the DH parameter file, in PEM format, to be installed. /nsconfig/ssl/ is the default path.
* `dhcount` - (Optional) Number of interactions, between the client and the Citrix ADC, after which the DH private-public pair is regenerated. A value of zero (0) specifies infinite use (no refresh).
* `dhkeyexpsizelimit` - (Optional) This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits. Possible values: [ ENABLED, DISABLED ]
* `ersa` - (Optional) State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts. Possible values: [ ENABLED, DISABLED ]
* `ersacount` - (Optional) Refresh count for regeneration of the RSA public-key and private-key pair. Zero (0) specifies infinite usage (no refresh).
* `sessreuse` - (Optional) State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client. Possible values: [ ENABLED, DISABLED ]
* `sesstimeout` - (Optional) Time, in seconds, for which to keep the session active. Any session resumption request received after the timeout period will require a fresh SSL handshake and establishment of a new SSL session.
* `cipherredirect` - (Optional) State of Cipher Redirect. If cipher redirect is enabled, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a cipher mismatch between the virtual server or service and the client. Possible values: [ ENABLED, DISABLED ]
* `cipherurl` - (Optional) The redirect URL to be used with the Cipher Redirect feature.
* `sslv2redirect` - (Optional) State of SSLv2 Redirect. If SSLv2 redirect is enabled, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a protocol version mismatch between the virtual server or service and the client. Possible values: [ ENABLED, DISABLED ]
* `sslv2url` - (Optional) URL of the page to which to redirect the client in case of a protocol version mismatch. Typically, this page has a clear explanation of the error or an alternative location that the transaction can continue from.
* `clientauth` - (Optional) State of client authentication. If client authentication is enabled, the virtual server terminates the SSL handshake if the SSL client does not provide a valid certificate. Possible values: [ ENABLED, DISABLED ]
* `clientcert` - (Optional) Type of client authentication. If this parameter is set to MANDATORY, the appliance terminates the SSL handshake if the SSL client does not provide a valid certificate. With the OPTIONAL setting, the appliance requests a certificate from the SSL clients but proceeds with the SSL transaction even if the client presents an invalid certificate. Caution: Define proper access control policies before changing this setting to Optional. Possible values: [ Mandatory, Optional ]
* `sslredirect` - (Optional) State of HTTPS redirects for the SSL virtual server. For an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect. If SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break. Possible values: [ ENABLED, DISABLED ]
* `redirectportrewrite` - (Optional) State of the port rewrite while performing HTTPS redirect. If this parameter is ENABLED and the URL from the server does not contain the standard port, the port is rewritten to the standard. Possible values: [ ENABLED, DISABLED ]
* `ssl2` - (Optional) State of SSLv2 protocol support for the SSL Virtual Server. Possible values: [ ENABLED, DISABLED ]
* `ssl3` - (Optional) State of SSLv3 protocol support for the SSL Virtual Server. Note: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED. Possible values: [ ENABLED, DISABLED ]
* `tls1` - (Optional) State of TLSv1.0 protocol support for the SSL Virtual Server. Possible values: [ ENABLED, DISABLED ]
* `tls11` - (Optional) State of TLSv1.1 protocol support for the SSL Virtual Server. Possible values: [ ENABLED, DISABLED ]
* `tls12` - (Optional) State of TLSv1.2 protocol support for the SSL Virtual Server. Possible values: [ ENABLED, DISABLED ]
* `tls13` - (Optional) State of TLSv1.3 protocol support for the SSL Virtual Server. Possible values: [ ENABLED, DISABLED ]
* `dtls1` - (Optional) State of DTLSv1.0 protocol support for the SSL Virtual Server. Possible values: [ ENABLED, DISABLED ]
* `dtls12` - (Optional) State of DTLSv1.2 protocol support for the SSL Virtual Server. Possible values: [ ENABLED, DISABLED ]
* `snienable` - (Optional) State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net. Possible values: [ ENABLED, DISABLED ]
* `ocspstapling` - (Optional) State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values: ENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake. DISABLED: The appliance does not check the status of the server certificate. . Possible values: [ ENABLED, DISABLED ]
* `pushenctrigger` - (Optional) Trigger encryption on the basis of the PUSH flag value. Available settings function as follows: * ALWAYS - Any PUSH packet triggers encryption. * IGNORE - Ignore PUSH packet for triggering encryption. * MERGE - For a consecutive sequence of PUSH packets, the last PUSH packet triggers encryption. * TIMER - PUSH packet triggering encryption is delayed by the time defined in the set ssl parameter command or in the Change Advanced SSL Settings dialog box. Possible values: [ Always, Merge, Ignore, Timer ]
* `sendclosenotify` - (Optional) Enable sending SSL Close-Notify at the end of a transaction. Possible values: [ YES, NO ]
* `dtlsprofilename` - (Optional) Name of the DTLS profile whose settings are to be applied to the virtual server.
* `sslprofile` - (Optional) Name of the SSL profile that contains SSL settings for the virtual server.
* `hsts` - (Optional) State of HSTS protocol support for the SSL Virtual Server. Using HSTS, a server can enforce the use of an HTTPS connection for all communication with a client. Possible values: [ ENABLED, DISABLED ]
* `maxage` - (Optional) Set the maximum time, in seconds, in the strict transport security (STS) header during which the client must send only HTTPS requests to the server.
* `includesubdomains` - (Optional) Enable HSTS for subdomains. If set to Yes, a client must send only HTTPS requests for subdomains. Possible values: [ YES, NO ]
* `preload` - (Optional) Flag indicates the consent of the site owner to have their domain preloaded. Possible values: [ YES, NO ]
* `strictsigdigestcheck` - (Optional) Parameter indicating to check whether peer entity certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Citrix ADC. Possible values: [ ENABLED, DISABLED ]
* `zerorttearlydata` - (Optional) State of TLS 1.3 0-RTT early data support for the SSL Virtual Server. This setting only has an effect if resumption is enabled, as early data cannot be sent along with an initial handshake. Early application data has significantly different security properties - in particular there is no guarantee that the data cannot be replayed. Possible values: [ ENABLED, DISABLED ]
* `tls13sessionticketsperauthcontext` - (Optional) Number of tickets the SSL Virtual Server will issue anytime TLS 1.3 is negotiated, ticket-based resumption is enabled, and either (1) a handshake completes or (2) post-handhsake client auth completes. This value can be increased to enable clients to open multiple parallel connections using a fresh ticket for each connection. No tickets are sent if resumption is disabled.
* `dhekeyexchangewithpsk` - (Optional) Whether or not the SSL Virtual Server will require a DHE key exchange to occur when a PSK is accepted during a TLS 1.3 resumption handshake. A DHE key exchange ensures forward secrecy even in the event that ticket keys are compromised, at the expense of an additional round trip and resources required to carry out the DHE key exchange. If disabled, a DHE key exchange will be performed when a PSK is accepted but only if requested by the client. If enabled, the server will require a DHE key exchange when a PSK is accepted regardless of whether the client supports combined PSK-DHE key exchange. This setting only has an effect when resumption is enabled. Possible values: [ YES, NO ]
* `defaultsni` - (Optional) Default domain name supported by the SSL virtual server. The parameter is effective, when zero touch certificate management is active for the SSL virtual server i.e. no manual SNI cert or default server cert is bound to the v-server. For SSL transactions, when SNI is not presented by the client, server-certificate corresponding to the default SNI, if available in the cert-store, is selected else connection is terminated.
* `sslclientlogs` - (Optional) This parameter is used to enable or disable the logging of additional information, such as the Session ID and SNI names, from SSL handshakes to the audit logs.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver. It has the same value as the `vservername` attribute.


## Import

A sslvserver can be imported using its name, e.g.

```shell
terraform import citrixadc_sslvserver.tf_sslvserver tf_vserver
```
