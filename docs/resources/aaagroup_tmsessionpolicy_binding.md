---
subcategory: "AAA"
---

# Resource: aaagroup_tmsessionpolicy_binding

The aaagroup_tmsessionpolicy_binding resource is used to create aaagroup_tmsessionpolicy_binding.


## Example usage

```hcl
# Since the tmsessionpolicy resource is not yet available on Terraform,
# the tf_tmsesspolicy policy must be created by hand(manually) in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli commands:
# add tmsessionaction tf_tmsessaction  -sessTimeout 30 -defaultAuthorization ALLOW
# add tmsessionpolicy tf_tmsesspolicy true tf_tmsessaction
resource "citrixadc_aaagroup_tmsessionpolicy_binding" "tf_aaagroup_tmsessionpolicy_binding" {
  groupname = "my_group"
  policy    = "tf_tmsesspolicy"
  priority  = 50
}

```


## Argument Reference

* `groupname` - (Required) Name of the group that you are binding. Minimum length =  1
* `policy` - (Required) The policy name.
* `priority` - (Required) Integer specifying the priority of the policy. A lower number indicates a higher priority. Policies are evaluated in the order of their priority numbers. Maximum value for default syntax policies is 2147483647 and for classic policies is 64000. Minimum value =  0 Maximum value =  2147483647
* `gotopriorityexpression` - (Optional) Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `type` - (Optional) Bindpoint to which the policy is bound. Possible values: [ REQUEST, UDP_REQUEST, DNS_REQUEST, ICMP_REQUEST ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaagroup_tmsessionpolicy_binding. It has the same value as the `groupname` and `policy` attributes separated by a comma.


## Import

A aaagroup_tmsessionpolicy_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_aaagroup_tmsessionpolicy_binding.tf_aaagroup_tmsessionpolicy_binding my_group,tf_tmsesspolicy
```
