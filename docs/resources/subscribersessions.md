---
subcategory: "Subscriber"
---

# Resource: subscribersessions

The subscribersessions resource performs the Citrix ADC "clear subscriber sessions" action. Applying it flushes entries from the subscriber session database that the ADC maintains for Subscriber/Gx (PCRF) sessions. Use it to force the ADC to purge stale or specific subscriber sessions so that fresh session state is negotiated with the PCRF on the next request.

This is an **action-only resource**. Each apply fires a `?action=clear` call against the ADC; there is no persistent object to manage. Read, Update, and Delete are no-ops, and the resource holds only a synthetic ID in state. Because the `ip` and `vlan` selectors are `RequiresReplace`, changing either one causes Terraform to destroy and recreate the resource, which re-fires the clear action.

~> **Caution: a bare apply flushes the ENTIRE subscriber session database.** If you configure the resource with neither `ip` nor `vlan`, the clear action removes every subscriber session on the ADC. Always set `ip` and/or `vlan` when you intend to clear only a specific subscriber session.

-> **Note:** The Subscriber/Gx/PCRF (Telco) feature must be licensed and enabled on the Citrix ADC for subscriber sessions to exist. If the feature is not active, there are no sessions to clear and the action is a no-op.


## Example usage

### Clear a specific subscriber session (recommended)

```hcl
resource "citrixadc_subscribersessions" "clear_one" {
  ip   = "198.51.100.25"
  vlan = 100
}
```

### Clear all sessions for a subscriber IP

```hcl
resource "citrixadc_subscribersessions" "clear_ip" {
  ip = "198.51.100.25"
}
```

### Clear the entire subscriber session database (use with extreme care)

```hcl
# WARNING: omitting both ip and vlan flushes ALL subscriber sessions.
resource "citrixadc_subscribersessions" "clear_all" {
}
```


## Argument Reference

Both arguments are Optional. Providing one or both narrows the clear to a specific session; omitting both clears the entire subscriber session database.

Note: Because this is an action-only resource whose Read is a no-op, these attributes are Optional only. They are never populated by the ADC.

* `ip` - (Optional) Subscriber IP address of the session to clear. Changing this value forces the resource to be recreated (re-firing the clear action).
* `vlan` - (Optional) The VLAN number on which the subscriber is located. Changing this value forces the resource to be recreated (re-firing the clear action).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for the clear action. It is derived from the supplied selectors: `subscribersessions-clear-all` when no selector is given, `subscribersessions-clear-<ip>`, `subscribersessions-clear-<ip>-<vlan>`, or `subscribersessions-clear-vlan-<vlan>`. It does not correspond to any persistent object on the ADC.


## Import

Because subscribersessions is an action-only resource with no persistent backing object, importing is not meaningful and is not supported in a useful way. To perform another clear, add or re-apply the resource configuration.
