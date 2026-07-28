---
subcategory: "NS"
---

# Resource: nsacls6_renumber

This resource is used to renumber IPv6 extended ACL (ACL6) rule priorities on the Citrix ADC.


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
