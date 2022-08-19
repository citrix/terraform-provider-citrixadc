---
subcategory: "AAA"
---

# Resource: aaauser_authorizationpolicy_binding

The aaauser_authorizationpolicy_binding resource is used to create aaauser_authorizationpolicy_binding.


## Example usage

```hcl
resource "citrixadc_aaauser_authorizationpolicy_binding" "tf_aaauser_authorizationpolicy_binding" {
  username = "user1"
  policy   = citrixadc_authorizationpolicy.tf_authorize.name
  priority = 100
}

resource "citrixadc_authorizationpolicy" "tf_authorize" {
  name   = "tp-authorize-1"
  rule   = "true"
  action = "ALLOW"
}
```


## Argument Reference

* `policy` - (Required) The policy Name.
* `priority` - (Required) Integer specifying the priority of the policy.  A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies max priority is 64000. . Minimum value =  0 Maximum value =  2147483647
* `type` - (Optional) Bindpoint to which the policy is bound. Possible values: [ REQUEST, UDP_REQUEST, DNS_REQUEST, ICMP_REQUEST ]
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `username` - (Required) User account to which to bind the policy. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaauser_authorizationpolicy_binding. It is the `username` and `policy` attributes separated by a comma.


## Import

A aaauser_authorizationpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_aaauser_authorizationpolicy_binding.tf_aaauser_authorizationpolicy_binding user1,tp-authorize-1
```
