---
subcategory: "VPN"
---

# Resource: vpnicaconnection_kill

The vpnicaconnection_kill resource terminates (kills) active ICA connections that are tunneled through Citrix ADC (NetScaler) Gateway. It is an action-only resource: applying it invokes the NITRO `kill` action against the matching active ICA connections. Use it to forcibly disconnect a specific user's ICA sessions, sessions of a particular transport protocol, or all active ICA connections at once.

This resource does not manage a persistent object on the appliance. Each apply performs the kill; changing any argument re-triggers the kill action.


## Example usage

### Terminate ICA connections for a specific user

```hcl
resource "citrixadc_vpnicaconnection_kill" "kill_user" {
  username   = "jdoe"
  transproto = "TCP"
}
```

### Terminate all active ICA connections

```hcl
resource "citrixadc_vpnicaconnection_kill" "kill_all" {
  all = true
}
```


## Argument Reference

At least one of the following arguments should be specified to select which ICA connections to terminate. Because the kill action is one-shot, changing any argument forces a new resource to be created (a new kill action is executed).

* `username` - (Optional) User name for which ICA connections need to be terminated. Changing this attribute forces a new resource to be created.
* `transproto` - (Optional) Transport type for the existing ICA connection to terminate. Possible values: [ TCP, UDP ]. Changing this attribute forces a new resource to be created.
* `all` - (Optional) Terminate all active ICA connections when set to `true`. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnicaconnection_kill resource. It is set to `vpnicaconnection_kill`.


Import is not supported for this action resource, as there is no persistent object on the Citrix ADC to import.
