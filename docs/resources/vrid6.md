---
subcategory: "Network"
---

# Resource: vrid6

The vrid6 resource is used to create Virtual Router ID resource.


## Example usage

```hcl
resource "citrixadc_vrid6" "tf_vrid6" {
  vrid6_id             = 3
  priority             = 30
  preemption           = "DISABLED"
  sharing              = "DISABLED"
  tracking             = "NONE"
  trackifnumpriority   = 0
  preemptiondelaytimer = 0
}
```


## Argument Reference

* `vrid6_id` - (Required) Integer value that uniquely identifies a VMAC6 address.
* `all` - (Optional) Remove all configured VMAC6 addresses from the Citrix ADC.
* `ownernode` - (Optional) In a cluster setup, assign a cluster node as the owner of this VMAC address for IP based VRRP configuration. If no owner is configured, ow ner node is displayed as ALL and one node is dynamically elected as the owner.
* `preemption` - (Optional) In an active-active mode configuration, make a backup VIP address the master if its priority becomes higher than that of a master VIP address bound to this VMAC address.              If you disable pre-emption while a backup VIP address is the master, the backup VIP address remains master until the original master VIP's priority becomes higher than that of the current master.
* `preemptiondelaytimer` - (Optional) Preemption delay time in seconds, in an active-active configuration. If any high priority node will come in network, it will wait for these many seconds before becoming master.
* `priority` - (Optional) Base priority (BP), in an active-active mode configuration, which ordinarily determines the master VIP address.
* `sharing` - (Optional) In an active-active mode configuration, enable the backup VIP address to process any traffic instead of dropping it.
* `trackifnumpriority` - (Optional) Priority by which the Effective priority will be reduced if any of the tracked interfaces goes down in an active-active configuration.
* `tracking` - (Optional) The effective priority (EP) value, relative to the base priority (BP) value in an active-active mode configuration. When EP is set to a value other than None, it is EP, not BP, which determines the master VIP address. Available settings function as follows: * NONE - No tracking. EP = BP * ALL -  If the status of all virtual servers is UP, EP = BP. Otherwise, EP = 0. * ONE - If the status of at least one virtual server is UP, EP = BP. Otherwise, EP = 0. * PROGRESSIVE - If the status of all virtual servers is UP, EP = BP. If the status of all virtual servers is DOWN, EP = 0. Otherwise EP = BP (1 - K/N), where N is the total number of virtual servers associated with the VIP address and K is the number of virtual servers for which the status is DOWN. Default: NONE.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vrid6. It has the same value as the `vrid6_id` attribute.


## Import

A vrid6 can be imported using its vrid6_id, e.g.

```shell
terraform import citrixadc_vrid6.tf_vrid6 3
```
