---
subcategory: "System"
---

# Resource: systemgroup_nspartition_binding

The systemgroup_nspartition_binding resource is used to create systemgroup_nspartition_binding.


## Example usage

```hcl
resource "citrixadc_systemgroup_nspartition_binding" "tf_systemgroup_nspartition_binding" {
  groupname     = citrixadc_systemgroup.tf_systemgroup.groupname
  partitionname = citrixadc_nspartition.tf_nspartition.partitionname
}

resource "citrixadc_systemgroup" "tf_systemgroup" {
  groupname = "tf_systemgroup"
  timeout   = 999
}

resource "citrixadc_nspartition" "tf_nspartition" {
  partitionname = "tf_nspartition"
  maxbandwidth  = 10240
  minbandwidth  = 512
  maxconn       = 512
  maxmemlimit   = 11
}

```


## Argument Reference

* `partitionname` - (Required) Name of the Partition to bind to the system group. Minimum length =  1
* `groupname` - (Required) Name of the system group. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemgroup_nspartition_binding. It has the same value as the `groupname` and `partitionname` attribute.


## Import

A systemgroup_nspartition_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_systemgroup_nspartition_binding.tf_systemgroup_nspartition_binding tf_systemgroup,tf_nspartition
```
