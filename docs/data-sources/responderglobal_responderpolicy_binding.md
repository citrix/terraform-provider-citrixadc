---
subcategory: "Responder"
---

# Data Source: responderglobal_responderpolicy_binding

The responderglobal_responderpolicy_binding data source allows you to retrieve information about a specific binding between a responder global entity and a responder policy.

## Example Usage

```terraform
data "citrixadc_responderglobal_responderpolicy_binding" "tf_responderglobal_responderpolicy_binding" {
  policyname = "tf_responderpolicy"
  type       = "REQ_DEFAULT"
}

output "globalbindtype" {
  value = data.citrixadc_responderglobal_responderpolicy_binding.tf_responderglobal_responderpolicy_binding.globalbindtype
}

output "priority" {
  value = data.citrixadc_responderglobal_responderpolicy_binding.tf_responderglobal_responderpolicy_binding.priority
}
```

## Argument Reference

* `policyname` - (Required) Name of the responder policy.
* `type` - (Required) Specifies the bind point whose policies you want to display. Possible values: [ REQ_OVERRIDE, REQ_DEFAULT, OTHERTCP_REQ_OVERRIDE, OTHERTCP_REQ_DEFAULT, SIPUDP_REQ_OVERRIDE, SIPUDP_REQ_DEFAULT, RADIUS_REQ_OVERRIDE, RADIUS_REQ_DEFAULT, MSSQL_REQ_OVERRIDE, MSSQL_REQ_DEFAULT, MYSQL_REQ_OVERRIDE, MYSQL_REQ_DEFAULT, HTTPQUIC_REQ_OVERRIDE, HTTPQUIC_REQ_DEFAULT ]

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `globalbindtype` - Applicable global bind point.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `id` - The id of the responderglobal_responderpolicy_binding. It is a system-generated identifier.
* `invoke` - If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.
* `labelname` - Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel.
* `labeltype` - Type of invocation. Possible values: [ vserver, policylabel ]
