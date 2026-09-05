---
subcategory: "ICA"
---

# Resource: icaglobal_icapolicy_binding

The icaglobal_icapolicy_binding resource is used to bind an ICA policy to the ICA global bind point.


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
* `globalbindtype` - (Optional) The global bind point to which the policy is bound. Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ] Defaults to `"SYSTEM_GLOBAL"`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the icaglobal_icapolicy_binding. It is the concatenation of `policyname` and `type` attributes separated by a comma.


## Import

A icaglobal_icapolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding my_ica_policy,ICA_REQ_DEFAULT
```
