---
subcategory: "NS"
---

# Resource: nsdhcpip

The nsdhcpip resource performs the NITRO `nsdhcpip` `release` action, which releases the DHCP lease for the appliance management IP. It is an action-only, zero-attribute resource: applying it triggers the release, and there are no configurable arguments.

~> **NOTE** There is no NITRO GET endpoint for `nsdhcpip`, so the resource cannot be read back or verified; `Read`/`Update` are no-ops and `Delete` only removes the resource from Terraform state.


## Example usage

```hcl
resource "citrixadc_nsdhcpip" "tf_nsdhcpip" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsdhcpip resource. It is a synthetic value (`nsdhcpip-config`), since the NITRO `nsdhcpip` action exposes no readable object.
