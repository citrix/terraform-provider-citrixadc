---
subcategory: "Cloud"
---

# Resource: cloudroutes

The cloudroutes resource is used to create cloudroutes.


## Example usage

```hcl
resource "citrixadc_cloudroutes" "tf_cloudroutes" {
  name             = "my_cloudroute"
  routesvpcnetwork = "client_vpc"
  vipsubnet        = "192.168.10.0/24"
  vipvpcnetwork    = "vip_vpc"
  clientipaddress  = "192.168.10.5"
}
```


## Argument Reference

* `name` - (Required) Name for the route.
* `routesvpcnetwork` - (Optional) client vpc network name.
* `vipsubnet` - (Optional) vip subnet in CIDR format.
* `vipvpcnetwork` - (Optional) vip vpc network name.
* `clientipaddress` - (Optional) IPv4 or IPv6 address attached to the nic interface towards vpc mentiond in vpcnetwork.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudroutes. It has the same value as the `name` attribute.


## Import

A `cloudroutes` can be imported using its name, e.g.

```shell
terraform import citrixadc_cloudroutes.tf_cloudroutes my_cloudroute
```
