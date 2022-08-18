---
subcategory: "Network"
---

# Resource: vlan_nsip6_binding

The vlan_nsip6_binding resource is used to create vlan_nsip6_binding.


## Example usage

```hcl
resource "citrixadc_vlan_nsip6_binding" "tf_vlan_nsip6_binding" {
  vlanid    = 2
  ipaddress = "2001::a/96"
}
```


## Argument Reference

* `vlanid` - (Required) Specifies the virtual LAN ID. Minimum value =  1 Maximum value =  4094
* `ipaddress` - (Required) The IP address assigned to the VLAN.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `ownergroup` - (Optional) The owner node group in a Cluster for this vlan. Minimum length =  1
* `netmask` - (Optional) Subnet mask for the network address defined for this VLAN. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vlan_nsip6_binding. It has the same value as the `vlanid` and `ipaddress` attributes separated by a comma.


## Import

A vlan_nsip6_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_vlan_nsip6_binding.tf_vlan_nsip6_binding 2,2001::a/96
```
