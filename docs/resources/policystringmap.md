---
subcategory: "Policy"
---

# Resource: policystringmap

The policystringmap resource is used to create policy string maps.


## Example usage

```hcl
resource "citrixadc_policystringmap" "tf_policystringmap" {
    name = "tf_policystringmap"
    comment = "Some comment"
}

resource "citrixadc_policystringmap_pattern_binding" "tf_bind1" {
    name = citrixadc_policystringmap.tf_policystringmap.name
    key = "key1"
    value = "value1"
}
```


## Argument Reference

* `name` - (Optional) Unique name for the string map. Not case sensitive. Must begin with an ASCII letter or underscore (\_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.
* `comment` - (Optional) Comments associated with the string map or key-value pair bound to this string map.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policystringmap. It has the same value as the `name` attribute.
