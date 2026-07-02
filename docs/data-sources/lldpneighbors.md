---
subcategory: "LLDP"
---

# Data Source: lldpneighbors

The lldpneighbors data source retrieves information about the Link Layer Discovery Protocol (LLDP) neighbors that the Citrix ADC has learned on its interfaces. It backs the NITRO `show lldp neighbors` (get-all) call and returns the read-only neighbor telemetry the ADC has discovered from directly connected devices.

An empty result is valid: if no LLDP peers have been learned (for example, because LLDP is not enabled on any interface or no neighboring device advertises LLDP), the neighbor list is empty and no error is raised. LLDP must be enabled per-interface for neighbor information to be learned.


## Example usage

```terraform
data "citrixadc_lldpneighbors" "tf_lldpneighbors" {
}

output "neighbor_ifnum" {
  value = data.citrixadc_lldpneighbors.tf_lldpneighbors.ifnum
}
```

Filter to the neighbor learned on a specific interface:

```terraform
data "citrixadc_lldpneighbors" "on_interface" {
  ifnum = "1/1"
}

output "neighbor_on_interface" {
  value = data.citrixadc_lldpneighbors.on_interface.ifnum
}
```

On a cluster, scope the lookup to a specific node:

```terraform
data "citrixadc_lldpneighbors" "on_node" {
  nodeid = 0
}
```


## Argument Reference

Both arguments are optional. When neither is supplied, the first learned neighbor (if any) is returned.

* `ifnum` - (Optional) Interface name to filter by (for example, `"1/1"`). When set, only the neighbor learned on that interface is returned.
* `nodeid` - (Optional) Unique number that identifies the cluster node to query. Applicable in a cluster deployment.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lldpneighbors data source. It is a fixed synthetic value, `lldpneighbors`.
* `ifnum` - Interface name on which the LLDP neighbor was learned.
* `nodeid` - Unique number that identifies the cluster node from which the neighbor information was read.
* `chassisidsubtype` - Chassis ID subtype of the LLDP neighbor.
* `chassisid` - Chassis ID of the LLDP neighbor.
* `portidsubtype` - Port ID subtype of the LLDP neighbor.
* `portid` - Port ID of the LLDP neighbor.
* `ttl` - Time to live of the LLDP neighbor advertisement.
* `portdescription` - Port description of the LLDP neighbor.
* `sys` - System name of the LLDP neighbor.
* `sysdesc` - System description of the LLDP neighbor.
* `mgmtaddresssubtype` - Management address subtype of the LLDP neighbor.
* `mgmtaddress` - Management address of the LLDP neighbor.
* `iftype` - Interface type of the LLDP neighbor.
* `ifnumber` - Interface number of the LLDP neighbor.
* `vlan` - VLAN of the LLDP neighbor.
* `vlanid` - VLAN ID of the LLDP neighbor.
* `portprotosupported` - Port protocol VLANs supported by the LLDP neighbor.
* `portprotoenabled` - Port protocol VLANs enabled on the LLDP neighbor.
* `portprotoid` - Port protocol VLAN ID of the LLDP neighbor.
* `portvlanid` - Port VLAN ID of the LLDP neighbor.
* `protocolid` - Protocol ID of the LLDP neighbor.
* `linkaggrcapable` - Whether the LLDP neighbor is link-aggregation capable.
* `linkaggrenabled` - Whether link aggregation is enabled on the LLDP neighbor.
* `linkaggrid` - Link aggregation ID of the LLDP neighbor.
* `flag` - Flag of the LLDP neighbor entry.
* `syscapabilities` - System capabilities of the LLDP neighbor.
* `syscapenabled` - Enabled system capabilities of the LLDP neighbor.
* `autonegsupport` - Whether auto-negotiation is supported by the LLDP neighbor.
* `autonegenabled` - Whether auto-negotiation is enabled on the LLDP neighbor.
* `autonegadvertised` - Auto-negotiation capabilities advertised by the LLDP neighbor.
* `autonegmautype` - Auto-negotiation MAU type of the LLDP neighbor.
* `mtu` - MTU of the LLDP neighbor.
