---
subcategory: "SSL"
---

# Data Source: sslcertbundle

The sslcertbundle data source allows you to retrieve information about an existing certificate bundle imported on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_sslcertbundle" "example" {
  name = "web-cert-bundle"
}

output "src" {
  value = data.citrixadc_sslcertbundle.example.src
}
```


## Argument Reference

* `name` - (Required) Name of the imported certificate bundle to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcertbundle. It has the same value as the `name` attribute.
* `src` - URL specifying the protocol, host, and path, including file name, of the imported certificate bundle.
