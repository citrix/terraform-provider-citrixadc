---
subcategory: "Network"
---

# Resource: nd6ravariables

The nd6ravariables resource is used to update nd6ravariables.


## Example usage

```hcl
resource "citrixadc_nd6ravariables" "tf_nd6ravariables" {
  vlan                     = 1
  ceaserouteradv           = "NO"
  onlyunicastrtadvresponse = "YES"
  srclinklayeraddroption   = "NO"
}
```


## Argument Reference

* `vlan` - (Required) The VLAN number. Minimum value =  1 Maximum value =  4094
* `ceaserouteradv` - (Optional) Cease router advertisements on this vlan. Possible values: [ YES, NO ]
* `sendrouteradv` - (Optional) whether the router sends periodic RAs and responds to Router Solicitations. Possible values: [ YES, NO ]
* `srclinklayeraddroption` - (Optional) Include source link layer address option in RA messages. Possible values: [ YES, NO ]
* `onlyunicastrtadvresponse` - (Optional) Send only Unicast Router Advertisements in respond to Router Solicitations. Possible values: [ YES, NO ]
* `managedaddrconfig` - (Optional) Value to be placed in the Managed address configuration flag field. Possible values: [ YES, NO ]
* `otheraddrconfig` - (Optional) Value to be placed in the Other configuration flag field. Possible values: [ YES, NO ]
* `currhoplimit` - (Optional) Current Hop limit. Minimum value =  0 Maximum value =  255
* `maxrtadvinterval` - (Optional) Maximum time allowed between unsolicited multicast RAs, in seconds. Minimum value =  4 Maximum value =  1800
* `minrtadvinterval` - (Optional) Minimum time interval between RA messages, in seconds. Minimum value =  3 Maximum value =  1350
* `linkmtu` - (Optional) The Link MTU. Minimum value =  0 Maximum value =  1500
* `reachabletime` - (Optional) Reachable time, in milliseconds. Minimum value =  0 Maximum value =  3600000
* `retranstime` - (Optional) Retransmission time, in milliseconds.
* `defaultlifetime` - (Optional) Default life time, in seconds. Minimum value =  0 Maximum value =  9000


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nd6ravariables. It has the same value as the `vlan` attribute.


## Import

A nd6ravariables can be imported using its name, e.g.

```shell
terraform import citrixadc_nd6ravariables.tf_nd6ravariables 1
```
