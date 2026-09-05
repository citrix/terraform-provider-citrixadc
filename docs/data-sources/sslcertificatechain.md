---
subcategory: "SSL"
---

# Data Source: sslcertificatechain

The sslcertificatechain data source allows you to retrieve information about a certificate chain formed for a certificate-key pair on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_sslcertificatechain" "example" {
  certkeyname = "servercert1"
}

output "certkeyname" {
  value = data.citrixadc_sslcertificatechain.example.certkeyname
}
```


## Argument Reference

* `certkeyname` - (Required) Name of the certificate-key pair to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertificatechain. It has the same value as the `certkeyname` attribute.

### Read-only sslcertificatechain metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_sslcertificatechain` resource). They are Computed / GET-only. Any attribute the appliance does not return is `null`.

* `chainlinked` - Certkeys which are currently in the SSL certificate chain. A list of strings.
* `chainpossiblelinks` - Certkeys which can be in the SSL certificate chain. A list of strings.
* `chainissuer` - Name of the issuer.
* `chaincomplete` - Is set to `1` if the SSL certificate chain is complete.
