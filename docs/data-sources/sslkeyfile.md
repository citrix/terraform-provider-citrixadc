---
subcategory: "SSL"
---

# Data Source: sslkeyfile

The sslkeyfile data source allows you to retrieve information about a private key file that has been imported onto the Citrix ADC.


## Example usage

```terraform
data "citrixadc_sslkeyfile" "example" {
  name = "servercert1key"
}

output "keyfile_src" {
  value = data.citrixadc_sslkeyfile.example.src
}
```


## Argument Reference

* `name` - (Required) Name of the imported key file to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslkeyfile. It has the same value as the `name` attribute.
* `src` - URL specifying the protocol, host, and path, including file name, to the key file that was imported.

Note: The key file passphrase (`password`) is a secret and is never returned by the NITRO API, so it is not available through this data source.
