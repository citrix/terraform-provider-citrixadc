---
subcategory: "GSLB"
---

# Resource: gslbvserver_lbpolicy_binding

The gslbvserver_lbpolicy_binding resource is used to bind lbpolicy to gslb vserver.


## Example usage

```hcl
resource "citrixadc_gslbvserver" "tf_gslbvserver" {
  name        = "tf_gslbvserver"
  servicetype = "HTTP"
}
resource "citrixadc_lbpolicy" "tf_pol" {
  name   = "tf_pol"
  rule   = "true"
  action = "NOLBACTION"
}

resource "citrixadc_gslbvserver_lbpolicy_binding" "tf_bind" {
  policyname = citrixadc_lbpolicy.tf_pol.name
  name       = citrixadc_gslbvserver.tf_gslbvserver.name
  priority   = 10
}

```


## Argument Reference

* `name` - (Required) Name of the virtual server on which to perform the binding operation.
* `policyname` - (Required) Name of the policy bound to the GSLB vserver.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
    * If gotoPriorityExpression is not present or if it is equal to END then the policy bank evaluation ends here
    * Else if the gotoPriorityExpression is equal to NEXT then the next policy in the priority order is evaluated.
    * Else gotoPriorityExpression is evaluated. 
    The result of gotoPriorityExpression (which has to be a number) is processed as follows:
    - An UNDEF event is triggered if 
        * gotoPriorityExpression cannot be evaluated 
        * gotoPriorityExpression evaluates to number which is smaller than the maximum priority in the policy bank but is not same as any policy's priority
        * gotoPriorityExpression evaluates to a priority that is smaller than the current policy's priority
    -	If the gotoPriorityExpression evaluates to the priority of the current policy then the next policy in the priority order is evaluated.
    -	If the gotoPriorityExpression evaluates to the priority of a policy further ahead in the list then that policy will be evaluated next.
    This field is applicable only to rewrite and responder policies.
* `order` - (Optional) Order number to be assigned to the service when it is bound to the lb vserver.
* `priority` - (Optional) Priority.
* `type` - (Optional) The bindpoint to which the policy is bound


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver_lbpolicy_binding is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A gslbvserver_lbpolicy_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_gslbvserver_lbpolicy_binding.tf_bind tf_gslbvserver,tf_pol
```
