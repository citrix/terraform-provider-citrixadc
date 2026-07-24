---
subcategory: "Subscriber"
---

# Resource: subscribersessions_clear

This resource is used to clear subscriber (Gx/PCRF) sessions on the Citrix ADC.

~> **Action-only:** Each apply fires a `clear` action; the `ip` and `vlan` selectors are immutable, so changing them re-triggers the action.


## Example usage

### Clear a specific subscriber session (recommended)

```hcl
resource "citrixadc_subscribersessions_clear" "clear_one" {
  ip   = "198.51.100.25"
  vlan = 100
}
```

### Clear all sessions for a subscriber IP

```hcl
resource "citrixadc_subscribersessions_clear" "clear_ip" {
  ip = "198.51.100.25"
}
```

### Clear the entire subscriber session database (use with extreme care)

```hcl
# WARNING: omitting both ip and vlan flushes ALL subscriber sessions.
resource "citrixadc_subscribersessions_clear" "clear_all" {
}
```


## Argument Reference

Both arguments are Optional. Providing one or both narrows the clear to a specific session; omitting both clears the entire subscriber session database.

* `ip` - (Optional) Subscriber IP address of the session to clear. Changing this value forces the resource to be recreated (re-firing the clear action).
* `vlan` - (Optional) The VLAN number on which the subscriber is located. Changing this value forces the resource to be recreated (re-firing the clear action).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the subscribersessions_clear resource. It is set to `subscribersessions_clear`.
