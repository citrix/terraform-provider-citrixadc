---
subcategory: "SSL"
---

# Data Source: sslprofile\_sslcertkey\_binding

The sslprofile\_sslcertkey\_binding data source allows you to retrieve information about the binding between an SSL profile and an SSL certificate key.


## Example usage

```terraform
data "citrixadc_sslprofile_sslcertkey_binding" "tf_binding" {
  name          = "tfUnit_sslprofile-hello"
  sslicacertkey = "tf_sslcertkey"
}

output "cipherpriority" {
  value = data.citrixadc_sslprofile_sslcertkey_binding.tf_binding.cipherpriority
}
```


## Argument Reference

* `name` - (Required) Name of the SSL profile.
* `sslicacertkey` - (Required) The certkey (CA certificate + private key) to be used for SSL interception.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `cipherpriority` - Priority of the cipher binding.
* `id` - The id of the sslprofile\_sslcertkey\_binding. It is the concatenation of the `name` and `sslicacertkey` attributes separated by a comma.
