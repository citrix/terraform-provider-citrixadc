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

* `id` - The id of the vpnvserver_appcontroller_binding. It is the concatenation of `name` and `appcontroller` attributes separated by a comma.

### Read-only vpnvserver_appcontroller_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_vpnvserver_appcontroller_binding` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `acttype` - The bound entity (action) type, as returned by the appliance.
