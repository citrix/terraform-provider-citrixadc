---
subcategory: "Cluster"
---

# Resource: cluster

The resource is used to create and modify Citrix ADC in cluster deployment.

If there is no cluster instantiated the resource will bootstrap
the first node so that the Cluster IP address is reachable
and then add cluster nodes according to the terraform execution plan.

If there already exists a cluster deployment with the Cluster IP address
reachable the resource will add, remove or modify cluster nodes
according to the terraform execution plan.

In both cases the Citrix ADC provider configuration should point to the
Cluster IP address.


## Example usage

```hcl
# cluster L2 deployment
resource "citrixadc_cluster" "tf_cluster" {
  clid          = 1
  clip          = "10.78.60.15"
  hellointerval = 200

  clusternode {
    nodeid     = 0
    delay      = 0
    priority   = 30
    endpoint   = "http://10.78.60.10"
    backplane  = "0/1/1"
    ipaddress  = "10.78.60.10"
    tunnelmode = "NONE"
    nodegroup  = "DEFAULT_NG"

    state = "ACTIVE"
  }

  clusternode {
    nodeid     = 1
    delay      = 0
    priority   = 31
    endpoint   = "http://10.78.60.11"
    ipaddress  = "10.78.60.11"
    backplane  = "1/1/1"
    tunnelmode = "NONE"
    nodegroup  = "DEFAULT_NG"

    state = "ACTIVE"
  }

}

# cluster L3 deployment
resource "citrixadc_cluster" "tf_cluster" {
  clid          = 1
  clip          = "192.168.5.55"
  hellointerval = 200
  inc           = "ENABLED"

  clusternodegroup {
    name   = "ng0"
    strict = "YES"
  }


  clusternode {
    nodeid     = 0
    delay      = 0
    priority   = 30
    endpoint   = "http://192.168.5.127"
    ipaddress  = "192.168.5.127"
    tunnelmode = "GRE"
    nodegroup  = "ng0"

    state = "ACTIVE"
  }

  clusternodegroup {
    name   = "ng1"
    strict = "YES"
  }
  clusternode {
    nodeid     = 1
    delay      = 0
    priority   = 31
    endpoint   = "http://192.168.7.146"
    ipaddress  = "192.168.7.146"
    tunnelmode = "GRE"
    nodegroup  = "ng1"

    state = "ACTIVE"
  }
  clusternodegroup {
    name   = "ng2"
    strict = "YES"
  }

  clusternode {
    nodeid     = 2
    delay      = 0
    priority   = 31
    endpoint   = "http://192.168.6.9"
    ipaddress  = "192.168.6.9"
    tunnelmode = "GRE"
    nodegroup  = "ng2"

    state = "ACTIVE"
  }
}
```


## Argument Reference

* `clip` - (Required) Cluster IP address. It will be added on cluster bootstrap on the first cluster node.
* `clid` - (Required) Unique number that identifies the cluster.
* `deadinterval` - (Optional) Amount of time, in seconds, after which nodes that do not respond to the heartbeats are assumed to be down.If the value is less than 3 sec, set the `hellointerval` parameter to 200 msec.
* `hellointerval` - (Optional) Interval, in milliseconds, at which heartbeats are sent to each cluster node to check the health status.Set the value to 200 msec, if the `deadinterval` parameter is less than 3 sec.
* `preemption` - (Optional) Preempt a cluster node that is configured as a SPARE if an ACTIVE node becomes available. Possible values: [ ENABLED, DISABLED ]
* `quorumtype` - (Optional) Quorum Configuration Choices - "MAJORITY" (recommended) requires majority of nodes to be online for the cluster to be UP. "NONE" relaxes this requirement. Possible values: [ MAJORITY, NONE ]
* `inc` - (Optional) This option is required if the cluster nodes reside on different networks. Possible values: [ ENABLED, DISABLED ]
* `processlocal` - (Optional) By turning on this option packets destined to a service in a cluster will not under go any steering. Possible values: [ ENABLED, DISABLED ]
* `retainconnectionsoncluster` - (Optional) This option enables you to retain existing connections on a node joining a Cluster system or when a node is being configured for passive timeout. By default, this option is disabled. Possible values: [ YES, NO ]
* `backplanebasedview` - (Optional) View based on heartbeat only on bkplane interface. Possible values: [ ENABLED, DISABLED ]
* `nodegroup` - (Optional) The node group in a Cluster system used for transition from L2 to L3.
* `bootstrap_poll_delay` - (Optional) Time duration to wait before the first poll for the Cluster IP during first node bootstrap. Defaults to `60s`.
* `bootstrap_poll_interval` - (Optional) Time duration between subsequent polls for the Cluster IP during first node bootstrap. Defaults to `60s`.
* `bootstrap_poll_timeout` - (Optional) Time duration that defines http request timeout for each Cluster IP poll during first node bootstrap. Defaults to `10s`.
* `bootstrap_total_timeout` - (Optional) Time duration that defines the timeout for the whole operation of the first node bootstrap. If the node has not been added when this timeout expires the operation will fail. Defaults to `10m`.
* `clip_migration_poll_delay` - (Optional) Time duration to wait before the first poll for the Cluster IP during a Cluster IP address migration. Defaults to `10s`.
* `clip_migration_poll_interval` - (Optional) Time duration between subsequent polls for the Cluster IP during a Cluster IP address migration. Defaults to `30s`.
* `clip_migration_poll_timeout` - (Optional) Time duration that defines http request timeout for each Cluster IP poll during a Cluster IP address migration. Defaults to `10s`.
* `clip_migration_total_timeout` - (Optional) Time duration that defines the timeout for the whole operation of Cluster IP address migration. If the node has not been added when this timeout expires the operation will fail. Defaults to `10m`.
* `node_add_poll_delay` - (Optional) Time duration to wait before the first poll for the Cluster IP during the addition of a node to the cluster. Defaults to `10s`.
* `node_add_poll_interval` - (Optional) Time duration between subsequent polls for the Cluster IP during the addition of a node to the cluster. Defaults to `30s`.
* `node_add_total_timeout` - (Optional) Time duration that defines the timeout for the whole operation of the addition of a node to the cluster. If the node has not been added when this timeout expires the operation will fail. Defaults to `10m`.
* `clusternode` - (Required) Cluster node configuration blocks. Documented below.
* `clusternodegroup` - (Optional) One cluster nodegroup configuration block. Documented below.

