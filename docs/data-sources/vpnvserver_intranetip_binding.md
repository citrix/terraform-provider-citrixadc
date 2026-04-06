---
subcategory: "VPN"
---

# Data Source: vpnvserver_intranetip_binding

The vpnvserver_intranetip_binding data source allows you to retrieve information about the binding between a VPN virtual server and an intranet IP address.

## Example Usage

```terraform
data "citrixadc_vpnvserver_intranetip_binding" "tf_bind" {
  name       = "tf_vserverexample"
  intranetip = "2.3.4.5"
}

output "name" {
  value = data.citrixadc_vpnvserver_intranetip_binding.tf_bind.name
}

output "intranetip" {
  value = data.citrixadc_vpnvserver_intranetip_binding.tf_bind.intranetip
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server.
* `intranetip` - (Required) The network ID for the range of intranet IP addresses or individual intranet IP addresses to be bound to the virtual server.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_intranetip_binding. It is the concatenation of `name` and `intranetip` attributes separated by comma.
* `netmask` - The netmask of the intranet IP address or range.
