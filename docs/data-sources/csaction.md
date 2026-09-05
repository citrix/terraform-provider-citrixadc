---
subcategory: "Content Switching"
---

# Data Source: csaction

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
* `id` - The id of the csaction. It has the same value as the `name` attribute.

### Read-only csaction metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_csaction` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `hits` - The number of times the action has been taken.
* `referencecount` - The number of references to the action.
* `undefhits` - The number of times the action resulted in UNDEF.
* `builtin` - Flag to determine whether the action is built-in (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this config.
