---
subcategory: "System"
---

# Resource: systemhwerror_check

Runs a hardware error check on the Citrix ADC as a one-shot diagnostic action. Use this resource when you want to trigger an on-demand disk error scan (for example, as part of a maintenance or troubleshooting workflow) and have the result surface in the ADC's diagnostic logs.

~>
* This is an **action-only** resource: applying it runs the hardware/disk error check once. It is a side effect, not a managed object.
* Because `diskcheck` is immutable, changing it forces the resource to be destroyed and recreated, which re-runs the check.
* **Import is not meaningful** for this resource because there is no underlying queryable object.


## Example usage

```hcl
resource "citrixadc_systemhwerror_check" "tf_systemhwerror_check" {
  diskcheck = true
}
```


## Argument Reference

* `diskcheck` - (Required) Perform only disk error checking. Changing this value forces a new resource to be created (re-runs the check).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemhwerror_check resource. It is set to `systemhwerror_check`.
