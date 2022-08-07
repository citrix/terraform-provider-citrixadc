---
subcategory: "Cluster"
---

# Resource: clusterinstance

The clusterinstanceresource is used to create clusterinstance.


## Example usage

```hcl
resource "citrixadc_clusterinstance" "tf_clusterinstance" {
  clid          = 1
  deadinterval  = 8
  hellointerval = 600
}
```


## Argument Reference

* `clid` - (Required) Unique number that identifies the cluster. Minimum value =  1 Maximum value =  16
* `deadinterval` - (Optional) Amount of time, in seconds, after which nodes that do not respond to the heartbeats are assumed to be down.If the value is less than 3 sec, set the helloInterval parameter to 200 msec. Minimum value =  1 Maximum value =  60
* `hellointerval` - (Optional) Interval, in milliseconds, at which heartbeats are sent to each cluster node to check the health status.Set the value to 200 msec, if the deadInterval parameter is less than 3 sec. Minimum value =  200 Maximum value =  1000
* `preemption` - (Optional) Preempt a cluster node that is configured as a SPARE if an ACTIVE node becomes available. Possible values: [ ENABLED, DISABLED ]
* `quorumtype` - (Optional) Quorum Configuration Choices  - "Majority" (recommended) requires majority of nodes to be online for the cluster to be UP. "None" relaxes this requirement. Possible values: [ MAJORITY, NONE ]
* `inc` - (Optional) This option is required if the cluster nodes reside on different networks. Possible values: [ ENABLED, DISABLED ]
* `processlocal` - (Optional) By turning on this option packets destined to a service in a cluster will not under go any steering. Possible values: [ ENABLED, DISABLED ]
* `retainconnectionsoncluster` - (Optional) This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled. Possible values: [ YES, NO ]
* `backplanebasedview` - (Optional) View based on heartbeat only on bkplane interface. Possible values: [ ENABLED, DISABLED ]
* `syncstatusstrictmode` - (Optional) strict mode for sync status of cluster. Depending on the the mode if there are any errors while applying config, sync status is displayed accordingly. By default the flag is disabled. Possible values: [ ENABLED, DISABLED ]
* `nodegroup` - (Optional) The node group in a Cluster system used for transition from L2 to L3.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusterinstance. It has the same value as the `clid` attribute.


## Import

A clusterinstance can be imported using its name, e.g.

```shell
terraform import citrixadc_clusterinstance.tf_clusterinstance tf_csaction 1
```
