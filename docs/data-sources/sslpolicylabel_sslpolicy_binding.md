---
subcategory: "SSL"
---

# Data Source: sslpolicylabel_sslpolicy_binding

The sslpolicylabel_sslpolicy_binding data source allows you to retrieve information about a binding between an SSL policy label and an SSL policy.

## Example Usage

```terraform
data "citrixadc_sslpolicylabel_sslpolicy_binding" "demo_sslpolicylabel_sslpolicy_binding" {
  labelname  = "ssl_pol_label"
  policyname = "certinsert_pol"
}

output "gotopriorityexpression" {
  value = data.citrixadc_sslpolicylabel_sslpolicy_binding.demo_sslpolicylabel_sslpolicy_binding.gotopriorityexpression
}

output "labeltype" {
  value = data.citrixadc_sslpolicylabel_sslpolicy_binding.demo_sslpolicylabel_sslpolicy_binding.labeltype
}
```

## Argument Reference

* `labelname` - (Required) Name of the SSL policy label to which to bind policies.
* `policyname` - (Required) Name of the SSL policy to bind to the policy label.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the sslpolicylabel_sslpolicy_binding. It is a system-generated identifier.
* `priority` - Specifies the priority of the policy.
* `invoke` - Invoke policies bound to a policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority.
* `invoke_labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE.
* `labeltype` - Type of policy label invocation. Possible values: [ vserver, service, policylabel ]
