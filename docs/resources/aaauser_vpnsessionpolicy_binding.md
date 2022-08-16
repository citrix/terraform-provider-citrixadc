---
subcategory: "AAA"
---

# Resource: aaauser_vpnsessionpolicy_binding

The aaauser_vpnsessionpolicy_binding resource is used to create aaauser_vpnsessionpolicy_binding.


## Example usage

```hcl
resource "citrixadc_aaauser_vpnsessionpolicy_binding" "tf_aaauser_vpnsessionpolicy_binding" {
  username = "user1"
  policy   = citrixadc_vpnsessionpolicy.tf_vpnsessionpolicy.name
  priority = 100
}

resource "citrixadc_vpnsessionaction" "tf_vpnsessionaction" {
  name                       = "newsession"
  sesstimeout                = "10"
  defaultauthorizationaction = "ALLOW"
}

resource "citrixadc_vpnsessionpolicy" "tf_vpnsessionpolicy" {
  name   = "tf_vpnsessionpolicy"
  rule   = "HTTP.REQ.HEADER(\"User-Agent\").CONTAINS(\"CitrixReceiver\").NOT"
  action = citrixadc_vpnsessionaction.tf_vpnsessionaction.name
}
```


## Argument Reference

* `policy` - (Required) The policy Name.
* `priority` - (Required) Integer specifying the priority of the policy.  A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies max priority is 64000. . Minimum value =  0 Maximum value =  2147483647
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `username` - (Required) User account to which to bind the policy. Minimum length =  1
* `type` - (Optional) Bindpoint to which the policy is bound. Possible values: [ REQUEST, UDP_REQUEST, DNS_REQUEST, ICMP_REQUEST ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaauser_vpnsessionpolicy_binding. It is the concatenation of `username` and `policy` attributes separated by a comma.


## Import

A aaauser_vpnsessionpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_aaauser_vpnsessionpolicy_binding.tf_aaauser_vpnsessionpolicy_binding user1,tf_vpnsessionpolicy
```
