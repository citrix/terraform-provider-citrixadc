---
subcategory: "SSL"
---

# Data Source: sslcrlfile

The sslcrlfile data source allows you to retrieve information about a Certificate Revocation List (CRL) file that has been imported onto the Citrix ADC.


## Example usage

```terraform
data "citrixadc_sslcrlfile" "example" {
  name = "crl1"
}

output "crl_src" {
  value = data.citrixadc_sslcrlfile.example.src
}
```


## Argument Reference

* `name` - (Required) Name of the imported CRL file to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslcrlfile. It has the same value as the `name` attribute.
* `src` - URL specifying the protocol, host, and path, including file name to the CRL file that was imported.
