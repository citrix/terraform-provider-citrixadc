---
subcategory: "Network"
---

# Resource: vridparam

This resource is used to manage the global VRRP (VRID) parameters.


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
