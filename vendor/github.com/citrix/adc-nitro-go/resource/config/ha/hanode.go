/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
*/

package ha

/**
* Configuration for node resource.
*/
type Hanode struct {
	/**
	* Number that uniquely identifies the node. For self node, it will always be 0. Peer node values can range from 1-64.
	*/
	Id int `json:"id,omitempty"`
	/**
	* The NSIP or NSIP6 address of the node to be added for an HA configuration. This setting is neither propagated nor synchronized.
	*/
	Ipaddress string `json:"ipaddress,omitempty"`
	/**
	* This option is required if the HA nodes reside on different networks. When this mode is enabled, the following independent network entities and configurations are neither propagated nor synced to the other node: MIPs, SNIPs, VLANs, routes (except LLB routes), route monitors, RNAT rules (except any RNAT rule with a VIP as the NAT IP), and dynamic routing configurations. They are maintained independently on each node.
	*/
	Inc string `json:"inc,omitempty"`
	/**
	* The HA status of the node. The HA status STAYSECONDARY is used to force the secondary device stay as secondary independent of the state of the Primary device. For example, in an existing HA setup, the Primary node has to be upgraded and this process would take few seconds. During the upgradation, it is possible that the Primary node may suffer from a downtime for a few seconds. However, the Secondary should not take over as the Primary node. Thus, the Secondary node should remain as Secondary even if there is a failure in the Primary node.
		STAYPRIMARY configuration keeps the node in primary state in case if it is healthy, even if the peer node was the primary node initially. If the node with STAYPRIMARY setting (and no peer node) is added to a primary node (which has this node as the peer) then this node takes over as the new primary and the older node becomes secondary. ENABLED state means normal HA operation without any constraints/preferences. DISABLED state disables the normal HA operation of the node.
	*/
	Hastatus string `json:"hastatus,omitempty"`
	/**
	* Automatically maintain synchronization by duplicating the configuration of the primary node on the secondary node. This setting is not propagated. Automatic synchronization requires that this setting be enabled (the default) on the current secondary node. Synchronization uses TCP port 3010.
	*/
	Hasync string `json:"hasync,omitempty"`
	/**
	* Automatically propagate all commands from the primary to the secondary node, except the following:
		* All HA configuration related commands. For example, add ha node, set ha node, and bind ha node. 
		* All Interface related commands. For example, set interface and unset interface.
		* All channels related commands. For example, add channel, set channel, and bind channel.
		The propagated command is executed on the secondary node before it is executed on the primary. If command propagation fails, or if command execution fails on the secondary, the primary node executes the command and logs an error.  Command propagation uses port 3010.
		Note: After enabling propagation, run force synchronization on either node.
	*/
	Haprop string `json:"haprop,omitempty"`
	/**
	* Interval, in milliseconds, between heartbeat messages sent to the peer node. The heartbeat messages are UDP packets sent to port 3003 of the peer node.
	*/
	Hellointerval int `json:"hellointerval,omitempty"`
	/**
	* Number of seconds after which a peer node is marked DOWN if heartbeat messages are not received from the peer node.
	*/
	Deadinterval int `json:"deadinterval,omitempty"`
	/**
	* Keep one node primary if both nodes fail the health check, so that a partially available node can back up data and handle traffic. This mode is set independently on each node.
	*/
	Failsafe string `json:"failsafe,omitempty"`
	/**
	* Max number of flips allowed before becoming sticky primary
	*/
	Maxflips int `json:"maxflips,omitempty"`
	/**
	* Interval after which flipping of node states can again start
	*/
	Maxfliptime int `json:"maxfliptime,omitempty"`
	/**
	* Vlan on which HA related communication is sent. This include sync, propagation , connection mirroring , LB persistency config sync, persistent session sync and session state sync. However HA heartbeats can go all interfaces.
	*/
	Syncvlan int `json:"syncvlan,omitempty"`
	/**
	* strict mode flag for sync status
	*/
	Syncstatusstrictmode string `json:"syncstatusstrictmode,omitempty"`

	//------- Read only Parameter ---------;

	Name string `json:"name,omitempty"`
	Flags string `json:"flags,omitempty"`
	State string `json:"state,omitempty"`
	Enaifaces string `json:"enaifaces,omitempty"`
	Disifaces string `json:"disifaces,omitempty"`
	Hamonifaces string `json:"hamonifaces,omitempty"`
	Haheartbeatifaces string `json:"haheartbeatifaces,omitempty"`
	Pfifaces string `json:"pfifaces,omitempty"`
	Ifaces string `json:"ifaces,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Ssl2 string `json:"ssl2,omitempty"`
	Masterstatetime string `json:"masterstatetime,omitempty"`
	Routemonitor string `json:"routemonitor,omitempty"`
	Curflips string `json:"curflips,omitempty"`
	Completedfliptime string `json:"completedfliptime,omitempty"`
	Routemonitorstate string `json:"routemonitorstate,omitempty"`
	Hasyncfailurereason string `json:"hasyncfailurereason,omitempty"`

}
