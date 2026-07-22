---
subcategory: "SSL"
---

# Data Source: sslservice_sslpolicy_binding

The sslservice_sslpolicy_binding data source allows you to retrieve information about the binding between an SSL service and an SSL policy.

## Example usage

```terraform
data "citrixadc_sslservice_sslpolicy_binding" "example" {
  servicename = "tf_service"
  policyname  = "tf_sslpolicy"
  priority    = 100
}

output "gotopriorityexpression" {
  value = data.citrixadc_sslservice_sslpolicy_binding.example.gotopriorityexpression
}

output "invoke" {
  value = data.citrixadc_sslservice_sslpolicy_binding.example.invoke
}
```

## Argument Reference

* `servicename` - (Required) Name of the SSL service for which to set advanced configuration.
* `policyname` - (Required) The SSL policy binding.
* `priority` - (Required) The priority of the policies bound to this SSL service.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_sslpolicy_binding.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke flag. This attribute is relevant only for ADVANCED policies.
* `labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE.
* `labeltype` - Type of policy label invocation.
