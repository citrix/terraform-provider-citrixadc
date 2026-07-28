---
subcategory: "Utility"
---

# Resource: traceroute6

This resource is used to run an IPv6 traceroute from the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_traceroute6" "tf_traceroute6" {
  host = "::1"
  m    = 3
}
```


## Argument Reference

* `host` - (Required) Destination host IP address or name.
* `n` - (Optional) Print hop addresses numerically rather than symbolically and numerically.
* `i` - (Optional) Use ICMP ECHO for probes (NITRO parameter: `I`).
* `r` - (Optional) Bypass normal routing tables and send directly to a host on an attached network. If the host is not on a directly attached network, an error is returned.
* `v` - (Optional) Verbose output. List received ICMP packets other than TIME_EXCEEDED and UNREACHABLE.
* `m` - (Optional) Maximum hop value for outgoing probe packets. For Nitro API, default value is taken as 10.
* `p` - (Optional) Base port number used in probes.
* `q` - (Optional) Number of probes per hop. For Nitro API, default value is taken as 1.
* `s` - (Optional) Source IP address to use in the outgoing query packets. If the IP address does not belong to this appliance, an error is returned and nothing is sent.
* `t` - (Optional) Traffic Domain Id (NITRO parameter: `T`).
* `w` - (Optional) Time (in seconds) to wait for a response to a query. For Nitro API, defalut value is set to 3.
* `packetlen` - (Optional) Length (in bytes) of the query packets.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the traceroute6 resource. It is set to `traceroute6-config`.
