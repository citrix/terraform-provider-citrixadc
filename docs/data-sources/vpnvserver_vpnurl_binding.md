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

* `id` - The id of the vpnvserver_vpnurl_binding. It is the concatenation of the `name` and `urlname` attributes separated by a comma.

### Read-only vpnvserver_vpnurl_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_vpnvserver_vpnurl_binding` resource) and are therefore GET-only / Computed. Any attribute the appliance does not return is `null`.

* `acttype` - Type of action associated with the bound URL.
