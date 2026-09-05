---
subcategory: "Load Balancing"
---

# Data Source: lbroute

This data source is used to retrieve information about an existing Load Balancing route.

## Example Usage

```hcl
data "citrixadc_lbroute" "example" {
  network = "55.0.0.0"
  netmask = "255.0.0.0"
  td      = 0
}
```

## Argument Reference

* `network` - (Required) The IP address of the network to which the route belongs.
* `netmask` - (Required) The netmask to which the route belongs.
* `td` - (Required) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the LB route.
* `gatewayname` - The name of the route (gateway name).

### Read-only lbroute metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_lbroute` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `flags` - State of the configured gateway (for example `UP`, `DOWN`, `OUT OF SERVICE`, `DISABLED`).
