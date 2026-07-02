---
subcategory: "SSL"
---

# Data Source: ssldhfile

The ssldhfile data source allows you to retrieve information about a Diffie-Hellman (DH) parameters file that has been imported onto the Citrix ADC.


## Example usage

```terraform
data "citrixadc_ssldhfile" "example" {
  name = "dh2048"
}

output "dh_src" {
  value = data.citrixadc_ssldhfile.example.src
}
```


## Argument Reference

* `name` - (Required) Name of the imported DH file to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ssldhfile. It has the same value as the `name` attribute.
* `src` - URL specifying the protocol, host, and path, including file name, to the DH file that was imported.
