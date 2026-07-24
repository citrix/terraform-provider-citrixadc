---
subcategory: "Network"
---

# Resource: rnatglobal_auditsyslogpolicy_binding

This resource is used to bind an audit syslog policy to the global RNAT configuration.


## Example usage

```hcl
resource "citrixadc_auditsyslogaction" "syslogaction1" {
  name       = "syslogaction1"
  serverip   = "10.78.60.33"
  serverport = 514
  loglevel = [
    "ERROR",
    "NOTICE",
  ]
}

resource "citrixadc_auditsyslogpolicy" "syslogpol1" {
  name   = "syslogpol1"
  rule   = "ns_true"
  action = citrixadc_auditsyslogaction.syslogaction1.name
}

resource "citrixadc_rnatglobal_auditsyslogpolicy_binding" "rnat_syslog" {
  policy   = citrixadc_auditsyslogpolicy.syslogpol1.name
  priority = 100
}
```


## Argument Reference

* `policy` - (Required) The name of the audit syslog policy to bind to the RNAT global configuration. This is the binding key. Maximum length = 31. Changing this value forces a new binding to be created.
* `priority` - (Optional, Computed) The priority assigned to the policy at the RNAT global scope. Policies are evaluated in ascending order of priority. Minimum value = 0. Maximum value = 4294967295. Changing this value forces a new binding to be created.
* `all` - (Optional, Computed) A delete-only flag. When set to `true`, deleting this resource removes all RNAT global bindings rather than just this policy binding. It is not sent as part of the bind (create) payload and only affects the unbind (delete) operation. Changing this value forces a new binding to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnatglobal_auditsyslogpolicy_binding. It has the same value as the `policy` attribute (the RNAT global object is a singleton, so no parent name is part of the id).


## Import

A rnatglobal_auditsyslogpolicy_binding can be imported using its `policy`, e.g.

```shell
terraform import citrixadc_rnatglobal_auditsyslogpolicy_binding.rnat_syslog syslogpol1
```
