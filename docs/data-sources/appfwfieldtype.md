---
subcategory: "Application Firewall"
---

# Data Source: appfwfieldtype

The appfwfieldtype data source allows you to retrieve information about an existing appfwfieldtype.


## Example usage

```terraform
data "citrixadc_appfwfieldtype" "tf_appfwfieldtype" {
  name = "my_field_type"
}

output "name" {
  value = data.citrixadc_appfwfieldtype.tf_appfwfieldtype.name
}

output "regex" {
  value = data.citrixadc_appfwfieldtype.tf_appfwfieldtype.regex
}

output "priority" {
  value = data.citrixadc_appfwfieldtype.tf_appfwfieldtype.priority
}
```


## Argument Reference

* `name` - (Required) Name for the field type.
Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the field type is added.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my field type" or 'my field type').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwfieldtype. It has the same value as the `name` attribute.
* `comment` - Comment describing the type of field that this field type is intended to match.
* `nocharmaps` - Will not show internal field types added as part of FieldFormat learn rules deployment.
* `priority` - Positive integer specifying the priority of the field type. A lower number specifies a higher priority. Field types are checked in the order of their priority numbers.
* `regex` - PCRE - format regular expression defining the characters and length allowed for this field type.

### Read-only appfwfieldtype metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_appfwfieldtype` resource). They are Computed/GET-only and are `null` when the appliance omits them.

* `builtin` - Flag to determine if the field type is built-in or not (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this configuration.
