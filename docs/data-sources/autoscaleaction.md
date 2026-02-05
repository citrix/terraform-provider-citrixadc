---
subcategory: "Autoscale"
---

# Data Source `autoscaleaction`

The autoscaleaction data source allows you to retrieve information about an existing autoscale action.


## Example usage

```terraform
data "citrixadc_autoscaleaction" "tf_autoscaleaction" {
  name = "my_autoscaleaction"
}

output "type" {
  value = data.citrixadc_autoscaleaction.tf_autoscaleaction.type
}

output "profilename" {
  value = data.citrixadc_autoscaleaction.tf_autoscaleaction.profilename
}

output "vserver" {
  value = data.citrixadc_autoscaleaction.tf_autoscaleaction.vserver
}
```


## Argument Reference

* `name` - (Required) ActionScale action name.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the autoscaleaction. It has the same value as the `name` attribute.
* `parameters` - Parameters to use in the action.
* `profilename` - AutoScale profile name.
* `quiettime` - Time in seconds no other policy is evaluated or action is taken.
* `type` - The type of action.
* `vmdestroygraceperiod` - Time in minutes a VM is kept in inactive state before destroying.
* `vserver` - Name of the vserver on which autoscale action has to be taken.


## Import

A autoscaleaction can be imported using its name, e.g.

```shell
terraform import citrixadc_autoscaleaction.tf_autoscaleaction my_autoscaleaction
```
