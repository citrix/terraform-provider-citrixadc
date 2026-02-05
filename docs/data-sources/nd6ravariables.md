---
subcategory: "Network"
---

# Data Source `nd6ravariables`

The nd6ravariables data source allows you to retrieve information about IPv6 Neighbor Discovery Router Advertisement (ND6 RA) variables configured on a VLAN.


## Example usage

```terraform
data "citrixadc_nd6ravariables" "tf_nd6ravariables" {
  vlan = 1
}

output "ceaserouteradv" {
  value = data.citrixadc_nd6ravariables.tf_nd6ravariables.ceaserouteradv
}

output "onlyunicastrtadvresponse" {
  value = data.citrixadc_nd6ravariables.tf_nd6ravariables.onlyunicastrtadvresponse
}
```


## Argument Reference

* `vlan` - (Required) The VLAN number.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `ceaserouteradv` - Cease router advertisements on this vlan.
* `currhoplimit` - Current Hop limit.
* `defaultlifetime` - Default life time, in seconds.
* `linkmtu` - The Link MTU.
* `managedaddrconfig` - Value to be placed in the Managed address configuration flag field.
* `maxrtadvinterval` - Maximum time allowed between unsolicited multicast RAs, in seconds.
* `minrtadvinterval` - Minimum time interval between RA messages, in seconds.
* `onlyunicastrtadvresponse` - Send only Unicast Router Advertisements in respond to Router Solicitations.
* `otheraddrconfig` - Value to be placed in the Other configuration flag field.
* `reachabletime` - Reachable time, in milliseconds.
* `retranstime` - Retransmission time, in milliseconds.
* `sendrouteradv` - whether the router sends periodic RAs and responds to Router Solicitations.
* `srclinklayeraddroption` - Include source link layer address option in RA messages.

## Attribute Reference

* `id` - The id of the nd6ravariables. It has the same value as the `vlan` attribute.
