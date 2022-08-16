---
subcategory: "AAA"
---

# Resource: aaagroup_authorizationpolicy_binding

The aaagroup_authorizationpolicy_binding resource is used to create aaagroup_authorizationpolicy_binding.


## Example usage

```hcl
resource "citrixadc_aaagroup_authorizationpolicy_binding" "tf_aaagroup_authorizationpolicy_binding" {
  groupname = "my_group"
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

* `groupname` - (Required) Name of the group that you are binding. Minimum length =  1
* `policy` - (Required) The policy name.
* `priority` - (Required) Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000. Minimum value =  0 Maximum value =  2147483647
* `type` - (Optional) Bindpoint to which the policy is bound. Possible values: [ REQUEST, UDP_REQUEST, DNS_REQUEST, ICMP_REQUEST ]
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaagroup_authorizationpolicy_binding. It is the concatenation of `grouopname` and `policy` attributes separated by a comma.


## Import

A aaagroup_authorizationpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_aaagroup_authorizationpolicy_binding.tf_aaagroup_authorizationpolicy_binding my_group,tp-authorize-1
```
