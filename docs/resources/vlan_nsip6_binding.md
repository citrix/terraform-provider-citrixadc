---
subcategory: "Network"
---

# Resource: vlan_nsip6_binding

The vlan_nsip6_binding resource is used to bind an IPv6 address to a VLAN.


## Example usage

```hcl
resource "citrixadc_vlan_nsip6_binding" "tf_vlan_nsip6_binding" {
  vlanid    = 2
  ipaddress = "2001::a/96"
}
```


## Argument Reference

* `vlanid` - (Required) Specifies the virtual LAN ID.
* `ipaddress` - (Required) The IP address assigned to the VLAN.
* `netmask` - (Optional) Subnet mask for the network address defined for this VLAN.
* `ownergroup` - (Optional) The owner node group in a Cluster for this VLAN.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vlan_nsip6_binding. It is the concatenation of the `vlanid` and `ipaddress` attributes separated by a comma.


## Import

A vlan_nsip6_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding 2,2001::a/96
```
