---
subcategory: "Network"
---

# Data Source: vridparam

The vridparam data source allows you to retrieve the global VRRP (Virtual Router Redundancy Protocol) parameters configured on the Citrix ADC, such as the master-forwarding behavior and the VRRP hello and dead-peer detection intervals used in active-active high-availability deployments.

Because vridparam is a singleton (one global object), no lookup attribute is required.


## Example usage

```terraform
data "citrixadc_vridparam" "tf_vridparam" {
}

output "vrrp_hellointerval" {
  value = data.citrixadc_vridparam.tf_vridparam.hellointerval
}
```


## Argument Reference

This data source has no required lookup arguments; it always refers to the single global vridparam object.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the vridparam resource. It is the synthetic constant string `vridparam-config`.
* `sendtomaster` - Forward packets to the master node, in an active-active mode configuration, if the virtual server is in the backup state and sharing is disabled. Possible values: [ ENABLED, DISABLED ].
* `hellointerval` - Interval, in milliseconds, between VRRP advertisement messages sent to the peer node in active-active mode.
* `deadinterval` - Number of seconds after which a peer node in active-active mode is marked down if VRRP advertisements are not received from the peer node.
