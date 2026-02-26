---
subcategory: "VPN"
---

# Data Source: vpnvserver_vpnurl_binding

The vpnvserver_vpnurl_binding data source allows you to retrieve information about a binding between a VPN virtual server and a VPN URL.

## Example Usage

```terraform
data "citrixadc_vpnvserver_vpnurl_binding" "tf_bind" {
  name    = "tf_examplevserver"
  urlname = "Firsturl"
}

output "name" {
  value = data.citrixadc_vpnvserver_vpnurl_binding.tf_bind.name
}

output "urlname" {
  value = data.citrixadc_vpnvserver_vpnurl_binding.tf_bind.urlname
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server.
* `urlname` - (Required) The intranet URL.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnurl_binding. It is a system-generated identifier.