Cluster node supports the following:

* `nodeid` - (Required) Unique number that identifies the cluster node.
* `ipaddress` - (Required) Citrix ADC IP (NSIP) address of the appliance to add to the cluster. Must be an IPv4 address.
* `state` - (Optional) Admin state of the cluster node. Possible values: [ ACTIVE, SPARE, PASSIVE ]
* `backplane` - (Optional) Interface through which the node communicates with the other nodes in the cluster. Must be specified in the three-tuple form n/c/u, where n represents the node ID and c/u refers to the interface on the appliance.
* `priority` - (Optional) Preference for selecting a node as the configuration coordinator. The node with the lowest priority value is selected as the configuration coordinator.

    When the current configuration coordinator goes down, the node with the next lowest priority is made the new configuration coordinator. When the original node comes back up, it will preempt the new configuration coordinator and take over as the configuration coordinator.

    -> When priority is not configured for any of the nodes or if multiple nodes have the same priority, the cluster elects one of the nodes as the configuration coordinator.

* `nodegroup` - (Optional) The default node group in a Cluster system.
* `delay` - (Optional) Applicable for Passive node and node becomes passive after this timeout (in minutes).
* `tunnelmode` - (Optional) To set the tunnel mode. Possible values: [ NONE, GRE, UDP ]
* `clearnodegroupconfig` - (Optional) Option to remove nodegroup config. Possible values: [ YES, NO ]

    -> When a standalone ADC instance joins a cluster the provider will issue a series of NITRO API calls against this particular node. The following arguments apply to the NITRO API client that will be used for this one time operation.
* `endpoint` - (Required) Defines the NITRO API endpoint prefix. Can use either `http` or `https` protocol.
* `username` - (Optional) Defines the username that will be used by the NITRO API for authentication. Defaults to the value of the same argument of the provider currently in effect.
* `password` - (Required) Defines the password that will be used by the NITRO API for authentication. Defaults to the value of the same argument of the provider currently in effect.
* `insecure_skip_verify` - (Optional) Boolean variable that defines if an error should be thrown if the target ADC's TLS certificate is not trusted. When `true` the error will be ignored. When `false` such an error will cause the failure of any provider operation. Defaults to `false`.
* `addsnip` - (Optional) Boolean variable that determines if a node SNIP should be added to the CLIP before joining the cluster.
* `snip_ipaddress` - (Optional) Node SNIP address to add to the CLIP. Applied only when `addsnip=true`.
* `snip_netmask` - (Optional) Node SNIP netmask to add to the CLIP. Applied only when `addsnip=true`.
* `vtysh_enable` - (Optional) Boolean variable that determines if vtysh commands should be applied to the CLIP before node joins the custer.
* `vtysh` - (Optional) Vtysh commands to add to the CLIP before node joins the cluster. Applied only when `vtysh_enable=true`.

Cluster nodegroup supports the following:

* `name` - (Optional) Name of the nodegroup. The name uniquely identifies the nodegroup on the cluster.
* `strict` - (Optional) Specifies whether cluster nodes, that are not part of the nodegroup, will be used as backup for the nodegroup. Possible values: [ YES, NO ]
* `sticky` - (Optional) Only one node can be bound to nodegroup with this option enabled. It specifies whether to prempt the traffic for the entities bound to nodegroup when owner node goes down and rejoins the cluster. Possible values: [ YES, NO ]
* `state` - (Optional) State of the nodegroup. All the nodes binding to this nodegroup must have the same state. Possible values: [ ACTIVE, SPARE, PASSIVE ]
* `priority` (Optional) Priority of Nodegroup. This priority is used for all the nodes bound to the nodegroup for Nodegroup selection.



## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cluster. It has the same value as the `clid` attribute.
