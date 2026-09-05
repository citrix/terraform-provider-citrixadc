---
subcategory: "Audit"
---

# Resource: auditnslogglobal_auditnslogpolicy_binding

The auditnslogglobal_auditnslogpolicy_binding resource is used to bind an audit nslog policy to the global audit nslog configuration.


## Example usage

```hcl
resource "citrixadc_auditnslogglobal_auditnslogpolicy_binding" "tf_auditnslogglobal_auditnslogpolicy_binding" {
  policyname     = citrixadc_auditnslogpolicy.tf_auditnslogpolicy.name
  priority       = 100
  globalbindtype = "SYSTEM_GLOBAL"
}

resource "citrixadc_auditnslogpolicy" "tf_auditnslogpolicy" {
  name   = "tf_auditnslogpolicy"
  rule   = "true"
  action = citrixadc_auditnslogaction.tf_nslogaction.name
}

resource "citrixadc_auditnslogaction" "tf_nslogaction" {
  name       = "tf_nslogaction"
  serverip   = "10.78.60.33"
  serverport = 514
  loglevel = [
    "ERROR",
    "NOTICE",
  ]
}
```


## Argument Reference

* `policyname` - (Required) Name of the audit nslog policy.
* `priority` - (Required) Specifies the priority of the policy.
* `globalbindtype` - (Optional) The global bind type. Defaults to `"SYSTEM_GLOBAL"`. Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the auditnslogglobal_auditnslogpolicy_binding. It is the concatenation of the `globalbindtype` and `policyname` attributes separated by a comma.


## Import

A auditnslogglobal_auditnslogpolicy_binding can be imported using the concatenation of the `globalbindtype` and `policyname` attributes separated by a comma, e.g.

```shell
terraform import citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding SYSTEM_GLOBAL,SETASLEARNNSLOG_ADV_POL
```
