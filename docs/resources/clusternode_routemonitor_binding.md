---
subcategory: "Cluster"
---

# Resource: clusternode_routemonitor_binding

The clusternode_routemonitor_binding resource is used to create clusternode_routemonitor_binding.


## Example usage

```hcl
resource "citrixadc_clusternode_routemonitor_binding" "tf_clusternode_routemonitor_binding" {
  nodeid       = 1
  routemonitor = "10.222.74.128"
  netmask      = "255.255.255.192"
}

```


## Argument Reference

* `routemonitor` - (Required) The IP address (IPv4 or IPv6). *NOTICE* This ip should be the same as the network address (gateway) in compliance with the netmask you provided. 
* `netmask` - (Requied) The netmask.
* `nodeid` - (Required) A number that uniquely identifies the cluster node. . Minimum value =  0 Maximum value =  31


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternode_routemonitor_binding. It is the concatenation of the `nodeid` and `routemonitor` attributes separated by a comma.


## Import

A routemonitor can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternode_routemonitor_binding.tf_clusternode_routemonitor_binding 1,10.222.74.128
```
