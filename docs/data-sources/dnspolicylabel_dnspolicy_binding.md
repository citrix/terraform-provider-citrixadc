---
subcategory: "DNS"
---

# Data Source: dnspolicylabel_dnspolicy_binding

The dnspolicylabel_dnspolicy_binding data source allows you to retrieve information about a DNS policy label to DNS policy binding.


## Example Usage

```terraform
data "citrixadc_dnspolicylabel_dnspolicy_binding" "tf_dnspolicylabel_dnspolicy_binding" {
  labelname  = "blue_label"
  policyname = "policy_A"
}

output "labelname" {
  value = data.citrixadc_dnspolicylabel_dnspolicy_binding.tf_dnspolicylabel_dnspolicy_binding.labelname
}

output "policyname" {
  value = data.citrixadc_dnspolicylabel_dnspolicy_binding.tf_dnspolicylabel_dnspolicy_binding.policyname
}

output "priority" {
  value = data.citrixadc_dnspolicylabel_dnspolicy_binding.tf_dnspolicylabel_dnspolicy_binding.priority
}
```


## Argument Reference

* `labelname` - (Required) Name of the DNS policy label.
* `policyname` - (Required) The DNS policy name.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dnspolicylabel_dnspolicy_binding. It is a combination of the `labelname` and `policyname` attributes.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke flag.
* `priority` - Specifies the priority of the policy.
* `invoke_labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE.
* `labeltype` - Type of policy label invocation.
