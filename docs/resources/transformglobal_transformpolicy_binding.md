---
subcategory: "Transform"
---

# Resource: transformglobal_transformpolicy_binding

The transformglobal_transformpolicy_binding resource is used to bind a transform policy globally on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_transformpolicy" "tf_trans_policy" {
  name        = "tf_trans_policy"
  profilename = "tf_trans_profile"
  rule        = "http.REQ.URL.CONTAINS(\"test_url\")"
}

resource "citrixadc_transformglobal_transformpolicy_binding" "tf_transformglobal_transformpolicy_binding" {
  policyname = citrixadc_transformpolicy.tf_trans_policy.name
  priority   = 2
  type       = "REQ_DEFAULT"
}
```


## Argument Reference

* `policyname` - (Required) Name of the transform policy.
* `priority` - (Required) Specifies the priority of the policy.
* `type` - (Optional) Specifies the bind point to which to bind the policy. Available settings function as follows: * REQ_OVERRIDE - Request override. Binds the policy to the priority request queue. * REQ_DEFAULT - Binds the policy to the default request queue. * HTTPQUIC_REQ_OVERRIDE - Binds the policy to the HTTP_QUIC override request queue. * HTTPQUIC_REQ_DEFAULT - Binds the policy to the HTTP_QUIC default request queue.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forwards the request or response to the specified virtual server or evaluates the specified policy label.
* `labelname` - (Optional) Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and the label type is Policy Label.
* `labeltype` - (Optional) Type of invocation. Available settings function as follows: * reqvserver - Send the request to the specified request virtual server. * resvserver - Send the response to the specified response virtual server. * policylabel - Invoke the specified policy label.
* `globalbindtype` - (Optional) The bind point at which the policy is bound globally. Defaults to `"SYSTEM_GLOBAL"`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the transformglobal_transformpolicy_binding. It is the concatenation of the `policyname` and `type` attributes separated by a comma.


## Import

A transformglobal_transformpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_transformglobal_transformpolicy_binding.tf_transformglobal_transformpolicy_binding tf_trans_policy,REQ_DEFAULT
```
