---
subcategory: "SSL"
---

# Data Source: citrixadc_sslservice

The sslservice data source allows you to retrieve information about an SSL service configuration.

## Example usage

```terraform
data "citrixadc_sslservice" "tf_sslservice_datasource" {
  servicename = "tf_service"
}

output "serverauth" {
  value = data.citrixadc_sslservice.tf_sslservice_datasource.serverauth
}

output "sessreuse" {
  value = data.citrixadc_sslservice.tf_sslservice_datasource.sessreuse
}

output "sesstimeout" {
  value = data.citrixadc_sslservice.tf_sslservice_datasource.sesstimeout
}
```

## Argument Reference

* `servicename` - (Required) Name of the SSL service.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `cipherredirect` - State of Cipher Redirect. If this parameter is set to ENABLED, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a cipher mismatch between the virtual server or service and the client. This parameter is not applicable when configuring a backend service.
* `cipherurl` - URL of the page to which to redirect the client in case of a cipher mismatch. Typically, this page has a clear explanation of the error or an alternative location that the transaction can continue from. This parameter is not applicable when configuring a backend service.
* `clientauth` - State of client authentication. In service-based SSL offload, the service terminates the SSL handshake if the SSL client does not provide a valid certificate. This parameter is not applicable when configuring a backend service.
* `clientcert` - Type of client authentication. If this parameter is set to MANDATORY, the appliance terminates the SSL handshake if the SSL client does not provide a valid certificate. With the OPTIONAL setting, the appliance requests a certificate from the SSL clients but proceeds with the SSL transaction even if the client presents an invalid certificate. This parameter is not applicable when configuring a backend SSL service. Caution: Define proper access control policies before changing this setting to Optional.
* `commonname` - Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server.
* `dh` - State of Diffie-Hellman (DH) key exchange. This parameter is not applicable when configuring a backend service.
* `dhcount` - Number of interactions, between the client and the Citrix ADC, after which the DH private-public pair is regenerated. A value of zero (0) specifies refresh every time. This parameter is not applicable when configuring a backend service. Allowed DH count values are 0 and >= 500.
* `dhfile` - Name for and, optionally, path to the PEM-format DH parameter file to be installed. /nsconfig/ssl/ is the default path. This parameter is not applicable when configuring a backend service.
* `dhkeyexpsizelimit` - This option enables the use of NIST recommended (NIST Special Publication 800-56A) bit size for private-key size. For example, for DH params of size 2048bit, the private-key size recommended is 224bits. This is rounded-up to 256bits.
* `dtls1` - State of DTLSv1.0 protocol support for the SSL service.
* `dtls12` - State of DTLSv1.2 protocol support for the SSL service.
* `dtlsprofilename` - Name of the DTLS profile that contains DTLS settings for the service.
* `ersa` - State of Ephemeral RSA (eRSA) key exchange. Ephemeral RSA allows clients that support only export ciphers to communicate with the secure server even if the server certificate does not support export clients. The ephemeral RSA key is automatically generated when you bind an export cipher to an SSL or TCP-based SSL virtual server or service. When you remove the export cipher, the eRSA key is not deleted. It is reused at a later date when another export cipher is bound to an SSL or TCP-based SSL virtual server or service. The eRSA key is deleted when the appliance restarts. This parameter is not applicable when configuring a backend service.
* `ersacount` - Refresh count for regeneration of RSA public-key and private-key pair. Zero (0) specifies infinite usage (no refresh). This parameter is not applicable when configuring a backend service.
* `id` - The id of the sslservice. It has the same value as the `servicename` attribute.
* `ocspstapling` - State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values: ENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake. DISABLED: The appliance does not check the status of the server certificate.
* `pushenctrigger` - Trigger encryption on the basis of the PUSH flag value. Available settings function as follows: ALWAYS - Any PUSH packet triggers encryption. IGNORE - Ignore PUSH packet for triggering encryption. MERGE - For a consecutive sequence of PUSH packets, the last PUSH packet triggers encryption. TIMER - PUSH packet triggering encryption is delayed by the time defined in the set ssl parameter command or in the Change Advanced SSL Settings dialog box.
* `redirectportrewrite` - State of the port rewrite while performing HTTPS redirect. If this parameter is set to ENABLED, and the URL from the server does not contain the standard port, the port is rewritten to the standard.
* `sendclosenotify` - Enable sending SSL Close-Notify at the end of a transaction.
* `serverauth` - State of server authentication support for the SSL service.
* `sessreuse` - State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client.
* `sesstimeout` - Time, in seconds, for which to keep the session active. Any session resumption request received after the timeout period will require a fresh SSL handshake and establishment of a new SSL session.
* `snienable` - State of the Server Name Indication (SNI) feature on the virtual server and service-based offload. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net.
* `ssl2` - State of SSLv2 protocol support for the SSL service. This parameter is not applicable when configuring a backend service.
* `ssl3` - State of SSLv3 protocol support for the SSL service. Note: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED.
* `sslclientlogs` - This parameter is used to enable or disable the logging of additional information, such as the Session ID and SNI name, from SSL handshakes to the audit logs.
* `sslprofile` - Name of the SSL profile that contains SSL settings for the service.
* `sslredirect` - State of HTTPS redirects for the SSL service. For an SSL session, if the client browser receives a redirect message, the browser tries to connect to the new location. However, the secure SSL session breaks if the object has moved from a secure site (https://) to an unsecure site (http://). Typically, a warning message appears on the screen, prompting the user to continue or disconnect. If SSL Redirect is ENABLED, the redirect message is automatically converted from http:// to https:// and the SSL session does not break. This parameter is not applicable when configuring a backend service.
* `sslv2redirect` - State of SSLv2 Redirect. If this parameter is set to ENABLED, you can configure an SSL virtual server or service to display meaningful error messages if the SSL handshake fails because of a protocol version mismatch between the virtual server or service and the client. This parameter is not applicable when configuring a backend service.
* `sslv2url` - URL of the page to which to redirect the client in case of a protocol version mismatch. Typically, this page has a clear explanation of the error or an alternative location that the transaction can continue from. This parameter is not applicable when configuring a backend service.
* `strictsigdigestcheck` - Parameter indicating to check whether peer's certificate during TLS1.2 handshake is signed with one of signature-hash combination supported by Citrix ADC.
* `tls1` - State of TLSv1.0 protocol support for the SSL service.
* `tls11` - State of TLSv1.1 protocol support for the SSL service.
* `tls12` - State of TLSv1.2 protocol support for the SSL service.
* `tls13` - State of TLSv1.3 protocol support for the SSL service.

## Import

A sslservice can be imported using its servicename, e.g.

```shell
terraform import citrixadc_sslservice.tf_sslservice tf_service
```
