---
subcategory: "DNS"
---

# Resource: dnsglobal_dnspolicy_binding

The dnsglobal_dnspolicy_binding resource is used to create DNS global transform policy binding.


## Example usage

```hcl
resource "citrixadc_dnspolicy" "dnspolicy" {
  name = "policy_A"
  rule = "CLIENT.IP.SRC.IN_SUBNET(1.1.1.1/24)"
  drop = "YES"
}
resource "citrixadc_dnsglobal_dnspolicy_binding" "dnsglobal_dnspolicy_binding" {
  policyname = citrixadc_dnspolicy.dnspolicy.name
  priority   = 30
  type       = "REQ_DEFAULT"
}
```


## Argument Reference

* `policyname` - (Required) Name of the dns policy.
* `priority` - (Required) Specifies the priority of the policy with which it is bound. Maximum allowed priority should be less than 65535.
* `globalbindtype` - (Optional) Global bind type. Defaults to `SYSTEM_GLOBAL`.
* `gotopriorityexpression` - (Optional) Expression or other value specifying the next policy to be evaluated if the current policy evaluates to TRUE. Specify one of the following values: NEXT - Evaluate the policy with the next higher priority number. END - End policy evaluation. USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT. An expression that evaluates to a number.
* `invoke` - (Optional) Invoke flag.
* `labelname` - (Optional) Name of the label to invoke if the current policy rule evaluates to TRUE.
* `labeltype` - (Optional) Type of policy label invocation.
* `type` - (Optional) Type of global bind point for which to show bound policies.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnsglobal_dnspolicy_binding. It is the concatenation of the `policyname` and `type` attributes separated by a comma.


## Import

A dnsglobal_dnspolicy_binding can be imported using the concatenation of its `policyname` and `type` attributes separated by a comma, e.g.

```shell
terraform import citrixadc_dnsglobal_dnspolicy_binding.dnsglobal_dnspolicy_binding policy_A,REQ_DEFAULT
```
