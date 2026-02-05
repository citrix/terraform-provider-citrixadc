---
subcategory: "Load Balancing"
---

# Data Source: citrixadc_lbroute6

This data source is used to retrieve information about an existing Load Balancing IPv6 route.

## Example Usage

```hcl
data "citrixadc_lbroute6" "example" {
  network = "66::/64"
  td      = 0
}
```

## Argument Reference

* `network` - (Required) The IPv6 destination network.
* `td` - (Required) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the LB route6.
* `gatewayname` - The name of the route (gateway name).
