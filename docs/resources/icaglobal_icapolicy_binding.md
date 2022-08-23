---
subcategory: "Ica"
---

# Resource: icaglobal_icapolicy_binding

The icaglobal_icapolicy_binding resource is used to create icaglobal_icapolicy_binding.


## Example usage

```hcl
resource "citrixadc_icaglobal_icapolicy_binding" "tf_icaglobal_icapolicy_binding" {
  policyname = "my_ica_policy"
  priority   = 100
  type       = "ICA_REQ_DEFAULT"
}
```


## Argument Reference

* `policyname` - (Required) Name of the ICA policy.
* `type` - (Required) Global bind point for which to show detailed information about the policies bound to the bind point. Possible values: [ ICA_REQ_OVERRIDE, ICA_REQ_DEFAULT ]
* `priority` - (Required) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `globalbindtype` - (Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the icaglobal_icapolicy_binding. It has the same value as the `policyname` attribute.


## Import

A icaglobal_icapolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding my_ica_policy
```
