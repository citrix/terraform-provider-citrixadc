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
