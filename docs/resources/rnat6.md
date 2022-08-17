---
subcategory: "Network"
---

# Resource: rnat6

The rnat6 resource is used to create rnat6.


## Example usage

```hcl
resource "citrixadc_rnat6" "tf_rnat6" {
  name             = "my_rnat6"
  network          = "2003::/64"
  srcippersistency = "ENABLED"
}
```


## Argument Reference

* `name` - (Required) Name for the RNAT6 rule. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. Choose a name that helps identify the RNAT6 rule. Minimum length =  1
* `network` - (Optional) IPv6 address of the network on whose traffic you want the Citrix ADC to do RNAT processing. Minimum length =  1
* `acl6name` - (Optional) Name of any configured ACL6 whose action is ALLOW. The rule of the ACL6 is used as an RNAT6 rule. Minimum length =  1
* `redirectport` - (Optional) Port number to which the IPv6 packets are redirected. Applicable to TCP and UDP protocols. Minimum value =  1 Maximum value =  65535
* `td` - (Optional) Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0. Minimum value =  0 Maximum value =  4094
* `srcippersistency` - (Optional) Enable source ip persistency, which enables the Citrix ADC to use the RNAT ips using source ip. Possible values: [ ENABLED, DISABLED ]
* `ownergroup` - (Optional) The owner node group in a Cluster for this rnat rule. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnat6. It has the same value as the `name` attribute.


## Import

A rnat6> can be imported using its name, e.g.

```shell
terraform import citrixadc_rnat6.tf_rnat6 my_rnat6
```
