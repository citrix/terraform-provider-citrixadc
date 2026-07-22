---
subcategory: "System"
---

# Resource: systemsession_kill

The systemsession_kill resource terminates administrative (NITRO/CLI/GUI) sessions on the Citrix ADC. Applying it invokes the NITRO `kill` action to forcibly log out either a single session identified by its session ID, or every administrative session on the appliance. Use it to revoke access for a stale or unauthorized session, or to clear out administrative sessions during maintenance.

This is a destructive, one-shot action resource: applying it performs the kill; it does not manage a persistent object. Each apply performs the kill; changing `sid` or `all` re-triggers the action (both attributes force replacement).

~> **WARNING:** Setting `all = true` terminates **ALL** administrative sessions except the current one — and depending on session reuse, this can include the provider's own NITRO session, causing subsequent operations in the same apply to fail with authentication errors. Specifying a `sid` kills only that single session. Use this resource deliberately; killing sessions immediately disconnects active administrators.

-> **NOTE:** Exactly one of `sid` or `all` must be specified. They are mutually exclusive — supplying both, or neither, results in a configuration validation error.


## Example usage

### Kill a single administrative session by its session ID

```hcl
resource "citrixadc_systemsession_kill" "kill_one" {
  sid = 12
}
```

### Kill all administrative sessions except the current one

```hcl
resource "citrixadc_systemsession_kill" "kill_all" {
  all = true
}
```


## Argument Reference

The following arguments are supported. Exactly one of `sid` or `all` must be set.

* `sid` - (Optional) ID of the system session to kill. Mutually exclusive with `all`. Changing this attribute re-triggers the kill action.
* `all` - (Optional) Terminate all the system sessions except the current session. Mutually exclusive with `sid`. See the warning above before using. Changing this attribute re-triggers the kill action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemsession_kill resource. It is set to `systemsession_kill`.
