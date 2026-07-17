---
subcategory: "NS"
---

# Resource: nsextension_change

The nsextension_change resource reloads (recompiles) an existing Citrix ADC extension object from its stored source file. Use it after you have updated the extension's source on the appliance and want the running configuration to pick up the new code without recreating the `nsextension` object itself (for example, to apply a fix to a policy/protocol extension function during development).

~> **One-shot action.** This resource maps to the NITRO change (reload) action, which NITRO exposes at `POST ?action=update` (the literal `change` verb is rejected by NITRO with errorcode 1240, so the provider invokes `update`). It does not create a persistent object on the appliance. Each `terraform apply` that creates or replaces this resource performs the reload once. There is no readable server-side object and no NITRO GET endpoint, so there is no corresponding data source: Read is a no-op, Delete only removes the resource from Terraform state, and changing `name` forces a new reload (replacement).


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

* `id` - The ID of the nsextension_change resource. It is a synthetic identifier with the format `nsextension_change-<name>` (for example, `nsextension_change-sample_extension`); it does not correspond to any object on the Citrix ADC.
