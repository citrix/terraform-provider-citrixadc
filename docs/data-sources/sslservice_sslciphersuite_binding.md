---
subcategory: "SSL"
---

# Data Source: sslservice_sslciphersuite_binding

The sslservice_sslciphersuite_binding data source allows you to retrieve information about the binding between an SSL service and an SSL cipher suite.

## Example Usage

```terraform
data "citrixadc_sslservice_sslciphersuite_binding" "tf_sslservice_sslciphersuite_binding" {
  servicename = "tf_service"
  ciphername  = "tfAccsslcipher"
}

output "description" {
  value = data.citrixadc_sslservice_sslciphersuite_binding.tf_sslservice_sslciphersuite_binding.description
}

output "cipherdefaulton" {
  value = data.citrixadc_sslservice_sslciphersuite_binding.tf_sslservice_sslciphersuite_binding.cipherdefaulton
}
```

## Argument Reference

* `servicename` - (Required) Name of the SSL service for which to set advanced configuration.
* `ciphername` - (Required) The cipher group/alias/individual cipher configuration.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `cipherdefaulton` - Flag indicating whether the bound cipher was the DEFAULT cipher, bound at boot time, or any other cipher from the CLI.
* `description` - The cipher suite description.

## Attribute Reference

* `id` - The id of the sslservice_sslciphersuite_binding. It is a system-generated identifier.
