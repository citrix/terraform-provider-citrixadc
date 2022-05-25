---
subcategory: "Network"
---

# Resource: vridparam

The vridparam resource is used to configure VR ID parameter resource..


## Example usage

```hcl
resource "citrixadc_vridparam" "tf_vridparam" {
  sendtomaster  = "DISABLED"
  hellointerval = 1000
  deadinterval  = 3
}
```


## Argument Reference

* `sendtomaster` - (Optional) Forward packets to the master node, in an active-active mode configuration, if the virtual server is in the backup state and sharing is disabled. Possible values: [ ENABLED, DISABLED ]
* `hellointerval` - (Optional) Interval, in milliseconds, between vrrp advertisement messages sent to the peer node in active-active mode. Minimum value =  200 Maximum value =  1000
* `deadinterval` - (Optional) Number of seconds after which a peer node in active-active mode is marked down if vrrp advertisements are not received from the peer node. Minimum value =  1 Maximum value =  60


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vridparam. It is a unique string prefixed with "tf-vridparam-"

