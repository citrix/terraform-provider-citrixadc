---
subcategory: "SSL"
---

# Data Source: sslcertkeybundle

The sslcertkeybundle data source allows you to retrieve information about an existing certificate-key bundle configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_sslcertkeybundle" "example" {
  certkeybundlename = "web-certkey-bundle"
}

output "bundlefile" {
  value = data.citrixadc_sslcertkeybundle.example.bundlefile
}
```


## Argument Reference

* `certkeybundlename` - (Required) Name given to the certKeyBundle to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertkeybundle. It has the same value as the `certkeybundlename` attribute.
* `bundlefile` - Name of and, optionally, path to the X509 certificate bundle file that is used to form the certificate-key bundle.

Note: The `passplain` pass phrase is a secret and is never returned by the NITRO API, so it is not available through this data source.
