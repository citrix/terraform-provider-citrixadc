---
subcategory: "AAA"
---

# Resource: aaasession

Terminates active AAA-TM/VPN sessions on the Citrix ADC. This resource performs an imperative `kill` action: applying it logs off the AAA sessions that match the supplied optional filters (user, group, intranet IP range, or session key), or all active sessions when `all` is set.

This is an action-only resource. It does not manage a persistent server-side object:

* **Apply runs the kill action.** Creating the resource sends a `POST ?action=kill` request to the ADC. Any matching sessions are terminated immediately.
* **There is no read-back.** Read is a no-op. A killed session is not a queryable managed object, so the provider cannot re-resolve it. State holds only the filters you supplied plus a synthetic ID.
* **There is no in-place update.** Every argument forces resource replacement. Changing any filter and re-applying re-runs the kill action against the new filter set.
* **Delete is state-only.** The `kill` action has no inverse endpoint, so destroying the resource simply removes it from Terraform state; no sessions are restored.

Because re-applying re-runs the kill, use this resource for deliberate, one-shot session-termination workflows rather than for declarative, drift-corrected configuration. To read active sessions without terminating them, use the `citrixadc_aaasession` data source instead.


## Example usage

Terminate all active AAA-TM/VPN sessions:

```hcl
resource "citrixadc_aaasession" "kill_all" {
  all = true
}
```

Terminate the sessions belonging to a specific user:

```hcl
resource "citrixadc_aaasession" "kill_user" {
  username = "jdoe"
}
```

Terminate sessions matching an intranet IP range:

```hcl
resource "citrixadc_aaasession" "kill_iprange" {
  iip     = "10.102.1.0"
  netmask = "255.255.255.0"
}
```


## Argument Reference

All arguments are optional filters that select which sessions to terminate. Supplying none combined with `all = true` terminates every active session. Changing any argument forces the resource to be recreated, which re-runs the kill action.

* `all` - (Optional) Terminate all active AAA-TM/VPN sessions.
* `username` - (Optional) Name of the AAA user whose sessions should be terminated.
* `groupname` - (Optional) Name of the AAA group whose member sessions should be terminated.
* `iip` - (Optional) IP address or the first address in the intranet IP range whose sessions should be terminated.
* `netmask` - (Optional) Subnet mask for the intranet IP range specified by `iip`.
* `sessionkey` - (Optional) Terminate the AAA session associated with the given session key.

Note: `nodeid` (the cluster-node GET filter) is not a valid argument for the kill action and is intentionally excluded from this resource. It is available only on the `citrixadc_aaasession` data source.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the constant value `"aaasession-kill"`. The ADC does not assign an ID to the kill action; this value is purely a Terraform state handle and is not a NITRO lookup key.

This resource is action-only and represents no importable server-side object, so there is no Import section.
