---
subcategory: "Network"
---

# Resource: vlan\_nsip\_binding

The vlan\_nsip\_binding resource is used to bind vlan to ipv4 nsip address.


## Example usage

```hcl
resource "citrixadc_vlan" "tf_vlan" {
    vlanid = 40
    aliasname = "Management VLAN"
}

resource "citrixadc_nsip" "tf_snip" {
    ipaddress = "10.222.74.146"
    type = "SNIP"
    netmask = "255.255.255.0"
    icmp = "ENABLED"
    state = "ENABLED"
}

resource "citrixadc_vlan_nsip_binding" "tf_bind" {
    vlanid = citrixadc_vlan.tf_vlan.vlanid
    ipaddress = citrixadc_nsip.tf_snip.ipaddress
    netmask = citrixadc_nsip.tf_snip.netmask
}
```


## Argument Reference

* `ipaddress` - (Required) The IP address assigned to the VLAN.
* `netmask` - (Optional) Subnet mask for the network address defined for this VLAN.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `ownergroup` - (Optional) The owner node group in a Cluster for this vlan.
* `vlanid` - (Required) Specifies the virtual LAN ID.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vlan\_nsip\_binding. It is the concatenation of the `vlanid` and `ipaddress` attributes separated by a comma.


## Import

A vlan\_nsip\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vlan_nsip_binding.tf_bind 40,10.222.74.146
```
