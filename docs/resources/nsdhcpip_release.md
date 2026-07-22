---
subcategory: "NS"
---

# Resource: nsdhcpip_release

The nsdhcpip_release resource releases the DHCP lease for the appliance management IP. It is an action-only, zero-attribute resource: applying it triggers the release, and there are no configurable arguments.

~> **NOTE** This is an action resource: applying it performs the release; it does not manage a persistent object, so re-applying re-runs the release.


## Example usage

```hcl
resource "citrixadc_nsdhcpip_release" "tf_nsdhcpip_release" {
}
```


## Argument Reference

This resource has no configurable arguments.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsdhcpip_release resource. It is set to `nsdhcpip_release`.
