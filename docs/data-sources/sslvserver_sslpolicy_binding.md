---
subcategory: "SSL"
---

# Data Source: sslvserver_sslpolicy_binding

The sslvserver_sslpolicy_binding data source allows you to retrieve information about a specific binding between an SSL virtual server and an SSL policy.

## Example Usage

```terraform
data "citrixadc_sslvserver_sslpolicy_binding" "tf_binding" {
  vservername = "tf_lbvserver"
  policyname  = "tf_sslpolicy"
  type        = "REQUEST"
}

output "gotopriorityexpression" {
  value = data.citrixadc_sslvserver_sslpolicy_binding.tf_binding.gotopriorityexpression
}

output "invoke" {
  value = data.citrixadc_sslvserver_sslpolicy_binding.tf_binding.invoke
}

output "labelname" {
  value = data.citrixadc_sslvserver_sslpolicy_binding.tf_binding.labelname
}
```

## Argument Reference

* `vservername` - (Required) Name of the SSL virtual server.
* `policyname` - (Required) The name of the SSL policy binding.
* `type` - (Required) Bind point to which to bind the policy. Possible Values: REQUEST, INTERCEPT_REQ and CLIENTHELLO_REQ. These bindpoints mean: 1. REQUEST: Policy evaluation will be done at appplication above SSL. This bindpoint is default and is used for actions based on clientauth and client cert. 2. INTERCEPT_REQ: Policy evaluation will be done during SSL handshake to decide whether to intercept or not. Actions allowed with this type are: INTERCEPT, BYPASS and RESET. 3. CLIENTHELLO_REQ: Policy evaluation will be done during handling of Client Hello Request. Action allowed with this type is: RESET, FORWARD and PICKCACERTGRP.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke flag. This attribute is relevant only for ADVANCED policies.
* `labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE.
* `priority` - The priority of the policies bound to this SSL service.
* `labeltype` - Type of policy label invocation.
* `id` - The id of the sslvserver_sslpolicy_binding. It is a system-generated identifier.
