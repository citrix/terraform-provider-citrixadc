---
subcategory: "Lsn"
---

# Resource: lsngroup_pcpserver_binding

The lsngroup_pcpserver_binding resource is used to create lsngroup_pcpserver_binding.


## Example usage

```hcl
resource "citrixadc_lsngroup_pcpserver_binding" "tf_lsngroup_pcpserver_binding" {
  groupname = "my_lsn_group"
  pcpserver = "my_pcpserver"
}
```


## Argument Reference

* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn group1" or 'lsn group1').
* `pcpserver` - (Required) Name of the PCP server to be associated with lsn group.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsngroup_pcpserver_binding. It is the concatenation of `groupname` and `pcpserver` attributes separated by a comma.


## Import

A lsngroup_pcpserver_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lsngroup_pcpserver_binding.tf_lsngroup_pcpserver_binding my_lsn_group,my_pcpserver
```
