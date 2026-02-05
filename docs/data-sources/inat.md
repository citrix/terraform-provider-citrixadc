---
subcategory: "Network"
---

# Data Source `inat`

The inat data source allows you to retrieve information about Inbound NAT (INAT) entries.


## Example usage

```terraform
data "citrixadc_inat" "tf_inat" {
  name = "my_inat"
}

output "privateip" {
  value = data.citrixadc_inat.tf_inat.privateip
}

output "publicip" {
  value = data.citrixadc_inat.tf_inat.publicip
}
```


## Argument Reference

* `name` - (Required) Name for the Inbound NAT (INAT) entry. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `connfailover` - Synchronize connection information with the secondary appliance in a high availability (HA) pair. That is, synchronize all connection-related information for the INAT session
* `ftp` - Enable the FTP protocol on the server for transferring files between the client and the server.
* `mode` - Stateless translation.
* `privateip` - IP address of the server to which the packet is sent by the Citrix ADC. Can be an IPv4 or IPv6 address.
* `proxyip` - Unique IP address used as the source IP address in packets sent to the server. Must be a MIP or SNIP address.
* `publicip` - Public IP address of packets received on the Citrix ADC. Can be aNetScaler-owned VIP or VIP6 address.
* `tcpproxy` - Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features.
* `td` - Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.
* `tftp` - To enable/disable TFTP (Default DISABLED).
* `useproxyport` - Enable the Citrix ADC to proxy the source port of packets before sending the packets to the server.
* `usip` - Enable the Citrix ADC to retain the source IP address of packets before sending the packets to the server.
* `usnip` - Enable the Citrix ADC to use a SNIP address as the source IP address of packets before sending the packets to the server.

## Attribute Reference

* `id` - The id of the inat. It has the same value as the `name` attribute.


## Import

An inat can be imported using its name, e.g.

```shell
terraform import citrixadc_inat.tf_inat my_inat
```
