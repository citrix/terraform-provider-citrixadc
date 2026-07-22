---
subcategory: "System"
---

# Resource: systemadmuserinfo

Sets the admin-user name that the Citrix ADC records in its syslogs for the management session established between the ADC and Citrix Application Delivery Management (ADM) over SSH. Use this resource when you want the syslog entries generated during the ADC-to-ADM SSH session to be attributed to a specific adm-user account.

~> **NOTE** Applying this resource sets the configured `username` for the admin-user info used in the ADC-to-ADM SSH session. Destroying the resource does not revert the adm-user info on the ADC, and importing it is not meaningful.


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
