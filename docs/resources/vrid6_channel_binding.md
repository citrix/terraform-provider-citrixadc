---
subcategory: "Network"
---

# Resource: vrid6_channel_binding

This resource is used to bind a channel interface to an IPv6 Virtual Router ID (VRID6) on the Citrix ADC.


## Example usage

```hcl
resource "citrixadc_vrid6" "tf_vrid6" {
  id = 100
}

resource "citrixadc_vrid6_channel_binding" "tf_vrid6_channel_binding" {
  vrid_id = citrixadc_vrid6.tf_vrid6.id
  ifnum   = "LA/1"
}
```


## Argument Reference

* `vrid_id` - (Required) Integer value that uniquely identifies a VMAC6 address. This is the ID of the parent VRID6 to which the channel is bound. Minimum value = `1`, Maximum value = `255`. Changing this attribute forces a new resource to be created.
* `ifnum` - (Required) Channel interface to bind to the VMAC6, specified in (slot/port) notation (for example, `LA/1`). Use spaces to separate multiple entries. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vrid6_channel_binding. It is the concatenation of the `vrid_id` and `ifnum` attributes in the form `id:<vrid_id>,ifnum:<ifnum>` (each value URL-encoded).


## Import

A vrid6_channel_binding can be imported using its id (the `id:<vrid_id>,ifnum:<ifnum>` composite key), e.g.

```shell
terraform import citrixadc_vrid6_channel_binding.tf_vrid6_channel_binding "id:100,ifnum:LA%2F1"
```
