---
subcategory: "Network"
---

# Resource: ip6tunnel

The ip6tunnel resource is used to create ip6 Tunnel resource.


## Example usage

```hcl
resource "citrixadc_nsip6" "test_nsip" {
  ipv6address = "23::30:20:23:34/64"
  type        = "VIP"
  icmp        = "DISABLED"
}
resource "citrixadc_ip6tunnel" "tf_ip6tunnel" {
  name   = "tf_ip6tunnel"
  remote = "2001:db8:0:b::/64"
  local  = trimsuffix(citrixadc_nsip6.test_nsip.ipv6address, "/64")
}
```


## Argument Reference

* `name` - (Required) Name for the IPv6 Tunnel. Cannot be changed after the service group is created. Must begin with a number or letter, and can consist of letters, numbers, and the @ _ - . (period) : (colon) # and space ( ) characters. Minimum length =  1 Maximum length =  31
* `remote` - (Required) An IPv6 address of the remote Citrix ADC used to set up the tunnel. Minimum length =  1
* `local` - (Required) An IPv6 address of the local Citrix ADC used to set up the tunnel.
* `ownergroup` - (Optional) The owner node group in a Cluster for the tunnel. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ip6tunnel. It has the same value as the `name` attribute.


## Import

A ip6tunnel can be imported using its name, e.g.

```shell
terraform import citrixadc_ip6tunnel.tf_ip6tunnel tf_ip6tunnel
```
