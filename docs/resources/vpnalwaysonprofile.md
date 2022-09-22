---
subcategory: "VPN"
---

# Resource: vpnalwaysonprofile

The vpnalwaysonprofile resource is used to ensure that users are always connected to the enterprise network.


## Example usage

```hcl
resource "citrixadc_vpnalwaysonprofile" "tf_vpnalwaysonprofile" {
	name = "tf_vpnalwaysonprofile"
	clientcontrol = "ALLOW"
	locationbasedvpn = "Everywhere"
	networkaccessonvpnfailure = "fullAccess"
}
```


## Argument Reference

* `name` - (Required) name of AlwaysON profile.
* `networkaccessonvpnfailure` - (Optional) Option to block network traffic when tunnel is not established(and the config requires that tunnel be established). When set to onlyToGateway, the network traffic to and from the client (except Gateway IP) is blocked. When set to fullAccess, the network traffic is not blocked. Possible values: [ onlyToGateway, fullAccess ]
* `clientcontrol` - (Optional) Allow/Deny user to log off and connect to another Gateway. Possible values: [ ALLOW, DENY ]
* `locationbasedvpn` - (Optional) Option to decide if tunnel should be established when in enterprise network. When locationBasedVPN is remote, client tries to detect if it is located in enterprise network or not and establishes the tunnel if not in enterprise network. Dns suffixes configured using -add dns suffix- are used to decide if the client is in the enterprise network or not. If the resolution of the DNS suffix results in private IP, client is said to be in enterprise network. When set to EveryWhere, the client skips the check to detect if it is on the enterprise network and tries to establish the tunnel. Possible values: [ Remote, Everywhere ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnalwaysonprofile. It has the same value as the `name` attribute.


## Import

A vpnalwaysonprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnalwaysonprofile.tf_vpnalwaysonprofile tf_vpnalwaysonprofile
```
