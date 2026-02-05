---
subcategory: "NS"
---

# Data Source `nssimpleacl`

The nssimpleacl data source allows you to retrieve information about a simple ACL rule.


## Example usage

```terraform
data "citrixadc_nssimpleacl" "my_simpleacl" {
  aclname = "my_simpleacl"
}

output "aclaction" {
  value = data.citrixadc_nssimpleacl.my_simpleacl.aclaction
}

output "srcip" {
  value = data.citrixadc_nssimpleacl.my_simpleacl.srcip
}
```


## Argument Reference

* `aclname` - (Required) Name for the simple ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `aclaction` - Drop incoming IPv4 packets that match the simple ACL rule.
* `destport` - Port number to match against the destination port number of an incoming IPv4 packet.
* `estsessions` - Specifies whether the ACL should match only established TCP sessions.
* `protocol` - Protocol to match against the protocol of an incoming IPv4 packet.
* `srcip` - IP address to match against the source IP address of an incoming IPv4 packet.
* `td` - Traffic Domain ID.
* `ttl` - Number of seconds after which this simple ACL rule expires.
* `id` - The id of the nssimpleacl. It has the same value as the `aclname` attribute.
