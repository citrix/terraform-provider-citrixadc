---
subcategory: "Subscriber"
---

# Resource: subscribersessions_clear

The subscribersessions_clear resource performs the Citrix ADC "clear subscriber sessions" action. Applying it flushes entries from the subscriber session database that the ADC maintains for Subscriber/Gx (PCRF) sessions. Use it to force the ADC to purge stale or specific subscriber sessions so that fresh session state is negotiated with the PCRF on the next request.

This is an **action-only resource**. Each apply fires a `?action=clear` call against the ADC; there is no persistent object to manage, and the resource holds only a synthetic ID in state. Because the `ip` and `vlan` selectors are immutable, changing either one causes Terraform to destroy and recreate the resource, which re-fires the clear action.

~> **Caution: a bare apply flushes the ENTIRE subscriber session database.** If you configure the resource with neither `ip` nor `vlan`, the clear action removes every subscriber session on the ADC. Always set `ip` and/or `vlan` when you intend to clear only a specific subscriber session.

-> **Note:** The Subscriber/Gx/PCRF (Telco) feature must be licensed and enabled on the Citrix ADC for subscriber sessions to exist. If the feature is not active, there are no sessions to clear.


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
