---
subcategory: "NS"
---

# citrixadc_nsassignment (Data Source)

This data source retrieves information about a specific nsassignment resource.

An nsassignment is used to manage variable assignments in the Citrix ADC. It allows you to set, add, append, or subtract values to/from variables that can be used in policy expressions and other configurations.

**Note:** The nsassignment resource is currently not fully implemented in the Terraform provider. The Create and Update operations are commented out in the resource implementation. This datasource is provided for future compatibility but may have limited functionality until the resource is fully implemented.

## Example Usage

```hcl
data "citrixadc_nsassignment" "example" {
  name = "my_assignment"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the assignment to retrieve.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the nsassignment resource.
* `add` - Right hand side of the assignment. The expression is evaluated and added to the left hand variable.
* `append` - Right hand side of the assignment. The expression is evaluated and appended to the left hand variable.
* `clear` - Clear the variable value. Deallocates a text value, and for a map, the text key.
* `comment` - Comment. Can be used to preserve information about this rewrite action.
* `newname` - New name for the assignment. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `set` - Right hand side of the assignment. The expression is evaluated and assigned to the left hand variable.
* `sub` - Right hand side of the assignment. The expression is evaluated and subtracted from the left hand variable.
* `variable` - Left hand side of the assignment, of the form $variable-name (for a singleton variable) or $variable-name[key-expression], where key-expression is an expression that evaluates to a text string and provides the key to select a map entry.
