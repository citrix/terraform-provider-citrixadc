---
subcategory: "ICA"
---

# Data Source: icaglobal_icapolicy_binding

The icaglobal_icapolicy_binding data source allows you to retrieve information about an ICA policy bound to the ICA global bind point.


## Example Usage

```terraform
data "citrixadc_icaglobal_icapolicy_binding" "tf_icaglobal_icapolicy_binding" {
  policyname = "my_ica_policy"
  type       = "ICA_REQ_DEFAULT"
}

output "globalbindtype" {
  value = data.citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding.globalbindtype
}

output "priority" {
  value = data.citrixadc_icaglobal_icapolicy_binding.tf_icaglobal_icapolicy_binding.priority
}
```


## Argument Reference

* `policyname` - (Required) Name of the ICA policy.
* `type` - (Required) Global bind point for which to show detailed information about the policies bound to the bind point. Possible values: [ ICA_REQ_OVERRIDE, ICA_REQ_DEFAULT ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the icaglobal_icapolicy_binding. It is the concatenation of `policyname` and `type` attributes separated by a comma.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `globalbindtype` - The global bind point to which the policy is bound. Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]
