---
subcategory: "Policy"
---

# Resource: policypatset\_pattern\_biding

The policypatset\_pattern\_biding resource is used to bind patterns to a policy pattern set.


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

* `string` - (Required) String of characters that constitutes a pattern. For more information about the characters that can be used, refer to the character set parameter. Note: Minimum length for pattern sets used in rewrite actions of type REPLACE\_ALL, DELETE\_ALL, INSERT\_AFTER\_ALL, and INSERT\_BEFORE\_ALL, is three characters.
* `feature` - (Optional) The feature to be checked while applying this config.
* `name` - (Required) Name of the pattern set to which to bind the string.
* `charset` - (Optional)    Character set associated with the characters in the string. Note: UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\xNN'. For example, the UTF-8 character 'ü' can be encoded as '\xC3\xBC'. 
* `index` - (Optional) The index of the string associated with the patset. 
* `comment` - (Optional) Any comments to preserve information about this patset or a pattern bound to this patset.




## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policypatset. It is the concatenation of the `name` and `string` attributes separated by a comma.
