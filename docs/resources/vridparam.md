---
subcategory: "Network"
---

# Resource: vridparam

Configures the global VRRP (Virtual Router Redundancy Protocol) parameters used in active-active high-availability deployments. These settings control whether a backup node forwards traffic to the master node, and the VRRP advertisement (hello) and dead-peer detection intervals that determine how quickly the Citrix ADC detects a failed peer and triggers a master takeover.

This is a singleton resource: a single global VRRP-parameter object always exists on the Citrix ADC. Applying this resource manages that global configuration; destroying it only removes it from Terraform state (the global object itself is not deleted from the appliance).


## Example usage

```hcl
resource "citrixadc_vridparam" "tf_vridparam" {
  sendtomaster  = "ENABLED"
  hellointerval = 800
  deadinterval  = 5
}
```


## Argument Reference

* `sendtomaster` - (Optional) Forward packets to the master node, in an active-active mode configuration, if the virtual server is in the backup state and sharing is disabled. Possible values: [ ENABLED, DISABLED ]. Defaults to `"DISABLED"`.
* `hellointerval` - (Optional) Interval, in milliseconds, between VRRP advertisement messages sent to the peer node in active-active mode. Minimum value = 200. Maximum value = 1000. Defaults to `1000`.
* `deadinterval` - (Optional) Number of seconds after which a peer node in active-active mode is marked down if VRRP advertisements are not received from the peer node. Minimum value = 1. Maximum value = 60. Defaults to `3`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the vridparam resource. Because this is a singleton, it is a synthetic constant string with the value `vridparam-config`.


## Import

Although this singleton always exists on the appliance and is created on first apply, it can be imported into Terraform state using its synthetic constant ID:

```shell
terraform import citrixadc_vridparam.tf_vridparam vridparam-config
```
