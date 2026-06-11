---
subcategory: "Network"
---

# Resource: fis_channel_binding

Binds an LA (link aggregation) channel to a Failover Interface Set (FIS) on the Citrix ADC. A FIS groups one or more interfaces or channels so that the failover state of the set as a whole follows the state of its members. Use this resource to add a channel (for example, `LA/1`) as a member of an existing FIS created with the `citrixadc_fis` resource.


## Example usage

```hcl
resource "citrixadc_fis" "tf_fis" {
  name = "fis1"
}

resource "citrixadc_fis_channel_binding" "tf_fis_channel_binding" {
  name  = citrixadc_fis.tf_fis.name
  ifnum = "LA/1"
}
```


## Argument Reference

* `name` - (Required) The name of the FIS to which you want to bind the channel. Changing this forces a new resource to be created.
* `ifnum` - (Required) Channel to be bound to the FIS, specified in slot/port notation (for example, `LA/1`). Changing this forces a new resource to be created.
* `ownernode` - (Optional) ID of the cluster node for which you are creating the FIS. Can be configured only through the cluster IP address. Changing this forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the fis_channel_binding. It is a composite key of the form `name:<name>,ifnum:<ifnum>`, where each value is URL-encoded (for example, `name:fis1,ifnum:LA%2F1`).


## Import

A fis_channel_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_fis_channel_binding.tf_fis_channel_binding "name:fis1,ifnum:LA%2F1"
```
