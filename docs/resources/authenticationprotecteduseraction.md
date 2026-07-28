---
subcategory: "Authentication"
---

# Resource: authenticationprotecteduseraction

This resource is used to manage protected-user actions on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_authenticationprotecteduseraction" "tf_protecteduseraction" {
  name               = "tf_protecteduseraction"
  realmstr           = "EXAMPLE.COM"
  maxconcurrentusers = 10
}
```


## Argument Reference

* `name` - (Required) Name of the action to configure. Changing this forces a new resource to be created.
* `realmstr` - (Required) Kerberos Realm.
* `maxconcurrentusers` - (Optional) Max number of concurrent users allowed. Defaults to `8`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationprotecteduseraction. It has the same value as the `name` attribute.


## Import

An authenticationprotecteduseraction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationprotecteduseraction.tf_protecteduseraction tf_protecteduseraction
```
