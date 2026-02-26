---
subcategory: "VPN"
---

# Data Source: vpnvserver_vpnnexthopserver_binding

The vpnvserver_vpnnexthopserver_binding data source allows you to retrieve information about a VPN virtual server to next hop server binding.

## Example Usage

```terraform
data "citrixadc_vpnvserver_vpnnexthopserver_binding" "tf_bind" {
  name          = "tf_exampleserver"
  nexthopserver = "tf_vpnnexthopserver"
}

output "name" {
  value = data.citrixadc_vpnvserver_vpnnexthopserver_binding.tf_bind.name
}

output "nexthopserver" {
  value = data.citrixadc_vpnvserver_vpnnexthopserver_binding.tf_bind.nexthopserver
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server.
* `nexthopserver` - (Required) The name of the next hop server bound to the VPN virtual server.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnnexthopserver_binding. It is a system-generated identifier.
