---
subcategory: "CI"
---

# Resource: contentinspectionpolicylabel_contentinspectionpolicy_binding

The contentinspectionpolicylabel_contentinspectionpolicy_binding resource is used to create contentinspectionpolicylabel_contentinspectionpolicy_binding.


## Example usage

```hcl
resource "citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding" "tf_ci_binding" {
  labelname  = "my_ci_label"
  policyname = "my_ci_policy"
  priority   = 100
}

```


## Argument Reference

* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) Suspend evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `invoke_labelname` - (Optional) * If labelType is policylabel, name of the policy label to invoke. * If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request or response.
* `labelname` - (Optional) Name of the contentInspection policy label to which to bind the policy.
* `labeltype` - (Optional) Type of invocation. Available settings function as follows: * reqvserver - Forward the request to the specified request virtual server. * resvserver - Forward the response to the specified response virtual server. * policylabel - Invoke the specified policy label.
* `policyname` - (Optional) Name of the contentInspection policy to bind to the policy label.
* `priority` - (Optional) Specifies the priority of the policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the contentinspectionpolicylabel_contentinspectionpolicy_binding. Itis the concatenation of  `labelname` and `policyname` attributes separated by a comma.


## Import

A contentinspectionpolicylabel_contentinspectionpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_contentinspectionpolicylabel_contentinspectionpolicy_binding.tf_ci_binding my_ci_label,my_ci_policy
```
