---
subcategory: "Network"
---

# Resource: vrid6_interface_binding

Binds a physical interface to an IPv6 Virtual Router ID (VRID6) so that the interface participates in IPv6 VRRP for the VMAC6 identified by the VRID. Use this resource to control which interfaces carry the virtual router's IPv6 traffic on the Citrix ADC.

Creating this resource performs a NITRO bind operation. Both `vrid_id` and `ifnum` force replacement, so any change recreates the binding (there is no in-place update). Deleting the resource removes the binding from the parent VRID6.


## Example usage

```hcl
resource "citrixadc_vrid6" "tf_vrid6" {
  id = 100
}

resource "citrixadc_vrid6_interface_binding" "tf_vrid6_interface_binding" {
  vrid_id = citrixadc_vrid6.tf_vrid6.id
  ifnum   = "1/2"
}
```


## Argument Reference

* `vrid_id` - (Required) Integer value that uniquely identifies a VMAC6 address. This is the ID of the parent VRID6 to which the interface is bound. Minimum value = `1`, Maximum value = `255`. Changing this attribute forces a new resource to be created.
* `ifnum` - (Required) Interface to bind to the VMAC6, specified in (slot/port) notation (for example, `1/2`). Use spaces to separate multiple entries. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vrid6_interface_binding. It is the concatenation of the `vrid_id` and `ifnum` attributes in the form `id:<vrid_id>,ifnum:<ifnum>` (each value URL-encoded).


## Import

A vrid6_interface_binding can be imported using its id (the `id:<vrid_id>,ifnum:<ifnum>` composite key), e.g.

```shell
terraform import citrixadc_vrid6_interface_binding.tf_vrid6_interface_binding "id:100,ifnum:1%2F2"
```
