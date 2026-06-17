---
subcategory: "NS"
---

# Resource: nsacls6

Commits the staged IPv6 extended ACL (ACL6) rule set so that newly added, modified, or removed `ns acl6` entries take effect on the Citrix ADC packet-processing engine. Until an apply is performed, ACL6 changes remain in a pending state and are not enforced. Use this resource as the final step in an ACL6 configuration workflow to activate the rules.

This is an **action-only** resource. Applying it triggers the NITRO `apply` action on the `nsacls6` endpoint (the same endpoint also backs the `clear` and `renumber` actions). There is no corresponding ADC object to read back, so this resource performs a one-shot action on create. Read, update, and delete are no-ops, and importing it is not meaningful.


## Example usage

```hcl
resource "citrixadc_nsacls6" "apply_acl6" {
  type = "CLASSIC"
}
```


## Argument Reference

* `type` - (Optional) Type of the ACL6 set to apply. Defaults to `CLASSIC`. Possible values: [ CLASSIC, DFD ]. `CLASSIC` applies the regular extended ACL6 rules; `DFD` applies cluster-specific ACL6 rules that specify the hash method used for steering packets within a cluster. Changing this value forces the action to be re-applied (replacement).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier with the fixed value `nsacls6-config`. Because `nsacls6` is an action-only resource with no NITRO GET endpoint, the ID does not correspond to a readable ADC object.


## Note

This resource models a one-shot action rather than a persistent ADC object. Applying it commits the current ACL6 rule set; the result cannot be read back from the ADC, so subsequent plans will not detect drift. To re-run the apply action (for example after further ACL6 changes), taint the resource or change the `type` argument. Destroying the resource only removes it from Terraform state and does not undo the applied ACL6 changes. Importing this resource is not meaningful.
