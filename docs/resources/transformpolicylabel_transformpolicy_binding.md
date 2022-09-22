---
subcategory: "Transform"
---

# Resource: transformpolicylabel_transformpolicy_binding

The transformpolicylabel_transformpolicy_binding resource is used to create Transform policylabel policy binding.


## Example usage

```hcl
resource "citrixadc_transformpolicy" "tf_trans_policy" {
  name        = "tf_trans_policy"
  profilename = "pro_1"
  rule        = "http.REQ.URL.CONTAINS(\"test_url\")"
}
resource "citrixadc_transformpolicylabel" "transformpolicylabel" {
  labelname       = "label_1"
  policylabeltype = "httpquic_req"
}
resource "citrixadc_transformpolicylabel_transformpolicy_binding" "transformpolicylabel_transformpolicy_binding" {
  policyname = citrixadc_transformpolicy.tf_trans_policy.name
  labelname  = citrixadc_transformpolicylabel.transformpolicylabel.labelname
  priority   = 2
}
```


## Argument Reference

* `labelname` - (Required) Name of the URL Transformation policy label to which to bind the policy.
* `policyname` - (Required) Name of the URL Transformation policy to bind to the policy label.
* `priority` - (Required) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `invoke_labelname` - (Optional) Name of the policy label.
* `labeltype` - (Optional) Type of invocation. Available settings function as follows: * reqvserver - Forward the request to the specified request virtual server. * policylabel - Invoke the specified policy label.



## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the transformpolicylabel_transformpolicy_binding. It has the same value as the `labelname,policyname` attribute.


## Import

A transformpolicylabel_transformpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_transformpolicylabel_transformpolicy_binding.transformpolicylabel_transformpolicy_binding label_1,tf_trans_policy
```
