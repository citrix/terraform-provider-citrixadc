---
subcategory: "Policy"
---

# Data Source: policypatset_pattern_binding

The policypatset_pattern_binding data source allows you to retrieve information about a specific pattern binding in a policy pattern set.

## Example usage

```terraform
data "citrixadc_policypatset_pattern_binding" "example" {
  name   = "my_patset"
  string = "Pattern"
}

output "index" {
  value = data.citrixadc_policypatset_pattern_binding.example.index
}

output "charset" {
  value = data.citrixadc_policypatset_pattern_binding.example.charset
}

output "comment" {
  value = data.citrixadc_policypatset_pattern_binding.example.comment
}
```

## Argument Reference

* `name` - (Required) Name of the pattern set to which to bind the string.
* `string` - (Optional) String of characters that constitutes a pattern. For more information about the characters that can be used, refer to the character set parameter. Note: Minimum length for pattern sets used in rewrite actions of type REPLACE_ALL, DELETE_ALL, INSERT_AFTER_ALL, and INSERT_BEFORE_ALL, is three characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `charset` - Character set associated with the characters in the string. Note: UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\xNN'. For example, the UTF-8 character 'ü' can be encoded as '\xC3\xBC'.
* `comment` - Any comments to preserve information about this patset or a pattern bound to this patset.
* `feature` - The feature to be checked while applying this config.
* `id` - The id of the policypatset_pattern_binding. It is the concatenation of the `name` and `string` attributes separated by a comma.
* `index` - The index of the string associated with the patset.
