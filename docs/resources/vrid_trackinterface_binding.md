---
subcategory: "Network"
---

# Resource: vrid_trackinterface_binding

This resource is used to bind a tracked interface to a VRID.


## Example usage

```hcl
resource "citrixadc_vrid" "tf_vrid" {
  id = 60
}

resource "citrixadc_vrid_trackinterface_binding" "tf_vrid_trackinterface_binding" {
  vrid_id    = citrixadc_vrid.tf_vrid.id
  trackifnum = "1/3"
}
```


## Argument Reference

* `vrid_id` - (Required) Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of `00:00:5e:00:01:<VRID>`. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is `00:00:5e:00:01:3c`, where `3c` is the hexadecimal representation of 60. This is the identifier of the parent `citrixadc_vrid` resource (an integer in the range 1-255). Changing this attribute forces a new resource to be created.
* `trackifnum` - (Required) Interface which needs to be tracked for this VRID, specified in (slot/port) notation (for example, `1/3`). Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vrid_trackinterface_binding. It is the concatenation of the `vrid_id` and `trackifnum` values in the form `id:<vrid_id>,trackifnum:<trackifnum>` (the `trackifnum` value is URL-encoded because interface names contain a `/`).


