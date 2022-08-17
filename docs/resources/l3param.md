---
subcategory: "Network"
---

# Resource: l3param

The l3param resource is used to update l3param.


## Example usage

```hcl
resource "citrixadc_l3param" "tf_l3param" {
  srcnat               = "ENABLED"
  icmpgenratethreshold = 150
  overridernat         = "ENABLED"
  dropdfflag           = "ENABLED"
}
```


## Argument Reference

* `srcnat` - (Optional) Perform NAT if only the source is in the private network. Possible values: [ ENABLED, DISABLED ]
* `icmpgenratethreshold` - (Optional) NS generated ICMP pkts per 10ms rate threshold.
* `overridernat` - (Optional) USNIP/USIP settings override RNAT settings for configured service/virtual server traffic.. . Possible values: [ ENABLED, DISABLED ]
* `dropdfflag` - (Optional) Enable dropping the IP DF flag. Possible values: [ ENABLED, DISABLED ]
* `miproundrobin` - (Optional) Enable round robin usage of mapped IPs. Possible values: [ ENABLED, DISABLED ]
* `externalloopback` - (Optional) Enable external loopback. Possible values: [ ENABLED, DISABLED ]
* `tnlpmtuwoconn` - (Optional) Enable/Disable learning PMTU of IP tunnel when ICMP error does not contain connection information. Possible values: [ ENABLED, DISABLED ]
* `usipserverstraypkt` - (Optional) Enable detection of stray server side pkts in USIP mode. Possible values: [ ENABLED, DISABLED ]
* `forwardicmpfragments` - (Optional) Enable forwarding of ICMP fragments. Possible values: [ ENABLED, DISABLED ]
* `dropipfragments` - (Optional) Enable dropping of IP fragments. Possible values: [ ENABLED, DISABLED ]
* `acllogtime` - (Optional) Parameter to tune acl logging time. Possible values: [ ENABLED, DISABLED ]
* `implicitaclallow` - (Optional) Do not apply ACLs for internal ports. Possible values: [ ENABLED, DISABLED ]
* `dynamicrouting` - (Optional) Enable/Disable Dynamic routing on partition. This configuration is not applicable to default partition. Possible values: [ ENABLED, DISABLED ]
* `ipv6dynamicrouting` - (Optional) Enable/Disable IPv6 Dynamic routing. Possible values: [ ENABLED, DISABLED ]
* `allowclasseipv4` - (Optional) Enable/Disable IPv4 Class E address clients. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the l3param. It is a unique string prefixed with "tf-l3param-"