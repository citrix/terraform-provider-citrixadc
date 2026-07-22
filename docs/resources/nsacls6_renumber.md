---
subcategory: "NS"
---

# Resource: nsacls6_renumber

Regenerates the priority numbers of the configured IPv6 extended ACL (ACL6) rules on the Citrix ADC, spacing them evenly so that new rules can be inserted between existing ones. Use this resource when the ACL6 priorities have become tightly packed and you need room to add rules at specific positions without manually reassigning priorities.

This is an action resource: applying it renumbers the current ACL6 rule priorities; it does not manage a persistent object, so re-applying re-runs the action.


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

* `id` - The id of the nsacls6_renumber resource. It is set to `nsacls6_renumber`.


## Note

To re-run the renumber action, taint the resource or change the `type` argument.
