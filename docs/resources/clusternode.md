---
subcategory: "Cluster"
---

# Resource: clusternode

The clusternode resource is used to create clusternode.


## Example usage

```hcl
resource "citrixadc_clusternode" "tf_clusternode" {
  nodeid             = 1
  ipaddress          = "10.222.74.150"
  state              = "ACTIVE"
}
```


## Argument Reference

* `nodeid` - (Required) Unique number that identifies the cluster node.
* `ipaddress` - (Required) Citrix ADC IP (NSIP) address of the appliance to add to the cluster. Must be an IPv4 address.
* `backplane` - (Optional) Interface through which the node communicates with the other nodes in the cluster. Must be specified in the three-tuple form n/c/u, where n represents the node ID and c/u refers to the interface on the appliance.
* `clearnodegroupconfig` - (Optional) Option to remove nodegroup config
* `delay` - (Optional) Applicable for Passive node and node becomes passive after this timeout (in minutes)
* `nodegroup` - (Optional) The default node group in a Cluster system.
* `priority` - (Optional) Preference for selecting a node as the configuration coordinator. The node with the lowest priority value is selected as the configuration coordinator. When the current configuration coordinator goes down, the node with the next lowest priority is made the new configuration coordinator. When the original node comes back up, it will preempt the new configuration coordinator and take over as the configuration coordinator. Note: When priority is not configured for any of the nodes or if multiple nodes have the same priority, the cluster elects one of the nodes as the configuration coordinator.
* `state` - (Optional) Admin state of the cluster node. The available settings function as follows: ACTIVE - The node serves traffic. SPARE - The node does not serve traffic unless an ACTIVE node goes down. PASSIVE - The node does not serve traffic, unless you change its state. PASSIVE state is useful during temporary maintenance activities in which you want the node to take part in the consensus protocol but not to serve traffic.
* `tunnelmode` - (Optional) To set the tunnel mode


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusternode. It has the same value as the `nodeid` attribute.


## Import

A clusternode can be imported using its name, e.g.

```shell
terraform import citrixadc_clusternode.tf_clusternode tf_csaction
```
