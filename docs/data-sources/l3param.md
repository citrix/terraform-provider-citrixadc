---
subcategory: "Network"
---

# Data Source `l3param`

The l3param data source allows you to retrieve information about Layer 3 parameters configuration.


## Example usage

```terraform
data "citrixadc_l3param" "tf_l3param" {
}

output "srcnat" {
  value = data.citrixadc_l3param.tf_l3param.srcnat
}

output "icmpgenratethreshold" {
  value = data.citrixadc_l3param.tf_l3param.icmpgenratethreshold
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `acllogtime` - Parameter to tune acl logging time.
* `allowclasseipv4` - Enable/Disable IPv4 Class E address clients. Possible values: `ENABLED`, `DISABLED`.
* `dropdfflag` - Enable dropping the IP DF flag. Possible values: `ENABLED`, `DISABLED`.
* `dropipfragments` - Enable dropping of IP fragments. Possible values: `ENABLED`, `DISABLED`.
* `dynamicrouting` - Enable/Disable Dynamic routing on partition. This configuration is not applicable to default partition. Possible values: `ENABLED`, `DISABLED`.
* `externalloopback` - Enable external loopback. Possible values: `ENABLED`, `DISABLED`.
* `forwardicmpfragments` - Enable forwarding of ICMP fragments. Possible values: `YES`, `NO`.
* `icmpgenratethreshold` - NS generated ICMP pkts per 10ms rate threshold.
* `implicitaclallow` - Do not apply ACLs for internal ports. Possible values: `ENABLED`, `DISABLED`.
* `implicitpbr` - Enable/Disable Policy Based Routing for control packets. Possible values: `ENABLED`, `DISABLED`.
* `ipv6dynamicrouting` - Enable/Disable IPv6 Dynamic routing. Possible values: `ENABLED`, `DISABLED`.
* `miproundrobin` - Enable round robin usage of mapped IPs. Possible values: `ENABLED`, `DISABLED`.
* `overridernat` - USNIP/USIP settings override RNAT settings for configured service/virtual server traffic. Possible values: `ENABLED`, `DISABLED`.
* `srcnat` - Perform NAT if only the source is in the private network. Possible values: `ENABLED`, `DISABLED`.
* `tnlpmtuwoconn` - Enable/Disable learning PMTU of IP tunnel when ICMP error does not contain connection information. Possible values: `ENABLED`, `DISABLED`.
* `usipserverstraypkt` - Enable detection of stray server side pkts in USIP mode. Possible values: `ENABLED`, `DISABLED`.
* `id` - The id of the l3param. It is a system-generated identifier.
