---
subcategory: "SSL"
---

# Resource: sslprofile_sslciphersuite_binding

This resource is used to manage the binding of cipher suites to an SSL profile.


## Example usage

```hcl
resource "citrixadc_sslprofile" "tf_sslprofile" {
  name        = "tf_sslprofile"
  sslprofiletype = "FrontEnd"
}

resource "citrixadc_sslprofile_sslciphersuite_binding" "tf_binding" {
  name           = citrixadc_sslprofile.tf_sslprofile.name
  ciphername     = "TLS1.2-ECDHE-RSA-AES256-GCM-SHA384"
  cipherpriority = 1
}
```


## Argument Reference

* `name` - (Required) Name of the SSL profile. Changing this attribute forces a new resource to be created.
* `ciphername` - (Required) The cipher group, alias, or individual cipher configuration to bind. Changing this attribute forces a new resource to be created.
* `cipherpriority` - (Optional) Cipher priority. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslprofile_sslciphersuite_binding. It is the concatenation of the `name` and `ciphername` attributes, formatted as `name:<name>,ciphername:<ciphername>`.
* `description` - (Read-only) The cipher suite description returned by the appliance.


## Import

A sslprofile_sslciphersuite_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_sslprofile_sslciphersuite_binding.tf_binding name:tf_sslprofile,ciphername:TLS1.2-ECDHE-RSA-AES256-GCM-SHA384
```
