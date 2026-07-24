---
subcategory: "VPN"
---

# Resource: vpnpcoipconnection_kill

This resource is used to terminate active PCoIP connections tunneled through Citrix ADC Gateway.

~> **Action-only:** Applying this resource forcibly disconnects matching PCoIP connections; each apply re-triggers the kill.


## Example usage

### Terminate PCoIP connections for a specific user

```hcl
resource "citrixadc_vpnpcoipconnection_kill" "kill_user" {
  username = "jdoe"
}
```

### Terminate all active PCoIP connections

```hcl
resource "citrixadc_vpnpcoipconnection_kill" "kill_all" {
  all = true
}
```


## Argument Reference

At least one of the following arguments should be specified to select which PCoIP connections to terminate. Because the kill action is one-shot, changing any argument forces a new resource to be created (a new kill action is executed).

* `username` - (Optional) User name for the PCoIP connections to be terminated. Changing this attribute forces a new resource to be created.
* `all` - (Optional) Terminate all active PCoIP connections when set to `true`. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnpcoipconnection_kill resource. It is set to `vpnpcoipconnection_kill`.


Import is not supported for this action resource, as there is no persistent object on the Citrix ADC to import.
