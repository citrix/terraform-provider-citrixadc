---
subcategory: "System"
---

# Resource: systemhwerror_check

This resource is used to run a hardware error check on the Citrix ADC.


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
