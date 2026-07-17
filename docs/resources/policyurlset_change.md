---
subcategory: "Policy"
---

# Resource: policyurlset_change

The policyurlset_change resource refreshes an existing URL set on the Citrix ADC by re-importing its entries from the source configured on the parent `policyurlset` object. Use it when the remote URL list backing a named url set has been updated and you want the appliance to pull the latest contents into the in-memory set without recreating the `policyurlset` itself.

~> **One-shot action.** This resource maps to the NITRO `update` action (`POST ?action=update`; the NITRO doc labels this section `change`, but the real HTTP action and CLI verb is `update`). It does not create a persistent object on the appliance. Each `terraform apply` that creates or replaces this resource performs the change once. There is no readable server-side object and no NITRO GET endpoint, so there is no corresponding data source: Read is a no-op, Delete only removes the resource from Terraform state, and changing `name` forces a new action (replacement).


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

* `id` - The ID of the policyurlset_change resource. It is a synthetic identifier with the format `policyurlset_change-<name>` (for example, `policyurlset_change-top_malware_urls`); it does not correspond to any object on the Citrix ADC.
