---
subcategory: "VPN"
---

# Data Source: vpnvserver_intranetip6_binding

The vpnvserver_intranetip6_binding data source allows you to retrieve information about vpnvserver_intranetip6_binding.

## Example Usage

```terraform
data "citrixadc_vpnvserver_intranetip6_binding" "tf_binding" {
  name        = "tf_vserverexample"
  intranetip6 = "2.3.4.5"
}

output "id" {
  value = data.citrixadc_vpnvserver_intranetip6_binding.tf_binding.id
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server.
* `intranetip6` - (Required) The network id for the range of intranet IP6 addresses or individual intranet ip to be bound to the vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_intranetip6_binding. It is the concatenation of the `name` and `intranetip6` attributes separated by a comma.
* `numaddr` - The number of ipv6 addresses
