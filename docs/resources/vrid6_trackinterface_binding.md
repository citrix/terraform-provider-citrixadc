---
subcategory: "Network"
---

# Resource: vrid6_trackinterface_binding

Binds a tracked interface to an IPv6 Virtual Router ID (VRID6). The state of a tracked interface influences the VRRP priority of the VMAC6: if a tracked interface goes down, the virtual router's effective priority is reduced, which can trigger a failover. Use this resource to make IPv6 VRRP failover decisions depend on the health of specific interfaces on the Citrix ADC.

Creating this resource performs a NITRO bind operation. Both `vrid_id` and `trackifnum` force replacement, so any change recreates the binding (there is no in-place update). Deleting the resource removes the binding from the parent VRID6.


## Example usage

```hcl
resource "citrixadc_vrid6" "tf_vrid6" {
  id = 100
}

resource "citrixadc_vrid6_trackinterface_binding" "tf_vrid6_trackinterface_binding" {
  vrid_id    = citrixadc_vrid6.tf_vrid6.id
  trackifnum = "1/3"
}
```


## Argument Reference

* `vrid_id` - (Required) Integer value that uniquely identifies a VMAC6 address. This is the ID of the parent VRID6 to which the tracked interface is bound. Minimum value = `1`, Maximum value = `255`. Changing this attribute forces a new resource to be created.
* `trackifnum` - (Required) Interface which needs to be tracked for this vrID, specified in (slot/port) notation (for example, `1/3`). Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vrid6_trackinterface_binding. It is the concatenation of the `vrid_id` and `trackifnum` attributes in the form `id:<vrid_id>,trackifnum:<trackifnum>` (each value URL-encoded).
