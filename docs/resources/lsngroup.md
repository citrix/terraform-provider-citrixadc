---
subcategory: "Lsn"
---

# Resource: lsngroup

The lsngroup resource is used to create lsngroup.


## Example usage

```hcl
resource "citrixadc_lsngroup" "tf_lsngroup" {
  groupname     = "my_lsngroup"
  clientname    = "my_lsnclient"
  logging       = "DISABLED"
  nattype       = "DYNAMIC"
}

```


## Argument Reference

* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn group1" or 'lsn group1').
* `clientname` - (Required) Name of the LSN client entity to be associated with the LSN group. You can associate only one LSN client entity with an LSN group.You cannot remove this association or replace with another LSN client entity once the LSN group is created.
* `allocpolicy` - (Optional) NAT IP and PORT block allocation policy for Deterministic NAT. Supported Policies are, 1: PORTS: Port blocks from single NATIP will be allocated to LSN subscribers sequentially. After all blocks are exhausted, port blocks from next NATIP will be allocated and so on. 2: IPADDRS(Default): One port block from each NATIP will be allocated and once all the NATIPs are over second port block from each NATIP will be allocated and so on. To understand better if we assume port blocks of all NAT IPs as two dimensional array, PORTS policy follows "row major order" and IPADDRS policy follows "column major order" while allocating port blocks. Example: Client IPs: 2.2.2.1, 2.2.2.2 and 2.2.2.3 NAT IPs and PORT Blocks:  4.4.4.1:PB1, PB2, PB3,., PBn 4.4.4.2: PB1, PB2, PB3,., PBn PORTS Policy:  2.2.2.1 => 4.4.4.1:PB1 2.2.2.2 => 4.4.4.1:PB2 2.2.2.3 => 4.4.4.1:PB3 IPADDRS Policy: 2.2.2.1 => 4.4.4.1:PB1 2.2.2.2 => 4.4.4.2:PB1 2.2.2.3 => 4.4.4.1:PB2
* `ftp` - (Optional) Enable Application Layer Gateway (ALG) for the FTP protocol. For some application-layer protocols, the IP addresses and protocol port numbers are usually communicated in the packet's payload. When acting as an ALG, the Citrix ADC changes the packet's payload to ensure that the protocol continues to work over LSN.   Note:  The Citrix ADC also includes ALG for ICMP and TFTP protocols. ALG for the ICMP protocol is enabled by default, and there is no provision to disable it. ALG for the TFTP protocol is disabled by default. ALG is enabled automatically for an LSN group when you bind a UDP LSN application profile, with endpoint-independent-mapping, endpoint-independent filtering, and destination port as 69 (well-known port for TFTP), to the LSN group.
* `ftpcm` - (Optional) Enable the FTP connection mirroring for specified LSN group. Connection mirroring (CM or connection failover) refers to keeping active an established TCP or UDP connection when a failover occurs.
* `ip6profile` - (Optional) Name of the LSN ip6 profile to associate with the specified LSN group. An ip6 profile can be associated with a group only during group creation.  By default, no LSN ip6 profile is associated with an LSN group during its creation. Only one ip6profile can be associated with a group.
* `logging` - (Optional) Log mapping entries and sessions created or deleted for this LSN group. The Citrix ADC logs LSN sessions for this LSN group only when both logging and session logging parameters are enabled.  The ADC uses its existing syslog and audit log framework to log LSN information. You must enable global level LSN logging by enabling the LSN parameter in the related NSLOG action and SYLOG action entities. When the Logging parameter is enabled, the Citrix ADC generates log messages related to LSN mappings and LSN sessions of this LSN group. The ADC then sends these log messages to servers associated with the NSLOG action and SYSLOG actions entities.   A log message for an LSN mapping entry consists of the following information: * NSIP address of the Citrix ADC * Time stamp * Entry type (MAPPING or SESSION) * Whether the LSN mapping entry is created or deleted * Subscriber's IP address, port, and traffic domain ID * NAT IP address and port * Protocol name * Destination IP address, port, and traffic domain ID might be  present, depending on the following conditions: ** Destination IP address and port are not logged for Endpoint-Independent mapping ** Only Destination IP address (and not port) is logged for Address-Dependent mapping ** Destination IP address and port are logged for Address-Port-Dependent mapping
* `nattype` - (Optional) Type of NAT IP address and port allocation (from the bound LSN pools) for subscribers:  Available options function as follows:  * Deterministic - Allocate a NAT IP address and a block of ports to each subscriber (of the LSN client bound to the LSN group). The Citrix ADC sequentially allocates NAT resources to these subscribers. The Citrix ADC ADC assigns the first block of ports (block size determined by the port block size parameter of the LSN group) on the beginning NAT IP address to the beginning subscriber IP address. The next range of ports is assigned to the next subscriber, and so on, until the NAT address does not have enough ports for the next subscriber. In this case, the first port block on the next NAT address is used for the subscriber, and so on.  Because each subscriber now receives a deterministic NAT IP address and a block of ports, a subscriber can be identified without any need for logging. For a connection, a subscriber can be identified based only on the NAT IP address and port, and the destination IP address and port. The maximum number of LSN subscribers allowed, globally, is 1 million.    * Dynamic - Allocate a random NAT IP address and a port from the LSN NAT pool for a subscriber's connection. If port block allocation is enabled (in LSN pool) and a port block size is specified (in the LSN group), the Citrix ADC allocates a random NAT IP address and a block of ports for a subscriber when it initiates a connection for the first time. The ADC allocates this NAT IP address and a port (from the allocated block of ports) for different connections from this subscriber. If all the ports are allocated (for different subscriber's connections) from the subscriber's allocated port block, the ADC allocates a new random port block for the subscriber.
* `portblocksize` - (Optional) Size of the NAT port block to be allocated for each subscriber.  To set this parameter for Dynamic NAT, you must enable the port block allocation parameter in the bound LSN pool. For Deterministic NAT, the port block allocation parameter is always  enabled, and you cannot disable it.  In Dynamic NAT, the Citrix ADC allocates a random NAT port block, from the available NAT port pool of an NAT IP address, for each subscriber. For a subscriber, if all the ports are allocated from the subscriber's allocated port block, the ADC allocates a new random port block for the subscriber.  The default port block size is 256 for Deterministic NAT, and 0 for Dynamic NAT.
* `pptp` - (Optional) Enable the PPTP Application Layer Gateway.
* `rtspalg` - (Optional) Enable the RTSP ALG.
* `sessionlogging` - (Optional) Log sessions created or deleted for the LSN group. The Citrix ADC logs LSN sessions for this LSN group only when both logging and session logging parameters are enabled.  A log message for an LSN session consists of the following information: * NSIP address of the Citrix ADC * Time stamp * Entry type (MAPPING or SESSION) * Whether the LSN session is created or removed * Subscriber's IP address, port, and traffic domain ID * NAT IP address and port * Protocol name * Destination IP address, port, and traffic domain ID
* `sessionsync` - (Optional) In a high availability (HA) deployment, synchronize information of all LSN sessions related to this LSN group with the secondary node. After a failover, established TCP connections and UDP packet flows are kept active and resumed on the secondary node (new primary).  For this setting to work, you must enable the global session synchronization parameter.
* `sipalg` - (Optional) Enable the SIP ALG.
* `snmptraplimit` - (Optional) Maximum number of SNMP Trap messages that can be generated for the LSN group in one minute.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsngroup. It has the same value as the `groupname` attribute.


## Import

A lsngroup can be imported using its name, e.g.

```shell
terraform import citrixadc_lsngroup.tf_lsngroup my_lsngroup
```
