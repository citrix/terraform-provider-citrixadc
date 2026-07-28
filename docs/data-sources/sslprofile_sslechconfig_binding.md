---
subcategory: "SSL"
---

# Data Source: sslprofile_sslechconfig_binding

The sslprofile_sslechconfig_binding data source allows you to retrieve information about the binding between an SSL profile and an Encrypted Client Hello (ECH) configuration on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_sslprofile_sslechconfig_binding" "tf_binding" {
  name          = "tf_sslprofile"
  echconfigname = "tf_echconfig"
}

output "cipherpriority" {
  value = data.citrixadc_sslprofile_sslechconfig_binding.tf_binding.cipherpriority
}
```


## Argument Reference

* `name` - (Required) Name of the SSL profile.
* `echconfigname` - (Required) Name of the Encrypted Client Hello (ECH) configuration.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile_sslechconfig_binding. It is the concatenation of the `name` and `echconfigname` attributes, formatted as `name:<name>,echconfigname:<echconfigname>`.
* `cipherpriority` - Priority of the cipher binding.
