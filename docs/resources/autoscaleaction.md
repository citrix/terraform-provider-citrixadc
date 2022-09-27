---
subcategory: "Autoscale"
---

# Resource: autoscaleaction

The autoscaleaction resource is used to create autoscaleaction.


## Example usage

```hcl
resource "citrixadc_autoscaleaction" "tf_autoscaleaction" {
  name        = "my_autoscaleaction"
  type        = "SCALE_UP"
  profilename = "my_profile"
  vserver     = "my_vserver"
  parameters  = "my_parameters"
}
```


## Argument Reference

* `name` - (Required) ActionScale action name.
* `parameters` - (Required) Parameters to use in the action
* `profilename` - (Required) AutoScale profile name.
* `type` - (Required) The type of action.
* `vserver` - (Required) Name of the vserver on which autoscale action has to be taken.
* `quiettime` - (Optional) Time in seconds no other policy is evaluated or action is taken
* `vmdestroygraceperiod` - (Optional) Time in minutes a VM is kept in inactive state before destroying


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the autoscaleaction. It has the same value as the `name` attribute.


## Import

A autoscaleaction can be imported using its name, e.g.

```shell
terraform import citrixadc_autoscaleaction.tfautoscaleaction my_autoscaleaction
```
