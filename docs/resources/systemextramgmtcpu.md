---
subcategory: "System"
---

# Resource: systemextramgmtcpu

The systemextramgmtcpu resource is used to enable or disable the extra management cpu on the target ADC.


## Example usage

```hcl
resource "citrixadc_systemextramgmtcpu" "tf_extramgmtcpu" {
    enabled = true
    reboot = true
}
```


## Argument Reference

* `enabled` - (Required) Boolean value indicating the desired state of the extra management CPU. Set to `true` to enable it.
* `reboot` - (Optional) Boolean value. Set to `true` to reboot after the application of the extra management cpu.
* `reachable_timeout` - (Optional) Total time to wait after reboot. Defaults to "10m".
* `reachable_poll_delay` - (Optional) Time to wait before the first poll after reboot. Defaults to "60s".
* `reachable_poll_interval` - (Optional) Interval between polls. Defaults to "60s".
* `reachable_poll_timeout` - (Optional) Timeout for a poll attempt. Default to "20s".


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemextramgmtcpu. It is a random string prefixed with `tf-systemextramgmtcpu-`
