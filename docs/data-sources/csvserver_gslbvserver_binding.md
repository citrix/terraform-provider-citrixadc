---
subcategory: "Content Switching"
---

# Data Source: csvserver_gslbvserver_binding

The csvserver_gslbvserver_binding data source allows you to retrieve information about a binding between a content switching virtual server and a GSLB virtual server.

## Example Usage

```terraform
data "citrixadc_csvserver_gslbvserver_binding" "tf_csvserver_gslbvserver_binding" {
  name    = "tf_csvserver"
  vserver = "tf_gslbvserver"
}

output "name" {
  value = data.citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding.name
}

output "vserver" {
  value = data.citrixadc_csvserver_gslbvserver_binding.tf_csvserver_gslbvserver_binding.vserver
}
```

## Argument Reference

* `name` - (Required) Name of the content switching virtual server to which the content switching policy applies.
* `vserver` - (Required) Name of the default gslb or vpn vserver bound to CS vserver of type GSLB/VPN. For Example: bind cs vserver cs1 -vserver gslb1 or bind cs vserver cs1 -vserver vpn1

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the csvserver_gslbvserver_binding. It is a system-generated identifier in the format `name,vserver`.
