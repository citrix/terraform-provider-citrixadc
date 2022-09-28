---
subcategory: "Integrated Caching"
---

# Resource: cachepolicylabel_cachepolicy_binding

The cachepolicylabel_cachepolicy_binding resource is used to create cachepolicylabel_cachepolicy_binding.


## Example usage

```hcl
resource "citrixadc_cachepolicylabel_cachepolicy_binding" "tf_policylabel_cachepolicy_binding" {
  labelname  = "my_cachepolicylabel"
  priority   = 100
  policyname = "my_cachepolicy"
}
```


## Argument Reference

* `policyname` - (Required) Name of the cache policy to bind to the policy label.
* `priority` - (Required) Specifies the priority of the policy.
* `labelname` - (Required) Name of the cache policy label to which to bind the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next-lower priority.
* `invoke_labelname` - (Optional) Name of the policy label to invoke if the current policy rule evaluates to TRUE.
* `labeltype` - (Optional) Type of policy label to invoke: an unnamed label associated with a virtual server, or user-defined policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cachepolicylabel_cachepolicy_binding. It has the same value as the `name` attribute.


## Import

A cachepolicylabel_cachepolicy_bindingcan be imported using its name, e.g.

```shell
terraform import citrixadc_cachepolicylabel_cachepolicy_binding.tf_cachepolicylabel_cachepolicy_binding my_cachepolicylabel,my_cachepolicy
```
