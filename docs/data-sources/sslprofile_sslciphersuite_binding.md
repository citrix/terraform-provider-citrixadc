---
subcategory: "SSL"
---

# Data Source: sslprofile_sslciphersuite_binding

The sslprofile_sslciphersuite_binding data source allows you to retrieve information about the binding between an SSL profile and a cipher suite on the Citrix ADC.


## Example Usage

```terraform
data "citrixadc_sslprofile_sslciphersuite_binding" "tf_binding" {
  name       = "tf_sslprofile"
  ciphername = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
}

output "cipherpriority" {
  value = data.citrixadc_sslprofile_sslciphersuite_binding.tf_binding.cipherpriority
}

output "description" {
  value = data.citrixadc_sslprofile_sslciphersuite_binding.tf_binding.description
}
```


## Argument Reference

* `name` - (Required) Name of the SSL profile.
* `ciphername` - (Required) The cipher group, alias, or individual cipher configuration.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile_sslciphersuite_binding. It is the concatenation of the `name` and `ciphername` attributes, formatted as `name:<name>,ciphername:<ciphername>`.
* `cipherpriority` - Cipher priority.
* `description` - The cipher suite description.
