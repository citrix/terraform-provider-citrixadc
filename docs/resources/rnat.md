---
subcategory: "Network"
---

# Resource: rnat

The rnat resource is used to create rnat.


## Example usage

```hcl
resource "citrixadc_rnat" "tfrnat" {
  name             = "tfrnat"
  network          = "10.2.2.0"
  netmask          = "255.255.255.255"
  useproxyport     = "ENABLED"
  srcippersistency = "DISABLED"
  connfailover     = "DISABLED"
}
```


## Argument Reference
1
* `name` - (Required) Name for the RNAT4 rule. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. Choose a name that helps identify the RNAT4 rule. Minimum length =  1
* `network` - (Optional) The network address defined for the RNAT entry. Minimum length =  1
* `netmask` - (Optional) The subnet mask for the network address. Minimum length =  1
* `aclname` - (Optional) An extended ACL defined for the RNAT entry. Minimum length =  1
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `ownergroup` - (Optional) The owner node group in a Cluster for this rnat rule. Minimum length =  
* `redirectport` - (Optional) Port number to which the IPv4 packets are redirected. Applicable to TCP and UDP protocols. Minimum value =  1 Maximum value =  65535
* `natip` - (Optional) Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range. Minimum length =  1
* `srcippersistency` - (Optional) Enables the Citrix ADC to use the same NAT IP address for all RNAT sessions initiated from a particular server. Possible values: [ ENABLED, DISABLED ]
* `useproxyport` - (Optional) Enable source port proxying, which enables the Citrix ADC to use the RNAT ips using proxied source port. Possible values: [ ENABLED, DISABLED ]
* `connfailover` - (Optional) Synchronize all connection-related information for the RNAT sessions with the secondary ADC in a high availability (HA) pair. Possible values: [ ENABLED, DISABLED ]
* `newname` - (Optional) New name for the RNAT4 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain       only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnat. It has the same value as the `name` attribute.


## Import

A rnat can be imported using its name, e.g.

```shell
terraform import citrixadc_rnat.tfrnat tfrnat
```
