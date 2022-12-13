---
subcategory: "Spillover"
---

# Resource: spilloverpolicy

The spilloverpolicy resource is used to create spillover policy.


## Example usage

```hcl
resource "citrixadc_spilloveraction" "tf_spilloveraction" {
  name   = "my_spilloveraction"
  action = "SPILLOVER"
}
resource "citrixadc_spilloverpolicy" "tf_spilloverpolicy" {
  name    = "tf_spilloverpolicy"
  rule    = "true"
  action  = citrixadc_spilloveraction.tf_spilloveraction.name
  comment = "This is example of spilloverpolicy"
}
```


## Argument Reference

* `name` - (Required) Name of the spillover policy.
* `rule` - (Required) Expression to be used by the spillover policy.
* `action` - (Required) Action for the spillover policy. Action is created using add spillover action command.
* `comment` - (Optional) Any comments that you might want to associate with the spillover policy.
* `newname` - (Optional) New name for the spillover policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Choose a name that reflects the function that the policy performs. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy'). Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the spilloverpolicy. It has the same value as the `name` attribute.


## Import

A spilloverpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_spilloverpolicy.tf_spilloverpolicy tf_spilloverpolicy
```
