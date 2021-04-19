---
subcategory: "Policy"
---

# Resource: policypatset

The policypatset resource is used to create policy pattern sets.


## Example usage

```hcl
resource "citrixadc_policypatset" "tf_patset" {
    name = "tf_patset"
    comment = "some comment"
}

resource "citrixadc_policypatset_pattern_binding" "tf_bind1" {
    name = citrixadc_policypatset.tf_patset.name
    string = "Pattern"
}
```


## Argument Reference

* `name` - (Optional) Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (\_) character and must contain only alphanumeric and underscore characters. Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.
* `indextype` - (Optional) Index type. Possible values: [ Auto-generated, User-defined ]
* `comment` - (Optional) Any comments to preserve information about this patset or a pattern bound to this patset.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policypatset. It has the same value as the `name` attribute.
