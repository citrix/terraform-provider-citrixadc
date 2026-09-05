---
subcategory: "VPN"
---

# Data Source: vpnvserver_vpnintranetapplication_binding

The vpnvserver_vpnintranetapplication_binding data source allows you to retrieve information about a vpnintranetapplication binding to a vpnvserver.


## Example usage

```terraform
data "citrixadc_vpnvserver_vpnintranetapplication_binding" "tf_bind" {
  name                = "tf_examplevserver"
  intranetapplication = "tf_vpnintranetapplication"
}

output "name" {
  value = data.citrixadc_vpnvserver_vpnintranetapplication_binding.tf_bind.name
}

output "intranetapplication" {
  value = data.citrixadc_vpnvserver_vpnintranetapplication_binding.tf_bind.intranetapplication
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `intranetapplication` - (Required) The intranet VPN application.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnintranetapplication_binding. It is the concatenation of `name` and `intranetapplication` attributes separated by a comma.

### Read-only vpnvserver_vpnintranetapplication_binding metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_vpnvserver_vpnintranetapplication_binding` resource). They are Computed/GET-only and are `null` when the appliance does not return them.

* `acttype` - Action type of the binding.
