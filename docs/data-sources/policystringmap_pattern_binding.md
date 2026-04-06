---
subcategory: "Policy"
---

# Data Source: policystringmap_pattern_binding

The policystringmap_pattern_binding data source allows you to retrieve information about a specific key-value pair binding in a policy string map.

## Example Usage

```terraform
data "citrixadc_policystringmap_pattern_binding" "example" {
  name = "my_stringmap"
  key  = "key1"
}

output "value" {
  value = data.citrixadc_policystringmap_pattern_binding.example.value
}

output "comment" {
  value = data.citrixadc_policystringmap_pattern_binding.example.comment
}
```

## Argument Reference

* `name` - (Required) Name of the string map to which to bind the key-value pair.
* `key` - (Required) Character string constituting the key to be bound to the string map. The key is matched against the data processed by the operation that uses the string map. The default character set is ASCII. UTF-8 characters can be included if the character set is UTF-8.  UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\xNN'. For example, the UTF-8 character 'ü' can be encoded as '\xC3\xBC'.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Comments associated with the string map or key-value pair bound to this string map.
* `id` - The id of the policystringmap_pattern_binding. It is a system-generated identifier.
* `value` - Character string constituting the value associated with the key. This value is returned when processed data matches the associated key. Refer to the key parameter for details of the value character set.
