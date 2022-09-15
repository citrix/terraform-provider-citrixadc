---
subcategory: "Lsn"
---

# Resource: lsngroup_lsnpool_binding

The lsngroup_lsnpool_binding resource is used to create lsngroup_lsnpool_binding.


## Example usage

```hcl
resource "citrixadc_lsngroup_lsnpool_binding" "tf_lsngroup_lsnpool_binding" {
  groupname = "my_lsn_group"
  poolname  = "my_pool"
}
```


## Argument Reference

* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn group1" or 'lsn group1').
* `poolname` - (Required) Name of the LSN pool to bind to the specified LSN group. Only LSN Pools and LSN groups with the same NAT type settings can be bound together. Multiples LSN pools can be bound to an LSN group.  For Deterministic NAT, pools bound to an LSN group cannot be bound to other LSN groups. For Dynamic NAT, pools bound to an LSN group can be bound to multiple LSN groups.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsngroup_lsnpool_binding. It is the concatenation of `groupname` and `poolname` attributes separated by a comma.


## Import

A lsngroup_lsnpool_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lsngroup_lsnpool_binding.tf_lsngroup_lsnpool_binding my_lsn_group,my_pool
```
