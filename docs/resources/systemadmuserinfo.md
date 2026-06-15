---
subcategory: "System"
---

# Resource: systemadmuserinfo

Sets the admin-user name that the Citrix ADC records in its syslogs for the management session established between the ADC and Citrix Application Delivery Management (ADM) over SSH. Use this resource when you want the syslog entries generated during the ADC-to-ADM SSH session to be attributed to a specific adm-user account.

~>
* This is an **update-only** resource. NITRO exposes only a `set` (PUT) operation for `systemadmuserinfo`; there is no `add` and no GET endpoint.
* **Create and Update both issue the NITRO `set`** call to apply the configured `username`. They effectively rename the admin-user info used in the ADC-to-ADM SSH session.
* **Read is a no-op.** Because NITRO has no GET endpoint for this object, the value cannot be read back from the ADC, so drift detection is not possible. The provider preserves the prior plan/state value unchanged.
* **Destroy is state-only.** There is no NITRO delete operation; destroying the resource only removes it from Terraform state and does not revert the adm-user info on the ADC.
* **Import is not meaningful** for this resource because there is no underlying queryable object to import.


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

* `id` - The id of the systemadmuserinfo resource. It is a synthetic constant string `"systemadmuserinfo-config"`. The ID is purely a Terraform state handle, not a NITRO lookup key (there is no GET endpoint for this resource).
