---
subcategory: "VPN"
---

# Data Source: vpnvserver_vpnepaprofile_binding

The vpnvserver_vpnepaprofile_binding data source allows you to retrieve information about a specific binding between a VPN virtual server and an advanced EPA profile on the Citrix ADC.

## Example Usage

```terraform
data "citrixadc_vpnvserver_vpnepaprofile_binding" "tf_binding" {
  name       = "tf.citrix.example.com"
  epaprofile = "tf_vpnepaprofile"
}

output "name" {
  value = data.citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding.name
}

output "epaprofile" {
  value = data.citrixadc_vpnvserver_vpnepaprofile_binding.tf_binding.epaprofile
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server.
* `epaprofile` - (Required) Advanced EPA profile bound to the virtual server.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnepaprofile_binding. It is a system-generated identifier.
* `epaprofileoptional` - Whether the EPA profile is marked optional for the preauthentication EPA profile.
