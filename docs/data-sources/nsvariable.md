---
subcategory: "NS"
---

# citrixadc_nsvariable (Data Source)

Data source for querying Citrix ADC variables. This data source retrieves information about variables configured on the ADC appliance, which can be used for policy expression evaluation.

## Example Usage

```hcl
data "citrixadc_nsvariable" "example" {
  name = "my_variable"
}

# Output variable attributes
output "variable_type" {
  value = data.citrixadc_nsvariable.example.type
}

output "variable_scope" {
  value = data.citrixadc_nsvariable.example.scope
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Variable name. This follows the same syntax rules as other expression entity names.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the nsvariable datasource.
* `comment` - Comments associated with this variable.
* `expires` - Value expiration in seconds. If the value is not referenced within the expiration period it will be deleted. 0 (the default) means no expiration.
* `iffull` - Action to perform if an assignment to a map exceeds its configured max-entries.
* `ifnovalue` - Action to perform if on a variable reference in an expression if the variable is single-valued and uninitialized or if the variable is a map and there is no value for the specified key.
* `ifvaluetoobig` - Action to perform if an value is assigned to a text variable that exceeds its configured max-size, or if a key is used that exceeds its configured max-size.
* `init` - Initialization value for this variable, to which a singleton variable or map entry will be set if it is referenced before an assignment action has assigned it a value.
* `scope` - Scope of the variable (global or transaction).
* `type` - Specification of the variable type (ulong, text, or map).

## Notes

NS variables are used in policy expressions to store and manipulate data during policy evaluation. They can be singleton variables or maps for more complex data structures.
