---
subcategory: "Network"
---

# Data Source: iptunnel

The iptunnel data source allows you to retrieve information about IP tunnels configured on the Citrix ADC.

## Example Usage

```terraform
data "citrixadc_iptunnel" "example" {
  name             = "my_iptunnel"
  remote           = "66.0.0.11"
  remotesubnetmask = "255.255.255.255"
}

output "protocol" {
  value = data.citrixadc_iptunnel.example.protocol
}

output "vnid" {
  value = data.citrixadc_iptunnel.example.vnid
}
```

## Argument Reference

The following arguments are required:

* `name` - (Required) Name for the IP tunnel. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).
* `remote` - (Required) Public IPv4 address, of the remote device, used to set up the tunnel. For this parameter, you can alternatively specify a network address.
* `remotesubnetmask` - (Required) Subnet mask of the remote IP address of the tunnel.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the iptunnel. It is a composite of `name`, `remote`, and `remotesubnetmask` separated by commas.
* `destport` - Specifies UDP destination port for Geneve packets. Default port is 6081.
* `grepayload` - The payload GRE will carry.
* `ipsecprofilename` - Name of IPSec profile to be associated.
* `local` - Type of Citrix ADC owned public IPv4 address, configured on the local Citrix ADC and used to set up the tunnel.
* `ownergroup` - The owner node group in a Cluster for the iptunnel.
* `protocol` - Name of the protocol to be used on this tunnel.
* `tosinherit` - Default behavior is to copy the ToS field of the internal IP Packet (Payload) to the outer IP packet (Transport packet). But the user can configure a new ToS field using this option.
* `vlan` - The vlan for multicast packets.
* `vlantagging` - Option to select Vlan Tagging.
* `vnid` - Virtual network identifier (VNID) is the value that identifies a specific virtual network in the data plane.

### Read-only iptunnel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_iptunnel` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `sysname` - The name of the ip tunnel.
* `type` - The type of this tunnel.
* `encapip` - The effective local IP address of the tunnel. Used as the source of the encapsulated packets.
* `channel` - The tunnel that is bound to a netbridge.
* `tunneltype` - Indicates that a tunnel is User-Configured, Internal or DELETE-IN-PROGRESS. A list of strings.
* `ipsectunnelstatus` - Whether the ipsec on this tunnel is up or down.
* `refcnt` - Number of PBRs to bound to this iptunnel.
