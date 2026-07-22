---
subcategory: "Policy"
---

# Resource: policypatsetfile_change

The policypatsetfile_change resource re-imports (refreshes) an existing pattern-set file on the Citrix ADC from its backing file, updating the in-memory patset to match the current file contents. Use it when the underlying pattern-set file has been edited out of band and you want the appliance to reload it under the same name, without deleting and re-creating the patset.

~> **One-shot action.** This resource maps to the NITRO `update` action (CLI: `update policy patsetFile -name <name>`); it does not manage a persistent object, so each `terraform apply` that creates or replaces this resource performs the update once, and changing `name` forces a new update (replacement).


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
