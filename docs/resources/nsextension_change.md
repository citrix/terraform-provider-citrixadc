---
subcategory: "NS"
---

# Resource: nsextension_change

The nsextension_change resource reloads (recompiles) an existing Citrix ADC extension object from its stored source file. Use it after you have updated the extension's source on the appliance and want the running configuration to pick up the new code without recreating the `nsextension` object itself (for example, to apply a fix to a policy/protocol extension function during development).

~> **One-shot action.** This is an action resource: applying it reloads (recompiles) the extension from its stored source file; it does not manage a persistent object, so re-applying re-runs the reload. Each `terraform apply` that creates or replaces this resource performs the reload once, and changing `name` forces a new reload (replacement).


## Example usage

```hcl
resource "citrixadc_nsextension_change" "tf_nsextension_change" {
  name = "sample_extension"
}
```


## Argument Reference

* `name` - (Required) Name of the extension object to reload from its stored source file. Changing this value forces the resource to be recreated (re-running the reload action against the new extension).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the nsextension_change resource. It has the format `nsextension_change-<name>` (for example, `nsextension_change-sample_extension`).
