---
subcategory: "AAA"
---

# Resource: aaauser

The aaauser resource is used to create aaauser.


## Example usage

```hcl
resource "citrixadc_aaauser" "tf_aaauser" {
  username = "john"
  password = "my_pass"
}

```


## Argument Reference

* `username` - (Required) Name for the user. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the user is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my aaa user" or "my aaa user"). Minimum length =  1
* `password` - (Optional) Password with which the user logs on. Required for any user account that does not exist on an external authentication server. If you are not using an external authentication server, all user accounts must have a password. If you are using an external authentication server, you must provide a password for local user accounts that do not exist on the authentication server. Minimum length =  1
* `loggedin` - (Optional) Show whether the user is logged in or not.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaauser. It has the same value as the `username` attribute.


## Import

A aaauser can be imported using its name, e.g.

```shell
terraform import citrixadc_aaauser.tf_aaauser john
```
