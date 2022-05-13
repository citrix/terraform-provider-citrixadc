---
subcategory: "Network"
---

# Resource: rnatglobal_auditsyslogpolicy_binding

The rnatglobal_auditsyslogpolicy_binding resource is used to bind auditsyslog policy to rnat global configuration.


## Example usage

```hcl
resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
  name       = "tf_syslogaction"
  serverip   = "10.78.60.33"
  serverport = 514
  loglevel = [
    "ERROR",
    "NOTICE",
  ]
}
resource "citrixadc_auditsyslogpolicy" "tf_policy" {
  name   = "tf_auditsyslogpolicy"
  rule   = "ns_true"
  action = citrixadc_auditsyslogaction.tf_syslogaction.name
}
resource "citrixadc_rnatglobal_auditsyslogpolicy_binding" "tf_binding" {
  policy   = citrixadc_auditsyslogpolicy.tf_policy.name
  priority = 50
}
```


## Argument Reference

* `policy` - (Required) The policy Name.
* `all` - (Optional) Remove all RNAT global config
* `priority` - (Optional) The priority of the policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnatglobal_auditsyslogpolicy_binding. It has the same value as the `policy` attribute.


## Import

A rnatglobal_auditsyslogpolicy_binding can be imported using its policy, e.g.

```shell
terraform import citrixadc_rnatglobal_auditsyslogpolicy_binding.tf_binding tf_auditsyslogpolicy
```
