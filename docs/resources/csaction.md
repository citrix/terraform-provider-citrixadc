---
subcategory: "Content Switching"
---

# Resource: csaction

The csaction resource is used to create content switching actions.


## Example usage

```hcl
resource "citrixadc_csaction" "tf_csaction" {
  name            = "tf_csaction"
  targetlbvserver = "image_lb"
  comment         = "Forwards image requests to the image_lb"
}
```


## Argument Reference

* `name` - (Required) Name of the content switching action.
* `targetlbvserver` - (Optional) Name of the load balancing virtual server to which the content is switched.
* `targetvserver` - (Optional) Name of the VPN, GSLB or Authentication virtual server to which the content is switched.
* `targetvserverexpr` - (Optional) Expression that evaluates to the target load balancing virtual server to which the content is switched.
* `comment` - (Optional) Comment associated with this content switching action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the content switching action. It has the same value as the `name` attribute.


## Import

A csaction can be imported using its name, e.g.

```shell
terraform import citrixadc_csaction.tf_csaction tf_csaction
```
