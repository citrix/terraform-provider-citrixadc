---
subcategory: "Utility"
---

# Resource: change_password

This resource is used to change the password of the Citrix ADC. default password reset operation.


## Example usage

```hcl
resource "citrixadc_change_password" "tf_change_password" {
  username                  = "nsroot"
  password                  = "secret"
  new_password              = "verysecret"
  first_time_password_reset = false
}
```


## Argument Reference

* `username` - (Required) User name for the operation.
* `password` - (Required) The default password.
* `new_password` - (Required) The new password
* `first_time_password_reset` - (Required) bool value.The value is `true` if the user wants to perform default password reset operation at first login, else `false` if the user wants to change the password at any point later. 

~>
│ Error: [ERROR] nitro-go: Failed to create resource of type login, name=, err=failed: 599 Netscaler specific error ({ "errorcode": 278, "message": "Invalid argument", "severity": "ERROR" })  
If the user observe above error than it may be possible that the value of `first_time_password_reset` is set to `true` by the user but if the user wants to change the password at any point later than it must be set `false`. So, change the value of `first_time_password_reset` to `false`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the password_resetter. It is a random string prefixed with "tf-change-password".
