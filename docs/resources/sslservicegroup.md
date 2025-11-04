---
subcategory: "SSL"
---

# Resource: sslservicegroup

The sslservicegroup resource is used to set the advanced SSL configuration for an SSL service group.


## Example usage

```hcl
resource "citrixadc_sslservicegroup" "tf_sslservicegroup" {
	servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
	sesstimeout = 50
	sessreuse = "ENABLED"
	ssl3 = "ENABLED"
	snienable = "ENABLED"
	serverauth = "ENABLED"
	sendclosenotify = "YES"
	strictsigdigestcheck = "ENABLED"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
	servicegroupname = "tf_servicegroup"
	servicetype = "SSL"
}
```


## Argument Reference

* `servicegroupname` - (Required) Name of the SSL service group for which to set advanced configuration.
* `sslprofile` - (Optional) Name of the SSL profile that contains SSL settings for the Service Group.
* `sessreuse` - (Optional) State of session reuse. Establishing the initial handshake requires CPU-intensive public key encryption operations. With the ENABLED setting, session key exchange is avoided for session resumption requests received from the client. Possible values: [ ENABLED, DISABLED ]
* `sesstimeout` - (Optional) Time, in seconds, for which to keep the session active. Any session resumption request received after the timeout period will require a fresh SSL handshake and establishment of a new SSL session.
* `ssl3` - (Optional) State of SSLv3 protocol support for the SSL service group. Note: On platforms with SSL acceleration chips, if the SSL chip does not support SSLv3, this parameter cannot be set to ENABLED. Possible values: [ ENABLED, DISABLED ]
* `tls1` - (Optional) State of TLSv1.0 protocol support for the SSL service group. Possible values: [ ENABLED, DISABLED ]
* `tls11` - (Optional) State of TLSv1.1 protocol support for the SSL service group. Possible values: [ ENABLED, DISABLED ]
* `tls12` - (Optional) State of TLSv1.2 protocol support for the SSL service group. Possible values: [ ENABLED, DISABLED ]
* `tls13` - (Optional) State of TLSv1.3 protocol support for the SSL service group. Possible values: [ ENABLED, DISABLED ]
* `snienable` - (Optional) State of the Server Name Indication (SNI) feature on the service. SNI helps to enable SSL encryption on multiple domains on a single virtual server or service if the domains are controlled by the same organization and share the same second-level domain name. For example, *.sports.net can be used to secure domains such as login.sports.net and help.sports.net. Possible values: [ ENABLED, DISABLED ]
* `ocspstapling` - (Optional) State of OCSP stapling support on the SSL virtual server. Supported only if the protocol used is higher than SSLv3. Possible values: ENABLED: The appliance sends a request to the OCSP responder to check the status of the server certificate and caches the response for the specified time. If the response is valid at the time of SSL handshake with the client, the OCSP-based server certificate status is sent to the client during the handshake. DISABLED: The appliance does not check the status of the server certificate. Possible values: [ ENABLED, DISABLED ]
* `serverauth` - (Optional) State of server authentication support for the SSL service group. Possible values: [ ENABLED, DISABLED ]
* `commonname` - (Optional) Name to be checked against the CommonName (CN) field in the server certificate bound to the SSL server.
* `sendclosenotify` - (Optional) Enable sending SSL Close-Notify at the end of a transaction. Possible values: [ YES, NO ]
* `strictsigdigestcheck` - (Optional) Parameter indicating to check whether peer's certificate is signed with one of signature-hash combination supported by Citrix ADC. Possible values: [ ENABLED, DISABLED ]
* `sslclientlogs` - (Optional) This parameter is used to enable or disable the logging of additional information, such as the Session ID and SNI names, from SSL handshakes to the audit logs.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservicegroup. It has the same value as the `servicegroupname` attribute.


## Import

A sslservicegroup can be imported using its name, e.g.

```shell
terraform import citrixadc_sslservicegroup.tf_sslservicegroup tf_servicegroup
```
