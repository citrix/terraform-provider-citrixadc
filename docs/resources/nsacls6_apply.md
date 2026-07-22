---
subcategory: "NS"
---

# Resource: nsacls6_apply

Commits the staged IPv6 extended ACL (ACL6) rule set so that newly added, modified, or removed `ns acl6` entries take effect on the Citrix ADC packet-processing engine. Until an apply is performed, ACL6 changes remain in a pending state and are not enforced. Use this resource as the final step in an ACL6 configuration workflow to activate the rules.

This is an action resource: applying it commits the current ACL6 rule set; it does not manage a persistent object, so re-applying re-runs the action.


## Example usage

### Apply the CLASSIC ACL6 rules

With no arguments, the `CLASSIC` ACL6 rule set is applied (the default).

```hcl
resource "citrixadc_nsacls6_apply" "apply_acl6" {}
```

### Apply a specific ACL6 type

```hcl
resource "citrixadc_nsacls6_apply" "apply_dfd" {
  type = "DFD"
}
```


## Argument Reference

* `type` - (Optional) Type of the ACL6 set to apply. Defaults to `CLASSIC`. Possible values: [ CLASSIC, DFD ]. `CLASSIC` applies the regular extended ACL6 rules; `DFD` applies cluster-specific ACL6 rules that specify the hash method used for steering packets within a cluster. Changing this value re-triggers the apply action (replacement).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsacls6_apply resource. It is set to `nsacls6_apply`.


## Note

To re-run the apply action (for example after further ACL6 changes), taint the resource or change the `type` argument.
