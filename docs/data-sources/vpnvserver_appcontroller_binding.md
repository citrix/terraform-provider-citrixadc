---
subcategory: "VPN"
---

# Data Source: vpnvserver_appcontroller_binding

The vpnvserver_appcontroller_binding data source allows you to retrieve information about the binding between a VPN virtual server and an App Controller server.

## Example Usage

```terraform
data "citrixadc_vpnvserver_appcontroller_binding" "tf_bind" {
  name          = "tf.citrix.example.com"
  appcontroller = "http://www.example.com"
}

output "binding_id" {
  value = data.citrixadc_vpnvserver_appcontroller_binding.tf_bind.id
}

output "vpnvserver_name" {
  value = data.citrixadc_vpnvserver_appcontroller_binding.tf_bind.name
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server.
* `appcontroller` - (Required) Configured App Controller server in XenMobile deployment.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_appcontroller_binding. It is a system-generated identifier.
