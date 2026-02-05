---
subcategory: "Network"
---

# Data Source: citrixadc_rnat

This data source retrieves information about a specific RNAT (Reverse NAT) rule.

## Example Usage

```hcl
data "citrixadc_rnat" "example" {
  name = "my_rnat_rule"
}

output "rnat_network" {
  value = data.citrixadc_rnat.example.network
}
```

## Argument Reference

* `name` - (Required) Name of the RNAT4 rule.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the RNAT rule.
* `network` - The network address defined for the RNAT entry.
* `netmask` - The subnet mask for the network address.
* `aclname` - An extended ACL defined for the RNAT entry.
* `connfailover` - Synchronize connection-related information for RNAT sessions with the secondary ADC in an HA pair.
* `natip` - NetScaler-owned IPv4 address used to replace source IP addresses of server-generated packets.
* `redirectport` - Port number to which the IPv4 packets are redirected (TCP/UDP protocols).
* `srcippersistency` - Enables using the same NAT IP address for all RNAT sessions from a particular server.
* `td` - Traffic domain ID.
* `useproxyport` - Enable source port proxying for RNAT IPs.
* `ownergroup` - The owner node group in a Cluster for this rnat rule.
* `newname` - New name for the RNAT4 rule.
