---
subcategory: "Cloud"
---

# Data Source: cloudtrafficroutes

The `citrixadc_cloudtrafficroutes` data source is used to retrieve information about a specific cloud traffic route configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_cloudtrafficroutes" "example" {
  name = "my_cloudtrafficroutes"
}
```

## Argument Reference

* `name` - (Required) Name for the traffic cloud route.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the cloudtrafficroutes.
* `targetvpcnetwork` - Target VPC network name.
* `destrange` - Destination IP range in CIDR format.
* `nexthopip` - Next hop IP address.
* `ownernode` - cluster owner node id for the nexthopipaddress.
