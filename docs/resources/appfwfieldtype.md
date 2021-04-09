---
subcategory: "Application Firewall"
---

# Resource: appfwfieldtype

The `appfwfieldtype` resource is used to create Application Firwall FieldType resource.

## Example usage

``` hcl
resource "citrixadc_appfwfieldtype" "demo_appfwfieldtype" {
  name = "demo_appfwfieldtype"
  regex = "test_.*regex"
  priority = "100"
  # comment = "test comment"
}
```

## Argument Reference

* `name` - Name for the field type. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the field type is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my field type" or 'my field type').
* `regex` - (Optional) PCRE - format regular expression defining the characters and length allowed for this field type.
* `priority` - (Optional) Positive integer specifying the priority of the field type. A lower number specifies a higher priority. Field types are checked in the order of their priority numbers.
* `comment` - (Optional) Comment describing the type of field that this field type is intended to match.
* `nocharmaps` - (Optional) will not show internal field types added as part of FieldFormat learn rules deployment.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the `appfwfieldtype`. It has the same value as the `name` attribute.
