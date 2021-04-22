---
subcategory: "Utility"
---

# Resource: rebooter

The rebooter resource is used to reboot the target ADC.


## Example usage

```hcl
resource "citrixadc_rebooter" "tf_rebooter" {
  timestamp            = timestamp()
  warm                 = true
  wait_until_reachable = true
}
```


## Argument Reference

* `warm` - (Required) Restarts the Citrix ADC software without rebooting the underlying operating system. The session terminates and you must log on to the appliance after it has restarted. Note: This argument is required only for nCore appliances. Classic appliances ignore this argument.
* `timestamp` - (Required) Time when reboot happened. Used to force the operation again in case all other attributes remain the same.
* `wait_until_reachable` - (Required) Boolean flag. If set to true the resource will wait for the ADC to become reachable.
* `reachable_timeout` - (Optional) Time period to wait for the ADC to become reachable. Defaults to "10m".
* `reachable_poll_delay` - (Optional) Time period to wait before first poll after reboot. Defaults to "60s".
* `reachable_poll_interval` - (Optional) Time interval between polls. Defaults to "60s".
* `reachable_poll_timeout` - (Optional) Time period to wait for each poll request. Defaults to "20s".


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rebooter. It is a unique string prefixed with "tf-rebooter-"
