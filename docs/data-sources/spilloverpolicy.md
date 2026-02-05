---
subcategory: "Spillover"
---

# Data Source `spilloverpolicy`

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
