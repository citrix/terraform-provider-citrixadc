---
subcategory: "SSL"
---

# Data Source: sslvserver_sslpolicy_binding

The sslvserver_sslpolicy_binding data source allows you to retrieve information about a specific binding between an SSL virtual server and an SSL policy.

## Example usage

```terraform
data "citrixadc_sslvserver_sslpolicy_binding" "tf_binding" {
  vservername = "tf_lbvserver"
  policyname  = "tf_sslpolicy"
  type        = "REQUEST"
}

output "priority" {
  value = data.citrixadc_sslvserver_sslpolicy_binding.tf_binding.priority
}

output "gotopriorityexpression" {
  value = data.citrixadc_sslvserver_sslpolicy_binding.tf_binding.gotopriorityexpression
}
```

## Argument Reference

The following arguments are required:

* `vservername` - (Required) Name of the SSL virtual server.
* `policyname` - (Required) The name of the SSL policy binding.
* `type` - (Required) Bind point to which to bind the policy. REQUEST: Policy evaluation is done at the application above SSL; this bind point is the default and is used for actions based on clientauth and client cert. INTERCEPT_REQ: Policy evaluation is done during the SSL handshake to decide whether to intercept or not (allowed actions: INTERCEPT, BYPASS and RESET). CLIENTHELLO_REQ: Policy evaluation is done while handling the Client Hello Request (allowed actions: RESET, FORWARD and PICKCACERTGRP).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver_sslpolicy_binding. It is the concatenation of the `policyname`, `priority`, `type` and `vservername` attributes separated by a comma.
* `priority` - The priority of the policies bound to this SSL service.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke flag. This attribute is relevant only for ADVANCED policies.
* `labeltype` - Type of policy label invocation.
* `labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE.
