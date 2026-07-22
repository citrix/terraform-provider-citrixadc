---
subcategory: "AAA"
---

# Resource: aaasession_kill

Terminates active AAA-TM/VPN sessions on the Citrix ADC. Applying this resource performs an imperative `kill` action: it logs off the AAA sessions that match the supplied optional filters (user, group, intranet IP range, or session key), or every active session when `all` is set. Use it for deliberate, one-shot session-termination workflows such as forcing a compromised user off the appliance or clearing sessions after a policy change.

This is an action resource: applying it performs the kill; it does not manage a persistent object, so re-applying re-runs the action. Every argument forces resource replacement. To read active sessions without terminating them, use the `citrixadc_aaasession` data source instead.


## Example usage

Terminate all active AAA-TM/VPN sessions:

```hcl
resource "citrixadc_aaasession_kill" "kill_all" {
  all = true
}
```

Terminate the sessions belonging to a specific user:

```hcl
resource "citrixadc_aaasession_kill" "kill_user" {
  username = "jdoe"
}
```

Terminate sessions matching an intranet IP range:

```hcl
resource "citrixadc_aaasession_kill" "kill_iprange" {
  iip     = "10.102.1.0"
  netmask = "255.255.255.0"
}
```

Terminate a single session by its session key:

```hcl
resource "citrixadc_aaasession_kill" "kill_session" {
  sessionkey = "a1b2c3d4e5f6"
}
```


## Argument Reference

All arguments are optional filters that select which sessions to terminate. Supplying `all = true` with no other filter terminates every active session. Changing any argument forces the resource to be recreated, which re-runs the kill action.

* `all` - (Optional) Terminate all active AAA-TM/VPN sessions. Changing this attribute re-triggers the kill action.
* `username` - (Optional) Name of the AAA user whose sessions should be terminated. Changing this attribute re-triggers the kill action.
* `groupname` - (Optional) Name of the AAA group whose member sessions should be terminated. Changing this attribute re-triggers the kill action.
* `iip` - (Optional) IP address or the first address in the intranet IP range whose sessions should be terminated. Changing this attribute re-triggers the kill action.
* `netmask` - (Optional) Subnet mask for the intranet IP range specified by `iip`. Changing this attribute re-triggers the kill action.
* `sessionkey` - (Optional) Terminate the AAA session associated with the given session key. Changing this attribute re-triggers the kill action.
* `nodeid` - (Optional) Unique number that identifies the cluster node. Setting it has no effect on which sessions are terminated; it is meaningful only on the `citrixadc_aaasession` data source. Changing this attribute re-triggers the kill action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaasession_kill resource. It is set to `aaasession_kill`.

This resource is action-only and represents no importable server-side object, so there is no Import section.
