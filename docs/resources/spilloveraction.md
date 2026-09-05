---
subcategory: "Spillover"
---

# Resource: spilloveraction

The spilloveraction resource is used to create spilloveraction.


## Example usage

```hcl
resource "citrixadc_spilloveraction" "tf_spilloveraction" {
  name         = "my_spilloveraction"
  action       = "SPILLOVER"
}
```


## Argument Reference

* `action` - (Required) Spillover action. Currently only type SPILLOVER is supported
* `name` - (Required) Name of the spillover action.
* `newname` - (Optional) New name for the spillover action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Choose a name that can be correlated with the function that the action performs. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the spilloveraction. It has the same value as the `name` attribute.


## Import

A spilloveraction can be imported using its name, e.g.

```shell
terraform import citrixadc_spilloveraction.tf_spilloveraction my_spilloveraction
```
