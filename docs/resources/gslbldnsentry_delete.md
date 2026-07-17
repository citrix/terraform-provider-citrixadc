---
subcategory: "GSLB"
---

# Resource: gslbldnsentry_delete

The gslbldnsentry_delete resource removes a runtime-learned LDNS (local DNS) entry from the Citrix ADC GSLB subsystem. When GSLB serves DNS requests, the ADC dynamically learns the LDNS servers that resolve on behalf of clients and uses these entries for proximity-based site selection. Use this resource to purge a stale or unwanted learned entry for a specific LDNS IP address (equivalent to the ADC CLI command `rm gslb ldnsentry <ipaddress>`).

~> **This is an action-only "delete-as-create" resource.** The NITRO API exposes only the `delete` verb for `gslbldnsentry` — there is no add, get, or update operation. As a result, **applying** this resource performs an imperative HTTP DELETE that removes the learned LDNS entry whose IP matches `ipaddress`. Read, Update, and Delete (Terraform destroy) are all no-ops on the appliance: there is nothing to read back, and destroying the Terraform resource does **not** re-add the entry — it only drops the resource from state. Because there is no NITRO GET endpoint, there is no corresponding data source. To remove a different LDNS entry, change `ipaddress`; because `ipaddress` forces replacement, this triggers another removal action for the new IP.


## Example usage

```hcl
resource "citrixadc_gslbldnsentry_delete" "tf_gslbldnsentry_delete" {
  ipaddress = "192.0.2.10"
}
```


## Argument Reference

* `ipaddress` - (Required) IP address of the LDNS server. Applying the resource removes the GSLB LDNS entry with this IP address. Changing this attribute forces a new resource to be created, which removes the entry for the new IP address.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `gslbldnsentry_delete`. It does not correspond to any object on the Citrix ADC.
</content>
</invoke>
