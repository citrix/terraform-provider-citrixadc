---
subcategory: "Utility"
---

# Resource: ping6

This resource is used to run the ping6 diagnostic action on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_ping6" "tf_ping6" {
  hostname = "::1"
  c        = 1
}
```


## Argument Reference

* `hostname` - (Required) Address of host to ping (NITRO parameter: `hostName`).
* `b` - (Optional) Set socket buffer size. If used, should be used with roughly +100 then the datalen (`-s` option). The default value is 8192.
* `c` - (Optional) Number of packets to send. The default value is infinite. For Nitro API, defalut value is taken as 5.
* `i` - (Optional) Waiting time, in seconds. The default value is 1 second (NITRO parameter: `i`).
* `i_upper` - (Optional) Network interface on which to ping, if you have multiple interfaces (NITRO parameter: `I`).
* `m` - (Optional) By default, ping6 asks the kernel to fragment packets to fit into the minimum IPv6 MTU. This option suppresses that behavior for unicast packets.
* `n` - (Optional) Numeric output only. No name resolution.
* `p` - (Optional) Pattern to fill in packets. Can be up to 16 bytes, useful for diagnosing data-dependent problems.
* `q` - (Optional) Quiet output. Only summary is printed. For Nitro API, this flag is set by default.
* `s` - (Optional) Data size, in bytes. The default value is 32 (NITRO parameter: `s`).
* `v` - (Optional) VLAN ID for link local address (NITRO parameter: `V`).
* `s_upper` - (Optional) Source IP address to be used in the outgoing query packets (NITRO parameter: `S`).
* `t_upper` - (Optional) Traffic Domain Id (NITRO parameter: `T`).
* `t` - (Optional) Timeout in seconds before ping6 exits (NITRO parameter: `t`).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ping6 resource. It is set to `ping6-config`.
