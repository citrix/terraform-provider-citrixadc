---
subcategory: "Network"
---

# Resource: ptp

The ptp resource is used to create Precision Time Protocol resource.


## Example usage

```hcl
resource "citrixadc_ptp" "tf_ptp" {
  state = "ENABLE"
}
```


## Argument Reference

* `state` - (Required) Enables or disables Precision Time Protocol (PTP) on the appliance. If you disable PTP, make sure you enable Network Time Protocol (NTP) on the cluster. Possible values: [ enable, disable ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ptp. It is a unique string prefixed with "tf-ptp-"

