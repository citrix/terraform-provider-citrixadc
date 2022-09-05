---
subcategory: "Audit"
---

# Resource: Auditnslogglobal_auditnslogpolicy_binding

The Auditnslogglobal_auditnslogpolicy_binding resource is used to create Auditnslogglobal_auditnslogpolicy_binding.


## Example usage

```hcl
resource "citrixadc_auditnslogglobal_auditnslogpolicy_binding" "tf_auditnslogglobal_auditnslogpolicy_binding" {
  policyname = "SETASLEARNNSLOG_ADV_POL"
  priority   = 100
  globalbindtype = "SYSTEM_GLOBAL"
}
```


## Argument Reference

* `policyname` - (Required) Name of the audit nslog policy.
* `priority` - (Required) Specifies the priority of the policy. Minimum value =  1 Maximum value =  2147483647
* `globalbindtype` - (Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]
* `builtin` - (Optional) Indicates that a variable is a built-in (SYSTEM INTERNAL) type. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the Auditnslogglobal_auditnslogpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A Auditnslogglobal_auditnslogpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_auditnslogglobal_auditnslogpolicy_binding.tf_auditnslogglobal_auditnslogpolicy_binding SETASLEARNNSLOG_ADV_POL
```
