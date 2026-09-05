---
subcategory: "Network"
---

# Data Source: ip6tunnel

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

### Read-only ip6tunnel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_ip6tunnel` resource). They are Computed / GET-only. Any attribute the appliance does not return is `null`.

* `remoteip` - The remote IP address or subnet of the tunnel.
* `type` - The type of this tunnel.
* `encapip` - The effective local IP address of the tunnel. Used as the source of the encapsulated packets.
