---
subcategory: "Front-end-optimization"
---

# Resource: feoglobal_feopolicy_binding

The feoglobal_feopolicy_binding resource is used to create feoglobal_feopolicy_binding.


## Example usage

```hcl
resource "citrixadc_feoglobal_feopolicy_binding" "tf_feoglobal_feopolicy_binding" {
  policyname = "my_feopolicy"
  type       = "REQ_DEFAULT"
  priority   = 100
}

```


## Argument Reference

* `policyname` - (Required) The name of the globally bound front end optimization policy.
* `priority` - (Optional) The priority assigned to the policy binding.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `type` - (Optional) Bindpoint to which the policy is bound. Possible values: [ REQ_OVERRIDE, REQ_DEFAULT, RES_OVERRIDE, RES_DEFAULT, HTTPQUIC_REQ_OVERRIDE, HTTPQUIC_REQ_DEFAULT, NONE ]
* `globalbindtype` - (Optional) . Possible values: [ SYSTEM_GLOBAL, VPN_GLOBAL, RNAT_GLOBAL ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the feoglobal_feopolicy_binding. It has the same value as the `policyname` attribute.


## Import

A feoglobal_feopolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_feoglobal_feopolicy_binding.tf_feoglobal_feopolicy_binding my_feopolicy
```
