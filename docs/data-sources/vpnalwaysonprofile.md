---
subcategory: "VPN"
---

# Data Source: citrixadc_vpnalwaysonprofile

The vpnalwaysonprofile data source allows you to retrieve information about a VPN AlwaysOn profile configured on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_vpnalwaysonprofile" "tf_vpnalwaysonprofile" {
  name = "my_alwayson_profile"
}

output "clientcontrol" {
  value = data.citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile.clientcontrol
}

output "locationbasedvpn" {
  value = data.citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile.locationbasedvpn
}
```

## Argument Reference

* `name` - (Required) name of AlwaysON profile

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnalwaysonprofile. It has the same value as the `name` attribute.
* `clientcontrol` - Allow/Deny user to log off and connect to another Gateway
* `locationbasedvpn` - Option to decide if tunnel should be established when in enterprise network. When locationBasedVPN is remote, client tries to detect if it is located in enterprise network or not and establishes the tunnel if not in enterprise network. Dns suffixes configured using -add dns suffix- are used to decide if the client is in the enterprise network or not. If the resolution of the DNS suffix results in private IP, client is said to be in enterprise network. When set to EveryWhere, the client skips the check to detect if it is on the enterprise network and tries to establish the tunnel
* `networkaccessonvpnfailure` - Option to block network traffic when tunnel is not established(and the config requires that tunnel be established). When set to onlyToGateway, the network traffic to and from the client (except Gateway IP) is blocked. When set to fullAccess, the network traffic is not blocked
