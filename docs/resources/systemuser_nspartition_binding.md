---
subcategory: "AAA"
---

# Resource: systemuser_nspartition_binding

The systemuser_nspartition_binding resource is used to create systemuser_nspartition_binding.


## Example usage

```hcl
resource "citrixadc_systemuser_nspartition_binding" "tf_systemuser_nspartition_binding" {
	username      = citrixadc_systemuser.user.username
	partitionname = citrixadc_nspartition.tf_nspartition.partitionname
  }
  
  resource "citrixadc_nspartition" "tf_nspartition" {
	partitionname = "tf_nspartition"
	maxbandwidth  = 10240
	minbandwidth  = 512
	maxconn       = 512
	maxmemlimit   = 11
  }
  
  
  resource "citrixadc_systemuser" "user" {
	username = "george"
	password = "12345"
	timeout  = 900
  }
```


## Argument Reference

* `partitionname` - (Required) Name of the Partition to bind to the system user.
* `username` - (Required) Name of the system-user entry to which to bind the command policy. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the systemuser_nspartition_binding. It is the concatenation of  `username` and `partitionname` attributes separated by a comma.


## Import

A systemuser_nspartition_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_systemuser_nspartition_binding.tf_systemuser_nspartition_binding george,tf_nspartition
```
