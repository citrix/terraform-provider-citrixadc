---
subcategory: "GSLB"
---

# Resource: gslbvserver_spilloverpolicy_binding

The gslbvserver_spilloverpolicy_binding resource is used to create gslbvserver_spilloverpolicy_binding.


## Example usage

```hcl
# Since the spilloverpolicy resource is not yet available on Terraform,
# the tf_spilloverpolicy policy must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add spillover policy tf_spilloverpolicy -rule TRUE -action SPILLOVER

resource "citrixadc_gslbvserver_spilloverpolicy_binding" "tf_gslbvserver_spilloverpolicy_binding" {
  name       = citrixadc_gslbvserver.tf_gslbvserver.name
  policyname = "tf_spilloverpolicy"
  priority   = 100

}

resource "citrixadc_gslbvserver" "tf_gslbvserver" {
  dnsrecordtype = "A"
  name          = "gslb_vserver"
  servicetype   = "HTTP"
  domain {
    domainname = "www.fooco.co"
    ttl        = "60"
  }
  domain {
    domainname = "www.barco.com"
    ttl        = "65"
  }
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server on which to perform the binding operation.
* `policyname` - (Required) Name of the policy bound to the GSLB vserver.
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE. 	o	If gotoPriorityExpression is not present or if it is equal to END then the policy bank evaluation ends here 	o	Else if the gotoPriorityExpression is equal to NEXT then the next policy in the priority order is evaluated. 	o	Else gotoPriorityExpression is evaluated. The result of gotoPriorityExpression (which has to be a number) is processed as follows: 		-	An UNDEF event is triggered if 			.	gotoPriorityExpression cannot be evaluated 			.	gotoPriorityExpression evaluates to number which is smaller than the maximum priority in the policy bank but is not same as any policy's priority 			.	gotoPriorityExpression evaluates to a priority that is smaller than the current policy's priority 		-	If the gotoPriorityExpression evaluates to the priority of the current policy then the next policy in the priority order is evaluated. 		-	If the gotoPriorityExpression evaluates to the priority of a policy further ahead in the list then that policy will be evaluated next. 		This field is applicable only to rewrite and responder policies.
* `priority` - (Optional) Priority.
* `type` - (Optional) The bindpoint to which the policy is bound
* `order` - (Optional) Order number to be assigned to the service when it is bound to the lb vserver


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbvserver_spilloverpolicy_binding is the concatenation of the `name` and `policyname` attributes separated by a comma.


## Import

A gslbvserver_spilloverpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_gslbvserver_spilloverpolicy_binding.tf_gslbvserver_spilloverpolicy_binding gslb_vserver,tf_spilloverpolicy
```
