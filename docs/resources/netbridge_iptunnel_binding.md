---
subcategory: "Network"
---

# Resource: netbridge_iptunnel_binding

The netbridge_iptunnel_binding resource is used to bind iptunnel to the netbridge resource.


## Example usage

```hcl
resource "citrixadc_vxlanvlanmap" "tf_vxlanvlanmp" {
  name = "tf_vxlanvlanmp"
}
resource "citrixadc_nsip" "nsip" {
  ipaddress = "2.2.2.1"
  type      = "VIP"
  netmask   = "255.255.255.0"
}
resource "citrixadc_netbridge" "tf_netbridge" {
  name         = "tf_netbridge"
  vxlanvlanmap = citrixadc_vxlanvlanmap.tf_vxlanvlanmp.name
}
resource "citrixadc_iptunnel" "tf_iptunnel" {
  name             = "tf_iptunnel"
  remote           = "66.0.0.11"
  remotesubnetmask = "255.255.255.255"
  local            = citrixadc_nsip.nsip.ipaddress
  protocol         = "GRE"
}
resource "citrixadc_netbridge_iptunnel_binding" "tf_binding" {
  name   = citrixadc_netbridge.tf_netbridge.name
  tunnel = citrixadc_iptunnel.tf_iptunnel.name
}
```


## Argument Reference

* `tunnel` - (Required) The name of the tunnel that is a part of this bridge.
* `name` - (Required) The name of the network bridge.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the netbridge_iptunnel_binding. It is the concatenation of `name` and `tunnel` attributes separated by comma.


## Import

A netbridge_iptunnel_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_netbridge_iptunnel_binding.tf_binding tf_netbridge,tf_iptunnel
```
