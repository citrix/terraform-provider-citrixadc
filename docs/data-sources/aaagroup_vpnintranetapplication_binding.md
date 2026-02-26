---
subcategory: "AAA"
---

# Data Source: aaagroup_vpnintranetapplication_binding

The aaagroup_vpnintranetapplication_binding data source allows you to retrieve information about the binding between an AAA group and a VPN intranet application.

## Example Usage

```terraform
data "citrixadc_aaagroup_vpnintranetapplication_binding" "tf_aaagroup_vpnintranetapplication_binding" {
  groupname           = "my_group"
  intranetapplication = "tf_vpnintranetapplication"
}

output "groupname" {
  value = data.citrixadc_aaagroup_vpnintranetapplication_binding.tf_aaagroup_vpnintranetapplication_binding.groupname
}

output "intranetapplication" {
  value = data.citrixadc_aaagroup_vpnintranetapplication_binding.tf_aaagroup_vpnintranetapplication_binding.intranetapplication
}
```

## Argument Reference

* `groupname` - (Required) Name of the group that you are binding.
* `intranetapplication` - (Required) Bind the group to the specified intranet VPN application.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaagroup_vpnintranetapplication_binding.
* `gotopriorityexpression` - Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number. If you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows: If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next. If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next. If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends. An UNDEF event is triggered if: The expression is invalid. The expression evaluates to a priority number that is numerically lower than the current policy's priority. The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.
