---
subcategory: "Utility"
---

# Resource: password_resetter

This resource is used to perform the default password reset operation.


## Example usage

```hcl
resource "citrixadc_password_resetter" "tf_resetter" {
    username = "nsroot"
    password = "nsroot"
    new_password = "newnsroot"
}
```


## Argument Reference

* `username` - (Required) User name for the operation.
* `password` - (Required) The default password.
* `new_password` - (Required) The new password


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the password_resetter. It is a random string prefixed with "tf-password-resetter-".
