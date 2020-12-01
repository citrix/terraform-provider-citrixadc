---
subcategory: "SSL"
---

# Resource: sslaction

The sslaction resource is used to create ssl actions.


## Example usage

```hcl
resource "citrixadc_sslaction" "tf_sslaction" {
    name                   = "tf_sslaction"
    clientauth             = "DOCLIENTAUTH"
    clientcertverification = "Mandatory"
}
```


## Argument Reference

* `name` - (Optional) Name for the SSL action.
* `clientauth` - (Optional) Perform client certificate authentication. Possible values: [ DOCLIENTAUTH, NOCLIENTAUTH ]
* `clientcertverification` - (Optional) Client certificate verification is mandatory or optional. Possible values: [ Mandatory, Optional ]
* `ssllogprofile` - (Optional) The name of the ssllogprofile.
* `clientcert` - (Optional) Insert the entire client certificate into the HTTP header of the request being sent to the web server. The certificate is inserted in ASCII (PEM) format. Possible values: [ ENABLED, DISABLED ]
* `certheader` - (Optional) Name of the header into which to insert the client certificate.
* `clientcertserialnumber` - (Optional) Insert the entire client serial number into the HTTP header of the request being sent to the web server. Possible values: [ ENABLED, DISABLED ]
* `certserialheader` - (Optional) Name of the header into which to insert the client serial number.
* `clientcertsubject` - (Optional) Insert the client certificate subject, also known as the distinguished name (DN), into the HTTP header of the request being sent to the web server. Possible values: [ ENABLED, DISABLED ]
* `certsubjectheader` - (Optional) Name of the header into which to insert the client certificate subject.
* `clientcerthash` - (Optional) Insert the certificate's signature into the HTTP header of the request being sent to the web server. The signature is the value extracted directly from the X.509 certificate signature field. All X.509 certificates contain a signature field. Possible values: [ ENABLED, DISABLED ]
* `certhashheader` - (Optional) Name of the header into which to insert the client certificate signature (hash).
* `clientcertfingerprint` - (Optional) Insert the certificate's fingerprint into the HTTP header of the request being sent to the web server. The fingerprint is derived by computing the specified hash value (SHA256, for example) of the DER-encoding of the client certificate. Possible values: [ ENABLED, DISABLED ]
* `certfingerprintheader` - (Optional) Name of the header into which to insert the client certificate fingerprint.
* `certfingerprintdigest` - (Optional) Digest algorithm used to compute the fingerprint of the client certificate. Possible values: [ SHA1, SHA224, SHA256, SHA384, SHA512 ]
* `clientcertissuer` - (Optional) Insert the certificate issuer details into the HTTP header of the request being sent to the web server. Possible values: [ ENABLED, DISABLED ]
* `certissuerheader` - (Optional) Name of the header into which to insert the client certificate issuer details.
* `sessionid` - (Optional) Insert the SSL session ID into the HTTP header of the request being sent to the web server. Every SSL connection that the client and the Citrix ADC share has a unique ID that identifies the specific connection. Possible values: [ ENABLED, DISABLED ]
* `sessionidheader` - (Optional) Name of the header into which to insert the Session ID.
* `cipher` - (Optional) Insert the cipher suite that the client and the Citrix ADC negotiated for the SSL session into the HTTP header of the request being sent to the web server. The appliance inserts the cipher-suite name, SSL protocol, export or non-export string, and cipher strength bit, depending on the type of browser connecting to the SSL virtual server or service (for example, Cipher-Suite: RC4- MD5 SSLv3 Non-Export 128-bit). Possible values: [ ENABLED, DISABLED ]
* `cipherheader` - (Optional) Name of the header into which to insert the name of the cipher suite.
* `clientcertnotbefore` - (Optional) Insert the date from which the certificate is valid into the HTTP header of the request being sent to the web server. Every certificate is configured with the date and time from which it is valid. Possible values: [ ENABLED, DISABLED ]
* `certnotbeforeheader` - (Optional) Name of the header into which to insert the date and time from which the certificate is valid.
* `clientcertnotafter` - (Optional) Insert the date of expiry of the certificate into the HTTP header of the request being sent to the web server. Every certificate is configured with the date and time at which the certificate expires. Possible values: [ ENABLED, DISABLED ]
* `certnotafterheader` - (Optional) Name of the header into which to insert the certificate's expiry date.
* `owasupport` - (Optional) If the appliance is in front of an Outlook Web Access (OWA) server, insert a special header field, FRONT-END-HTTPS: ON, into the HTTP requests going to the OWA server. This header communicates to the server that the transaction is HTTPS and not HTTP. Possible values: [ ENABLED, DISABLED ]
* `forward` - (Optional) This action takes an argument a vserver name, to this vserver one will be able to forward all the packets.
* `cacertgrpname` - (Optional) This action will allow to pick CA(s) from the specific CA group, to verify the client certificate.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslaction. It has the same value as the `name` attribute.


## Import

A sslaction can be imported using its name, e.g.

```shell
terraform import citrixadc_sslaction.tf_sslaction tf_sslaction
```
