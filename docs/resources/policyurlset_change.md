---
subcategory: "Policy"
---

# Resource: policyurlset_change

This resource is used to refresh an existing URL set by re-importing its entries.


## Example usage

```hcl
resource "citrixadc_policyurlset_change" "tf_policyurlset_change" {
  name = "top_malware_urls"
}
```


## Argument Reference

* `name` - (Required) Unique name of the url set to refresh. Maximum length: 127. Changing this value forces the resource to be recreated (re-running the change action against the new url set).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the policyurlset_change resource. It has the format `policyurlset_change-<name>` (for example, `policyurlset_change-top_malware_urls`).
