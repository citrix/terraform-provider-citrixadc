---
subcategory: "System"
---

# Resource: systemsession_kill

This resource is used to terminate administrative sessions on the Citrix ADC.

!> **WARNING:** This forcibly logs out a single admin session (by `sid`) or every admin session (`all`). Each apply re-triggers the kill.


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
