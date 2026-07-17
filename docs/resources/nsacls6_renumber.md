---
subcategory: "NS"
---

# Resource: nsacls6_renumber

Regenerates the priority numbers of the configured IPv6 extended ACL (ACL6) rules on the Citrix ADC, spacing them evenly so that new rules can be inserted between existing ones. Use this resource when the ACL6 priorities have become tightly packed and you need room to add rules at specific positions without manually reassigning priorities.

This is an **action-only** resource. Applying it triggers the NITRO `renumber` action on the `nsacls6` endpoint (the same endpoint also backs the `apply` and `clear` actions). There is no corresponding ADC object to read back, so this resource performs a one-shot action on create. Read, update, and delete are no-ops, and importing it is not meaningful. There is no NITRO GET endpoint for this action, so there is no corresponding data source.


## Example usage

### Renumber the CLASSIC ACL6 rules

With no arguments, the `CLASSIC` ACL6 rule set is renumbered (the default).

```hcl
resource "citrixadc_nsacls6_renumber" "renumber_acl6" {}
```

### Renumber a specific ACL6 type

```hcl
resource "citrixadc_nsacls6_renumber" "renumber_dfd" {
  type = "DFD"
}
```


## Argument Reference

* `type` - (Optional) Type of the ACL6 set to renumber. Defaults to `CLASSIC`. Possible values: [ CLASSIC, DFD ]. `CLASSIC` renumbers the regular extended ACL6 rules; `DFD` renumbers cluster-specific ACL6 rules that specify the hash method used for steering packets within a cluster. Changing this value re-triggers the renumber action (replacement).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for this action-only resource. It is a fixed string with the value `nsacls6_renumber`. Because `nsacls6_renumber` performs a NITRO action with no GET endpoint, the ID does not correspond to any readable object on the Citrix ADC.


## Note

This resource models a one-shot action rather than a persistent ADC object. Applying it renumbers the current ACL6 rule priorities; the result cannot be read back from the ADC, so subsequent plans will not detect drift. To re-run the renumber action, taint the resource or change the `type` argument. Destroying the resource only removes it from Terraform state and does not revert the renumbered priorities. Importing this resource is not meaningful.
