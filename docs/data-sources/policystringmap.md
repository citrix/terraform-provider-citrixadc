---
subcategory: "Policy"
---

# Data Source `policystringmap`

The policystringmap data source allows you to retrieve information about an existing policy string map configuration.

## Example usage

```terraform
data "citrixadc_policystringmap" "tf_stringmap" {
  name = "my_stringmap"
}

output "stringmap_name" {
  value = data.citrixadc_policystringmap.tf_stringmap.name
}

output "comment" {
  value = data.citrixadc_policystringmap.tf_stringmap.comment
}
```

## Argument Reference

* `name` - (Required) Unique name for the string map. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policystringmap resource. It has the same value as the `name` attribute.
* `comment` - Comments associated with the string map or key-value pair bound to this string map.
