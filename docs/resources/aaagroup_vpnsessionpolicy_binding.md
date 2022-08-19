---
subcategory: "AAA"
---

# Resource: aaagroup_vpnsessionpolicy_binding

The aaagroup_vpnsessionpolicy_binding resource is used to create aaagroup_vpnsessionpolicy_binding.


## Example usage

```hcl
resource "citrixadc_aaagroup_vpnsessionpolicy_binding" "tf_aaagroup_vpnsessionpolicy_binding" {
  groupname = "my_group"
  policy    = citrixadc_vpnsessionpolicy.tf_vpnsessionpolicy.name
  priority  = 100
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

* `policy` - (Required) The policy name.
* `priority` - (Required) Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000. Minimum value =  0 Maximum value =  2147483647
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `groupname` - (Required) Name of the group that you are binding. Minimum length =  1
* `type` - (Optional) Bindpoint to which the policy is bound. Possible values: [ REQUEST, UDP_REQUEST, DNS_REQUEST, ICMP_REQUEST ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaagroup_vpnsessionpolicy_binding. It is the concatenation of  `groupname` and `policy` attributes separated by a comma.


## Import

A aaagroup_vpnsessionpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_aaagroup_vpnsessionpolicy_binding.tf_aaagroup_vpnsessionpolicy_binding my_group,tf_vpnsessionpolicy
```
