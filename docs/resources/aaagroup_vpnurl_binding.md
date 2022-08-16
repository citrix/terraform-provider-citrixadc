---
subcategory: "AAA"
---

# Resource: aaagroup_vpnurl_binding

The aaagroup_vpnurl_binding resource is used to create aaagroup_vpnurl_binding.


## Example usage

```hcl
resource "citrixadc_aaagroup_vpnurl_binding" "tf_aaagroup_vpnurl_binding" {
  groupname = "my_group"
  urlname   = citrixadc_vpnurl.tf_url.actualurl
}

resource "citrixadc_vpnurl" "tf_url" {
  actualurl        = "www.citrix.com"
  appjson          = "xyz"
  applicationtype  = "CVPN"
  clientlessaccess = "OFF"
  comment          = "Testing"
  linkname         = "Description"
  ssotype          = "unifiedgateway"
  urlname          = "Firsturl"
  vservername      = "server1"
}
```


## Argument Reference

* `urlname` - (Required) The intranet url.
* `groupname` - (Required) Name of the group that you are binding. Minimum length =  1
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values: * NEXT - Evaluate the policy with the next higher priority number. * END - End policy evaluation. * USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. * An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: *  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. * If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. * If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends. An UNDEF event is triggered if: * The expression is invalid. * The expression evaluates to a priority number that is numerically lower than the current policy's priority. * The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaagroup_vpnurl_binding. It is the concatenation of `groupname` and `urlname` attributes separated by a comma.


## Import

A aaagroup_vpnurl_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_aaagroup_vpnurl_binding.tf_aaagroup_vpnurl_binding my_group,Firsturl
```
