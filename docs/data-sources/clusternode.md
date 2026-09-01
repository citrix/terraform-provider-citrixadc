---
subcategory: "Cluster"
---

# Data Source: clusternode

The clusternode data source allows you to retrieve information about a cluster node.


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

### Read-only clusternode metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_clusternode` resource). Any attribute the appliance does not return is `null`.

* `clusterhealth` - Node clusterd state.
* `effectivestate` - Node effective health state. Possible values: `UP`, `NOT UP`, `UNKNOWN`, `INIT`.
* `operationalsyncstate` - Node Operational Reconciliation state.
* `syncfailurereason` - Additional information along with cluster sync status.
* `masterstate` - Node Master state. Possible values: `INACTIVE`, `ACTIVE`, `UNKNOWN`.
* `health` - Node Health state.
* `syncstate` - Whether synchronization of cluster configurations on the node is enabled/disabled. Possible values: `ENABLED`, `DISABLED`.
* `isconfigurationcoordinator` - Whether the node is configuration coordinator (CCO).
* `islocalnode` - Whether it is the local node.
* `nodersskeymismatch` - Whether there is an RSS key mismatch at cluster node level.
* `nodelicensemismatch` - Whether there is a License mismatch at cluster node level.
* `nodejumbonotsupported` - Whether the Jumbo framework is not supported at cluster node level.
* `nodelist` - Nodelist for displaying Heartbeat not seen interfaces on a cluster node. A list of strings.
* `ifaceslist` - Interface list corresponding to `nodelist` for Heartbeat not seen interfaces on a cluster node. A list of strings.
* `enabledifaces` - Enabled Interfaces on a cluster node.
* `disabledifaces` - Disabled Interfaces on a cluster node.
* `partialfailifaces` - Partial Failure Interfaces on a cluster node.
* `hamonifaces` - Hamon Interfaces on a cluster node.
* `name` - Name of the state specific nodegroup.
* `cfgflags` - Flag indicating whether the node is bound to a cluster nodegroup.
* `routemonitor` - The IP address (IPv4 or IPv6).
* `netmask` - The netmask.

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
