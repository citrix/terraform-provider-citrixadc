---
subcategory: "SSL"
---

# Data Source: sslservice_sslcipher_binding

The sslservice_sslcipher_binding data source allows you to retrieve information about the binding between an SSL service and a cipher, cipher group, or cipher alias.

## Example usage

```terraform
data "citrixadc_sslservice_sslcipher_binding" "example" {
  servicename = "tf_service"
  ciphername  = "HIGH"
}

output "cipheraliasname" {
  value = data.citrixadc_sslservice_sslcipher_binding.example.cipheraliasname
}

output "description" {
  value = data.citrixadc_sslservice_sslcipher_binding.example.description
}
```

## Argument Reference

* `servicename` - (Required) Name of the SSL service for which to set advanced configuration.
* `ciphername` - (Required) Name of the individual cipher, user-defined cipher group, or predefined (built-in) cipher alias.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_sslcipher_binding.
* `cipheraliasname` - The cipher group/alias/individual cipher configuration.
* `cipherdefaulton` - Flag indicating whether the bound cipher was the DEFAULT cipher, bound at boot time, or any other cipher from the CLI.
* `description` - The cipher suite description.
