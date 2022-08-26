---
subcategory: "autoscale"
---

# Resource: autoscaleaction

The `autoscaleaction` resource is used to create autoscaleaction.


## Example usage

```hcl
resource "citrixadc_autoscaleaction" "action1" {
    name = "action1"
    type = "SCALE_UP"
    profilename = "profile1"
    vserver = "server1"
    parameters = "abc123"
}
```


## Argument Reference

* `name` - (Optional) ActionScale action name.
* `parameters` - (Optional) Parameters to use in the action
* `profilename` - (Optional) AutoScale profile name.
* `quiettime` - (Optional) Time in seconds no other policy is evaluated or action is taken
* `type` - (Optional) The type of action.
* `vmdestroygraceperiod` - (Optional) Time in minutes a VM is kept in inactive state before destroying
* `vserver` - (Optional) Name of the vserver on which autoscale action has to be taken.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `autoscaleaction`. It has the same value as the `name` attribute.


## Import

A `autoscaleaction` can be imported using its name, e.g.

```shell
terraform import citrixadc_csaction.tf_csaction tf_csaction
```
