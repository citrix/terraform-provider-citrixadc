---
subcategory: "NS"
---

# Resource: nssimpleacl

The nssimpleacl resource is used to create simple ACL resource.


## Example usage

```hcl
resource "citrixadc_nssimpleacl" "tf_nssimpleacl" {
  aclname   = "tf_nssimpleacl"
  aclaction = "DENY"
  srcip     = "1.2.3.1"
  destport  = 123
  protocol  = "UDP"
  ttl       = 600
}
```


## Argument Reference

* `aclname` - (Required) Name for the simple ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the simple ACL rule is created. Minimum length =  1
* `aclaction` - (Required) Drop incoming IPv4 packets that match the simple ACL rule. Possible values: [ DENY ]4094
* `srcip` - (Required) IP address to match against the source IP address of an incoming IPv4 packet.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  
* `destport` - (Optional) Port number to match against the destination port number of an incoming IPv4 packet. DestPort is mandatory while setting Protocol. Omitting the port number and protocol creates an all-ports  and all protocols simple ACL rule, which matches any port and any protocol. In that case, you cannot create another simple ACL rule specifying a specific port and the same source IPv4 address. Minimum value =  1 Maximum value =  65535
* `protocol` - (Optional) Protocol to match against the protocol of an incoming IPv4 packet. You must set this parameter if you have set the Destination Port parameter. Possible values: [ TCP, UDP ]
* `ttl` - (Optional) Number of seconds, in multiples of four, after which the simple ACL rule expires. If you do not want the simple ACL rule to expire, do not specify a TTL value. Minimum value =  4 Maximum value =  2147483647
* `estsessions` - (Optional) .


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nssimpleacl. It has the same value as the `aclname` attribute.


## Import

A nssimpleacl can be imported using its aclname, e.g.

```shell
terraform import citrixadc_nssimpleacl.tf_nssimpleacl tf_nssimpleacl
```
