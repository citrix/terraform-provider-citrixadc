---
subcategory: "SSL"
---

# Data Source `sslaction`

The sslaction data source allows you to retrieve information about an SSL action.


## Example usage

```terraform
data "citrixadc_sslaction" "tf_sslaction" {
  name = "my_sslaction"
}

output "clientauth" {
  value = data.citrixadc_sslaction.tf_sslaction.clientauth
}

output "clientcertverification" {
  value = data.citrixadc_sslaction.tf_sslaction.clientcertverification
}
```


## Argument Reference

* `name` - (Required) Name for the SSL action.

## Attribute Reference

The following attributes are available:

* `name` - Name for the SSL action.
* `clientauth` - Perform client certificate authentication.
* `clientcertverification` - Client certificate verification is mandatory or optional.
* `cacertgrpname` - This action will allow to pick CA(s) from the specific CA group, to verify the client certificate.
* `certfingerprintdigest` - Digest algorithm used to compute the fingerprint of the client certificate.
* `certfingerprintheader` - Name of the header into which to insert the client certificate fingerprint.
* `certhashheader` - Name of the header into which to insert the client certificate signature (hash).
* `certheader` - Name of the header into which to insert the client certificate.
* `certissuerheader` - Name of the header into which to insert the client certificate issuer details.
* `certnotafterheader` - Name of the header into which to insert the certificate's expiry date.
* `certnotbeforeheader` - Name of the header into which to insert the date and time from which the certificate is valid.
* `certserialheader` - Name of the header into which to insert the client serial number.
* `certsubjectheader` - Name of the header into which to insert the client certificate subject.
* `cipher` - Name of the cipher to use.
* `cipherurl` - URL of the cipher suite.
* `clientcert` - Insert the entire client certificate into the header in PEM format.
* `clientcertfingerprint` - Insert the certificate fingerprint into the header.
* `clientcerthash` - Insert the certificate signature into the header.
* `clientcertissuer` - Insert the certificate issuer details into the header.
* `clientcertnotafter` - Insert the certificate's expiry date into the header.
* `clientcertnotbefore` - Insert the date from which the certificate is valid into the header.
* `clientcertserialnumber` - Insert the client serial number into the header.
* `clientcertsubject` - Insert the client certificate subject into the header.
* `forward` - Action to perform.
* `owasupport` - If the appliance is in front of an Outlook Web Access (OWA) server, turn on this option.
* `sessionid` - Session ID to use.
* `sessionidhash` - Hash algorithm to generate hash of the session ID.
* `skipcommoninitialhandshake` - Skip initial SSL handshake with the Common Name.
* `snienable` - State of TLS Server Name Indication (SNI) processing on the virtual server. If SNI is enabled on a virtual server, a domain name (FQDN) can be associated with that virtual server for performing SNI-based certificate selections or to reject requests with an invalid SNI.
* `ssllogprofile` - The name of the ssllogprofile.
* `id` - The id of the sslaction. It is a system-generated identifier.
