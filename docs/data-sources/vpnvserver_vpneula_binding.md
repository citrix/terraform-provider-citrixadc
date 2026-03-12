---
subcategory: "VPN"
---

# Data Source: vpnvserver_vpneula_binding

The vpnvserver_vpneula_binding data source allows you to retrieve information about the binding between a VPN virtual server and a VPN EULA.


## Example Usage

```terraform
data "citrixadc_vpnvserver_vpneula_binding" "tf_bind" {
  name = "tf_examplevserver"
  eula = "tf_vpneula"
}

output "name" {
  value = data.citrixadc_vpnvserver_vpneula_binding.tf_bind.name
}

output "eula" {
  value = data.citrixadc_vpnvserver_vpneula_binding.tf_bind.eula
}
```


## Argument Reference

* `name` - (Required) Name of the virtual server.
* `eula` - (Required) Name of the EULA bound to VPN vserver.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpneula_binding. It is the concatenation of `name` and `eula` attributes separated by comma.
