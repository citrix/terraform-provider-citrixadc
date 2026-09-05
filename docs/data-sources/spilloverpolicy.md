---
subcategory: "Spillover"
---

# Data Source: spilloverpolicy

The spilloverpolicy data source allows you to retrieve information about a spillover policy.


## Example usage

```terraform
data "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
  name = "my_spilloverpolicy"
}

output "rule" {
  value = data.citrixadc_spilloverpolicy.tf_spilloverpolicy.rule
}

output "action" {
  value = data.citrixadc_spilloverpolicy.tf_spilloverpolicy.action
}
```


## Argument Reference

* `name` - (Required) Name of the spillover policy.

## Attribute Reference

The following attributes are available:

* `name` - Name of the spillover policy.
* `rule` - Expression to be used by the spillover policy.
* `action` - Action for the spillover policy. Action is created using add spillover action command.
* `comment` - Any comments that you might want to associate with the spillover policy.
* `newname` - New name for the spillover policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `id` - The id of the spilloverpolicy. It is a system-generated identifier.

### Read-only spilloverpolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_spilloverpolicy` resource) and are Computed. Any attribute the appliance does not return is `null`.

* `hits` - The number of times the policy has been hit.
* `undefhits` - Number of policy UNDEF hits.
* `builtin` - Flag to determine if the spillover policy is builtin or not (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this configuration.
