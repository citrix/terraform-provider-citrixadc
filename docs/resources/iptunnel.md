---
subcategory: "Network"
---

# Resource: iptunnel

The iptunnel resource is used to create ipv4 network tunnels.


## Example usage

```hcl
resource "citrixadc_iptunnel" "tf_iptunnel" {
    name = "tf_iptunnel"
    remote = "66.0.0.11"
    remotesubnetmask = "255.255.255.255"
    local = "*"
}
```


## Argument Reference

* `name` - (Required) Name for the IP tunnel. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).
* `remote` - (Optional) Public IPv4 address, of the remote device, used to set up the tunnel. For this parameter, you can alternatively specify a network address.
* `remotesubnetmask` - (Optional) Subnet mask of the remote IP address of the tunnel.
* `local` - (Optional) Type of Citrix ADC owned public IPv4 address, configured on the local Citrix ADC and used to set up the tunnel.
* `protocol` - (Optional) Name of the protocol to be used on this tunnel. Possible values: [ IPIP, GRE, IPSEC, UDP ]
* `grepayload` - (Optional) The payload GRE will carry. Possible values: [ ETHERNETwithDOT1Q, ETHERNET, IP ]
* `ipsecprofilename` - (Optional) Name of IPSec profile to be associated.
* `vlan` - (Optional) The vlan for mulicast packets.
* `ownergroup` - (Optional) The owner node group in a Cluster for the iptunnel.
* `destport` - (Optional) Specifies UDP destination port for Geneve packets. Default port is 6081.
* `tosinherit` - (Optional) Default behavior is to copy the ToS field of the internal IP Packet (Payload) to the outer IP packet (Transport packet). But the user can configure a new ToS field using this option.
* `vlantagging` - (Optional) Option to select Vlan Tagging.
* `vnid` - (Optional) Virtual network identifier (VNID) is the value that identifies a specific virtual network in the data plane.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the iptunnel. It has the same value as the `name` attribute.


## Import

A iptunnel can be imported using its name, e.g.

```shell
terraform import citrixadc_iptunnel.tf_iptunnel tf_iptunnel
```
