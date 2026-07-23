---
subcategory: "System"
---

# Resource: systemadmuserinfo

This resource is used to manage the admin-user info logged for the ADC-to-ADM SSH session.


## Example usage

```hcl
resource "citrixadc_systemadmuserinfo" "tf_systemadmuserinfo" {
  username = "admuser1"
}
```


## Argument Reference

* `username` - (Required) Name of adm-user to log in syslogs.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemadmuserinfo resource. It is set to `systemadmuserinfo-config`.
