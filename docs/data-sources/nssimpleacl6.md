---
subcategory: "NS"
---

# Data Source `nssimpleacl6`

The nssimpleacl6 data source allows you to retrieve information about a simple ACL6 rule for IPv6 traffic.


## Example usage

```terraform
data "citrixadc_nssimpleacl6" "my_simpleacl6" {
  aclname = "my_simpleacl6"
}

output "aclaction" {
  value = data.citrixadc_nssimpleacl6.my_simpleacl6.aclaction
}

output "srcipv6" {
  value = data.citrixadc_nssimpleacl6.my_simpleacl6.srcipv6
}
```


## Argument Reference

* `aclname` - (Required) Name for the simple ACL6 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `aclaction` - Drop incoming IPv6 packets that match the simple ACL6 rule.
* `destport` - Port number to match against the destination port number of an incoming IPv6 packet.
* `estsessions` - Specifies whether the ACL should match only established TCP sessions.
* `protocol` - Protocol to match against the protocol of an incoming IPv6 packet.
* `srcipv6` - IPv6 address to match against the source IPv6 address of an incoming IPv6 packet.
* `td` - Traffic Domain ID.
* `ttl` - Number of seconds after which this simple ACL6 rule expires.
* `id` - The id of the nssimpleacl6. It has the same value as the `aclname` attribute.
