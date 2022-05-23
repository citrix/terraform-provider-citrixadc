---
subcategory: "NS"
---

# Resource: nsassignment

The nsassignment resource is used to create assignment resource.


## Example usage

```hcl
resource "citrixadc_nsvariable" "tf_nsvariable" {
  name          = "tf_nsvariable"
  type          = "text(20)"
  scope         = "global"
  iffull        = "undef"
  ifvaluetoobig = "undef"
  ifnovalue     = "init"
  comment       = "Testing"
}
resource "citrixadc_nsassignment" "tf_nsassignment" {
  name     = "tf_nsassignment"
  variable = join("", ["$", citrixadc_nsvariable.tf_nsvariable.name])
  set      = 1
  comment  = "Testing"
}
```


## Argument Reference

* `name` - (Required) Name for the assignment. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the assignment is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my assignment" or my assignment).
* `variable` - (Required) Left hand side of the assigment, of the form $variable-name (for a singleton variabled) or $variable-name[key-expression], where key-expression is an expression that evaluates to a text string and provides the key to select a map entry.
* `set` - (Optional) Right hand side of the assignment. The expression is evaluated and assigned to the left hand variable.
* `Add` - (Optional) Right hand side of the assignment. The expression is evaluated and added to the left hand variable.
* `sub` - (Optional) Right hand side of the assignment. The expression is evaluated and subtracted from the left hand variable.
* `append` - (Optional) Right hand side of the assignment. The expression is evaluated and appended to the left hand variable.
* `clear` - (Optional) Clear the variable value. Deallocates a text value, and for a map, the text key.
* `comment` - (Optional) Comment. Can be used to preserve information about this rewrite action.
* `newname` - (Optional) New name for the assignment. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my assignment" or my assignment). Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsassignment. It has the same value as the `name` attribute.


## Import

A nsassignment can be imported using its name, e.g.

```shell
terraform import citrixadc_nsassignment.tf_nsassignment tf_nsassignment
```
