---
subcategory: "RDP"
---

# Resource: rdpconnections_kill

The rdpconnections_kill resource terminates active RDP proxy connections on the Citrix ADC. It is an action-only resource: applying it fires the NITRO `kill` action against the appliance to disconnect one or more live RDP proxy sessions established through the NetScaler Gateway / VPN RDP proxy feature. This is useful for administratively forcing users off stale or unauthorized RDP proxy sessions.

Since every apply re-fires the kill, both arguments are marked as forcing replacement. To kill connections again after a previous apply, use `terraform taint` (or `-replace`) or change a selector.

~> **Note** The RDP Proxy feature (NetScaler Gateway / VPN RDP proxy) must be in use for any RDP connections to exist. If there are no active RDP proxy sessions, the kill action succeeds with nothing to terminate.


## Example usage

### Kill all active RDP proxy connections

A bare apply (or `all = true`) terminates every active RDP proxy connection.

```hcl
resource "citrixadc_rdpconnections_kill" "kill_all" {
  all = true
}
```

### Kill RDP proxy connections for a specific user

```hcl
resource "citrixadc_rdpconnections_kill" "kill_user" {
  username = "jdoe"
}
```


## Argument Reference

Both arguments are optional. Provide `username` to target a single user's connections, or `all = true` (or neither) to terminate all active connections.

* `username` - (Optional) User name whose active RDP proxy connections should be terminated. Changing this attribute forces the kill action to be re-fired (a new resource is created).
* `all` - (Optional) When set to `true`, terminates all active RDP proxy connections. Changing this attribute forces the kill action to be re-fired (a new resource is created).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rdpconnections_kill resource. It is set to `rdpconnections_kill`.
