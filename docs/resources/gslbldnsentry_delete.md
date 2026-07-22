---
subcategory: "GSLB"
---

# Resource: gslbldnsentry_delete

The gslbldnsentry_delete resource removes a runtime-learned LDNS (local DNS) entry from the Citrix ADC GSLB subsystem. When GSLB serves DNS requests, the ADC dynamically learns the LDNS servers that resolve on behalf of clients and uses these entries for proximity-based site selection. Use this resource to purge a stale or unwanted learned entry for a specific LDNS IP address (equivalent to the ADC CLI command `rm gslb ldnsentry <ipaddress>`).

~> **This is an action resource.** Applying it removes the learned LDNS entry whose IP matches `ipaddress`; it does not manage a persistent object. Destroying the Terraform resource does **not** re-add the entry. To remove a different LDNS entry, change `ipaddress`; because `ipaddress` forces replacement, this triggers another removal action for the new IP.


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

* `id` - The id of the gslbldnsentry_delete resource. It is set to `gslbldnsentry_delete`.
