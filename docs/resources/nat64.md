---
subcategory: "Network"
---

# Resource: nat64

The nat64 resource is used to create nat64 config resource.


## Example usage

```hcl
resource "citrixadc_nsacl6" "tf_nsacl6" {
  acl6name   = "tf_nsacl6"
  acl6action = "ALLOW"
  logstate   = "ENABLED"
  stateful   = "NO"
  ratelimit  = 120
  state      = "ENABLED"
  priority   = 20
  protocol   = "TCP"
}
resource "citrixadc_netprofile" "tf_netprofile" {
  name                   = "tf_netprofile"
  proxyprotocol          = "ENABLED"
  proxyprotocoltxversion = "V1"
}
resource "citrixadc_nat64" "tf_nat64" {
  name       = "tf_nat64"
  acl6name   = citrixadc_nsacl6.tf_nsacl6.acl6name
  netprofile = citrixadc_netprofile.tf_netprofile.name
}
```


## Argument Reference

* `name` - (Required) Name for the NAT64 rule. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. Choose a name that helps identify the NAT64 rule. Minimum length =  1
* `acl6name` - (Required) Name of any configured ACL6 whose action is ALLOW.  IPv6 Packets matching the condition of this ACL6 rule and destination IP address of these packets matching the NAT64 IPv6 prefix are considered for NAT64 translation. Minimum length =  1
* `netprofile` - (Optional) Name of the configured netprofile. The Citrix ADC selects one of the IP address in the netprofile as the source IP address of the translated IPv4 packet to be sent to the IPv4 server. Minimum length =  1 Maximum length =  127


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nat64. It has the same value as the `name` attribute.


## Import

A nat64 can be imported using its name, e.g.

```shell
terraform import citrixadc_nat64.tf_nat64 tf_nat64
```
