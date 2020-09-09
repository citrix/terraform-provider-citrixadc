---
subcategory: "Network"
---

# Resource: inat

The inat resource is used to create inbound nat resources.


## Example usage

```hcl
resource "citrixadc_inat" "inat_resource" {
  name = "ip4ip"
  privateip = "192.168.1.1"
  publicip = "172.16.1.2"
  tcpproxy = "ENABLED"
  usnip = "ON"
}
```


## Argument Reference

* `name` - (Required) Name for the Inbound NAT (INAT) entry.
* `publicip` - (Optional) Public IP address of packets received on the Citrix ADC. Can be aNetScaler-owned VIP or VIP6 address.
* `privateip` - (Optional) IP address of the server to which the packet is sent by the Citrix ADC. Can be an IPv4 or IPv6 address.
* `mode` - (Optional) Stateless translation. Possible values: [ STATELESS ]
* `tcpproxy` - (Optional) Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features. Possible values: [ ENABLED, DISABLED ]
* `ftp` - (Optional) Enable the FTP protocol on the server for transferring files between the client and the server. Possible values: [ ENABLED, DISABLED ]
* `tftp` - (Optional) To enable/disable TFTP (Default DISABLED). Possible values: [ ENABLED, DISABLED ]
* `usip` - (Optional) Enable the Citrix ADC to retain the source IP address of packets before sending the packets to the server. Possible values: [ on, off ]
* `usnip` - (Optional) Enable the Citrix ADC to use a SNIP address as the source IP address of packets before sending the packets to the server. Possible values: [ on, off ]
* `proxyip` - (Optional) Unique IP address used as the source IP address in packets sent to the server. Must be a MIP or SNIP address.
* `useproxyport` - (Optional) Enable the Citrix ADC to proxy the source port of packets before sending the packets to the server. Possible values: [ ENABLED, DISABLED ]
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the inat resource. It has the same value as the `name` attribute.


## Import

A <resource> can be imported using its name, e.g.

```shell
terraform import citrixadc_inat.inat_resource ip4ip
```
