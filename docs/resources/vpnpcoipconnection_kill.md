---
subcategory: "VPN"
---

# Resource: vpnpcoipconnection_kill

The vpnpcoipconnection_kill resource terminates (kills) active PCoIP connections that are tunneled through Citrix ADC (NetScaler) Gateway. It is an action-only resource: applying it invokes the NITRO `kill` action against the matching active PCoIP connections. Use it to forcibly disconnect a specific user's PCoIP sessions or to tear down all active PCoIP connections at once, for example when reclaiming licenses or evicting a stuck session.

This resource does not create, read, or manage a persistent object on the appliance. There is no NITRO GET endpoint for PCoIP connections, so there is no corresponding data source. Each apply performs the kill; changing any argument re-triggers the kill action.


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

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `vpnpcoipconnection_kill`. It does not correspond to any object on the Citrix ADC.


## Behavior notes

This is an action resource, so its lifecycle differs from a normal managed object:

* **Apply (create)** executes the `kill` action against the matching active PCoIP connections.
* **Read** is a no-op. The killed connection is not a persistent, GET-backed object, so the provider preserves state rather than re-fetching.
* **Update** is a no-op. There is no NITRO update endpoint, and every argument is replace-only, so any change triggers a new kill action instead of an in-place update.
* **Destroy** only removes the resource from Terraform state. The `kill` action has no inverse NITRO endpoint, so nothing is reverted on the Citrix ADC.

Import is not supported for this action resource, as there is no persistent object on the Citrix ADC to import.
