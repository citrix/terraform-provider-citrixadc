---
subcategory: "Cluster"
---

# Data Source: citrixadc_clusternode

Use this data source to retrieve information about an existing Cluster Node.

The `citrixadc_clusternode` data source allows you to retrieve details of a cluster node by its node ID. This is useful for referencing existing cluster nodes in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing cluster node
data "citrixadc_clusternode" "example" {
  nodeid = 1
}

# Reference node attributes
output "node_ipaddress" {
  value = data.citrixadc_clusternode.example.ipaddress
}

output "node_state" {
  value = data.citrixadc_clusternode.example.state
}
```

## Argument Reference

The following arguments are supported:

* `nodeid` - (Required) Unique number that identifies the cluster node.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the cluster node (same as nodeid).

* `backplane` - Interface through which the node communicates with the other nodes in the cluster. Must be specified in the three-tuple form n/c/u, where n represents the node ID and c/u refers to the interface on the appliance.

* `clearnodegroupconfig` - Option to remove nodegroup config. Possible values: "YES", "NO".

* `delay` - Applicable for Passive node and node becomes passive after this timeout (in minutes).

* `force` - Node will be removed from cluster without prompting for user confirmation.

* `ipaddress` - Citrix ADC IP (NSIP) address of the appliance to add to the cluster. Must be an IPv4 address.

* `nodegroup` - The default node group in a Cluster system.

* `priority` - Preference for selecting a node as the configuration coordinator. The node with the lowest priority value is selected as the configuration coordinator. When the current configuration coordinator goes down, the node with the next lowest priority is made the new configuration coordinator. When the original node comes back up, it will preempt the new configuration coordinator and take over as the configuration coordinator. Note: When priority is not configured for any of the nodes or if multiple nodes have the same priority, the cluster elects one of the nodes as the configuration coordinator.

* `state` - Admin state of the cluster node. The available settings function as follows:
  - ACTIVE - The node serves traffic.
  - SPARE - The node does not serve traffic unless an ACTIVE node goes down.
  - PASSIVE - The node does not serve traffic, unless you change its state. PASSIVE state is useful during temporary maintenance activities in which you want the node to take part in the consensus protocol but not to serve traffic.

* `tunnelmode` - To set the tunnel mode.

## Common Use Cases

### Retrieve Cluster Node Information

```hcl
data "citrixadc_clusternode" "node1" {
  nodeid = 1
}

output "node1_details" {
  value = {
    ipaddress = data.citrixadc_clusternode.node1.ipaddress
    state     = data.citrixadc_clusternode.node1.state
    priority  = data.citrixadc_clusternode.node1.priority
  }
}
```

### Reference Node in Configuration

```hcl
data "citrixadc_clusternode" "coordinator" {
  nodeid = 0
}

# Use node information for validation or conditional logic
locals {
  is_coordinator_active = data.citrixadc_clusternode.coordinator.state == "ACTIVE"
}

output "coordinator_status" {
  value       = local.is_coordinator_active
  description = "Whether the coordinator node is active"
}
```

### Multiple Node Configuration

```hcl
data "citrixadc_clusternode" "node1" {
  nodeid = 1
}

data "citrixadc_clusternode" "node2" {
  nodeid = 2
}

output "cluster_nodes" {
  value = {
    node1 = {
      ipaddress = data.citrixadc_clusternode.node1.ipaddress
      state     = data.citrixadc_clusternode.node1.state
    }
    node2 = {
      ipaddress = data.citrixadc_clusternode.node2.ipaddress
      state     = data.citrixadc_clusternode.node2.state
    }
  }
}
```

## Notes

* The cluster node must already exist in the Citrix ADC cluster configuration before it can be retrieved using this data source.
* Cluster nodes are identified by their unique `nodeid`.
* The node state can be ACTIVE, SPARE, or PASSIVE, each serving different purposes in cluster operation.
* Priority values determine the configuration coordinator selection in a cluster setup.
* This data source requires a CLUSTER testbed environment for proper operation.
