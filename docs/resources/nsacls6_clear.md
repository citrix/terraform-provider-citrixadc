---
subcategory: "NS"
---

# Resource: nsacls6_clear

Removes all configured IPv6 extended ACL (ACL6) rules from the Citrix ADC in one operation. Use this resource to reset the ACL6 rule set — for example, before reloading a fresh set of rules or when tearing down an ACL6 configuration — instead of deleting each `ns acl6` entry individually.

This is an **action-only** resource. Applying it triggers the NITRO `clear` action on the `nsacls6` endpoint (the same endpoint also backs the `apply` and `renumber` actions). There is no corresponding ADC object to read back, so this resource performs a one-shot action on create. Read, update, and delete are no-ops, and importing it is not meaningful. There is no NITRO GET endpoint for this action, so there is no corresponding data source.


## Example usage

### Clear the CLASSIC ACL6 rules

With no arguments, the `CLASSIC` ACL6 rule set is cleared (the default).

```hcl
resource "citrixadc_nsacls6_clear" "clear_acl6" {}
```

### Clear a specific ACL6 type

```hcl
resource "citrixadc_nsacls6_clear" "clear_dfd" {
  type = "DFD"
}
```


## Argument Reference

* `type` - (Optional) Type of the ACL6 set to clear. Defaults to `CLASSIC`. Possible values: [ CLASSIC, DFD ]. `CLASSIC` clears the regular extended ACL6 rules; `DFD` clears cluster-specific ACL6 rules that specify the hash method used for steering packets within a cluster. Changing this value re-triggers the clear action (replacement).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `nsacls6_clear`. Because `nsacls6_clear` performs a NITRO action with no GET endpoint, the ID does not correspond to any readable object on the Citrix ADC.


## Note

This resource models a one-shot action rather than a persistent ADC object. Applying it clears the current ACL6 rule set; the result cannot be read back from the ADC, so subsequent plans will not detect drift. To re-run the clear action, taint the resource or change the `type` argument. Destroying the resource only removes it from Terraform state and does not restore the cleared ACL6 rules. Importing this resource is not meaningful.
