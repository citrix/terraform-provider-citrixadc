---
subcategory: "Network"
---

# Resource: ip6tunnelparam

The ip6tunnelparam resource is used to create ip6 tunnel parameter resource.


## Example usage

```hcl
resource "citrixadc_ip6tunnelparam" "tf_ip6tunnelparam" {
  srcip                = "2001:db8:100::fb"
  dropfrag             = "NO"
  dropfragcputhreshold = 1
  srciproundrobin      = "NO"
  useclientsourceipv6  = "NO"
}
```


## Argument Reference

* `srcip` - (Optional) Common source IPv6 address for all IPv6 tunnels. Must be a SNIP6 or VIP6 address. Minimum length =  1
* `dropfrag` - (Optional) Drop any packet that requires fragmentation. Possible values: [ YES, NO ]
* `dropfragcputhreshold` - (Optional) Threshold value, as a percentage of CPU usage, at which to drop packets that require fragmentation. Applies only if dropFragparameter is set to NO. Minimum value =  1 Maximum value =  100
* `srciproundrobin` - (Optional) Use a different source IPv6 address for each new session through a particular IPv6 tunnel, as determined by round robin selection of one of the SNIP6 addresses. This setting is ignored if a common global source IPv6 address has been specified for all the IPv6 tunnels. This setting does not apply to a tunnel for which a source IPv6 address has been specified. Possible values: [ YES, NO ]
* `useclientsourceipv6` - (Optional) Use client source IPv6 address as source IPv6 address for outer tunnel IPv6 header. Possible values: [ YES, NO ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ip6tunnelparam. It is a unique string prefixed with "tf-ip6tunnelparam-"

