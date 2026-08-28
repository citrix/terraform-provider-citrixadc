---
subcategory: "Cloud"
---

# Data Source: cloudroutes

The `citrixadc_cloudroutes` data source is used to retrieve information about a specific cloud route configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_cloudroutes" "example" {
  name = "my_cloudroute"
}
```

## Argument Reference

* `name` - (Required) Name for the route.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the cloudroutes.
* `routesvpcnetwork` - client vpc network name.
* `vipsubnet` - vip subnet in CIDR format.
* `vipvpcnetwork` - vip vpc network name.
* `clientipaddress` - IPv4 or IPv6 address attached to the nic interface towards vpc mentiond in vpcnetwork.
