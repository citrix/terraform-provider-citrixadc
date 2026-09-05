---
subcategory: "Cloud"
---

# Resource: cloudtrafficroutes

The cloudtrafficroutes resource is used to create cloud traffic routes.

## Example usage

```hcl
resource "citrixadc_cloudtrafficroutes" "tf_cloudtrafficroutes" {
  name             = "my_cloudtrafficroutes"
  targetvpcnetwork = "my_vpc_network"
  destrange        = "10.0.0.0/24"
  nexthopip        = "192.168.1.1"
  ownernode        = 0
}
```

## Argument Reference

* `name` - (Required) Name for the traffic cloud route. Changing this forces a new resource to be created.
* `targetvpcnetwork` - (Optional) Target VPC network name.
* `destrange` - (Optional) Destination IP range in CIDR format.
* `nexthopip` - (Optional) Next hop IP address.
* `ownernode` - (Optional) cluster owner node id for the nexthopipaddress. Minimum value = 0, Maximum value = 31.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The id of the cloudtrafficroutes. It has the same value as the `name` attribute.

## Import

A cloudtrafficroutes can be imported using its name, e.g.

```shell
terraform import citrixadc_cloudtrafficroutes.tf_cloudtrafficroutes my_cloudtrafficroutes
```
