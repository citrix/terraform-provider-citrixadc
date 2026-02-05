---
subcategory: "Network"
---

# Data Source: citrixadc_rnat6

This data source retrieves information about a specific RNAT6 (Reverse NAT for IPv6) rule.

## Example Usage

```hcl
data "citrixadc_rnat6" "example" {
  name = "my_rnat6_rule"
}

output "rnat6_network" {
  value = data.citrixadc_rnat6.example.network
}
```

## Argument Reference

* `name` - (Required) Name of the RNAT6 rule.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the RNAT6 rule.
* `network` - IPv6 address of the network on whose traffic the Citrix ADC performs RNAT processing.
* `acl6name` - Name of any configured ACL6 whose action is ALLOW. The rule of the ACL6 is used as an RNAT6 rule.
* `redirectport` - Port number to which the IPv6 packets are redirected (TCP/UDP protocols).
* `srcippersistency` - Enable source IP persistency, which enables the Citrix ADC to use the RNAT IPs using source IP.
* `td` - Traffic domain ID.
* `ownergroup` - The owner node group in a Cluster for this rnat rule.
