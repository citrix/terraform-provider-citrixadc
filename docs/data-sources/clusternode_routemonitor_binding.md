---
subcategory: "Cluster"
---

# Data Source: clusternode_routemonitor_binding

The clusternode_routemonitor_binding data source allows you to retrieve information about a cluster node to route monitor binding.


## Example Usage

```hcl
data "citrixadc_clusternode_routemonitor_binding" "tf_clusternode_routemonitor_binding" {
    nodeid       = 0
    routemonitor = "10.222.74.128"
    netmask      = "255.255.255.192"
}
```

## Argument Reference

* `nodeid` - (Required) A number that uniquely identifies the cluster node.
* `routemonitor` - (Required) The IP address (IPv4 or IPv6).
* `netmask` - (Required) The netmask.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the binding. It is the concatenation of `nodeid`, `routemonitor`, and `netmask` attributes separated by comma.
