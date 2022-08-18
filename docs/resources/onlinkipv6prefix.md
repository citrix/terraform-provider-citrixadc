---
subcategory: "Network"
---

# Resource: onlinkipv6prefix

The onlinkipv6prefix resource is used to create onlinkipv6prefix.


## Example usage

```hcl
resource "citrixadc_onlinkipv6prefix" "tf_onlinkipv6prefix" {
  ipv6prefix      = "8000::/64"
  onlinkprefix    = "YES"
  autonomusprefix = "NO"
}
```


## Argument Reference

* `ipv6prefix` - (Required) Onlink prefixes for RA messages.
* `onlinkprefix` - (Optional) RA Prefix onlink flag. Possible values: [ YES, NO ]
* `autonomusprefix` - (Optional) RA Prefix Autonomus flag. Possible values: [ YES, NO ]
* `depricateprefix` - (Optional) Depricate the prefix. Possible values: [ YES, NO ]
* `decrementprefixlifetimes` - (Optional) RA Prefix Autonomus flag. Possible values: [ YES, NO ]
* `prefixvalidelifetime` - (Optional) Valide life time of the prefix, in seconds.
* `prefixpreferredlifetime` - (Optional) Preferred life time of the prefix, in seconds.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the onlinkipv6prefix. It has the same value as the `ipv6prefix` attribute.


## Import

A onlinkipv6prefix can be imported using its name, e.g.

```shell
terraform import citrixadc_onlinkipv6prefix.tf_onlinkipv6prefix 8000::/64
```
