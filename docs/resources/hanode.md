---
subcategory: "High Availability"
---

# Resource: hanode

The hanode resource is used to create hanode.


## Example usage

```hcl
resource "citrixadc_hanode" "local_node" {
  hanode_id     = 0       //the id of local_node is always 0
  hellointerval = 400
  deadinterval = 30
}

resource "citrixadc_hanode" "remote_node" {
  //we can only add remote add and not configure anything
  //if we want to configure this, we have to connect to the remote_node NSIP
  hanode_id = 3          
  ipaddress = "10.222.74.145"
}
```


## Argument Reference

* `hanode_id` - (Required) Number that uniquely identifies the node. For self node, it will always be 0. Peer node values can range from 1-64.
* `deadinterval` - (Optional) Number of seconds after which a peer node is marked DOWN if heartbeat messages are not received from the peer node.
* `failsafe` - (Optional) Keep one node primary if both nodes fail the health check, so that a partially available node can back up data and handle traffic. This mode is set independently on each node.
* `haprop` - (Optional) Automatically propagate all commands from the primary to the secondary node, except the following: * All HA configuration related commands. For example, add ha node, set ha node, and bind ha node.  * All Interface related commands. For example, set interface and unset interface. * All channels related commands. For example, add channel, set channel, and bind channel. The propagated command is executed on the secondary node before it is executed on the primary. If command propagation fails, or if command execution fails on the secondary, the primary node executes the command and logs an error.  Command propagation uses port 3010. Note: After enabling propagation, run force synchronization on either node.
* `hastatus` - (Optional) The HA status of the node. The HA status STAYSECONDARY is used to force the secondary device stay as secondary independent of the state of the Primary device. For example, in an existing HA setup, the Primary node has to be upgraded and this process would take few seconds. During the upgradation, it is possible that the Primary node may suffer from a downtime for a few seconds. However, the Secondary should not take over as the Primary node. Thus, the Secondary node should remain as Secondary even if there is a failure in the Primary node. 	 STAYPRIMARY configuration keeps the node in primary state in case if it is healthy, even if the peer node was the primary node initially. If the node with STAYPRIMARY setting (and no peer node) is added to a primary node (which has this node as the peer) then this node takes over as the new primary and the older node becomes secondary. ENABLED state means normal HA operation without any constraints/preferences. DISABLED state disables the normal HA operation of the node.
* `hasync` - (Optional) Automatically maintain synchronization by duplicating the configuration of the primary node on the secondary node. This setting is not propagated. Automatic synchronization requires that this setting be enabled (the default) on the current secondary node. Synchronization uses TCP port 3010.
* `hellointerval` - (Optional) Interval, in milliseconds, between heartbeat messages sent to the peer node. The heartbeat messages are UDP packets sent to port 3003 of the peer node.
* `inc` - (Optional) This option is required if the HA nodes reside on different networks. When this mode is enabled, the following independent network entities and configurations are neither propagated nor synced to the other node: MIPs, SNIPs, VLANs, routes (except LLB routes), route monitors, RNAT rules (except any RNAT rule with a VIP as the NAT IP), and dynamic routing configurations. They are maintained independently on each node.
* `ipaddress` - (Optional) The NSIP or NSIP6 address of the node to be added for an HA configuration. This setting is neither propagated nor synchronized.
* `maxflips` - (Optional) Max number of flips allowed before becoming sticky primary
* `maxfliptime` - (Optional) Interval after which flipping of node states can again start
* `syncstatusstrictmode` - (Optional) strict mode flag for sync status
* `syncvlan` - (Optional) Vlan on which HA related communication is sent. This include sync, propagation , connection mirroring , LB persistency config sync, persistent session sync and session state sync. However HA heartbeats can go all interfaces.
* `rpcnodepassword` - (Optional) Password to be used in authentication with the peer rpc node.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the hanode. It has the same value as the `hanode_id` attribute.


## Import

A hanode can be imported using its name, e.g.

```shell
terraform import citrixadc_hanode.local_node 0
```
