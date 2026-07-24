---
subcategory: "NS"
---

# Resource: nsacls6_clear

This resource is used to clear all IPv6 extended ACL (ACL6) rules.


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
