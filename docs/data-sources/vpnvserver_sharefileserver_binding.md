---
subcategory: "VPN"
---

# Data Source: vpnvserver_sharefileserver_binding

The vpnvserver_sharefileserver_binding data source allows you to retrieve information about a ShareFile server binding to a VPN virtual server.

## Example Usage

```terraform
data "citrixadc_vpnvserver_sharefileserver_binding" "tf_bind" {
  name      = "tf_vpnvserver"
  sharefile = "3.3.4.3:90"
}

output "name" {
  value = data.citrixadc_vpnvserver_sharefileserver_binding.tf_bind.name
}

output "sharefile" {
  value = data.citrixadc_vpnvserver_sharefileserver_binding.tf_bind.sharefile
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server.
* `sharefile` - (Required) Configured ShareFile server in XenMobile deployment. Format IP:PORT / FQDN:PORT

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_sharefileserver_binding. It is the concatenation of `name` and `sharefile` attributes separated by a comma.
