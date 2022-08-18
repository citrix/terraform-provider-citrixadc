---
subcategory: "Network"
---

# Resource: netbridge_nsip6_binding

The netbridge_nsip6_binding resource is used to create netbridge_nsip6_binding.


## Example usage

```hcl
resource "citrixadc_netbridge_nsip6_binding" "tf_netbridge_nsip6_binding" {
  name      = "my_netbridge"
  ipaddress = "dea:97c5:d381:e72b::/64"
}
```


## Argument Reference

* `ipaddress` - (Required) The subnet that is extended by this network bridge. Minimum length =  1
* `name` - (Required) The name of the network bridge.
* `netmask` - (Optional) The network mask for the subnet.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netbridge_nsip6_binding. It has the same value as the `name` and `ipaddress`attributes separated by a comma.


## Import

A netbridge_nsip6_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_netbridge_nsip6_binding.tf_netbridge_nsip6_binding my_netbridge,dea:97c5:d381:e72b::/64
```
