---
subcategory: "Network"
---

# Data Source: citrixadc_bridgetable

Use this data source to retrieve information about an existing Bridge Table entry.

The `citrixadc_bridgetable` data source allows you to retrieve details of a bridge table entry. Bridge table entries map MAC addresses to VXLAN tunnel endpoints (VTEPs) in a VXLAN network configuration.

## Example usage

```hcl
# Retrieve an existing bridge table entry
data "citrixadc_bridgetable" "example" {
  mac   = "00:00:00:00:00:01"
  vxlan = 123
  vtep  = "2.34.5.6"
}

# Reference bridge table attributes
output "bridge_age" {
  value = data.citrixadc_bridgetable.example.bridgeage
}
```

## Argument Reference

The following arguments are supported:

* `mac` - (Optional) The MAC address of the target.

* `vxlan` - (Optional) The VXLAN to which this address is associated.

* `vtep` - (Optional) The IP address of the destination VXLAN tunnel endpoint where the Ethernet MAC ADDRESS resides.

* `vlan` - (Optional) VLAN whose entries are to be removed.

* `ifnum` - (Optional) INTERFACE whose entries are to be removed.

* `nodeid` - (Optional) Unique number that identifies the cluster node.

* `vni` - (Optional) The VXLAN VNI Network Identifier (or VXLAN Segment ID) to use to connect to the remote VXLAN tunnel endpoint. If omitted the value specified as vxlan will be used.

* `devicevlan` - (Optional) The vlan on which to send multicast packets when the VXLAN tunnel endpoint is a muticast group address.

* `bridgeage` - (Optional) Time-out value for the bridge table entries, in seconds. The new value applies only to the entries that are dynamically learned after the new value is set. Previously existing bridge table entries expire after the previously configured time-out value.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the bridge table entry. For bridgetable, this is a composite key of `mac,vxlan,vtep`.

## Common Use Cases

### Retrieve Bridge Table Entry for VXLAN Configuration

```hcl
# Retrieve bridge table entry for a specific MAC in a VXLAN
data "citrixadc_bridgetable" "vxlan_entry" {
  mac   = "00:00:00:00:00:01"
  vxlan = 123
  vtep  = "2.34.5.6"
}

# Use in other configurations
output "bridge_table_id" {
  value = data.citrixadc_bridgetable.vxlan_entry.id
}
```

### Query Bridge Table for Monitoring

```hcl
# Retrieve bridge table to monitor VXLAN learning
data "citrixadc_bridgetable" "monitor" {
}

output "bridge_mac" {
  value       = data.citrixadc_bridgetable.monitor.mac
  description = "MAC address learned in bridge table"
}

output "bridge_vtep" {
  value       = data.citrixadc_bridgetable.monitor.vtep
  description = "VTEP IP address for the MAC"
}
```

## Notes

* Bridge table entries are typically learned dynamically when VXLAN tunnels are configured.
* The datasource retrieves bridge table information from the Citrix ADC.
* Multiple bridge table entries may exist; filter appropriately using the optional parameters.
* The `bridgeage` parameter determines how long entries remain in the bridge table before timing out.
