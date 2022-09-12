---
subcategory: "CI"
---

# Resource: contentinspectionglobal_contentinspectionpolicy_binding

The contentinspectionglobal_contentinspectionpolicy_binding resource is used to create contentinspectionglobal_contentinspectionpolicy_binding.


## Example usage

```hcl
resource "citrixadc_contentinspectionglobal_contentinspectionpolicy_binding" "tf_ci_binding" {
  policyname = "my_ci_policy"
  priority   = 100
}

```


## Argument Reference

* `policyname` - (Required) Name of the contentInspection policy.
* `priority` - (Required) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labelname` - (Optional) * If labelType is policylabel, name of the policy label to invoke. * If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request of response.
* `labeltype` - (Optional) Type of invocation. Available settings function as follows: * reqvserver - Forward the request to the specified request virtual server. * resvserver - Forward the response to the specified response virtual server. * policylabel - Invoke the specified policy label.
* `type` - (Optional) The bindpoint to which to policy is bound.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the contentinspectionglobal_contentinspectionpolicy_binding. It has the same value as the `policyname` attribute.


## Import

A contentinspectionglobal_contentinspectionpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_contentinspectionglobal_contentinspectionpolicy_binding.tf_ci_binding my_ci_policy
```
