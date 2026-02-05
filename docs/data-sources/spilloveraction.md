---
subcategory: "Spillover"
---

# Data Source `spilloveraction`

The spilloveraction data source allows you to retrieve information about a spillover action.


## Example usage

```terraform
data "citrixadc_spilloveraction" "tf_spilloveraction" {
  name = "my_spilloveraction"
}

output "action" {
  value = data.citrixadc_spilloveraction.tf_spilloveraction.action
}
```


## Argument Reference

* `name` - (Required) Name of the spillover action.

## Attribute Reference

The following attributes are available:

* `name` - Name of the spillover action.
* `action` - Spillover action. Currently only type SPILLOVER is supported.
* `newname` - New name for the spillover action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `id` - The id of the spilloveraction. It is a system-generated identifier.
