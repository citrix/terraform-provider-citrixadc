---
subcategory: "High Availability"
---

# Data Source: hanode

The hanode data source allows you to retrieve information about High Availability (HA) nodes configured on the Citrix ADC appliance.

## Example usage

```terraform
data "citrixadc_hanode" "local_node" {
  hanode_id = 0
}

output "hellointerval" {
  value = data.citrixadc_hanode.local_node.hellointerval
}

output "deadinterval" {
  value = data.citrixadc_hanode.local_node.deadinterval
}

output "hastatus" {
  value = data.citrixadc_hanode.local_node.hastatus
}
```

## Argument Reference

* `hanode_id` - (Required) Number that uniquely identifies the node. For self node, it will always be 0. Peer node values can range from 1-64.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `deadinterval` - Number of seconds after which a peer node is marked DOWN if heartbeat messages are not received from the peer node.
* `failsafe` - Keep one node primary if both nodes fail the health check, so that a partially available node can back up data and handle traffic. This mode is set independently on each node.
* `haprop` - Automatically propagate all commands from the primary to the secondary node, except the following: All HA configuration related commands, All Interface related commands, All channels related commands. The propagated command is executed on the secondary node before it is executed on the primary.
* `hastatus` - The HA status of the node. The HA status STAYSECONDARY is used to force the secondary device stay as secondary independent of the state of the Primary device. STAYPRIMARY configuration keeps the node in primary state in case if it is healthy, even if the peer node was the primary node initially. ENABLED state means normal HA operation without any constraints/preferences. DISABLED state disables the normal HA operation of the node. Possible values: [ ENABLED, STAYSECONDARY, STAYPRIMARY, DISABLED ]
* `hasync` - Automatically maintain synchronization by duplicating the configuration of the primary node on the secondary node. This setting is not propagated. Automatic synchronization requires that this setting be enabled (the default) on the current secondary node.
* `hellointerval` - Interval, in milliseconds, between heartbeat messages sent to the peer node. The heartbeat messages are UDP packets sent to port 3003 of the peer node.
* `inc` - This option is required if the HA nodes reside on different networks. When this mode is enabled, the following independent network entities and configurations are neither propagated nor synced to the other node: MIPs, SNIPs, VLANs, routes (except LLB routes), route monitors, RNAT rules, and dynamic routing configurations.
* `ipaddress` - The NSIP or NSIP6 address of the node to be added for an HA configuration. This setting is neither propagated nor synchronized.
* `maxflips` - Max number of flips allowed before becoming sticky primary.
* `maxfliptime` - Interval after which flipping of node states can again start.
* `rpcnodepassword` - Password to be used in authentication with the peer rpc node.
* `syncstatusstrictmode` - Strict mode flag for sync status.
* `syncvlan` - Vlan on which HA related communication is sent. This include sync, propagation, connection mirroring, LB persistency config sync, persistent session sync and session state sync. However HA heartbeats can go all interfaces.
* `id` - The id of the hanode. It has the same value as the `hanode_id` attribute.

## Import

A hanode can be imported using its hanode_id, e.g.

```shell
terraform import citrixadc_hanode.local_node 0
```