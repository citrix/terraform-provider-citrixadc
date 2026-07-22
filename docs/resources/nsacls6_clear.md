---
subcategory: "NS"
---

# Resource: nsacls6_clear

Removes all configured IPv6 extended ACL (ACL6) rules from the Citrix ADC in one operation. Use this resource to reset the ACL6 rule set — for example, before reloading a fresh set of rules or when tearing down an ACL6 configuration — instead of deleting each `ns acl6` entry individually.

This is an action resource: applying it clears the current ACL6 rule set; it does not manage a persistent object, so re-applying re-runs the action.


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

* `id` - The id of the nsacls6_clear resource. It is set to `nsacls6_clear`.


## Note

To re-run the clear action, taint the resource or change the `type` argument.
