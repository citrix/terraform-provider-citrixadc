---
subcategory: "System"
---

# Resource: systemsession

Terminates one or more administrative (NITRO/CLI/GUI) sessions on the Citrix ADC. Applying this resource invokes the NITRO `kill` action to forcibly log out either a single session identified by its session ID, or every administrative session on the appliance.

This is a **destructive, one-shot action resource**. NITRO exposes no add/update/delete API for it, so:

* `Create` performs the `kill` action.
* `Read` and `Update` are no-ops; the resource is not reconciled against live state (administrative sessions are transient runtime objects).
* `Delete` removes the resource from Terraform state only — there is no inverse "un-kill" operation.

To re-run the kill action, taint the resource or change `sid`/`all` (both attributes force replacement).

~> **WARNING:** Setting `all = true` terminates **ALL** administrative sessions except the current one — and depending on session reuse, this can include the provider's own NITRO session, causing subsequent operations in the same apply to fail with authentication errors. Specifying a `sid` kills only that single session. Use this resource deliberately; killing sessions immediately disconnects active administrators.

-> **NOTE:** Exactly one of `sid` or `all` must be specified. They are mutually exclusive — supplying both, or neither, results in a configuration validation error.


## Example usage

Kill a single administrative session by its session ID:

```hcl
resource "citrixadc_systemsession" "kill_one" {
  sid = 12
}
```

Kill all administrative sessions except the current one:

```hcl
resource "citrixadc_systemsession" "kill_all" {
  all = true
}
```


## Argument Reference

The following arguments are supported. Exactly one of `sid` or `all` must be set.

* `sid` - (Optional) ID of the system session to kill. Mutually exclusive with `all`. Changing this value forces a new resource to be created.
* `all` - (Optional) Terminate all the system sessions except the current session. Mutually exclusive with `sid`. See the warning above before using. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The identifier of the kill action: the `sid` value when a single session was killed, or the literal string `"all"` when `all = true`.
