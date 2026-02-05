---
subcategory: "Network"
---

# Data Source `ip6tunnelparam`

The ip6tunnelparam data source allows you to retrieve information about IPv6 tunnel parameters configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_ip6tunnelparam" "tf_ip6tunnelparam" {
}

output "ip6tunnelparam_srcip" {
  value = data.citrixadc_ip6tunnelparam.tf_ip6tunnelparam.srcip
}
```


## Argument Reference

This data source does not require any arguments as it retrieves the global IPv6 tunnel parameters.

## Attribute Reference

The following attributes are exported:

* `id` - The id of the ip6tunnelparam.
* `srcip` - Common source IPv6 address for all IPv6 tunnels. Must be a SNIP6 or VIP6 address.
* `dropfrag` - Drop any packet that requires fragmentation.
* `dropfragcputhreshold` - Threshold value, as a percentage of CPU usage, at which to drop packets that require fragmentation. Applies only if dropFragparameter is set to NO.
* `srciproundrobin` - Use a different source IPv6 address for each new session through a particular IPv6 tunnel, as determined by round robin selection of one of the SNIP6 addresses. This setting is ignored if a common global source IPv6 address has been specified for all the IPv6 tunnels. This setting does not apply to a tunnel for which a source IPv6 address has been specified.
* `useclientsourceipv6` - Use client source IPv6 address as source IPv6 address for outer tunnel IPv6 header
