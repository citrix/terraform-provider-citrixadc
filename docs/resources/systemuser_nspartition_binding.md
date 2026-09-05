---
subcategory: "System"
---

# Resource: systemuser_nspartition_binding

The systemuser_nspartition_binding resource is used to bind a system user to an admin partition.


## Example usage

```hcl
resource "citrixadc_systemuser" "tf_user" {
  username = "george"
  password = "tf_password"
  timeout  = 900
}

resource "citrixadc_nspartition" "tf_nspartition" {
  partitionname = "tf_nspartition"
  maxbandwidth  = 10240
  minbandwidth  = 512
  maxconn       = 512
  maxmemlimit   = 11
}

resource "citrixadc_systemuser_nspartition_binding" "tf_bind" {
  username      = citrixadc_systemuser.tf_user.username
  partitionname = citrixadc_nspartition.tf_nspartition.partitionname
}
```


## Argument Reference

* `partitionname` - (Required) Name of the Partition to bind to the system user.
* `username` - (Required) Name of the system-user entry to which to bind the command policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemuser_nspartition_binding. It is the concatenation of  `username` and `partitionname` attributes separated by a comma.


## Import

A systemuser_nspartition_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_systemuser_nspartition_binding.tf_bind george,tf_nspartition
```
