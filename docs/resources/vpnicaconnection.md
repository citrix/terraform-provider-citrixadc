---
subcategory: "VPN"
---

# Resource: vpnicaconnection

The vpnicaconnection resource terminates (kills) active ICA connections that are tunneled through Citrix ADC (NetScaler) Gateway. It is an **action resource**: applying it invokes the NITRO `kill` action against the matching active ICA connections. Use it to forcibly disconnect a specific user's ICA sessions, sessions of a particular transport protocol, or all active ICA connections at once.

This resource does not represent a persistent managed object. Applying it performs a one-shot kill action; there is no underlying configuration object that the Citrix ADC stores and returns afterward.


## Example usage

### Terminate ICA connections for a specific user

```hcl
resource "citrixadc_vpnicaconnection" "tf_vpnicaconnection" {
  username   = "jdoe"
  transproto = "TCP"
}
```

### Terminate all active ICA connections

```hcl
resource "citrixadc_vpnicaconnection" "tf_vpnicaconnection_all" {
  all = true
}
```


## Argument Reference

At least one of the following arguments should be specified to select which ICA connections to terminate. Because the kill action is one-shot, changing any argument forces a new resource to be created (a new kill action is executed).

* `username` - (Optional) User name for which ICA connections need to be terminated. Changing this attribute forces a new resource to be created.
* `transproto` - (Optional) Transport protocol of the existing ICA connection to terminate. Possible values: [ TCP, UDP ]. Changing this attribute forces a new resource to be created.
* `all` - (Optional) Terminate all active ICA connections when set to `true`. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic Terraform state identifier composed from the kill arguments (`transproto:<value>,username:<value>`). Because the killed connection is not a queryable managed object, this ID is only a Terraform state handle and not a NITRO lookup key.


## Behavior notes

This is an action resource, so its lifecycle differs from a normal managed object:

* **Apply (create)** executes the `kill` action against the matching active ICA connections.
* **Read** is a no-op. The killed connection is not a persistent, GET-backed object, so the provider preserves state rather than re-fetching.
* **Update** is a no-op. There is no NITRO update endpoint, and every argument is replace-only, so any change triggers a new kill action instead of an in-place update.
* **Destroy** only removes the resource from Terraform state. The `kill` action has no inverse NITRO endpoint, so nothing is reverted on the Citrix ADC.

Import is not supported for this action resource, as there is no persistent object on the Citrix ADC to import.
