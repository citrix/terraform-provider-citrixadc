---
subcategory: "Content Switching"
---

# Data Source: csvserver_lbvserver_binding

The csvserver_lbvserver_binding data source allows you to retrieve information about a binding between a content switching virtual server and a load balancing virtual server.

## Example Usage

```terraform
data "citrixadc_csvserver_lbvserver_binding" "tf_csvserver_lbvserver_binding" {
  name      = "tf_csvserver"
  lbvserver = "tf_lbvserver"
}

output "name" {
  value = data.citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding.name
}

output "lbvserver" {
  value = data.citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding.lbvserver
}

output "targetvserver" {
  value = data.citrixadc_csvserver_lbvserver_binding.tf_csvserver_lbvserver_binding.targetvserver
}
```

## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `lbvserver` - (Required) Name of the default lb vserver bound. Use this param for Default binding only. For Example: bind cs vserver cs1 -lbvserver lb1.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_lbvserver_binding. It is a system-generated identifier in the format `name,lbvserver`.
* `targetvserver` - The virtual server name (created with the add lb vserver command) to which content will be switched.
