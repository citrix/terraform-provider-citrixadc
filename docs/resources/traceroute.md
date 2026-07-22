---
subcategory: "Utility"
---

# Resource: traceroute

The traceroute resource performs the NITRO `traceroute` action, which runs a traceroute from the target Citrix ADC. It is an action-only diagnostic resource: applying it runs the traceroute once. Every argument is one-shot and forces replacement when changed.

-> **Attribute naming** Several NITRO parameters differ only by letter case (e.g. `s`/`S`, `m`/`M`, `p`/`P`, `t`/`T`). Because Terraform attribute names must be lowercase and unique, the upper-case NITRO parameter is exposed with an `_upper` suffix (for example NITRO `S` -> `s_upper`, `M` -> `m_upper`, `P` -> `p_upper`, `T` -> `t_upper`).


## Example usage

```hcl
resource "citrixadc_traceroute" "tf_traceroute" {
  host = "127.0.0.1"
  m    = 3
}
```


## Argument Reference

* `host` - (Required) Destination host IP address or name.
* `s_upper` - (Optional) Print a summary of how many probes were not answered for each hop (NITRO parameter: `S`).
* `n` - (Optional) Print hop addresses numerically instead of symbolically and numerically.
* `r` - (Optional) Bypass normal routing tables and send directly to a host on an attached network. If the host is not on a directly attached network, an error is returned.
* `v` - (Optional) Verbose output. List received ICMP packets other than TIME_EXCEEDED and UNREACHABLE.
* `m_upper` - (Optional) Minimum TTL value used in outgoing probe packets (NITRO parameter: `M`).
* `m` - (Optional) Maximum TTL value used in outgoing probe packets. For Nitro API, default value is taken as 10 (NITRO parameter: `m`).
* `p_upper` - (Optional) Send packets of specified IP protocol. The currently supported protocols are UDP and ICMP (NITRO parameter: `P`).
* `p` - (Optional) Base port number used in probes (NITRO parameter: `p`).
* `q` - (Optional) Number of queries per hop. For Nitro API, defalut value is taken as 1.
* `s` - (Optional) Source IP address to use in the outgoing query packets. If the IP address does not belong to this appliance, an error is returned and nothing is sent (NITRO parameter: `s`).
* `t_upper` - (Optional) Traffic Domain Id (NITRO parameter: `T`).
* `t` - (Optional) Type-of-service in query packets (NITRO parameter: `t`).
* `w` - (Optional) Time (in seconds) to wait for a response to a query. For Nitro API, defalut value is set to 3.
* `packetlen` - (Optional) Length (in bytes) of the query packets.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the traceroute resource. It is set to `traceroute-config`.
