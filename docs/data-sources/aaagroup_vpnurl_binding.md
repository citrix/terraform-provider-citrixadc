---
subcategory: "AAA"
---

# Data Source: aaagroup_vpnurl_binding

The aaagroup_vpnurl_binding data source allows you to retrieve information about a specific aaagroup vpnurl binding.

## Example Usage

```terraform
data "citrixadc_aaagroup_vpnurl_binding" "tf_aaagroup_vpnurl_binding" {
  groupname = citrixadc_aaagroup_vpnurl_binding.tf_aaagroup_vpnurl_binding.groupname
  urlname   = citrixadc_aaagroup_vpnurl_binding.tf_aaagroup_vpnurl_binding.urlname
}

output "groupname" {
  value = data.citrixadc_aaagroup_vpnurl_binding.tf_aaagroup_vpnurl_binding.groupname
}

output "urlname" {
  value = data.citrixadc_aaagroup_vpnurl_binding.tf_aaagroup_vpnurl_binding.urlname
}
```

## Argument Reference

* `groupname` - (Required) Name of the group that you are binding.
* `urlname` - (Required) The intranet url.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate.
* `id` - The id of the aaagroup_vpnurl_binding. It is a system-generated identifier.
