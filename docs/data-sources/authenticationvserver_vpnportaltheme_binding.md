---
subcategory: "Authentication"
---

# Data Source: authenticationvserver_vpnportaltheme_binding

The authenticationvserver_vpnportaltheme_binding data source allows you to retrieve information about the binding between an authentication virtual server and a VPN portal theme.

## Example Usage

```terraform
data "citrixadc_authenticationvserver_vpnportaltheme_binding" "tf_bind" {
  name        = "tf_authenticationvserver"
  portaltheme = "tf_vpnportaltheme"
}

output "name" {
  value = data.citrixadc_authenticationvserver_vpnportaltheme_binding.tf_bind.name
}

output "portaltheme" {
  value = data.citrixadc_authenticationvserver_vpnportaltheme_binding.tf_bind.portaltheme
}
```

## Argument Reference

* `name` - (Required) Name of the authentication virtual server to which to bind the policy.
* `portaltheme` - (Required) Theme for Authentication virtual server Login portal.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationvserver_vpnportaltheme_binding. It is a system-generated identifier.
