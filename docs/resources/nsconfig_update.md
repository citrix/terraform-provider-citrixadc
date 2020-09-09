---
subcategory: "NS"
---

# Resource: nsconfig_update

The nsconfig_update resource is used to apply the update operation for ns config.


## Example usage

```hcl
resource "citrixadc_nsconfig_update" "tf_nsupdate" {
    ipaddress = "10.0.1.164"
    netmask   = "255.255.255.0"
}  
```


## Argument Reference

* `ipaddress` - (Optional) IP address of the Citrix ADC. Commonly referred to as NSIP address. This parameter is mandatory to bring up the appliance.
* `netmask` - (Optional) Netmask corresponding to the IP address. This parameter is mandatory to bring up the appliance.
* `nsvlan` - (Optional) VLAN (NSVLAN) for the subnet on which the IP address resides.
* `ifnum` - (Optional) Interfaces of the appliances that must be bound to the NSVLAN.
* `tagged` - (Optional) Specifies that the interfaces will be added as 802.1q tagged interfaces. Packets sent on these interface on this VLAN will have an additional 4-byte 802.1q tag which identifies the VLAN. To use 802.1q tagging, the switch connected to the appliance's interfaces must also be configured for tagging. Possible values: [ YES, NO ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsconfig_update. It is a random string prefixed with "tf-nsconfig-update"
