---
subcategory: "LSN"
---

# Data Source: lsngroup

The lsngroup data source allows you to retrieve information about an LSN (Large Scale NAT) group.

## Example usage

```terraform
data "citrixadc_lsngroup" "tf_lsngroup_ds" {
  groupname = "my_lsngroup_ds"
}

output "clientname" {
  value = data.citrixadc_lsngroup.tf_lsngroup_ds.clientname
}

output "nattype" {
  value = data.citrixadc_lsngroup.tf_lsngroup_ds.nattype
}

output "logging" {
  value = data.citrixadc_lsngroup.tf_lsngroup_ds.logging
}
```

## Argument Reference

* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `allocpolicy` - NAT IP and PORT block allocation policy for Deterministic NAT. Supported Policies are: 1: PORTS - Port blocks from single NATIP will be allocated to LSN subscribers sequentially. After all blocks are exhausted, port blocks from next NATIP will be allocated and so on. 2: IPADDRS(Default) - One port block from each NATIP will be allocated and once all the NATIPs are over second port block from each NATIP will be allocated and so on. Possible values: [ PORTS, IPADDRS ]
* `clientname` - Name of the LSN client entity to be associated with the LSN group. You can associate only one LSN client entity with an LSN group. You cannot remove this association or replace with another LSN client entity once the LSN group is created.
* `ftp` - Enable Application Layer Gateway (ALG) for the FTP protocol. For some application-layer protocols, the IP addresses and protocol port numbers are usually communicated in the packet's payload. When acting as an ALG, the Citrix ADC changes the packet's payload to ensure that the protocol continues to work over LSN. Possible values: [ ENABLED, DISABLED ]
* `ftpcm` - Enable the FTP connection mirroring for specified LSN group. Connection mirroring (CM or connection failover) refers to keeping active an established TCP or UDP connection when a failover occurs. Possible values: [ ENABLED, DISABLED ]
* `id` - The id of the lsngroup. It has the same value as the `groupname` attribute.
* `ip6profile` - Name of the LSN ip6 profile to associate with the specified LSN group. An ip6 profile can be associated with a group only during group creation. By default, no LSN ip6 profile is associated with an LSN group during its creation. Only one ip6profile can be associated with a group.
* `logging` - Log mapping entries and sessions created or deleted for this LSN group. The Citrix ADC logs LSN sessions for this LSN group only when both logging and session logging parameters are enabled. Possible values: [ ENABLED, DISABLED ]
* `nattype` - Type of NAT IP address and port allocation (from the bound LSN pools) for subscribers. Available options are: Deterministic - Allocate a NAT IP address and a block of ports to each subscriber. Dynamic - Allocate a random NAT IP address and a port from the LSN NAT pool for a subscriber's connection. Possible values: [ DYNAMIC, DETERMINISTIC ]
* `portblocksize` - Size of the NAT port block to be allocated for each subscriber. To set this parameter for Dynamic NAT, you must enable the port block allocation parameter in the bound LSN pool. For Deterministic NAT, the port block allocation parameter is always enabled. The default port block size is 256 for Deterministic NAT, and 0 for Dynamic NAT.
* `pptp` - Enable the PPTP Application Layer Gateway. Possible values: [ ENABLED, DISABLED ]
* `rtspalg` - Enable the RTSP ALG. Possible values: [ ENABLED, DISABLED ]
* `sessionlogging` - Log sessions created or deleted for the LSN group. The Citrix ADC logs LSN sessions for this LSN group only when both logging and session logging parameters are enabled. Possible values: [ ENABLED, DISABLED ]
* `sessionsync` - In a high availability (HA) deployment, synchronize information of all LSN sessions related to this LSN group with the secondary node. After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary node (new primary). Possible values: [ ENABLED, DISABLED ]
* `sipalg` - Enable the SIP ALG. Possible values: [ ENABLED, DISABLED ]
* `snmptraplimit` - Maximum number of SNMP Trap messages that can be generated for the LSN group in one minute.

## Import

A lsngroup can be imported using its groupname, e.g.

```shell
terraform import citrixadc_lsngroup.tf_lsngroup my_lsngroup_ds
```
