---
subcategory: "VPN"
---

# Data Source: vpnvserver_analyticsprofile_binding

The vpnvserver_analyticsprofile_binding data source allows you to retrieve information about the binding between a VPN virtual server and an analytics profile.

## Example Usage

```terraform
data "citrixadc_vpnvserver_analyticsprofile_binding" "tf_bind" {
  name             = "tf_vserver"
  analyticsprofile = "new_profile"
}

output "binding_id" {
  value = data.citrixadc_vpnvserver_analyticsprofile_binding.tf_bind.id
}

output "vpnvserver_name" {
  value = data.citrixadc_vpnvserver_analyticsprofile_binding.tf_bind.name
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server.
* `analyticsprofile` - (Required) Name of the analytics profile bound to the VPN Vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_analyticsprofile_binding. It is a system-generated identifier.
