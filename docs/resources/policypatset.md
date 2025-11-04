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
* `dynamic` - (Optional) This is used to populate internal patset information so that the patset can also be used dynamically in an expression. Here dynamically means the patset name can also be derived using an expression. For example for a given patset name "allow_test" it can be used dynamically as http.req.url.contains_any("allow_" + http.req.url.path.get(1)). This cannot be used with default patsets.
* `patsetfile` - (Optional) File which contains list of patterns that needs to be bound to the patset. A patsetfile cannot be associated with multiple patsets.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policypatset. It has the same value as the `name` attribute.

## Import

A policypatset can be imported using its name, e.g.

```shell
terraform import citrixadc_policypatset.tf_patset tf_patset
```