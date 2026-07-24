---
subcategory: "NS"
---

# Resource: nsextension_change

This resource is used to reload an existing Citrix ADC extension from its stored source file.


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
