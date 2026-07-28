---
subcategory: "SSL"
---

# Data Source: sslcacertbundle

The sslcacertbundle data source allows you to retrieve information about an existing CA certificate bundle configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_sslcacertbundle" "example" {
  cacertbundlename = "trusted-ca-bundle"
}

output "bundlefile" {
  value = data.citrixadc_sslcacertbundle.example.bundlefile
}
```


## Argument Reference

* `cacertbundlename` - (Required) Name given to the CA certbundle to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcacertbundle. It has the same value as the `cacertbundlename` attribute.
* `bundlefile` - Name of and, optionally, path to the X509 CA certificate bundle file that is used to form the cacertbundle entity.
