---
subcategory: "Network"
---

# Resource: linkset_channel_binding

The linkset_channel_binding resource is used to create linkset_channel_binding.


## Example usage

```hcl
esource "citrixadc_linkset_channel_binding" "tf_linkset_channel_binding" {
	linkset_id = citrixadc_linkset.tf_linkset.linkset_id
	ifnum      = citrixadc_channel.tf_channel.channel_id
  }
  
  
  resource "citrixadc_linkset" "tf_linkset"{
	  linkset_id = "LS/3"
  }
  
  resource "citrixadc_channel" "tf_channel"{
	  channel_id = "0/LA/2"
  }
```


## Argument Reference

* `linkset_id` - (Required) ID of the linkset to which to bind the interfaces.
* `ifnum` - (Required) The interfaces to be bound to the linkset.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the linkset_channel_binding. It is the concatenation of `linkset_id` and `ifnum` attributes separated by a comma.


## Import

A linkset_channel_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_linkset_channel_binding.tf_linkset_channel_binding LS/3,0/LA/2
```
