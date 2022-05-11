---
subcategory: "Network"
---

# Resource: forwardingsession

The forwardingsession resource is used to create session forward resource.


## Example usage

```hcl
resource "citrixadc_forwardingsession" "tf_forwarding" {
  name             = "tf_forwarding"
  network          = "10.102.105.90"
  netmask          = "255.255.255.255"
  connfailover     = "ENABLED"
  sourceroutecache = "ENABLED"
  processlocal     = "DISABLED"
}
```


## Argument Reference

* `name` - (required) Name for the forwarding session rule. Can begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my rule" or 'my rule').
* `acl6name` - (Optional) Name of any configured ACL6 whose action is ALLOW. The rule of the ACL6 is used as a forwarding session rule.
* `aclname` - (Optional) Name of any configured ACL whose action is ALLOW. The rule of the ACL is used as a forwarding session rule.
* `connfailover` - (Optional) Synchronize connection information with the secondary appliance in a high availability (HA) pair. That is, synchronize all connection-related information for the forwarding session.
* `netmask` - (Optional) Subnet mask associated with the network.
* `network` - (Optional) An IPv4 network address or IPv6 prefix of a network from which the forwarded traffic originates or to which it is destined.
* `processlocal` - (Optional) Enabling this option on forwarding session will not steer the packet to flow processor. Instead, packet will be routed.
* `sourceroutecache` - (Optional) Cache the source ip address and mac address of the DA servers.
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the forwardingsession. It has the same value as the `name` attribute.


## Import

A forwardingsession can be imported using its name, e.g.

```shell
terraform import citrixadc_forwardingsession.tf_forwarding tf_forwarding
```
