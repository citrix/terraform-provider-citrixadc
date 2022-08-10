---
subcategory: "Network"
---

# Resource: netbridge_nsip_binding

The netbridge_nsip_binding resource is used to create netbridge_nsip_binding.


## Example usage

```hcl
resource "citrixadc_netbridge_nsip_binding" "tf_netbridge_nsip_binding" {
  name      = "my_netbridge"
  netmask   = "255.255.255.192"
  ipaddress = "10.222.74.128"
}

```


## Argument Reference

* `ipaddress` - (Required) The subnet that is extended by this network bridge. Notice, this value must have the network ip address of the subnet Minimum length =  1
* `netmask` - (Required) The network mask for the subnet.
* `name` - (Required) The name of the network bridge.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netbridge_nsip_binding. It has the same value as the `name` and `ipaddress` attributes separated by a comma.


## Import

A netbridge_nsip_bindingcan be imported using its name, e.g.

```shell
terraform import citrixadc_netbridge_nsip_binding.tf_netbridge_nsip_binding my_netbridge,10.222.74.128
```
