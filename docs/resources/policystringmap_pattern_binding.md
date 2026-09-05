---
subcategory: "Policy"
---

# Resource: policystringmap\_pattern\_binding

The policystringmap\_pattern\_binding resource is used to bind a key-value pattern to a policy string map.


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
    comment = "Some comment"
}
```


## Argument Reference

* `name` - (Required) Name of the string map to which to bind the key-value pair.
* `value` - (Required) Character string constituting the value associated with the key. This value is returned when processed data matches the associated key. Refer to the key parameter for details of the value character set.
* `key` - (Required) Character string constituting the key to be bound to the string map. The key is matched against the data processed by the operation that uses the string map. The default character set is ASCII. UTF-8 characters can be included if the character set is UTF-8. UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\xNN'. For example, the UTF-8 character 'ü' can be encoded as '\xC3\xBC'.
* `comment` - (Optional) Comments associated with the string map or key-value pair bound to this string map.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policystringmap_pattern_binding. It is the concatenation of the `name` and `key` attributes separated by a comma.


## Import

A policystringmap_pattern_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_policystringmap_pattern_binding.tf_bind1 tf_policystringmap,key1
```
