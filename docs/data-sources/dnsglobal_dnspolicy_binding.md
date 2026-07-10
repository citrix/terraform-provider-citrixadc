---
subcategory: "DNS"
---

# Data Source: dnsglobal_dnspolicy_binding

The dnsglobal_dnspolicy_binding data source allows you to retrieve information about DNS global transform policy bindings.


## Example usage

```terraform
data "citrixadc_dnsglobal_dnspolicy_binding" "dnsglobal_dnspolicy_binding" {
  policyname = "policy_A"
  type       = "REQ_DEFAULT"
}

output "priority" {
  value = data.citrixadc_dnsglobal_dnspolicy_binding.dnsglobal_dnspolicy_binding.priority
}

output "globalbindtype" {
  value = data.citrixadc_dnsglobal_dnspolicy_binding.dnsglobal_dnspolicy_binding.globalbindtype
}
```


## Argument Reference

* `policyname` - (Required) Name of the dns policy.
* `type` - (Required) Type of global bind point for which to show bound policies.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsglobal_dnspolicy_binding. It is the concatenation of the `policyname` and `type` attributes separated by a comma.
* `globalbindtype` - Global bind type.
* `gotopriorityexpression` - Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number.
* `invoke` - Invoke flag.
* `labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE.
* `labeltype` - Type of policy label invocation.
* `priority` - Specifies the priority of the policy with which it is bound. Maximum allowed priority should be less than 65535.
