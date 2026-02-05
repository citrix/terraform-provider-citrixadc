---
subcategory: "Policy"
---

# Data Source `policypatset`

The policypatset data source allows you to retrieve information about an existing policy pattern set configuration.

## Example usage

```terraform
data "citrixadc_policypatset" "tf_patset" {
  name = "my_patset"
}

output "patset_name" {
  value = data.citrixadc_policypatset.tf_patset.name
}

output "comment" {
  value = data.citrixadc_policypatset.tf_patset.comment
}
```

## Argument Reference

* `name` - (Required) Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policypatset resource. It has the same value as the `name` attribute.
* `comment` - Any comments to preserve information about this patset or a pattern bound to this patset.
* `dynamic` - This is used to populate internal patset information so that the patset can also be used dynamically in an expression.
* `dynamiconly` - Shows only dynamic patsets when set true.
* `patsetfile` - File which contains list of patterns that needs to be bound to the patset. A patsetfile cannot be associated with multiple patsets.
