---
subcategory: "Network"
---

# Data Source `ip6tunnel`

The ip6tunnel data source allows you to retrieve information about IPv6 tunnels configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_ip6tunnel" "tf_ip6tunnel" {
  name   = "my_ip6tunnel"
  remote = "2001:db8:0:b::/64"
}

output "local_address" {
  value = data.citrixadc_ip6tunnel.tf_ip6tunnel.local
}

output "remote_address" {
  value = data.citrixadc_ip6tunnel.tf_ip6tunnel.remote
}
```


## Argument Reference

* `name` - (Required) Name for the IPv6 Tunnel. Cannot be changed after the service group is created. Must begin with a number or letter, and can consist of letters, numbers, and the @ _ - . (period) : (colon) # and space ( ) characters.
* `remote` - (Required) An IPv6 address of the remote Citrix ADC used to set up the tunnel.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ip6tunnel resource.
* `local` - An IPv6 address of the local Citrix ADC used to set up the tunnel.
* `ownergroup` - The owner node group in a Cluster for the tunnel.
