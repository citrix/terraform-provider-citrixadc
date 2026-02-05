---
subcategory: "Cluster"
---

# Data Source `clusterinstance`

The clusterinstance data source allows you to retrieve information about a Citrix ADC cluster instance configuration.


## Example usage

```terraform
data "citrixadc_clusterinstance" "tf_clusterinstance" {
  clid = 1
}

output "clusterinstance_id" {
  value = data.citrixadc_clusterinstance.tf_clusterinstance.id
}
```


## Argument Reference

* `clid` - (Required) Unique number that identifies the cluster.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusterinstance. It is the same as the `clid` attribute.
* `backplanebasedview` - View based on heartbeat only on bkplane interface.
* `clusterproxyarp` - This field controls the proxy arp feature in cluster. By default the flag is enabled.
* `deadinterval` - Amount of time, in seconds, after which nodes that do not respond to the heartbeats are assumed to be down.If the value is less than 3 sec, set the helloInterval parameter to 200 msec.
* `dfdretainl2params` - flag to add ext l2 header during steering. By default the flag is disabled.
* `hellointerval` - Interval, in milliseconds, at which heartbeats are sent to each cluster node to check the health status.Set the value to 200 msec, if the deadInterval parameter is less than 3 sec.
* `inc` - This option is required if the cluster nodes reside on different networks.
* `nodegroup` - The node group in a Cluster system used for transition from L2 to L3.
* `preemption` - Preempt a cluster node that is configured as a SPARE if an ACTIVE node becomes available.
* `processlocal` - By turning on this option packets destined to a service in a cluster will not under go any steering.
* `quorumtype` - Quorum Configuration Choices  - "Majority" (recommended) requires majority of nodes to be online for the cluster to be UP. "None" relaxes this requirement.
* `retainconnectionsoncluster` - This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled.
* `secureheartbeats` - By turning on this option cluster heartbeats will have security enabled.
* `syncstatusstrictmode` - strict mode for sync status of cluster. Depending on the the mode if there are any errors while applying config, sync status is displayed accordingly. By default the flag is disabled.

## Import

A clusterinstance can be imported using its clid, e.g.

```shell
terraform import citrixadc_clusterinstance.tf_clusterinstance 1
```
