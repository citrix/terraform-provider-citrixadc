---
subcategory: "SSL"
---

# Resource: sslprofile_sslechconfig_binding

This resource is used to bind an Encrypted Client Hello (ECH) configuration to an SSL profile.


## Example usage

```hcl
resource "citrixadc_sslprofile" "tf_sslprofile" {
  name           = "tf_sslprofile"
  sslprofiletype = "FrontEnd"
}

resource "citrixadc_sslechconfig" "tf_echconfig" {
  name        = "tf_echconfig"
  publicname  = "ech.example.com"
  keyfilename = "ech_key"
}

resource "citrixadc_sslprofile_sslechconfig_binding" "tf_binding" {
  name           = citrixadc_sslprofile.tf_sslprofile.name
  echconfigname  = citrixadc_sslechconfig.tf_echconfig.name
  cipherpriority = 1
}
```


## Argument Reference

* `name` - (Required) Name of the SSL profile. Changing this attribute forces a new resource to be created.
* `echconfigname` - (Required) Name of the Encrypted Client Hello (ECH) configuration to bind. Changing this attribute forces a new resource to be created.
* `cipherpriority` - (Optional) Priority of the cipher binding. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile_sslechconfig_binding. It is the concatenation of the `name` and `echconfigname` attributes, formatted as `name:<name>,echconfigname:<echconfigname>`.


## Import

A sslprofile_sslechconfig_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslprofile_sslechconfig_binding.tf_binding name:tf_sslprofile,echconfigname:tf_echconfig
```
