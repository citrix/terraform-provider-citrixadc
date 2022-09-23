---
subcategory: "DB"
---

# Resource: dbuser

The dbuser resource is used to create dbuser.


## Example usage

```hcl
resource "citrixadc_dbuser" "tf_dbuser" {
  username = "user1"
  password = "13456"
}
```


## Argument Reference

* `username` - (Required) Name of the database user. Must be the same as the user name specified in the database. Minimum length =  1
* `password` - (Optional) Password for logging on to the database. Must be the same as the password specified in the database. Minimum length =  1

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dbuser. It has the same value as the `username` attribute.


## Import

A dbuser can be imported using its name, e.g.

```shell
terraform import citrixadc_dbuser.tf_dbuser user1
```
