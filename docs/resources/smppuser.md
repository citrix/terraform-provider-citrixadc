---
subcategory: "SMPP"
---

# Resource: smppuser

The smppuser resource is used to create smppuser.


## Example usage

```hcl
resource "citrixadc_smppuser" "tf_smppuser" {
  username = "user1"
  password = "abc"
}
```


## Argument Reference

* `password` - (Required) Password for binding to the SMPP server. Must be the same as the password specified in the SMPP server.
* `username` - (Required) Name of the SMPP user. Must be the same as the user name specified in the SMPP server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the smppuser. It has the same value as the `name` attribute.


## Import

A smppuser can be imported using its name, e.g.

```shell
terraform import citrixadc_smppuser.tf_smppuser user1
```
