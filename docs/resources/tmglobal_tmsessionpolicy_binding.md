---
subcategory: "Traffic Management"
---

# Resource: tmglobal_tmsessionpolicy_binding

Binds a Traffic Management (TM) session policy to the global TM bind point so the policy is evaluated for all traffic, letting you enforce session settings (such as session timeout, single sign-on, and authorization) globally rather than per virtual server.

This binding is immutable. Changing any attribute forces the binding to be recreated (unbind and re-bind), and there is no in-place update.


## Example usage

```hcl
resource "citrixadc_tmglobal_tmsessionpolicy_binding" "tf_tmglobal_tmsessionpolicy_binding" {
  policyname             = "tf_tmsessionpolicy"
  priority               = 100
  gotopriorityexpression = "NEXT"
}
```


## Argument Reference

* `policyname` - (Required) The name of the TM session policy to bind to TM global. Changing this forces a new resource to be created.
* `priority` - (Optional) The priority of the policy that determines the order in which policies are evaluated at the global bind point. Changing this forces a new resource to be created.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tmglobal_tmsessionpolicy_binding. It has the same value as the `policyname` attribute.
* `feature` - The feature to be checked while applying this config. This is a read-only value returned by the ADC.


## Import

A tmglobal_tmsessionpolicy_binding can be imported using its policyname, e.g.

```shell
terraform import citrixadc_tmglobal_tmsessionpolicy_binding.tf_tmglobal_tmsessionpolicy_binding tf_tmsessionpolicy
```
