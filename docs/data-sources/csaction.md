---
subcategory: "Content Switching"
---

# Data Source `csaction`

The csaction data source allows you to retrieve information about content switching actions.


## Example usage

```terraform
data "citrixadc_csaction" "tf_csaction" {
  name = "my_csaction"
}

output "targetlbvserver" {
  value = data.citrixadc_csaction.tf_csaction.targetlbvserver
}

output "comment" {
  value = data.citrixadc_csaction.tf_csaction.comment
}
```


## Argument Reference

* `name` - (Required) Name for the content switching action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the content switching action is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Comments associated with this cs action.
* `newname` - New name for the content switching action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.
* `targetlbvserver` - Name of the load balancing virtual server to which the content is switched.
* `targetvserver` - Name of the VPN, GSLB or Authentication virtual server to which the content is switched.
* `targetvserverexpr` - Information about this content switching action.

## Attribute Reference

* `id` - The id of the csaction. It has the same value as the `name` attribute.


## Import

A csaction can be imported using its name, e.g.

```shell
terraform import citrixadc_csaction.tf_csaction my_csaction
```
