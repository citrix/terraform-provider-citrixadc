---
subcategory: "Network"
---

# Data Source `l4param`

The l4param data source allows you to retrieve information about L4 parameters configuration.


## Example usage

```terraform
data "citrixadc_l4param" "tf_l4param" {
}

output "l2connmethod" {
  value = data.citrixadc_l4param.tf_l4param.l2connmethod
}

output "l4switch" {
  value = data.citrixadc_l4param.tf_l4param.l4switch
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `l2connmethod` - Layer 2 connection method based on the combination of  channel number, MAC address and VLAN. It is tuned with l2conn param of lb vserver. If l2conn of lb vserver is ON then method specified here will be used to identify a connection in addition to the 4-tuple (<source IP>:<source port>::<destination IP>:<destination port>). Possible values: `Channel`, `Vlan`, `VlanChannel`, `Mac`, `MacChannel`, `MacVlanChannel`.
* `l4switch` - In L4 switch topology, always clients and servers are on the same side. Enable l4switch to allow such connections. Possible values: `ENABLED`, `DISABLED`.
* `id` - The id of the l4param. It is a system-generated identifier.
