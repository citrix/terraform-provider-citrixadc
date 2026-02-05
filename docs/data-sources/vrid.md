---
subcategory: "Network"
---

# Data Source: vrid

The vrid data source allows you to retrieve information about a Virtual Router ID (VMAC address).

## Example usage

```terraform
data "citrixadc_vrid" "example" {
  vrid_id = 3
}

output "priority" {
  value = data.citrixadc_vrid.example.priority
}

output "preemption" {
  value = data.citrixadc_vrid.example.preemption
}
```

## Argument Reference

* `vrid_id` - (Required) Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of 00:00:5e:00:01:<VRID>. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is 00:00:5e:00:01:3c, where 3c is the hexadecimal representation of 60.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `all` - Remove all the configured VMAC addresses from the Citrix ADC.
* `ownernode` - In a cluster setup, assign a cluster node as the owner of this VMAC address for IP based VRRP configuration. If no owner is configured, owner node is displayed as ALL and one node is dynamically elected as the owner.
* `preemption` - In an active-active mode configuration, make a backup VIP address the master if its priority becomes higher than that of a master VIP address bound to this VMAC address. If you disable pre-emption while a backup VIP address is the master, the backup VIP address remains master until the original master VIP's priority becomes higher than that of the current master.
* `preemptiondelaytimer` - Preemption delay time, in seconds, in an active-active configuration. If any high priority node will come in network, it will wait for these many seconds before becoming master.
* `priority` - Base priority (BP), in an active-active mode configuration, which ordinarily determines the master VIP address.
* `sharing` - In an active-active mode configuration, enable the backup VIP address to process any traffic instead of dropping it.
* `trackifnumpriority` - Priority by which the Effective priority will be reduced if any of the tracked interfaces goes down in an active-active configuration.
* `tracking` - The effective priority (EP) value, relative to the base priority (BP) value in an active-active mode configuration. When EP is set to a value other than None, it is EP, not BP, which determines the master VIP address. Available settings function as follows: NONE - No tracking. EP = BP; ALL - If the status of all virtual servers is UP, EP = BP. Otherwise, EP = 0; ONE - If the status of at least one virtual server is UP, EP = BP. Otherwise, EP = 0; PROGRESSIVE - If the status of all virtual servers is UP, EP = BP. If the status of all virtual servers is DOWN, EP = 0. Otherwise EP = BP (1 - K/N), where N is the total number of virtual servers associated with the VIP address and K is the number of virtual servers for which the status is DOWN. Default: NONE.
