---
subcategory: "Network"
---

# Resource: vxlan_nsip_binding

The vxlan_nsip_binding resource is used to bind nsip to vxlan resource.


## Example usage

```hcl
resource "citrixadc_vxlan" "tf_vxlan" {
  vxlanid            = 123
  port               = 33
  dynamicrouting     = "DISABLED"
  ipv6dynamicrouting = "DISABLED"
  innervlantagging   = "ENABLED"
}
resource "citrixadc_nsip" "tf_snip" {
  ipaddress = "10.222.74.146"
  type      = "SNIP"
  netmask   = "255.255.255.0"
  icmp      = "ENABLED"
  state     = "ENABLED"
}
resource "citrixadc_vxlan_nsip_binding" "tf_binding" {
  vxlanid   = citrixadc_vxlan.tf_vxlan.vxlanid
  ipaddress = citrixadc_nsip.tf_snip.ipaddress
  netmask   = citrixadc_nsip.tf_snip.netmask
}
```


## Argument Reference

* `vxlanid` - (Required) A positive integer, which is also called VXLAN Network Identifier (VNI), that uniquely identifies a VXLAN.
* `ipaddress` - (Required) The IP address assigned to the VXLAN.
* `netmask` - (Optional) Subnet mask for the network address defined for this VXLAN.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vxlan_nsip_binding. It has the same value as the `vxlanid` and `ipaddress` attributes seperated by comma.


## Import

A vxlan_nsip_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_vxlan_nsip_binding.tf_binding 123,10.222.74.146
```
