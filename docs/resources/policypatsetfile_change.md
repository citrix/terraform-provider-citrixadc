---
subcategory: "Policy"
---

# Resource: policypatsetfile_change

This resource is used to reimport an existing pattern-set file on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_policypatsetfile_change" "tf_policypatsetfile_change" {
  name = "my_patset_file"
}
```


## Argument Reference

* `name` - (Required) Name to assign to the imported patset file. Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters. Changing this value forces the resource to be recreated (re-running the update action against the new patset file name).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the policypatsetfile_change resource. It has the format `policypatsetfile_change-<name>` (for example, `policypatsetfile_change-my_patset_file`).
