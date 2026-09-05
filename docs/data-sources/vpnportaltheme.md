---
subcategory: "VPN"
---

# Data Source: vpnportaltheme

The vpnportaltheme data source allows you to retrieve information about a VPN portal theme.

## Example usage

```terraform
data "citrixadc_vpnportaltheme" "tf_vpnportaltheme" {
  name = "tf_vpnportaltheme"
}

output "basetheme" {
  value = data.citrixadc_vpnportaltheme.tf_vpnportaltheme.basetheme
}
```

## Argument Reference

* `name` - (Required) Name of the uitheme

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnportaltheme. It is the same as the `name` attribute.
* `basetheme` - Base theme for the portal

### Read-only vpnportaltheme metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_vpnportaltheme` resource). They are GET-only/Computed. Any attribute the appliance does not return is `null`.

* `feature` - The feature to be checked while applying this config.
