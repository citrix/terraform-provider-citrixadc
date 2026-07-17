---
subcategory: "Endpoint"
---

# Resource: endpointinfo_clear

The endpointinfo_clear resource clears the cached endpoint information that the Citrix ADC maintains for discovered endpoints. Use it when you want the appliance to discard the currently learned endpoint records (for example, IP endpoints) so that endpoint information is re-learned from scratch, which is useful after topology changes or when stale endpoint entries need to be flushed.

~> **One-shot action.** This resource maps to the NITRO `clear` action (`POST ?action=clear`, CLI: `clear endpointInfo -endpointKind <kind>`); it does not create a persistent object on the appliance. Each `terraform apply` that creates or replaces this resource performs the clear once. There is no readable server-side object and no NITRO GET endpoint, so there is no corresponding data source: Read is a no-op, Delete only removes the resource from Terraform state, and changing `endpointkind` forces a new clear (replacement).


## Example usage

```hcl
resource "citrixadc_endpointinfo_clear" "tf_endpointinfo_clear" {
  endpointkind = "IP"
}
```


## Argument Reference

* `endpointkind` - (Optional) Endpoint kind whose information to clear. Currently, IP endpoints are supported. Possible values: [ IP ]. Changing this value forces the resource to be recreated (re-running the clear action against the new endpoint kind).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the endpointinfo_clear resource. It is a synthetic identifier with the format `endpointinfo_clear-<endpointkind>` (for example, `endpointinfo_clear-IP`); it does not correspond to any object on the Citrix ADC.
