---
subcategory: "Cluster"
---

# Resource: clustefiles\_syncer

This resource is used to manually trigger the cluster files synchronization
operation.

It is the equivalent of running
```shell
sync cluster files
```
on the command prompt of the Cluster IP address.

By its nature this resource does not have a remote state to read or modify.

Any change in the local state will trigger the operation once more.



## Example usage

```hcl
resource "citrixadc_clusterfiles_syncer" "syncer" {
    timestamp = timestamp()
    mode = ["all", "misc"]
}
```


## Argument Reference

* `timestamp` - (Required) A string representing the current time by convention. Its main purpose is to enable the user to trigger the operation again without having to manually taint the resource.
* `mode` - (Required) A list of values defining which directories and files will be synchronized. Possible values: [ all, bookmarks, ssl, htmlinjection, imports, misc, dns, krb, AAA, app\_catalog, all\_plus\_misc, all\_minus\_misc ]



## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the clusterfiles\_syncer. It has the same value as the `timestamp` attribute.
