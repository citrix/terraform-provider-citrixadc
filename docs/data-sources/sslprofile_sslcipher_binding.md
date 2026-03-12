---
subcategory: "SSL"
---

# Data Source: sslprofile_sslcipher_binding

The sslprofile_sslcipher_binding data source allows you to retrieve information about the binding between an SSL profile and an SSL cipher.

## Example Usage

```terraform
data "citrixadc_sslprofile_sslcipher_binding" "tf_binding" {
  name       = "tf_sslprofile"
  ciphername = "HIGH"
}

output "cipheraliasname" {
  value = data.citrixadc_sslprofile_sslcipher_binding.tf_binding.cipheraliasname
}

output "cipherpriority" {
  value = data.citrixadc_sslprofile_sslcipher_binding.tf_binding.cipherpriority
}
```

## Argument Reference

* `name` - (Required) Name of the SSL profile.
* `ciphername` - (Required) Name of the cipher.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `cipheraliasname` - The name of the cipher group/alias/individual cipher bindings.
* `cipherpriority` - Cipher priority.
* `description` - The cipher suite description.
* `id` - The id of the sslprofile_sslcipher_binding. It is a system-generated identifier.
