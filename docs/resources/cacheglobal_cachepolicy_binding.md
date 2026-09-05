---
subcategory: "Integrated Caching"
---

# Resource: cacheglobal_cachepolicy_binding

The cacheglobal_cachepolicy_binding resource is used to create cacheglobal_cachepolicy_binding.


## Example usage

```hcl
resource "citrixadc_cacheglobal_cachepolicy_binding" "tf_cacheglobal_cachepolicy_binding" {
  policy   = "my_cachepolicy"
  priority = 100
  type     = "REQ_DEFAULT"
}
```


## Argument Reference

* `policy` - (Required) Name of the cache policy.
* `priority` - (Required) Specifies the priority of the policy.
* `type` - (Required) The bind point to which policy is bound. When you specify the type, detailed information about that bind point appears.
* `globalbindtype` - (Optional) Global bind type.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority. Applicable only to default-syntax policies.
* `labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE. (To invoke a label associated with a virtual server, specify the name of the virtual server.)
* `labeltype` - (Optional) Type of policy label to invoke.
* `precededefrules` - (Optional) Specify whether this policy should be evaluated.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cacheglobal_cachepolicy_binding. It is the concatenation of the `policy` and `type` attributes separated by a comma.


## Import

A cacheglobal_cachepolicy_binding can be imported using the concatenation of the `policy` and `type` attributes separated by a comma, e.g.

```shell
terraform import citrixadc_cacheglobal_cachepolicy_binding.tf_cacheglobal_cachepolicy_binding my_cachepolicy,REQ_DEFAULT
```
