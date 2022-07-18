---
subcategory: "AUTHORIZATION"
---

# Resource: authorizationpolicylabel_authorizationpolicy_binding

The authorizationpolicylabel_authorizationpolicy_binding resource is used to create Authorization policylabel policy binding.


## Example usage

```hcl
resource "citrixadc_authorizationpolicy" "authorize" {
  name   = "tp-authorize-1"
  rule   = "true"
  action = "DENY"
}
resource "citrixadc_authorizationpolicylabel" "authorizationpolicylabel" {
  labelname = "trans_http_url"
}
resource "citrixadc_authorizationpolicylabel_authorizationpolicy_binding" "authorizationpolicylabel_authorizationpolicy_binding" {
  policyname = citrixadc_authorizationpolicy.authorize.name
  labelname  = citrixadc_authorizationpolicylabel.authorizationpolicylabel.labelname
  priority   = 2
}
```


## Argument Reference

* `labelname` - (Required) Name of the authorization policy label to which to bind the policy.
* `policyname` - (Required) Name of the authorization policy to bind to the policy label.
* `priority` - (Required) Specifies the priority of the policy.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - (Optional) If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then either forward the request or response to the specified virtual server or evaluate the specified policy label.
* `invoke_labelname` - (Optional) Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.
* `labeltype` - (Optional) Type of invocation. Available settings function as follows: * reqvserver - Send the request to the specified request virtual server. * resvserver - Send the response to the specified response virtual server. * policylabel - Invoke the specified policy label.



## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authorizationpolicylabel_authorizationpolicy_binding. It has the same value as the `labelname,policyname` attribute.


## Import

A authorizationpolicylabel_authorizationpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_authorizationpolicylabel_authorizationpolicy_binding.authorizationpolicylabel_authorizationpolicy_binding trans_http_url,tp-authorize-1
```
