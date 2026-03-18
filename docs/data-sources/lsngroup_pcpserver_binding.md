---
subcategory: "LSN"
---

# Data Source: lsngroup_pcpserver_binding

The lsngroup_pcpserver_binding data source allows you to retrieve information about LSN group PCP server bindings.

## Example Usage

```terraform
data "citrixadc_lsngroup_pcpserver_binding" "tf_lsngroup_pcpserver_binding" {
  groupname = "my_lsn_group"
  pcpserver = "my_pcpserver"
}

output "groupname" {
  value = data.citrixadc_lsngroup_pcpserver_binding.tf_lsngroup_pcpserver_binding.groupname
}

output "pcpserver" {
  value = data.citrixadc_lsngroup_pcpserver_binding.tf_lsngroup_pcpserver_binding.pcpserver
}
```

## Argument Reference

* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn group1" or 'lsn group1').
* `pcpserver` - (Required) Name of the PCP server to be associated with lsn group.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsngroup_pcpserver_binding. It is a system-generated identifier.
