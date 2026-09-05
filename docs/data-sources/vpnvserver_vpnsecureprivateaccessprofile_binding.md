---
subcategory: "VPN"
---

# Data Source: vpnvserver_vpnsecureprivateaccessprofile_binding

The vpnvserver_vpnsecureprivateaccessprofile_binding data source allows you to retrieve information about the binding between a VPN virtual server and a Secure Private Access profile.

## Example Usage

```terraform
data "citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding" "tf_bind" {
  name                       = "tf_vserver"
  secureprivateaccessprofile = "tf_spa_profile"
}

output "binding_id" {
  value = data.citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind.id
}

output "vpnvserver_name" {
  value = data.citrixadc_vpnvserver_vpnsecureprivateaccessprofile_binding.tf_bind.name
}
```

## Argument Reference

* `name` - (Required) Name of the virtual server.
* `secureprivateaccessprofile` - (Required) Name of the Secure Private Access profile bound to the vserver.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnvserver_vpnsecureprivateaccessprofile_binding. It is the concatenation of the `name` and `secureprivateaccessprofile` attributes separated by a comma.
