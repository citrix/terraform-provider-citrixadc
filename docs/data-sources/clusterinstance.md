---
subcategory: "Cluster"
---

# Data Source: clusterinstance

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

### Read-only clusterinstance metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_clusterinstance` resource). Any attribute the appliance does not return is `null`.

* `adminstate` - Cluster Admin State. Possible values: `ENABLED`, `DISABLED`.
* `propstate` - Whether execution of commands on the cluster is enabled/disabled. This does not impact command execution on individual cluster nodes by using the NSIP.
* `validmtu` - Correct MTU value that has to be set on the backplane.
* `heterogeneousflag` - Indicates whether heterogeneity is detected in the cluster system. Possible values: `YES`, `NO`.
* `operationalstate` - Cluster Operational State. Possible values: `ENABLED`, `DISABLED`.
* `status` - Cluster operational status. Possible values: `DOWN`, `UP`, `PARTIAL-UP`, `UNKNOWN`.
* `rsskeymismatch` - Whether there is an RSS key mismatch at cluster instance level.
* `penummismatch` - Whether there is a PE number mismatch at cluster instance level.
* `nodegroupstatewarning` - Whether all the cluster nodes are bound to a nodegroup with state set.
* `licensemismatch` - Whether there is a License mismatch at cluster instance level.
* `jumbonotsupported` - Whether the Jumbo framework is not supported at cluster instance level.
* `clustertunnelmodemismatch` - Whether a different tunnel mode is configured on cluster nodes.
* `clusternoheartbeatonnode` - Whether heartbeat is not seen on the backplane interface of a member node.
* `clusternolinksetmbf` - Whether MBF is enabled but linkset is not configured.
* `clusternospottedip` - Whether there is no spotted SNIP or MIP.
* `clusterclipfailure` - Whether CLIP movement failed (CLIP is not attached to CCO).
* `clusterhbhmacerrordetected` - Whether a cluster heartbeat HMAC error was detected (could be due to version mismatch).
* `nodepenummismatch` - Whether there is a PE mismatch at cluster node level.
* `operationalpropstate` - Cluster Operational Propagation State. Possible values: `UNKNOWN`, `ENABLED`, `DISABLED`, `AUTO DISABLED`, `AUTO DISABLED (Disk Encryption Mismatch)`.
