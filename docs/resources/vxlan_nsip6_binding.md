---
subcategory: "Network"
---

# Resource: vxlan_nsip6_binding

The vxlan_nsip6_binding resource is used to bind nsip6 to vxlan resource.


## Example usage

```hcl
resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_nsip6" "test_nsip" {
  ipv6address = "2001:db8:100::fb/64"
  type        = "VIP"
  icmp        = "DISABLED"
}
resource "citrixadc_vxlan_nsip6_binding" "tf_binding" {
  vxlanid   = citrixadc_vxlan.tf_vxlan.vxlanid
  ipaddress = citrixadc_nsip6.test_nsip.ipv6address
  netmask   = "255.255.255.0"
}
```


## Argument Reference

* `vxlanid` - (Required) A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.
* `ipaddress` - (Required) The IP address assigned to the VXLAN.
* `netmask` - (Optional) Subnet mask for the network address defined for this VXLAN.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vxlan_nsip6_binding. It is the concatenation of `vxlanid` and `ipaddress` attributes seperated by comma.


## Import

A vxlan_nsip6_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vxlan_nsip6_binding.tf_binding 123,2001:db8:100::fb/64
```
