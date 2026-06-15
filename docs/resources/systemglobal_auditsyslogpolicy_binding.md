---
subcategory: "System"
---

# Resource: systemglobal_auditsyslogpolicy_binding

Binds an audit syslog policy to the Citrix ADC system global scope. A syslog policy bound at the system global level applies its logging rules to all traffic processed by the appliance, letting you forward audit events to one or more syslog servers without binding the policy to an individual virtual server. The binding is keyless on the global side: the policy is identified solely by its `policyname`.


## Example usage

```hcl
resource "citrixadc_auditsyslogpolicy" "syslogpol1" {
  name   = "syslogpol1"
  rule   = "true"
  action = "SYSLOG_ACT_1"
}

resource "citrixadc_systemglobal_auditsyslogpolicy_binding" "syslog_global" {
  policyname = citrixadc_auditsyslogpolicy.syslogpol1.name
  priority   = 100
}
```


## Argument Reference

* `policyname` - (Required) The name of the audit syslog policy to bind to the system global scope. This is the binding key. Changing this value forces a new binding to be created.
* `priority` - (Optional, Computed) The priority assigned to the policy at the global scope. Policies are evaluated in ascending order of priority. Changing this value forces a new binding to be created.
* `nextfactor` - (Optional, Computed) On-success invoke label. Applicable for advanced authentication policy binding. Changing this value forces a new binding to be created.
* `gotopriorityexpression` - (Optional, Computed) Applicable only to advanced authentication policy. Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT (evaluate the policy with the next higher priority number) or END (end policy evaluation). Changing this value forces a new binding to be created.

The following attributes are read-only discriminators returned by the NITRO server. They describe the binding scope and are not accepted as bind arguments, so the provider does not send them in the create payload:

* `feature` - (Read-only, Computed) The feature checked while applying this configuration.
* `globalbindtype` - (Read-only, Computed) The global bind type for this binding. The NITRO server reports this as `SYSTEM_GLOBAL` for system global bindings.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemglobal_auditsyslogpolicy_binding. It is a single token that has the same value as the `policyname` attribute.


## Import

A systemglobal_auditsyslogpolicy_binding can be imported using its `policyname`, e.g.

```shell
terraform import citrixadc_systemglobal_auditsyslogpolicy_binding.syslog_global syslogpol1
```
