---
subcategory: "Network"
---

# Data Source: vrid6

The vrid6 data source allows you to retrieve information about an IPv6 Virtual Router ID (VMAC6 address).

## Example usage

```terraform
data "citrixadc_vrid6" "example" {
  vrid6_id = 3
}

output "priority" {
  value = data.citrixadc_vrid6.example.priority
}

output "preemption" {
  value = data.citrixadc_vrid6.example.preemption
}
```

## Argument Reference

* `vrid6_id` - (Required) Integer value that uniquely identifies a VMAC6 address.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `all` - Remove all configured VMAC6 addresses from the Citrix ADC.
* `ownernode` - In a cluster setup, assign a cluster node as the owner of this VMAC address for IP based VRRP configuration. If no owner is configured, owner node is displayed as ALL and one node is dynamically elected as the owner.
* `preemption` - In an active-active mode configuration, make a backup VIP address the master if its priority becomes higher than that of a master VIP address bound to this VMAC address. If you disable pre-emption while a backup VIP address is the master, the backup VIP address remains master until the original master VIP's priority becomes higher than that of the current master.
* `preemptiondelaytimer` - Preemption delay time in seconds, in an active-active configuration. If any high priority node will come in network, it will wait for these many seconds before becoming master.
* `priority` - Base priority (BP), in an active-active mode configuration, which ordinarily determines the master VIP address.
* `sharing` - In an active-active mode configuration, enable the backup VIP address to process any traffic instead of dropping it.
* `trackifnumpriority` - Priority by which the Effective priority will be reduced if any of the tracked interfaces goes down in an active-active configuration.
* `tracking` - The effective priority (EP) value, relative to the base priority (BP) value in an active-active mode configuration. When EP is set to a value other than None, it is EP, not BP, which determines the master VIP address. Available settings function as follows: NONE - No tracking. EP = BP; ALL - If the status of all virtual servers is UP, EP = BP. Otherwise, EP = 0; ONE - If the status of at least one virtual server is UP, EP = BP. Otherwise, EP = 0; PROGRESSIVE - If the status of all virtual servers is UP, EP = BP. If the status of all virtual servers is DOWN, EP = 0. Otherwise EP = BP (1 - K/N), where N is the total number of virtual servers associated with the VIP address and K is the number of virtual servers for which the status is DOWN. Default: NONE.
